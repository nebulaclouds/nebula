package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
)

//go:generate mockery -name=NodeExecutionEventRepoInterface -output=../mocks -case=underscore

type NodeExecutionEventRepoInterface interface {
	// Inserts a node execution event into the database store.
	Create(ctx context.Context, input models.NodeExecutionEvent) error
}
