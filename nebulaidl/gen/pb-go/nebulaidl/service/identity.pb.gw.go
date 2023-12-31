// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: nebulaidl/service/identity.proto

/*
Package service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package service

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_IdentityService_UserInfo_0(ctx context.Context, marshaler runtime.Marshaler, client IdentityServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UserInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := client.UserInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterIdentityServiceHandlerFromEndpoint is same as RegisterIdentityServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterIdentityServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterIdentityServiceHandler(ctx, mux, conn)
}

// RegisterIdentityServiceHandler registers the http handlers for service IdentityService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterIdentityServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterIdentityServiceHandlerClient(ctx, mux, NewIdentityServiceClient(conn))
}

// RegisterIdentityServiceHandlerClient registers the http handlers for service IdentityService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "IdentityServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "IdentityServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "IdentityServiceClient" to call the correct interceptors.
func RegisterIdentityServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client IdentityServiceClient) error {

	mux.Handle("GET", pattern_IdentityService_UserInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IdentityService_UserInfo_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IdentityService_UserInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_IdentityService_UserInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"me"}, ""))
)

var (
	forward_IdentityService_UserInfo_0 = runtime.ForwardResponseMessage
)
