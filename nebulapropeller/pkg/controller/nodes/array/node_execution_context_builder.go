package array

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/io"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/executors"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/nodes/interfaces"
)

type arrayNodeExecutionContextBuilder struct {
	nCtxBuilder        interfaces.NodeExecutionContextBuilder
	subNodeID          v1alpha1.NodeID
	subNodeIndex       int
	subNodeStatus      *v1alpha1.NodeStatus
	inputReader        io.InputReader
	currentParallelism *uint32
	maxParallelism     uint32
	eventRecorder      arrayEventRecorder
}

func (a *arrayNodeExecutionContextBuilder) BuildNodeExecutionContext(ctx context.Context, executionContext executors.ExecutionContext,
	nl executors.NodeLookup, currentNodeID v1alpha1.NodeID) (interfaces.NodeExecutionContext, error) {

	// create base NodeExecutionContext
	nCtx, err := a.nCtxBuilder.BuildNodeExecutionContext(ctx, executionContext, nl, currentNodeID)
	if err != nil {
		return nil, err
	}

	if currentNodeID == a.subNodeID {
		// overwrite NodeExecutionContext for ArrayNode execution
		nCtx = newArrayNodeExecutionContext(nCtx, a.inputReader, a.eventRecorder, a.subNodeIndex, a.subNodeStatus, a.currentParallelism, a.maxParallelism)
	}

	return nCtx, nil
}

func newArrayNodeExecutionContextBuilder(nCtxBuilder interfaces.NodeExecutionContextBuilder, subNodeID v1alpha1.NodeID, subNodeIndex int, subNodeStatus *v1alpha1.NodeStatus,
	inputReader io.InputReader, currentParallelism *uint32, maxParallelism uint32, eventRecorder arrayEventRecorder) interfaces.NodeExecutionContextBuilder {

	return &arrayNodeExecutionContextBuilder{
		nCtxBuilder:        nCtxBuilder,
		subNodeID:          subNodeID,
		subNodeIndex:       subNodeIndex,
		subNodeStatus:      subNodeStatus,
		inputReader:        inputReader,
		currentParallelism: currentParallelism,
		maxParallelism:     maxParallelism,
		eventRecorder:      eventRecorder,
	}
}
