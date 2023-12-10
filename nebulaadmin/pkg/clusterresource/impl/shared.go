package impl

import (
	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
)

func NewMissingEntityError(entity string) error {
	return errors.NewNebulaAdminErrorf(codes.NotFound, "Failed to find [%s]", entity)
}
