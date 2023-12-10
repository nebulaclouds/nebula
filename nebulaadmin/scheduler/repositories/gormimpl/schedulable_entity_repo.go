package gormimpl

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	adminErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

// SchedulableEntityRepo Implementation of SchedulableEntityRepoInterface.
type SchedulableEntityRepo struct {
	db               *gorm.DB
	errorTransformer adminErrors.ErrorTransformer
	metrics          gormMetrics
}

func (r *SchedulableEntityRepo) Create(ctx context.Context, input models.SchedulableEntity) error {
	timer := r.metrics.GetDuration.Start()
	var record models.SchedulableEntity
	tx := r.db.Omit("id").FirstOrCreate(&record, input)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

func (r *SchedulableEntityRepo) Activate(ctx context.Context, input models.SchedulableEntity) error {
	var schedulableEntity models.SchedulableEntity
	timer := r.metrics.GetDuration.Start()
	// Find the existence of a scheduled entity
	tx := r.db.Where(&models.SchedulableEntity{
		SchedulableEntityKey: models.SchedulableEntityKey{
			Project: input.Project,
			Domain:  input.Domain,
			Name:    input.Name,
			Version: input.Version,
		},
	}).Take(&schedulableEntity)
	timer.Stop()

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			// Not found and hence create one
			return r.Create(ctx, input)
		}
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	// Activate the already existing schedule
	return activateOrDeactivate(r, input.SchedulableEntityKey, true)
}

func (r *SchedulableEntityRepo) Deactivate(ctx context.Context, ID models.SchedulableEntityKey) error {
	// Activate the schedule
	return activateOrDeactivate(r, ID, false)
}

func (r *SchedulableEntityRepo) GetAll(ctx context.Context) ([]models.SchedulableEntity, error) {
	var schedulableEntities []models.SchedulableEntity
	timer := r.metrics.GetDuration.Start()

	tx := r.db.Find(&schedulableEntities)

	timer.Stop()

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil,
				fmt.Errorf("no active schedulable entities found")
		}
		return nil, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	return schedulableEntities, nil
}

func (r *SchedulableEntityRepo) Get(ctx context.Context, ID models.SchedulableEntityKey) (models.SchedulableEntity, error) {
	var schedulableEntity models.SchedulableEntity
	timer := r.metrics.GetDuration.Start()
	tx := r.db.Where(&models.SchedulableEntity{
		SchedulableEntityKey: models.SchedulableEntityKey{
			Project: ID.Project,
			Domain:  ID.Domain,
			Name:    ID.Name,
			Version: ID.Version,
		},
	}).Take(&schedulableEntity)
	timer.Stop()

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return models.SchedulableEntity{},
				adminErrors.GetMissingEntityError("schedulable entity", &core.Identifier{
					Project: ID.Project,
					Domain:  ID.Domain,
					Name:    ID.Name,
					Version: ID.Version,
				})
		}
		return models.SchedulableEntity{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	return schedulableEntity, nil
}

// Helper function to activate and deactivate a schedule
func activateOrDeactivate(r *SchedulableEntityRepo, ID models.SchedulableEntityKey, activate bool) error {
	timer := r.metrics.GetDuration.Start()
	tx := r.db.Model(&models.SchedulableEntity{}).Where(&models.SchedulableEntity{
		SchedulableEntityKey: models.SchedulableEntityKey{
			Project: ID.Project,
			Domain:  ID.Domain,
			Name:    ID.Name,
			Version: ID.Version,
		},
	}).Update("active", activate)
	timer.Stop()
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return adminErrors.GetMissingEntityError("schedulable entity", &core.Identifier{
				Project: ID.Project,
				Domain:  ID.Domain,
				Name:    ID.Name,
				Version: ID.Version,
			})
		}
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

// NewSchedulableEntityRepo Returns an instance of SchedulableEntityRepoInterface
func NewSchedulableEntityRepo(
	db *gorm.DB, errorTransformer adminErrors.ErrorTransformer, scope promutils.Scope) interfaces.SchedulableEntityRepoInterface {
	metrics := newMetrics(scope)
	return &SchedulableEntityRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
