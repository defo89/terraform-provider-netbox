// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/payments_account.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A payments account, which can be used to set up billing for an Ads customer.
type PaymentsAccount struct {
	// Immutable. The resource name of the payments account.
	// PaymentsAccount resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. A 16 digit ID used to identify a payments account.
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,2,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// Output only. The name of the payments account.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The currency code of the payments account.
	// A subset of the currency codes derived from the ISO 4217 standard is
	// supported.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// Output only. A 12 digit ID used to identify the payments profile associated with the
	// payments account.
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,5,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// Output only. A secondary payments profile ID present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *wrappers.StringValue `protobuf:"bytes,6,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3" json:"secondary_payments_profile_id,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}              `json:"-"`
	XXX_unrecognized           []byte                `json:"-"`
	XXX_sizecache              int32                 `json:"-"`
}

func (m *PaymentsAccount) Reset()         { *m = PaymentsAccount{} }
func (m *PaymentsAccount) String() string { return proto.CompactTextString(m) }
func (*PaymentsAccount) ProtoMessage()    {}
func (*PaymentsAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_154b57c389e0899a, []int{0}
}

func (m *PaymentsAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentsAccount.Unmarshal(m, b)
}
func (m *PaymentsAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentsAccount.Marshal(b, m, deterministic)
}
func (m *PaymentsAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentsAccount.Merge(m, src)
}
func (m *PaymentsAccount) XXX_Size() int {
	return xxx_messageInfo_PaymentsAccount.Size(m)
}
func (m *PaymentsAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentsAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentsAccount proto.InternalMessageInfo

func (m *PaymentsAccount) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *PaymentsAccount) GetPaymentsAccountId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountId
	}
	return nil
}

func (m *PaymentsAccount) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *PaymentsAccount) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *PaymentsAccount) GetPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileId
	}
	return nil
}

func (m *PaymentsAccount) GetSecondaryPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.SecondaryPaymentsProfileId
	}
	return nil
}

func init() {
	proto.RegisterType((*PaymentsAccount)(nil), "google.ads.googleads.v1.resources.PaymentsAccount")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/payments_account.proto", fileDescriptor_154b57c389e0899a)
}

var fileDescriptor_154b57c389e0899a = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6a, 0xd4, 0x40,
	0x1c, 0xc6, 0xd9, 0x6c, 0x5b, 0x70, 0x6c, 0x11, 0xa3, 0x87, 0x75, 0xa9, 0xda, 0x0a, 0x85, 0xc5,
	0xc3, 0x8c, 0x51, 0x84, 0x32, 0x9e, 0x66, 0x05, 0x4b, 0x3d, 0x48, 0x5c, 0x71, 0x0f, 0xb2, 0x10,
	0x66, 0x67, 0x66, 0x63, 0x20, 0x99, 0x7f, 0x98, 0x49, 0x56, 0x96, 0xd2, 0x83, 0xaf, 0xe2, 0xd1,
	0x47, 0xf1, 0x19, 0x3c, 0xf4, 0xdc, 0x47, 0xf0, 0x24, 0x9b, 0x64, 0xd2, 0x36, 0x88, 0xae, 0xb7,
	0x2f, 0xfc, 0xbf, 0xef, 0x37, 0xff, 0x2f, 0xc3, 0xa0, 0xe3, 0x18, 0x20, 0x4e, 0x15, 0xe1, 0xd2,
	0x92, 0x5a, 0xae, 0xd5, 0x32, 0x20, 0x46, 0x59, 0x28, 0x8d, 0x50, 0x96, 0xe4, 0x7c, 0x95, 0x29,
	0x5d, 0xd8, 0x88, 0x0b, 0x01, 0xa5, 0x2e, 0x70, 0x6e, 0xa0, 0x00, 0xff, 0xb0, 0xb6, 0x63, 0x2e,
	0x2d, 0x6e, 0x93, 0x78, 0x19, 0xe0, 0x36, 0x39, 0x7c, 0xec, 0xe0, 0x79, 0x42, 0x16, 0x89, 0x4a,
	0x65, 0x34, 0x57, 0x9f, 0xf9, 0x32, 0x01, 0x53, 0x33, 0x86, 0x0f, 0xae, 0x19, 0x5c, 0xac, 0x19,
	0x3d, 0x6a, 0x46, 0xd5, 0xd7, 0xbc, 0x5c, 0x90, 0x2f, 0x86, 0xe7, 0xb9, 0x32, 0xb6, 0x99, 0xef,
	0x5f, 0x8b, 0x72, 0xad, 0xa1, 0xe0, 0x45, 0x02, 0xba, 0x99, 0x3e, 0xf9, 0xb9, 0x85, 0xee, 0x84,
	0xcd, 0xde, 0xac, 0x5e, 0xdb, 0xff, 0x88, 0xf6, 0xdc, 0x19, 0x91, 0xe6, 0x99, 0x1a, 0xf4, 0x0e,
	0x7a, 0xa3, 0x5b, 0xe3, 0x67, 0x17, 0x6c, 0xfb, 0x17, 0x7b, 0x8a, 0x46, 0x57, 0x25, 0x1a, 0x95,
	0x27, 0x16, 0x0b, 0xc8, 0x48, 0x07, 0x34, 0xd9, 0x75, 0x98, 0x77, 0x3c, 0x53, 0xfe, 0x7b, 0x74,
	0xaf, 0xfb, 0x87, 0xa2, 0x44, 0x0e, 0xbc, 0x83, 0xde, 0xe8, 0xf6, 0xf3, 0xfd, 0x86, 0x85, 0x5d,
	0x0d, 0xfc, 0xa1, 0x30, 0x89, 0x8e, 0xa7, 0x3c, 0x2d, 0xd5, 0xb8, 0x7f, 0xc1, 0xfa, 0x93, 0xbb,
	0xf9, 0x4d, 0xfc, 0xa9, 0xf4, 0x5f, 0xa2, 0xad, 0x6a, 0xc1, 0xfe, 0xa6, 0x8c, 0xca, 0xee, 0xbf,
	0x41, 0x7b, 0xa2, 0x34, 0x46, 0x69, 0xb1, 0x8a, 0x04, 0x48, 0x35, 0xd8, 0xda, 0x34, 0xbf, 0xeb,
	0x72, 0xaf, 0x41, 0xde, 0x6c, 0x94, 0x1b, 0x58, 0x24, 0xa9, 0x5a, 0x37, 0xda, 0xfe, 0xef, 0x46,
	0x61, 0x1d, 0x3e, 0x95, 0xbe, 0x44, 0x0f, 0xad, 0x12, 0xa0, 0x25, 0x37, 0xab, 0xe8, 0x4f, 0xf0,
	0x9d, 0x4d, 0xe1, 0xc3, 0x96, 0x13, 0x76, 0x4f, 0xa1, 0xf1, 0x25, 0x93, 0x9b, 0xdf, 0xa3, 0x7f,
	0x2c, 0x4a, 0x5b, 0x40, 0xa6, 0x8c, 0x25, 0x67, 0x4e, 0x9e, 0x93, 0xce, 0x75, 0x58, 0x72, 0xd6,
	0xbd, 0xde, 0xf3, 0xf1, 0x57, 0x0f, 0x1d, 0x09, 0xc8, 0xf0, 0x3f, 0x9f, 0xc0, 0xf8, 0x7e, 0xe7,
	0xd0, 0x70, 0x5d, 0x2c, 0xec, 0x7d, 0x7a, 0xdb, 0x44, 0x63, 0x48, 0xb9, 0x8e, 0x31, 0x98, 0x98,
	0xc4, 0x4a, 0x57, 0xb5, 0xc9, 0xd5, 0xda, 0x7f, 0x79, 0x96, 0xaf, 0x5a, 0xf5, 0xcd, 0xeb, 0x9f,
	0x30, 0xf6, 0xdd, 0x3b, 0x3c, 0xa9, 0x91, 0x4c, 0x5a, 0x5c, 0xcb, 0xb5, 0x9a, 0x06, 0x78, 0xe2,
	0x9c, 0x3f, 0x9c, 0x67, 0xc6, 0xa4, 0x9d, 0xb5, 0x9e, 0xd9, 0x34, 0x98, 0xb5, 0x9e, 0x4b, 0xef,
	0xa8, 0x1e, 0x50, 0xca, 0xa4, 0xa5, 0xb4, 0x75, 0x51, 0x3a, 0x0d, 0x28, 0x6d, 0x7d, 0xf3, 0x9d,
	0x6a, 0xd9, 0x17, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x15, 0x90, 0x4e, 0xb1, 0x42, 0x04, 0x00,
	0x00,
}
