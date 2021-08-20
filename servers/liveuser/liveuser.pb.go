// Code generated by protoc-gen-go. DO NOT EDIT.
// source: liveuser.proto

package liveuser

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

type AddUserReq struct {
	User                 *UserRecord `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddUserReq) Reset()         { *m = AddUserReq{} }
func (m *AddUserReq) String() string { return proto.CompactTextString(m) }
func (*AddUserReq) ProtoMessage()    {}
func (*AddUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{0}
}

func (m *AddUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserReq.Unmarshal(m, b)
}
func (m *AddUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserReq.Marshal(b, m, deterministic)
}
func (m *AddUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserReq.Merge(m, src)
}
func (m *AddUserReq) XXX_Size() int {
	return xxx_messageInfo_AddUserReq.Size(m)
}
func (m *AddUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserReq proto.InternalMessageInfo

func (m *AddUserReq) GetUser() *UserRecord {
	if m != nil {
		return m.User
	}
	return nil
}

type AddUserRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRes) Reset()         { *m = AddUserRes{} }
func (m *AddUserRes) String() string { return proto.CompactTextString(m) }
func (*AddUserRes) ProtoMessage()    {}
func (*AddUserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{1}
}

func (m *AddUserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRes.Unmarshal(m, b)
}
func (m *AddUserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRes.Marshal(b, m, deterministic)
}
func (m *AddUserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRes.Merge(m, src)
}
func (m *AddUserRes) XXX_Size() int {
	return xxx_messageInfo_AddUserRes.Size(m)
}
func (m *AddUserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRes.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRes proto.InternalMessageInfo

type QueryUserReq struct {
	Uids                 []string `protobuf:"bytes,1,rep,name=uids,proto3" json:"uids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserReq) Reset()         { *m = QueryUserReq{} }
func (m *QueryUserReq) String() string { return proto.CompactTextString(m) }
func (*QueryUserReq) ProtoMessage()    {}
func (*QueryUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{2}
}

func (m *QueryUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserReq.Unmarshal(m, b)
}
func (m *QueryUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserReq.Marshal(b, m, deterministic)
}
func (m *QueryUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserReq.Merge(m, src)
}
func (m *QueryUserReq) XXX_Size() int {
	return xxx_messageInfo_QueryUserReq.Size(m)
}
func (m *QueryUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserReq proto.InternalMessageInfo

func (m *QueryUserReq) GetUids() []string {
	if m != nil {
		return m.Uids
	}
	return nil
}

type QueryUserRes struct {
	Users                []*UserRecord `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueryUserRes) Reset()         { *m = QueryUserRes{} }
func (m *QueryUserRes) String() string { return proto.CompactTextString(m) }
func (*QueryUserRes) ProtoMessage()    {}
func (*QueryUserRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{3}
}

func (m *QueryUserRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserRes.Unmarshal(m, b)
}
func (m *QueryUserRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserRes.Marshal(b, m, deterministic)
}
func (m *QueryUserRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserRes.Merge(m, src)
}
func (m *QueryUserRes) XXX_Size() int {
	return xxx_messageInfo_QueryUserRes.Size(m)
}
func (m *QueryUserRes) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserRes.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserRes proto.InternalMessageInfo

func (m *QueryUserRes) GetUsers() []*UserRecord {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserRecord struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Agent                string   `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Avatar               string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRecord) Reset()         { *m = UserRecord{} }
func (m *UserRecord) String() string { return proto.CompactTextString(m) }
func (*UserRecord) ProtoMessage()    {}
func (*UserRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{4}
}

func (m *UserRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRecord.Unmarshal(m, b)
}
func (m *UserRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRecord.Marshal(b, m, deterministic)
}
func (m *UserRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRecord.Merge(m, src)
}
func (m *UserRecord) XXX_Size() int {
	return xxx_messageInfo_UserRecord.Size(m)
}
func (m *UserRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRecord.DiscardUnknown(m)
}

var xxx_messageInfo_UserRecord proto.InternalMessageInfo

func (m *UserRecord) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UserRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRecord) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *UserRecord) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRecord) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func init() {
	proto.RegisterType((*AddUserReq)(nil), "liveuser.AddUserReq")
	proto.RegisterType((*AddUserRes)(nil), "liveuser.AddUserRes")
	proto.RegisterType((*QueryUserReq)(nil), "liveuser.QueryUserReq")
	proto.RegisterType((*QueryUserRes)(nil), "liveuser.QueryUserRes")
	proto.RegisterType((*UserRecord)(nil), "liveuser.UserRecord")
}

func init() {
	proto.RegisterFile("liveuser.proto", fileDescriptor_6125f2c4dfd0d26d)
}

var fileDescriptor_6125f2c4dfd0d26d = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x37, 0x5b, 0xdd, 0xb1, 0x88, 0x0c, 0xa5, 0x84, 0x9e, 0x4a, 0x4e, 0x8b, 0x87,
	0x1e, 0x2a, 0x7a, 0xd0, 0x93, 0x8f, 0x60, 0xc0, 0x07, 0x88, 0x66, 0x90, 0x40, 0xdb, 0xd5, 0x64,
	0xb3, 0xa0, 0x4f, 0x2f, 0x99, 0xb4, 0xdd, 0x15, 0xda, 0xdb, 0xfc, 0x1f, 0x7f, 0xe6, 0x9f, 0x99,
	0xc0, 0xcd, 0xc6, 0xf5, 0x14, 0x03, 0xf9, 0xd5, 0x97, 0x6f, 0xbb, 0x16, 0xaf, 0x0e, 0x5a, 0x3d,
	0x02, 0xbc, 0x58, 0xfb, 0x16, 0xc8, 0x6b, 0xfa, 0xc6, 0x06, 0x44, 0xa2, 0xb2, 0x58, 0x16, 0xcd,
	0xf5, 0x7a, 0xb6, 0x3a, 0x3e, 0xcb, 0x86, 0x8f, 0xd6, 0x5b, 0xcd, 0x0e, 0x35, 0x1d, 0xbd, 0x0b,
	0x4a, 0xc1, 0xf4, 0x35, 0x92, 0xff, 0x39, 0xf4, 0x41, 0x10, 0xd1, 0xd9, 0x20, 0x8b, 0x65, 0xd9,
	0xd4, 0x9a, 0x6b, 0xf5, 0xf4, 0xcf, 0x13, 0xf0, 0x0e, 0xaa, 0xd4, 0x29, 0x9b, 0xce, 0x85, 0x65,
	0x8b, 0xea, 0x01, 0x06, 0x88, 0xb7, 0x50, 0x46, 0x67, 0x79, 0xc8, 0x5a, 0xa7, 0x32, 0xe5, 0xed,
	0xcc, 0x96, 0xe4, 0x05, 0x23, 0xae, 0x71, 0x06, 0x95, 0xf9, 0xa4, 0x5d, 0x27, 0x4b, 0x86, 0x59,
	0x24, 0x4a, 0x5b, 0xe3, 0x36, 0x52, 0x64, 0xca, 0x02, 0xe7, 0x30, 0x31, 0xbd, 0xe9, 0x8c, 0x97,
	0x15, 0xe3, 0xbd, 0x5a, 0xff, 0x82, 0x48, 0xb9, 0xf8, 0x00, 0x97, 0xfb, 0x6d, 0x71, 0x34, 0xe7,
	0x70, 0xb8, 0xc5, 0x29, 0x1a, 0xf0, 0x19, 0xea, 0xe3, 0xca, 0x38, 0x1f, 0x2c, 0xe3, 0x5b, 0x2d,
	0x4e, 0xf3, 0xf0, 0x3e, 0xe1, 0xaf, 0xba, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x39, 0xa6, 0xd8,
	0x7c, 0xbc, 0x01, 0x00, 0x00,
}
