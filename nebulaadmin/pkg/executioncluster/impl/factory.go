package impl

import (
	executioncluster_interface "github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster/interfaces"
	repositoryInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

func GetExecutionCluster(scope promutils.Scope, kubeConfig, master string, config interfaces.Configuration, db repositoryInterfaces.Repository) executioncluster_interface.ClusterInterface {
	initializationErrorCounter := scope.MustNewCounter(
		"nebulaclient_initialization_error",
		"count of errors encountered initializing a nebula client from kube config")
	switch len(config.ClusterConfiguration().GetClusterConfigs()) {
	case 0:
		cluster, err := NewInCluster(initializationErrorCounter, kubeConfig, master)
		if err != nil {
			panic(err)
		}
		return cluster
	default:
		listTargetsProvider, err := NewListTargets(initializationErrorCounter, NewExecutionTargetProvider(), config.ClusterConfiguration())
		if err != nil {
			panic(err)
		}
		cluster, err := NewRandomClusterSelector(listTargetsProvider, config, db)
		if err != nil {
			panic(err)
		}
		return cluster
	}
}
