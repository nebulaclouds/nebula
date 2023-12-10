package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
)

// Defines an interface for fetching pre-signed URLs.
type RemoteURLInterface interface {
	// TODO: Refactor for URI to be of type DataReference. We should package a FromString-like function in nebulastdlib
	Get(ctx context.Context, uri string) (admin.UrlBlob, error)
}
