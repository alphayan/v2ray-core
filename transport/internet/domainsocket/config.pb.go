package domainsocket

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
	// Path of the domain socket. This overrides the IP/Port parameter from upstream caller.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Abstract speicifies whether to use abstract namespace or not.
	// Traditionally Unix domain socket is file system based. Abstract domain socket can be used without acquiring file lock.
	Abstract             bool     `protobuf:"varint,2,opt,name=abstract,proto3" json:"abstract,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_76473d52e3e3815d, []int{0}
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

func (m *Config) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Config) GetAbstract() bool {
	if m != nil {
		return m.Abstract
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.domainsocket.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/domainsocket/config.proto", fileDescriptor_76473d52e3e3815d)
}

var fileDescriptor_76473d52e3e3815d = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2e, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0x4f, 0xc9,
	0xcf, 0x4d, 0xcc, 0xcc, 0x2b, 0xce, 0x4f, 0xce, 0x4e, 0x2d, 0xd1, 0x4f, 0xce, 0xcf, 0x4b, 0xcb,
	0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x82, 0x69, 0x2e, 0x4a, 0xd5, 0x83, 0x6b,
	0xd4, 0x83, 0x69, 0xd4, 0x43, 0xd6, 0xa8, 0x64, 0xc1, 0xc5, 0xe6, 0x0c, 0xd6, 0x2b, 0x24, 0xc4,
	0xc5, 0x52, 0x90, 0x58, 0x92, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x49,
	0x71, 0x71, 0x24, 0x26, 0x15, 0x97, 0x14, 0x25, 0x26, 0x97, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x04, 0xc1, 0xf9, 0x4e, 0xb9, 0x5c, 0x20, 0xe7, 0xe9, 0x11, 0x6f, 0x57, 0x00, 0x63, 0x14, 0x0f,
	0x32, 0x7f, 0x15, 0x93, 0x56, 0x98, 0x51, 0x50, 0x62, 0xa5, 0x9e, 0x33, 0x48, 0x73, 0x08, 0x5c,
	0xb3, 0x27, 0x4c, 0xb3, 0x0b, 0x58, 0x71, 0x30, 0x58, 0x71, 0x12, 0x1b, 0xd8, 0x6f, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xb3, 0xe6, 0x90, 0x1a, 0x01, 0x00, 0x00,
}
