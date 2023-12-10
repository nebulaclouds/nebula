package implementations

import (
	"testing"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	event2 "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/event"
)

func TestWorkflowExecutionEventWriter(t *testing.T) {
	db := mocks.NewMockRepository()

	event := admin.WorkflowExecutionEventRequest{
		RequestId: "request_id",
		Event: &event2.WorkflowExecutionEvent{
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Project: "project",
				Domain:  "domain",
				Name:    "exec_name",
			},
		},
	}

	workflowExecEventRepo := mocks.ExecutionEventRepoInterface{}
	workflowExecEventRepo.On("Create", event).Return(nil)
	db.(*mocks.MockRepository).ExecutionEventRepoIface = &workflowExecEventRepo
	writer := NewWorkflowExecutionEventWriter(db, 100)
	// Assert we can write an event using the buffered channel without holding up this process.
	writer.Write(event)
	go func() { writer.Run() }()
	close(writer.(*workflowExecutionEventWriter).events)
}
