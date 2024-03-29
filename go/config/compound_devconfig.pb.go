// Copyright(c) 2023 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: config/compound_devconfig.proto

package config

import (
	auth "github.com/lf-edge/eve-api/go/auth"
	profile "github.com/lf-edge/eve-api/go/profile"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompoundEdgeDevConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// See the details about the field in the corresponding messages
	// `LocalDevInfo` and `LocalAppInfo` defined in `local_profile.proto`.
	LastCmdTimestamp uint64         `protobuf:"varint,1,opt,name=last_cmd_timestamp,json=lastCmdTimestamp,proto3" json:"last_cmd_timestamp,omitempty"`
	CfgReq           *ConfigRequest `protobuf:"bytes,2,opt,name=cfg_req,json=cfgReq,proto3" json:"cfg_req,omitempty"`
}

func (x *CompoundEdgeDevConfigRequest) Reset() {
	*x = CompoundEdgeDevConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_compound_devconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompoundEdgeDevConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompoundEdgeDevConfigRequest) ProtoMessage() {}

func (x *CompoundEdgeDevConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_compound_devconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompoundEdgeDevConfigRequest.ProtoReflect.Descriptor instead.
func (*CompoundEdgeDevConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_compound_devconfig_proto_rawDescGZIP(), []int{0}
}

func (x *CompoundEdgeDevConfigRequest) GetLastCmdTimestamp() uint64 {
	if x != nil {
		return x.LastCmdTimestamp
	}
	return 0
}

func (x *CompoundEdgeDevConfigRequest) GetCfgReq() *ConfigRequest {
	if x != nil {
		return x.CfgReq
	}
	return nil
}

// CompoundEdgeDevConfig message combines regular edge config, packed in auth
// envelope, node/applications commands and radio config. This message is
// needed for the LOC case, when a single "/compound-config" endpoint can be
// used.
type CompoundEdgeDevConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// See the details about the field in the corresponding messages
	// `LocalDevCommand` and `AppCommand` defined in `local_profile.proto`.
	Timestamp       uint64                   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ProtectedConfig *auth.AuthContainer      `protobuf:"bytes,2,opt,name=protected_config,json=protectedConfig,proto3" json:"protected_config,omitempty"`
	DevCmd          *profile.LocalDevCmd     `protobuf:"bytes,3,opt,name=dev_cmd,json=devCmd,proto3" json:"dev_cmd,omitempty"`
	AppCmdList      *profile.LocalAppCmdList `protobuf:"bytes,4,opt,name=app_cmd_list,json=appCmdList,proto3" json:"app_cmd_list,omitempty"`
	RadioConfig     *profile.RadioConfig     `protobuf:"bytes,5,opt,name=radio_config,json=radioConfig,proto3" json:"radio_config,omitempty"`
}

func (x *CompoundEdgeDevConfig) Reset() {
	*x = CompoundEdgeDevConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_compound_devconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompoundEdgeDevConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompoundEdgeDevConfig) ProtoMessage() {}

func (x *CompoundEdgeDevConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_compound_devconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompoundEdgeDevConfig.ProtoReflect.Descriptor instead.
func (*CompoundEdgeDevConfig) Descriptor() ([]byte, []int) {
	return file_config_compound_devconfig_proto_rawDescGZIP(), []int{1}
}

func (x *CompoundEdgeDevConfig) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CompoundEdgeDevConfig) GetProtectedConfig() *auth.AuthContainer {
	if x != nil {
		return x.ProtectedConfig
	}
	return nil
}

func (x *CompoundEdgeDevConfig) GetDevCmd() *profile.LocalDevCmd {
	if x != nil {
		return x.DevCmd
	}
	return nil
}

func (x *CompoundEdgeDevConfig) GetAppCmdList() *profile.LocalAppCmdList {
	if x != nil {
		return x.AppCmdList
	}
	return nil
}

func (x *CompoundEdgeDevConfig) GetRadioConfig() *profile.RadioConfig {
	if x != nil {
		return x.RadioConfig
	}
	return nil
}

var File_config_compound_devconfig_proto protoreflect.FileDescriptor

var file_config_compound_devconfig_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x64, 0x65, 0x76, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x64, 0x67, 0x65, 0x44,
	0x65, 0x76, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x43, 0x6d, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a,
	0x07, 0x63, 0x66, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x63, 0x66, 0x67, 0x52, 0x65, 0x71, 0x22, 0xd5, 0x02, 0x0a,
	0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x64, 0x67, 0x65, 0x44, 0x65, 0x76,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x4d, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x3c, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x63, 0x6d, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x44, 0x65, 0x76, 0x43, 0x6d, 0x64, 0x52, 0x06, 0x64, 0x65, 0x76, 0x43, 0x6d,
	0x64, 0x12, 0x49, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x5f, 0x63, 0x6d, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0a, 0x61, 0x70, 0x70, 0x43, 0x6d, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0c,
	0x72, 0x61, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x61, 0x64, 0x69,
	0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_compound_devconfig_proto_rawDescOnce sync.Once
	file_config_compound_devconfig_proto_rawDescData = file_config_compound_devconfig_proto_rawDesc
)

func file_config_compound_devconfig_proto_rawDescGZIP() []byte {
	file_config_compound_devconfig_proto_rawDescOnce.Do(func() {
		file_config_compound_devconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_compound_devconfig_proto_rawDescData)
	})
	return file_config_compound_devconfig_proto_rawDescData
}

var file_config_compound_devconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_compound_devconfig_proto_goTypes = []interface{}{
	(*CompoundEdgeDevConfigRequest)(nil), // 0: org.lfedge.eve.config.CompoundEdgeDevConfigRequest
	(*CompoundEdgeDevConfig)(nil),        // 1: org.lfedge.eve.config.CompoundEdgeDevConfig
	(*ConfigRequest)(nil),                // 2: org.lfedge.eve.config.ConfigRequest
	(*auth.AuthContainer)(nil),           // 3: org.lfedge.eve.auth.AuthContainer
	(*profile.LocalDevCmd)(nil),          // 4: org.lfedge.eve.profile.LocalDevCmd
	(*profile.LocalAppCmdList)(nil),      // 5: org.lfedge.eve.profile.LocalAppCmdList
	(*profile.RadioConfig)(nil),          // 6: org.lfedge.eve.profile.RadioConfig
}
var file_config_compound_devconfig_proto_depIdxs = []int32{
	2, // 0: org.lfedge.eve.config.CompoundEdgeDevConfigRequest.cfg_req:type_name -> org.lfedge.eve.config.ConfigRequest
	3, // 1: org.lfedge.eve.config.CompoundEdgeDevConfig.protected_config:type_name -> org.lfedge.eve.auth.AuthContainer
	4, // 2: org.lfedge.eve.config.CompoundEdgeDevConfig.dev_cmd:type_name -> org.lfedge.eve.profile.LocalDevCmd
	5, // 3: org.lfedge.eve.config.CompoundEdgeDevConfig.app_cmd_list:type_name -> org.lfedge.eve.profile.LocalAppCmdList
	6, // 4: org.lfedge.eve.config.CompoundEdgeDevConfig.radio_config:type_name -> org.lfedge.eve.profile.RadioConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_compound_devconfig_proto_init() }
func file_config_compound_devconfig_proto_init() {
	if File_config_compound_devconfig_proto != nil {
		return
	}
	file_config_devconfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_compound_devconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompoundEdgeDevConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_compound_devconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompoundEdgeDevConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_compound_devconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_compound_devconfig_proto_goTypes,
		DependencyIndexes: file_config_compound_devconfig_proto_depIdxs,
		MessageInfos:      file_config_compound_devconfig_proto_msgTypes,
	}.Build()
	File_config_compound_devconfig_proto = out.File
	file_config_compound_devconfig_proto_rawDesc = nil
	file_config_compound_devconfig_proto_goTypes = nil
	file_config_compound_devconfig_proto_depIdxs = nil
}
