// Code generated by protoc-gen-go. DO NOT EDIT.
// source: liveroom.proto

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

// 消息类型
type RoomMessage_Type int32

const (
	RoomMessage_UNSPECIFIED RoomMessage_Type = 0
	RoomMessage_ENTER_ROOM  RoomMessage_Type = 1
	RoomMessage_LEAVE_ROOM  RoomMessage_Type = 2
	RoomMessage_PLAIN_TEXT  RoomMessage_Type = 3
)

var RoomMessage_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "ENTER_ROOM",
	2: "LEAVE_ROOM",
	3: "PLAIN_TEXT",
}

var RoomMessage_Type_value = map[string]int32{
	"UNSPECIFIED": 0,
	"ENTER_ROOM":  1,
	"LEAVE_ROOM":  2,
	"PLAIN_TEXT":  3,
}

func (x RoomMessage_Type) String() string {
	return proto.EnumName(RoomMessage_Type_name, int32(x))
}

func (RoomMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{8, 0}
}

type EnterReq struct {
	Room                 string        `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Info                 *MsgEnterRoom `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EnterReq) Reset()         { *m = EnterReq{} }
func (m *EnterReq) String() string { return proto.CompactTextString(m) }
func (*EnterReq) ProtoMessage()    {}
func (*EnterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{0}
}

func (m *EnterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterReq.Unmarshal(m, b)
}
func (m *EnterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterReq.Marshal(b, m, deterministic)
}
func (m *EnterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterReq.Merge(m, src)
}
func (m *EnterReq) XXX_Size() int {
	return xxx_messageInfo_EnterReq.Size(m)
}
func (m *EnterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterReq.DiscardUnknown(m)
}

var xxx_messageInfo_EnterReq proto.InternalMessageInfo

func (m *EnterReq) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *EnterReq) GetInfo() *MsgEnterRoom {
	if m != nil {
		return m.Info
	}
	return nil
}

type EnterRes struct {
	Uids                 []string `protobuf:"bytes,1,rep,name=uids,proto3" json:"uids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterRes) Reset()         { *m = EnterRes{} }
func (m *EnterRes) String() string { return proto.CompactTextString(m) }
func (*EnterRes) ProtoMessage()    {}
func (*EnterRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{1}
}

func (m *EnterRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterRes.Unmarshal(m, b)
}
func (m *EnterRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterRes.Marshal(b, m, deterministic)
}
func (m *EnterRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterRes.Merge(m, src)
}
func (m *EnterRes) XXX_Size() int {
	return xxx_messageInfo_EnterRes.Size(m)
}
func (m *EnterRes) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterRes.DiscardUnknown(m)
}

var xxx_messageInfo_EnterRes proto.InternalMessageInfo

func (m *EnterRes) GetUids() []string {
	if m != nil {
		return m.Uids
	}
	return nil
}

type LeaveReq struct {
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveReq) Reset()         { *m = LeaveReq{} }
func (m *LeaveReq) String() string { return proto.CompactTextString(m) }
func (*LeaveReq) ProtoMessage()    {}
func (*LeaveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{2}
}

func (m *LeaveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveReq.Unmarshal(m, b)
}
func (m *LeaveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveReq.Marshal(b, m, deterministic)
}
func (m *LeaveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveReq.Merge(m, src)
}
func (m *LeaveReq) XXX_Size() int {
	return xxx_messageInfo_LeaveReq.Size(m)
}
func (m *LeaveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveReq.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveReq proto.InternalMessageInfo

func (m *LeaveReq) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

type LeaveRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRes) Reset()         { *m = LeaveRes{} }
func (m *LeaveRes) String() string { return proto.CompactTextString(m) }
func (*LeaveRes) ProtoMessage()    {}
func (*LeaveRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{3}
}

func (m *LeaveRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRes.Unmarshal(m, b)
}
func (m *LeaveRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRes.Marshal(b, m, deterministic)
}
func (m *LeaveRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRes.Merge(m, src)
}
func (m *LeaveRes) XXX_Size() int {
	return xxx_messageInfo_LeaveRes.Size(m)
}
func (m *LeaveRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRes.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRes proto.InternalMessageInfo

type SendReq struct {
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendReq) Reset()         { *m = SendReq{} }
func (m *SendReq) String() string { return proto.CompactTextString(m) }
func (*SendReq) ProtoMessage()    {}
func (*SendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{4}
}

func (m *SendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendReq.Unmarshal(m, b)
}
func (m *SendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendReq.Marshal(b, m, deterministic)
}
func (m *SendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendReq.Merge(m, src)
}
func (m *SendReq) XXX_Size() int {
	return xxx_messageInfo_SendReq.Size(m)
}
func (m *SendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendReq proto.InternalMessageInfo

func (m *SendReq) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *SendReq) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SendRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRes) Reset()         { *m = SendRes{} }
func (m *SendRes) String() string { return proto.CompactTextString(m) }
func (*SendRes) ProtoMessage()    {}
func (*SendRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{5}
}

func (m *SendRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRes.Unmarshal(m, b)
}
func (m *SendRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRes.Marshal(b, m, deterministic)
}
func (m *SendRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRes.Merge(m, src)
}
func (m *SendRes) XXX_Size() int {
	return xxx_messageInfo_SendRes.Size(m)
}
func (m *SendRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendRes proto.InternalMessageInfo

type RecvReq struct {
	// 聊天室ID
	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	// 客户端记住的已经收取到的最后一条消息的序列号
	LastSeq uint64 `protobuf:"varint,2,opt,name=last_seq,json=lastSeq,proto3" json:"last_seq,omitempty"`
	// 是否从最新的消息开始接收。这个选项只在客户端首次收取消息，也就是last_seq=0时，有效。
	// 为了防止漏掉消息，客户端下次收取必须从前一次收取到的消息的last_seq开始。
	OffsetNewest         bool     `protobuf:"varint,3,opt,name=offset_newest,json=offsetNewest,proto3" json:"offset_newest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecvReq) Reset()         { *m = RecvReq{} }
func (m *RecvReq) String() string { return proto.CompactTextString(m) }
func (*RecvReq) ProtoMessage()    {}
func (*RecvReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{6}
}

func (m *RecvReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecvReq.Unmarshal(m, b)
}
func (m *RecvReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecvReq.Marshal(b, m, deterministic)
}
func (m *RecvReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvReq.Merge(m, src)
}
func (m *RecvReq) XXX_Size() int {
	return xxx_messageInfo_RecvReq.Size(m)
}
func (m *RecvReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvReq.DiscardUnknown(m)
}

var xxx_messageInfo_RecvReq proto.InternalMessageInfo

func (m *RecvReq) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *RecvReq) GetLastSeq() uint64 {
	if m != nil {
		return m.LastSeq
	}
	return 0
}

func (m *RecvReq) GetOffsetNewest() bool {
	if m != nil {
		return m.OffsetNewest
	}
	return false
}

type RecvRes struct {
	// 本次收取的所有未读消息，按seq排序，最小的seq必须比last_seq大
	Msgs                 []*RoomMessage `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RecvRes) Reset()         { *m = RecvRes{} }
func (m *RecvRes) String() string { return proto.CompactTextString(m) }
func (*RecvRes) ProtoMessage()    {}
func (*RecvRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{7}
}

func (m *RecvRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecvRes.Unmarshal(m, b)
}
func (m *RecvRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecvRes.Marshal(b, m, deterministic)
}
func (m *RecvRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecvRes.Merge(m, src)
}
func (m *RecvRes) XXX_Size() int {
	return xxx_messageInfo_RecvRes.Size(m)
}
func (m *RecvRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RecvRes.DiscardUnknown(m)
}

var xxx_messageInfo_RecvRes proto.InternalMessageInfo

func (m *RecvRes) GetMsgs() []*RoomMessage {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// 聊天室消息
type RoomMessage struct {
	// // 聊天室ID
	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	// 聊天室里的每条消息都有唯一的seq，新消息的seq总是更大
	Seq uint64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	// 消息类型
	Type RoomMessage_Type `protobuf:"varint,3,opt,name=type,proto3,enum=liveroom.RoomMessage_Type" json:"type,omitempty"`
	// 谁发出的消息
	Uid string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	// 何时发出的消息（毫秒时间戳）
	SendAt int64 `protobuf:"varint,5,opt,name=send_at,json=sendAt,proto3" json:"send_at,omitempty"`
	// 具体消息对象，与消息类型对应
	EnterRoom            *MsgEnterRoom `protobuf:"bytes,6,opt,name=enter_room,json=enterRoom,proto3" json:"enter_room,omitempty"`
	LeaveRoom            *MsgLeaveRoom `protobuf:"bytes,7,opt,name=leave_room,json=leaveRoom,proto3" json:"leave_room,omitempty"`
	PlainText            *MsgPlainText `protobuf:"bytes,8,opt,name=plain_text,json=plainText,proto3" json:"plain_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RoomMessage) Reset()         { *m = RoomMessage{} }
func (m *RoomMessage) String() string { return proto.CompactTextString(m) }
func (*RoomMessage) ProtoMessage()    {}
func (*RoomMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{8}
}

func (m *RoomMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomMessage.Unmarshal(m, b)
}
func (m *RoomMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomMessage.Marshal(b, m, deterministic)
}
func (m *RoomMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomMessage.Merge(m, src)
}
func (m *RoomMessage) XXX_Size() int {
	return xxx_messageInfo_RoomMessage.Size(m)
}
func (m *RoomMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RoomMessage proto.InternalMessageInfo

func (m *RoomMessage) GetRoom() string {
	if m != nil {
		return m.Room
	}
	return ""
}

func (m *RoomMessage) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *RoomMessage) GetType() RoomMessage_Type {
	if m != nil {
		return m.Type
	}
	return RoomMessage_UNSPECIFIED
}

func (m *RoomMessage) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RoomMessage) GetSendAt() int64 {
	if m != nil {
		return m.SendAt
	}
	return 0
}

func (m *RoomMessage) GetEnterRoom() *MsgEnterRoom {
	if m != nil {
		return m.EnterRoom
	}
	return nil
}

func (m *RoomMessage) GetLeaveRoom() *MsgLeaveRoom {
	if m != nil {
		return m.LeaveRoom
	}
	return nil
}

func (m *RoomMessage) GetPlainText() *MsgPlainText {
	if m != nil {
		return m.PlainText
	}
	return nil
}

// 进入聊天室
type MsgEnterRoom struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Agent                string   `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Ipaddr               string   `protobuf:"bytes,4,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Location             string   `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgEnterRoom) Reset()         { *m = MsgEnterRoom{} }
func (m *MsgEnterRoom) String() string { return proto.CompactTextString(m) }
func (*MsgEnterRoom) ProtoMessage()    {}
func (*MsgEnterRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{9}
}

func (m *MsgEnterRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgEnterRoom.Unmarshal(m, b)
}
func (m *MsgEnterRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgEnterRoom.Marshal(b, m, deterministic)
}
func (m *MsgEnterRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEnterRoom.Merge(m, src)
}
func (m *MsgEnterRoom) XXX_Size() int {
	return xxx_messageInfo_MsgEnterRoom.Size(m)
}
func (m *MsgEnterRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEnterRoom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEnterRoom proto.InternalMessageInfo

func (m *MsgEnterRoom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgEnterRoom) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *MsgEnterRoom) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *MsgEnterRoom) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

func (m *MsgEnterRoom) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

// 退出聊天室
type MsgLeaveRoom struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgLeaveRoom) Reset()         { *m = MsgLeaveRoom{} }
func (m *MsgLeaveRoom) String() string { return proto.CompactTextString(m) }
func (*MsgLeaveRoom) ProtoMessage()    {}
func (*MsgLeaveRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{10}
}

func (m *MsgLeaveRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgLeaveRoom.Unmarshal(m, b)
}
func (m *MsgLeaveRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgLeaveRoom.Marshal(b, m, deterministic)
}
func (m *MsgLeaveRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLeaveRoom.Merge(m, src)
}
func (m *MsgLeaveRoom) XXX_Size() int {
	return xxx_messageInfo_MsgLeaveRoom.Size(m)
}
func (m *MsgLeaveRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLeaveRoom.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLeaveRoom proto.InternalMessageInfo

// 聊天室文本消息
type MsgPlainText struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgPlainText) Reset()         { *m = MsgPlainText{} }
func (m *MsgPlainText) String() string { return proto.CompactTextString(m) }
func (*MsgPlainText) ProtoMessage()    {}
func (*MsgPlainText) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c9dc16fcaa0b3d, []int{11}
}

func (m *MsgPlainText) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgPlainText.Unmarshal(m, b)
}
func (m *MsgPlainText) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgPlainText.Marshal(b, m, deterministic)
}
func (m *MsgPlainText) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlainText.Merge(m, src)
}
func (m *MsgPlainText) XXX_Size() int {
	return xxx_messageInfo_MsgPlainText.Size(m)
}
func (m *MsgPlainText) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlainText.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlainText proto.InternalMessageInfo

func (m *MsgPlainText) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterEnum("liveroom.RoomMessage_Type", RoomMessage_Type_name, RoomMessage_Type_value)
	proto.RegisterType((*EnterReq)(nil), "liveroom.EnterReq")
	proto.RegisterType((*EnterRes)(nil), "liveroom.EnterRes")
	proto.RegisterType((*LeaveReq)(nil), "liveroom.LeaveReq")
	proto.RegisterType((*LeaveRes)(nil), "liveroom.LeaveRes")
	proto.RegisterType((*SendReq)(nil), "liveroom.SendReq")
	proto.RegisterType((*SendRes)(nil), "liveroom.SendRes")
	proto.RegisterType((*RecvReq)(nil), "liveroom.RecvReq")
	proto.RegisterType((*RecvRes)(nil), "liveroom.RecvRes")
	proto.RegisterType((*RoomMessage)(nil), "liveroom.RoomMessage")
	proto.RegisterType((*MsgEnterRoom)(nil), "liveroom.MsgEnterRoom")
	proto.RegisterType((*MsgLeaveRoom)(nil), "liveroom.MsgLeaveRoom")
	proto.RegisterType((*MsgPlainText)(nil), "liveroom.MsgPlainText")
}

func init() {
	proto.RegisterFile("liveroom.proto", fileDescriptor_b2c9dc16fcaa0b3d)
}

var fileDescriptor_b2c9dc16fcaa0b3d = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x51, 0x8f, 0xd2, 0x40,
	0x10, 0xb6, 0x47, 0x0f, 0xda, 0xe1, 0x44, 0xdc, 0xe8, 0x59, 0xfb, 0x60, 0xc8, 0xfa, 0x82, 0xc6,
	0x60, 0x44, 0xfd, 0x01, 0x17, 0xad, 0xe6, 0x0c, 0x70, 0x64, 0x41, 0xe3, 0x8b, 0x69, 0xd6, 0xeb,
	0x40, 0x9a, 0x94, 0xb6, 0xb0, 0x7b, 0x78, 0xf7, 0xee, 0x4f, 0xf3, 0x67, 0xf9, 0x60, 0x66, 0x4b,
	0x0b, 0x09, 0xe0, 0xdb, 0x7c, 0xdf, 0xee, 0x37, 0xfd, 0x66, 0x66, 0xa7, 0xd0, 0x4a, 0xe2, 0x35,
	0xae, 0xb2, 0x6c, 0xd1, 0xcb, 0x57, 0x99, 0xce, 0x98, 0x53, 0x62, 0xfe, 0x05, 0x9c, 0x20, 0xd5,
	0xb8, 0x12, 0xb8, 0x64, 0x0c, 0x6c, 0xe2, 0x3c, 0xab, 0x63, 0x75, 0x5d, 0x61, 0x62, 0xf6, 0x12,
	0xec, 0x38, 0x9d, 0x65, 0xde, 0x49, 0xc7, 0xea, 0x36, 0xfb, 0xe7, 0xbd, 0x2a, 0xd1, 0x50, 0xcd,
	0x0b, 0x61, 0x96, 0x2d, 0x84, 0xb9, 0xc3, 0x9f, 0x55, 0xb9, 0x14, 0xe5, 0xba, 0x89, 0x23, 0xe5,
	0x59, 0x9d, 0x1a, 0xe5, 0xa2, 0x98, 0xce, 0x07, 0x28, 0xd7, 0x78, 0xe4, 0x5b, 0x1c, 0xaa, 0x73,
	0xc5, 0xdf, 0x40, 0x63, 0x82, 0x69, 0x74, 0xcc, 0x16, 0x03, 0x5b, 0xe3, 0xad, 0x36, 0xb6, 0x5c,
	0x61, 0x62, 0xee, 0x96, 0x12, 0xc5, 0x7f, 0x40, 0x43, 0xe0, 0xf5, 0xfa, 0x98, 0xfa, 0x29, 0x38,
	0x89, 0x54, 0x3a, 0x54, 0xb8, 0x34, 0x19, 0x6c, 0xd1, 0x20, 0x3c, 0xc1, 0x25, 0x7b, 0x0e, 0xf7,
	0xb3, 0xd9, 0x4c, 0xa1, 0x0e, 0x53, 0xfc, 0x85, 0x4a, 0x7b, 0xb5, 0x8e, 0xd5, 0x75, 0xc4, 0x59,
	0x41, 0x8e, 0x0c, 0xc7, 0xdf, 0x95, 0xe9, 0x15, 0x7b, 0x01, 0xf6, 0x42, 0xcd, 0x8b, 0x3a, 0x9b,
	0xfd, 0xc7, 0xdb, 0xfe, 0x50, 0x5f, 0x86, 0xa8, 0x94, 0x9c, 0xa3, 0x30, 0x57, 0xf8, 0xdf, 0x13,
	0x68, 0xee, 0xb0, 0x07, 0x9d, 0xb5, 0xa1, 0xb6, 0x35, 0x45, 0x21, 0xeb, 0x81, 0xad, 0xef, 0x72,
	0x34, 0x3e, 0x5a, 0x7d, 0xff, 0xe0, 0x07, 0x7a, 0xd3, 0xbb, 0x1c, 0x85, 0xb9, 0x47, 0x19, 0x6e,
	0xe2, 0xc8, 0xb3, 0x4d, 0x52, 0x0a, 0xd9, 0x13, 0x68, 0x28, 0x4c, 0xa3, 0x50, 0x6a, 0xef, 0xb4,
	0x63, 0x75, 0x6b, 0xa2, 0x4e, 0xf0, 0x42, 0xb3, 0xf7, 0x00, 0x48, 0xf3, 0x0a, 0x8d, 0x8d, 0xfa,
	0x7f, 0x27, 0xec, 0x62, 0x19, 0x92, 0x2c, 0xa1, 0x31, 0x15, 0xb2, 0xc6, 0x01, 0x59, 0x31, 0x45,
	0x23, 0x4b, 0xca, 0x90, 0x64, 0x79, 0x22, 0xe3, 0x34, 0x34, 0x83, 0x73, 0x0e, 0xc8, 0xc6, 0x74,
	0x3c, 0xc5, 0x5b, 0x2d, 0xdc, 0xbc, 0x0c, 0xf9, 0x67, 0xb0, 0xa9, 0x3a, 0xf6, 0x00, 0x9a, 0x5f,
	0x47, 0x93, 0x71, 0xf0, 0xe1, 0xf2, 0xd3, 0x65, 0xf0, 0xb1, 0x7d, 0x8f, 0xb5, 0x00, 0x82, 0xd1,
	0x34, 0x10, 0xa1, 0xb8, 0xba, 0x1a, 0xb6, 0x2d, 0xc2, 0x83, 0xe0, 0xe2, 0x5b, 0x50, 0xe0, 0x13,
	0xc2, 0xe3, 0xc1, 0xc5, 0xe5, 0x28, 0x9c, 0x06, 0xdf, 0xa7, 0xed, 0x1a, 0xff, 0x6d, 0xc1, 0xd9,
	0x6e, 0x49, 0xd4, 0xff, 0x54, 0x2e, 0xb0, 0xec, 0x3f, 0xc5, 0xec, 0x11, 0x9c, 0xca, 0x39, 0xa6,
	0xe5, 0xc3, 0x2a, 0x00, 0x3b, 0x87, 0xba, 0x5c, 0x4b, 0x2d, 0x57, 0x66, 0x0a, 0xae, 0xd8, 0x20,
	0xe2, 0xe3, 0x5c, 0x46, 0xd1, 0x6a, 0xd3, 0xee, 0x0d, 0x62, 0x3e, 0x38, 0x49, 0x76, 0x2d, 0x75,
	0x9c, 0xa5, 0xa6, 0xe5, 0xae, 0xa8, 0x30, 0x6f, 0x19, 0x17, 0x55, 0x87, 0x38, 0x37, 0xb8, 0x2a,
	0xbd, 0x7a, 0xd9, 0xd6, 0xf6, 0x65, 0xf7, 0xff, 0x58, 0x60, 0x1b, 0xcb, 0xaf, 0xe1, 0xd4, 0xf8,
	0x67, 0x6c, 0xdb, 0xb8, 0x72, 0x7d, 0xfd, 0x7d, 0x4e, 0xb1, 0x57, 0x60, 0xd3, 0x4e, 0xb0, 0x87,
	0xdb, 0xb3, 0xcd, 0x5a, 0xf9, 0x7b, 0x94, 0xb9, 0x4d, 0xef, 0x7a, 0xf7, 0xf6, 0x66, 0x8d, 0xfc,
	0x3d, 0x4a, 0x91, 0x19, 0x53, 0xc6, 0xae, 0x99, 0x72, 0xbf, 0xfd, 0x7d, 0x4e, 0xfd, 0xac, 0x9b,
	0x9f, 0xcf, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x2a, 0xea, 0xbb, 0x8e, 0x04, 0x00,
	0x00,
}
