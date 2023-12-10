package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

//go:generate mockery -name=MetricsInterface -output=../mocks -case=underscore

// Interface for managing Nebula execution metrics
type MetricsInterface interface {
	GetExecutionMetrics(ctx context.Context, request admin.WorkflowExecutionGetMetricsRequest) (
		*admin.WorkflowExecutionGetMetricsResponse, error)
}
