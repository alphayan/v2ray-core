package wechat

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

type VideoConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoConfig) Reset()         { *m = VideoConfig{} }
func (m *VideoConfig) String() string { return proto.CompactTextString(m) }
func (*VideoConfig) ProtoMessage()    {}
func (*VideoConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad582a12d5e4846, []int{0}
}

func (m *VideoConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoConfig.Unmarshal(m, b)
}
func (m *VideoConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoConfig.Marshal(b, m, deterministic)
}
func (m *VideoConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoConfig.Merge(m, src)
}
func (m *VideoConfig) XXX_Size() int {
	return xxx_messageInfo_VideoConfig.Size(m)
}
func (m *VideoConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VideoConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VideoConfig)(nil), "v2ray.core.transport.internet.headers.wechat.VideoConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/wechat/config.proto", fileDescriptor_0ad582a12d5e4846)
}

var fileDescriptor_0ad582a12d5e4846 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2d, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0xcf, 0x48,
	0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0x2f, 0x4f, 0x4d, 0xce, 0x48, 0x2c, 0xd1, 0x4f, 0xce, 0xcf,
	0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x81, 0x69, 0x2f, 0x4a, 0xd5,
	0x83, 0x6b, 0xd5, 0x83, 0x69, 0xd5, 0x83, 0x6a, 0xd5, 0x83, 0x68, 0x55, 0xe2, 0xe5, 0xe2, 0x0e,
	0xcb, 0x4c, 0x49, 0xcd, 0x77, 0x06, 0x1b, 0xe1, 0x94, 0xcd, 0x65, 0x90, 0x9c, 0x9f, 0xab, 0x47,
	0x8a, 0x11, 0x01, 0x8c, 0x51, 0x6c, 0x10, 0xd6, 0x2a, 0x26, 0x9d, 0x30, 0xa3, 0xa0, 0xc4, 0x4a,
	0x3d, 0x67, 0x90, 0xc6, 0x10, 0xb8, 0x46, 0x4f, 0x98, 0x46, 0x0f, 0xa8, 0xc6, 0x70, 0xb0, 0xf2,
	0x24, 0x36, 0xb0, 0x83, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x46, 0x75, 0xf8, 0xf1,
	0x00, 0x00, 0x00,
}
