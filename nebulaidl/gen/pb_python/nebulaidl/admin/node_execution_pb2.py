# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/admin/node_execution.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.admin import common_pb2 as nebulaidl_dot_admin_dot_common__pb2
from nebulaidl.core import execution_pb2 as nebulaidl_dot_core_dot_execution__pb2
from nebulaidl.core import catalog_pb2 as nebulaidl_dot_core_dot_catalog__pb2
from nebulaidl.core import compiler_pb2 as nebulaidl_dot_core_dot_compiler__pb2
from nebulaidl.core import identifier_pb2 as nebulaidl_dot_core_dot_identifier__pb2
from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#nebulaidl/admin/node_execution.proto\x12\x0e\x66lyteidl.admin\x1a\x1b\x66lyteidl/admin/common.proto\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1b\x66lyteidl/core/catalog.proto\x1a\x1c\x66lyteidl/core/compiler.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"Q\n\x17NodeExecutionGetRequest\x12\x36\n\x02id\x18\x01 \x01(\x0b\x32&.nebulaidl.core.NodeExecutionIdentifierR\x02id\"\x99\x02\n\x18NodeExecutionListRequest\x12^\n\x15workflow_execution_id\x18\x01 \x01(\x0b\x32*.nebulaidl.core.WorkflowExecutionIdentifierR\x13workflowExecutionId\x12\x14\n\x05limit\x18\x02 \x01(\rR\x05limit\x12\x14\n\x05token\x18\x03 \x01(\tR\x05token\x12\x18\n\x07\x66ilters\x18\x04 \x01(\tR\x07\x66ilters\x12-\n\x07sort_by\x18\x05 \x01(\x0b\x32\x14.nebulaidl.admin.SortR\x06sortBy\x12(\n\x10unique_parent_id\x18\x06 \x01(\tR\x0euniqueParentId\"\xea\x01\n\x1fNodeExecutionForTaskListRequest\x12R\n\x11task_execution_id\x18\x01 \x01(\x0b\x32&.nebulaidl.core.TaskExecutionIdentifierR\x0ftaskExecutionId\x12\x14\n\x05limit\x18\x02 \x01(\rR\x05limit\x12\x14\n\x05token\x18\x03 \x01(\tR\x05token\x12\x18\n\x07\x66ilters\x18\x04 \x01(\tR\x07\x66ilters\x12-\n\x07sort_by\x18\x05 \x01(\x0b\x32\x14.nebulaidl.admin.SortR\x06sortBy\"\xe7\x01\n\rNodeExecution\x12\x36\n\x02id\x18\x01 \x01(\x0b\x32&.nebulaidl.core.NodeExecutionIdentifierR\x02id\x12\x1b\n\tinput_uri\x18\x02 \x01(\tR\x08inputUri\x12>\n\x07\x63losure\x18\x03 \x01(\x0b\x32$.nebulaidl.admin.NodeExecutionClosureR\x07\x63losure\x12\x41\n\x08metadata\x18\x04 \x01(\x0b\x32%.nebulaidl.admin.NodeExecutionMetaDataR\x08metadata\"\xba\x01\n\x15NodeExecutionMetaData\x12\x1f\n\x0bretry_group\x18\x01 \x01(\tR\nretryGroup\x12$\n\x0eis_parent_node\x18\x02 \x01(\x08R\x0cisParentNode\x12 \n\x0cspec_node_id\x18\x03 \x01(\tR\nspecNodeId\x12\x1d\n\nis_dynamic\x18\x04 \x01(\x08R\tisDynamic\x12\x19\n\x08is_array\x18\x05 \x01(\x08R\x07isArray\"q\n\x11NodeExecutionList\x12\x46\n\x0fnode_executions\x18\x01 \x03(\x0b\x32\x1d.nebulaidl.admin.NodeExecutionR\x0enodeExecutions\x12\x14\n\x05token\x18\x02 \x01(\tR\x05token\"\xf6\x05\n\x14NodeExecutionClosure\x12#\n\noutput_uri\x18\x01 \x01(\tB\x02\x18\x01H\x00R\toutputUri\x12\x35\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x1d.nebulaidl.core.ExecutionErrorH\x00R\x05\x65rror\x12@\n\x0boutput_data\x18\n \x01(\x0b\x32\x19.nebulaidl.core.LiteralMapB\x02\x18\x01H\x00R\noutputData\x12\x38\n\x05phase\x18\x03 \x01(\x0e\x32\".nebulaidl.core.NodeExecution.PhaseR\x05phase\x12\x39\n\nstarted_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12\x35\n\x08\x64uration\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x64uration\x12\x39\n\ncreated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\\\n\x16workflow_node_metadata\x18\x08 \x01(\x0b\x32$.nebulaidl.admin.WorkflowNodeMetadataH\x01R\x14workflowNodeMetadata\x12P\n\x12task_node_metadata\x18\t \x01(\x0b\x32 .nebulaidl.admin.TaskNodeMetadataH\x01R\x10taskNodeMetadata\x12\x19\n\x08\x64\x65\x63k_uri\x18\x0b \x01(\tR\x07\x64\x65\x63kUri\x12/\n\x14\x64ynamic_job_spec_uri\x18\x0c \x01(\tR\x11\x64ynamicJobSpecUriB\x0f\n\routput_resultB\x11\n\x0ftarget_metadata\"d\n\x14WorkflowNodeMetadata\x12L\n\x0b\x65xecutionId\x18\x01 \x01(\x0b\x32*.nebulaidl.core.WorkflowExecutionIdentifierR\x0b\x65xecutionId\"\xc0\x01\n\x10TaskNodeMetadata\x12\x44\n\x0c\x63\x61\x63he_status\x18\x01 \x01(\x0e\x32!.nebulaidl.core.CatalogCacheStatusR\x0b\x63\x61\x63heStatus\x12?\n\x0b\x63\x61talog_key\x18\x02 \x01(\x0b\x32\x1e.nebulaidl.core.CatalogMetadataR\ncatalogKey\x12%\n\x0e\x63heckpoint_uri\x18\x04 \x01(\tR\rcheckpointUri\"\xce\x01\n\x1b\x44ynamicWorkflowNodeMetadata\x12)\n\x02id\x18\x01 \x01(\x0b\x32\x19.nebulaidl.core.IdentifierR\x02id\x12S\n\x11\x63ompiled_workflow\x18\x02 \x01(\x0b\x32&.nebulaidl.core.CompiledWorkflowClosureR\x10\x63ompiledWorkflow\x12/\n\x14\x64ynamic_job_spec_uri\x18\x03 \x01(\tR\x11\x64ynamicJobSpecUri\"U\n\x1bNodeExecutionGetDataRequest\x12\x36\n\x02id\x18\x01 \x01(\x0b\x32&.nebulaidl.core.NodeExecutionIdentifierR\x02id\"\x96\x03\n\x1cNodeExecutionGetDataResponse\x12\x33\n\x06inputs\x18\x01 \x01(\x0b\x32\x17.nebulaidl.admin.UrlBlobB\x02\x18\x01R\x06inputs\x12\x35\n\x07outputs\x18\x02 \x01(\x0b\x32\x17.nebulaidl.admin.UrlBlobB\x02\x18\x01R\x07outputs\x12:\n\x0b\x66ull_inputs\x18\x03 \x01(\x0b\x32\x19.nebulaidl.core.LiteralMapR\nfullInputs\x12<\n\x0c\x66ull_outputs\x18\x04 \x01(\x0b\x32\x19.nebulaidl.core.LiteralMapR\x0b\x66ullOutputs\x12V\n\x10\x64ynamic_workflow\x18\x10 \x01(\x0b\x32+.nebulaidl.admin.DynamicWorkflowNodeMetadataR\x0f\x64ynamicWorkflow\x12\x38\n\nnebula_urls\x18\x11 \x01(\x0b\x32\x19.nebulaidl.admin.NebulaURLsR\tnebulaUrlsB\xbe\x01\n\x12\x63om.nebulaidl.adminB\x12NodeExecutionProtoP\x01Z;github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin\xa2\x02\x03\x46\x41X\xaa\x02\x0e\x46lyteidl.Admin\xca\x02\x0e\x46lyteidl\\Admin\xe2\x02\x1a\x46lyteidl\\Admin\\GPBMetadata\xea\x02\x0f\x46lyteidl::Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.admin.node_execution_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.nebulaidl.adminB\022NodeExecutionProtoP\001Z;github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin\242\002\003FAX\252\002\016Nebulaidl.Admin\312\002\016Nebulaidl\\Admin\342\002\032Nebulaidl\\Admin\\GPBMetadata\352\002\017Nebulaidl::Admin'
  _NODEEXECUTIONCLOSURE.fields_by_name['output_uri']._options = None
  _NODEEXECUTIONCLOSURE.fields_by_name['output_uri']._serialized_options = b'\030\001'
  _NODEEXECUTIONCLOSURE.fields_by_name['output_data']._options = None
  _NODEEXECUTIONCLOSURE.fields_by_name['output_data']._serialized_options = b'\030\001'
  _NODEEXECUTIONGETDATARESPONSE.fields_by_name['inputs']._options = None
  _NODEEXECUTIONGETDATARESPONSE.fields_by_name['inputs']._serialized_options = b'\030\001'
  _NODEEXECUTIONGETDATARESPONSE.fields_by_name['outputs']._options = None
  _NODEEXECUTIONGETDATARESPONSE.fields_by_name['outputs']._serialized_options = b'\030\001'
  _globals['_NODEEXECUTIONGETREQUEST']._serialized_start=301
  _globals['_NODEEXECUTIONGETREQUEST']._serialized_end=382
  _globals['_NODEEXECUTIONLISTREQUEST']._serialized_start=385
  _globals['_NODEEXECUTIONLISTREQUEST']._serialized_end=666
  _globals['_NODEEXECUTIONFORTASKLISTREQUEST']._serialized_start=669
  _globals['_NODEEXECUTIONFORTASKLISTREQUEST']._serialized_end=903
  _globals['_NODEEXECUTION']._serialized_start=906
  _globals['_NODEEXECUTION']._serialized_end=1137
  _globals['_NODEEXECUTIONMETADATA']._serialized_start=1140
  _globals['_NODEEXECUTIONMETADATA']._serialized_end=1326
  _globals['_NODEEXECUTIONLIST']._serialized_start=1328
  _globals['_NODEEXECUTIONLIST']._serialized_end=1441
  _globals['_NODEEXECUTIONCLOSURE']._serialized_start=1444
  _globals['_NODEEXECUTIONCLOSURE']._serialized_end=2202
  _globals['_WORKFLOWNODEMETADATA']._serialized_start=2204
  _globals['_WORKFLOWNODEMETADATA']._serialized_end=2304
  _globals['_TASKNODEMETADATA']._serialized_start=2307
  _globals['_TASKNODEMETADATA']._serialized_end=2499
  _globals['_DYNAMICWORKFLOWNODEMETADATA']._serialized_start=2502
  _globals['_DYNAMICWORKFLOWNODEMETADATA']._serialized_end=2708
  _globals['_NODEEXECUTIONGETDATAREQUEST']._serialized_start=2710
  _globals['_NODEEXECUTIONGETDATAREQUEST']._serialized_end=2795
  _globals['_NODEEXECUTIONGETDATARESPONSE']._serialized_start=2798
  _globals['_NODEEXECUTIONGETDATARESPONSE']._serialized_end=3204
# @@protoc_insertion_point(module_scope)
