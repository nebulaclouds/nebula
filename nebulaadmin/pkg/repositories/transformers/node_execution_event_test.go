package transformers

import (
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/event"
)

func TestCreateNodeExecutionEventModel(t *testing.T) {
	occurredAt := time.Now().UTC()
	occurredAtProto, _ := ptypes.TimestampProto(occurredAt)
	request := admin.NodeExecutionEventRequest{
		RequestId: "request id",
		Event: &event.NodeExecutionEvent{
			Id: &core.NodeExecutionIdentifier{
				NodeId: "nodey",
				ExecutionId: &core.WorkflowExecutionIdentifier{
					Project: "project",
					Domain:  "domain",
					Name:    "name",
				},
			},
			Phase:      core.NodeExecution_ABORTED,
			OccurredAt: occurredAtProto,
		},
	}

	nodeExecutionEventModel, err := CreateNodeExecutionEventModel(request)
	assert.Nil(t, err)
	assert.Equal(t, &models.NodeExecutionEvent{
		NodeExecutionKey: models.NodeExecutionKey{
			NodeID: "nodey",
			ExecutionKey: models.ExecutionKey{
				Project: "project",
				Domain:  "domain",
				Name:    "name",
			},
		},
		RequestID:  "request id",
		OccurredAt: occurredAt,
		Phase:      "ABORTED",
	}, nodeExecutionEventModel)
}
