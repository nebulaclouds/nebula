package gormimpl

import (
	"context"
	"errors"

	"gorm.io/gorm"

	nebulaSchedulerDbErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	interfaces2 "github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/scheduler/repositories/models"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

// ScheduleEntitiesSnapshotRepo Implementation of ScheduleEntitiesSnapshotRepoInterface.
type ScheduleEntitiesSnapshotRepo struct {
	db               *gorm.DB
	errorTransformer nebulaSchedulerDbErrors.ErrorTransformer
	metrics          gormMetrics
}

// TODO : always overwrite the existing snapshot instead of creating new rows
func (r *ScheduleEntitiesSnapshotRepo) Write(ctx context.Context, input models.ScheduleEntitiesSnapshot) error {
	timer := r.metrics.GetDuration.Start()
	tx := r.db.Omit("id").Create(&input)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

func (r *ScheduleEntitiesSnapshotRepo) Read(ctx context.Context) (models.ScheduleEntitiesSnapshot, error) {
	var schedulableEntitiesSnapshot models.ScheduleEntitiesSnapshot
	timer := r.metrics.GetDuration.Start()
	tx := r.db.Last(&schedulableEntitiesSnapshot)
	timer.Stop()

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return models.ScheduleEntitiesSnapshot{}, nebulaSchedulerDbErrors.GetSingletonMissingEntityError("schedule_entities_snapshots")
		}
		return models.ScheduleEntitiesSnapshot{}, r.errorTransformer.ToNebulaAdminError(tx.Error)
	}

	return schedulableEntitiesSnapshot, nil
}

// NewScheduleEntitiesSnapshotRepo Returns an instance of ScheduleEntitiesSnapshotRepoInterface
func NewScheduleEntitiesSnapshotRepo(
	db *gorm.DB, errorTransformer nebulaSchedulerDbErrors.ErrorTransformer, scope promutils.Scope) interfaces2.ScheduleEntitiesSnapShotRepoInterface {
	metrics := newMetrics(scope)
	return &ScheduleEntitiesSnapshotRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
