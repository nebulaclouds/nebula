package gormimpl

import (
	"context"

	"gorm.io/gorm"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

type NodeExecutionEventRepo struct {
	db               *gorm.DB
	errorTransformer errors.ErrorTransformer
	metrics          gormMetrics
}

func (r *NodeExecutionEventRepo) Create(ctx context.Context, input models.NodeExecutionEvent) error {
	timer := r.metrics.CreateDuration.Start()
	tx := r.db.WithContext(ctx).Omit("id").Create(&input)
	timer.Stop()
	if tx.Error != nil {
		return r.errorTransformer.ToNebulaAdminError(tx.Error)
	}
	return nil
}

// Returns an instance of NodeExecutionRepoInterface
func NewNodeExecutionEventRepo(
	db *gorm.DB, errorTransformer errors.ErrorTransformer, scope promutils.Scope) interfaces.NodeExecutionEventRepoInterface {
	metrics := newMetrics(scope)
	return &NodeExecutionEventRepo{
		db:               db,
		errorTransformer: errorTransformer,
		metrics:          metrics,
	}
}
