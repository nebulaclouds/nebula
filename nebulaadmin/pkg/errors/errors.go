// Defines the error messages used within NebulaAdmin categorized by common error codes.
package errors

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
)

type NebulaAdminError interface {
	Error() string
	Code() codes.Code
	GRPCStatus() *status.Status
	WithDetails(details proto.Message) (NebulaAdminError, error)
	String() string
}
type nebulaAdminErrorImpl struct {
	status *status.Status
}

func (e *nebulaAdminErrorImpl) Error() string {
	return e.status.Message()
}

func (e *nebulaAdminErrorImpl) Code() codes.Code {
	return e.status.Code()
}

func (e *nebulaAdminErrorImpl) GRPCStatus() *status.Status {
	return e.status
}

func (e *nebulaAdminErrorImpl) String() string {
	return fmt.Sprintf("status: %v", e.status)
}

// enclose the error in the format that grpc server expect from golang:
//
//	https://github.com/grpc/grpc-go/blob/master/status/status.go#L133
func (e *nebulaAdminErrorImpl) WithDetails(details proto.Message) (NebulaAdminError, error) {
	s, err := e.status.WithDetails(details)
	if err != nil {
		return nil, err
	}
	return NewNebulaAdminErrorFromStatus(s), nil
}

func NewNebulaAdminErrorFromStatus(status *status.Status) NebulaAdminError {
	return &nebulaAdminErrorImpl{
		status: status,
	}
}

func NewNebulaAdminError(code codes.Code, message string) NebulaAdminError {
	return &nebulaAdminErrorImpl{
		status: status.New(code, message),
	}
}

func NewNebulaAdminErrorf(code codes.Code, format string, a ...interface{}) NebulaAdminError {
	return NewNebulaAdminError(code, fmt.Sprintf(format, a...))
}

func toStringSlice(errors []error) []string {
	errSlice := make([]string, len(errors))
	for idx, err := range errors {
		errSlice[idx] = err.Error()
	}
	return errSlice
}

func NewCollectedNebulaAdminError(code codes.Code, errors []error) NebulaAdminError {
	return NewNebulaAdminError(code, strings.Join(toStringSlice(errors), ", "))
}

func NewAlreadyInTerminalStateError(ctx context.Context, errorMsg string, curPhase string) NebulaAdminError {
	logger.Warn(ctx, errorMsg)
	alreadyInTerminalPhase := &admin.EventErrorAlreadyInTerminalState{CurrentPhase: curPhase}
	reason := &admin.EventFailureReason{
		Reason: &admin.EventFailureReason_AlreadyInTerminalState{AlreadyInTerminalState: alreadyInTerminalPhase},
	}
	statusErr, transformationErr := NewNebulaAdminError(codes.FailedPrecondition, errorMsg).WithDetails(reason)
	if transformationErr != nil {
		logger.Panicf(ctx, "Failed to wrap grpc status in type 'Error': %v", transformationErr)
		return NewNebulaAdminErrorf(codes.FailedPrecondition, errorMsg)
	}
	return statusErr
}

func NewIncompatibleClusterError(ctx context.Context, errorMsg, curCluster string) NebulaAdminError {
	statusErr, transformationErr := NewNebulaAdminError(codes.FailedPrecondition, errorMsg).WithDetails(&admin.EventFailureReason{
		Reason: &admin.EventFailureReason_IncompatibleCluster{
			IncompatibleCluster: &admin.EventErrorIncompatibleCluster{
				Cluster: curCluster,
			},
		},
	})
	if transformationErr != nil {
		logger.Panicf(ctx, "Failed to wrap grpc status in type 'Error': %v", transformationErr)
		return NewNebulaAdminErrorf(codes.FailedPrecondition, errorMsg)
	}
	return statusErr
}

func NewWorkflowExistsDifferentStructureError(ctx context.Context, request *admin.WorkflowCreateRequest) NebulaAdminError {
	errorMsg := "workflow with different structure already exists"
	statusErr, transformationErr := NewNebulaAdminError(codes.InvalidArgument, errorMsg).WithDetails(&admin.CreateWorkflowFailureReason{
		Reason: &admin.CreateWorkflowFailureReason_ExistsDifferentStructure{
			ExistsDifferentStructure: &admin.WorkflowErrorExistsDifferentStructure{
				Id: request.Id,
			},
		},
	})
	if transformationErr != nil {
		logger.Panicf(ctx, "Failed to wrap grpc status in type 'Error': %v", transformationErr)
		return NewNebulaAdminErrorf(codes.InvalidArgument, errorMsg)
	}
	return statusErr
}

func NewWorkflowExistsIdenticalStructureError(ctx context.Context, request *admin.WorkflowCreateRequest) NebulaAdminError {
	errorMsg := "workflow with identical structure already exists"
	statusErr, transformationErr := NewNebulaAdminError(codes.AlreadyExists, errorMsg).WithDetails(&admin.CreateWorkflowFailureReason{
		Reason: &admin.CreateWorkflowFailureReason_ExistsIdenticalStructure{
			ExistsIdenticalStructure: &admin.WorkflowErrorExistsIdenticalStructure{
				Id: request.Id,
			},
		},
	})
	if transformationErr != nil {
		logger.Panicf(ctx, "Failed to wrap grpc status in type 'Error': %v", transformationErr)
		return NewNebulaAdminErrorf(codes.AlreadyExists, errorMsg)
	}
	return statusErr
}

func IsDoesNotExistError(err error) bool {
	adminError, ok := err.(NebulaAdminError)
	return ok && adminError.Code() == codes.NotFound
}
