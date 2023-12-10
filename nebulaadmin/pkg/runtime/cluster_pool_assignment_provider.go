package runtime

import (
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulastdlib/config"
)

const clusterPoolsKey = "clusterPools"

var clusterPoolsConfig = config.MustRegisterSection(clusterPoolsKey, &interfaces.ClusterPoolAssignmentConfig{
	ClusterPoolAssignments: make(interfaces.ClusterPoolAssignments),
})

// Implementation of an interfaces.ClusterPoolAssignmentConfiguration
type ClusterPoolAssignmentConfigurationProvider struct{}

func (p *ClusterPoolAssignmentConfigurationProvider) GetClusterPoolAssignments() interfaces.ClusterPoolAssignments {
	return clusterPoolsConfig.GetConfig().(*interfaces.ClusterPoolAssignmentConfig).ClusterPoolAssignments
}

func NewClusterPoolAssignmentConfigurationProvider() interfaces.ClusterPoolAssignmentConfiguration {
	return &ClusterPoolAssignmentConfigurationProvider{}
}
