package gormimpl

import (
	"context"
	"errors"

	"gorm.io/gorm"

	nebulaAdminDbErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

// Implementation of WorkflowRepoInterface.
type WorkflowRepo struct {
	db               *gorm.DB
	errorTransformer nebulaAdminDbErrors.ErrorTransformer
	metrics          gormMetrics
}

func (r *WorkflowRepo) Create(ctx context.Context, input models.Workflow, descriptionEntity *models.DescriptionEntity) error {
	timer := r.metrics.CreateDuration.Start()
	err := r.db.WithContext(ctx).Transaction(func(_ *gorm.DB) error {
		if descriptionEntity != nil {
			tx := r.db.WithContext(ctx).Omit("id").Create(descriptionEntity)
			if tx.Error != nil {
				return r.errorTransformer.ToNebulaAdminError(tx.Error)
			}
		}
		tx := r.db.WithContext(ctx).Omit("id").Create(&input)
		if tx.Error != nil {
			return r.errorTransformer.ToNebulaAdminError(tx.Error)
		}

		return nil
	})
	timer.Stop()
	return err
}

func (r *WorkflowRepo) Get(ctx context.Context, input interfaces.Identifier) (models.Workflow, error) {
	var workflow models.Workflow
	timer := r.metrics.GetDuration.Start()
	tx := r.db.WithContext(ctx).Where(&models.Workflow{
		WorkflowKey: models.WorkflowKey{
			Project: input.Project,
			Domain:  input.Domain,
			Name:    input.Name,
			Version: input.Version,
		},
	}).Take(&workflow)
	timer.Stop()

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.Workflow{}, nebulaAdminDbErrors.GetMissingEntityError(core.ResourceType_WORKFLOW.String(), &core.Identifier{
			Project: input.Project,
			Domain:  input.Domain,
			Name:    input.Name,
			Version: input.Version,
		})
	} else if tx.Error != nil {
		return models.Workflow{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return workflow, nil
}

func (r *WorkflowRepo) List(
	ctx context.Context, input interfaces.ListResourceInput) (interfaces.WorkflowCollectionOutput, error) {
	// First validate input.
	if err := ValidateListInput(input); err != nil {
		return interfaces.WorkflowCollectionOutput{}, err
	}
	var workflows []models.Workflow
	tx := r.db.WithContext(ctx).Limit(input.Limit).Offset(input.Offset)

	// Apply filters
	tx, err := applyFilters(tx, input.InlineFilters, input.MapFilters)
	if err != nil {
		return interfaces.WorkflowCollectionOutput{}, err
	}
	// Apply sort ordering.
	if input.SortParameter != nil {
		tx = tx.Order(input.SortParameter.GetGormOrderExpr())
	}
	timer := r.metrics.ListDuration.Start()
	tx.Find(&workflows)
	timer.Stop()
	if tx.Error != nil {
		return interfaces.WorkflowCollectionOutput{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return interfaces.WorkflowCollectionOutput{
		Workflows: workflows,
	}, nil
}

func (r *WorkflowRepo) ListIdentifiers(ctx context.Context, input interfaces.ListResourceInput) (
	interfaces.WorkflowCollectionOutput, error) {

	// Validate input.
	if err := ValidateListInput(input); err != nil {
		return interfaces.WorkflowCollectionOutput{}, err
	}

	tx := r.db.WithContext(ctx).Model(models.Workflow{}).Limit(input.Limit).Offset(input.Offset)

	// Apply filters
	tx, err := applyFilters(tx, input.InlineFilters, input.MapFilters)
	if err != nil {
		return interfaces.WorkflowCollectionOutput{}, err
	}

	// Apply sort ordering.
	if input.SortParameter != nil {
		tx = tx.Order(input.SortParameter.GetGormOrderExpr())
	}

	// Scan the results into a list of workflows
	var workflows []models.Workflow
	timer := r.metrics.ListIdentifiersDuration.Start()
	tx.Select([]string{Project, Domain, Name}).Group(identifierGroupBy).Scan(&workflows)
	timer.Stop()
	if tx.Error != nil {
		return interfaces.WorkflowCollectionOutput{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	return interfaces.WorkflowCollectionOutput{
		Workflows: workflows,
	}, nil
}

// Returns an instance of WorkflowRepoInterface
func NewWorkflowRepo(
	db *gorm.DB, errorTransformer nebulaAdminDbErrors.ErrorTransformer, scope promutils.Scope) interfaces.WorkflowRepoInterface {
	metrics := newMetrics(scope)
	return &WorkflowRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
