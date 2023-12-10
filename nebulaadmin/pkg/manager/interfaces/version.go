package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

//go:generate mockery -name VersionInterface -output=../mocks -case=underscore

// Interface for managing Nebula admin version
type VersionInterface interface {
	GetVersion(ctx context.Context, r *admin.GetVersionRequest) (*admin.GetVersionResponse, error)
}
