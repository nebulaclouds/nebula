package errors

import (
	admin_errors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
)

// Defines the basic error transformer interface that all database types must implement.
type ErrorTransformer interface {
	ToNebulaAdminError(err error) admin_errors.NebulaAdminError
}
