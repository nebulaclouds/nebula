package executors

import (
	"context"
	"testing"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1/mocks"
	"github.com/stretchr/testify/assert"
)

type nl struct {
	NodeLookup
}

type en struct {
	v1alpha1.ExecutableNode
}

type ns struct {
	v1alpha1.ExecutableNodeStatus
}

func TestNewFailureNodeLookup(t *testing.T) {
	nl := nl{}
	en := en{}
	ns := ns{}
	nodeLoopUp := NewFailureNodeLookup(nl, en, ns)
	assert.NotNil(t, nl)
	typed := nodeLoopUp.(FailureNodeLookup)
	assert.Equal(t, nl, typed.NodeLookup)
	assert.Equal(t, en, typed.FailureNode)
	assert.Equal(t, ns, typed.FailureNodeStatus)
}

func TestNewTestFailureNodeLookup(t *testing.T) {
	n := &mocks.ExecutableNode{}
	ns := &mocks.ExecutableNodeStatus{}
	failureNodeID := "fn1"
	nl := NewTestNodeLookup(
		map[string]v1alpha1.ExecutableNode{v1alpha1.StartNodeID: n, failureNodeID: n},
		map[string]v1alpha1.ExecutableNodeStatus{v1alpha1.StartNodeID: ns, failureNodeID: ns},
	)

	assert.NotNil(t, nl)

	failureNodeLookup := NewFailureNodeLookup(nl, n, ns)
	r, ok := failureNodeLookup.GetNode(v1alpha1.StartNodeID)
	assert.True(t, ok)
	assert.Equal(t, n, r)
	assert.Equal(t, ns, failureNodeLookup.GetNodeExecutionStatus(context.TODO(), v1alpha1.StartNodeID))

	r, ok = failureNodeLookup.GetNode(failureNodeID)
	assert.True(t, ok)
	assert.Equal(t, n, r)
	assert.Equal(t, ns, failureNodeLookup.GetNodeExecutionStatus(context.TODO(), failureNodeID))

	nodeIDs, err := failureNodeLookup.ToNode(failureNodeID)
	assert.Equal(t, len(nodeIDs), 1)
	assert.Equal(t, nodeIDs[0], v1alpha1.StartNodeID)
	assert.Nil(t, err)

	nodeIDs, err = failureNodeLookup.FromNode(failureNodeID)
	assert.Nil(t, nodeIDs)
	assert.Nil(t, err)
}
