// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comet.proto

package comet

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

type MsgType int32

const (
	// Heartbeat is sent on downlink and uplink, to keep the persistent connection alive
	MsgType_HB MsgType = 0
	// Auth is sent on uplink as the first message for Comet.Subscribe
	MsgType_AUTH MsgType = 1
	// JoinRoom is sent on uplink to join the specified room
	MsgType_JOIN MsgType = 2
	// ServerPush is sent on downlink to push event to client
	MsgType_PUSH MsgType = 3
)

var MsgType_name = map[int32]string{
	0: "HB",
	1: "AUTH",
	2: "JOIN",
	3: "PUSH",
}

var MsgType_value = map[string]int32{
	"HB":   0,
	"AUTH": 1,
	"JOIN": 2,
	"PUSH": 3,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}

func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{0}
}

type PublishReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Evt                  string   `protobuf:"bytes,2,opt,name=evt,proto3" json:"evt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishReq) Reset()         { *m = PublishReq{} }
func (m *PublishReq) String() string { return proto.CompactTextString(m) }
func (*PublishReq) ProtoMessage()    {}
func (*PublishReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{0}
}

func (m *PublishReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishReq.Unmarshal(m, b)
}
func (m *PublishReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishReq.Marshal(b, m, deterministic)
}
func (m *PublishReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishReq.Merge(m, src)
}
func (m *PublishReq) XXX_Size() int {
	return xxx_messageInfo_PublishReq.Size(m)
}
func (m *PublishReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishReq.DiscardUnknown(m)
}

var xxx_messageInfo_PublishReq proto.InternalMessageInfo

func (m *PublishReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *PublishReq) GetEvt() string {
	if m != nil {
		return m.Evt
	}
	return ""
}

type PublishRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRes) Reset()         { *m = PublishRes{} }
func (m *PublishRes) String() string { return proto.CompactTextString(m) }
func (*PublishRes) ProtoMessage()    {}
func (*PublishRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{1}
}

func (m *PublishRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRes.Unmarshal(m, b)
}
func (m *PublishRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRes.Marshal(b, m, deterministic)
}
func (m *PublishRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRes.Merge(m, src)
}
func (m *PublishRes) XXX_Size() int {
	return xxx_messageInfo_PublishRes.Size(m)
}
func (m *PublishRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRes.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRes proto.InternalMessageInfo

type BroadcastReq struct {
	Rid                  string   `protobuf:"bytes,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Evt                  string   `protobuf:"bytes,2,opt,name=evt,proto3" json:"evt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastReq) Reset()         { *m = BroadcastReq{} }
func (m *BroadcastReq) String() string { return proto.CompactTextString(m) }
func (*BroadcastReq) ProtoMessage()    {}
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{2}
}

func (m *BroadcastReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastReq.Unmarshal(m, b)
}
func (m *BroadcastReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastReq.Marshal(b, m, deterministic)
}
func (m *BroadcastReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastReq.Merge(m, src)
}
func (m *BroadcastReq) XXX_Size() int {
	return xxx_messageInfo_BroadcastReq.Size(m)
}
func (m *BroadcastReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastReq.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastReq proto.InternalMessageInfo

func (m *BroadcastReq) GetRid() string {
	if m != nil {
		return m.Rid
	}
	return ""
}

func (m *BroadcastReq) GetEvt() string {
	if m != nil {
		return m.Evt
	}
	return ""
}

type BroadcastRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastRes) Reset()         { *m = BroadcastRes{} }
func (m *BroadcastRes) String() string { return proto.CompactTextString(m) }
func (*BroadcastRes) ProtoMessage()    {}
func (*BroadcastRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{3}
}

func (m *BroadcastRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastRes.Unmarshal(m, b)
}
func (m *BroadcastRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastRes.Marshal(b, m, deterministic)
}
func (m *BroadcastRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastRes.Merge(m, src)
}
func (m *BroadcastRes) XXX_Size() int {
	return xxx_messageInfo_BroadcastRes.Size(m)
}
func (m *BroadcastRes) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastRes.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastRes proto.InternalMessageInfo

type Uplink struct {
	T                    MsgType    `protobuf:"varint,1,opt,name=t,proto3,enum=comet.MsgType" json:"t,omitempty"`
	Hb                   *Heartbeat `protobuf:"bytes,2,opt,name=hb,proto3" json:"hb,omitempty"`
	Auth                 *Auth      `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Join                 *JoinRoom  `protobuf:"bytes,4,opt,name=join,proto3" json:"join,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Uplink) Reset()         { *m = Uplink{} }
func (m *Uplink) String() string { return proto.CompactTextString(m) }
func (*Uplink) ProtoMessage()    {}
func (*Uplink) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{4}
}

func (m *Uplink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uplink.Unmarshal(m, b)
}
func (m *Uplink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uplink.Marshal(b, m, deterministic)
}
func (m *Uplink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uplink.Merge(m, src)
}
func (m *Uplink) XXX_Size() int {
	return xxx_messageInfo_Uplink.Size(m)
}
func (m *Uplink) XXX_DiscardUnknown() {
	xxx_messageInfo_Uplink.DiscardUnknown(m)
}

var xxx_messageInfo_Uplink proto.InternalMessageInfo

func (m *Uplink) GetT() MsgType {
	if m != nil {
		return m.T
	}
	return MsgType_HB
}

func (m *Uplink) GetHb() *Heartbeat {
	if m != nil {
		return m.Hb
	}
	return nil
}

func (m *Uplink) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Uplink) GetJoin() *JoinRoom {
	if m != nil {
		return m.Join
	}
	return nil
}

type Downlink struct {
	T                    MsgType     `protobuf:"varint,1,opt,name=t,proto3,enum=comet.MsgType" json:"t,omitempty"`
	Hb                   *Heartbeat  `protobuf:"bytes,2,opt,name=hb,proto3" json:"hb,omitempty"`
	Push                 *ServerPush `protobuf:"bytes,3,opt,name=push,proto3" json:"push,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Downlink) Reset()         { *m = Downlink{} }
func (m *Downlink) String() string { return proto.CompactTextString(m) }
func (*Downlink) ProtoMessage()    {}
func (*Downlink) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{5}
}

func (m *Downlink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Downlink.Unmarshal(m, b)
}
func (m *Downlink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Downlink.Marshal(b, m, deterministic)
}
func (m *Downlink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Downlink.Merge(m, src)
}
func (m *Downlink) XXX_Size() int {
	return xxx_messageInfo_Downlink.Size(m)
}
func (m *Downlink) XXX_DiscardUnknown() {
	xxx_messageInfo_Downlink.DiscardUnknown(m)
}

var xxx_messageInfo_Downlink proto.InternalMessageInfo

func (m *Downlink) GetT() MsgType {
	if m != nil {
		return m.T
	}
	return MsgType_HB
}

func (m *Downlink) GetHb() *Heartbeat {
	if m != nil {
		return m.Hb
	}
	return nil
}

func (m *Downlink) GetPush() *ServerPush {
	if m != nil {
		return m.Push
	}
	return nil
}

type Heartbeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{6}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

type Auth struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{7}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type JoinRoom struct {
	Rid                  string   `protobuf:"bytes,1,opt,name=rid,proto3" json:"rid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoom) Reset()         { *m = JoinRoom{} }
func (m *JoinRoom) String() string { return proto.CompactTextString(m) }
func (*JoinRoom) ProtoMessage()    {}
func (*JoinRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{8}
}

func (m *JoinRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoom.Unmarshal(m, b)
}
func (m *JoinRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoom.Marshal(b, m, deterministic)
}
func (m *JoinRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoom.Merge(m, src)
}
func (m *JoinRoom) XXX_Size() int {
	return xxx_messageInfo_JoinRoom.Size(m)
}
func (m *JoinRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoom.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoom proto.InternalMessageInfo

func (m *JoinRoom) GetRid() string {
	if m != nil {
		return m.Rid
	}
	return ""
}

type ServerPush struct {
	Evt                  string   `protobuf:"bytes,1,opt,name=evt,proto3" json:"evt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerPush) Reset()         { *m = ServerPush{} }
func (m *ServerPush) String() string { return proto.CompactTextString(m) }
func (*ServerPush) ProtoMessage()    {}
func (*ServerPush) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac08d4044ac8a660, []int{9}
}

func (m *ServerPush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerPush.Unmarshal(m, b)
}
func (m *ServerPush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerPush.Marshal(b, m, deterministic)
}
func (m *ServerPush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerPush.Merge(m, src)
}
func (m *ServerPush) XXX_Size() int {
	return xxx_messageInfo_ServerPush.Size(m)
}
func (m *ServerPush) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerPush.DiscardUnknown(m)
}

var xxx_messageInfo_ServerPush proto.InternalMessageInfo

func (m *ServerPush) GetEvt() string {
	if m != nil {
		return m.Evt
	}
	return ""
}

func init() {
	proto.RegisterEnum("comet.MsgType", MsgType_name, MsgType_value)
	proto.RegisterType((*PublishReq)(nil), "comet.PublishReq")
	proto.RegisterType((*PublishRes)(nil), "comet.PublishRes")
	proto.RegisterType((*BroadcastReq)(nil), "comet.BroadcastReq")
	proto.RegisterType((*BroadcastRes)(nil), "comet.BroadcastRes")
	proto.RegisterType((*Uplink)(nil), "comet.Uplink")
	proto.RegisterType((*Downlink)(nil), "comet.Downlink")
	proto.RegisterType((*Heartbeat)(nil), "comet.Heartbeat")
	proto.RegisterType((*Auth)(nil), "comet.Auth")
	proto.RegisterType((*JoinRoom)(nil), "comet.JoinRoom")
	proto.RegisterType((*ServerPush)(nil), "comet.ServerPush")
}

func init() {
	proto.RegisterFile("comet.proto", fileDescriptor_ac08d4044ac8a660)
}

var fileDescriptor_ac08d4044ac8a660 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xdd, 0x0e, 0xd2, 0x30,
	0x14, 0xb6, 0x63, 0xfc, 0xec, 0x0c, 0x71, 0x56, 0x2f, 0x16, 0x42, 0x94, 0xd4, 0x98, 0x10, 0x2f,
	0x80, 0xcc, 0xf8, 0x00, 0xa0, 0x17, 0x93, 0x44, 0x25, 0x05, 0x1e, 0x60, 0x83, 0xc6, 0x4d, 0x60,
	0x1d, 0x6d, 0x87, 0xf1, 0x25, 0x7c, 0x06, 0x1f, 0xd5, 0xac, 0xeb, 0x00, 0x03, 0x97, 0xde, 0x9d,
	0x9e, 0xef, 0x3b, 0x5f, 0xbe, 0x7e, 0xe7, 0x80, 0xbb, 0xe5, 0x47, 0xa6, 0xc6, 0xb9, 0xe0, 0x8a,
	0xe3, 0xa6, 0x7e, 0x90, 0x29, 0xc0, 0xb2, 0x88, 0x0f, 0xa9, 0x4c, 0x28, 0x3b, 0x61, 0x0f, 0x1a,
	0x45, 0xba, 0xf3, 0xd1, 0x10, 0x8d, 0x1c, 0x5a, 0x96, 0x65, 0x87, 0x9d, 0x95, 0x6f, 0x55, 0x1d,
	0x76, 0x56, 0xa4, 0x7b, 0x33, 0x21, 0x49, 0x00, 0xdd, 0xb9, 0xe0, 0xd1, 0x6e, 0x1b, 0x49, 0x65,
	0x14, 0xc4, 0x55, 0x41, 0x3c, 0x54, 0xe8, 0xfd, 0x33, 0x23, 0xc9, 0x6f, 0x04, 0xad, 0x4d, 0x7e,
	0x48, 0xb3, 0x3d, 0x1e, 0x00, 0x52, 0x7a, 0xb8, 0x17, 0xf4, 0xc6, 0x95, 0xdd, 0x2f, 0xf2, 0xfb,
	0xfa, 0x57, 0xce, 0x28, 0x52, 0x78, 0x08, 0x56, 0x12, 0x6b, 0x25, 0x37, 0xf0, 0x0c, 0x1c, 0xb2,
	0x48, 0xa8, 0x98, 0x45, 0x8a, 0x5a, 0x49, 0x8c, 0x5f, 0x83, 0x1d, 0x15, 0x2a, 0xf1, 0x1b, 0x9a,
	0xe3, 0x1a, 0xce, 0xac, 0x50, 0x09, 0xd5, 0x00, 0x7e, 0x03, 0xf6, 0x0f, 0x9e, 0x66, 0xbe, 0xad,
	0x09, 0xcf, 0x0c, 0x61, 0xc1, 0xd3, 0x8c, 0x72, 0x7e, 0xa4, 0x1a, 0x24, 0x27, 0xe8, 0x7c, 0xe2,
	0x3f, 0xb3, 0xff, 0xe2, 0xe8, 0x2d, 0xd8, 0x79, 0x21, 0x6b, 0x47, 0xcf, 0x0d, 0x67, 0xc5, 0xc4,
	0x99, 0x89, 0x65, 0x21, 0x13, 0xaa, 0x61, 0xe2, 0x82, 0x73, 0x99, 0x23, 0x03, 0xb0, 0x4b, 0xcb,
	0xf8, 0x25, 0x34, 0x15, 0xdf, 0xb3, 0xcc, 0xc4, 0x59, 0x3d, 0xc8, 0x00, 0x3a, 0xb5, 0xdf, 0xfb,
	0xb8, 0xc9, 0x2b, 0x80, 0xab, 0x78, 0x1d, 0x3e, 0xba, 0x84, 0xff, 0x6e, 0x02, 0x6d, 0xe3, 0x1f,
	0xb7, 0xc0, 0x0a, 0xe7, 0xde, 0x13, 0xdc, 0x01, 0x7b, 0xb6, 0x59, 0x87, 0x1e, 0x2a, 0xab, 0xc5,
	0xb7, 0xcf, 0x5f, 0x3d, 0xab, 0xac, 0x96, 0x9b, 0x55, 0xe8, 0x35, 0x82, 0x3f, 0x08, 0x9a, 0x1f,
	0x4b, 0xd3, 0x78, 0x02, 0x6d, 0xb3, 0x79, 0x5c, 0xff, 0xe3, 0x7a, 0x3b, 0xfd, 0xbb, 0x96, 0xc4,
	0x13, 0x70, 0x56, 0x45, 0x2c, 0xb7, 0x22, 0x8d, 0x19, 0x7e, 0x6a, 0xf0, 0x6a, 0xd3, 0xfd, 0x3a,
	0xfa, 0x3a, 0xe8, 0x11, 0x9a, 0x22, 0xfc, 0x01, 0x9c, 0xcb, 0x65, 0xe0, 0x17, 0x86, 0x71, 0x7b,
	0x5f, 0xfd, 0x07, 0x4d, 0x19, 0xb7, 0xf4, 0x49, 0xbf, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0x13, 0x7a, 0x9a, 0xe1, 0x02, 0x00, 0x00,
}
