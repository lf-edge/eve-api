# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/vm.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0f\x63onfig/vm.proto\x12\x15org.lfedge.eve.config\"\xf3\x03\n\x08VmConfig\x12\x0e\n\x06kernel\x18\x01 \x01(\t\x12\x0f\n\x07ramdisk\x18\x02 \x01(\t\x12\x0e\n\x06memory\x18\x03 \x01(\r\x12\x0e\n\x06maxmem\x18\x04 \x01(\r\x12\r\n\x05vcpus\x18\x05 \x01(\r\x12\x0f\n\x07maxcpus\x18\x06 \x01(\r\x12\x0f\n\x07rootdev\x18\x07 \x01(\t\x12\x11\n\textraargs\x18\x08 \x01(\t\x12\x12\n\nbootloader\x18\t \x01(\t\x12\x0c\n\x04\x63pus\x18\n \x01(\t\x12\x12\n\ndevicetree\x18\x0b \x01(\t\x12\r\n\x05\x64tdev\x18\x0c \x03(\t\x12\x0c\n\x04irqs\x18\r \x03(\r\x12\r\n\x05iomem\x18\x0e \x03(\t\x12\x39\n\x12virtualizationMode\x18\x0f \x01(\x0e\x32\x1d.org.lfedge.eve.config.VmMode\x12\x12\n\nenable_vnc\x18\x10 \x01(\x08\x12\x12\n\nvncDisplay\x18\x11 \x01(\r\x12\x11\n\tvncPasswd\x18\x12 \x01(\t\x12\x13\n\x0b\x64isableLogs\x18\x13 \x01(\x08\x12\x0f\n\x07pin_cpu\x18\x14 \x01(\x08\x12\x12\n\nvmm_maxmem\x18\x15 \x01(\r\x12\x1a\n\x12\x65nable_vnc_shim_vm\x18\x16 \x01(\x08\x12\x34\n\tboot_mode\x18\x17 \x01(\x0e\x32!.org.lfedge.eve.config.VmBootMode*G\n\x06VmMode\x12\x06\n\x02PV\x10\x00\x12\x07\n\x03HVM\x10\x01\x12\n\n\x06\x46iller\x10\x02\x12\x07\n\x03\x46ML\x10\x03\x12\x0b\n\x07NOHYPER\x10\x04\x12\n\n\x06LEGACY\x10\x05*Z\n\nVmBootMode\x12\x1c\n\x18VM_BOOT_MODE_UNSPECIFIED\x10\x00\x12\x17\n\x13VM_BOOT_MODE_LEGACY\x10\x01\x12\x15\n\x11VM_BOOT_MODE_UEFI\x10\x02\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config.vm_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/config'
  _globals['_VMMODE']._serialized_start=544
  _globals['_VMMODE']._serialized_end=615
  _globals['_VMBOOTMODE']._serialized_start=617
  _globals['_VMBOOTMODE']._serialized_end=707
  _globals['_VMCONFIG']._serialized_start=43
  _globals['_VMCONFIG']._serialized_end=542
# @@protoc_insertion_point(module_scope)
