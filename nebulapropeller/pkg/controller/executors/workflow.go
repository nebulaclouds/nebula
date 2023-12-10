package executors

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

//go:generate mockery -name Workflow -case=underscore

type Workflow interface {
	Initialize(ctx context.Context) error
	HandleNebulaWorkflow(ctx context.Context, w *v1alpha1.NebulaWorkflow) error
	HandleAbortedWorkflow(ctx context.Context, w *v1alpha1.NebulaWorkflow, maxRetries uint32) error
}
