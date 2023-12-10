package gormimpl

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

	nebulaAdminErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	nebulaAdminDbErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type ProjectRepo struct {
	db               *gorm.DB
	errorTransformer nebulaAdminDbErrors.ErrorTransformer
	metrics          gormMetrics
}

func (r *ProjectRepo) Create(ctx context.Context, project models.Project) error {
	timer := r.metrics.CreateDuration.Start()
	tx := r.db.WithContext(ctx).Omit("id").Create(&project)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

func (r *ProjectRepo) Get(ctx context.Context, projectID string) (models.Project, error) {
	var project models.Project
	timer := r.metrics.GetDuration.Start()
	tx := r.db.WithContext(ctx).Where(&models.Project{
		Identifier: projectID,
	}).Take(&project)
	timer.Stop()
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.Project{}, nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound, "project [%s] not found", projectID)
	}

	if tx.Error != nil {
		return models.Project{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	return project, nil
}

func (r *ProjectRepo) List(ctx context.Context, input interfaces.ListResourceInput) ([]models.Project, error) {
	var projects []models.Project

	tx := r.db.WithContext(ctx).Offset(input.Offset)
	if input.Limit != 0 {
		tx = tx.Limit(input.Limit)
	}

	// Apply filters
	// If no filter provided, default to filtering out archived projects
	if len(input.InlineFilters) == 0 && len(input.MapFilters) == 0 {
		tx = tx.Where("state != ?", int32(admin.Project_ARCHIVED))
	} else {
		var err error
		tx, err = applyFilters(tx, input.InlineFilters, input.MapFilters)
		if err != nil {
			return nil, err
		}
	}

	// Apply sort ordering
	if input.SortParameter != nil {
		tx = tx.Order(input.SortParameter.GetGormOrderExpr())
	}

	timer := r.metrics.ListDuration.Start()
	tx.Find(&projects)
	timer.Stop()

	if tx.Error != nil {
		return nil, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return projects, nil
}

func NewProjectRepo(db *gorm.DB, errorTransformer nebulaAdminDbErrors.ErrorTransformer,
	scope promutils.Scope) interfaces.ProjectRepoInterface {
	metrics := newMetrics(scope)
	return &ProjectRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}

func (r *ProjectRepo) UpdateProject(ctx context.Context, projectUpdate models.Project) error {
	// Use gorm client to update the two fields that are changed.
	writeTx := r.db.WithContext(ctx).Model(&projectUpdate).Updates(projectUpdate)

	// Return error if applies.
	if writeTx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(writeTx.Error)
	}

	return nil
}
