package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
)

//go:generate mockery -name=PartitionRepo -output=../mocks -case=underscore

type PartitionRepo interface {
	Create(ctx context.Context, in models.Partition) error
}
