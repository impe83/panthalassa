// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dapp/registry/pb/messages.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_Type int32

const (
	Message_LOG  Message_Type = 0
	Message_DApp Message_Type = 1
)

var Message_Type_name = map[int32]string{
	0: "LOG",
	1: "DApp",
}
var Message_Type_value = map[string]int32{
	"LOG":  0,
	"DApp": 1,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}
func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_messages_fea11ca9f69536ab, []int{0, 0}
}

type Message struct {
	Type                 Message_Type `protobuf:"varint,1,opt,name=type,enum=api.Message_Type" json:"type,omitempty"`
	Log                  []byte       `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty"`
	DApp                 []byte       `protobuf:"bytes,3,opt,name=dApp,proto3" json:"dApp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_fea11ca9f69536ab, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_LOG
}

func (m *Message) GetLog() []byte {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Message) GetDApp() []byte {
	if m != nil {
		return m.DApp
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "api.Message")
	proto.RegisterEnum("api.Message_Type", Message_Type_name, Message_Type_value)
}

func init() {
	proto.RegisterFile("dapp/registry/pb/messages.proto", fileDescriptor_messages_fea11ca9f69536ab)
}

var fileDescriptor_messages_fea11ca9f69536ab = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x49, 0x2c, 0x28,
	0xd0, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2f, 0x48, 0xd2, 0xcf, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8,
	0x54, 0x2a, 0xe4, 0x62, 0xf7, 0x85, 0x08, 0x0b, 0xa9, 0x72, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x09, 0xea, 0x25, 0x16, 0x64, 0xea, 0x41, 0xe5, 0xf4, 0x42,
	0x2a, 0x0b, 0x52, 0x83, 0xc0, 0xd2, 0x42, 0x02, 0x5c, 0xcc, 0x39, 0xf9, 0xe9, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x20, 0xa6, 0x90, 0x10, 0x17, 0x4b, 0x8a, 0x63, 0x41, 0x81, 0x04, 0x33,
	0x58, 0x08, 0xcc, 0x56, 0x92, 0xe4, 0x62, 0x01, 0xe9, 0x11, 0x62, 0xe7, 0x62, 0xf6, 0xf1, 0x77,
	0x17, 0x60, 0x10, 0xe2, 0xe0, 0x62, 0x71, 0x71, 0x2c, 0x28, 0x10, 0x60, 0x4c, 0x62, 0x03, 0x5b,
	0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xee, 0xbd, 0xd2, 0x39, 0xa1, 0x00, 0x00, 0x00,
}
