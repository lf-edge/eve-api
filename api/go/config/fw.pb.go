// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fw.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ACEMatch struct {
	// FIXME: We should convert this to enum
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACEMatch) Reset()         { *m = ACEMatch{} }
func (m *ACEMatch) String() string { return proto.CompactTextString(m) }
func (*ACEMatch) ProtoMessage()    {}
func (*ACEMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_505e7efac08d3ba9, []int{0}
}

func (m *ACEMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACEMatch.Unmarshal(m, b)
}
func (m *ACEMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACEMatch.Marshal(b, m, deterministic)
}
func (m *ACEMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACEMatch.Merge(m, src)
}
func (m *ACEMatch) XXX_Size() int {
	return xxx_messageInfo_ACEMatch.Size(m)
}
func (m *ACEMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ACEMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ACEMatch proto.InternalMessageInfo

func (m *ACEMatch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ACEMatch) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ACEAction struct {
	Drop bool `protobuf:"varint,1,opt,name=drop,proto3" json:"drop,omitempty"`
	// limit action, and its associated parameter
	Limit      bool   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Limitrate  uint32 `protobuf:"varint,3,opt,name=limitrate,proto3" json:"limitrate,omitempty"`
	Limitunit  string `protobuf:"bytes,4,opt,name=limitunit,proto3" json:"limitunit,omitempty"`
	Limitburst uint32 `protobuf:"varint,5,opt,name=limitburst,proto3" json:"limitburst,omitempty"`
	// port map action, and its assoicated paramtere
	Portmap              bool     `protobuf:"varint,6,opt,name=portmap,proto3" json:"portmap,omitempty"`
	AppPort              uint32   `protobuf:"varint,7,opt,name=appPort,proto3" json:"appPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACEAction) Reset()         { *m = ACEAction{} }
func (m *ACEAction) String() string { return proto.CompactTextString(m) }
func (*ACEAction) ProtoMessage()    {}
func (*ACEAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_505e7efac08d3ba9, []int{1}
}

func (m *ACEAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACEAction.Unmarshal(m, b)
}
func (m *ACEAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACEAction.Marshal(b, m, deterministic)
}
func (m *ACEAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACEAction.Merge(m, src)
}
func (m *ACEAction) XXX_Size() int {
	return xxx_messageInfo_ACEAction.Size(m)
}
func (m *ACEAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ACEAction.DiscardUnknown(m)
}

var xxx_messageInfo_ACEAction proto.InternalMessageInfo

func (m *ACEAction) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *ACEAction) GetLimit() bool {
	if m != nil {
		return m.Limit
	}
	return false
}

func (m *ACEAction) GetLimitrate() uint32 {
	if m != nil {
		return m.Limitrate
	}
	return 0
}

func (m *ACEAction) GetLimitunit() string {
	if m != nil {
		return m.Limitunit
	}
	return ""
}

func (m *ACEAction) GetLimitburst() uint32 {
	if m != nil {
		return m.Limitburst
	}
	return 0
}

func (m *ACEAction) GetPortmap() bool {
	if m != nil {
		return m.Portmap
	}
	return false
}

func (m *ACEAction) GetAppPort() uint32 {
	if m != nil {
		return m.AppPort
	}
	return 0
}

type ACE struct {
	Matches []*ACEMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	// Expect only single action...repeated here is
	// for future work.
	Actions              []*ACEAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ACE) Reset()         { *m = ACE{} }
func (m *ACE) String() string { return proto.CompactTextString(m) }
func (*ACE) ProtoMessage()    {}
func (*ACE) Descriptor() ([]byte, []int) {
	return fileDescriptor_505e7efac08d3ba9, []int{2}
}

func (m *ACE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACE.Unmarshal(m, b)
}
func (m *ACE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACE.Marshal(b, m, deterministic)
}
func (m *ACE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACE.Merge(m, src)
}
func (m *ACE) XXX_Size() int {
	return xxx_messageInfo_ACE.Size(m)
}
func (m *ACE) XXX_DiscardUnknown() {
	xxx_messageInfo_ACE.DiscardUnknown(m)
}

var xxx_messageInfo_ACE proto.InternalMessageInfo

func (m *ACE) GetMatches() []*ACEMatch {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *ACE) GetActions() []*ACEAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ACEMatch)(nil), "ACEMatch")
	proto.RegisterType((*ACEAction)(nil), "ACEAction")
	proto.RegisterType((*ACE)(nil), "ACE")
}

func init() { proto.RegisterFile("fw.proto", fileDescriptor_505e7efac08d3ba9) }

var fileDescriptor_505e7efac08d3ba9 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x95, 0xfe, 0x25, 0x39, 0x9f, 0xbe, 0xc5, 0x62, 0xf0, 0x80, 0xa0, 0x6a, 0x19, 0x3a,
	0x39, 0x12, 0x70, 0x03, 0xa1, 0x2a, 0x1b, 0x52, 0xe5, 0x91, 0xcd, 0x71, 0xdc, 0xd4, 0x22, 0x89,
	0x2d, 0xc7, 0x0e, 0x82, 0xfb, 0xe3, 0xbe, 0x50, 0x4e, 0x94, 0x96, 0xed, 0xbc, 0x3f, 0x8f, 0x2c,
	0xbf, 0x90, 0x9c, 0x3e, 0x99, 0x75, 0xc6, 0x9b, 0xcd, 0x33, 0x24, 0xf9, 0xfe, 0xf0, 0x26, 0xbc,
	0x3c, 0x13, 0x02, 0x0b, 0xff, 0x65, 0x15, 0x8d, 0xd6, 0xd1, 0x2e, 0xe5, 0x78, 0x93, 0x1b, 0x58,
	0xf6, 0xa2, 0x0e, 0x8a, 0xce, 0xd0, 0x1c, 0xc5, 0xe6, 0x27, 0x82, 0x34, 0xdf, 0x1f, 0x72, 0xe9,
	0xb5, 0x69, 0x07, 0xae, 0x74, 0xc6, 0x22, 0x97, 0x70, 0xbc, 0x07, 0xae, 0xd6, 0x8d, 0xf6, 0xc8,
	0x25, 0x7c, 0x14, 0xe4, 0x16, 0x52, 0x3c, 0x9c, 0xf0, 0x8a, 0xce, 0xd7, 0xd1, 0xee, 0x3f, 0xbf,
	0x1a, 0x97, 0x34, 0xb4, 0xda, 0xd3, 0x05, 0xbe, 0x77, 0x35, 0xc8, 0x1d, 0x00, 0x8a, 0x22, 0xb8,
	0xce, 0xd3, 0x25, 0xc2, 0x7f, 0x1c, 0x42, 0x21, 0xb6, 0xc6, 0xf9, 0x46, 0x58, 0xba, 0xc2, 0x37,
	0x27, 0x39, 0x24, 0xc2, 0xda, 0xa3, 0x71, 0x9e, 0xc6, 0x88, 0x4d, 0x72, 0x73, 0x84, 0x79, 0xbe,
	0x3f, 0x90, 0x2d, 0xc4, 0xcd, 0xb0, 0x80, 0xea, 0x68, 0xb4, 0x9e, 0xef, 0xfe, 0x3d, 0xa6, 0x6c,
	0x1a, 0x85, 0x4f, 0x09, 0x79, 0x80, 0x58, 0xe0, 0x7f, 0x3b, 0x3a, 0xc3, 0x12, 0xb0, 0xcb, 0x04,
	0x7c, 0x8a, 0x5e, 0x5e, 0xe1, 0x5e, 0x9a, 0x86, 0x7d, 0xab, 0x52, 0x95, 0x82, 0xc9, 0xda, 0x84,
	0x92, 0x85, 0x4e, 0xb9, 0x5e, 0x4b, 0x35, 0x4e, 0xfe, 0xbe, 0xad, 0xb4, 0x3f, 0x87, 0x82, 0x49,
	0xd3, 0x64, 0x63, 0x2f, 0x53, 0xbd, 0xca, 0xba, 0xf2, 0x23, 0xab, 0x4c, 0x26, 0x4d, 0x7b, 0xd2,
	0x55, 0xb1, 0xc2, 0xee, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x89, 0x2f, 0x04, 0xaa,
	0x01, 0x00, 0x00,
}
