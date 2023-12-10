package transformers

import (
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Transforms a ExecutionEventCreateRequest to a ExecutionEvent model
func CreateExecutionEventModel(request admin.WorkflowExecutionEventRequest) (*models.ExecutionEvent, error) {
	occurredAt, err := ptypes.Timestamp(request.Event.OccurredAt)
	if err != nil {
		return nil, errors.NewNebulaAdminErrorf(codes.Internal, "failed to marshal occurred at timestamp")
	}
	return &models.ExecutionEvent{
		ExecutionKey: models.ExecutionKey{
			Project: request.Event.ExecutionId.Project,
			Domain:  request.Event.ExecutionId.Domain,
			Name:    request.Event.ExecutionId.Name,
		},
		RequestID:  request.RequestId,
		OccurredAt: occurredAt,
		Phase:      request.Event.Phase.String(),
	}, nil
}
