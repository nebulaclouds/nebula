package nodes

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

type taskReader struct {
	*core.TaskTemplate
}

func (t taskReader) GetTaskType() v1alpha1.TaskType {
	return t.TaskTemplate.Type
}

func (t taskReader) GetTaskID() *core.Identifier {
	return t.Id
}

func (t taskReader) Read(ctx context.Context) (*core.TaskTemplate, error) {
	return t.TaskTemplate, nil
}
