package executioncluster

import (
	"k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	nebulaclient "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned"
	"github.com/nebulaclouds/nebula/nebulastdlib/random"
)

// Spec to determine the execution target
type ExecutionTargetSpec struct {
	TargetID    string
	ExecutionID string
	Project     string
	Domain      string
	Workflow    string
	LaunchPlan  string
}

// Client object of the target execution cluster
type ExecutionTarget struct {
	ID            string
	NebulaClient   nebulaclient.Interface
	Client        client.Client
	DynamicClient dynamic.Interface
	Enabled       bool
	Config        restclient.Config
}

func (e ExecutionTarget) Compare(to random.Comparable) bool {
	return e.ID < to.(ExecutionTarget).ID
}
