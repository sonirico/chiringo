// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p.proto

package main

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

type RequestChain struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestChain) Reset()         { *m = RequestChain{} }
func (m *RequestChain) String() string { return proto.CompactTextString(m) }
func (*RequestChain) ProtoMessage()    {}
func (*RequestChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{0}
}

func (m *RequestChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestChain.Unmarshal(m, b)
}
func (m *RequestChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestChain.Marshal(b, m, deterministic)
}
func (m *RequestChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestChain.Merge(m, src)
}
func (m *RequestChain) XXX_Size() int {
	return xxx_messageInfo_RequestChain.Size(m)
}
func (m *RequestChain) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestChain.DiscardUnknown(m)
}

var xxx_messageInfo_RequestChain proto.InternalMessageInfo

type ResponseChain struct {
	Chain                *Chain   `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseChain) Reset()         { *m = ResponseChain{} }
func (m *ResponseChain) String() string { return proto.CompactTextString(m) }
func (*ResponseChain) ProtoMessage()    {}
func (*ResponseChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{1}
}

func (m *ResponseChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseChain.Unmarshal(m, b)
}
func (m *ResponseChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseChain.Marshal(b, m, deterministic)
}
func (m *ResponseChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseChain.Merge(m, src)
}
func (m *ResponseChain) XXX_Size() int {
	return xxx_messageInfo_ResponseChain.Size(m)
}
func (m *ResponseChain) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseChain.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseChain proto.InternalMessageInfo

func (m *ResponseChain) GetChain() *Chain {
	if m != nil {
		return m.Chain
	}
	return nil
}

type RequestLatestBlock struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestLatestBlock) Reset()         { *m = RequestLatestBlock{} }
func (m *RequestLatestBlock) String() string { return proto.CompactTextString(m) }
func (*RequestLatestBlock) ProtoMessage()    {}
func (*RequestLatestBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{2}
}

func (m *RequestLatestBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestLatestBlock.Unmarshal(m, b)
}
func (m *RequestLatestBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestLatestBlock.Marshal(b, m, deterministic)
}
func (m *RequestLatestBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestLatestBlock.Merge(m, src)
}
func (m *RequestLatestBlock) XXX_Size() int {
	return xxx_messageInfo_RequestLatestBlock.Size(m)
}
func (m *RequestLatestBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestLatestBlock.DiscardUnknown(m)
}

var xxx_messageInfo_RequestLatestBlock proto.InternalMessageInfo

type ResponseLatestBlock struct {
	Block                *Block   `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseLatestBlock) Reset()         { *m = ResponseLatestBlock{} }
func (m *ResponseLatestBlock) String() string { return proto.CompactTextString(m) }
func (*ResponseLatestBlock) ProtoMessage()    {}
func (*ResponseLatestBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7fdddb109e6467a, []int{3}
}

func (m *ResponseLatestBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseLatestBlock.Unmarshal(m, b)
}
func (m *ResponseLatestBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseLatestBlock.Marshal(b, m, deterministic)
}
func (m *ResponseLatestBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseLatestBlock.Merge(m, src)
}
func (m *ResponseLatestBlock) XXX_Size() int {
	return xxx_messageInfo_ResponseLatestBlock.Size(m)
}
func (m *ResponseLatestBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseLatestBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseLatestBlock proto.InternalMessageInfo

func (m *ResponseLatestBlock) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestChain)(nil), "main.RequestChain")
	proto.RegisterType((*ResponseChain)(nil), "main.ResponseChain")
	proto.RegisterType((*RequestLatestBlock)(nil), "main.RequestLatestBlock")
	proto.RegisterType((*ResponseLatestBlock)(nil), "main.ResponseLatestBlock")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_e7fdddb109e6467a) }

var fileDescriptor_e7fdddb109e6467a = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x30, 0x2a, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93, 0xe2, 0x4e, 0xce, 0x48,
	0xcc, 0xcc, 0x83, 0x08, 0x29, 0xf1, 0x71, 0xf1, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x38,
	0x83, 0x44, 0x95, 0x8c, 0xb8, 0x78, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xc1, 0x02,
	0x42, 0x8a, 0x5c, 0xac, 0x60, 0xf5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xdc, 0x7a, 0x20,
	0x33, 0xf4, 0xc0, 0x72, 0x41, 0x10, 0x19, 0x25, 0x11, 0x2e, 0x21, 0xa8, 0x19, 0x3e, 0x89, 0x25,
	0xa9, 0xc5, 0x25, 0x4e, 0x39, 0xf9, 0xc9, 0xd9, 0x4a, 0x16, 0x5c, 0xc2, 0x30, 0x93, 0x90, 0x84,
	0x41, 0xe6, 0x25, 0x81, 0x18, 0xa8, 0xe6, 0x81, 0xe5, 0x82, 0x20, 0x32, 0x49, 0x6c, 0x60, 0xa7,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x31, 0xb8, 0x23, 0xba, 0x00, 0x00, 0x00,
}
