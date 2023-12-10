package impl

import (
	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/workflowengine/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/compiler/common"
)

type workflowCompiler struct {
}

func (c *workflowCompiler) CompileTask(task *core.TaskTemplate) (*core.CompiledTask, error) {
	compiledTask, err := compiler.CompileTask(task)
	if err != nil {
		return nil, errors.NewNebulaAdminErrorf(codes.InvalidArgument, "task failed compilation with %v", err)
	}
	return compiledTask, nil
}

func (c *workflowCompiler) GetRequirements(fg *core.WorkflowTemplate, subWfs []*core.WorkflowTemplate) (
	compiler.WorkflowExecutionRequirements, error) {
	requirements, err := compiler.GetRequirements(fg, subWfs)
	if err != nil {
		return compiler.WorkflowExecutionRequirements{},
			errors.NewNebulaAdminErrorf(codes.InvalidArgument, "failed to get workflow requirements with err %v", err)
	}
	return requirements, nil
}

func (c *workflowCompiler) CompileWorkflow(
	primaryWf *core.WorkflowTemplate, subworkflows []*core.WorkflowTemplate, tasks []*core.CompiledTask,
	launchPlans []common.InterfaceProvider) (*core.CompiledWorkflowClosure, error) {

	compiledWorkflowClosure, err := compiler.CompileWorkflow(primaryWf, subworkflows, tasks, launchPlans)
	if err != nil {
		return nil, errors.NewNebulaAdminErrorf(codes.InvalidArgument, "failed to compile workflow with err %v", err)
	}
	return compiledWorkflowClosure, nil
}

func NewCompiler() interfaces.Compiler {
	return &workflowCompiler{}
}
