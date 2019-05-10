// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vm.proto

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

// For now we need to tell the device which virtualization mode
// to use. Later we might use a single one for all VMs (on any particular
// ISA). If we end up keeping this we should make the names be less
// tied to a particular hypervisor.
type VmMode int32

const (
	VmMode_PV  VmMode = 0
	VmMode_HVM VmMode = 1
)

var VmMode_name = map[int32]string{
	0: "PV",
	1: "HVM",
}

var VmMode_value = map[string]int32{
	"PV":  0,
	"HVM": 1,
}

func (x VmMode) String() string {
	return proto.EnumName(VmMode_name, int32(x))
}

func (VmMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cab246c8c7c5372d, []int{0}
}

type VmConfig struct {
	Kernel               string   `protobuf:"bytes,1,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Ramdisk              string   `protobuf:"bytes,2,opt,name=ramdisk,proto3" json:"ramdisk,omitempty"`
	Memory               uint32   `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Maxmem               uint32   `protobuf:"varint,4,opt,name=maxmem,proto3" json:"maxmem,omitempty"`
	Vcpus                uint32   `protobuf:"varint,5,opt,name=vcpus,proto3" json:"vcpus,omitempty"`
	Maxcpus              uint32   `protobuf:"varint,6,opt,name=maxcpus,proto3" json:"maxcpus,omitempty"`
	Rootdev              string   `protobuf:"bytes,7,opt,name=rootdev,proto3" json:"rootdev,omitempty"`
	Extraargs            string   `protobuf:"bytes,8,opt,name=extraargs,proto3" json:"extraargs,omitempty"`
	Bootloader           string   `protobuf:"bytes,9,opt,name=bootloader,proto3" json:"bootloader,omitempty"`
	Cpus                 string   `protobuf:"bytes,10,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Devicetree           string   `protobuf:"bytes,11,opt,name=devicetree,proto3" json:"devicetree,omitempty"`
	Dtdev                []string `protobuf:"bytes,12,rep,name=dtdev,proto3" json:"dtdev,omitempty"`
	Irqs                 []uint32 `protobuf:"varint,13,rep,packed,name=irqs,proto3" json:"irqs,omitempty"`
	Iomem                []string `protobuf:"bytes,14,rep,name=iomem,proto3" json:"iomem,omitempty"`
	VirtualizationMode   VmMode   `protobuf:"varint,15,opt,name=virtualizationMode,proto3,enum=VmMode" json:"virtualizationMode,omitempty"`
	EnableVnc            bool     `protobuf:"varint,16,opt,name=enableVnc,proto3" json:"enableVnc,omitempty"`
	VncDisplay           uint32   `protobuf:"varint,17,opt,name=vncDisplay,proto3" json:"vncDisplay,omitempty"`
	VncPasswd            string   `protobuf:"bytes,18,opt,name=vncPasswd,proto3" json:"vncPasswd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmConfig) Reset()         { *m = VmConfig{} }
func (m *VmConfig) String() string { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()    {}
func (*VmConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab246c8c7c5372d, []int{0}
}

func (m *VmConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmConfig.Unmarshal(m, b)
}
func (m *VmConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmConfig.Marshal(b, m, deterministic)
}
func (m *VmConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmConfig.Merge(m, src)
}
func (m *VmConfig) XXX_Size() int {
	return xxx_messageInfo_VmConfig.Size(m)
}
func (m *VmConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VmConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VmConfig proto.InternalMessageInfo

func (m *VmConfig) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *VmConfig) GetRamdisk() string {
	if m != nil {
		return m.Ramdisk
	}
	return ""
}

func (m *VmConfig) GetMemory() uint32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *VmConfig) GetMaxmem() uint32 {
	if m != nil {
		return m.Maxmem
	}
	return 0
}

func (m *VmConfig) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *VmConfig) GetMaxcpus() uint32 {
	if m != nil {
		return m.Maxcpus
	}
	return 0
}

func (m *VmConfig) GetRootdev() string {
	if m != nil {
		return m.Rootdev
	}
	return ""
}

func (m *VmConfig) GetExtraargs() string {
	if m != nil {
		return m.Extraargs
	}
	return ""
}

func (m *VmConfig) GetBootloader() string {
	if m != nil {
		return m.Bootloader
	}
	return ""
}

func (m *VmConfig) GetCpus() string {
	if m != nil {
		return m.Cpus
	}
	return ""
}

func (m *VmConfig) GetDevicetree() string {
	if m != nil {
		return m.Devicetree
	}
	return ""
}

func (m *VmConfig) GetDtdev() []string {
	if m != nil {
		return m.Dtdev
	}
	return nil
}

func (m *VmConfig) GetIrqs() []uint32 {
	if m != nil {
		return m.Irqs
	}
	return nil
}

func (m *VmConfig) GetIomem() []string {
	if m != nil {
		return m.Iomem
	}
	return nil
}

func (m *VmConfig) GetVirtualizationMode() VmMode {
	if m != nil {
		return m.VirtualizationMode
	}
	return VmMode_PV
}

func (m *VmConfig) GetEnableVnc() bool {
	if m != nil {
		return m.EnableVnc
	}
	return false
}

func (m *VmConfig) GetVncDisplay() uint32 {
	if m != nil {
		return m.VncDisplay
	}
	return 0
}

func (m *VmConfig) GetVncPasswd() string {
	if m != nil {
		return m.VncPasswd
	}
	return ""
}

func init() {
	proto.RegisterEnum("VmMode", VmMode_name, VmMode_value)
	proto.RegisterType((*VmConfig)(nil), "VmConfig")
}

func init() { proto.RegisterFile("vm.proto", fileDescriptor_cab246c8c7c5372d) }

var fileDescriptor_cab246c8c7c5372d = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x40, 0x49, 0xd3, 0x66, 0x77, 0x0d, 0x5b, 0x16, 0x0b, 0x21, 0x23, 0x21, 0x88, 0xe0, 0x12,
	0x71, 0x48, 0x24, 0x38, 0x70, 0x07, 0x84, 0xb8, 0x54, 0xaa, 0x72, 0xc8, 0x81, 0x9b, 0xd7, 0x1e,
	0x16, 0x6b, 0xe3, 0xcc, 0x62, 0x27, 0x66, 0xdb, 0x7f, 0x47, 0x42, 0x1e, 0xa7, 0x6c, 0x0f, 0xbd,
	0xf9, 0xbd, 0x99, 0xf1, 0x8c, 0x47, 0x66, 0xcb, 0x60, 0xeb, 0x83, 0xc3, 0x11, 0xdf, 0xfe, 0xcd,
	0xd9, 0xb2, 0xb3, 0x5f, 0x70, 0xf8, 0x69, 0x76, 0xfc, 0x05, 0x2b, 0xf6, 0xe0, 0x06, 0xe8, 0x45,
	0x56, 0x66, 0xd5, 0xaa, 0x9d, 0x89, 0x0b, 0xb6, 0x70, 0xd2, 0x6a, 0xe3, 0xf7, 0xe2, 0x8c, 0x02,
	0x77, 0x18, 0x2b, 0x2c, 0x58, 0x74, 0x37, 0x22, 0x2f, 0xb3, 0x6a, 0xdd, 0xce, 0x44, 0x5e, 0x1e,
	0x2d, 0x58, 0x71, 0x3e, 0x7b, 0x22, 0xfe, 0x9c, 0x5d, 0x04, 0x75, 0x98, 0xbc, 0xb8, 0x20, 0x9d,
	0x20, 0xde, 0x6f, 0xe5, 0x91, 0x7c, 0x41, 0xfe, 0x0e, 0xa9, 0x33, 0xe2, 0xa8, 0x21, 0x88, 0xc5,
	0xdc, 0x39, 0x21, 0x7f, 0xc5, 0x56, 0x70, 0x1c, 0x9d, 0x94, 0x6e, 0xe7, 0xc5, 0x92, 0x62, 0x27,
	0xc1, 0x5f, 0x33, 0xb6, 0x45, 0x1c, 0x7b, 0x94, 0x1a, 0x9c, 0x58, 0x51, 0xf8, 0x9e, 0xe1, 0x9c,
	0x9d, 0x53, 0x3b, 0x46, 0x11, 0x3a, 0xc7, 0x1a, 0x0d, 0xc1, 0x28, 0x18, 0x1d, 0x80, 0x78, 0x9c,
	0x6a, 0x4e, 0x26, 0xce, 0xae, 0x69, 0x92, 0x27, 0x65, 0x5e, 0xad, 0xda, 0x04, 0xf1, 0x26, 0xe3,
	0x7e, 0x7b, 0xb1, 0x2e, 0xf3, 0x6a, 0xdd, 0xd2, 0x39, 0x66, 0x1a, 0x8c, 0x8f, 0xbf, 0x4c, 0x99,
	0x04, 0xfc, 0x13, 0xe3, 0xc1, 0xb8, 0x71, 0x92, 0xbd, 0xb9, 0x95, 0xa3, 0xc1, 0xe1, 0x0a, 0x35,
	0x88, 0xa7, 0x65, 0x56, 0x5d, 0x7e, 0x58, 0xd4, 0x9d, 0x8d, 0xd8, 0x3e, 0x90, 0x42, 0x4f, 0x1d,
	0xe4, 0xb6, 0x87, 0x6e, 0x50, 0x62, 0x53, 0x66, 0xd5, 0xb2, 0x3d, 0x89, 0x38, 0x76, 0x18, 0xd4,
	0x57, 0xe3, 0x0f, 0xbd, 0xbc, 0x11, 0xcf, 0x68, 0x7f, 0xf7, 0x4c, 0xac, 0x0e, 0x83, 0xba, 0x96,
	0xde, 0xff, 0xd1, 0x82, 0xa7, 0x45, 0xfd, 0x17, 0xef, 0x5f, 0xb2, 0x22, 0x75, 0xe6, 0x05, 0x3b,
	0xbb, 0xee, 0x36, 0x8f, 0xf8, 0x82, 0xe5, 0xdf, 0xbb, 0xab, 0x4d, 0xf6, 0xf9, 0x1b, 0x7b, 0xa3,
	0xd0, 0xd6, 0xb7, 0xa0, 0x41, 0xcb, 0x5a, 0xf5, 0x38, 0xe9, 0x7a, 0xf2, 0xe0, 0xe2, 0x42, 0xd2,
	0xef, 0xf9, 0xf1, 0x6e, 0x67, 0xc6, 0x5f, 0xd3, 0xb6, 0x56, 0x68, 0x9b, 0x94, 0xd7, 0x40, 0x80,
	0xc6, 0xeb, 0x7d, 0xb3, 0xc3, 0x46, 0xd1, 0xaf, 0xda, 0x16, 0x94, 0xfb, 0xf1, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8e, 0x17, 0xa4, 0xf9, 0x75, 0x02, 0x00, 0x00,
}
