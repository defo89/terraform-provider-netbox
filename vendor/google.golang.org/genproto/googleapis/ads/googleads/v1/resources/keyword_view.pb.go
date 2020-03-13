// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/keyword_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A keyword view.
type KeywordView struct {
	// Immutable. The resource name of the keyword view.
	// Keyword view resource names have the form:
	//
	// `customers/{customer_id}/keywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordView) Reset()         { *m = KeywordView{} }
func (m *KeywordView) String() string { return proto.CompactTextString(m) }
func (*KeywordView) ProtoMessage()    {}
func (*KeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c4b375a791b3e16, []int{0}
}

func (m *KeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordView.Unmarshal(m, b)
}
func (m *KeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordView.Marshal(b, m, deterministic)
}
func (m *KeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordView.Merge(m, src)
}
func (m *KeywordView) XXX_Size() int {
	return xxx_messageInfo_KeywordView.Size(m)
}
func (m *KeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordView proto.InternalMessageInfo

func (m *KeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*KeywordView)(nil), "google.ads.googleads.v1.resources.KeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/keyword_view.proto", fileDescriptor_4c4b375a791b3e16)
}

var fileDescriptor_4c4b375a791b3e16 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x2f, 0xf7, 0xc2, 0x8d, 0x0a, 0xd2, 0x95, 0x16, 0x41, 0x2b, 0x56, 0x5c, 0xc8,
	0x8c, 0x41, 0x57, 0xe3, 0x6a, 0xba, 0x29, 0x28, 0x48, 0xed, 0x22, 0x88, 0x06, 0xca, 0x34, 0x33,
	0xc6, 0xc1, 0x26, 0x7f, 0x99, 0x49, 0x53, 0xa4, 0xf4, 0x65, 0x5c, 0xfa, 0x0e, 0xbe, 0x80, 0x4f,
	0xe1, 0xba, 0x8f, 0x20, 0x08, 0x92, 0x4e, 0x67, 0x9a, 0x95, 0xba, 0x3b, 0xf0, 0x7f, 0xe7, 0xe4,
	0xe4, 0x8c, 0x7f, 0x96, 0x00, 0x24, 0x43, 0x81, 0x19, 0xd7, 0xd8, 0xc8, 0x52, 0x15, 0x01, 0x56,
	0x42, 0xc3, 0x58, 0xc5, 0x42, 0xe3, 0x47, 0xf1, 0x34, 0x01, 0xc5, 0xfb, 0x85, 0x14, 0x13, 0x34,
	0x52, 0x90, 0x43, 0xbd, 0x69, 0x50, 0xc4, 0xb8, 0x46, 0xce, 0x85, 0x8a, 0x00, 0x39, 0x57, 0x63,
	0xd7, 0x06, 0x8f, 0x24, 0xbe, 0x97, 0x62, 0xc8, 0xfb, 0x03, 0xf1, 0xc0, 0x0a, 0x09, 0xca, 0x64,
	0x34, 0xb6, 0x2b, 0x80, 0xb5, 0x2d, 0x4f, 0x3b, 0x95, 0x13, 0xcb, 0x32, 0xc8, 0x59, 0x2e, 0x21,
	0xd3, 0xe6, 0xba, 0xff, 0xea, 0xf9, 0x6b, 0x97, 0xa6, 0x53, 0x28, 0xc5, 0xa4, 0x7e, 0xed, 0x6f,
	0x58, 0x7f, 0x3f, 0x63, 0xa9, 0xd8, 0xf2, 0xf6, 0xbc, 0xa3, 0xff, 0xed, 0xe3, 0x77, 0xfa, 0xf7,
	0x83, 0x1e, 0xfa, 0x07, 0xab, 0x82, 0x4b, 0x35, 0x92, 0x1a, 0xc5, 0x90, 0xe2, 0x4a, 0x48, 0x6f,
	0xdd, 0x46, 0x5c, 0xb1, 0x54, 0x90, 0xbb, 0x39, 0xbd, 0xf9, 0x9d, 0xb1, 0x7e, 0x12, 0x8f, 0x75,
	0x0e, 0xa9, 0x50, 0x1a, 0x4f, 0xad, 0x9c, 0xd9, 0xcd, 0x4a, 0x42, 0xe3, 0x69, 0x75, 0xc1, 0x59,
	0xfb, 0xd3, 0xf3, 0x5b, 0x31, 0xa4, 0xe8, 0xc7, 0x0d, 0xdb, 0x9b, 0x95, 0x0f, 0x75, 0xcb, 0x7f,
	0xef, 0x7a, 0xb7, 0x17, 0x4b, 0x5b, 0x02, 0x43, 0x96, 0x25, 0x08, 0x54, 0x82, 0x13, 0x91, 0x2d,
	0x96, 0xc1, 0xab, 0x9a, 0xdf, 0xbc, 0xe7, 0xb9, 0x53, 0xcf, 0xb5, 0x3f, 0x1d, 0x4a, 0x5f, 0x6a,
	0xcd, 0x8e, 0x89, 0xa4, 0x5c, 0x23, 0x23, 0x4b, 0x15, 0x06, 0xa8, 0x67, 0xc9, 0x37, 0xcb, 0x44,
	0x94, 0xeb, 0xc8, 0x31, 0x51, 0x18, 0x44, 0x8e, 0x99, 0xd7, 0x5a, 0xe6, 0x40, 0x08, 0xe5, 0x9a,
	0x10, 0x47, 0x11, 0x12, 0x06, 0x84, 0x38, 0x6e, 0xf0, 0x6f, 0x51, 0xf6, 0xf4, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0x3c, 0xbf, 0x72, 0x7b, 0x02, 0x00, 0x00,
}
