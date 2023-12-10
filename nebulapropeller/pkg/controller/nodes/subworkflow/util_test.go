package subworkflow

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
)

func TestGetChildWorkflowExecutionID(t *testing.T) {
	id, err := GetChildWorkflowExecutionID(
		&core.NodeExecutionIdentifier{
			NodeId: "hello-world",
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Project: "project",
				Domain:  "domain",
				Name:    "first-name-is-pretty-large",
			},
		},
		1)
	assert.Equal(t, id.Name, "fav2uxxi")
	assert.NoError(t, err)
}
