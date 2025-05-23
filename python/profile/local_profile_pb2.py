# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: profile/local_profile.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from info import info_pb2 as info_dot_info__pb2
from metrics import metrics_pb2 as metrics_dot_metrics__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bprofile/local_profile.proto\x12\x16org.lfedge.eve.profile\x1a\x0finfo/info.proto\x1a\x15metrics/metrics.proto\x1a\x1fgoogle/protobuf/timestamp.proto\";\n\x0cLocalProfile\x12\x15\n\rlocal_profile\x18\x01 \x01(\t\x12\x14\n\x0cserver_token\x18\x02 \x01(\t\"\xbd\x01\n\x0bRadioStatus\x12\x15\n\rradio_silence\x18\x01 \x01(\x08\x12\x14\n\x0c\x63onfig_error\x18\x02 \x01(\t\x12?\n\x0f\x63\x65llular_status\x18\x03 \x03(\x0b\x32&.org.lfedge.eve.profile.CellularStatus\x12@\n\x10\x63\x65llular_metrics\x18\x04 \x03(\x0b\x32&.org.lfedge.eve.metrics.CellularMetric\"\xfc\x01\n\x0e\x43\x65llularStatus\x12\x14\n\x0clogicallabel\x18\x01 \x01(\t\x12\x38\n\x06module\x18\x02 \x01(\x0b\x32(.org.lfedge.eve.info.ZCellularModuleInfo\x12\x34\n\tsim_cards\x18\x03 \x03(\x0b\x32!.org.lfedge.eve.info.ZSimcardInfo\x12\x39\n\tproviders\x18\x04 \x03(\x0b\x32&.org.lfedge.eve.info.ZCellularProvider\x12\x14\n\x0c\x63onfig_error\x18\n \x01(\t\x12\x13\n\x0bprobe_error\x18\x0b \x01(\t\":\n\x0bRadioConfig\x12\x14\n\x0cserver_token\x18\x01 \x01(\t\x12\x15\n\rradio_silence\x18\x02 \x01(\x08\"K\n\x10LocalAppInfoList\x12\x37\n\tapps_info\x18\x01 \x03(\x0b\x32$.org.lfedge.eve.profile.LocalAppInfo\"\xb0\x01\n\x0cLocalAppInfo\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12+\n\x03\x65rr\x18\x04 \x01(\x0b\x32\x1e.org.lfedge.eve.info.ErrorInfo\x12,\n\x05state\x18\x05 \x01(\x0e\x32\x1d.org.lfedge.eve.info.ZSwState\x12\x1a\n\x12last_cmd_timestamp\x18\x06 \x01(\x04\"a\n\x0fLocalAppCmdList\x12\x14\n\x0cserver_token\x18\x01 \x01(\t\x12\x38\n\x0c\x61pp_commands\x18\x02 \x03(\x0b\x32\".org.lfedge.eve.profile.AppCommand\"\xc9\x01\n\nAppCommand\x12\n\n\x02id\x18\x01 \x01(\t\x12\x13\n\x0b\x64isplayname\x18\x02 \x01(\t\x12\x11\n\ttimestamp\x18\x03 \x01(\x04\x12;\n\x07\x63ommand\x18\x04 \x01(\x0e\x32*.org.lfedge.eve.profile.AppCommand.Command\"J\n\x07\x43ommand\x12\x17\n\x13\x43OMMAND_UNSPECIFIED\x10\x00\x12\x13\n\x0f\x43OMMAND_RESTART\x10\x01\x12\x11\n\rCOMMAND_PURGE\x10\x02\"\xa9\x02\n\x0cLocalDevInfo\x12\x13\n\x0b\x64\x65vice_uuid\x18\x01 \x01(\t\x12\x30\n\x05state\x18\x02 \x01(\x0e\x32!.org.lfedge.eve.info.ZDeviceState\x12L\n\x18maintenance_mode_reasons\x18\x03 \x03(\x0e\x32*.org.lfedge.eve.info.MaintenanceModeReason\x12-\n\tboot_time\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x39\n\x10last_boot_reason\x18\x05 \x01(\x0e\x32\x1f.org.lfedge.eve.info.BootReason\x12\x1a\n\x12last_cmd_timestamp\x18\n \x01(\x04\"\xac\x02\n\x0bLocalDevCmd\x12\x14\n\x0cserver_token\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x04\x12<\n\x07\x63ommand\x18\x03 \x01(\x0e\x32+.org.lfedge.eve.profile.LocalDevCmd.Command\"\xb5\x01\n\x07\x43ommand\x12\x17\n\x13\x43OMMAND_UNSPECIFIED\x10\x00\x12\x14\n\x10\x43OMMAND_SHUTDOWN\x10\x01\x12\x1d\n\x19\x43OMMAND_GRACEFUL_POWEROFF\x10\x02\x12!\n\x19\x43OMMAND_SHUTDOWN_POWEROFF\x10\x02\x1a\x02\x08\x01\x12\x1b\n\x17\x43OMMAND_GRACEFUL_REBOOT\x10\x03\x12\x18\n\x14\x43OMMAND_COLLECT_INFO\x10\x04\x1a\x02\x10\x01\x42?\n\x16org.lfedge.eve.profileZ%github.com/lf-edge/eve-api/go/profileb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'profile.local_profile_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\026org.lfedge.eve.profileZ%github.com/lf-edge/eve-api/go/profile'
  _LOCALDEVCMD_COMMAND._options = None
  _LOCALDEVCMD_COMMAND._serialized_options = b'\020\001'
  _LOCALDEVCMD_COMMAND.values_by_name["COMMAND_SHUTDOWN_POWEROFF"]._options = None
  _LOCALDEVCMD_COMMAND.values_by_name["COMMAND_SHUTDOWN_POWEROFF"]._serialized_options = b'\010\001'
  _globals['_LOCALPROFILE']._serialized_start=128
  _globals['_LOCALPROFILE']._serialized_end=187
  _globals['_RADIOSTATUS']._serialized_start=190
  _globals['_RADIOSTATUS']._serialized_end=379
  _globals['_CELLULARSTATUS']._serialized_start=382
  _globals['_CELLULARSTATUS']._serialized_end=634
  _globals['_RADIOCONFIG']._serialized_start=636
  _globals['_RADIOCONFIG']._serialized_end=694
  _globals['_LOCALAPPINFOLIST']._serialized_start=696
  _globals['_LOCALAPPINFOLIST']._serialized_end=771
  _globals['_LOCALAPPINFO']._serialized_start=774
  _globals['_LOCALAPPINFO']._serialized_end=950
  _globals['_LOCALAPPCMDLIST']._serialized_start=952
  _globals['_LOCALAPPCMDLIST']._serialized_end=1049
  _globals['_APPCOMMAND']._serialized_start=1052
  _globals['_APPCOMMAND']._serialized_end=1253
  _globals['_APPCOMMAND_COMMAND']._serialized_start=1179
  _globals['_APPCOMMAND_COMMAND']._serialized_end=1253
  _globals['_LOCALDEVINFO']._serialized_start=1256
  _globals['_LOCALDEVINFO']._serialized_end=1553
  _globals['_LOCALDEVCMD']._serialized_start=1556
  _globals['_LOCALDEVCMD']._serialized_end=1856
  _globals['_LOCALDEVCMD_COMMAND']._serialized_start=1675
  _globals['_LOCALDEVCMD_COMMAND']._serialized_end=1856
# @@protoc_insertion_point(module_scope)
