package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Interface for managing Nebula Workflow NodeExecutions
type NodeExecutionInterface interface {
	CreateNodeEvent(ctx context.Context, request admin.NodeExecutionEventRequest) (
		*admin.NodeExecutionEventResponse, error)
	GetNodeExecution(ctx context.Context, request admin.NodeExecutionGetRequest) (*admin.NodeExecution, error)
	ListNodeExecutions(ctx context.Context, request admin.NodeExecutionListRequest) (*admin.NodeExecutionList, error)
	ListNodeExecutionsForTask(ctx context.Context, request admin.NodeExecutionForTaskListRequest) (*admin.NodeExecutionList, error)
	GetNodeExecutionData(
		ctx context.Context, request admin.NodeExecutionGetDataRequest) (*admin.NodeExecutionGetDataResponse, error)
}
