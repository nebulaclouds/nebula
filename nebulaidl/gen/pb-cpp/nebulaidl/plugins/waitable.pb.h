// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/plugins/waitable.proto

#ifndef PROTOBUF_INCLUDED_nebulaidl_2fplugins_2fwaitable_2eproto
#define PROTOBUF_INCLUDED_nebulaidl_2fplugins_2fwaitable_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include "nebulaidl/core/execution.pb.h"
#include "nebulaidl/core/identifier.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_nebulaidl_2fplugins_2fwaitable_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_nebulaidl_2fplugins_2fwaitable_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_nebulaidl_2fplugins_2fwaitable_2eproto();
namespace nebulaidl {
namespace plugins {
class Waitable;
class WaitableDefaultTypeInternal;
extern WaitableDefaultTypeInternal _Waitable_default_instance_;
}  // namespace plugins
}  // namespace nebulaidl
namespace google {
namespace protobuf {
template<> ::nebulaidl::plugins::Waitable* Arena::CreateMaybeMessage<::nebulaidl::plugins::Waitable>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace nebulaidl {
namespace plugins {

// ===================================================================

class Waitable final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:nebulaidl.plugins.Waitable) */ {
 public:
  Waitable();
  virtual ~Waitable();

  Waitable(const Waitable& from);

  inline Waitable& operator=(const Waitable& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  Waitable(Waitable&& from) noexcept
    : Waitable() {
    *this = ::std::move(from);
  }

  inline Waitable& operator=(Waitable&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const Waitable& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const Waitable* internal_default_instance() {
    return reinterpret_cast<const Waitable*>(
               &_Waitable_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(Waitable* other);
  friend void swap(Waitable& a, Waitable& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline Waitable* New() const final {
    return CreateMaybeMessage<Waitable>(nullptr);
  }

  Waitable* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<Waitable>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const Waitable& from);
  void MergeFrom(const Waitable& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Waitable* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // string workflow_id = 3;
  void clear_workflow_id();
  static const int kWorkflowIdFieldNumber = 3;
  const ::std::string& workflow_id() const;
  void set_workflow_id(const ::std::string& value);
  #if LANG_CXX11
  void set_workflow_id(::std::string&& value);
  #endif
  void set_workflow_id(const char* value);
  void set_workflow_id(const char* value, size_t size);
  ::std::string* mutable_workflow_id();
  ::std::string* release_workflow_id();
  void set_allocated_workflow_id(::std::string* workflow_id);

  // .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
  bool has_wf_exec_id() const;
  void clear_wf_exec_id();
  static const int kWfExecIdFieldNumber = 1;
  const ::nebulaidl::core::WorkflowExecutionIdentifier& wf_exec_id() const;
  ::nebulaidl::core::WorkflowExecutionIdentifier* release_wf_exec_id();
  ::nebulaidl::core::WorkflowExecutionIdentifier* mutable_wf_exec_id();
  void set_allocated_wf_exec_id(::nebulaidl::core::WorkflowExecutionIdentifier* wf_exec_id);

  // .nebulaidl.core.WorkflowExecution.Phase phase = 2;
  void clear_phase();
  static const int kPhaseFieldNumber = 2;
  ::nebulaidl::core::WorkflowExecution_Phase phase() const;
  void set_phase(::nebulaidl::core::WorkflowExecution_Phase value);

  // @@protoc_insertion_point(class_scope:nebulaidl.plugins.Waitable)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::internal::ArenaStringPtr workflow_id_;
  ::nebulaidl::core::WorkflowExecutionIdentifier* wf_exec_id_;
  int phase_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_nebulaidl_2fplugins_2fwaitable_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Waitable

// .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
inline bool Waitable::has_wf_exec_id() const {
  return this != internal_default_instance() && wf_exec_id_ != nullptr;
}
inline const ::nebulaidl::core::WorkflowExecutionIdentifier& Waitable::wf_exec_id() const {
  const ::nebulaidl::core::WorkflowExecutionIdentifier* p = wf_exec_id_;
  // @@protoc_insertion_point(field_get:nebulaidl.plugins.Waitable.wf_exec_id)
  return p != nullptr ? *p : *reinterpret_cast<const ::nebulaidl::core::WorkflowExecutionIdentifier*>(
      &::nebulaidl::core::_WorkflowExecutionIdentifier_default_instance_);
}
inline ::nebulaidl::core::WorkflowExecutionIdentifier* Waitable::release_wf_exec_id() {
  // @@protoc_insertion_point(field_release:nebulaidl.plugins.Waitable.wf_exec_id)
  
  ::nebulaidl::core::WorkflowExecutionIdentifier* temp = wf_exec_id_;
  wf_exec_id_ = nullptr;
  return temp;
}
inline ::nebulaidl::core::WorkflowExecutionIdentifier* Waitable::mutable_wf_exec_id() {
  
  if (wf_exec_id_ == nullptr) {
    auto* p = CreateMaybeMessage<::nebulaidl::core::WorkflowExecutionIdentifier>(GetArenaNoVirtual());
    wf_exec_id_ = p;
  }
  // @@protoc_insertion_point(field_mutable:nebulaidl.plugins.Waitable.wf_exec_id)
  return wf_exec_id_;
}
inline void Waitable::set_allocated_wf_exec_id(::nebulaidl::core::WorkflowExecutionIdentifier* wf_exec_id) {
  ::google::protobuf::Arena* message_arena = GetArenaNoVirtual();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::google::protobuf::MessageLite*>(wf_exec_id_);
  }
  if (wf_exec_id) {
    ::google::protobuf::Arena* submessage_arena = nullptr;
    if (message_arena != submessage_arena) {
      wf_exec_id = ::google::protobuf::internal::GetOwnedMessage(
          message_arena, wf_exec_id, submessage_arena);
    }
    
  } else {
    
  }
  wf_exec_id_ = wf_exec_id;
  // @@protoc_insertion_point(field_set_allocated:nebulaidl.plugins.Waitable.wf_exec_id)
}

// .nebulaidl.core.WorkflowExecution.Phase phase = 2;
inline void Waitable::clear_phase() {
  phase_ = 0;
}
inline ::nebulaidl::core::WorkflowExecution_Phase Waitable::phase() const {
  // @@protoc_insertion_point(field_get:nebulaidl.plugins.Waitable.phase)
  return static_cast< ::nebulaidl::core::WorkflowExecution_Phase >(phase_);
}
inline void Waitable::set_phase(::nebulaidl::core::WorkflowExecution_Phase value) {
  
  phase_ = value;
  // @@protoc_insertion_point(field_set:nebulaidl.plugins.Waitable.phase)
}

// string workflow_id = 3;
inline void Waitable::clear_workflow_id() {
  workflow_id_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline const ::std::string& Waitable::workflow_id() const {
  // @@protoc_insertion_point(field_get:nebulaidl.plugins.Waitable.workflow_id)
  return workflow_id_.GetNoArena();
}
inline void Waitable::set_workflow_id(const ::std::string& value) {
  
  workflow_id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), value);
  // @@protoc_insertion_point(field_set:nebulaidl.plugins.Waitable.workflow_id)
}
#if LANG_CXX11
inline void Waitable::set_workflow_id(::std::string&& value) {
  
  workflow_id_.SetNoArena(
    &::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::move(value));
  // @@protoc_insertion_point(field_set_rvalue:nebulaidl.plugins.Waitable.workflow_id)
}
#endif
inline void Waitable::set_workflow_id(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  
  workflow_id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), ::std::string(value));
  // @@protoc_insertion_point(field_set_char:nebulaidl.plugins.Waitable.workflow_id)
}
inline void Waitable::set_workflow_id(const char* value, size_t size) {
  
  workflow_id_.SetNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(),
      ::std::string(reinterpret_cast<const char*>(value), size));
  // @@protoc_insertion_point(field_set_pointer:nebulaidl.plugins.Waitable.workflow_id)
}
inline ::std::string* Waitable::mutable_workflow_id() {
  
  // @@protoc_insertion_point(field_mutable:nebulaidl.plugins.Waitable.workflow_id)
  return workflow_id_.MutableNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline ::std::string* Waitable::release_workflow_id() {
  // @@protoc_insertion_point(field_release:nebulaidl.plugins.Waitable.workflow_id)
  
  return workflow_id_.ReleaseNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}
inline void Waitable::set_allocated_workflow_id(::std::string* workflow_id) {
  if (workflow_id != nullptr) {
    
  } else {
    
  }
  workflow_id_.SetAllocatedNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), workflow_id);
  // @@protoc_insertion_point(field_set_allocated:nebulaidl.plugins.Waitable.workflow_id)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace plugins
}  // namespace nebulaidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_nebulaidl_2fplugins_2fwaitable_2eproto