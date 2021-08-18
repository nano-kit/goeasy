// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package liveroom

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

type HelloReq struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	Sleep                int32    `protobuf:"varint,2,opt,name=sleep,proto3" json:"sleep,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

func (m *HelloReq) GetSleep() int32 {
	if m != nil {
		return m.Sleep
	}
	return 0
}

type HelloRes struct {
	Ack                  string   `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
	Account              []*KV    `protobuf:"bytes,2,rep,name=account,proto3" json:"account,omitempty"`
	Time                 float32  `protobuf:"fixed32,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRes) Reset()         { *m = HelloRes{} }
func (m *HelloRes) String() string { return proto.CompactTextString(m) }
func (*HelloRes) ProtoMessage()    {}
func (*HelloRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *HelloRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRes.Unmarshal(m, b)
}
func (m *HelloRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRes.Marshal(b, m, deterministic)
}
func (m *HelloRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRes.Merge(m, src)
}
func (m *HelloRes) XXX_Size() int {
	return xxx_messageInfo_HelloRes.Size(m)
}
func (m *HelloRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRes.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRes proto.InternalMessageInfo

func (m *HelloRes) GetAck() string {
	if m != nil {
		return m.Ack
	}
	return ""
}

func (m *HelloRes) GetAccount() []*KV {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *HelloRes) GetTime() float32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "liveroom.HelloReq")
	proto.RegisterType((*HelloRes)(nil), "liveroom.HelloRes")
	proto.RegisterType((*KV)(nil), "liveroom.KV")
}

func init() {
	proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d)
}

var fileDescriptor_ca53982754088a9d = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0xcd, 0xcd,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xc9, 0x2c, 0x4b, 0x2d, 0xca, 0xcf, 0xcf,
	0x55, 0x32, 0xe2, 0xe2, 0xf0, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62,
	0x2e, 0x4e, 0xac, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58,
	0x8b, 0x73, 0x52, 0x53, 0x0b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x20, 0x1c, 0xa5, 0x08,
	0xb8, 0x9e, 0x62, 0x90, 0x9e, 0xc4, 0xe4, 0x6c, 0x98, 0x9e, 0xc4, 0xe4, 0x6c, 0x21, 0x35, 0x2e,
	0xf6, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x1e,
	0x3d, 0x98, 0x6d, 0x7a, 0xde, 0x61, 0x41, 0x30, 0x49, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xcc, 0xdc,
	0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xa6, 0x20, 0x30, 0x5b, 0x49, 0x87, 0x8b, 0xc9, 0x3b, 0x0c,
	0x64, 0x66, 0x76, 0x2a, 0xdc, 0x1d, 0xd9, 0xa9, 0x60, 0x77, 0x94, 0x25, 0xe6, 0x94, 0xa6, 0x82,
	0xdd, 0xc1, 0x19, 0x04, 0xe1, 0x18, 0x99, 0x73, 0xb1, 0xb8, 0xa4, 0xe6, 0xe6, 0x0b, 0xe9, 0x73,
	0xb1, 0x82, 0xdd, 0x23, 0x24, 0x84, 0xb0, 0x09, 0xe6, 0x29, 0x29, 0x4c, 0xb1, 0xe2, 0x24, 0x36,
	0x70, 0x28, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x78, 0xf1, 0xb2, 0xdc, 0x13, 0x01, 0x00,
	0x00,
}
