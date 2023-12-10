package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Interface for managing Nebula Workflows
type WorkflowInterface interface {
	CreateWorkflow(ctx context.Context, request admin.WorkflowCreateRequest) (*admin.WorkflowCreateResponse, error)
	GetWorkflow(ctx context.Context, request admin.ObjectGetRequest) (*admin.Workflow, error)
	ListWorkflows(ctx context.Context, request admin.ResourceListRequest) (*admin.WorkflowList, error)
	ListWorkflowIdentifiers(ctx context.Context, request admin.NamedEntityIdentifierListRequest) (
		*admin.NamedEntityIdentifierList, error)
}
