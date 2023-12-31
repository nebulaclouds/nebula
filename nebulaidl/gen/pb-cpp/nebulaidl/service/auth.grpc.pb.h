// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: nebulaidl/service/auth.proto
#ifndef GRPC_nebulaidl_2fservice_2fauth_2eproto__INCLUDED
#define GRPC_nebulaidl_2fservice_2fauth_2eproto__INCLUDED

#include "nebulaidl/service/auth.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_generic_service.h>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/method_handler_impl.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/codegen/rpc_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/impl/codegen/stub_options.h>
#include <grpcpp/impl/codegen/sync_stream.h>

namespace grpc_impl {
class Channel;
class CompletionQueue;
class ServerCompletionQueue;
}  // namespace grpc_impl

namespace grpc {
namespace experimental {
template <typename RequestT, typename ResponseT>
class MessageAllocator;
}  // namespace experimental
}  // namespace grpc_impl

namespace grpc {
class ServerContext;
}  // namespace grpc

namespace nebulaidl {
namespace service {

// The following defines an RPC service that is also served over HTTP via grpc-gateway.
// Standard response codes for both are defined here: https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go
// RPCs defined in this service must be anonymously accessible.
class AuthMetadataService final {
 public:
  static constexpr char const* service_full_name() {
    return "nebulaidl.service.AuthMetadataService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // Anonymously accessible. Retrieves local or external oauth authorization server metadata.
    virtual ::grpc::Status GetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::nebulaidl::service::OAuth2MetadataResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::OAuth2MetadataResponse>> AsyncGetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::OAuth2MetadataResponse>>(AsyncGetOAuth2MetadataRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::OAuth2MetadataResponse>> PrepareAsyncGetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::OAuth2MetadataResponse>>(PrepareAsyncGetOAuth2MetadataRaw(context, request, cq));
    }
    // Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization
    // requests.
    virtual ::grpc::Status GetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::PublicClientAuthConfigResponse>> AsyncGetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::PublicClientAuthConfigResponse>>(AsyncGetPublicClientConfigRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::PublicClientAuthConfigResponse>> PrepareAsyncGetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::PublicClientAuthConfigResponse>>(PrepareAsyncGetPublicClientConfigRaw(context, request, cq));
    }
    class experimental_async_interface {
     public:
      virtual ~experimental_async_interface() {}
      // Anonymously accessible. Retrieves local or external oauth authorization server metadata.
      virtual void GetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetOAuth2Metadata(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::OAuth2MetadataResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      virtual void GetOAuth2Metadata(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::OAuth2MetadataResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      // Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization
      // requests.
      virtual void GetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetPublicClientConfig(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
      virtual void GetPublicClientConfig(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) = 0;
    };
    virtual class experimental_async_interface* experimental_async() { return nullptr; }
  private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::OAuth2MetadataResponse>* AsyncGetOAuth2MetadataRaw(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::OAuth2MetadataResponse>* PrepareAsyncGetOAuth2MetadataRaw(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::PublicClientAuthConfigResponse>* AsyncGetPublicClientConfigRaw(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::nebulaidl::service::PublicClientAuthConfigResponse>* PrepareAsyncGetPublicClientConfigRaw(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel);
    ::grpc::Status GetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::nebulaidl::service::OAuth2MetadataResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::OAuth2MetadataResponse>> AsyncGetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::OAuth2MetadataResponse>>(AsyncGetOAuth2MetadataRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::OAuth2MetadataResponse>> PrepareAsyncGetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::OAuth2MetadataResponse>>(PrepareAsyncGetOAuth2MetadataRaw(context, request, cq));
    }
    ::grpc::Status GetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::PublicClientAuthConfigResponse>> AsyncGetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::PublicClientAuthConfigResponse>>(AsyncGetPublicClientConfigRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::PublicClientAuthConfigResponse>> PrepareAsyncGetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::PublicClientAuthConfigResponse>>(PrepareAsyncGetPublicClientConfigRaw(context, request, cq));
    }
    class experimental_async final :
      public StubInterface::experimental_async_interface {
     public:
      void GetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response, std::function<void(::grpc::Status)>) override;
      void GetOAuth2Metadata(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::OAuth2MetadataResponse* response, std::function<void(::grpc::Status)>) override;
      void GetOAuth2Metadata(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void GetOAuth2Metadata(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::OAuth2MetadataResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void GetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, std::function<void(::grpc::Status)>) override;
      void GetPublicClientConfig(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, std::function<void(::grpc::Status)>) override;
      void GetPublicClientConfig(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
      void GetPublicClientConfig(::grpc::ClientContext* context, const ::grpc::ByteBuffer* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, ::grpc::experimental::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit experimental_async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class experimental_async_interface* experimental_async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class experimental_async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::OAuth2MetadataResponse>* AsyncGetOAuth2MetadataRaw(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::OAuth2MetadataResponse>* PrepareAsyncGetOAuth2MetadataRaw(::grpc::ClientContext* context, const ::nebulaidl::service::OAuth2MetadataRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::PublicClientAuthConfigResponse>* AsyncGetPublicClientConfigRaw(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::nebulaidl::service::PublicClientAuthConfigResponse>* PrepareAsyncGetPublicClientConfigRaw(::grpc::ClientContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_GetOAuth2Metadata_;
    const ::grpc::internal::RpcMethod rpcmethod_GetPublicClientConfig_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // Anonymously accessible. Retrieves local or external oauth authorization server metadata.
    virtual ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response);
    // Anonymously accessible. Retrieves the client information clients should use when initiating OAuth2 authorization
    // requests.
    virtual ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_GetOAuth2Metadata : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_GetOAuth2Metadata() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_GetOAuth2Metadata() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetOAuth2Metadata(::grpc::ServerContext* context, ::nebulaidl::service::OAuth2MetadataRequest* request, ::grpc::ServerAsyncResponseWriter< ::nebulaidl::service::OAuth2MetadataResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_GetPublicClientConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithAsyncMethod_GetPublicClientConfig() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_GetPublicClientConfig() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetPublicClientConfig(::grpc::ServerContext* context, ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::grpc::ServerAsyncResponseWriter< ::nebulaidl::service::PublicClientAuthConfigResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_GetOAuth2Metadata<WithAsyncMethod_GetPublicClientConfig<Service > > AsyncService;
  template <class BaseClass>
  class ExperimentalWithCallbackMethod_GetOAuth2Metadata : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithCallbackMethod_GetOAuth2Metadata() {
      ::grpc::Service::experimental().MarkMethodCallback(0,
        new ::grpc::internal::CallbackUnaryHandler< ::nebulaidl::service::OAuth2MetadataRequest, ::nebulaidl::service::OAuth2MetadataResponse>(
          [this](::grpc::ServerContext* context,
                 const ::nebulaidl::service::OAuth2MetadataRequest* request,
                 ::nebulaidl::service::OAuth2MetadataResponse* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   return this->GetOAuth2Metadata(context, request, response, controller);
                 }));
    }
    void SetMessageAllocatorFor_GetOAuth2Metadata(
        ::grpc::experimental::MessageAllocator< ::nebulaidl::service::OAuth2MetadataRequest, ::nebulaidl::service::OAuth2MetadataResponse>* allocator) {
      static_cast<::grpc::internal::CallbackUnaryHandler< ::nebulaidl::service::OAuth2MetadataRequest, ::nebulaidl::service::OAuth2MetadataResponse>*>(
          ::grpc::Service::experimental().GetHandler(0))
              ->SetMessageAllocator(allocator);
    }
    ~ExperimentalWithCallbackMethod_GetOAuth2Metadata() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class ExperimentalWithCallbackMethod_GetPublicClientConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithCallbackMethod_GetPublicClientConfig() {
      ::grpc::Service::experimental().MarkMethodCallback(1,
        new ::grpc::internal::CallbackUnaryHandler< ::nebulaidl::service::PublicClientAuthConfigRequest, ::nebulaidl::service::PublicClientAuthConfigResponse>(
          [this](::grpc::ServerContext* context,
                 const ::nebulaidl::service::PublicClientAuthConfigRequest* request,
                 ::nebulaidl::service::PublicClientAuthConfigResponse* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   return this->GetPublicClientConfig(context, request, response, controller);
                 }));
    }
    void SetMessageAllocatorFor_GetPublicClientConfig(
        ::grpc::experimental::MessageAllocator< ::nebulaidl::service::PublicClientAuthConfigRequest, ::nebulaidl::service::PublicClientAuthConfigResponse>* allocator) {
      static_cast<::grpc::internal::CallbackUnaryHandler< ::nebulaidl::service::PublicClientAuthConfigRequest, ::nebulaidl::service::PublicClientAuthConfigResponse>*>(
          ::grpc::Service::experimental().GetHandler(1))
              ->SetMessageAllocator(allocator);
    }
    ~ExperimentalWithCallbackMethod_GetPublicClientConfig() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  typedef ExperimentalWithCallbackMethod_GetOAuth2Metadata<ExperimentalWithCallbackMethod_GetPublicClientConfig<Service > > ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_GetOAuth2Metadata : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_GetOAuth2Metadata() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_GetOAuth2Metadata() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_GetPublicClientConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithGenericMethod_GetPublicClientConfig() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_GetPublicClientConfig() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetOAuth2Metadata : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithRawMethod_GetOAuth2Metadata() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_GetOAuth2Metadata() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetOAuth2Metadata(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetPublicClientConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithRawMethod_GetPublicClientConfig() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_GetPublicClientConfig() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetPublicClientConfig(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class ExperimentalWithRawCallbackMethod_GetOAuth2Metadata : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithRawCallbackMethod_GetOAuth2Metadata() {
      ::grpc::Service::experimental().MarkMethodRawCallback(0,
        new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
          [this](::grpc::ServerContext* context,
                 const ::grpc::ByteBuffer* request,
                 ::grpc::ByteBuffer* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   this->GetOAuth2Metadata(context, request, response, controller);
                 }));
    }
    ~ExperimentalWithRawCallbackMethod_GetOAuth2Metadata() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void GetOAuth2Metadata(::grpc::ServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class ExperimentalWithRawCallbackMethod_GetPublicClientConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    ExperimentalWithRawCallbackMethod_GetPublicClientConfig() {
      ::grpc::Service::experimental().MarkMethodRawCallback(1,
        new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
          [this](::grpc::ServerContext* context,
                 const ::grpc::ByteBuffer* request,
                 ::grpc::ByteBuffer* response,
                 ::grpc::experimental::ServerCallbackRpcController* controller) {
                   this->GetPublicClientConfig(context, request, response, controller);
                 }));
    }
    ~ExperimentalWithRawCallbackMethod_GetPublicClientConfig() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual void GetPublicClientConfig(::grpc::ServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response, ::grpc::experimental::ServerCallbackRpcController* controller) { controller->Finish(::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "")); }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetOAuth2Metadata : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_GetOAuth2Metadata() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler< ::nebulaidl::service::OAuth2MetadataRequest, ::nebulaidl::service::OAuth2MetadataResponse>(std::bind(&WithStreamedUnaryMethod_GetOAuth2Metadata<BaseClass>::StreamedGetOAuth2Metadata, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_GetOAuth2Metadata() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetOAuth2Metadata(::grpc::ServerContext* context, const ::nebulaidl::service::OAuth2MetadataRequest* request, ::nebulaidl::service::OAuth2MetadataResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetOAuth2Metadata(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::nebulaidl::service::OAuth2MetadataRequest,::nebulaidl::service::OAuth2MetadataResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetPublicClientConfig : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service *service) {}
   public:
    WithStreamedUnaryMethod_GetPublicClientConfig() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler< ::nebulaidl::service::PublicClientAuthConfigRequest, ::nebulaidl::service::PublicClientAuthConfigResponse>(std::bind(&WithStreamedUnaryMethod_GetPublicClientConfig<BaseClass>::StreamedGetPublicClientConfig, this, std::placeholders::_1, std::placeholders::_2)));
    }
    ~WithStreamedUnaryMethod_GetPublicClientConfig() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetPublicClientConfig(::grpc::ServerContext* context, const ::nebulaidl::service::PublicClientAuthConfigRequest* request, ::nebulaidl::service::PublicClientAuthConfigResponse* response) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetPublicClientConfig(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::nebulaidl::service::PublicClientAuthConfigRequest,::nebulaidl::service::PublicClientAuthConfigResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_GetOAuth2Metadata<WithStreamedUnaryMethod_GetPublicClientConfig<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_GetOAuth2Metadata<WithStreamedUnaryMethod_GetPublicClientConfig<Service > > StreamedService;
};

}  // namespace service
}  // namespace nebulaidl


#endif  // GRPC_nebulaidl_2fservice_2fauth_2eproto__INCLUDED
