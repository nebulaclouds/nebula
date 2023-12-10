// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: nebulaidl/datacatalog/datacatalog.proto

#include "nebulaidl/datacatalog/datacatalog.pb.h"
#include "nebulaidl/datacatalog/datacatalog.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace datacatalog {

static const char* DataCatalog_method_names[] = {
  "/datacatalog.DataCatalog/CreateDataset",
  "/datacatalog.DataCatalog/GetDataset",
  "/datacatalog.DataCatalog/CreateArtifact",
  "/datacatalog.DataCatalog/GetArtifact",
  "/datacatalog.DataCatalog/AddTag",
  "/datacatalog.DataCatalog/ListArtifacts",
  "/datacatalog.DataCatalog/ListDatasets",
  "/datacatalog.DataCatalog/UpdateArtifact",
  "/datacatalog.DataCatalog/GetOrExtendReservation",
  "/datacatalog.DataCatalog/ReleaseReservation",
};

std::unique_ptr< DataCatalog::Stub> DataCatalog::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< DataCatalog::Stub> stub(new DataCatalog::Stub(channel));
  return stub;
}

DataCatalog::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_CreateDataset_(DataCatalog_method_names[0], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetDataset_(DataCatalog_method_names[1], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_CreateArtifact_(DataCatalog_method_names[2], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetArtifact_(DataCatalog_method_names[3], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_AddTag_(DataCatalog_method_names[4], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ListArtifacts_(DataCatalog_method_names[5], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ListDatasets_(DataCatalog_method_names[6], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_UpdateArtifact_(DataCatalog_method_names[7], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_GetOrExtendReservation_(DataCatalog_method_names[8], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_ReleaseReservation_(DataCatalog_method_names[9], ::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status DataCatalog::Stub::CreateDataset(::grpc::ClientContext* context, const ::datacatalog::CreateDatasetRequest& request, ::datacatalog::CreateDatasetResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_CreateDataset_, context, request, response);
}

void DataCatalog::Stub::experimental_async::CreateDataset(::grpc::ClientContext* context, const ::datacatalog::CreateDatasetRequest* request, ::datacatalog::CreateDatasetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateDataset_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::CreateDataset(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::CreateDatasetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateDataset_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::CreateDataset(::grpc::ClientContext* context, const ::datacatalog::CreateDatasetRequest* request, ::datacatalog::CreateDatasetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateDataset_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::CreateDataset(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::CreateDatasetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateDataset_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::CreateDatasetResponse>* DataCatalog::Stub::AsyncCreateDatasetRaw(::grpc::ClientContext* context, const ::datacatalog::CreateDatasetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::CreateDatasetResponse>::Create(channel_.get(), cq, rpcmethod_CreateDataset_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::CreateDatasetResponse>* DataCatalog::Stub::PrepareAsyncCreateDatasetRaw(::grpc::ClientContext* context, const ::datacatalog::CreateDatasetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::CreateDatasetResponse>::Create(channel_.get(), cq, rpcmethod_CreateDataset_, context, request, false);
}

::grpc::Status DataCatalog::Stub::GetDataset(::grpc::ClientContext* context, const ::datacatalog::GetDatasetRequest& request, ::datacatalog::GetDatasetResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetDataset_, context, request, response);
}

void DataCatalog::Stub::experimental_async::GetDataset(::grpc::ClientContext* context, const ::datacatalog::GetDatasetRequest* request, ::datacatalog::GetDatasetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetDataset_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::GetDataset(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::GetDatasetResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetDataset_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::GetDataset(::grpc::ClientContext* context, const ::datacatalog::GetDatasetRequest* request, ::datacatalog::GetDatasetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetDataset_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::GetDataset(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::GetDatasetResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetDataset_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::GetDatasetResponse>* DataCatalog::Stub::AsyncGetDatasetRaw(::grpc::ClientContext* context, const ::datacatalog::GetDatasetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::GetDatasetResponse>::Create(channel_.get(), cq, rpcmethod_GetDataset_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::GetDatasetResponse>* DataCatalog::Stub::PrepareAsyncGetDatasetRaw(::grpc::ClientContext* context, const ::datacatalog::GetDatasetRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::GetDatasetResponse>::Create(channel_.get(), cq, rpcmethod_GetDataset_, context, request, false);
}

::grpc::Status DataCatalog::Stub::CreateArtifact(::grpc::ClientContext* context, const ::datacatalog::CreateArtifactRequest& request, ::datacatalog::CreateArtifactResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_CreateArtifact_, context, request, response);
}

void DataCatalog::Stub::experimental_async::CreateArtifact(::grpc::ClientContext* context, const ::datacatalog::CreateArtifactRequest* request, ::datacatalog::CreateArtifactResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateArtifact_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::CreateArtifact(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::CreateArtifactResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_CreateArtifact_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::CreateArtifact(::grpc::ClientContext* context, const ::datacatalog::CreateArtifactRequest* request, ::datacatalog::CreateArtifactResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateArtifact_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::CreateArtifact(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::CreateArtifactResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_CreateArtifact_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::CreateArtifactResponse>* DataCatalog::Stub::AsyncCreateArtifactRaw(::grpc::ClientContext* context, const ::datacatalog::CreateArtifactRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::CreateArtifactResponse>::Create(channel_.get(), cq, rpcmethod_CreateArtifact_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::CreateArtifactResponse>* DataCatalog::Stub::PrepareAsyncCreateArtifactRaw(::grpc::ClientContext* context, const ::datacatalog::CreateArtifactRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::CreateArtifactResponse>::Create(channel_.get(), cq, rpcmethod_CreateArtifact_, context, request, false);
}

::grpc::Status DataCatalog::Stub::GetArtifact(::grpc::ClientContext* context, const ::datacatalog::GetArtifactRequest& request, ::datacatalog::GetArtifactResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetArtifact_, context, request, response);
}

void DataCatalog::Stub::experimental_async::GetArtifact(::grpc::ClientContext* context, const ::datacatalog::GetArtifactRequest* request, ::datacatalog::GetArtifactResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetArtifact_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::GetArtifact(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::GetArtifactResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetArtifact_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::GetArtifact(::grpc::ClientContext* context, const ::datacatalog::GetArtifactRequest* request, ::datacatalog::GetArtifactResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetArtifact_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::GetArtifact(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::GetArtifactResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetArtifact_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::GetArtifactResponse>* DataCatalog::Stub::AsyncGetArtifactRaw(::grpc::ClientContext* context, const ::datacatalog::GetArtifactRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::GetArtifactResponse>::Create(channel_.get(), cq, rpcmethod_GetArtifact_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::GetArtifactResponse>* DataCatalog::Stub::PrepareAsyncGetArtifactRaw(::grpc::ClientContext* context, const ::datacatalog::GetArtifactRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::GetArtifactResponse>::Create(channel_.get(), cq, rpcmethod_GetArtifact_, context, request, false);
}

::grpc::Status DataCatalog::Stub::AddTag(::grpc::ClientContext* context, const ::datacatalog::AddTagRequest& request, ::datacatalog::AddTagResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_AddTag_, context, request, response);
}

void DataCatalog::Stub::experimental_async::AddTag(::grpc::ClientContext* context, const ::datacatalog::AddTagRequest* request, ::datacatalog::AddTagResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_AddTag_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::AddTag(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::AddTagResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_AddTag_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::AddTag(::grpc::ClientContext* context, const ::datacatalog::AddTagRequest* request, ::datacatalog::AddTagResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_AddTag_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::AddTag(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::AddTagResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_AddTag_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::AddTagResponse>* DataCatalog::Stub::AsyncAddTagRaw(::grpc::ClientContext* context, const ::datacatalog::AddTagRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::AddTagResponse>::Create(channel_.get(), cq, rpcmethod_AddTag_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::AddTagResponse>* DataCatalog::Stub::PrepareAsyncAddTagRaw(::grpc::ClientContext* context, const ::datacatalog::AddTagRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::AddTagResponse>::Create(channel_.get(), cq, rpcmethod_AddTag_, context, request, false);
}

::grpc::Status DataCatalog::Stub::ListArtifacts(::grpc::ClientContext* context, const ::datacatalog::ListArtifactsRequest& request, ::datacatalog::ListArtifactsResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_ListArtifacts_, context, request, response);
}

void DataCatalog::Stub::experimental_async::ListArtifacts(::grpc::ClientContext* context, const ::datacatalog::ListArtifactsRequest* request, ::datacatalog::ListArtifactsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_ListArtifacts_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::ListArtifacts(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::ListArtifactsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_ListArtifacts_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::ListArtifacts(::grpc::ClientContext* context, const ::datacatalog::ListArtifactsRequest* request, ::datacatalog::ListArtifactsResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_ListArtifacts_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::ListArtifacts(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::ListArtifactsResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_ListArtifacts_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::ListArtifactsResponse>* DataCatalog::Stub::AsyncListArtifactsRaw(::grpc::ClientContext* context, const ::datacatalog::ListArtifactsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::ListArtifactsResponse>::Create(channel_.get(), cq, rpcmethod_ListArtifacts_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::ListArtifactsResponse>* DataCatalog::Stub::PrepareAsyncListArtifactsRaw(::grpc::ClientContext* context, const ::datacatalog::ListArtifactsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::ListArtifactsResponse>::Create(channel_.get(), cq, rpcmethod_ListArtifacts_, context, request, false);
}

::grpc::Status DataCatalog::Stub::ListDatasets(::grpc::ClientContext* context, const ::datacatalog::ListDatasetsRequest& request, ::datacatalog::ListDatasetsResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_ListDatasets_, context, request, response);
}

void DataCatalog::Stub::experimental_async::ListDatasets(::grpc::ClientContext* context, const ::datacatalog::ListDatasetsRequest* request, ::datacatalog::ListDatasetsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_ListDatasets_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::ListDatasets(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::ListDatasetsResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_ListDatasets_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::ListDatasets(::grpc::ClientContext* context, const ::datacatalog::ListDatasetsRequest* request, ::datacatalog::ListDatasetsResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_ListDatasets_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::ListDatasets(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::ListDatasetsResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_ListDatasets_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::ListDatasetsResponse>* DataCatalog::Stub::AsyncListDatasetsRaw(::grpc::ClientContext* context, const ::datacatalog::ListDatasetsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::ListDatasetsResponse>::Create(channel_.get(), cq, rpcmethod_ListDatasets_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::ListDatasetsResponse>* DataCatalog::Stub::PrepareAsyncListDatasetsRaw(::grpc::ClientContext* context, const ::datacatalog::ListDatasetsRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::ListDatasetsResponse>::Create(channel_.get(), cq, rpcmethod_ListDatasets_, context, request, false);
}

::grpc::Status DataCatalog::Stub::UpdateArtifact(::grpc::ClientContext* context, const ::datacatalog::UpdateArtifactRequest& request, ::datacatalog::UpdateArtifactResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_UpdateArtifact_, context, request, response);
}

void DataCatalog::Stub::experimental_async::UpdateArtifact(::grpc::ClientContext* context, const ::datacatalog::UpdateArtifactRequest* request, ::datacatalog::UpdateArtifactResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_UpdateArtifact_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::UpdateArtifact(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::UpdateArtifactResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_UpdateArtifact_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::UpdateArtifact(::grpc::ClientContext* context, const ::datacatalog::UpdateArtifactRequest* request, ::datacatalog::UpdateArtifactResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_UpdateArtifact_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::UpdateArtifact(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::UpdateArtifactResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_UpdateArtifact_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::UpdateArtifactResponse>* DataCatalog::Stub::AsyncUpdateArtifactRaw(::grpc::ClientContext* context, const ::datacatalog::UpdateArtifactRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::UpdateArtifactResponse>::Create(channel_.get(), cq, rpcmethod_UpdateArtifact_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::UpdateArtifactResponse>* DataCatalog::Stub::PrepareAsyncUpdateArtifactRaw(::grpc::ClientContext* context, const ::datacatalog::UpdateArtifactRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::UpdateArtifactResponse>::Create(channel_.get(), cq, rpcmethod_UpdateArtifact_, context, request, false);
}

::grpc::Status DataCatalog::Stub::GetOrExtendReservation(::grpc::ClientContext* context, const ::datacatalog::GetOrExtendReservationRequest& request, ::datacatalog::GetOrExtendReservationResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_GetOrExtendReservation_, context, request, response);
}

void DataCatalog::Stub::experimental_async::GetOrExtendReservation(::grpc::ClientContext* context, const ::datacatalog::GetOrExtendReservationRequest* request, ::datacatalog::GetOrExtendReservationResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetOrExtendReservation_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::GetOrExtendReservation(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::GetOrExtendReservationResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_GetOrExtendReservation_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::GetOrExtendReservation(::grpc::ClientContext* context, const ::datacatalog::GetOrExtendReservationRequest* request, ::datacatalog::GetOrExtendReservationResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetOrExtendReservation_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::GetOrExtendReservation(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::GetOrExtendReservationResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_GetOrExtendReservation_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::GetOrExtendReservationResponse>* DataCatalog::Stub::AsyncGetOrExtendReservationRaw(::grpc::ClientContext* context, const ::datacatalog::GetOrExtendReservationRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::GetOrExtendReservationResponse>::Create(channel_.get(), cq, rpcmethod_GetOrExtendReservation_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::GetOrExtendReservationResponse>* DataCatalog::Stub::PrepareAsyncGetOrExtendReservationRaw(::grpc::ClientContext* context, const ::datacatalog::GetOrExtendReservationRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::GetOrExtendReservationResponse>::Create(channel_.get(), cq, rpcmethod_GetOrExtendReservation_, context, request, false);
}

::grpc::Status DataCatalog::Stub::ReleaseReservation(::grpc::ClientContext* context, const ::datacatalog::ReleaseReservationRequest& request, ::datacatalog::ReleaseReservationResponse* response) {
  return ::grpc::internal::BlockingUnaryCall(channel_.get(), rpcmethod_ReleaseReservation_, context, request, response);
}

void DataCatalog::Stub::experimental_async::ReleaseReservation(::grpc::ClientContext* context, const ::datacatalog::ReleaseReservationRequest* request, ::datacatalog::ReleaseReservationResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_ReleaseReservation_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::ReleaseReservation(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::ReleaseReservationResponse* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall(stub_->channel_.get(), stub_->rpcmethod_ReleaseReservation_, context, request, response, std::move(f));
}

void DataCatalog::Stub::experimental_async::ReleaseReservation(::grpc::ClientContext* context, const ::datacatalog::ReleaseReservationRequest* request, ::datacatalog::ReleaseReservationResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_ReleaseReservation_, context, request, response, reactor);
}

void DataCatalog::Stub::experimental_async::ReleaseReservation(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::datacatalog::ReleaseReservationResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create(stub_->channel_.get(), stub_->rpcmethod_ReleaseReservation_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::ReleaseReservationResponse>* DataCatalog::Stub::AsyncReleaseReservationRaw(::grpc::ClientContext* context, const ::datacatalog::ReleaseReservationRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::ReleaseReservationResponse>::Create(channel_.get(), cq, rpcmethod_ReleaseReservation_, context, request, true);
}

::grpc::ClientAsyncResponseReader< ::datacatalog::ReleaseReservationResponse>* DataCatalog::Stub::PrepareAsyncReleaseReservationRaw(::grpc::ClientContext* context, const ::datacatalog::ReleaseReservationRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderFactory< ::datacatalog::ReleaseReservationResponse>::Create(channel_.get(), cq, rpcmethod_ReleaseReservation_, context, request, false);
}

DataCatalog::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::CreateDatasetRequest, ::datacatalog::CreateDatasetResponse>(
          std::mem_fn(&DataCatalog::Service::CreateDataset), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[1],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::GetDatasetRequest, ::datacatalog::GetDatasetResponse>(
          std::mem_fn(&DataCatalog::Service::GetDataset), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[2],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::CreateArtifactRequest, ::datacatalog::CreateArtifactResponse>(
          std::mem_fn(&DataCatalog::Service::CreateArtifact), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[3],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::GetArtifactRequest, ::datacatalog::GetArtifactResponse>(
          std::mem_fn(&DataCatalog::Service::GetArtifact), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[4],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::AddTagRequest, ::datacatalog::AddTagResponse>(
          std::mem_fn(&DataCatalog::Service::AddTag), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[5],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::ListArtifactsRequest, ::datacatalog::ListArtifactsResponse>(
          std::mem_fn(&DataCatalog::Service::ListArtifacts), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[6],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::ListDatasetsRequest, ::datacatalog::ListDatasetsResponse>(
          std::mem_fn(&DataCatalog::Service::ListDatasets), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[7],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::UpdateArtifactRequest, ::datacatalog::UpdateArtifactResponse>(
          std::mem_fn(&DataCatalog::Service::UpdateArtifact), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[8],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::GetOrExtendReservationRequest, ::datacatalog::GetOrExtendReservationResponse>(
          std::mem_fn(&DataCatalog::Service::GetOrExtendReservation), this)));
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      DataCatalog_method_names[9],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< DataCatalog::Service, ::datacatalog::ReleaseReservationRequest, ::datacatalog::ReleaseReservationResponse>(
          std::mem_fn(&DataCatalog::Service::ReleaseReservation), this)));
}

DataCatalog::Service::~Service() {
}

::grpc::Status DataCatalog::Service::CreateDataset(::grpc::ServerContext* context, const ::datacatalog::CreateDatasetRequest* request, ::datacatalog::CreateDatasetResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::GetDataset(::grpc::ServerContext* context, const ::datacatalog::GetDatasetRequest* request, ::datacatalog::GetDatasetResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::CreateArtifact(::grpc::ServerContext* context, const ::datacatalog::CreateArtifactRequest* request, ::datacatalog::CreateArtifactResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::GetArtifact(::grpc::ServerContext* context, const ::datacatalog::GetArtifactRequest* request, ::datacatalog::GetArtifactResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::AddTag(::grpc::ServerContext* context, const ::datacatalog::AddTagRequest* request, ::datacatalog::AddTagResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::ListArtifacts(::grpc::ServerContext* context, const ::datacatalog::ListArtifactsRequest* request, ::datacatalog::ListArtifactsResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::ListDatasets(::grpc::ServerContext* context, const ::datacatalog::ListDatasetsRequest* request, ::datacatalog::ListDatasetsResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::UpdateArtifact(::grpc::ServerContext* context, const ::datacatalog::UpdateArtifactRequest* request, ::datacatalog::UpdateArtifactResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::GetOrExtendReservation(::grpc::ServerContext* context, const ::datacatalog::GetOrExtendReservationRequest* request, ::datacatalog::GetOrExtendReservationResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status DataCatalog::Service::ReleaseReservation(::grpc::ServerContext* context, const ::datacatalog::ReleaseReservationRequest* request, ::datacatalog::ReleaseReservationResponse* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace datacatalog
