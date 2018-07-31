// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/pb/request.proto

package api_proto

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

type Request struct {
	RequestID               string                           `protobuf:"bytes,1,opt,name=requestID" json:"requestID,omitempty"`
	ShowModal               *Request_ShowModal               `protobuf:"bytes,8,opt,name=showModal" json:"showModal,omitempty"`
	SendEthereumTransaction *Request_SendEthereumTransaction `protobuf:"bytes,9,opt,name=sendEthereumTransaction" json:"sendEthereumTransaction,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                         `json:"-"`
	XXX_unrecognized        []byte                           `json:"-"`
	XXX_sizecache           int32                            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_c94f20042dcff75c, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetRequestID() string {
	if m != nil {
		return m.RequestID
	}
	return ""
}

func (m *Request) GetShowModal() *Request_ShowModal {
	if m != nil {
		return m.ShowModal
	}
	return nil
}

func (m *Request) GetSendEthereumTransaction() *Request_SendEthereumTransaction {
	if m != nil {
		return m.SendEthereumTransaction
	}
	return nil
}

type Request_ShowModal struct {
	DAppPublicKey        []byte   `protobuf:"bytes,1,opt,name=dAppPublicKey,proto3" json:"dAppPublicKey,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Layout               string   `protobuf:"bytes,3,opt,name=layout" json:"layout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_ShowModal) Reset()         { *m = Request_ShowModal{} }
func (m *Request_ShowModal) String() string { return proto.CompactTextString(m) }
func (*Request_ShowModal) ProtoMessage()    {}
func (*Request_ShowModal) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_c94f20042dcff75c, []int{0, 0}
}
func (m *Request_ShowModal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_ShowModal.Unmarshal(m, b)
}
func (m *Request_ShowModal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_ShowModal.Marshal(b, m, deterministic)
}
func (dst *Request_ShowModal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_ShowModal.Merge(dst, src)
}
func (m *Request_ShowModal) XXX_Size() int {
	return xxx_messageInfo_Request_ShowModal.Size(m)
}
func (m *Request_ShowModal) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_ShowModal.DiscardUnknown(m)
}

var xxx_messageInfo_Request_ShowModal proto.InternalMessageInfo

func (m *Request_ShowModal) GetDAppPublicKey() []byte {
	if m != nil {
		return m.DAppPublicKey
	}
	return nil
}

func (m *Request_ShowModal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Request_ShowModal) GetLayout() string {
	if m != nil {
		return m.Layout
	}
	return ""
}

type Request_SendEthereumTransaction struct {
	Value                string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_SendEthereumTransaction) Reset()         { *m = Request_SendEthereumTransaction{} }
func (m *Request_SendEthereumTransaction) String() string { return proto.CompactTextString(m) }
func (*Request_SendEthereumTransaction) ProtoMessage()    {}
func (*Request_SendEthereumTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_c94f20042dcff75c, []int{0, 1}
}
func (m *Request_SendEthereumTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_SendEthereumTransaction.Unmarshal(m, b)
}
func (m *Request_SendEthereumTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_SendEthereumTransaction.Marshal(b, m, deterministic)
}
func (dst *Request_SendEthereumTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_SendEthereumTransaction.Merge(dst, src)
}
func (m *Request_SendEthereumTransaction) XXX_Size() int {
	return xxx_messageInfo_Request_SendEthereumTransaction.Size(m)
}
func (m *Request_SendEthereumTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_SendEthereumTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_Request_SendEthereumTransaction proto.InternalMessageInfo

func (m *Request_SendEthereumTransaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Request_SendEthereumTransaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *Request_SendEthereumTransaction) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "api_proto.Request")
	proto.RegisterType((*Request_ShowModal)(nil), "api_proto.Request.ShowModal")
	proto.RegisterType((*Request_SendEthereumTransaction)(nil), "api_proto.Request.SendEthereumTransaction")
}

func init() { proto.RegisterFile("api/pb/request.proto", fileDescriptor_request_c94f20042dcff75c) }

var fileDescriptor_request_c94f20042dcff75c = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x55, 0xab, 0x19, 0xff, 0x1c, 0x86, 0x62, 0x97, 0xd2, 0x43, 0x11, 0x0f, 0xc5,
	0xc3, 0x16, 0xf4, 0xe6, 0x4d, 0xd0, 0x83, 0x88, 0x20, 0xa9, 0xf7, 0x32, 0xdb, 0x04, 0x1a, 0x88,
	0x9b, 0xb8, 0x3b, 0x51, 0xfa, 0xe9, 0x15, 0xb3, 0xb1, 0xa5, 0xd0, 0xde, 0xe6, 0xbd, 0xbc, 0xfc,
	0x92, 0x37, 0xd0, 0x27, 0x6f, 0xa6, 0xbe, 0x9a, 0x36, 0xfa, 0x33, 0xe8, 0x96, 0x4b, 0xdf, 0x38,
	0x76, 0x28, 0xc8, 0x9b, 0x79, 0x1c, 0xaf, 0x7e, 0x72, 0x38, 0x96, 0xdd, 0x21, 0x8e, 0x40, 0xa4,
	0xdc, 0xf3, 0x63, 0x91, 0x8d, 0xb3, 0x89, 0x90, 0x1b, 0x03, 0xef, 0x41, 0xb4, 0x4b, 0xf7, 0xfd,
	0xea, 0x14, 0xd9, 0xe2, 0x64, 0x9c, 0x4d, 0x4e, 0x6f, 0x47, 0xe5, 0x1a, 0x54, 0x26, 0x48, 0x39,
	0xfb, 0xcf, 0xc8, 0x4d, 0x1c, 0x15, 0x0c, 0x5a, 0x5d, 0xab, 0x27, 0x5e, 0xea, 0x46, 0x87, 0x8f,
	0xf7, 0x86, 0xea, 0x96, 0x16, 0x6c, 0x5c, 0x5d, 0x88, 0x48, 0xba, 0xd9, 0x45, 0xda, 0x7d, 0x43,
	0xee, 0x43, 0x0d, 0xe7, 0x20, 0xd6, 0xaf, 0xe3, 0x35, 0x9c, 0xab, 0x07, 0xef, 0xdf, 0x42, 0x65,
	0xcd, 0xe2, 0x45, 0xaf, 0x62, 0xa1, 0x33, 0xb9, 0x6d, 0x62, 0x1f, 0x8e, 0xd8, 0xb0, 0xd5, 0x45,
	0x1e, 0xeb, 0x76, 0x02, 0x2f, 0xa1, 0x67, 0x69, 0xe5, 0x02, 0x17, 0x07, 0xd1, 0x4e, 0x6a, 0x38,
	0x83, 0xc1, 0x9e, 0x4f, 0xfd, 0x81, 0xbe, 0xc8, 0x06, 0x9d, 0xf6, 0xd6, 0x09, 0xbc, 0x80, 0x9c,
	0x5d, 0x62, 0xe7, 0xec, 0x10, 0xe1, 0x50, 0x11, 0x53, 0xc2, 0xc6, 0xb9, 0xea, 0xc5, 0xd6, 0x77,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59, 0x09, 0x03, 0x6e, 0xab, 0x01, 0x00, 0x00,
}
