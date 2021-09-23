// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catalog.proto

package catalog

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

type ListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReq) Reset()         { *m = ListReq{} }
func (m *ListReq) String() string { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()    {}
func (*ListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{0}
}

func (m *ListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReq.Unmarshal(m, b)
}
func (m *ListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReq.Marshal(b, m, deterministic)
}
func (m *ListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReq.Merge(m, src)
}
func (m *ListReq) XXX_Size() int {
	return xxx_messageInfo_ListReq.Size(m)
}
func (m *ListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListReq proto.InternalMessageInfo

type ListRes struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListRes) Reset()         { *m = ListRes{} }
func (m *ListRes) String() string { return proto.CompactTextString(m) }
func (*ListRes) ProtoMessage()    {}
func (*ListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{1}
}

func (m *ListRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRes.Unmarshal(m, b)
}
func (m *ListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRes.Marshal(b, m, deterministic)
}
func (m *ListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRes.Merge(m, src)
}
func (m *ListRes) XXX_Size() int {
	return xxx_messageInfo_ListRes.Size(m)
}
func (m *ListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListRes proto.InternalMessageInfo

func (m *ListRes) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type Product struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PriceCent int32  `protobuf:"varint,3,opt,name=price_cent,json=priceCent,proto3" json:"price_cent,omitempty"`
	// 产品快照编号：记录历史时间线上的一个产品，被订单所引用。
	Snapshot             uint64   `protobuf:"varint,1001,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	CreatedAt            int64    `protobuf:"varint,1002,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,1003,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            int64    `protobuf:"varint,1004,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Operator             string   `protobuf:"bytes,1005,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{2}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPriceCent() int32 {
	if m != nil {
		return m.PriceCent
	}
	return 0
}

func (m *Product) GetSnapshot() uint64 {
	if m != nil {
		return m.Snapshot
	}
	return 0
}

func (m *Product) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Product) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Product) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func (m *Product) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type SetReq struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReq) Reset()         { *m = SetReq{} }
func (m *SetReq) String() string { return proto.CompactTextString(m) }
func (*SetReq) ProtoMessage()    {}
func (*SetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{3}
}

func (m *SetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReq.Unmarshal(m, b)
}
func (m *SetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReq.Marshal(b, m, deterministic)
}
func (m *SetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReq.Merge(m, src)
}
func (m *SetReq) XXX_Size() int {
	return xxx_messageInfo_SetReq.Size(m)
}
func (m *SetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetReq proto.InternalMessageInfo

func (m *SetReq) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type SetRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRes) Reset()         { *m = SetRes{} }
func (m *SetRes) String() string { return proto.CompactTextString(m) }
func (*SetRes) ProtoMessage()    {}
func (*SetRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{4}
}

func (m *SetRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRes.Unmarshal(m, b)
}
func (m *SetRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRes.Marshal(b, m, deterministic)
}
func (m *SetRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRes.Merge(m, src)
}
func (m *SetRes) XXX_Size() int {
	return xxx_messageInfo_SetRes.Size(m)
}
func (m *SetRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRes.DiscardUnknown(m)
}

var xxx_messageInfo_SetRes proto.InternalMessageInfo

type DeleteReq struct {
	ProductId            string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{5}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

type DeleteRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRes) Reset()         { *m = DeleteRes{} }
func (m *DeleteRes) String() string { return proto.CompactTextString(m) }
func (*DeleteRes) ProtoMessage()    {}
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{6}
}

func (m *DeleteRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRes.Unmarshal(m, b)
}
func (m *DeleteRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRes.Marshal(b, m, deterministic)
}
func (m *DeleteRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRes.Merge(m, src)
}
func (m *DeleteRes) XXX_Size() int {
	return xxx_messageInfo_DeleteRes.Size(m)
}
func (m *DeleteRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRes proto.InternalMessageInfo

type FindByIDReq struct {
	ProductIds           []string `protobuf:"bytes,1,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDReq) Reset()         { *m = FindByIDReq{} }
func (m *FindByIDReq) String() string { return proto.CompactTextString(m) }
func (*FindByIDReq) ProtoMessage()    {}
func (*FindByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{7}
}

func (m *FindByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDReq.Unmarshal(m, b)
}
func (m *FindByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDReq.Marshal(b, m, deterministic)
}
func (m *FindByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDReq.Merge(m, src)
}
func (m *FindByIDReq) XXX_Size() int {
	return xxx_messageInfo_FindByIDReq.Size(m)
}
func (m *FindByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDReq proto.InternalMessageInfo

func (m *FindByIDReq) GetProductIds() []string {
	if m != nil {
		return m.ProductIds
	}
	return nil
}

type FindByIDRes struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindByIDRes) Reset()         { *m = FindByIDRes{} }
func (m *FindByIDRes) String() string { return proto.CompactTextString(m) }
func (*FindByIDRes) ProtoMessage()    {}
func (*FindByIDRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_0abbfcf058acdf89, []int{8}
}

func (m *FindByIDRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDRes.Unmarshal(m, b)
}
func (m *FindByIDRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDRes.Marshal(b, m, deterministic)
}
func (m *FindByIDRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDRes.Merge(m, src)
}
func (m *FindByIDRes) XXX_Size() int {
	return xxx_messageInfo_FindByIDRes.Size(m)
}
func (m *FindByIDRes) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDRes.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDRes proto.InternalMessageInfo

func (m *FindByIDRes) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*ListReq)(nil), "catalog.ListReq")
	proto.RegisterType((*ListRes)(nil), "catalog.ListRes")
	proto.RegisterType((*Product)(nil), "catalog.Product")
	proto.RegisterType((*SetReq)(nil), "catalog.SetReq")
	proto.RegisterType((*SetRes)(nil), "catalog.SetRes")
	proto.RegisterType((*DeleteReq)(nil), "catalog.DeleteReq")
	proto.RegisterType((*DeleteRes)(nil), "catalog.DeleteRes")
	proto.RegisterType((*FindByIDReq)(nil), "catalog.FindByIDReq")
	proto.RegisterType((*FindByIDRes)(nil), "catalog.FindByIDRes")
}

func init() {
	proto.RegisterFile("catalog.proto", fileDescriptor_0abbfcf058acdf89)
}

var fileDescriptor_0abbfcf058acdf89 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x87, 0x49, 0xd3, 0xdb, 0x24, 0x27, 0xdc, 0x7b, 0xe5, 0xe0, 0x22, 0x44, 0xd4, 0x30, 0x1b,
	0x43, 0x91, 0x22, 0x55, 0x74, 0xe1, 0xaa, 0xb6, 0x08, 0x05, 0x17, 0x92, 0x3e, 0x40, 0x89, 0x99,
	0x41, 0x03, 0x35, 0x13, 0x33, 0xd3, 0x85, 0x6f, 0xe8, 0x93, 0x08, 0xfe, 0x7b, 0x07, 0xc9, 0x64,
	0x32, 0x2d, 0xb5, 0x1b, 0x77, 0x93, 0xef, 0x3b, 0x67, 0x26, 0xe7, 0x37, 0x03, 0x7f, 0xb3, 0x54,
	0xa6, 0x0b, 0x7e, 0x3f, 0x28, 0x2b, 0x2e, 0x39, 0x3a, 0xfa, 0x93, 0x78, 0xe0, 0xdc, 0xe4, 0x42,
	0x26, 0xec, 0x89, 0x5c, 0xb4, 0x4b, 0x81, 0xc7, 0xe0, 0x96, 0x15, 0xa7, 0xcb, 0x4c, 0x8a, 0xc0,
	0x8a, 0xec, 0xd8, 0x1f, 0xee, 0x0c, 0xda, 0x0d, 0x6e, 0x1b, 0x91, 0x98, 0x0a, 0xf2, 0x6a, 0x81,
	0xa3, 0x29, 0xfe, 0x83, 0x4e, 0x4e, 0x03, 0x2b, 0xb2, 0x62, 0x2f, 0xe9, 0xe4, 0x14, 0x11, 0xba,
	0x45, 0xfa, 0xc8, 0x82, 0x8e, 0x22, 0x6a, 0x8d, 0xfb, 0x00, 0x65, 0x95, 0x67, 0x6c, 0x9e, 0xb1,
	0x42, 0x06, 0x76, 0x64, 0xc5, 0x7f, 0x12, 0x4f, 0x91, 0x31, 0x2b, 0x24, 0xee, 0x81, 0x2b, 0x8a,
	0xb4, 0x14, 0x0f, 0x5c, 0x06, 0x6f, 0x4e, 0x64, 0xc5, 0xdd, 0xc4, 0x00, 0x3c, 0x00, 0xc8, 0x2a,
	0x96, 0x4a, 0x46, 0xe7, 0xa9, 0x0c, 0xde, 0x6b, 0x6d, 0x27, 0x9e, 0x46, 0x23, 0xe5, 0x97, 0x25,
	0x6d, 0xfd, 0x87, 0xf6, 0x1a, 0x35, 0x9e, 0xb2, 0x05, 0xd3, 0xfe, 0x53, 0x7b, 0x8d, 0x46, 0xea,
	0x70, 0x5e, 0xb2, 0x2a, 0x95, 0xbc, 0x0a, 0xbe, 0x1c, 0xf5, 0xd3, 0x06, 0x90, 0x33, 0xe8, 0xcd,
	0x58, 0x9d, 0x15, 0xf6, 0xc1, 0xd1, 0xe3, 0xab, 0x59, 0xb7, 0xe5, 0xd3, 0x16, 0x10, 0x57, 0x77,
	0x09, 0xd2, 0x07, 0x6f, 0xa2, 0x4e, 0xaa, 0xb7, 0x50, 0x29, 0xa8, 0x8a, 0xb9, 0x49, 0xcc, 0xd3,
	0x64, 0x4a, 0x89, 0xbf, 0xaa, 0x15, 0x64, 0x00, 0xfe, 0x75, 0x5e, 0xd0, 0xab, 0xe7, 0xe9, 0xa4,
	0x6e, 0x3d, 0x04, 0x7f, 0xd5, 0xda, 0xdc, 0x90, 0x97, 0x80, 0xe9, 0x15, 0xe4, 0x72, 0xbd, 0xfe,
	0x97, 0xd7, 0x39, 0x7c, 0xb1, 0xc0, 0x19, 0x37, 0x16, 0xfb, 0xd0, 0xad, 0xdf, 0x04, 0xae, 0xea,
	0xf5, 0x6b, 0x09, 0x37, 0x89, 0xc0, 0x23, 0xb0, 0x67, 0x4c, 0xe2, 0x7f, 0x23, 0x9a, 0xac, 0xc2,
	0x0d, 0x20, 0xf0, 0x04, 0x7a, 0xcd, 0x68, 0x88, 0x46, 0x99, 0x5c, 0xc2, 0x9f, 0x4c, 0xe0, 0x39,
	0xb8, 0xed, 0x3c, 0xb8, 0x6b, 0xfc, 0x5a, 0x24, 0xe1, 0x36, 0x2a, 0xee, 0x7a, 0xea, 0xb5, 0x9f,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x35, 0x46, 0xbe, 0xfe, 0x02, 0x00, 0x00,
}
