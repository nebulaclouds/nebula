// Convenience methods for shared errors.
package shared

import (
	"fmt"

	"google.golang.org/grpc/codes"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
)

const missingFieldFormat = "missing %s"
const invalidArgFormat = "invalid value for %s"

func GetMissingArgumentError(field string) error {
	return errors.NewNebulaAdminError(codes.InvalidArgument, fmt.Sprintf(missingFieldFormat, field))
}

func GetInvalidArgumentError(field string) error {
	return errors.NewNebulaAdminError(codes.InvalidArgument, fmt.Sprintf(invalidArgFormat, field))
}
