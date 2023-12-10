package gormimpl

import (
	"context"

	"gorm.io/gorm"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type ExecutionEventRepo struct {
	db               *gorm.DB
	errorTransformer errors.ErrorTransformer
	metrics          gormMetrics
}

func (r *ExecutionEventRepo) Create(ctx context.Context, input models.ExecutionEvent) error {
	timer := r.metrics.CreateDuration.Start()
	tx := r.db.WithContext(ctx).Omit("id").Create(&input)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

// Returns an instance of ExecutionRepoInterface
func NewExecutionEventRepo(
	db *gorm.DB, errorTransformer errors.ErrorTransformer, scope promutils.Scope) interfaces.ExecutionEventRepoInterface {
	metrics := newMetrics(scope)
	return &ExecutionEventRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
