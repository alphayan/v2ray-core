package dokodemo

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	net "v2ray.com/core/common/net"
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

type Config struct {
	Address *net.IPOrDomain `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32          `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// List of networks that the Dokodemo accepts.
	// Deprecated. Use networks.
	NetworkList *net.NetworkList `protobuf:"bytes,3,opt,name=network_list,json=networkList,proto3" json:"network_list,omitempty"` // Deprecated: Do not use.
	// List of networks that the Dokodemo accepts.
	Networks             []net.Network `protobuf:"varint,7,rep,packed,name=networks,proto3,enum=v2ray.core.common.net.Network" json:"networks,omitempty"`
	Timeout              uint32        `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"` // Deprecated: Do not use.
	FollowRedirect       bool          `protobuf:"varint,5,opt,name=follow_redirect,json=followRedirect,proto3" json:"follow_redirect,omitempty"`
	UserLevel            uint32        `protobuf:"varint,6,opt,name=user_level,json=userLevel,proto3" json:"user_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_de04411d7254f312, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetAddress() *net.IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Config) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Deprecated: Do not use.
func (m *Config) GetNetworkList() *net.NetworkList {
	if m != nil {
		return m.NetworkList
	}
	return nil
}

func (m *Config) GetNetworks() []net.Network {
	if m != nil {
		return m.Networks
	}
	return nil
}

// Deprecated: Do not use.
func (m *Config) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Config) GetFollowRedirect() bool {
	if m != nil {
		return m.FollowRedirect
	}
	return false
}

func (m *Config) GetUserLevel() uint32 {
	if m != nil {
		return m.UserLevel
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.proxy.dokodemo.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/proxy/dokodemo/config.proto", fileDescriptor_de04411d7254f312)
}

var fileDescriptor_de04411d7254f312 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0xd3, 0xc2, 0x03, 0xde, 0xf2, 0x1e, 0x26, 0x7b, 0x5a, 0x8c, 0x98, 0xca, 0x85, 0xc6,
	0xc3, 0x36, 0xa9, 0x37, 0xbd, 0x01, 0x89, 0x21, 0x21, 0x4a, 0xf6, 0xe0, 0xc1, 0x0b, 0xa9, 0xed,
	0x60, 0x1a, 0xba, 0x3b, 0x64, 0xbb, 0x80, 0x7c, 0x25, 0x3f, 0x88, 0x9f, 0xcb, 0x74, 0xdb, 0xaa,
	0x31, 0x81, 0xdb, 0xec, 0x7f, 0x7f, 0xfb, 0x9b, 0xd9, 0x0c, 0xb9, 0xde, 0x85, 0x3a, 0x3a, 0xf0,
	0x18, 0x65, 0x10, 0xa3, 0x86, 0x60, 0xa3, 0xf1, 0xed, 0x10, 0x24, 0xb8, 0xc6, 0x04, 0x24, 0x06,
	0x31, 0xaa, 0x55, 0xfa, 0xca, 0x37, 0x1a, 0x0d, 0xd2, 0x7e, 0xcd, 0x6a, 0xe0, 0x96, 0xe3, 0x35,
	0x77, 0x3e, 0xfa, 0xa5, 0x89, 0x51, 0x4a, 0x54, 0x81, 0x02, 0x13, 0x44, 0x49, 0xa2, 0x21, 0xcf,
	0x4b, 0xc7, 0x29, 0x50, 0x81, 0xd9, 0xa3, 0x5e, 0x97, 0xe0, 0xf0, 0xc3, 0x25, 0xad, 0x89, 0xed,
	0x4e, 0xef, 0x48, 0xbb, 0x92, 0x30, 0xc7, 0x73, 0xfc, 0x6e, 0x78, 0xc5, 0x7f, 0x4c, 0x52, 0x1a,
	0xb8, 0x02, 0xc3, 0x67, 0x8b, 0x47, 0x3d, 0x45, 0x19, 0xa5, 0x4a, 0xd4, 0x2f, 0x28, 0x25, 0xcd,
	0x0d, 0x6a, 0xc3, 0x5c, 0xcf, 0xf1, 0xff, 0x0b, 0x5b, 0xd3, 0x19, 0xf9, 0x57, 0x35, 0x5b, 0x66,
	0x69, 0x6e, 0x58, 0xc3, 0x5a, 0x87, 0x47, 0xac, 0x0f, 0x25, 0x3a, 0x4f, 0x73, 0x33, 0x76, 0x99,
	0x23, 0xba, 0xea, 0x3b, 0xa0, 0xb7, 0xa4, 0x53, 0x1d, 0x73, 0xd6, 0xf6, 0x1a, 0x7e, 0x2f, 0xbc,
	0x3c, 0xad, 0x11, 0x5f, 0x3c, 0xbd, 0x20, 0x6d, 0x93, 0x4a, 0xc0, 0xad, 0x61, 0xcd, 0x62, 0x3a,
	0x6b, 0xaf, 0x23, 0x3a, 0x22, 0x67, 0x2b, 0xcc, 0x32, 0xdc, 0x2f, 0x35, 0x24, 0xa9, 0x86, 0xd8,
	0xb0, 0x3f, 0x9e, 0xe3, 0x77, 0x44, 0xaf, 0x8c, 0x45, 0x95, 0xd2, 0x01, 0x21, 0xdb, 0x1c, 0xf4,
	0x32, 0x83, 0x1d, 0x64, 0xac, 0x65, 0xff, 0xf9, 0xb7, 0x48, 0xe6, 0x45, 0x30, 0xbe, 0x27, 0x83,
	0x18, 0x25, 0x3f, 0xba, 0xbb, 0x85, 0xf3, 0xdc, 0xa9, 0xeb, 0x77, 0xb7, 0xff, 0x14, 0x8a, 0xe8,
	0xc0, 0x27, 0x05, 0xb7, 0xb0, 0xdc, 0xb4, 0xba, 0x7b, 0x69, 0xd9, 0xc5, 0xdc, 0x7c, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xec, 0xf1, 0xbb, 0x96, 0x33, 0x02, 0x00, 0x00,
}
