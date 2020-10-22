# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: host.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import meta_pb2 as meta__pb2
import driver_unknown_pb2 as driver__unknown__pb2
import driver_bridge_pb2 as driver__bridge__pb2
import driver_wireguard_pb2 as driver__wireguard__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='host.proto',
  package='abbot',
  syntax='proto3',
  serialized_options=b'Z\037arhat.dev/abbot-proto/abbotgopb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\nhost.proto\x12\x05\x61\x62\x62ot\x1a\nmeta.proto\x1a\x14\x64river_unknown.proto\x1a\x13\x64river_bridge.proto\x1a\x16\x64river_wireguard.proto\"\xda\x01\n\x14HostNetworkInterface\x12)\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.abbot.NetworkInterface\x12\x10\n\x08provider\x18\x02 \x01(\t\x12\'\n\x07unknown\x18\n \x01(\x0b\x32\x14.abbot.DriverUnknownH\x00\x12%\n\x06\x62ridge\x18\x0b \x01(\x0b\x32\x13.abbot.DriverBridgeH\x00\x12+\n\twireguard\x18\x0c \x01(\x0b\x32\x16.abbot.DriverWireguardH\x00\x42\x08\n\x06\x63onfig\"O\n\x1eHostNetworkConfigEnsureRequest\x12-\n\x08\x65xpected\x18\x01 \x03(\x0b\x32\x1b.abbot.HostNetworkInterface\"2\n\x1dHostNetworkConfigQueryRequest\x12\x11\n\tproviders\x18\x01 \x03(\t\"H\n\x19HostNetworkConfigResponse\x12+\n\x06\x61\x63tual\x18\x01 \x03(\x0b\x32\x1b.abbot.HostNetworkInterfaceB!Z\x1f\x61rhat.dev/abbot-proto/abbotgopbb\x06proto3'
  ,
  dependencies=[meta__pb2.DESCRIPTOR,driver__unknown__pb2.DESCRIPTOR,driver__bridge__pb2.DESCRIPTOR,driver__wireguard__pb2.DESCRIPTOR,])




_HOSTNETWORKINTERFACE = _descriptor.Descriptor(
  name='HostNetworkInterface',
  full_name='abbot.HostNetworkInterface',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='metadata', full_name='abbot.HostNetworkInterface.metadata', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='provider', full_name='abbot.HostNetworkInterface.provider', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='unknown', full_name='abbot.HostNetworkInterface.unknown', index=2,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='bridge', full_name='abbot.HostNetworkInterface.bridge', index=3,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='wireguard', full_name='abbot.HostNetworkInterface.wireguard', index=4,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
      name='config', full_name='abbot.HostNetworkInterface.config',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=101,
  serialized_end=319,
)


_HOSTNETWORKCONFIGENSUREREQUEST = _descriptor.Descriptor(
  name='HostNetworkConfigEnsureRequest',
  full_name='abbot.HostNetworkConfigEnsureRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='expected', full_name='abbot.HostNetworkConfigEnsureRequest.expected', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=321,
  serialized_end=400,
)


_HOSTNETWORKCONFIGQUERYREQUEST = _descriptor.Descriptor(
  name='HostNetworkConfigQueryRequest',
  full_name='abbot.HostNetworkConfigQueryRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='providers', full_name='abbot.HostNetworkConfigQueryRequest.providers', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=402,
  serialized_end=452,
)


_HOSTNETWORKCONFIGRESPONSE = _descriptor.Descriptor(
  name='HostNetworkConfigResponse',
  full_name='abbot.HostNetworkConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='actual', full_name='abbot.HostNetworkConfigResponse.actual', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=454,
  serialized_end=526,
)

_HOSTNETWORKINTERFACE.fields_by_name['metadata'].message_type = meta__pb2._NETWORKINTERFACE
_HOSTNETWORKINTERFACE.fields_by_name['unknown'].message_type = driver__unknown__pb2._DRIVERUNKNOWN
_HOSTNETWORKINTERFACE.fields_by_name['bridge'].message_type = driver__bridge__pb2._DRIVERBRIDGE
_HOSTNETWORKINTERFACE.fields_by_name['wireguard'].message_type = driver__wireguard__pb2._DRIVERWIREGUARD
_HOSTNETWORKINTERFACE.oneofs_by_name['config'].fields.append(
  _HOSTNETWORKINTERFACE.fields_by_name['unknown'])
_HOSTNETWORKINTERFACE.fields_by_name['unknown'].containing_oneof = _HOSTNETWORKINTERFACE.oneofs_by_name['config']
_HOSTNETWORKINTERFACE.oneofs_by_name['config'].fields.append(
  _HOSTNETWORKINTERFACE.fields_by_name['bridge'])
_HOSTNETWORKINTERFACE.fields_by_name['bridge'].containing_oneof = _HOSTNETWORKINTERFACE.oneofs_by_name['config']
_HOSTNETWORKINTERFACE.oneofs_by_name['config'].fields.append(
  _HOSTNETWORKINTERFACE.fields_by_name['wireguard'])
_HOSTNETWORKINTERFACE.fields_by_name['wireguard'].containing_oneof = _HOSTNETWORKINTERFACE.oneofs_by_name['config']
_HOSTNETWORKCONFIGENSUREREQUEST.fields_by_name['expected'].message_type = _HOSTNETWORKINTERFACE
_HOSTNETWORKCONFIGRESPONSE.fields_by_name['actual'].message_type = _HOSTNETWORKINTERFACE
DESCRIPTOR.message_types_by_name['HostNetworkInterface'] = _HOSTNETWORKINTERFACE
DESCRIPTOR.message_types_by_name['HostNetworkConfigEnsureRequest'] = _HOSTNETWORKCONFIGENSUREREQUEST
DESCRIPTOR.message_types_by_name['HostNetworkConfigQueryRequest'] = _HOSTNETWORKCONFIGQUERYREQUEST
DESCRIPTOR.message_types_by_name['HostNetworkConfigResponse'] = _HOSTNETWORKCONFIGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

HostNetworkInterface = _reflection.GeneratedProtocolMessageType('HostNetworkInterface', (_message.Message,), {
  'DESCRIPTOR' : _HOSTNETWORKINTERFACE,
  '__module__' : 'host_pb2'
  # @@protoc_insertion_point(class_scope:abbot.HostNetworkInterface)
  })
_sym_db.RegisterMessage(HostNetworkInterface)

HostNetworkConfigEnsureRequest = _reflection.GeneratedProtocolMessageType('HostNetworkConfigEnsureRequest', (_message.Message,), {
  'DESCRIPTOR' : _HOSTNETWORKCONFIGENSUREREQUEST,
  '__module__' : 'host_pb2'
  # @@protoc_insertion_point(class_scope:abbot.HostNetworkConfigEnsureRequest)
  })
_sym_db.RegisterMessage(HostNetworkConfigEnsureRequest)

HostNetworkConfigQueryRequest = _reflection.GeneratedProtocolMessageType('HostNetworkConfigQueryRequest', (_message.Message,), {
  'DESCRIPTOR' : _HOSTNETWORKCONFIGQUERYREQUEST,
  '__module__' : 'host_pb2'
  # @@protoc_insertion_point(class_scope:abbot.HostNetworkConfigQueryRequest)
  })
_sym_db.RegisterMessage(HostNetworkConfigQueryRequest)

HostNetworkConfigResponse = _reflection.GeneratedProtocolMessageType('HostNetworkConfigResponse', (_message.Message,), {
  'DESCRIPTOR' : _HOSTNETWORKCONFIGRESPONSE,
  '__module__' : 'host_pb2'
  # @@protoc_insertion_point(class_scope:abbot.HostNetworkConfigResponse)
  })
_sym_db.RegisterMessage(HostNetworkConfigResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
