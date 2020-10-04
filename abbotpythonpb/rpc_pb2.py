# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: rpc.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import proto_pb2 as proto__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='rpc.proto',
  package='abbot',
  syntax='proto3',
  serialized_options=b'Z\037arhat.dev/abbot-proto/abbotgopb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\trpc.proto\x12\x05\x61\x62\x62ot\x1a\x0bproto.proto2<\n\x0eNetworkManager\x12*\n\x07Process\x12\x0e.abbot.Request\x1a\x0f.abbot.ResponseB!Z\x1f\x61rhat.dev/abbot-proto/abbotgopbb\x06proto3'
  ,
  dependencies=[proto__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_NETWORKMANAGER = _descriptor.ServiceDescriptor(
  name='NetworkManager',
  full_name='abbot.NetworkManager',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=33,
  serialized_end=93,
  methods=[
  _descriptor.MethodDescriptor(
    name='Process',
    full_name='abbot.NetworkManager.Process',
    index=0,
    containing_service=None,
    input_type=proto__pb2._REQUEST,
    output_type=proto__pb2._RESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_NETWORKMANAGER)

DESCRIPTOR.services_by_name['NetworkManager'] = _NETWORKMANAGER

# @@protoc_insertion_point(module_scope)