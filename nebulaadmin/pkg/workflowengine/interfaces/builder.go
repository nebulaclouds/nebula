package interfaces

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

//go:generate mockery -name NebulaWorkflowBuilder -output=../mocks -case=underscore

// NebulaWorkflowBuilder produces a v1alpha1.NebulaWorkflow definition from a compiled workflow closure and execution inputs
type NebulaWorkflowBuilder interface {
	Build(
		wfClosure *core.CompiledWorkflowClosure, inputs *core.LiteralMap, executionID *core.WorkflowExecutionIdentifier,
		namespace string) (*v1alpha1.NebulaWorkflow, error)
}
