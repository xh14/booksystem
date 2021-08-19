// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package proto

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

type MessageHead struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageHead) Reset()         { *m = MessageHead{} }
func (m *MessageHead) String() string { return proto.CompactTextString(m) }
func (*MessageHead) ProtoMessage()    {}
func (*MessageHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *MessageHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageHead.Unmarshal(m, b)
}
func (m *MessageHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageHead.Marshal(b, m, deterministic)
}
func (m *MessageHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageHead.Merge(m, src)
}
func (m *MessageHead) XXX_Size() int {
	return xxx_messageInfo_MessageHead.Size(m)
}
func (m *MessageHead) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageHead.DiscardUnknown(m)
}

var xxx_messageInfo_MessageHead proto.InternalMessageInfo

func (m *MessageHead) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MessageHead) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type MessageResponse struct {
	Head                 *MessageHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetHead() *MessageHead {
	if m != nil {
		return m.Head
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageHead)(nil), "proto.MessageHead")
	proto.RegisterType((*MessageResponse)(nil), "proto.MessageResponse")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xa6, 0x5c, 0xdc,
	0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x1e, 0xa9, 0x89, 0x29, 0x42, 0x7c, 0x5c, 0x4c, 0x9e,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x4c, 0x9e, 0x29, 0x42, 0x62, 0x5c, 0x6c, 0xae,
	0x45, 0x45, 0xbe, 0xc5, 0xe9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x25,
	0x17, 0x3f, 0x54, 0x5b, 0x50, 0x6a, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x1a, 0x17, 0x0b,
	0xc8, 0x08, 0xb0, 0x66, 0x6e, 0x23, 0x21, 0x88, 0x35, 0x7a, 0x48, 0x86, 0x07, 0x81, 0xe5, 0x93,
	0xd8, 0xc0, 0x12, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xf6, 0xdc, 0xb9, 0x8f, 0x00,
	0x00, 0x00,
}
