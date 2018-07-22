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
	DRKeyStoreGet           *Request_DRKeyStoreGet           `protobuf:"bytes,2,opt,name=dRKeyStoreGet" json:"dRKeyStoreGet,omitempty"`
	DRKeyStorePut           *Request_DRKeyStorePut           `protobuf:"bytes,3,opt,name=dRKeyStorePut" json:"dRKeyStorePut,omitempty"`
	DRKeyStoreDeleteMK      *Request_DRKeyStoreDeleteMK      `protobuf:"bytes,4,opt,name=dRKeyStoreDeleteMK" json:"dRKeyStoreDeleteMK,omitempty"`
	DRKeyStoreDeleteKeys    *Request_DRKeyStoreDeleteKeys    `protobuf:"bytes,5,opt,name=dRKeyStoreDeleteKeys" json:"dRKeyStoreDeleteKeys,omitempty"`
	DRKeyStoreCount         *Request_DRKeyStoreCount         `protobuf:"bytes,6,opt,name=dRKeyStoreCount" json:"dRKeyStoreCount,omitempty"`
	DRKeyStoreAll           *Request_DRKeyStoreAll           `protobuf:"bytes,7,opt,name=dRKeyStoreAll" json:"dRKeyStoreAll,omitempty"`
	ShowModal               *Request_RenderModal             `protobuf:"bytes,8,opt,name=showModal" json:"showModal,omitempty"`
	SendEthereumTransaction *Request_SendEthereumTransaction `protobuf:"bytes,9,opt,name=sendEthereumTransaction" json:"sendEthereumTransaction,omitempty"`
	SaveDApp                *Request_SaveDApp                `protobuf:"bytes,10,opt,name=saveDApp" json:"saveDApp,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                         `json:"-"`
	XXX_unrecognized        []byte                           `json:"-"`
	XXX_sizecache           int32                            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0}
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

func (m *Request) GetDRKeyStoreGet() *Request_DRKeyStoreGet {
	if m != nil {
		return m.DRKeyStoreGet
	}
	return nil
}

func (m *Request) GetDRKeyStorePut() *Request_DRKeyStorePut {
	if m != nil {
		return m.DRKeyStorePut
	}
	return nil
}

func (m *Request) GetDRKeyStoreDeleteMK() *Request_DRKeyStoreDeleteMK {
	if m != nil {
		return m.DRKeyStoreDeleteMK
	}
	return nil
}

func (m *Request) GetDRKeyStoreDeleteKeys() *Request_DRKeyStoreDeleteKeys {
	if m != nil {
		return m.DRKeyStoreDeleteKeys
	}
	return nil
}

func (m *Request) GetDRKeyStoreCount() *Request_DRKeyStoreCount {
	if m != nil {
		return m.DRKeyStoreCount
	}
	return nil
}

func (m *Request) GetDRKeyStoreAll() *Request_DRKeyStoreAll {
	if m != nil {
		return m.DRKeyStoreAll
	}
	return nil
}

func (m *Request) GetShowModal() *Request_RenderModal {
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

func (m *Request) GetSaveDApp() *Request_SaveDApp {
	if m != nil {
		return m.SaveDApp
	}
	return nil
}

type Request_DRKeyStoreGet struct {
	DrKey                []byte   `protobuf:"bytes,1,opt,name=drKey,proto3" json:"drKey,omitempty"`
	MessageNumber        uint64   `protobuf:"varint,2,opt,name=messageNumber" json:"messageNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_DRKeyStoreGet) Reset()         { *m = Request_DRKeyStoreGet{} }
func (m *Request_DRKeyStoreGet) String() string { return proto.CompactTextString(m) }
func (*Request_DRKeyStoreGet) ProtoMessage()    {}
func (*Request_DRKeyStoreGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 0}
}
func (m *Request_DRKeyStoreGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_DRKeyStoreGet.Unmarshal(m, b)
}
func (m *Request_DRKeyStoreGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_DRKeyStoreGet.Marshal(b, m, deterministic)
}
func (dst *Request_DRKeyStoreGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_DRKeyStoreGet.Merge(dst, src)
}
func (m *Request_DRKeyStoreGet) XXX_Size() int {
	return xxx_messageInfo_Request_DRKeyStoreGet.Size(m)
}
func (m *Request_DRKeyStoreGet) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_DRKeyStoreGet.DiscardUnknown(m)
}

var xxx_messageInfo_Request_DRKeyStoreGet proto.InternalMessageInfo

func (m *Request_DRKeyStoreGet) GetDrKey() []byte {
	if m != nil {
		return m.DrKey
	}
	return nil
}

func (m *Request_DRKeyStoreGet) GetMessageNumber() uint64 {
	if m != nil {
		return m.MessageNumber
	}
	return 0
}

type Request_DRKeyStorePut struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MessageNumber        uint64   `protobuf:"varint,2,opt,name=messageNumber" json:"messageNumber,omitempty"`
	MessageKey           []byte   `protobuf:"bytes,3,opt,name=messageKey,proto3" json:"messageKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_DRKeyStorePut) Reset()         { *m = Request_DRKeyStorePut{} }
func (m *Request_DRKeyStorePut) String() string { return proto.CompactTextString(m) }
func (*Request_DRKeyStorePut) ProtoMessage()    {}
func (*Request_DRKeyStorePut) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 1}
}
func (m *Request_DRKeyStorePut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_DRKeyStorePut.Unmarshal(m, b)
}
func (m *Request_DRKeyStorePut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_DRKeyStorePut.Marshal(b, m, deterministic)
}
func (dst *Request_DRKeyStorePut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_DRKeyStorePut.Merge(dst, src)
}
func (m *Request_DRKeyStorePut) XXX_Size() int {
	return xxx_messageInfo_Request_DRKeyStorePut.Size(m)
}
func (m *Request_DRKeyStorePut) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_DRKeyStorePut.DiscardUnknown(m)
}

var xxx_messageInfo_Request_DRKeyStorePut proto.InternalMessageInfo

func (m *Request_DRKeyStorePut) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request_DRKeyStorePut) GetMessageNumber() uint64 {
	if m != nil {
		return m.MessageNumber
	}
	return 0
}

func (m *Request_DRKeyStorePut) GetMessageKey() []byte {
	if m != nil {
		return m.MessageKey
	}
	return nil
}

type Request_DRKeyStoreDeleteMK struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MsgNum               uint64   `protobuf:"varint,2,opt,name=msgNum" json:"msgNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_DRKeyStoreDeleteMK) Reset()         { *m = Request_DRKeyStoreDeleteMK{} }
func (m *Request_DRKeyStoreDeleteMK) String() string { return proto.CompactTextString(m) }
func (*Request_DRKeyStoreDeleteMK) ProtoMessage()    {}
func (*Request_DRKeyStoreDeleteMK) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 2}
}
func (m *Request_DRKeyStoreDeleteMK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_DRKeyStoreDeleteMK.Unmarshal(m, b)
}
func (m *Request_DRKeyStoreDeleteMK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_DRKeyStoreDeleteMK.Marshal(b, m, deterministic)
}
func (dst *Request_DRKeyStoreDeleteMK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_DRKeyStoreDeleteMK.Merge(dst, src)
}
func (m *Request_DRKeyStoreDeleteMK) XXX_Size() int {
	return xxx_messageInfo_Request_DRKeyStoreDeleteMK.Size(m)
}
func (m *Request_DRKeyStoreDeleteMK) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_DRKeyStoreDeleteMK.DiscardUnknown(m)
}

var xxx_messageInfo_Request_DRKeyStoreDeleteMK proto.InternalMessageInfo

func (m *Request_DRKeyStoreDeleteMK) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request_DRKeyStoreDeleteMK) GetMsgNum() uint64 {
	if m != nil {
		return m.MsgNum
	}
	return 0
}

type Request_DRKeyStoreDeleteKeys struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_DRKeyStoreDeleteKeys) Reset()         { *m = Request_DRKeyStoreDeleteKeys{} }
func (m *Request_DRKeyStoreDeleteKeys) String() string { return proto.CompactTextString(m) }
func (*Request_DRKeyStoreDeleteKeys) ProtoMessage()    {}
func (*Request_DRKeyStoreDeleteKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 3}
}
func (m *Request_DRKeyStoreDeleteKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_DRKeyStoreDeleteKeys.Unmarshal(m, b)
}
func (m *Request_DRKeyStoreDeleteKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_DRKeyStoreDeleteKeys.Marshal(b, m, deterministic)
}
func (dst *Request_DRKeyStoreDeleteKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_DRKeyStoreDeleteKeys.Merge(dst, src)
}
func (m *Request_DRKeyStoreDeleteKeys) XXX_Size() int {
	return xxx_messageInfo_Request_DRKeyStoreDeleteKeys.Size(m)
}
func (m *Request_DRKeyStoreDeleteKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_DRKeyStoreDeleteKeys.DiscardUnknown(m)
}

var xxx_messageInfo_Request_DRKeyStoreDeleteKeys proto.InternalMessageInfo

func (m *Request_DRKeyStoreDeleteKeys) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Request_DRKeyStoreCount struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_DRKeyStoreCount) Reset()         { *m = Request_DRKeyStoreCount{} }
func (m *Request_DRKeyStoreCount) String() string { return proto.CompactTextString(m) }
func (*Request_DRKeyStoreCount) ProtoMessage()    {}
func (*Request_DRKeyStoreCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 4}
}
func (m *Request_DRKeyStoreCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_DRKeyStoreCount.Unmarshal(m, b)
}
func (m *Request_DRKeyStoreCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_DRKeyStoreCount.Marshal(b, m, deterministic)
}
func (dst *Request_DRKeyStoreCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_DRKeyStoreCount.Merge(dst, src)
}
func (m *Request_DRKeyStoreCount) XXX_Size() int {
	return xxx_messageInfo_Request_DRKeyStoreCount.Size(m)
}
func (m *Request_DRKeyStoreCount) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_DRKeyStoreCount.DiscardUnknown(m)
}

var xxx_messageInfo_Request_DRKeyStoreCount proto.InternalMessageInfo

func (m *Request_DRKeyStoreCount) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Request_DRKeyStoreAll struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_DRKeyStoreAll) Reset()         { *m = Request_DRKeyStoreAll{} }
func (m *Request_DRKeyStoreAll) String() string { return proto.CompactTextString(m) }
func (*Request_DRKeyStoreAll) ProtoMessage()    {}
func (*Request_DRKeyStoreAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 5}
}
func (m *Request_DRKeyStoreAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_DRKeyStoreAll.Unmarshal(m, b)
}
func (m *Request_DRKeyStoreAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_DRKeyStoreAll.Marshal(b, m, deterministic)
}
func (dst *Request_DRKeyStoreAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_DRKeyStoreAll.Merge(dst, src)
}
func (m *Request_DRKeyStoreAll) XXX_Size() int {
	return xxx_messageInfo_Request_DRKeyStoreAll.Size(m)
}
func (m *Request_DRKeyStoreAll) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_DRKeyStoreAll.DiscardUnknown(m)
}

var xxx_messageInfo_Request_DRKeyStoreAll proto.InternalMessageInfo

type Request_RenderModal struct {
	DAppPublicKey        []byte   `protobuf:"bytes,1,opt,name=dAppPublicKey,proto3" json:"dAppPublicKey,omitempty"`
	UiID                 string   `protobuf:"bytes,2,opt,name=uiID" json:"uiID,omitempty"`
	Layout               string   `protobuf:"bytes,3,opt,name=layout" json:"layout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_RenderModal) Reset()         { *m = Request_RenderModal{} }
func (m *Request_RenderModal) String() string { return proto.CompactTextString(m) }
func (*Request_RenderModal) ProtoMessage()    {}
func (*Request_RenderModal) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 6}
}
func (m *Request_RenderModal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_RenderModal.Unmarshal(m, b)
}
func (m *Request_RenderModal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_RenderModal.Marshal(b, m, deterministic)
}
func (dst *Request_RenderModal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_RenderModal.Merge(dst, src)
}
func (m *Request_RenderModal) XXX_Size() int {
	return xxx_messageInfo_Request_RenderModal.Size(m)
}
func (m *Request_RenderModal) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_RenderModal.DiscardUnknown(m)
}

var xxx_messageInfo_Request_RenderModal proto.InternalMessageInfo

func (m *Request_RenderModal) GetDAppPublicKey() []byte {
	if m != nil {
		return m.DAppPublicKey
	}
	return nil
}

func (m *Request_RenderModal) GetUiID() string {
	if m != nil {
		return m.UiID
	}
	return ""
}

func (m *Request_RenderModal) GetLayout() string {
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
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 7}
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

type Request_SaveDApp struct {
	AppName              string   `protobuf:"bytes,1,opt,name=appName" json:"appName,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Signature            string   `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
	SigningPublicKey     string   `protobuf:"bytes,4,opt,name=signingPublicKey" json:"signingPublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_SaveDApp) Reset()         { *m = Request_SaveDApp{} }
func (m *Request_SaveDApp) String() string { return proto.CompactTextString(m) }
func (*Request_SaveDApp) ProtoMessage()    {}
func (*Request_SaveDApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_60f52d7d44df6391, []int{0, 8}
}
func (m *Request_SaveDApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_SaveDApp.Unmarshal(m, b)
}
func (m *Request_SaveDApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_SaveDApp.Marshal(b, m, deterministic)
}
func (dst *Request_SaveDApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_SaveDApp.Merge(dst, src)
}
func (m *Request_SaveDApp) XXX_Size() int {
	return xxx_messageInfo_Request_SaveDApp.Size(m)
}
func (m *Request_SaveDApp) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_SaveDApp.DiscardUnknown(m)
}

var xxx_messageInfo_Request_SaveDApp proto.InternalMessageInfo

func (m *Request_SaveDApp) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *Request_SaveDApp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Request_SaveDApp) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Request_SaveDApp) GetSigningPublicKey() string {
	if m != nil {
		return m.SigningPublicKey
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "api_proto.Request")
	proto.RegisterType((*Request_DRKeyStoreGet)(nil), "api_proto.Request.DRKeyStoreGet")
	proto.RegisterType((*Request_DRKeyStorePut)(nil), "api_proto.Request.DRKeyStorePut")
	proto.RegisterType((*Request_DRKeyStoreDeleteMK)(nil), "api_proto.Request.DRKeyStoreDeleteMK")
	proto.RegisterType((*Request_DRKeyStoreDeleteKeys)(nil), "api_proto.Request.DRKeyStoreDeleteKeys")
	proto.RegisterType((*Request_DRKeyStoreCount)(nil), "api_proto.Request.DRKeyStoreCount")
	proto.RegisterType((*Request_DRKeyStoreAll)(nil), "api_proto.Request.DRKeyStoreAll")
	proto.RegisterType((*Request_RenderModal)(nil), "api_proto.Request.RenderModal")
	proto.RegisterType((*Request_SendEthereumTransaction)(nil), "api_proto.Request.SendEthereumTransaction")
	proto.RegisterType((*Request_SaveDApp)(nil), "api_proto.Request.SaveDApp")
}

func init() { proto.RegisterFile("api/pb/request.proto", fileDescriptor_request_60f52d7d44df6391) }

var fileDescriptor_request_60f52d7d44df6391 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x5f, 0x6f, 0x12, 0x41,
	0x14, 0xc5, 0x43, 0xa1, 0xc0, 0xde, 0x8a, 0x34, 0x13, 0x62, 0x27, 0xab, 0x69, 0x48, 0xd5, 0x48,
	0xfa, 0x40, 0x13, 0x7d, 0xf0, 0xc5, 0x98, 0x10, 0x51, 0xd3, 0x60, 0x09, 0x19, 0xf4, 0xc9, 0x87,
	0x66, 0x60, 0x6f, 0xe8, 0xc6, 0xd9, 0x9d, 0x75, 0x77, 0xa6, 0x86, 0x57, 0xbf, 0xb0, 0x5f, 0xc1,
	0xcc, 0x30, 0xb0, 0x2c, 0x6c, 0x1b, 0xde, 0xee, 0x3d, 0x7b, 0xee, 0x6f, 0xfe, 0xec, 0xc9, 0x40,
	0x87, 0x27, 0xe1, 0x55, 0x32, 0xbb, 0x4a, 0xf1, 0xb7, 0xc6, 0x4c, 0xf5, 0x93, 0x54, 0x2a, 0x49,
	0x3c, 0x9e, 0x84, 0xb7, 0xb6, 0xbc, 0xf8, 0x07, 0xd0, 0x60, 0xab, 0x8f, 0xe4, 0x05, 0x78, 0xce,
	0x77, 0x3d, 0xa4, 0x95, 0x6e, 0xa5, 0xe7, 0xb1, 0x5c, 0x20, 0x5f, 0xa0, 0x15, 0xb0, 0x11, 0x2e,
	0xa7, 0x4a, 0xa6, 0xf8, 0x15, 0x15, 0x3d, 0xea, 0x56, 0x7a, 0x27, 0x6f, 0xbb, 0xfd, 0x0d, 0xac,
	0xef, 0x40, 0xfd, 0xe1, 0xb6, 0x8f, 0x15, 0xc7, 0x8a, 0x9c, 0x89, 0x56, 0xb4, 0x7a, 0x00, 0x67,
	0xa2, 0x0b, 0x9c, 0x89, 0x56, 0xe4, 0x07, 0x90, 0x5c, 0x18, 0xa2, 0x40, 0x85, 0x37, 0x23, 0x5a,
	0xb3, 0xb0, 0xd7, 0x8f, 0xc2, 0xd6, 0x66, 0x56, 0x02, 0x20, 0x3f, 0xa1, 0xb3, 0xab, 0x8e, 0x70,
	0x99, 0xd1, 0x63, 0x0b, 0x7e, 0x73, 0x00, 0xd8, 0xd8, 0x59, 0x29, 0x84, 0x7c, 0x83, 0x76, 0xae,
	0x7f, 0x92, 0x3a, 0x56, 0xb4, 0x6e, 0xb9, 0x17, 0x8f, 0x72, 0xad, 0x93, 0xed, 0x8e, 0x16, 0x6f,
	0x72, 0x20, 0x04, 0x6d, 0x1c, 0x70, 0x93, 0x03, 0x21, 0x58, 0x71, 0x8c, 0x7c, 0x00, 0x2f, 0xbb,
	0x93, 0x7f, 0x6e, 0x64, 0xc0, 0x05, 0x6d, 0x5a, 0xc6, 0x79, 0x09, 0x83, 0x61, 0x1c, 0x60, 0x6a,
	0x5d, 0x2c, 0x1f, 0x20, 0x01, 0x9c, 0x65, 0x18, 0x07, 0x9f, 0xd5, 0x1d, 0xa6, 0xa8, 0xa3, 0xef,
	0x29, 0x8f, 0x33, 0x3e, 0x57, 0xa1, 0x8c, 0xa9, 0x67, 0x59, 0x97, 0x25, 0xac, 0x69, 0xf9, 0x04,
	0x7b, 0x08, 0x45, 0xde, 0x43, 0x33, 0xe3, 0xf7, 0x38, 0x1c, 0x24, 0x09, 0x05, 0x8b, 0x7d, 0x5e,
	0x86, 0x75, 0x16, 0xb6, 0x31, 0xfb, 0x23, 0x68, 0x15, 0xe2, 0x48, 0x3a, 0x70, 0x1c, 0xa4, 0x23,
	0x5c, 0xda, 0x84, 0x3f, 0x61, 0xab, 0x86, 0xbc, 0x82, 0x56, 0x84, 0x59, 0xc6, 0x17, 0x38, 0xd6,
	0xd1, 0x0c, 0x53, 0x9b, 0xee, 0x1a, 0x2b, 0x8a, 0xfe, 0x62, 0x1b, 0x66, 0x42, 0x78, 0x0a, 0xd5,
	0x5f, 0x1b, 0x94, 0x29, 0x0f, 0x03, 0x91, 0x73, 0x00, 0x27, 0x98, 0x9d, 0x54, 0xed, 0xf8, 0x96,
	0xe2, 0x7f, 0x04, 0xb2, 0x9f, 0xd7, 0x92, 0xd5, 0x9e, 0x41, 0x3d, 0xca, 0x16, 0x63, 0x1d, 0xb9,
	0x65, 0x5c, 0xe7, 0xf7, 0xa0, 0x53, 0x16, 0xcb, 0x7d, 0x82, 0xff, 0x12, 0xda, 0x3b, 0x41, 0x2b,
	0x31, 0xb5, 0xb7, 0xcf, 0x3d, 0x10, 0xc2, 0xbf, 0x85, 0x93, 0xad, 0x38, 0x98, 0x43, 0x07, 0x83,
	0x24, 0x99, 0xe8, 0x99, 0x08, 0xe7, 0xf9, 0xdd, 0x16, 0x45, 0x42, 0xa0, 0xa6, 0xc3, 0xeb, 0xa1,
	0xdd, 0xaa, 0xc7, 0x6c, 0x6d, 0x0e, 0x20, 0xf8, 0x52, 0xba, 0x67, 0xc0, 0x63, 0xae, 0xf3, 0xa7,
	0x70, 0xf6, 0x40, 0x46, 0xcc, 0x0f, 0xbc, 0xe7, 0x42, 0xa3, 0x7b, 0xa2, 0x56, 0x0d, 0x79, 0x0a,
	0x47, 0x4a, 0x3a, 0xf4, 0x91, 0x92, 0x66, 0xb1, 0x80, 0x2b, 0xee, 0xb0, 0xb6, 0xf6, 0xff, 0x56,
	0xa0, 0xb9, 0x8e, 0x08, 0xa1, 0xd0, 0xe0, 0x49, 0x32, 0xe6, 0xd1, 0x1a, 0xb4, 0x6e, 0xcd, 0xe8,
	0x5c, 0x06, 0xb8, 0xde, 0xa7, 0xa9, 0xcd, 0xdb, 0x98, 0x85, 0x8b, 0x98, 0x2b, 0x9d, 0xa2, 0x63,
	0xe6, 0x02, 0xb9, 0x84, 0x53, 0xd3, 0x84, 0xf1, 0x22, 0xbf, 0x82, 0x9a, 0x35, 0xed, 0xe9, 0xb3,
	0xba, 0x8d, 0xec, 0xbb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x95, 0x00, 0x5b, 0x9b, 0x05,
	0x00, 0x00,
}
