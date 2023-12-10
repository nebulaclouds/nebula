# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/core/compiler.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import workflow_pb2 as nebulaidl_dot_core_dot_workflow__pb2
from nebulaidl.core import tasks_pb2 as nebulaidl_dot_core_dot_tasks__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x66lyteidl/core/compiler.proto\x12\rnebulaidl.core\x1a\x1c\x66lyteidl/core/workflow.proto\x1a\x19\x66lyteidl/core/tasks.proto\"\x87\x03\n\rConnectionSet\x12L\n\ndownstream\x18\x07 \x03(\x0b\x32,.nebulaidl.core.ConnectionSet.DownstreamEntryR\ndownstream\x12\x46\n\x08upstream\x18\x08 \x03(\x0b\x32*.nebulaidl.core.ConnectionSet.UpstreamEntryR\x08upstream\x1a\x1a\n\x06IdList\x12\x10\n\x03ids\x18\x01 \x03(\tR\x03ids\x1a\x62\n\x0f\x44ownstreamEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x39\n\x05value\x18\x02 \x01(\x0b\x32#.nebulaidl.core.ConnectionSet.IdListR\x05value:\x02\x38\x01\x1a`\n\rUpstreamEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x39\n\x05value\x18\x02 \x01(\x0b\x32#.nebulaidl.core.ConnectionSet.IdListR\x05value:\x02\x38\x01\"\x8f\x01\n\x10\x43ompiledWorkflow\x12;\n\x08template\x18\x01 \x01(\x0b\x32\x1f.nebulaidl.core.WorkflowTemplateR\x08template\x12>\n\x0b\x63onnections\x18\x02 \x01(\x0b\x32\x1c.nebulaidl.core.ConnectionSetR\x0b\x63onnections\"G\n\x0c\x43ompiledTask\x12\x37\n\x08template\x18\x01 \x01(\x0b\x32\x1b.nebulaidl.core.TaskTemplateR\x08template\"\xcd\x01\n\x17\x43ompiledWorkflowClosure\x12\x39\n\x07primary\x18\x01 \x01(\x0b\x32\x1f.nebulaidl.core.CompiledWorkflowR\x07primary\x12\x44\n\rsub_workflows\x18\x02 \x03(\x0b\x32\x1f.nebulaidl.core.CompiledWorkflowR\x0csubWorkflows\x12\x31\n\x05tasks\x18\x03 \x03(\x0b\x32\x1b.nebulaidl.core.CompiledTaskR\x05tasksB\xb3\x01\n\x11\x63om.nebulaidl.coreB\rCompilerProtoP\x01Z:github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core\xa2\x02\x03\x46\x43X\xaa\x02\rNebulaidl.Core\xca\x02\rNebulaidl\\Core\xe2\x02\x19\x46lyteidl\\Core\\GPBMetadata\xea\x02\x0e\x46lyteidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.core.compiler_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.nebulaidl.coreB\rCompilerProtoP\001Z:github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core\242\002\003FCX\252\002\rNebulaidl.Core\312\002\rNebulaidl\\Core\342\002\031Nebulaidl\\Core\\GPBMetadata\352\002\016Nebulaidl::Core'
  _CONNECTIONSET_DOWNSTREAMENTRY._options = None
  _CONNECTIONSET_DOWNSTREAMENTRY._serialized_options = b'8\001'
  _CONNECTIONSET_UPSTREAMENTRY._options = None
  _CONNECTIONSET_UPSTREAMENTRY._serialized_options = b'8\001'
  _globals['_CONNECTIONSET']._serialized_start=105
  _globals['_CONNECTIONSET']._serialized_end=496
  _globals['_CONNECTIONSET_IDLIST']._serialized_start=272
  _globals['_CONNECTIONSET_IDLIST']._serialized_end=298
  _globals['_CONNECTIONSET_DOWNSTREAMENTRY']._serialized_start=300
  _globals['_CONNECTIONSET_DOWNSTREAMENTRY']._serialized_end=398
  _globals['_CONNECTIONSET_UPSTREAMENTRY']._serialized_start=400
  _globals['_CONNECTIONSET_UPSTREAMENTRY']._serialized_end=496
  _globals['_COMPILEDWORKFLOW']._serialized_start=499
  _globals['_COMPILEDWORKFLOW']._serialized_end=642
  _globals['_COMPILEDTASK']._serialized_start=644
  _globals['_COMPILEDTASK']._serialized_end=715
  _globals['_COMPILEDWORKFLOWCLOSURE']._serialized_start=718
  _globals['_COMPILEDWORKFLOWCLOSURE']._serialized_end=923
# @@protoc_insertion_point(module_scope)
