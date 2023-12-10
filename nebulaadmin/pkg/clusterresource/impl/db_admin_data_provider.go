package impl

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/clusterresource/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/common"
	managerInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/interfaces"
	repositoryInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/transformers"
	runtimeInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Implementation of an interfaces.NebulaAdminDataProvider which fetches data directly from the provided database connection.
type dbAdminProvider struct {
	db              repositoryInterfaces.Repository
	config          runtimeInterfaces.Configuration
	resourceManager managerInterfaces.ResourceInterface
}

func (p dbAdminProvider) GetClusterResourceAttributes(ctx context.Context, project, domain string) (*admin.ClusterResourceAttributes, error) {
	resource, err := p.resourceManager.GetResource(ctx, managerInterfaces.ResourceRequest{
		Project:      project,
		Domain:       domain,
		ResourceType: admin.MatchableResource_CLUSTER_RESOURCE,
	})
	if err != nil {
		return nil, err
	}
	if resource != nil && resource.Attributes != nil && resource.Attributes.GetClusterResourceAttributes() != nil {
		return resource.Attributes.GetClusterResourceAttributes(), nil
	}
	return nil, NewMissingEntityError("cluster resource attributes")
}

func (p dbAdminProvider) getDomains() []*admin.Domain {
	configDomains := p.config.ApplicationConfiguration().GetDomainsConfig()
	var domains = make([]*admin.Domain, len(*configDomains))
	for index, configDomain := range *configDomains {
		domains[index] = &admin.Domain{
			Id:   configDomain.ID,
			Name: configDomain.Name,
		}
	}
	return domains
}

func (p dbAdminProvider) GetProjects(ctx context.Context) (*admin.Projects, error) {
	filter, err := common.NewSingleValueFilter(common.Project, common.NotEqual, "state", int32(admin.Project_ARCHIVED))
	if err != nil {
		return nil, err
	}
	projectModels, err := p.db.ProjectRepo().List(ctx, repositoryInterfaces.ListResourceInput{
		SortParameter: descCreatedAtSortDBParam,
		InlineFilters: []common.InlineFilter{filter},
	})
	if err != nil {
		return nil, err
	}
	projects := transformers.FromProjectModels(projectModels, p.getDomains())
	return &admin.Projects{
		Projects: projects,
	}, nil
}

func NewDatabaseAdminDataProvider(db repositoryInterfaces.Repository, config runtimeInterfaces.Configuration, resourceManager managerInterfaces.ResourceInterface) interfaces.NebulaAdminDataProvider {
	return &dbAdminProvider{
		db:              db,
		config:          config,
		resourceManager: resourceManager,
	}
}
