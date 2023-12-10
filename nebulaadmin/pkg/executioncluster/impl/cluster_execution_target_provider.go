package impl

import (
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/nebulak8s"
	runtime "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	nebulaclient "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned"
)

type clusterExecutionTargetProvider struct{}

// Creates a new Execution target for a cluster based on config passed in.
func (c *clusterExecutionTargetProvider) GetExecutionTarget(initializationErrorCounter prometheus.Counter, k8sCluster runtime.ClusterConfig) (*executioncluster.ExecutionTarget, error) {
	kubeConf, err := nebulak8s.GetRestClientConfig("", "", &k8sCluster)
	if err != nil {
		return nil, err
	}
	nebulaClient, err := getRestClientFromKubeConfig(initializationErrorCounter, kubeConf)
	if err != nil {
		return nil, err
	}
	client, err := client.New(kubeConf, client.Options{})
	if err != nil {
		return nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(kubeConf)
	if err != nil {
		return nil, err
	}
	return &executioncluster.ExecutionTarget{
		NebulaClient:   nebulaClient,
		Client:        client,
		DynamicClient: dynamicClient,
		ID:            k8sCluster.Name,
		Enabled:       k8sCluster.Enabled,
		Config:        *kubeConf,
	}, nil
}

func getRestClientFromKubeConfig(initializationErrorCounter prometheus.Counter, kubeConfiguration *rest.Config) (*nebulaclient.Clientset, error) {
	fc, err := nebulaclient.NewForConfig(kubeConfiguration)
	if err != nil {
		initializationErrorCounter.Inc()
		return nil, err
	}
	return fc, nil
}

func NewExecutionTargetProvider() interfaces.ExecutionTargetProvider {
	return &clusterExecutionTargetProvider{}
}
