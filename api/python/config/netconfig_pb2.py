# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/netconfig.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import acipherinfo_pb2 as config_dot_acipherinfo__pb2
from config import fw_pb2 as config_dot_fw__pb2
from config import netcmn_pb2 as config_dot_netcmn__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/netconfig.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config'),
  serialized_pb=_b('\n\x16\x63onfig/netconfig.proto\x1a\x18\x63onfig/acipherinfo.proto\x1a\x0f\x63onfig/fw.proto\x1a\x13\x63onfig/netcmn.proto\"\xb1\x01\n\rNetworkConfig\x12\n\n\x02id\x18\x01 \x01(\t\x12\x1a\n\x04type\x18\x05 \x01(\x0e\x32\x0c.NetworkType\x12\x13\n\x02ip\x18\x06 \x01(\x0b\x32\x07.ipspec\x12 \n\x03\x64ns\x18\x07 \x03(\x0b\x32\x13.ZnetStaticDNSEntry\x12\x1e\n\x08\x65ntProxy\x18\x08 \x01(\x0b\x32\x0c.ProxyConfig\x12!\n\x08wireless\x18\n \x01(\x0b\x32\x0f.WirelessConfig\"\xcb\x01\n\x0eNetworkAdapter\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tnetworkId\x18\x03 \x01(\t\x12\x0c\n\x04\x61\x64\x64r\x18\x04 \x01(\t\x12\x10\n\x08hostname\x18\x05 \x01(\t\x12\x11\n\tcryptoEid\x18\n \x01(\t\x12\x15\n\rlispsignature\x18\x06 \x01(\t\x12\x0f\n\x07pemcert\x18\x07 \x01(\x0c\x12\x15\n\rpemprivatekey\x18\x08 \x01(\x0c\x12\x12\n\nmacAddress\x18\t \x01(\t\x12\x12\n\x04\x61\x63ls\x18( \x03(\x0b\x32\x04.ACE\"q\n\x0eWirelessConfig\x12\x1b\n\x04type\x18\x01 \x01(\x0e\x32\r.WirelessType\x12$\n\x0b\x63\x65llularCfg\x18\x05 \x03(\x0b\x32\x0f.CellularConfig\x12\x1c\n\x07wifiCfg\x18\n \x03(\x0b\x32\x0b.WifiConfig\"\x1d\n\x0e\x43\x65llularConfig\x12\x0b\n\x03\x41PN\x18\x01 \x01(\t\"\xf5\x01\n\nWifiConfig\x12\x10\n\x08wifiSSID\x18\x01 \x01(\t\x12!\n\tkeyScheme\x18\x02 \x01(\x0e\x32\x0e.WiFiKeyScheme\x12\x10\n\x08identity\x18\x05 \x01(\t\x12\x10\n\x08password\x18\n \x01(\t\x12\'\n\x06\x63rypto\x18\x14 \x01(\x0b\x32\x17.WifiConfig.cryptoblock\x12\x10\n\x08priority\x18\x19 \x01(\x05\x12 \n\ncipherData\x18\x1e \x01(\x0b\x32\x0c.CipherBlock\x1a\x31\n\x0b\x63ryptoblock\x12\x10\n\x08identity\x18\x0b \x01(\t\x12\x10\n\x08password\x18\x0c \x01(\tB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3')
  ,
  dependencies=[config_dot_acipherinfo__pb2.DESCRIPTOR,config_dot_fw__pb2.DESCRIPTOR,config_dot_netcmn__pb2.DESCRIPTOR,])




_NETWORKCONFIG = _descriptor.Descriptor(
  name='NetworkConfig',
  full_name='NetworkConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='NetworkConfig.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='NetworkConfig.type', index=1,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ip', full_name='NetworkConfig.ip', index=2,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dns', full_name='NetworkConfig.dns', index=3,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entProxy', full_name='NetworkConfig.entProxy', index=4,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='wireless', full_name='NetworkConfig.wireless', index=5,
      number=10, type=11, cpp_type=10, label=1,
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
  ],
  serialized_start=91,
  serialized_end=268,
)


_NETWORKADAPTER = _descriptor.Descriptor(
  name='NetworkAdapter',
  full_name='NetworkAdapter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='NetworkAdapter.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='networkId', full_name='NetworkAdapter.networkId', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addr', full_name='NetworkAdapter.addr', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hostname', full_name='NetworkAdapter.hostname', index=3,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cryptoEid', full_name='NetworkAdapter.cryptoEid', index=4,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lispsignature', full_name='NetworkAdapter.lispsignature', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pemcert', full_name='NetworkAdapter.pemcert', index=6,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pemprivatekey', full_name='NetworkAdapter.pemprivatekey', index=7,
      number=8, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='macAddress', full_name='NetworkAdapter.macAddress', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='acls', full_name='NetworkAdapter.acls', index=9,
      number=40, type=11, cpp_type=10, label=3,
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
  serialized_start=271,
  serialized_end=474,
)


_WIRELESSCONFIG = _descriptor.Descriptor(
  name='WirelessConfig',
  full_name='WirelessConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='WirelessConfig.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cellularCfg', full_name='WirelessConfig.cellularCfg', index=1,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='wifiCfg', full_name='WirelessConfig.wifiCfg', index=2,
      number=10, type=11, cpp_type=10, label=3,
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
  serialized_start=476,
  serialized_end=589,
)


_CELLULARCONFIG = _descriptor.Descriptor(
  name='CellularConfig',
  full_name='CellularConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='APN', full_name='CellularConfig.APN', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=591,
  serialized_end=620,
)


_WIFICONFIG_CRYPTOBLOCK = _descriptor.Descriptor(
  name='cryptoblock',
  full_name='WifiConfig.cryptoblock',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='identity', full_name='WifiConfig.cryptoblock.identity', index=0,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='password', full_name='WifiConfig.cryptoblock.password', index=1,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=819,
  serialized_end=868,
)

_WIFICONFIG = _descriptor.Descriptor(
  name='WifiConfig',
  full_name='WifiConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='wifiSSID', full_name='WifiConfig.wifiSSID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='keyScheme', full_name='WifiConfig.keyScheme', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='identity', full_name='WifiConfig.identity', index=2,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='password', full_name='WifiConfig.password', index=3,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='crypto', full_name='WifiConfig.crypto', index=4,
      number=20, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='priority', full_name='WifiConfig.priority', index=5,
      number=25, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cipherData', full_name='WifiConfig.cipherData', index=6,
      number=30, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_WIFICONFIG_CRYPTOBLOCK, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=623,
  serialized_end=868,
)

_NETWORKCONFIG.fields_by_name['type'].enum_type = config_dot_netcmn__pb2._NETWORKTYPE
_NETWORKCONFIG.fields_by_name['ip'].message_type = config_dot_netcmn__pb2._IPSPEC
_NETWORKCONFIG.fields_by_name['dns'].message_type = config_dot_netcmn__pb2._ZNETSTATICDNSENTRY
_NETWORKCONFIG.fields_by_name['entProxy'].message_type = config_dot_netcmn__pb2._PROXYCONFIG
_NETWORKCONFIG.fields_by_name['wireless'].message_type = _WIRELESSCONFIG
_NETWORKADAPTER.fields_by_name['acls'].message_type = config_dot_fw__pb2._ACE
_WIRELESSCONFIG.fields_by_name['type'].enum_type = config_dot_netcmn__pb2._WIRELESSTYPE
_WIRELESSCONFIG.fields_by_name['cellularCfg'].message_type = _CELLULARCONFIG
_WIRELESSCONFIG.fields_by_name['wifiCfg'].message_type = _WIFICONFIG
_WIFICONFIG_CRYPTOBLOCK.containing_type = _WIFICONFIG
_WIFICONFIG.fields_by_name['keyScheme'].enum_type = config_dot_netcmn__pb2._WIFIKEYSCHEME
_WIFICONFIG.fields_by_name['crypto'].message_type = _WIFICONFIG_CRYPTOBLOCK
_WIFICONFIG.fields_by_name['cipherData'].message_type = config_dot_acipherinfo__pb2._CIPHERBLOCK
DESCRIPTOR.message_types_by_name['NetworkConfig'] = _NETWORKCONFIG
DESCRIPTOR.message_types_by_name['NetworkAdapter'] = _NETWORKADAPTER
DESCRIPTOR.message_types_by_name['WirelessConfig'] = _WIRELESSCONFIG
DESCRIPTOR.message_types_by_name['CellularConfig'] = _CELLULARCONFIG
DESCRIPTOR.message_types_by_name['WifiConfig'] = _WIFICONFIG
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

NetworkConfig = _reflection.GeneratedProtocolMessageType('NetworkConfig', (_message.Message,), dict(
  DESCRIPTOR = _NETWORKCONFIG,
  __module__ = 'config.netconfig_pb2'
  # @@protoc_insertion_point(class_scope:NetworkConfig)
  ))
_sym_db.RegisterMessage(NetworkConfig)

NetworkAdapter = _reflection.GeneratedProtocolMessageType('NetworkAdapter', (_message.Message,), dict(
  DESCRIPTOR = _NETWORKADAPTER,
  __module__ = 'config.netconfig_pb2'
  # @@protoc_insertion_point(class_scope:NetworkAdapter)
  ))
_sym_db.RegisterMessage(NetworkAdapter)

WirelessConfig = _reflection.GeneratedProtocolMessageType('WirelessConfig', (_message.Message,), dict(
  DESCRIPTOR = _WIRELESSCONFIG,
  __module__ = 'config.netconfig_pb2'
  # @@protoc_insertion_point(class_scope:WirelessConfig)
  ))
_sym_db.RegisterMessage(WirelessConfig)

CellularConfig = _reflection.GeneratedProtocolMessageType('CellularConfig', (_message.Message,), dict(
  DESCRIPTOR = _CELLULARCONFIG,
  __module__ = 'config.netconfig_pb2'
  # @@protoc_insertion_point(class_scope:CellularConfig)
  ))
_sym_db.RegisterMessage(CellularConfig)

WifiConfig = _reflection.GeneratedProtocolMessageType('WifiConfig', (_message.Message,), dict(

  cryptoblock = _reflection.GeneratedProtocolMessageType('cryptoblock', (_message.Message,), dict(
    DESCRIPTOR = _WIFICONFIG_CRYPTOBLOCK,
    __module__ = 'config.netconfig_pb2'
    # @@protoc_insertion_point(class_scope:WifiConfig.cryptoblock)
    ))
  ,
  DESCRIPTOR = _WIFICONFIG,
  __module__ = 'config.netconfig_pb2'
  # @@protoc_insertion_point(class_scope:WifiConfig)
  ))
_sym_db.RegisterMessage(WifiConfig)
_sym_db.RegisterMessage(WifiConfig.cryptoblock)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
