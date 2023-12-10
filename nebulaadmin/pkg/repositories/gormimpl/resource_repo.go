package gormimpl

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"gorm.io/gorm"

	nebulaAdminErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	nebulaAdminDbErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type ResourceRepo struct {
	db               *gorm.DB
	errorTransformer nebulaAdminDbErrors.ErrorTransformer
	metrics          gormMetrics
}

const priorityDescending = "priority desc"

/*
The data in the Resource repo maps to the following rules:
* Domain and ResourceType can never be empty.
* Empty string can be interpreted as all. Example: "" for Project field can be interpreted as all Projects for a domain.
* One cannot provide specific value for Project, unless a specific value for Domain is provided.
** Project is always scoped within a domain.
**	Example: Domain="" Project="Lyft" is invalid.
* One cannot provide specific value for Workflow, unless a specific value for Domain and Project is provided.
** Workflow is always scoped within a domain and project.
**	Example: Domain="staging" Project="" Workflow="W1" is invalid.
* One cannot provide specific value for Launch plan, unless a specific value for Domain, Project and Workflow is provided.
** Launch plan is always scoped within a domain, project and workflow.
**	Example: Domain="staging" Project="Lyft" Workflow="" LaunchPlan= "l1" is invalid.
*/
func validateCreateOrUpdateResourceInput(project, domain, workflow, launchPlan, resourceType string) bool {
	if resourceType == "" {
		return false
	}
	if domain == "" && project == "" {
		return false
	}
	if project == "" && (workflow != "" || launchPlan != "") {
		return false
	}
	if workflow == "" && launchPlan != "" {
		return false
	}
	return true
}

func (r *ResourceRepo) CreateOrUpdate(ctx context.Context, input models.Resource) error {
	if !validateCreateOrUpdateResourceInput(input.Project, input.Domain, input.Workflow, input.LaunchPlan, input.ResourceType) {
		return nebulaAdminDbErrors.GetInvalidInputError(fmt.Sprintf("%v", input))
	}
	if input.Priority == 0 {
		return nebulaAdminDbErrors.GetInvalidInputError(fmt.Sprintf("invalid priority %v", input))
	}
	timer := r.metrics.GetDuration.Start()
	var record models.Resource
	tx := r.db.WithContext(ctx).FirstOrCreate(&record, models.Resource{
		Project:      input.Project,
		Domain:       input.Domain,
		Workflow:     input.Workflow,
		LaunchPlan:   input.LaunchPlan,
		ResourceType: input.ResourceType,
		Priority:     input.Priority,
	})
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	timer = r.metrics.UpdateDuration.Start()
	record.Attributes = input.Attributes
	tx = r.db.WithContext(ctx).Save(&record)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

// Get returns the most-specific attribute setting for the given ResourceType.
func (r *ResourceRepo) Get(ctx context.Context, ID interfaces.ResourceID) (models.Resource, error) {
	if !validateCreateOrUpdateResourceInput(ID.Project, ID.Domain, ID.Workflow, ID.LaunchPlan, ID.ResourceType) {
		return models.Resource{}, r.errorTransformer.ToNebulaAdminError(nebulaAdminDbErrors.GetInvalidInputError(fmt.Sprintf("%v", ID)))
	}
	var resources []models.Resource
	timer := r.metrics.GetDuration.Start()

	txWhereClause := "resource_type = ? AND domain IN (?) AND project IN (?) AND workflow IN (?) AND launch_plan IN (?)"
	project := []string{""}
	if ID.Project != "" {
		project = append(project, ID.Project)
	}

	domain := []string{""}
	if ID.Domain != "" {
		domain = append(domain, ID.Domain)
	}

	workflow := []string{""}
	if ID.Workflow != "" {
		workflow = append(workflow, ID.Workflow)
	}

	launchPlan := []string{""}
	if ID.LaunchPlan != "" {
		launchPlan = append(launchPlan, ID.LaunchPlan)
	}

	tx := r.db.WithContext(ctx).Where(txWhereClause, ID.ResourceType, domain, project, workflow, launchPlan)
	tx.Order(priorityDescending).First(&resources)
	timer.Stop()

	if (tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound)) || len(resources) == 0 {
		return models.Resource{}, nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound,
			"Resource [%+v] not found", ID)
	} else if tx.Error != nil {
		return models.Resource{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return resources[0], nil
}

// GetProjectLevel differs from Get in that it returns only the project-level attribute setting for the
// given ResourceType if it exists. The reason this exists is because we want to return project level
// attributes to Nebula Console, regardless of whether a more specific setting exists.
func (r *ResourceRepo) GetProjectLevel(ctx context.Context, ID interfaces.ResourceID) (models.Resource, error) {
	if ID.Project == "" {
		return models.Resource{}, r.errorTransformer.ToNebulaAdminError(nebulaAdminDbErrors.GetInvalidInputError(fmt.Sprintf("%v", ID)))
	}

	var resources []models.Resource
	timer := r.metrics.GetDuration.Start()

	txWhereClause := "resource_type = ? AND domain = '' AND project = ? AND workflow = '' AND launch_plan = ''"

	tx := r.db.WithContext(ctx).Where(txWhereClause, ID.ResourceType, ID.Project)
	tx.Order(priorityDescending).First(&resources)
	timer.Stop()

	if (tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound)) || len(resources) == 0 {
		return models.Resource{}, nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound,
			"Resource [%+v] not found", ID)
	} else if tx.Error != nil {
		return models.Resource{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return resources[0], nil
}

func (r *ResourceRepo) GetRaw(ctx context.Context, ID interfaces.ResourceID) (models.Resource, error) {
	if ID.Domain == "" || ID.ResourceType == "" {
		return models.Resource{}, r.errorTransformer.ToNebulaAdminError(nebulaAdminDbErrors.GetInvalidInputError(fmt.Sprintf("%v", ID)))
	}
	var model models.Resource
	timer := r.metrics.GetDuration.Start()
	tx := r.db.WithContext(ctx).Where(&models.Resource{
		Project:      ID.Project,
		Domain:       ID.Domain,
		Workflow:     ID.Workflow,
		LaunchPlan:   ID.LaunchPlan,
		ResourceType: ID.ResourceType,
	}).First(&model)
	timer.Stop()

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.Resource{}, nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound,
			"%v", ID)
	} else if tx.Error != nil {
		return models.Resource{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return model, nil
}

func (r *ResourceRepo) ListAll(ctx context.Context, resourceType string) ([]models.Resource, error) {
	var resources []models.Resource
	timer := r.metrics.ListDuration.Start()

	tx := r.db.WithContext(ctx).Where(&models.Resource{ResourceType: resourceType}).Order(priorityDescending).Find(&resources)
	timer.Stop()

	if tx.Error != nil {
		return nil, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return resources, nil
}

func (r *ResourceRepo) Delete(ctx context.Context, ID interfaces.ResourceID) error {
	var tx *gorm.DB
	r.metrics.DeleteDuration.Time(func() {
		tx = r.db.WithContext(ctx).Where(&models.Resource{
			Project:      ID.Project,
			Domain:       ID.Domain,
			Workflow:     ID.Workflow,
			LaunchPlan:   ID.LaunchPlan,
			ResourceType: ID.ResourceType,
		}).Unscoped().Delete(models.Resource{})
	})

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound, "%v", ID)
	} else if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

func NewResourceRepo(db *gorm.DB, errorTransformer nebulaAdminDbErrors.ErrorTransformer,
	scope promutils.Scope) interfaces.ResourceRepoInterface {
	metrics := newMetrics(scope)
	return &ResourceRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
