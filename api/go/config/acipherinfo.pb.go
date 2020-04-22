// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.6.1
// source: config/acipherinfo.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Security Key Exchange Method
type KeyExchangeScheme int32

const (
	KeyExchangeScheme_KEA_NONE KeyExchangeScheme = 0
	KeyExchangeScheme_KEA_ECDH KeyExchangeScheme = 1
)

// Enum value maps for KeyExchangeScheme.
var (
	KeyExchangeScheme_name = map[int32]string{
		0: "KEA_NONE",
		1: "KEA_ECDH",
	}
	KeyExchangeScheme_value = map[string]int32{
		"KEA_NONE": 0,
		"KEA_ECDH": 1,
	}
)

func (x KeyExchangeScheme) Enum() *KeyExchangeScheme {
	p := new(KeyExchangeScheme)
	*p = x
	return p
}

func (x KeyExchangeScheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyExchangeScheme) Descriptor() protoreflect.EnumDescriptor {
	return file_config_acipherinfo_proto_enumTypes[0].Descriptor()
}

func (KeyExchangeScheme) Type() protoreflect.EnumType {
	return &file_config_acipherinfo_proto_enumTypes[0]
}

func (x KeyExchangeScheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyExchangeScheme.Descriptor instead.
func (KeyExchangeScheme) EnumDescriptor() ([]byte, []int) {
	return file_config_acipherinfo_proto_rawDescGZIP(), []int{0}
}

// Encryption Scheme for Cipher Payload
type EncryptionScheme int32

const (
	EncryptionScheme_SA_NONE        EncryptionScheme = 0
	EncryptionScheme_SA_AES_256_CFB EncryptionScheme = 1
)

// Enum value maps for EncryptionScheme.
var (
	EncryptionScheme_name = map[int32]string{
		0: "SA_NONE",
		1: "SA_AES_256_CFB",
	}
	EncryptionScheme_value = map[string]int32{
		"SA_NONE":        0,
		"SA_AES_256_CFB": 1,
	}
)

func (x EncryptionScheme) Enum() *EncryptionScheme {
	p := new(EncryptionScheme)
	*p = x
	return p
}

func (x EncryptionScheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptionScheme) Descriptor() protoreflect.EnumDescriptor {
	return file_config_acipherinfo_proto_enumTypes[1].Descriptor()
}

func (EncryptionScheme) Type() protoreflect.EnumType {
	return &file_config_acipherinfo_proto_enumTypes[1]
}

func (x EncryptionScheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptionScheme.Descriptor instead.
func (EncryptionScheme) EnumDescriptor() ([]byte, []int) {
	return file_config_acipherinfo_proto_rawDescGZIP(), []int{1}
}

type CipherHashAlgorithm int32

const (
	CipherHashAlgorithm_HASH_NONE           CipherHashAlgorithm = 0
	CipherHashAlgorithm_HASH_SHA256_16bytes CipherHashAlgorithm = 1 // hash with sha256, the 1st 16 bytes of result in 'senderCertHash'
	CipherHashAlgorithm_HASH_SHA256_32bytes CipherHashAlgorithm = 2 // hash with sha256, the whole 32 bytes of hash result
)

// Enum value maps for CipherHashAlgorithm.
var (
	CipherHashAlgorithm_name = map[int32]string{
		0: "HASH_NONE",
		1: "HASH_SHA256_16bytes",
		2: "HASH_SHA256_32bytes",
	}
	CipherHashAlgorithm_value = map[string]int32{
		"HASH_NONE":           0,
		"HASH_SHA256_16bytes": 1,
		"HASH_SHA256_32bytes": 2,
	}
)

func (x CipherHashAlgorithm) Enum() *CipherHashAlgorithm {
	p := new(CipherHashAlgorithm)
	*p = x
	return p
}

func (x CipherHashAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CipherHashAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_config_acipherinfo_proto_enumTypes[2].Descriptor()
}

func (CipherHashAlgorithm) Type() protoreflect.EnumType {
	return &file_config_acipherinfo_proto_enumTypes[2]
}

func (x CipherHashAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CipherHashAlgorithm.Descriptor instead.
func (CipherHashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_config_acipherinfo_proto_rawDescGZIP(), []int{2}
}

// Cipher information to decrypt Sensitive Data
type CipherContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cipher context id, key to this structure
	ContextId string `protobuf:"bytes,1,opt,name=contextId,proto3" json:"contextId,omitempty"`
	// algorithm used to compute hash for certificates
<<<<<<< HEAD
	HashScheme CipherHashAlgorithm `protobuf:"varint,2,opt,name=hashScheme,proto3,enum=CipherHashAlgorithm" json:"hashScheme,omitempty"`
=======
	HashScheme common.HashAlgorithm `protobuf:"varint,2,opt,name=hashScheme,proto3,enum=common.HashAlgorithm" json:"hashScheme,omitempty"`
>>>>>>> 47bb3b73... fixup! Removed DupliCate enum definitions and moved common nes to proto/common/common.go
	// for key exchange scheme, like ECDH etc.
	KeyExchangeScheme KeyExchangeScheme `protobuf:"varint,3,opt,name=keyExchangeScheme,proto3,enum=KeyExchangeScheme" json:"keyExchangeScheme,omitempty"`
	// for encrypting sensitive data, like AES256 etc.
	EncryptionScheme EncryptionScheme `protobuf:"varint,4,opt,name=encryptionScheme,proto3,enum=EncryptionScheme" json:"encryptionScheme,omitempty"`
	// device public certificate hash
	DeviceCertHash []byte `protobuf:"bytes,5,opt,name=deviceCertHash,proto3" json:"deviceCertHash,omitempty"`
	// controller certificate hash
	ControllerCertHash []byte `protobuf:"bytes,6,opt,name=controllerCertHash,proto3" json:"controllerCertHash,omitempty"`
}

func (x *CipherContext) Reset() {
	*x = CipherContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_acipherinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherContext) ProtoMessage() {}

func (x *CipherContext) ProtoReflect() protoreflect.Message {
	mi := &file_config_acipherinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherContext.ProtoReflect.Descriptor instead.
func (*CipherContext) Descriptor() ([]byte, []int) {
	return file_config_acipherinfo_proto_rawDescGZIP(), []int{0}
}

func (x *CipherContext) GetContextId() string {
	if x != nil {
		return x.ContextId
	}
	return ""
}

func (x *CipherContext) GetHashScheme() CipherHashAlgorithm {
	if x != nil {
		return x.HashScheme
	}
	return CipherHashAlgorithm_HASH_NONE
}

func (x *CipherContext) GetKeyExchangeScheme() KeyExchangeScheme {
	if x != nil {
		return x.KeyExchangeScheme
	}
	return KeyExchangeScheme_KEA_NONE
}

func (x *CipherContext) GetEncryptionScheme() EncryptionScheme {
	if x != nil {
		return x.EncryptionScheme
	}
	return EncryptionScheme_SA_NONE
}

func (x *CipherContext) GetDeviceCertHash() []byte {
	if x != nil {
		return x.DeviceCertHash
	}
	return nil
}

func (x *CipherContext) GetControllerCertHash() []byte {
	if x != nil {
		return x.ControllerCertHash
	}
	return nil
}

// Encrypted sensitive data information
type CipherBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cipher context id
	CipherContextId string `protobuf:"bytes,1,opt,name=cipherContextId,proto3" json:"cipherContextId,omitempty"`
	// Initial Value for Symmetric Key derivation
	InitialValue []byte `protobuf:"bytes,2,opt,name=initialValue,proto3" json:"initialValue,omitempty"`
	// encrypted sensitive data
	CipherData []byte `protobuf:"bytes,3,opt,name=cipherData,proto3" json:"cipherData,omitempty"`
	// sha256 of the plaintext sensitive data
	ClearTextSha256 []byte `protobuf:"bytes,4,opt,name=clearTextSha256,proto3" json:"clearTextSha256,omitempty"`
}

func (x *CipherBlock) Reset() {
	*x = CipherBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_acipherinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherBlock) ProtoMessage() {}

func (x *CipherBlock) ProtoReflect() protoreflect.Message {
	mi := &file_config_acipherinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherBlock.ProtoReflect.Descriptor instead.
func (*CipherBlock) Descriptor() ([]byte, []int) {
	return file_config_acipherinfo_proto_rawDescGZIP(), []int{1}
}

func (x *CipherBlock) GetCipherContextId() string {
	if x != nil {
		return x.CipherContextId
	}
	return ""
}

func (x *CipherBlock) GetInitialValue() []byte {
	if x != nil {
		return x.InitialValue
	}
	return nil
}

func (x *CipherBlock) GetCipherData() []byte {
	if x != nil {
		return x.CipherData
	}
	return nil
}

func (x *CipherBlock) GetClearTextSha256() []byte {
	if x != nil {
		return x.ClearTextSha256
	}
	return nil
}

// This message will be filled with the
// details to be encrypted and shared across
// wire for data in transit, by the controller
// for encryption
type EncryptionBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DsAPIKey          string `protobuf:"bytes,1,opt,name=dsAPIKey,proto3" json:"dsAPIKey,omitempty"`
	DsPassword        string `protobuf:"bytes,2,opt,name=dsPassword,proto3" json:"dsPassword,omitempty"`
	WifiUserName      string `protobuf:"bytes,3,opt,name=wifiUserName,proto3" json:"wifiUserName,omitempty"` // If the authentication type is EAP
	WifiPassword      string `protobuf:"bytes,4,opt,name=wifiPassword,proto3" json:"wifiPassword,omitempty"`
	ProtectedUserData string `protobuf:"bytes,5,opt,name=protectedUserData,proto3" json:"protectedUserData,omitempty"`
}

func (x *EncryptionBlock) Reset() {
	*x = EncryptionBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_acipherinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionBlock) ProtoMessage() {}

func (x *EncryptionBlock) ProtoReflect() protoreflect.Message {
	mi := &file_config_acipherinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionBlock.ProtoReflect.Descriptor instead.
func (*EncryptionBlock) Descriptor() ([]byte, []int) {
	return file_config_acipherinfo_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptionBlock) GetDsAPIKey() string {
	if x != nil {
		return x.DsAPIKey
	}
	return ""
}

func (x *EncryptionBlock) GetDsPassword() string {
	if x != nil {
		return x.DsPassword
	}
	return ""
}

func (x *EncryptionBlock) GetWifiUserName() string {
	if x != nil {
		return x.WifiUserName
	}
	return ""
}

func (x *EncryptionBlock) GetWifiPassword() string {
	if x != nil {
		return x.WifiPassword
	}
	return ""
}

func (x *EncryptionBlock) GetProtectedUserData() string {
	if x != nil {
		return x.ProtectedUserData
	}
	return ""
}

var File_config_acipherinfo_proto protoreflect.FileDescriptor

var file_config_acipherinfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
<<<<<<< HEAD
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x0d, 0x43,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x68, 0x61,
	0x73, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x12, 0x40, 0x0a, 0x11, 0x6b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x4b, 0x65,
	0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x11, 0x6b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52,
	0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0xa5, 0x01, 0x0a, 0x0b, 0x43, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72,
	0x54, 0x65, 0x78, 0x74, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x65, 0x78, 0x74, 0x53, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x73, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x73, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x69, 0x66, 0x69, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x69, 0x66, 0x69, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x69, 0x66, 0x69, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x69, 0x66,
	0x69, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x2f, 0x0a, 0x11, 0x4b, 0x65, 0x79, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x08,
	0x4b, 0x45, 0x41, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45,
	0x41, 0x5f, 0x45, 0x43, 0x44, 0x48, 0x10, 0x01, 0x2a, 0x33, 0x0a, 0x10, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x41, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x41, 0x5f,
	0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x01, 0x2a, 0x56, 0x0a,
	0x13, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x32,
	0x35, 0x36, 0x5f, 0x31, 0x36, 0x62, 0x79, 0x74, 0x65, 0x73, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13,
	0x48, 0x41, 0x53, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x5f, 0x33, 0x32, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x10, 0x02, 0x42, 0x47, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x65, 0x64,
	0x65, 0x64, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
=======
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbd, 0x02, 0x0a, 0x0d, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73,
	0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x0a, 0x68, 0x61, 0x73, 0x68,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x6b, 0x65, 0x79, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x12, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x11, 0x6b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x65, 0x72,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22,
	0xa5, 0x01, 0x0a, 0x0b, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x28, 0x0a, 0x0f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x65, 0x78, 0x74, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x65, 0x78,
	0x74, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x73, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x73, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x73, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x73, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x69, 0x66, 0x69, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77,
	0x69, 0x66, 0x69, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77,
	0x69, 0x66, 0x69, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x77, 0x69, 0x66, 0x69, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2a, 0x2f, 0x0a,
	0x11, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x41, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x45, 0x41, 0x5f, 0x45, 0x43, 0x44, 0x48, 0x10, 0x01, 0x2a, 0x33,
	0x0a, 0x10, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x41, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x41, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x43, 0x46,
	0x42, 0x10, 0x01, 0x42, 0x47, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x65, 0x64, 0x65, 0x64,
	0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
>>>>>>> 47bb3b73... fixup! Removed DupliCate enum definitions and moved common nes to proto/common/common.go
}

var (
	file_config_acipherinfo_proto_rawDescOnce sync.Once
	file_config_acipherinfo_proto_rawDescData = file_config_acipherinfo_proto_rawDesc
)

func file_config_acipherinfo_proto_rawDescGZIP() []byte {
	file_config_acipherinfo_proto_rawDescOnce.Do(func() {
		file_config_acipherinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_acipherinfo_proto_rawDescData)
	})
	return file_config_acipherinfo_proto_rawDescData
}

var file_config_acipherinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_config_acipherinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_acipherinfo_proto_goTypes = []interface{}{
<<<<<<< HEAD
	(KeyExchangeScheme)(0),   // 0: KeyExchangeScheme
	(EncryptionScheme)(0),    // 1: EncryptionScheme
	(CipherHashAlgorithm)(0), // 2: cipherHashAlgorithm
	(*CipherContext)(nil),    // 3: CipherContext
	(*CipherBlock)(nil),      // 4: CipherBlock
	(*EncryptionBlock)(nil),  // 5: EncryptionBlock
}
var file_config_acipherinfo_proto_depIdxs = []int32{
	2, // 0: CipherContext.hashScheme:type_name -> cipherHashAlgorithm
=======
	(KeyExchangeScheme)(0),    // 0: KeyExchangeScheme
	(EncryptionScheme)(0),     // 1: EncryptionScheme
	(*CipherContext)(nil),     // 2: CipherContext
	(*CipherBlock)(nil),       // 3: CipherBlock
	(*EncryptionBlock)(nil),   // 4: EncryptionBlock
	(common.HashAlgorithm)(0), // 5: common.HashAlgorithm
}
var file_config_acipherinfo_proto_depIdxs = []int32{
	5, // 0: CipherContext.hashScheme:type_name -> common.HashAlgorithm
>>>>>>> 47bb3b73... fixup! Removed DupliCate enum definitions and moved common nes to proto/common/common.go
	0, // 1: CipherContext.keyExchangeScheme:type_name -> KeyExchangeScheme
	1, // 2: CipherContext.encryptionScheme:type_name -> EncryptionScheme
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_config_acipherinfo_proto_init() }
func file_config_acipherinfo_proto_init() {
	if File_config_acipherinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_acipherinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherContext); i {
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
		file_config_acipherinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherBlock); i {
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
		file_config_acipherinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionBlock); i {
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
			RawDescriptor: file_config_acipherinfo_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_acipherinfo_proto_goTypes,
		DependencyIndexes: file_config_acipherinfo_proto_depIdxs,
		EnumInfos:         file_config_acipherinfo_proto_enumTypes,
		MessageInfos:      file_config_acipherinfo_proto_msgTypes,
	}.Build()
	File_config_acipherinfo_proto = out.File
	file_config_acipherinfo_proto_rawDesc = nil
	file_config_acipherinfo_proto_goTypes = nil
	file_config_acipherinfo_proto_depIdxs = nil
}