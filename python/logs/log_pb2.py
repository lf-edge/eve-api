# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: logs/log.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0elogs/log.proto\x12\x13org.lfedge.eve.logs\x1a\x1fgoogle/protobuf/timestamp.proto\"\x90\x02\n\x08LogEntry\x12\x10\n\x08severity\x18\x01 \x01(\t\x12\x0e\n\x06source\x18\x02 \x01(\t\x12\x0b\n\x03iid\x18\x03 \x01(\t\x12\x0f\n\x07\x63ontent\x18\x04 \x01(\t\x12\r\n\x05msgid\x18\x05 \x01(\x04\x12\x35\n\x04tags\x18\x06 \x03(\x0b\x32\'.org.lfedge.eve.logs.LogEntry.TagsEntry\x12-\n\ttimestamp\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x10\n\x08\x66ilename\x18\x08 \x01(\t\x12\x10\n\x08\x66unction\x18\t \x01(\t\x1a+\n\tTagsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x98\x01\n\tLogBundle\x12\r\n\x05\x64\x65vID\x18\x01 \x01(\t\x12\r\n\x05image\x18\x02 \x01(\t\x12*\n\x03log\x18\x03 \x03(\x0b\x32\x1d.org.lfedge.eve.logs.LogEntry\x12-\n\ttimestamp\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x12\n\neveVersion\x18\x05 \x01(\t\"q\n\x14\x41ppInstanceLogBundle\x12*\n\x03log\x18\x01 \x03(\x0b\x32\x1d.org.lfedge.eve.logs.LogEntry\x12-\n\ttimestamp\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"G\n\rServerMetrics\x12\x16\n\x0e\x63pu_percentage\x18\x01 \x01(\x02\x12\x1e\n\x16log_process_delay_msec\x18\x02 \x01(\rB9\n\x13org.lfedge.eve.logsZ\"github.com/lf-edge/eve-api/go/logsb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'logs.log_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023org.lfedge.eve.logsZ\"github.com/lf-edge/eve-api/go/logs'
  _LOGENTRY_TAGSENTRY._options = None
  _LOGENTRY_TAGSENTRY._serialized_options = b'8\001'
  _LOGENTRY._serialized_start=73
  _LOGENTRY._serialized_end=345
  _LOGENTRY_TAGSENTRY._serialized_start=302
  _LOGENTRY_TAGSENTRY._serialized_end=345
  _LOGBUNDLE._serialized_start=348
  _LOGBUNDLE._serialized_end=500
  _APPINSTANCELOGBUNDLE._serialized_start=502
  _APPINSTANCELOGBUNDLE._serialized_end=615
  _SERVERMETRICS._serialized_start=617
  _SERVERMETRICS._serialized_end=688
# @@protoc_insertion_point(module_scope)
