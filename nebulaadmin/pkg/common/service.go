package common

import (
	"context"

	"google.golang.org/grpc"
)

// RegisterAdditionalGRPCService is the interface for the plugin hook for additional GRPC service handlers which
// should be also served on the nebulaadmin gRPC server.
type RegisterAdditionalGRPCService = func(ctx context.Context, grpcServer *grpc.Server) error
