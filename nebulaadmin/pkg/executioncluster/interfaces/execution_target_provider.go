package interfaces

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
)

//go:generate mockery -all -case=underscore -output=../mocks -case=underscore

type ExecutionTargetProvider interface {
	GetExecutionTarget(initializationErrorCounter prometheus.Counter, k8sCluster interfaces.ClusterConfig) (*executioncluster.ExecutionTarget, error)
}
