# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/plugins/ray.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1a\x66lyteidl/plugins/ray.proto\x12\x10\x66lyteidl.plugins\"h\n\x06RayJob\x12=\n\x0bray_cluster\x18\x01 \x01(\x0b\x32\x1c.nebulaidl.plugins.RayClusterR\nrayCluster\x12\x1f\n\x0bruntime_env\x18\x02 \x01(\tR\nruntimeEnv\"\xa4\x01\n\nRayCluster\x12G\n\x0fhead_group_spec\x18\x01 \x01(\x0b\x32\x1f.nebulaidl.plugins.HeadGroupSpecR\rheadGroupSpec\x12M\n\x11worker_group_spec\x18\x02 \x03(\x0b\x32!.nebulaidl.plugins.WorkerGroupSpecR\x0fworkerGroupSpec\"\xb1\x01\n\rHeadGroupSpec\x12]\n\x10ray_start_params\x18\x01 \x03(\x0b\x32\x33.nebulaidl.plugins.HeadGroupSpec.RayStartParamsEntryR\x0erayStartParams\x1a\x41\n\x13RayStartParamsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\xb6\x02\n\x0fWorkerGroupSpec\x12\x1d\n\ngroup_name\x18\x01 \x01(\tR\tgroupName\x12\x1a\n\x08replicas\x18\x02 \x01(\x05R\x08replicas\x12!\n\x0cmin_replicas\x18\x03 \x01(\x05R\x0bminReplicas\x12!\n\x0cmax_replicas\x18\x04 \x01(\x05R\x0bmaxReplicas\x12_\n\x10ray_start_params\x18\x05 \x03(\x0b\x32\x35.nebulaidl.plugins.WorkerGroupSpec.RayStartParamsEntryR\x0erayStartParams\x1a\x41\n\x13RayStartParamsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x42\xc0\x01\n\x14\x63om.nebulaidl.pluginsB\x08RayProtoP\x01Z=github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/plugins\xa2\x02\x03\x46PX\xaa\x02\x10\x46lyteidl.Plugins\xca\x02\x10\x46lyteidl\\Plugins\xe2\x02\x1c\x46lyteidl\\Plugins\\GPBMetadata\xea\x02\x11\x46lyteidl::Pluginsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.plugins.ray_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024com.nebulaidl.pluginsB\010RayProtoP\001Z=github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/plugins\242\002\003FPX\252\002\020Nebulaidl.Plugins\312\002\020Nebulaidl\\Plugins\342\002\034Nebulaidl\\Plugins\\GPBMetadata\352\002\021Nebulaidl::Plugins'
  _HEADGROUPSPEC_RAYSTARTPARAMSENTRY._options = None
  _HEADGROUPSPEC_RAYSTARTPARAMSENTRY._serialized_options = b'8\001'
  _WORKERGROUPSPEC_RAYSTARTPARAMSENTRY._options = None
  _WORKERGROUPSPEC_RAYSTARTPARAMSENTRY._serialized_options = b'8\001'
  _globals['_RAYJOB']._serialized_start=48
  _globals['_RAYJOB']._serialized_end=152
  _globals['_RAYCLUSTER']._serialized_start=155
  _globals['_RAYCLUSTER']._serialized_end=319
  _globals['_HEADGROUPSPEC']._serialized_start=322
  _globals['_HEADGROUPSPEC']._serialized_end=499
  _globals['_HEADGROUPSPEC_RAYSTARTPARAMSENTRY']._serialized_start=434
  _globals['_HEADGROUPSPEC_RAYSTARTPARAMSENTRY']._serialized_end=499
  _globals['_WORKERGROUPSPEC']._serialized_start=502
  _globals['_WORKERGROUPSPEC']._serialized_end=812
  _globals['_WORKERGROUPSPEC_RAYSTARTPARAMSENTRY']._serialized_start=434
  _globals['_WORKERGROUPSPEC_RAYSTARTPARAMSENTRY']._serialized_end=499
# @@protoc_insertion_point(module_scope)