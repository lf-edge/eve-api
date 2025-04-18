# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: config/netcmn.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13\x63onfig/netcmn.proto\x12\x15org.lfedge.eve.config\"%\n\x07ipRange\x12\r\n\x05start\x18\x01 \x01(\t\x12\x0b\n\x03\x65nd\x18\x02 \x01(\t\"]\n\x0bProxyServer\x12\x30\n\x05proto\x18\x01 \x01(\x0e\x32!.org.lfedge.eve.config.proxyProto\x12\x0e\n\x06server\x18\x02 \x01(\t\x12\x0c\n\x04port\x18\x03 \x01(\r\"\xb2\x01\n\x0bProxyConfig\x12\x1a\n\x12networkProxyEnable\x18\x01 \x01(\x08\x12\x33\n\x07proxies\x18\x02 \x03(\x0b\x32\".org.lfedge.eve.config.ProxyServer\x12\x12\n\nexceptions\x18\x03 \x01(\t\x12\x0f\n\x07pacfile\x18\x04 \x01(\t\x12\x17\n\x0fnetworkProxyURL\x18\x05 \x01(\t\x12\x14\n\x0cproxyCertPEM\x18\x06 \x03(\x0c\"*\n\tZedServer\x12\x10\n\x08HostName\x18\x01 \x01(\t\x12\x0b\n\x03\x45ID\x18\x02 \x03(\t\"7\n\x12ZnetStaticDNSEntry\x12\x10\n\x08HostName\x18\x01 \x01(\t\x12\x0f\n\x07\x41\x64\x64ress\x18\x02 \x03(\t\"\x8e\x02\n\x06ipspec\x12-\n\x04\x64hcp\x18\x02 \x01(\x0e\x32\x1f.org.lfedge.eve.config.DHCPType\x12\x0e\n\x06subnet\x18\x03 \x01(\t\x12\x0f\n\x07gateway\x18\x05 \x01(\t\x12\x0e\n\x06\x64omain\x18\x06 \x01(\t\x12\x0b\n\x03ntp\x18\x07 \x01(\t\x12\x10\n\x08more_ntp\x18\n \x03(\t\x12\x0b\n\x03\x64ns\x18\x08 \x03(\t\x12\x31\n\tdhcpRange\x18\t \x01(\x0b\x32\x1e.org.lfedge.eve.config.ipRange\x12\x45\n\x13\x64hcp_options_ignore\x18\x0b \x01(\x0b\x32(.org.lfedge.eve.config.DhcpOptionsIgnore\"3\n\x11\x44hcpOptionsIgnore\x12\x1e\n\x16ntp_server_exclusively\x18\x01 \x01(\x08\"+\n\rProbeEndpoint\x12\x0c\n\x04host\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\r\"\x97\x01\n\x11\x43onnectivityProbe\x12\x44\n\x0cprobe_method\x18\x01 \x01(\x0e\x32..org.lfedge.eve.config.ConnectivityProbeMethod\x12<\n\x0eprobe_endpoint\x18\x02 \x01(\x0b\x32$.org.lfedge.eve.config.ProbeEndpoint*_\n\nproxyProto\x12\x0e\n\nPROXY_HTTP\x10\x00\x12\x0f\n\x0bPROXY_HTTPS\x10\x01\x12\x0f\n\x0bPROXY_SOCKS\x10\x02\x12\r\n\tPROXY_FTP\x10\x03\x12\x10\n\x0bPROXY_OTHER\x10\xff\x01*>\n\x08\x44HCPType\x12\x0c\n\x08\x44HCPNoop\x10\x00\x12\n\n\x06Static\x10\x01\x12\x0c\n\x08\x44HCPNone\x10\x02\x12\n\n\x06\x43lient\x10\x04*\x83\x01\n\x0bNetworkType\x12\x13\n\x0fNETWORKTYPENOOP\x10\x00\x12\x06\n\x02V4\x10\x04\x12\x06\n\x02V6\x10\x06\x12\x0c\n\x08\x43ryptoV4\x10\x18\x12\x0c\n\x08\x43ryptoV6\x10\x1a\x12\r\n\tCryptoEID\x10\x0e\x12\n\n\x06V4Only\x10\x07\x12\n\n\x06V6Only\x10\x08\x12\x0c\n\x08\x44ualV4V6\x10\t*4\n\x0cWirelessType\x12\x0c\n\x08TypeNOOP\x10\x00\x12\x08\n\x04WiFi\x10\x01\x12\x0c\n\x08\x43\x65llular\x10\x02*7\n\rWiFiKeyScheme\x12\x0e\n\nSchemeNOOP\x10\x00\x12\n\n\x06WPAPSK\x10\x01\x12\n\n\x06WPAEAP\x10\x02*\x8b\x01\n\x17\x43onnectivityProbeMethod\x12)\n%CONNECTIVITY_PROBE_METHOD_UNSPECIFIED\x10\x00\x12\"\n\x1e\x43ONNECTIVITY_PROBE_METHOD_ICMP\x10\x01\x12!\n\x1d\x43ONNECTIVITY_PROBE_METHOD_TCP\x10\x02\x42=\n\x15org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/configb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config.netcmn_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025org.lfedge.eve.configZ$github.com/lf-edge/eve-api/go/config'
  _globals['_PROXYPROTO']._serialized_start=987
  _globals['_PROXYPROTO']._serialized_end=1082
  _globals['_DHCPTYPE']._serialized_start=1084
  _globals['_DHCPTYPE']._serialized_end=1146
  _globals['_NETWORKTYPE']._serialized_start=1149
  _globals['_NETWORKTYPE']._serialized_end=1280
  _globals['_WIRELESSTYPE']._serialized_start=1282
  _globals['_WIRELESSTYPE']._serialized_end=1334
  _globals['_WIFIKEYSCHEME']._serialized_start=1336
  _globals['_WIFIKEYSCHEME']._serialized_end=1391
  _globals['_CONNECTIVITYPROBEMETHOD']._serialized_start=1394
  _globals['_CONNECTIVITYPROBEMETHOD']._serialized_end=1533
  _globals['_IPRANGE']._serialized_start=46
  _globals['_IPRANGE']._serialized_end=83
  _globals['_PROXYSERVER']._serialized_start=85
  _globals['_PROXYSERVER']._serialized_end=178
  _globals['_PROXYCONFIG']._serialized_start=181
  _globals['_PROXYCONFIG']._serialized_end=359
  _globals['_ZEDSERVER']._serialized_start=361
  _globals['_ZEDSERVER']._serialized_end=403
  _globals['_ZNETSTATICDNSENTRY']._serialized_start=405
  _globals['_ZNETSTATICDNSENTRY']._serialized_end=460
  _globals['_IPSPEC']._serialized_start=463
  _globals['_IPSPEC']._serialized_end=733
  _globals['_DHCPOPTIONSIGNORE']._serialized_start=735
  _globals['_DHCPOPTIONSIGNORE']._serialized_end=786
  _globals['_PROBEENDPOINT']._serialized_start=788
  _globals['_PROBEENDPOINT']._serialized_end=831
  _globals['_CONNECTIVITYPROBE']._serialized_start=834
  _globals['_CONNECTIVITYPROBE']._serialized_end=985
# @@protoc_insertion_point(module_scope)
