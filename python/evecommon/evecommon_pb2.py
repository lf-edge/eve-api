# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: evecommon/evecommon.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x65vecommon/evecommon.proto\x12\x15org.lfedge.eve.common\"E\n\x0f\x44iskDescription\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x14\n\x0clogical_name\x18\x02 \x01(\t\x12\x0e\n\x06serial\x18\x03 \x01(\t*q\n\rHashAlgorithm\x12\x1a\n\x16HASH_ALGORITHM_INVALID\x10\x00\x12!\n\x1dHASH_ALGORITHM_SHA256_16BYTES\x10\x01\x12!\n\x1dHASH_ALGORITHM_SHA256_32BYTES\x10\x02*\xc6\x01\n\x15RadioAccessTechnology\x12\'\n#RADIO_ACCESS_TECHNOLOGY_UNSPECIFIED\x10\x00\x12\x1f\n\x1bRADIO_ACCESS_TECHNOLOGY_GSM\x10\x01\x12 \n\x1cRADIO_ACCESS_TECHNOLOGY_UMTS\x10\x02\x12\x1f\n\x1bRADIO_ACCESS_TECHNOLOGY_LTE\x10\x03\x12 \n\x1cRADIO_ACCESS_TECHNOLOGY_5GNR\x10\x04*\x8c\x01\n\x0e\x43\x65llularIPType\x12 \n\x1c\x43\x45LLULAR_IP_TYPE_UNSPECIFIED\x10\x00\x12\x19\n\x15\x43\x45LLULAR_IP_TYPE_IPV4\x10\x01\x12\"\n\x1e\x43\x45LLULAR_IP_TYPE_IPV4_AND_IPV6\x10\x02\x12\x19\n\x15\x43\x45LLULAR_IP_TYPE_IPV6\x10\x03*u\n\nBearerType\x12\x1b\n\x17\x42\x45\x41RER_TYPE_UNSPECIFIED\x10\x00\x12\x16\n\x12\x42\x45\x41RER_TYPE_ATTACH\x10\x01\x12\x17\n\x13\x42\x45\x41RER_TYPE_DEFAULT\x10\x02\x12\x19\n\x15\x42\x45\x41RER_TYPE_DEDICATED\x10\x03\x42M\n\x15org.lfedge.eve.commonB\tEvecommonP\x01Z\'github.com/lf-edge/eve-api/go/evecommonb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'evecommon.evecommon_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.commonB\tEvecommonP\001Z\'github.com/lf-edge/eve-api/go/evecommon'
  _globals['_HASHALGORITHM']._serialized_start=123
  _globals['_HASHALGORITHM']._serialized_end=236
  _globals['_RADIOACCESSTECHNOLOGY']._serialized_start=239
  _globals['_RADIOACCESSTECHNOLOGY']._serialized_end=437
  _globals['_CELLULARIPTYPE']._serialized_start=440
  _globals['_CELLULARIPTYPE']._serialized_end=580
  _globals['_BEARERTYPE']._serialized_start=582
  _globals['_BEARERTYPE']._serialized_end=699
  _globals['_DISKDESCRIPTION']._serialized_start=52
  _globals['_DISKDESCRIPTION']._serialized_end=121
# @@protoc_insertion_point(module_scope)
