package interfaces

import (
	"context"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
)

type TagManager interface {
	AddTag(ctx context.Context, request *datacatalog.AddTagRequest) (*datacatalog.AddTagResponse, error)
}
