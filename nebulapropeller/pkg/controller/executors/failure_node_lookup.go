package executors

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

type FailureNodeLookup struct {
	NodeLookup
	FailureNode       v1alpha1.ExecutableNode
	FailureNodeStatus v1alpha1.ExecutableNodeStatus
}

func (f FailureNodeLookup) GetNode(nodeID v1alpha1.NodeID) (v1alpha1.ExecutableNode, bool) {
	if nodeID == v1alpha1.StartNodeID {
		return f.NodeLookup.GetNode(nodeID)
	}
	return f.FailureNode, true
}

func (f FailureNodeLookup) GetNodeExecutionStatus(ctx context.Context, id v1alpha1.NodeID) v1alpha1.ExecutableNodeStatus {
	if id == v1alpha1.StartNodeID {
		return f.NodeLookup.GetNodeExecutionStatus(ctx, id)
	}
	return f.FailureNodeStatus
}

func (f FailureNodeLookup) ToNode(id v1alpha1.NodeID) ([]v1alpha1.NodeID, error) {
	// The upstream node of the failure node is always the start node
	return []v1alpha1.NodeID{v1alpha1.StartNodeID}, nil
}

func (f FailureNodeLookup) FromNode(id v1alpha1.NodeID) ([]v1alpha1.NodeID, error) {
	return nil, nil
}

func NewFailureNodeLookup(nodeLookup NodeLookup, failureNode v1alpha1.ExecutableNode, failureNodeStatus v1alpha1.ExecutableNodeStatus) NodeLookup {
	return FailureNodeLookup{
		NodeLookup:        nodeLookup,
		FailureNode:       failureNode,
		FailureNodeStatus: failureNodeStatus,
	}
}
