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

type OrderRecord_State int32

const (
	OrderRecord_CREATED OrderRecord_State = 0
	OrderRecord_PAID    OrderRecord_State = 1
)

var OrderRecord_State_name = map[int32]string{
	0: "CREATED",
	1: "PAID",
}

var OrderRecord_State_value = map[string]int32{
	"CREATED": 0,
	"PAID":    1,
}

func (x OrderRecord_State) String() string {
	return proto.EnumName(OrderRecord_State_name, int32(x))
}

func (OrderRecord_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{18, 0}
}

type SetUserInfoReq struct {
	// 需要更新的用户信息。其中 uid, update_at 可以不填。
	User                 *UserRecord `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetUserInfoReq) Reset()         { *m = SetUserInfoReq{} }
func (m *SetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*SetUserInfoReq) ProtoMessage()    {}
func (*SetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{0}
}

func (m *SetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserInfoReq.Unmarshal(m, b)
}
func (m *SetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserInfoReq.Marshal(b, m, deterministic)
}
func (m *SetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserInfoReq.Merge(m, src)
}
func (m *SetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_SetUserInfoReq.Size(m)
}
func (m *SetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserInfoReq proto.InternalMessageInfo

func (m *SetUserInfoReq) GetUser() *UserRecord {
	if m != nil {
		return m.User
	}
	return nil
}

type SetUserInfoRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserInfoRes) Reset()         { *m = SetUserInfoRes{} }
func (m *SetUserInfoRes) String() string { return proto.CompactTextString(m) }
func (*SetUserInfoRes) ProtoMessage()    {}
func (*SetUserInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{1}
}

func (m *SetUserInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserInfoRes.Unmarshal(m, b)
}
func (m *SetUserInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserInfoRes.Marshal(b, m, deterministic)
}
func (m *SetUserInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserInfoRes.Merge(m, src)
}
func (m *SetUserInfoRes) XXX_Size() int {
	return xxx_messageInfo_SetUserInfoRes.Size(m)
}
func (m *SetUserInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserInfoRes proto.InternalMessageInfo

type GetUserInfoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoReq) Reset()         { *m = GetUserInfoReq{} }
func (m *GetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReq) ProtoMessage()    {}
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{2}
}

func (m *GetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReq.Unmarshal(m, b)
}
func (m *GetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReq.Marshal(b, m, deterministic)
}
func (m *GetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReq.Merge(m, src)
}
func (m *GetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReq.Size(m)
}
func (m *GetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReq proto.InternalMessageInfo

type GetUserInfoRes struct {
	User                 *UserRecord `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetUserInfoRes) Reset()         { *m = GetUserInfoRes{} }
func (m *GetUserInfoRes) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRes) ProtoMessage()    {}
func (*GetUserInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{3}
}

func (m *GetUserInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoRes.Unmarshal(m, b)
}
func (m *GetUserInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoRes.Marshal(b, m, deterministic)
}
func (m *GetUserInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRes.Merge(m, src)
}
func (m *GetUserInfoRes) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoRes.Size(m)
}
func (m *GetUserInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRes proto.InternalMessageInfo

func (m *GetUserInfoRes) GetUser() *UserRecord {
	if m != nil {
		return m.User
	}
	return nil
}

// UserRecord 是用户信息
type UserRecord struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Agent                string   `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	UpdateAt             int64    `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
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

func (m *UserRecord) GetUpdateAt() int64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

func (m *UserRecord) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type LoginReq struct {
	// 用户登录凭证（有效期五分钟）。开发者需要在开发者服务器后台调用 auth.code2Session，
	// 使用 code 换取 openid、unionid、session_key 等信息
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{5}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type LoginRes struct {
	// 该用户调用开发者服务器后台的凭据，用来识别用户身份
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// 用来换取新的 access_token，客户端应该保存在本地存储
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// access_token 凭证到期的时间，格式为Unix时间戳
	Expiry               int64    `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{6}
}

func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginRes) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LoginRes) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type RenewTokenReq struct {
	// 客户端保存在本地存储的上次的 refresh_token
	RefreshToken         string   `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewTokenReq) Reset()         { *m = RenewTokenReq{} }
func (m *RenewTokenReq) String() string { return proto.CompactTextString(m) }
func (*RenewTokenReq) ProtoMessage()    {}
func (*RenewTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{7}
}

func (m *RenewTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewTokenReq.Unmarshal(m, b)
}
func (m *RenewTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewTokenReq.Marshal(b, m, deterministic)
}
func (m *RenewTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewTokenReq.Merge(m, src)
}
func (m *RenewTokenReq) XXX_Size() int {
	return xxx_messageInfo_RenewTokenReq.Size(m)
}
func (m *RenewTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_RenewTokenReq proto.InternalMessageInfo

func (m *RenewTokenReq) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RenewTokenRes struct {
	// 该用户调用开发者服务器后台的凭据，用来识别用户身份
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// 用来换取新的 access_token，客户端应该保存在本地存储。
	// 取决于是否开启了 Refresh Token Rotation，它可能与请求时的 refresh_token 不同
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// access_token 凭证到期的时间，格式为Unix时间戳
	Expiry               int64    `protobuf:"varint,3,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewTokenRes) Reset()         { *m = RenewTokenRes{} }
func (m *RenewTokenRes) String() string { return proto.CompactTextString(m) }
func (*RenewTokenRes) ProtoMessage()    {}
func (*RenewTokenRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{8}
}

func (m *RenewTokenRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewTokenRes.Unmarshal(m, b)
}
func (m *RenewTokenRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewTokenRes.Marshal(b, m, deterministic)
}
func (m *RenewTokenRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewTokenRes.Merge(m, src)
}
func (m *RenewTokenRes) XXX_Size() int {
	return xxx_messageInfo_RenewTokenRes.Size(m)
}
func (m *RenewTokenRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewTokenRes.DiscardUnknown(m)
}

var xxx_messageInfo_RenewTokenRes proto.InternalMessageInfo

func (m *RenewTokenRes) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *RenewTokenRes) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RenewTokenRes) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type PrepayReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepayReq) Reset()         { *m = PrepayReq{} }
func (m *PrepayReq) String() string { return proto.CompactTextString(m) }
func (*PrepayReq) ProtoMessage()    {}
func (*PrepayReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{9}
}

func (m *PrepayReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepayReq.Unmarshal(m, b)
}
func (m *PrepayReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepayReq.Marshal(b, m, deterministic)
}
func (m *PrepayReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepayReq.Merge(m, src)
}
func (m *PrepayReq) XXX_Size() int {
	return xxx_messageInfo_PrepayReq.Size(m)
}
func (m *PrepayReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepayReq.DiscardUnknown(m)
}

var xxx_messageInfo_PrepayReq proto.InternalMessageInfo

type PrepayRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepayRes) Reset()         { *m = PrepayRes{} }
func (m *PrepayRes) String() string { return proto.CompactTextString(m) }
func (*PrepayRes) ProtoMessage()    {}
func (*PrepayRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{10}
}

func (m *PrepayRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepayRes.Unmarshal(m, b)
}
func (m *PrepayRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepayRes.Marshal(b, m, deterministic)
}
func (m *PrepayRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepayRes.Merge(m, src)
}
func (m *PrepayRes) XXX_Size() int {
	return xxx_messageInfo_PrepayRes.Size(m)
}
func (m *PrepayRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepayRes.DiscardUnknown(m)
}

var xxx_messageInfo_PrepayRes proto.InternalMessageInfo

type PostpayReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostpayReq) Reset()         { *m = PostpayReq{} }
func (m *PostpayReq) String() string { return proto.CompactTextString(m) }
func (*PostpayReq) ProtoMessage()    {}
func (*PostpayReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{11}
}

func (m *PostpayReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostpayReq.Unmarshal(m, b)
}
func (m *PostpayReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostpayReq.Marshal(b, m, deterministic)
}
func (m *PostpayReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostpayReq.Merge(m, src)
}
func (m *PostpayReq) XXX_Size() int {
	return xxx_messageInfo_PostpayReq.Size(m)
}
func (m *PostpayReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PostpayReq.DiscardUnknown(m)
}

var xxx_messageInfo_PostpayReq proto.InternalMessageInfo

type PostpayRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostpayRes) Reset()         { *m = PostpayRes{} }
func (m *PostpayRes) String() string { return proto.CompactTextString(m) }
func (*PostpayRes) ProtoMessage()    {}
func (*PostpayRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{12}
}

func (m *PostpayRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostpayRes.Unmarshal(m, b)
}
func (m *PostpayRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostpayRes.Marshal(b, m, deterministic)
}
func (m *PostpayRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostpayRes.Merge(m, src)
}
func (m *PostpayRes) XXX_Size() int {
	return xxx_messageInfo_PostpayRes.Size(m)
}
func (m *PostpayRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PostpayRes.DiscardUnknown(m)
}

var xxx_messageInfo_PostpayRes proto.InternalMessageInfo

type CreateOrderReq struct {
	Products             []*OrderProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateOrderReq) Reset()         { *m = CreateOrderReq{} }
func (m *CreateOrderReq) String() string { return proto.CompactTextString(m) }
func (*CreateOrderReq) ProtoMessage()    {}
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{13}
}

func (m *CreateOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderReq.Unmarshal(m, b)
}
func (m *CreateOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderReq.Marshal(b, m, deterministic)
}
func (m *CreateOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderReq.Merge(m, src)
}
func (m *CreateOrderReq) XXX_Size() int {
	return xxx_messageInfo_CreateOrderReq.Size(m)
}
func (m *CreateOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderReq proto.InternalMessageInfo

func (m *CreateOrderReq) GetProducts() []*OrderProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

type CreateOrderRes struct {
	Order                *OrderRecord `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOrderRes) Reset()         { *m = CreateOrderRes{} }
func (m *CreateOrderRes) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRes) ProtoMessage()    {}
func (*CreateOrderRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{14}
}

func (m *CreateOrderRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRes.Unmarshal(m, b)
}
func (m *CreateOrderRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRes.Marshal(b, m, deterministic)
}
func (m *CreateOrderRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRes.Merge(m, src)
}
func (m *CreateOrderRes) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRes.Size(m)
}
func (m *CreateOrderRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRes proto.InternalMessageInfo

func (m *CreateOrderRes) GetOrder() *OrderRecord {
	if m != nil {
		return m.Order
	}
	return nil
}

type ListOrderReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrderReq) Reset()         { *m = ListOrderReq{} }
func (m *ListOrderReq) String() string { return proto.CompactTextString(m) }
func (*ListOrderReq) ProtoMessage()    {}
func (*ListOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{15}
}

func (m *ListOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrderReq.Unmarshal(m, b)
}
func (m *ListOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrderReq.Marshal(b, m, deterministic)
}
func (m *ListOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrderReq.Merge(m, src)
}
func (m *ListOrderReq) XXX_Size() int {
	return xxx_messageInfo_ListOrderReq.Size(m)
}
func (m *ListOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrderReq proto.InternalMessageInfo

type ListOrderRes struct {
	Orders               []*OrderRecord `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListOrderRes) Reset()         { *m = ListOrderRes{} }
func (m *ListOrderRes) String() string { return proto.CompactTextString(m) }
func (*ListOrderRes) ProtoMessage()    {}
func (*ListOrderRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{16}
}

func (m *ListOrderRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrderRes.Unmarshal(m, b)
}
func (m *ListOrderRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrderRes.Marshal(b, m, deterministic)
}
func (m *ListOrderRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrderRes.Merge(m, src)
}
func (m *ListOrderRes) XXX_Size() int {
	return xxx_messageInfo_ListOrderRes.Size(m)
}
func (m *ListOrderRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrderRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrderRes proto.InternalMessageInfo

func (m *ListOrderRes) GetOrders() []*OrderRecord {
	if m != nil {
		return m.Orders
	}
	return nil
}

type OrderProduct struct {
	OrderId              uint64   `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId            string   `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price                int32    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Count                int32    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	ProductSnapshot      uint64   `protobuf:"varint,6,opt,name=product_snapshot,json=productSnapshot,proto3" json:"product_snapshot,omitempty"`
	Detail               string   `protobuf:"bytes,7,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderProduct) Reset()         { *m = OrderProduct{} }
func (m *OrderProduct) String() string { return proto.CompactTextString(m) }
func (*OrderProduct) ProtoMessage()    {}
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{17}
}

func (m *OrderProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderProduct.Unmarshal(m, b)
}
func (m *OrderProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderProduct.Marshal(b, m, deterministic)
}
func (m *OrderProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderProduct.Merge(m, src)
}
func (m *OrderProduct) XXX_Size() int {
	return xxx_messageInfo_OrderProduct.Size(m)
}
func (m *OrderProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderProduct.DiscardUnknown(m)
}

var xxx_messageInfo_OrderProduct proto.InternalMessageInfo

func (m *OrderProduct) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderProduct) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *OrderProduct) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderProduct) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderProduct) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OrderProduct) GetProductSnapshot() uint64 {
	if m != nil {
		return m.ProductSnapshot
	}
	return 0
}

func (m *OrderProduct) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type OrderRecord struct {
	Id                   uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string            `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	State                OrderRecord_State `protobuf:"varint,3,opt,name=state,proto3,enum=liveuser.OrderRecord_State" json:"state,omitempty"`
	Amount               int32             `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Discount             int32             `protobuf:"varint,5,opt,name=discount,proto3" json:"discount,omitempty"`
	Pay                  int32             `protobuf:"varint,6,opt,name=pay,proto3" json:"pay,omitempty"`
	PayAt                int64             `protobuf:"varint,7,opt,name=pay_at,json=payAt,proto3" json:"pay_at,omitempty"`
	Products             []*OrderProduct   `protobuf:"bytes,8,rep,name=products,proto3" json:"products,omitempty"`
	CreatedAt            int64             `protobuf:"varint,1002,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64             `protobuf:"varint,1003,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            int64             `protobuf:"varint,1004,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderRecord) Reset()         { *m = OrderRecord{} }
func (m *OrderRecord) String() string { return proto.CompactTextString(m) }
func (*OrderRecord) ProtoMessage()    {}
func (*OrderRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_6125f2c4dfd0d26d, []int{18}
}

func (m *OrderRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRecord.Unmarshal(m, b)
}
func (m *OrderRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRecord.Marshal(b, m, deterministic)
}
func (m *OrderRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRecord.Merge(m, src)
}
func (m *OrderRecord) XXX_Size() int {
	return xxx_messageInfo_OrderRecord.Size(m)
}
func (m *OrderRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRecord.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRecord proto.InternalMessageInfo

func (m *OrderRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderRecord) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *OrderRecord) GetState() OrderRecord_State {
	if m != nil {
		return m.State
	}
	return OrderRecord_CREATED
}

func (m *OrderRecord) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *OrderRecord) GetDiscount() int32 {
	if m != nil {
		return m.Discount
	}
	return 0
}

func (m *OrderRecord) GetPay() int32 {
	if m != nil {
		return m.Pay
	}
	return 0
}

func (m *OrderRecord) GetPayAt() int64 {
	if m != nil {
		return m.PayAt
	}
	return 0
}

func (m *OrderRecord) GetProducts() []*OrderProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *OrderRecord) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *OrderRecord) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *OrderRecord) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("liveuser.OrderRecord_State", OrderRecord_State_name, OrderRecord_State_value)
	proto.RegisterType((*SetUserInfoReq)(nil), "liveuser.SetUserInfoReq")
	proto.RegisterType((*SetUserInfoRes)(nil), "liveuser.SetUserInfoRes")
	proto.RegisterType((*GetUserInfoReq)(nil), "liveuser.GetUserInfoReq")
	proto.RegisterType((*GetUserInfoRes)(nil), "liveuser.GetUserInfoRes")
	proto.RegisterType((*UserRecord)(nil), "liveuser.UserRecord")
	proto.RegisterType((*LoginReq)(nil), "liveuser.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "liveuser.LoginRes")
	proto.RegisterType((*RenewTokenReq)(nil), "liveuser.RenewTokenReq")
	proto.RegisterType((*RenewTokenRes)(nil), "liveuser.RenewTokenRes")
	proto.RegisterType((*PrepayReq)(nil), "liveuser.PrepayReq")
	proto.RegisterType((*PrepayRes)(nil), "liveuser.PrepayRes")
	proto.RegisterType((*PostpayReq)(nil), "liveuser.PostpayReq")
	proto.RegisterType((*PostpayRes)(nil), "liveuser.PostpayRes")
	proto.RegisterType((*CreateOrderReq)(nil), "liveuser.CreateOrderReq")
	proto.RegisterType((*CreateOrderRes)(nil), "liveuser.CreateOrderRes")
	proto.RegisterType((*ListOrderReq)(nil), "liveuser.ListOrderReq")
	proto.RegisterType((*ListOrderRes)(nil), "liveuser.ListOrderRes")
	proto.RegisterType((*OrderProduct)(nil), "liveuser.OrderProduct")
	proto.RegisterType((*OrderRecord)(nil), "liveuser.OrderRecord")
}

func init() {
	proto.RegisterFile("liveuser.proto", fileDescriptor_6125f2c4dfd0d26d)
}

var fileDescriptor_6125f2c4dfd0d26d = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x2d, 0x25, 0x92, 0xa2, 0x46, 0xb2, 0x2a, 0x4c, 0x6d, 0x97, 0xa5, 0x51, 0xc3, 0x65, 0x2f,
	0x2a, 0x8a, 0xba, 0x28, 0xfb, 0x01, 0x34, 0x88, 0x03, 0x08, 0xb6, 0x61, 0x08, 0x30, 0x10, 0x81,
	0x72, 0x90, 0xa3, 0xb1, 0x21, 0xd7, 0x36, 0x13, 0x9b, 0xa4, 0xb9, 0x2b, 0xc7, 0x02, 0xe2, 0x5f,
	0x98, 0x3f, 0x91, 0x73, 0x92, 0x5b, 0xfe, 0x40, 0xb0, 0x1f, 0x12, 0x49, 0x5b, 0x02, 0x92, 0x43,
	0x6e, 0x3b, 0xef, 0xbd, 0x19, 0xce, 0xee, 0xce, 0x3e, 0x42, 0xef, 0x32, 0xb9, 0xa1, 0x53, 0x46,
	0x8b, 0xdd, 0xbc, 0xc8, 0x78, 0x86, 0xce, 0x3c, 0xf6, 0x1f, 0x41, 0x6f, 0x42, 0xf9, 0x33, 0x46,
	0x8b, 0x51, 0x7a, 0x96, 0x85, 0xf4, 0x1a, 0x07, 0x60, 0x0a, 0xc6, 0x35, 0x76, 0x8c, 0x41, 0x27,
	0x58, 0xdf, 0x5d, 0xa4, 0x0a, 0x51, 0x48, 0xa3, 0xac, 0x88, 0x43, 0xa9, 0xf0, 0xfb, 0xf7, 0x72,
	0x99, 0x40, 0x8e, 0x6a, 0xd5, 0x44, 0xfd, 0x1a, 0xc2, 0xbe, 0xa2, 0xfe, 0x1d, 0x40, 0x89, 0x61,
	0x1f, 0x9a, 0xd3, 0x24, 0x96, 0x69, 0xed, 0x50, 0x2c, 0x11, 0xc1, 0x4c, 0xc9, 0x15, 0x75, 0x1b,
	0x12, 0x92, 0x6b, 0x5c, 0x07, 0x8b, 0x9c, 0xd3, 0x94, 0xbb, 0x4d, 0x09, 0xaa, 0x00, 0xb7, 0xa0,
	0x3d, 0xcd, 0x63, 0xc2, 0xe9, 0x29, 0xe1, 0xae, 0xb9, 0x63, 0x0c, 0x9a, 0xa1, 0xa3, 0x80, 0x21,
	0xc7, 0x4d, 0xb0, 0xc9, 0x0d, 0xe1, 0xa4, 0x70, 0x2d, 0x99, 0xa3, 0x23, 0x7f, 0x1b, 0x9c, 0xe3,
	0xec, 0x3c, 0x49, 0xc5, 0xa1, 0x20, 0x98, 0x51, 0x16, 0x53, 0xfd, 0x75, 0xb9, 0xf6, 0x5f, 0x2e,
	0x78, 0x86, 0xbf, 0x40, 0x97, 0x44, 0x11, 0x65, 0xec, 0x94, 0x67, 0xaf, 0x68, 0xaa, 0x75, 0x1d,
	0x85, 0x9d, 0x08, 0x08, 0x7f, 0x85, 0xb5, 0x82, 0x9e, 0x15, 0x94, 0x5d, 0x68, 0x8d, 0x6a, 0xbb,
	0xab, 0x41, 0x25, 0xda, 0x04, 0x9b, 0xde, 0xe6, 0x49, 0x31, 0x93, 0xfd, 0x37, 0x43, 0x1d, 0xf9,
	0xff, 0xc0, 0x5a, 0x48, 0x53, 0xfa, 0x5a, 0xaa, 0x44, 0x43, 0x0f, 0xaa, 0x19, 0x0f, 0xab, 0xf9,
	0x59, 0x3d, 0xeb, 0xdb, 0xb7, 0xd9, 0x81, 0xf6, 0xb8, 0xa0, 0x39, 0x99, 0x89, 0xab, 0xaf, 0x04,
	0xcc, 0xef, 0x02, 0x8c, 0x33, 0xc6, 0x35, 0x55, 0x8d, 0x98, 0x7f, 0x00, 0xbd, 0xfd, 0x82, 0x12,
	0x4e, 0x9f, 0x16, 0xb1, 0xb8, 0xee, 0x6b, 0x0c, 0xc0, 0xc9, 0x8b, 0x2c, 0x9e, 0x46, 0x9c, 0xb9,
	0xc6, 0x4e, 0x73, 0xd0, 0x09, 0x36, 0xcb, 0x39, 0x91, 0xaa, 0xb1, 0xa2, 0xc3, 0x85, 0xce, 0xdf,
	0xbb, 0x57, 0x85, 0xe1, 0xef, 0x60, 0x65, 0x62, 0xad, 0x47, 0x6d, 0xe3, 0x5e, 0x09, 0x3d, 0x6b,
	0x4a, 0xe3, 0xf7, 0xa0, 0x7b, 0x9c, 0x30, 0x3e, 0x6f, 0xc1, 0xdf, 0xab, 0xc5, 0x0c, 0xff, 0x00,
	0x5b, 0x0a, 0xe7, 0x0d, 0xad, 0xa8, 0xa6, 0x45, 0xfe, 0x5b, 0x03, 0xba, 0xd5, 0x46, 0xf1, 0x27,
	0x70, 0x24, 0x75, 0xaa, 0x67, 0xd8, 0x0c, 0x5b, 0x32, 0x1e, 0xc5, 0xf8, 0x33, 0x80, 0xde, 0x85,
	0x20, 0xd5, 0x79, 0xb7, 0x35, 0x32, 0x2a, 0xc7, 0xbc, 0x59, 0x1f, 0xf3, 0xbc, 0x48, 0x22, 0x2a,
	0x87, 0xd9, 0x0a, 0x55, 0x20, 0xd0, 0x28, 0x9b, 0xa6, 0x5c, 0x0e, 0xb2, 0x15, 0xaa, 0x00, 0x7f,
	0x83, 0xfe, 0xbc, 0x3c, 0x4b, 0x49, 0xce, 0x2e, 0x32, 0xee, 0xda, 0xb2, 0x83, 0xef, 0x35, 0x3e,
	0xd1, 0xb0, 0xb8, 0xd7, 0x98, 0x72, 0x92, 0x5c, 0xba, 0x2d, 0xf5, 0x14, 0x54, 0xe4, 0x7f, 0x6a,
	0x40, 0xa7, 0xb2, 0x4b, 0xec, 0x41, 0x63, 0xb1, 0x8d, 0x46, 0xb2, 0x78, 0x9b, 0x8d, 0xf2, 0x6d,
	0xfe, 0x05, 0x16, 0xe3, 0x84, 0xab, 0xae, 0x7b, 0xc1, 0xd6, 0xd2, 0xd3, 0xda, 0x9d, 0x08, 0x49,
	0xa8, 0x94, 0xf2, 0x1d, 0x5e, 0xc9, 0xf6, 0xd5, 0xa6, 0x74, 0x84, 0x1e, 0x38, 0x71, 0xc2, 0xaa,
	0x1b, 0x5b, 0xc4, 0xe2, 0xc3, 0x39, 0x99, 0xc9, 0xed, 0x58, 0xa1, 0x58, 0xe2, 0x06, 0xd8, 0x39,
	0x99, 0x89, 0x77, 0xde, 0x92, 0xa3, 0x69, 0xe5, 0x64, 0x36, 0xe4, 0xb5, 0x89, 0x72, 0xbe, 0x6c,
	0xa2, 0x70, 0x1b, 0x20, 0x92, 0x13, 0x15, 0x8b, 0x72, 0xef, 0x55, 0xbd, 0xb6, 0x86, 0x86, 0x5c,
	0xf0, 0xca, 0x44, 0x24, 0xff, 0x41, 0xf3, 0x1a, 0x52, 0x7c, 0x4c, 0x2f, 0xa9, 0xe6, 0x3f, 0x6a,
	0x5e, 0x43, 0x43, 0xee, 0x6f, 0x83, 0x25, 0x0f, 0x00, 0x3b, 0xd0, 0xda, 0x0f, 0x0f, 0x87, 0x27,
	0x87, 0x07, 0xfd, 0xef, 0xd0, 0x01, 0x73, 0x3c, 0x1c, 0x1d, 0xf4, 0x8d, 0xe0, 0x0d, 0x98, 0xc2,
	0xff, 0xf0, 0x7f, 0x68, 0x4e, 0x28, 0x47, 0xb7, 0x6c, 0xb8, 0x6e, 0xd9, 0xde, 0x2a, 0x86, 0x89,
	0xd4, 0xa3, 0x7a, 0xea, 0xd1, 0xca, 0xd4, 0xba, 0x4f, 0x07, 0xef, 0x0c, 0x68, 0x3c, 0xbf, 0xc5,
	0x3f, 0xc1, 0x92, 0x2e, 0x87, 0x58, 0x2a, 0xe7, 0xb6, 0xe8, 0x3d, 0xc4, 0x18, 0x3e, 0x01, 0x28,
	0x4d, 0x07, 0x7f, 0x2c, 0x15, 0x35, 0x03, 0xf3, 0x56, 0x10, 0x0c, 0x03, 0xb0, 0x95, 0x6d, 0xe0,
	0x0f, 0xa5, 0x64, 0xe1, 0x2a, 0xde, 0x12, 0x90, 0xe1, 0xbf, 0xd0, 0xd2, 0x7e, 0x82, 0x95, 0x1f,
	0x4a, 0x69, 0x38, 0xde, 0x32, 0x94, 0x05, 0x77, 0x60, 0xc9, 0xab, 0xc7, 0xc7, 0x60, 0x2b, 0xef,
	0xa8, 0x9e, 0x54, 0xdd, 0x93, 0xbc, 0x55, 0x0c, 0xc3, 0xff, 0xc0, 0x14, 0x56, 0x81, 0x95, 0x89,
	0xaa, 0x5a, 0x89, 0xb7, 0x1c, 0x67, 0x2f, 0x6c, 0xf9, 0x33, 0xfe, 0xfb, 0x73, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x24, 0xe9, 0x62, 0x1b, 0x9e, 0x07, 0x00, 0x00,
}
