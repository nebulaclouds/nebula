package interfaces

import (
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

//go:generate mockery -name=NodeExecutionEventWriter -output=../mocks -case=underscore

type NodeExecutionEventWriter interface {
	Run()
	Write(nodeExecutionEvent admin.NodeExecutionEventRequest)
}
