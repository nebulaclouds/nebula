# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/admin/project_domain_attributes.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.admin import matchable_resource_pb2 as nebulaidl_dot_admin_dot_matchable__resource__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.nebulaidl/admin/project_domain_attributes.proto\x12\x0e\x66lyteidl.admin\x1a\'nebulaidl/admin/matchable_resource.proto\"\xa0\x01\n\x17ProjectDomainAttributes\x12\x18\n\x07project\x18\x01 \x01(\tR\x07project\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12S\n\x13matching_attributes\x18\x03 \x01(\x0b\x32\".nebulaidl.admin.MatchingAttributesR\x12matchingAttributes\"o\n$ProjectDomainAttributesUpdateRequest\x12G\n\nattributes\x18\x01 \x01(\x0b\x32\'.nebulaidl.admin.ProjectDomainAttributesR\nattributes\"\'\n%ProjectDomainAttributesUpdateResponse\"\x9d\x01\n!ProjectDomainAttributesGetRequest\x12\x18\n\x07project\x18\x01 \x01(\tR\x07project\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12\x46\n\rresource_type\x18\x03 \x01(\x0e\x32!.nebulaidl.admin.MatchableResourceR\x0cresourceType\"m\n\"ProjectDomainAttributesGetResponse\x12G\n\nattributes\x18\x01 \x01(\x0b\x32\'.nebulaidl.admin.ProjectDomainAttributesR\nattributes\"\xa0\x01\n$ProjectDomainAttributesDeleteRequest\x12\x18\n\x07project\x18\x01 \x01(\tR\x07project\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12\x46\n\rresource_type\x18\x03 \x01(\x0e\x32!.nebulaidl.admin.MatchableResourceR\x0cresourceType\"\'\n%ProjectDomainAttributesDeleteResponseB\xc8\x01\n\x12\x63om.nebulaidl.adminB\x1cProjectDomainAttributesProtoP\x01Z;github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin\xa2\x02\x03\x46\x41X\xaa\x02\x0e\x46lyteidl.Admin\xca\x02\x0e\x46lyteidl\\Admin\xe2\x02\x1a\x46lyteidl\\Admin\\GPBMetadata\xea\x02\x0f\x46lyteidl::Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.admin.project_domain_attributes_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.nebulaidl.adminB\034ProjectDomainAttributesProtoP\001Z;github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin\242\002\003FAX\252\002\016Nebulaidl.Admin\312\002\016Nebulaidl\\Admin\342\002\032Nebulaidl\\Admin\\GPBMetadata\352\002\017Nebulaidl::Admin'
  _globals['_PROJECTDOMAINATTRIBUTES']._serialized_start=108
  _globals['_PROJECTDOMAINATTRIBUTES']._serialized_end=268
  _globals['_PROJECTDOMAINATTRIBUTESUPDATEREQUEST']._serialized_start=270
  _globals['_PROJECTDOMAINATTRIBUTESUPDATEREQUEST']._serialized_end=381
  _globals['_PROJECTDOMAINATTRIBUTESUPDATERESPONSE']._serialized_start=383
  _globals['_PROJECTDOMAINATTRIBUTESUPDATERESPONSE']._serialized_end=422
  _globals['_PROJECTDOMAINATTRIBUTESGETREQUEST']._serialized_start=425
  _globals['_PROJECTDOMAINATTRIBUTESGETREQUEST']._serialized_end=582
  _globals['_PROJECTDOMAINATTRIBUTESGETRESPONSE']._serialized_start=584
  _globals['_PROJECTDOMAINATTRIBUTESGETRESPONSE']._serialized_end=693
  _globals['_PROJECTDOMAINATTRIBUTESDELETEREQUEST']._serialized_start=696
  _globals['_PROJECTDOMAINATTRIBUTESDELETEREQUEST']._serialized_end=856
  _globals['_PROJECTDOMAINATTRIBUTESDELETERESPONSE']._serialized_start=858
  _globals['_PROJECTDOMAINATTRIBUTESDELETERESPONSE']._serialized_end=897
# @@protoc_insertion_point(module_scope)
