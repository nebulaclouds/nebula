// Generic errors used across files in repositories/
package errors

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
)

const (
	singletonNotFound = "missing singleton entity of type %s"
	notFound          = "missing entity of type %s with identifier %v"
	idNotFound        = "missing entity of type %s"
	invalidInput      = "missing and/or invalid parameters: %s"
)

func GetMissingEntityError(entityType string, identifier proto.Message) errors.NebulaAdminError {
	return errors.NewNebulaAdminErrorf(codes.NotFound, notFound, entityType, identifier)
}

func GetSingletonMissingEntityError(entityType string) errors.NebulaAdminError {
	return errors.NewNebulaAdminErrorf(codes.NotFound, singletonNotFound, entityType)
}

func GetMissingEntityByIDError(entityType string) errors.NebulaAdminError {
	return errors.NewNebulaAdminErrorf(codes.NotFound, idNotFound, entityType)
}

func GetInvalidInputError(input string) errors.NebulaAdminError {
	return errors.NewNebulaAdminErrorf(codes.InvalidArgument, invalidInput, input)
}
