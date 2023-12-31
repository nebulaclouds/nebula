package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
)

//go:generate mockery -name=ExecutionEventRepoInterface -output=../mocks -case=underscore

type ExecutionEventRepoInterface interface {
	// Inserts a workflow execution event into the database store.
	Create(ctx context.Context, input models.ExecutionEvent) error
}
