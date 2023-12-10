package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Interface for managing metadata associated with NamedEntityIdentifiers
type NamedEntityInterface interface {
	GetNamedEntity(ctx context.Context, request admin.NamedEntityGetRequest) (*admin.NamedEntity, error)
	UpdateNamedEntity(ctx context.Context, request admin.NamedEntityUpdateRequest) (*admin.NamedEntityUpdateResponse, error)
	ListNamedEntities(ctx context.Context, request admin.NamedEntityListRequest) (*admin.NamedEntityList, error)
}
