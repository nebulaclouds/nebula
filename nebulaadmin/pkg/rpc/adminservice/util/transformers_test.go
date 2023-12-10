package util

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/common"
	adminErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	mockScope "github.com/nebulaclouds/nebula/nebulastdlib/promutils"
)

var testRequestMetrics = NewRequestMetrics(mockScope.NewTestScope(), "foo")

func TestTransformError_NebulaAdminError(t *testing.T) {
	invalidArgError := adminErrors.NewNebulaAdminError(codes.InvalidArgument, "invalid arg")

	transformedError := TransformAndRecordError(invalidArgError, &testRequestMetrics)

	transormerStatus, ok := status.FromError(transformedError)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, transormerStatus.Code())
}

func TestTransformError_NebulaAdminErrorWithDetails(t *testing.T) {
	terminalStateError := adminErrors.NewAlreadyInTerminalStateError(context.Background(), "terminal state", "curPhase")

	transformedError := TransformAndRecordError(terminalStateError, &testRequestMetrics)

	transormerStatus, ok := status.FromError(transformedError)
	assert.True(t, ok)
	assert.Equal(t, codes.FailedPrecondition, transormerStatus.Code())
	assert.Equal(t, 1, len(transormerStatus.Details()))
}

func TestTransformError_GRPCError(t *testing.T) {
	err := status.Error(codes.InvalidArgument, strings.Repeat("X", common.MaxResponseStatusBytes+1))

	transformedError := TransformAndRecordError(err, &testRequestMetrics)

	transormerStatus, ok := status.FromError(transformedError)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, transormerStatus.Code())
	assert.Len(t, transormerStatus.Message(), common.MaxResponseStatusBytes)
}

func TestTransformError_BasicError(t *testing.T) {
	err := errors.New("some error")

	transformedError := TransformAndRecordError(err, &testRequestMetrics)

	transormerStatus, ok := status.FromError(transformedError)
	assert.True(t, ok)
	assert.Equal(t, codes.Internal, transormerStatus.Code())
}

func TestTruncateErrorMessage(t *testing.T) {
	err := adminErrors.NewNebulaAdminError(codes.InvalidArgument, strings.Repeat("X", common.MaxResponseStatusBytes+1))

	transformedError := TransformAndRecordError(err, &testRequestMetrics)

	transormerStatus, ok := status.FromError(transformedError)
	assert.True(t, ok)
	assert.Equal(t, codes.InvalidArgument, transormerStatus.Code())
	assert.Len(t, transormerStatus.Message(), common.MaxResponseStatusBytes)
}
