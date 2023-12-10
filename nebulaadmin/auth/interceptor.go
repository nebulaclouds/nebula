package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nebulaclouds/nebula/nebulastdlib/logger"
)

func BlanketAuthorization(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {

	identityContext := IdentityContextFromContext(ctx)
	if identityContext.IsEmpty() {
		return handler(ctx, req)
	}

	if !identityContext.Scopes().Has(ScopeAll) {
		logger.Debugf(ctx, "authenticated user doesn't have required scope")
		return nil, status.Errorf(codes.Unauthenticated, "authenticated user doesn't have required scope")
	}

	return handler(ctx, req)
}

// ExecutionUserIdentifierInterceptor injects identityContext.UserID() to identityContext.executionIdentity
func ExecutionUserIdentifierInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	identityContext := IdentityContextFromContext(ctx)
	identityContext = identityContext.WithExecutionUserIdentifier(identityContext.UserID())
	ctx = identityContext.WithContext(ctx)
	return handler(ctx, req)
}
