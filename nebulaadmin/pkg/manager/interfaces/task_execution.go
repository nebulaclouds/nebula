package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Interface for managing Nebula Workflow TaskExecutions
type TaskExecutionInterface interface {
	CreateTaskExecutionEvent(ctx context.Context, request admin.TaskExecutionEventRequest) (
		*admin.TaskExecutionEventResponse, error)
	GetTaskExecution(ctx context.Context, request admin.TaskExecutionGetRequest) (*admin.TaskExecution, error)
	ListTaskExecutions(ctx context.Context, request admin.TaskExecutionListRequest) (*admin.TaskExecutionList, error)
	GetTaskExecutionData(
		ctx context.Context, request admin.TaskExecutionGetDataRequest) (*admin.TaskExecutionGetDataResponse, error)
}
