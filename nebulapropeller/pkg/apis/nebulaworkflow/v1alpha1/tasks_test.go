package v1alpha1_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

func TestTaskSpec(t *testing.T) {
	j, err := ReadYamlFileAsJSON("testdata/task.yaml")
	assert.NoError(t, err)

	task := &v1alpha1.TaskSpec{}
	assert.NoError(t, json.Unmarshal(j, task))

	assert.NotNil(t, task.CoreTask())
	assert.Equal(t, "demo", task.TaskType())
}
