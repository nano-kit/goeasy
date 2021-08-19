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
	Room                 string   `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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
	LastSeq              uint64   `protobuf:"varint,2,opt,name=last_seq,json=lastSeq,proto3" json:"last_seq,omitempty"`
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
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x41, 0xaf, 0xd2, 0x40,
	0x10, 0x76, 0xe9, 0x3e, 0x0a, 0x83, 0x41, 0xdc, 0x44, 0xad, 0x3d, 0x98, 0x66, 0x4f, 0x98, 0x18,
	0x8c, 0xa8, 0x89, 0x57, 0xa2, 0xd5, 0x90, 0x00, 0x8f, 0x2c, 0xd5, 0x78, 0x6b, 0xaa, 0x4c, 0x08,
	0x49, 0xa1, 0x85, 0xdd, 0x47, 0xde, 0xfb, 0x7f, 0xfe, 0x2c, 0x0f, 0x66, 0xb6, 0xb4, 0x25, 0x02,
	0xb7, 0xf9, 0xbe, 0x99, 0x6f, 0x67, 0xa6, 0xf3, 0x15, 0xba, 0xe9, 0xfa, 0x80, 0xfb, 0x2c, 0xdb,
	0x0c, 0xf2, 0x7d, 0x66, 0x32, 0xd1, 0x2a, 0xb1, 0x7c, 0x05, 0xad, 0x70, 0x6b, 0x70, 0xaf, 0x70,
	0x27, 0x04, 0x70, 0xe2, 0x3c, 0x16, 0xb0, 0x7e, 0x5b, 0xf1, 0xff, 0xf2, 0x9a, 0xf2, 0x77, 0xeb,
	0xa5, 0xf6, 0x58, 0xe0, 0x50, 0x9e, 0x62, 0xca, 0x4f, 0x30, 0x39, 0xe0, 0x35, 0x3d, 0x54, 0x79,
	0x2d, 0xdf, 0x81, 0xbb, 0xc0, 0xed, 0xf2, 0x4a, 0x29, 0x71, 0x06, 0xef, 0x8d, 0xd7, 0x28, 0x38,
	0x8a, 0x65, 0xbb, 0x94, 0x68, 0xf9, 0x09, 0x5c, 0x85, 0xbf, 0x0f, 0xd7, 0xd4, 0x2f, 0xa1, 0x95,
	0x26, 0xda, 0xc4, 0x1a, 0x77, 0xf6, 0x05, 0xae, 0x5c, 0xc2, 0x0b, 0xdc, 0xc9, 0x0f, 0xa5, 0x52,
	0x8b, 0xd7, 0xc0, 0x37, 0x7a, 0x55, 0xac, 0xd0, 0x19, 0x3e, 0x1b, 0x54, 0xdf, 0x45, 0x65, 0xd9,
	0x66, 0x8a, 0x5a, 0x27, 0x2b, 0x54, 0xb6, 0x44, 0xfe, 0x6d, 0x40, 0xe7, 0x84, 0xbd, 0xd8, 0xb4,
	0x07, 0x4e, 0xdd, 0x8f, 0x42, 0x31, 0x00, 0x6e, 0x1e, 0x72, 0xf4, 0x9c, 0x80, 0xf5, 0xbb, 0x43,
	0xff, 0x62, 0x83, 0x41, 0xf4, 0x90, 0xa3, 0xb2, 0x75, 0xf4, 0xc2, 0xdd, 0x7a, 0xe9, 0x71, 0xfb,
	0x28, 0x85, 0xe2, 0x05, 0xb8, 0x1a, 0xb7, 0xcb, 0x38, 0x31, 0xde, 0x4d, 0xc0, 0xfa, 0x8e, 0x6a,
	0x12, 0x1c, 0x19, 0xf1, 0x11, 0x00, 0xe9, 0x14, 0xb1, 0x1d, 0xa3, 0x19, 0xb0, 0x7e, 0x67, 0xf8,
	0xbc, 0x6e, 0x30, 0xd5, 0xab, 0xe2, 0x52, 0x59, 0xb6, 0x51, 0x6d, 0x2c, 0x43, 0x92, 0xa5, 0x74,
	0x81, 0x42, 0xe6, 0x5e, 0x90, 0x15, 0x07, 0xb2, 0xb2, 0xb4, 0x0c, 0x49, 0x96, 0xa7, 0xc9, 0x7a,
	0x1b, 0xdb, 0x9b, 0xb4, 0x2e, 0xc8, 0xe6, 0x94, 0x8e, 0xf0, 0xde, 0xa8, 0x76, 0x5e, 0x86, 0xf2,
	0x1b, 0x70, 0xda, 0x4e, 0x3c, 0x81, 0xce, 0xf7, 0xd9, 0x62, 0x1e, 0x7e, 0x1e, 0x7f, 0x1d, 0x87,
	0x5f, 0x7a, 0x8f, 0x44, 0x17, 0x20, 0x9c, 0x45, 0xa1, 0x8a, 0xd5, 0xed, 0xed, 0xb4, 0xc7, 0x08,
	0x4f, 0xc2, 0xd1, 0x8f, 0xb0, 0xc0, 0x0d, 0xc2, 0xf3, 0xc9, 0x68, 0x3c, 0x8b, 0xa3, 0xf0, 0x67,
	0xd4, 0x73, 0x64, 0x17, 0x1e, 0x9f, 0x6e, 0x74, 0xc4, 0xd5, 0xa8, 0x52, 0x5a, 0x5c, 0xcd, 0x50,
	0xb9, 0x87, 0xd5, 0xee, 0x19, 0xfe, 0x61, 0xc0, 0xed, 0x32, 0x6f, 0xe1, 0xc6, 0xbe, 0x24, 0x44,
	0xbd, 0x41, 0x69, 0x7b, 0xff, 0x9c, 0xd3, 0xe2, 0x0d, 0x70, 0xf2, 0x9d, 0x78, 0x5a, 0xe7, 0x8e,
	0xd6, 0xf5, 0xcf, 0x28, 0x5b, 0x4d, 0x06, 0x3b, 0xad, 0x3e, 0x5a, 0xd5, 0x3f, 0xa3, 0x34, 0x0d,
	0x63, 0xd7, 0x38, 0x1d, 0xa6, 0xfc, 0x87, 0xfc, 0x73, 0x4e, 0xff, 0x6a, 0xda, 0x9f, 0xf6, 0xfd,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x4b, 0xa2, 0xc2, 0xc6, 0x03, 0x00, 0x00,
}
