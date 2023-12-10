package interfaces

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler/common"
)

// Workflow compiler interface.
type Compiler interface {
	CompileTask(task *core.TaskTemplate) (*core.CompiledTask, error)
	GetRequirements(fg *core.WorkflowTemplate, subWfs []*core.WorkflowTemplate) (
		compiler.WorkflowExecutionRequirements, error)
	CompileWorkflow(primaryWf *core.WorkflowTemplate, subworkflows []*core.WorkflowTemplate, tasks []*core.CompiledTask,
		launchPlans []common.InterfaceProvider) (*core.CompiledWorkflowClosure, error)
}
