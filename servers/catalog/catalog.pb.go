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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PriceCent            int32    `protobuf:"varint,3,opt,name=price_cent,json=priceCent,proto3" json:"price_cent,omitempty"`
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

func init() {
	proto.RegisterType((*ListReq)(nil), "catalog.ListReq")
	proto.RegisterType((*ListRes)(nil), "catalog.ListRes")
	proto.RegisterType((*Product)(nil), "catalog.Product")
	proto.RegisterType((*SetReq)(nil), "catalog.SetReq")
	proto.RegisterType((*SetRes)(nil), "catalog.SetRes")
	proto.RegisterType((*DeleteReq)(nil), "catalog.DeleteReq")
	proto.RegisterType((*DeleteRes)(nil), "catalog.DeleteRes")
}

func init() {
	proto.RegisterFile("catalog.proto", fileDescriptor_0abbfcf058acdf89)
}

var fileDescriptor_0abbfcf058acdf89 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdd, 0x6a, 0x83, 0x40,
	0x10, 0x85, 0x59, 0x4d, 0xdd, 0xec, 0x84, 0xfe, 0x30, 0x57, 0x22, 0x14, 0x64, 0x6f, 0x2a, 0x52,
	0x42, 0x49, 0x0b, 0x7d, 0x80, 0xf4, 0xa6, 0x90, 0x8b, 0x62, 0x1e, 0x20, 0x58, 0x1d, 0x8a, 0x90,
	0xaa, 0x75, 0xa7, 0xef, 0xd1, 0x47, 0x2e, 0xae, 0xeb, 0x06, 0xd2, 0xde, 0x8d, 0xdf, 0x1c, 0xcf,
	0x39, 0xcc, 0xc2, 0x65, 0x55, 0x72, 0x79, 0xec, 0x3e, 0xd6, 0xfd, 0xd0, 0x71, 0x87, 0xd2, 0x7d,
	0x6a, 0x05, 0x72, 0xd7, 0x18, 0x2e, 0xe8, 0x4b, 0x3f, 0xcf, 0xa3, 0xc1, 0x7b, 0x58, 0xf6, 0x43,
	0x57, 0x7f, 0x57, 0x6c, 0x62, 0x91, 0x86, 0xd9, 0x6a, 0x73, 0xb3, 0x9e, 0x0d, 0xde, 0xa6, 0x45,
	0xe1, 0x15, 0x7a, 0x07, 0xd2, 0x41, 0xbc, 0x82, 0xa0, 0xa9, 0x63, 0x91, 0x8a, 0x4c, 0x15, 0x41,
	0x53, 0x23, 0xc2, 0xa2, 0x2d, 0x3f, 0x29, 0x0e, 0x2c, 0xb1, 0x33, 0xde, 0x02, 0xf4, 0x43, 0x53,
	0xd1, 0xa1, 0xa2, 0x96, 0xe3, 0x30, 0x15, 0xd9, 0x45, 0xa1, 0x2c, 0xd9, 0x52, 0xcb, 0xfa, 0x09,
	0xa2, 0x3d, 0x8d, 0x85, 0x30, 0x07, 0xe9, 0x32, 0xac, 0xe3, 0x7f, 0x25, 0x66, 0x81, 0x5e, 0xba,
	0xbf, 0x8c, 0xce, 0x41, 0xbd, 0xd0, 0x91, 0x98, 0x46, 0x0b, 0x9b, 0x65, 0x15, 0x07, 0xdf, 0x4b,
	0x39, 0xf2, 0x5a, 0xeb, 0xd5, 0x49, 0x6b, 0x36, 0x3f, 0x02, 0xe4, 0x76, 0xf2, 0xc7, 0x1c, 0x16,
	0xe3, 0x2d, 0xf0, 0x94, 0xe8, 0xae, 0x94, 0x9c, 0x13, 0x83, 0x77, 0x10, 0xee, 0x89, 0xf1, 0xda,
	0x2f, 0xa6, 0xfa, 0xc9, 0x19, 0x30, 0xf8, 0x00, 0xd1, 0x94, 0x86, 0xe8, 0x57, 0xbe, 0x6a, 0xf2,
	0x97, 0x99, 0xf7, 0xc8, 0xbe, 0xd6, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x5d, 0x39,
	0x63, 0xbe, 0x01, 0x00, 0x00,
}
