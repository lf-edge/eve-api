# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/devconfig.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from config import acipherinfo_pb2 as config_dot_acipherinfo__pb2
from config import appconfig_pb2 as config_dot_appconfig__pb2
from config import baseosconfig_pb2 as config_dot_baseosconfig__pb2
from config import devcommon_pb2 as config_dot_devcommon__pb2
from config import devmodel_pb2 as config_dot_devmodel__pb2
from config import mesh_pb2 as config_dot_mesh__pb2
from config import netconfig_pb2 as config_dot_netconfig__pb2
from config import netinst_pb2 as config_dot_netinst__pb2
from config import storage_pb2 as config_dot_storage__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='config/devconfig.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/config'),
  serialized_pb=_b('\n\x16\x63onfig/devconfig.proto\x1a\x18\x63onfig/acipherinfo.proto\x1a\x16\x63onfig/appconfig.proto\x1a\x19\x63onfig/baseosconfig.proto\x1a\x16\x63onfig/devcommon.proto\x1a\x15\x63onfig/devmodel.proto\x1a\x11\x63onfig/mesh.proto\x1a\x16\x63onfig/netconfig.proto\x1a\x14\x63onfig/netinst.proto\x1a\x14\x63onfig/storage.proto\"\x8e\x05\n\rEdgeDevConfig\x12\x1b\n\x02id\x18\x01 \x01(\x0b\x32\x0f.UUIDandVersion\x12 \n\x04\x61pps\x18\x04 \x03(\x0b\x32\x12.AppInstanceConfig\x12 \n\x08networks\x18\x05 \x03(\x0b\x32\x0e.NetworkConfig\x12$\n\ndatastores\x18\x06 \x03(\x0b\x32\x10.DatastoreConfig\x12$\n\x08lispInfo\x18\x07 \x01(\x0b\x32\x12.DeviceLispDetails\x12\x1b\n\x04\x62\x61se\x18\x08 \x03(\x0b\x32\r.BaseOSConfig\x12\x1d\n\x06reboot\x18\t \x01(\x0b\x32\r.DeviceOpsCmd\x12\x1d\n\x06\x62\x61\x63kup\x18\n \x01(\x0b\x32\r.DeviceOpsCmd\x12 \n\x0b\x63onfigItems\x18\x0b \x03(\x0b\x32\x0b.ConfigItem\x12)\n\x11systemAdapterList\x18\x0c \x03(\x0b\x32\x0e.SystemAdapter\x12!\n\x0c\x64\x65viceIoList\x18\r \x03(\x0b\x32\x0b.PhysicalIO\x12\x14\n\x0cmanufacturer\x18\x0e \x01(\t\x12\x13\n\x0bproductName\x18\x0f \x01(\t\x12\x30\n\x10networkInstances\x18\x10 \x03(\x0b\x32\x16.NetworkInstanceConfig\x12\x12\n\nenterprise\x18\x11 \x01(\t\x12\x0c\n\x04name\x18\x12 \x01(\t\x12&\n\x0e\x63ipherContexts\x18\x13 \x03(\x0b\x32\x0e.CipherContext\x12!\n\x0b\x63ontentInfo\x18\x14 \x03(\x0b\x32\x0c.ContentTree\x12\x18\n\x07volumes\x18\x15 \x03(\x0b\x32\x07.Volume\x12!\n\x19\x63ontrollercert_confighash\x18\x16 \x01(\t\"#\n\rConfigRequest\x12\x12\n\nconfigHash\x18\x01 \x01(\t\"D\n\x0e\x43onfigResponse\x12\x1e\n\x06\x63onfig\x18\x01 \x01(\x0b\x32\x0e.EdgeDevConfig\x12\x12\n\nconfigHash\x18\x02 \x01(\tB=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve/api/go/configb\x06proto3')
  ,
  dependencies=[config_dot_acipherinfo__pb2.DESCRIPTOR,config_dot_appconfig__pb2.DESCRIPTOR,config_dot_baseosconfig__pb2.DESCRIPTOR,config_dot_devcommon__pb2.DESCRIPTOR,config_dot_devmodel__pb2.DESCRIPTOR,config_dot_mesh__pb2.DESCRIPTOR,config_dot_netconfig__pb2.DESCRIPTOR,config_dot_netinst__pb2.DESCRIPTOR,config_dot_storage__pb2.DESCRIPTOR,])




_EDGEDEVCONFIG = _descriptor.Descriptor(
  name='EdgeDevConfig',
  full_name='EdgeDevConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='EdgeDevConfig.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='apps', full_name='EdgeDevConfig.apps', index=1,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='networks', full_name='EdgeDevConfig.networks', index=2,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='datastores', full_name='EdgeDevConfig.datastores', index=3,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='lispInfo', full_name='EdgeDevConfig.lispInfo', index=4,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='base', full_name='EdgeDevConfig.base', index=5,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='reboot', full_name='EdgeDevConfig.reboot', index=6,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='backup', full_name='EdgeDevConfig.backup', index=7,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='configItems', full_name='EdgeDevConfig.configItems', index=8,
      number=11, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='systemAdapterList', full_name='EdgeDevConfig.systemAdapterList', index=9,
      number=12, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='deviceIoList', full_name='EdgeDevConfig.deviceIoList', index=10,
      number=13, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='manufacturer', full_name='EdgeDevConfig.manufacturer', index=11,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='productName', full_name='EdgeDevConfig.productName', index=12,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='networkInstances', full_name='EdgeDevConfig.networkInstances', index=13,
      number=16, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enterprise', full_name='EdgeDevConfig.enterprise', index=14,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='EdgeDevConfig.name', index=15,
      number=18, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cipherContexts', full_name='EdgeDevConfig.cipherContexts', index=16,
      number=19, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='contentInfo', full_name='EdgeDevConfig.contentInfo', index=17,
      number=20, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='volumes', full_name='EdgeDevConfig.volumes', index=18,
      number=21, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='controllercert_confighash', full_name='EdgeDevConfig.controllercert_confighash', index=19,
      number=22, type=9, cpp_type=9, label=1,
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
  serialized_start=238,
  serialized_end=892,
)


_CONFIGREQUEST = _descriptor.Descriptor(
  name='ConfigRequest',
  full_name='ConfigRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='configHash', full_name='ConfigRequest.configHash', index=0,
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
  serialized_start=894,
  serialized_end=929,
)


_CONFIGRESPONSE = _descriptor.Descriptor(
  name='ConfigResponse',
  full_name='ConfigResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='config', full_name='ConfigResponse.config', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='configHash', full_name='ConfigResponse.configHash', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=931,
  serialized_end=999,
)

_EDGEDEVCONFIG.fields_by_name['id'].message_type = config_dot_devcommon__pb2._UUIDANDVERSION
_EDGEDEVCONFIG.fields_by_name['apps'].message_type = config_dot_appconfig__pb2._APPINSTANCECONFIG
_EDGEDEVCONFIG.fields_by_name['networks'].message_type = config_dot_netconfig__pb2._NETWORKCONFIG
_EDGEDEVCONFIG.fields_by_name['datastores'].message_type = config_dot_storage__pb2._DATASTORECONFIG
_EDGEDEVCONFIG.fields_by_name['lispInfo'].message_type = config_dot_mesh__pb2._DEVICELISPDETAILS
_EDGEDEVCONFIG.fields_by_name['base'].message_type = config_dot_baseosconfig__pb2._BASEOSCONFIG
_EDGEDEVCONFIG.fields_by_name['reboot'].message_type = config_dot_devcommon__pb2._DEVICEOPSCMD
_EDGEDEVCONFIG.fields_by_name['backup'].message_type = config_dot_devcommon__pb2._DEVICEOPSCMD
_EDGEDEVCONFIG.fields_by_name['configItems'].message_type = config_dot_devcommon__pb2._CONFIGITEM
_EDGEDEVCONFIG.fields_by_name['systemAdapterList'].message_type = config_dot_devmodel__pb2._SYSTEMADAPTER
_EDGEDEVCONFIG.fields_by_name['deviceIoList'].message_type = config_dot_devmodel__pb2._PHYSICALIO
_EDGEDEVCONFIG.fields_by_name['networkInstances'].message_type = config_dot_netinst__pb2._NETWORKINSTANCECONFIG
_EDGEDEVCONFIG.fields_by_name['cipherContexts'].message_type = config_dot_acipherinfo__pb2._CIPHERCONTEXT
_EDGEDEVCONFIG.fields_by_name['contentInfo'].message_type = config_dot_storage__pb2._CONTENTTREE
_EDGEDEVCONFIG.fields_by_name['volumes'].message_type = config_dot_storage__pb2._VOLUME
_CONFIGRESPONSE.fields_by_name['config'].message_type = _EDGEDEVCONFIG
DESCRIPTOR.message_types_by_name['EdgeDevConfig'] = _EDGEDEVCONFIG
DESCRIPTOR.message_types_by_name['ConfigRequest'] = _CONFIGREQUEST
DESCRIPTOR.message_types_by_name['ConfigResponse'] = _CONFIGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EdgeDevConfig = _reflection.GeneratedProtocolMessageType('EdgeDevConfig', (_message.Message,), dict(
  DESCRIPTOR = _EDGEDEVCONFIG,
  __module__ = 'config.devconfig_pb2'
  # @@protoc_insertion_point(class_scope:EdgeDevConfig)
  ))
_sym_db.RegisterMessage(EdgeDevConfig)

ConfigRequest = _reflection.GeneratedProtocolMessageType('ConfigRequest', (_message.Message,), dict(
  DESCRIPTOR = _CONFIGREQUEST,
  __module__ = 'config.devconfig_pb2'
  # @@protoc_insertion_point(class_scope:ConfigRequest)
  ))
_sym_db.RegisterMessage(ConfigRequest)

ConfigResponse = _reflection.GeneratedProtocolMessageType('ConfigResponse', (_message.Message,), dict(
  DESCRIPTOR = _CONFIGRESPONSE,
  __module__ = 'config.devconfig_pb2'
  # @@protoc_insertion_point(class_scope:ConfigResponse)
  ))
_sym_db.RegisterMessage(ConfigResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
