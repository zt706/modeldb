# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/public/uac/Collaborator.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from ...public.uac import UACService_pb2 as protos_dot_public_dot_uac_dot_UACService__pb2
from ...public.uac import Organization_pb2 as protos_dot_public_dot_uac_dot_Organization__pb2
from ...public.uac import Team_pb2 as protos_dot_public_dot_uac_dot_Team__pb2
from ...public.uac import RoleService_pb2 as protos_dot_public_dot_uac_dot_RoleService__pb2
from ...public.common import CommonService_pb2 as protos_dot_public_dot_common_dot_CommonService__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/public/uac/Collaborator.proto',
  package='ai.verta.uac',
  syntax='proto3',
  serialized_options=b'P\001Z:github.com/VertaAI/modeldb/protos/gen/go/protos/public/uac',
  serialized_pb=b'\n$protos/public/uac/Collaborator.proto\x12\x0c\x61i.verta.uac\x1a\x1cgoogle/api/annotations.proto\x1a\"protos/public/uac/UACService.proto\x1a$protos/public/uac/Organization.proto\x1a\x1cprotos/public/uac/Team.proto\x1a#protos/public/uac/RoleService.proto\x1a(protos/public/common/CommonService.proto\"\xdc\x04\n\x16\x41\x64\x64\x43ollaboratorRequest\x12\x12\n\nentity_ids\x18\x01 \x03(\t\x12\x12\n\nshare_with\x18\x02 \x01(\t\x12Q\n\x11\x63ollaborator_type\x18\x03 \x01(\x0e\x32\x36.ai.verta.common.CollaboratorTypeEnum.CollaboratorType\x12\x0f\n\x07message\x18\x04 \x01(\t\x12\x14\n\x0c\x64\x61te_created\x18\x05 \x01(\x04\x12\x14\n\x0c\x64\x61te_updated\x18\x06 \x01(\x04\x12\x38\n\ncan_deploy\x18\x07 \x01(\x0e\x32$.ai.verta.common.TernaryEnum.Ternary\x12\x43\n\x11\x61uthz_entity_type\x18\x08 \x01(\x0e\x32(.ai.verta.uac.EntitiesEnum.EntitiesTypes\x1a\x8a\x02\n\x08Response\x12\x32\n\x14self_allowed_actions\x18\x05 \x03(\x0b\x32\x14.ai.verta.uac.Action\x12\x0e\n\x06status\x18\x01 \x01(\x08\x12\x38\n\x16\x63ollaborator_user_info\x18\x02 \x01(\x0b\x32\x16.ai.verta.uac.UserInfoH\x00\x12?\n\x19\x63ollaborator_organization\x18\x03 \x01(\x0b\x32\x1a.ai.verta.uac.OrganizationH\x00\x12/\n\x11\x63ollaborator_team\x18\x04 \x01(\x0b\x32\x12.ai.verta.uac.TeamH\x00\x42\x0e\n\x0c\x63ollaborator\"\xe6\x01\n\x12RemoveCollaborator\x12\x11\n\tentity_id\x18\x01 \x01(\t\x12\x12\n\nshare_with\x18\x02 \x01(\t\x12\x14\n\x0c\x64\x61te_deleted\x18\x03 \x01(\x04\x12\x43\n\x11\x61uthz_entity_type\x18\x04 \x01(\x0e\x32(.ai.verta.uac.EntitiesEnum.EntitiesTypes\x1aN\n\x08Response\x12\x0e\n\x06status\x18\x01 \x01(\x08\x12\x32\n\x14self_allowed_actions\x18\x05 \x03(\x0b\x32\x14.ai.verta.uac.Action\"\xc6\x02\n\x17GetCollaboratorResponse\x12\x13\n\x07user_id\x18\x01 \x01(\tB\x02\x18\x01\x12Q\n\x11\x63ollaborator_type\x18\x02 \x01(\x0e\x32\x36.ai.verta.common.CollaboratorTypeEnum.CollaboratorType\x12\x32\n\x0eshare_via_type\x18\x03 \x01(\x0e\x32\x1a.ai.verta.uac.ShareViaEnum\x12\x10\n\x08verta_id\x18\x04 \x01(\t\x12\x38\n\ncan_deploy\x18\x05 \x01(\x0e\x32$.ai.verta.common.TernaryEnum.Ternary\x12\x43\n\x11\x61uthz_entity_type\x18\x06 \x01(\x0e\x32(.ai.verta.uac.EntitiesEnum.EntitiesTypes\"m\n\x0fGetCollaborator\x12\x11\n\tentity_id\x18\x01 \x01(\t\x1aG\n\x08Response\x12;\n\x0cshared_users\x18\x01 \x03(\x0b\x32%.ai.verta.uac.GetCollaboratorResponse*7\n\x0cShareViaEnum\x12\x0b\n\x07USER_ID\x10\x00\x12\x0c\n\x08\x45MAIL_ID\x10\x01\x12\x0c\n\x08USERNAME\x10\x02\x32\xdf\x0b\n\x13\x43ollaboratorService\x12\xb1\x01\n\x1e\x61\x64\x64OrUpdateProjectCollaborator\x12$.ai.verta.uac.AddCollaboratorRequest\x1a-.ai.verta.uac.AddCollaboratorRequest.Response\":\x82\xd3\xe4\x93\x02\x34\"//v1/collaborator/addOrUpdateProjectCollaborator:\x01*\x12\x9c\x01\n\x19removeProjectCollaborator\x12 .ai.verta.uac.RemoveCollaborator\x1a).ai.verta.uac.RemoveCollaborator.Response\"2\x82\xd3\xe4\x93\x02,**/v1/collaborator/removeProjectCollaborator\x12\x92\x01\n\x17getProjectCollaborators\x12\x1d.ai.verta.uac.GetCollaborator\x1a&.ai.verta.uac.GetCollaborator.Response\"0\x82\xd3\xe4\x93\x02*\x12(/v1/collaborator/getProjectCollaborators\x12\xb1\x01\n\x1e\x61\x64\x64OrUpdateDatasetCollaborator\x12$.ai.verta.uac.AddCollaboratorRequest\x1a-.ai.verta.uac.AddCollaboratorRequest.Response\":\x82\xd3\xe4\x93\x02\x34\"//v1/collaborator/addOrUpdateDatasetCollaborator:\x01*\x12\x9c\x01\n\x19removeDatasetCollaborator\x12 .ai.verta.uac.RemoveCollaborator\x1a).ai.verta.uac.RemoveCollaborator.Response\"2\x82\xd3\xe4\x93\x02,**/v1/collaborator/removeDatasetCollaborator\x12\x92\x01\n\x17getDatasetCollaborators\x12\x1d.ai.verta.uac.GetCollaborator\x1a&.ai.verta.uac.GetCollaborator.Response\"0\x82\xd3\xe4\x93\x02*\x12(/v1/collaborator/getDatasetCollaborators\x12\xb7\x01\n!addOrUpdateRepositoryCollaborator\x12$.ai.verta.uac.AddCollaboratorRequest\x1a-.ai.verta.uac.AddCollaboratorRequest.Response\"=\x82\xd3\xe4\x93\x02\x37\"2/v1/collaborator/addOrUpdateRepositoryCollaborator:\x01*\x12\xa2\x01\n\x1cremoveRepositoryCollaborator\x12 .ai.verta.uac.RemoveCollaborator\x1a).ai.verta.uac.RemoveCollaborator.Response\"5\x82\xd3\xe4\x93\x02/*-/v1/collaborator/removeRepositoryCollaborator\x12\x98\x01\n\x1agetRepositoryCollaborators\x12\x1d.ai.verta.uac.GetCollaborator\x1a&.ai.verta.uac.GetCollaborator.Response\"3\x82\xd3\xe4\x93\x02-\x12+/v1/collaborator/getRepositoryCollaboratorsB>P\x01Z:github.com/VertaAI/modeldb/protos/gen/go/protos/public/uacb\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,protos_dot_public_dot_uac_dot_UACService__pb2.DESCRIPTOR,protos_dot_public_dot_uac_dot_Organization__pb2.DESCRIPTOR,protos_dot_public_dot_uac_dot_Team__pb2.DESCRIPTOR,protos_dot_public_dot_uac_dot_RoleService__pb2.DESCRIPTOR,protos_dot_public_dot_common_dot_CommonService__pb2.DESCRIPTOR,])

_SHAREVIAENUM = _descriptor.EnumDescriptor(
  name='ShareViaEnum',
  full_name='ai.verta.uac.ShareViaEnum',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='USER_ID', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='EMAIL_ID', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='USERNAME', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1547,
  serialized_end=1602,
)
_sym_db.RegisterEnumDescriptor(_SHAREVIAENUM)

ShareViaEnum = enum_type_wrapper.EnumTypeWrapper(_SHAREVIAENUM)
USER_ID = 0
EMAIL_ID = 1
USERNAME = 2



_ADDCOLLABORATORREQUEST_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='ai.verta.uac.AddCollaboratorRequest.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='self_allowed_actions', full_name='ai.verta.uac.AddCollaboratorRequest.Response.self_allowed_actions', index=0,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='ai.verta.uac.AddCollaboratorRequest.Response.status', index=1,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collaborator_user_info', full_name='ai.verta.uac.AddCollaboratorRequest.Response.collaborator_user_info', index=2,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collaborator_organization', full_name='ai.verta.uac.AddCollaboratorRequest.Response.collaborator_organization', index=3,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collaborator_team', full_name='ai.verta.uac.AddCollaboratorRequest.Response.collaborator_team', index=4,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='collaborator', full_name='ai.verta.uac.AddCollaboratorRequest.Response.collaborator',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=606,
  serialized_end=872,
)

_ADDCOLLABORATORREQUEST = _descriptor.Descriptor(
  name='AddCollaboratorRequest',
  full_name='ai.verta.uac.AddCollaboratorRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity_ids', full_name='ai.verta.uac.AddCollaboratorRequest.entity_ids', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='share_with', full_name='ai.verta.uac.AddCollaboratorRequest.share_with', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collaborator_type', full_name='ai.verta.uac.AddCollaboratorRequest.collaborator_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='ai.verta.uac.AddCollaboratorRequest.message', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='date_created', full_name='ai.verta.uac.AddCollaboratorRequest.date_created', index=4,
      number=5, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='date_updated', full_name='ai.verta.uac.AddCollaboratorRequest.date_updated', index=5,
      number=6, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='can_deploy', full_name='ai.verta.uac.AddCollaboratorRequest.can_deploy', index=6,
      number=7, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='authz_entity_type', full_name='ai.verta.uac.AddCollaboratorRequest.authz_entity_type', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_ADDCOLLABORATORREQUEST_RESPONSE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=268,
  serialized_end=872,
)


_REMOVECOLLABORATOR_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='ai.verta.uac.RemoveCollaborator.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ai.verta.uac.RemoveCollaborator.Response.status', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='self_allowed_actions', full_name='ai.verta.uac.RemoveCollaborator.Response.self_allowed_actions', index=1,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1027,
  serialized_end=1105,
)

_REMOVECOLLABORATOR = _descriptor.Descriptor(
  name='RemoveCollaborator',
  full_name='ai.verta.uac.RemoveCollaborator',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity_id', full_name='ai.verta.uac.RemoveCollaborator.entity_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='share_with', full_name='ai.verta.uac.RemoveCollaborator.share_with', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='date_deleted', full_name='ai.verta.uac.RemoveCollaborator.date_deleted', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='authz_entity_type', full_name='ai.verta.uac.RemoveCollaborator.authz_entity_type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_REMOVECOLLABORATOR_RESPONSE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=875,
  serialized_end=1105,
)


_GETCOLLABORATORRESPONSE = _descriptor.Descriptor(
  name='GetCollaboratorResponse',
  full_name='ai.verta.uac.GetCollaboratorResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='ai.verta.uac.GetCollaboratorResponse.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\030\001', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='collaborator_type', full_name='ai.verta.uac.GetCollaboratorResponse.collaborator_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='share_via_type', full_name='ai.verta.uac.GetCollaboratorResponse.share_via_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='verta_id', full_name='ai.verta.uac.GetCollaboratorResponse.verta_id', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='can_deploy', full_name='ai.verta.uac.GetCollaboratorResponse.can_deploy', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='authz_entity_type', full_name='ai.verta.uac.GetCollaboratorResponse.authz_entity_type', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1108,
  serialized_end=1434,
)


_GETCOLLABORATOR_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='ai.verta.uac.GetCollaborator.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='shared_users', full_name='ai.verta.uac.GetCollaborator.Response.shared_users', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1474,
  serialized_end=1545,
)

_GETCOLLABORATOR = _descriptor.Descriptor(
  name='GetCollaborator',
  full_name='ai.verta.uac.GetCollaborator',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity_id', full_name='ai.verta.uac.GetCollaborator.entity_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_GETCOLLABORATOR_RESPONSE, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1436,
  serialized_end=1545,
)

_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['self_allowed_actions'].message_type = protos_dot_public_dot_uac_dot_RoleService__pb2._ACTION
_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_user_info'].message_type = protos_dot_public_dot_uac_dot_UACService__pb2._USERINFO
_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_organization'].message_type = protos_dot_public_dot_uac_dot_Organization__pb2._ORGANIZATION
_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_team'].message_type = protos_dot_public_dot_uac_dot_Team__pb2._TEAM
_ADDCOLLABORATORREQUEST_RESPONSE.containing_type = _ADDCOLLABORATORREQUEST
_ADDCOLLABORATORREQUEST_RESPONSE.oneofs_by_name['collaborator'].fields.append(
  _ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_user_info'])
_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_user_info'].containing_oneof = _ADDCOLLABORATORREQUEST_RESPONSE.oneofs_by_name['collaborator']
_ADDCOLLABORATORREQUEST_RESPONSE.oneofs_by_name['collaborator'].fields.append(
  _ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_organization'])
_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_organization'].containing_oneof = _ADDCOLLABORATORREQUEST_RESPONSE.oneofs_by_name['collaborator']
_ADDCOLLABORATORREQUEST_RESPONSE.oneofs_by_name['collaborator'].fields.append(
  _ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_team'])
_ADDCOLLABORATORREQUEST_RESPONSE.fields_by_name['collaborator_team'].containing_oneof = _ADDCOLLABORATORREQUEST_RESPONSE.oneofs_by_name['collaborator']
_ADDCOLLABORATORREQUEST.fields_by_name['collaborator_type'].enum_type = protos_dot_public_dot_common_dot_CommonService__pb2._COLLABORATORTYPEENUM_COLLABORATORTYPE
_ADDCOLLABORATORREQUEST.fields_by_name['can_deploy'].enum_type = protos_dot_public_dot_common_dot_CommonService__pb2._TERNARYENUM_TERNARY
_ADDCOLLABORATORREQUEST.fields_by_name['authz_entity_type'].enum_type = protos_dot_public_dot_uac_dot_RoleService__pb2._ENTITIESENUM_ENTITIESTYPES
_REMOVECOLLABORATOR_RESPONSE.fields_by_name['self_allowed_actions'].message_type = protos_dot_public_dot_uac_dot_RoleService__pb2._ACTION
_REMOVECOLLABORATOR_RESPONSE.containing_type = _REMOVECOLLABORATOR
_REMOVECOLLABORATOR.fields_by_name['authz_entity_type'].enum_type = protos_dot_public_dot_uac_dot_RoleService__pb2._ENTITIESENUM_ENTITIESTYPES
_GETCOLLABORATORRESPONSE.fields_by_name['collaborator_type'].enum_type = protos_dot_public_dot_common_dot_CommonService__pb2._COLLABORATORTYPEENUM_COLLABORATORTYPE
_GETCOLLABORATORRESPONSE.fields_by_name['share_via_type'].enum_type = _SHAREVIAENUM
_GETCOLLABORATORRESPONSE.fields_by_name['can_deploy'].enum_type = protos_dot_public_dot_common_dot_CommonService__pb2._TERNARYENUM_TERNARY
_GETCOLLABORATORRESPONSE.fields_by_name['authz_entity_type'].enum_type = protos_dot_public_dot_uac_dot_RoleService__pb2._ENTITIESENUM_ENTITIESTYPES
_GETCOLLABORATOR_RESPONSE.fields_by_name['shared_users'].message_type = _GETCOLLABORATORRESPONSE
_GETCOLLABORATOR_RESPONSE.containing_type = _GETCOLLABORATOR
DESCRIPTOR.message_types_by_name['AddCollaboratorRequest'] = _ADDCOLLABORATORREQUEST
DESCRIPTOR.message_types_by_name['RemoveCollaborator'] = _REMOVECOLLABORATOR
DESCRIPTOR.message_types_by_name['GetCollaboratorResponse'] = _GETCOLLABORATORRESPONSE
DESCRIPTOR.message_types_by_name['GetCollaborator'] = _GETCOLLABORATOR
DESCRIPTOR.enum_types_by_name['ShareViaEnum'] = _SHAREVIAENUM
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AddCollaboratorRequest = _reflection.GeneratedProtocolMessageType('AddCollaboratorRequest', (_message.Message,), {

  'Response' : _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
    'DESCRIPTOR' : _ADDCOLLABORATORREQUEST_RESPONSE,
    '__module__' : 'protos.public.uac.Collaborator_pb2'
    # @@protoc_insertion_point(class_scope:ai.verta.uac.AddCollaboratorRequest.Response)
    })
  ,
  'DESCRIPTOR' : _ADDCOLLABORATORREQUEST,
  '__module__' : 'protos.public.uac.Collaborator_pb2'
  # @@protoc_insertion_point(class_scope:ai.verta.uac.AddCollaboratorRequest)
  })
_sym_db.RegisterMessage(AddCollaboratorRequest)
_sym_db.RegisterMessage(AddCollaboratorRequest.Response)

RemoveCollaborator = _reflection.GeneratedProtocolMessageType('RemoveCollaborator', (_message.Message,), {

  'Response' : _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
    'DESCRIPTOR' : _REMOVECOLLABORATOR_RESPONSE,
    '__module__' : 'protos.public.uac.Collaborator_pb2'
    # @@protoc_insertion_point(class_scope:ai.verta.uac.RemoveCollaborator.Response)
    })
  ,
  'DESCRIPTOR' : _REMOVECOLLABORATOR,
  '__module__' : 'protos.public.uac.Collaborator_pb2'
  # @@protoc_insertion_point(class_scope:ai.verta.uac.RemoveCollaborator)
  })
_sym_db.RegisterMessage(RemoveCollaborator)
_sym_db.RegisterMessage(RemoveCollaborator.Response)

GetCollaboratorResponse = _reflection.GeneratedProtocolMessageType('GetCollaboratorResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETCOLLABORATORRESPONSE,
  '__module__' : 'protos.public.uac.Collaborator_pb2'
  # @@protoc_insertion_point(class_scope:ai.verta.uac.GetCollaboratorResponse)
  })
_sym_db.RegisterMessage(GetCollaboratorResponse)

GetCollaborator = _reflection.GeneratedProtocolMessageType('GetCollaborator', (_message.Message,), {

  'Response' : _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
    'DESCRIPTOR' : _GETCOLLABORATOR_RESPONSE,
    '__module__' : 'protos.public.uac.Collaborator_pb2'
    # @@protoc_insertion_point(class_scope:ai.verta.uac.GetCollaborator.Response)
    })
  ,
  'DESCRIPTOR' : _GETCOLLABORATOR,
  '__module__' : 'protos.public.uac.Collaborator_pb2'
  # @@protoc_insertion_point(class_scope:ai.verta.uac.GetCollaborator)
  })
_sym_db.RegisterMessage(GetCollaborator)
_sym_db.RegisterMessage(GetCollaborator.Response)


DESCRIPTOR._options = None
_GETCOLLABORATORRESPONSE.fields_by_name['user_id']._options = None

_COLLABORATORSERVICE = _descriptor.ServiceDescriptor(
  name='CollaboratorService',
  full_name='ai.verta.uac.CollaboratorService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1605,
  serialized_end=3108,
  methods=[
  _descriptor.MethodDescriptor(
    name='addOrUpdateProjectCollaborator',
    full_name='ai.verta.uac.CollaboratorService.addOrUpdateProjectCollaborator',
    index=0,
    containing_service=None,
    input_type=_ADDCOLLABORATORREQUEST,
    output_type=_ADDCOLLABORATORREQUEST_RESPONSE,
    serialized_options=b'\202\323\344\223\0024\"//v1/collaborator/addOrUpdateProjectCollaborator:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='removeProjectCollaborator',
    full_name='ai.verta.uac.CollaboratorService.removeProjectCollaborator',
    index=1,
    containing_service=None,
    input_type=_REMOVECOLLABORATOR,
    output_type=_REMOVECOLLABORATOR_RESPONSE,
    serialized_options=b'\202\323\344\223\002,**/v1/collaborator/removeProjectCollaborator',
  ),
  _descriptor.MethodDescriptor(
    name='getProjectCollaborators',
    full_name='ai.verta.uac.CollaboratorService.getProjectCollaborators',
    index=2,
    containing_service=None,
    input_type=_GETCOLLABORATOR,
    output_type=_GETCOLLABORATOR_RESPONSE,
    serialized_options=b'\202\323\344\223\002*\022(/v1/collaborator/getProjectCollaborators',
  ),
  _descriptor.MethodDescriptor(
    name='addOrUpdateDatasetCollaborator',
    full_name='ai.verta.uac.CollaboratorService.addOrUpdateDatasetCollaborator',
    index=3,
    containing_service=None,
    input_type=_ADDCOLLABORATORREQUEST,
    output_type=_ADDCOLLABORATORREQUEST_RESPONSE,
    serialized_options=b'\202\323\344\223\0024\"//v1/collaborator/addOrUpdateDatasetCollaborator:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='removeDatasetCollaborator',
    full_name='ai.verta.uac.CollaboratorService.removeDatasetCollaborator',
    index=4,
    containing_service=None,
    input_type=_REMOVECOLLABORATOR,
    output_type=_REMOVECOLLABORATOR_RESPONSE,
    serialized_options=b'\202\323\344\223\002,**/v1/collaborator/removeDatasetCollaborator',
  ),
  _descriptor.MethodDescriptor(
    name='getDatasetCollaborators',
    full_name='ai.verta.uac.CollaboratorService.getDatasetCollaborators',
    index=5,
    containing_service=None,
    input_type=_GETCOLLABORATOR,
    output_type=_GETCOLLABORATOR_RESPONSE,
    serialized_options=b'\202\323\344\223\002*\022(/v1/collaborator/getDatasetCollaborators',
  ),
  _descriptor.MethodDescriptor(
    name='addOrUpdateRepositoryCollaborator',
    full_name='ai.verta.uac.CollaboratorService.addOrUpdateRepositoryCollaborator',
    index=6,
    containing_service=None,
    input_type=_ADDCOLLABORATORREQUEST,
    output_type=_ADDCOLLABORATORREQUEST_RESPONSE,
    serialized_options=b'\202\323\344\223\0027\"2/v1/collaborator/addOrUpdateRepositoryCollaborator:\001*',
  ),
  _descriptor.MethodDescriptor(
    name='removeRepositoryCollaborator',
    full_name='ai.verta.uac.CollaboratorService.removeRepositoryCollaborator',
    index=7,
    containing_service=None,
    input_type=_REMOVECOLLABORATOR,
    output_type=_REMOVECOLLABORATOR_RESPONSE,
    serialized_options=b'\202\323\344\223\002/*-/v1/collaborator/removeRepositoryCollaborator',
  ),
  _descriptor.MethodDescriptor(
    name='getRepositoryCollaborators',
    full_name='ai.verta.uac.CollaboratorService.getRepositoryCollaborators',
    index=8,
    containing_service=None,
    input_type=_GETCOLLABORATOR,
    output_type=_GETCOLLABORATOR_RESPONSE,
    serialized_options=b'\202\323\344\223\002-\022+/v1/collaborator/getRepositoryCollaborators',
  ),
])
_sym_db.RegisterServiceDescriptor(_COLLABORATORSERVICE)

DESCRIPTOR.services_by_name['CollaboratorService'] = _COLLABORATORSERVICE

# @@protoc_insertion_point(module_scope)
