package interfaces

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

//go:generate mockery -name=WorkflowExecutionEventWriter -output=../mocks -case=underscore

type WorkflowExecutionEventWriter interface {
	Run()
	Write(workflowExecutionEvent admin.WorkflowExecutionEventRequest)
}
