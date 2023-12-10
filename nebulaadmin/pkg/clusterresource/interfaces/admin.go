package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

//go:generate mockery -name NebulaAdminDataProvider -output=../mocks -case=underscore

type NebulaAdminDataProvider interface {
	GetClusterResourceAttributes(ctx context.Context, project, domain string) (*admin.ClusterResourceAttributes, error)
	GetProjects(ctx context.Context) (*admin.Projects, error)
}
