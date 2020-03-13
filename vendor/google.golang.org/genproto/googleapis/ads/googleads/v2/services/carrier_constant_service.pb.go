// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/carrier_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [CarrierConstantService.GetCarrierConstant][google.ads.googleads.v2.services.CarrierConstantService.GetCarrierConstant].
type GetCarrierConstantRequest struct {
	// Required. Resource name of the carrier constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCarrierConstantRequest) Reset()         { *m = GetCarrierConstantRequest{} }
func (m *GetCarrierConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetCarrierConstantRequest) ProtoMessage()    {}
func (*GetCarrierConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b7cfe7382d5c66, []int{0}
}

func (m *GetCarrierConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCarrierConstantRequest.Unmarshal(m, b)
}
func (m *GetCarrierConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCarrierConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetCarrierConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCarrierConstantRequest.Merge(m, src)
}
func (m *GetCarrierConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetCarrierConstantRequest.Size(m)
}
func (m *GetCarrierConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCarrierConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCarrierConstantRequest proto.InternalMessageInfo

func (m *GetCarrierConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCarrierConstantRequest)(nil), "google.ads.googleads.v2.services.GetCarrierConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/carrier_constant_service.proto", fileDescriptor_08b7cfe7382d5c66)
}

var fileDescriptor_08b7cfe7382d5c66 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x8a, 0xdb, 0x30,
	0x18, 0xc7, 0x0e, 0x14, 0x6a, 0xda, 0xc5, 0x43, 0x9b, 0x3a, 0x85, 0x86, 0x10, 0x4a, 0xe8, 0x20,
	0x81, 0xbb, 0x14, 0x85, 0xb6, 0x28, 0xa1, 0xa4, 0x53, 0x09, 0x29, 0x64, 0x28, 0x06, 0xa3, 0xd8,
	0xaa, 0x2b, 0xb0, 0xa5, 0x54, 0x72, 0xbc, 0x94, 0x2e, 0x79, 0x85, 0xbe, 0x41, 0xc7, 0x7b, 0x8f,
	0x5b, 0xb2, 0xde, 0x76, 0xd3, 0x0d, 0x37, 0xdd, 0x3b, 0x1c, 0x1c, 0x8e, 0x2c, 0xe7, 0xcf, 0xc5,
	0x64, 0xfb, 0xe1, 0xdf, 0xbf, 0xef, 0xfb, 0x64, 0xe7, 0x73, 0x22, 0x44, 0x92, 0x52, 0x48, 0x62,
	0x05, 0x35, 0x2c, 0x51, 0xe1, 0x43, 0x45, 0x65, 0xc1, 0x22, 0xaa, 0x60, 0x44, 0xa4, 0x64, 0x54,
	0x86, 0x91, 0xe0, 0x2a, 0x27, 0x3c, 0x0f, 0x2b, 0x06, 0x2c, 0xa5, 0xc8, 0x85, 0xdb, 0xd5, 0x2e,
	0x40, 0x62, 0x05, 0xea, 0x00, 0x50, 0xf8, 0xc0, 0x04, 0x78, 0x1f, 0x9a, 0x2a, 0x24, 0x55, 0x62,
	0x25, 0x4f, 0x75, 0xe8, 0x6c, 0xef, 0xb5, 0x71, 0x2e, 0x19, 0x24, 0x9c, 0x8b, 0x9c, 0xe4, 0x4c,
	0x70, 0x55, 0xb1, 0x2f, 0xf7, 0xd8, 0x28, 0x65, 0xb4, 0xb6, 0xbd, 0xd9, 0x23, 0x7e, 0x32, 0x9a,
	0xc6, 0xe1, 0x82, 0xfe, 0x22, 0x05, 0x13, 0x52, 0x0b, 0x7a, 0x5f, 0x9c, 0x57, 0x13, 0x9a, 0x8f,
	0x75, 0xe9, 0xb8, 0xea, 0x9c, 0xd1, 0xdf, 0x2b, 0xaa, 0x72, 0x77, 0xe0, 0x3c, 0x37, 0x83, 0x85,
	0x9c, 0x64, 0xb4, 0x6d, 0x75, 0xad, 0xc1, 0xd3, 0x51, 0xeb, 0x06, 0xdb, 0xb3, 0x67, 0x86, 0xf9,
	0x46, 0x32, 0xea, 0xdf, 0x5b, 0xce, 0x8b, 0xa3, 0x90, 0xef, 0x7a, 0x69, 0xf7, 0xd2, 0x72, 0xdc,
	0xc7, 0x15, 0xee, 0x10, 0x9c, 0xbb, 0x16, 0x68, 0x1c, 0xcc, 0xf3, 0x1b, 0xcd, 0xf5, 0x21, 0xc1,
	0x91, 0xb5, 0xf7, 0xe9, 0x1a, 0x1f, 0x6e, 0xb3, 0xbe, 0xba, 0xfd, 0x67, 0x0f, 0xdc, 0xb7, 0xe5,
	0xfd, 0xff, 0x1c, 0x30, 0x1f, 0xa3, 0x43, 0xaf, 0x82, 0xef, 0xfe, 0x7a, 0x9d, 0x0d, 0x6e, 0xef,
	0xaa, 0x2a, 0xb4, 0x64, 0x0a, 0x44, 0x22, 0x1b, 0xad, 0x6d, 0xa7, 0x1f, 0x89, 0xec, 0xec, 0x4e,
	0xa3, 0xce, 0xe9, 0x2b, 0x4d, 0xcb, 0xc7, 0x98, 0x5a, 0x3f, 0xbe, 0x56, 0x01, 0x89, 0x48, 0x09,
	0x4f, 0x80, 0x90, 0x09, 0x4c, 0x28, 0xdf, 0x3e, 0x15, 0xdc, 0x55, 0x36, 0xff, 0xa2, 0x43, 0x03,
	0xfe, 0xdb, 0xad, 0x09, 0xc6, 0x17, 0x76, 0x77, 0xa2, 0x03, 0x71, 0xac, 0x80, 0x86, 0x25, 0x9a,
	0xfb, 0xa0, 0x2a, 0x56, 0x1b, 0x23, 0x09, 0x70, 0xac, 0x82, 0x5a, 0x12, 0xcc, 0xfd, 0xc0, 0x48,
	0xee, 0xec, 0xbe, 0xfe, 0x8e, 0x10, 0x8e, 0x15, 0x42, 0xb5, 0x08, 0xa1, 0xb9, 0x8f, 0x90, 0x91,
	0x2d, 0x9e, 0x6c, 0xe7, 0x7c, 0xff, 0x10, 0x00, 0x00, 0xff, 0xff, 0x25, 0x64, 0x9a, 0x57, 0x49,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarrierConstantServiceClient is the client API for CarrierConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarrierConstantServiceClient interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error)
}

type carrierConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarrierConstantServiceClient(cc grpc.ClientConnInterface) CarrierConstantServiceClient {
	return &carrierConstantServiceClient{cc}
}

func (c *carrierConstantServiceClient) GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error) {
	out := new(resources.CarrierConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CarrierConstantService/GetCarrierConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarrierConstantServiceServer is the server API for CarrierConstantService service.
type CarrierConstantServiceServer interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(context.Context, *GetCarrierConstantRequest) (*resources.CarrierConstant, error)
}

// UnimplementedCarrierConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarrierConstantServiceServer struct {
}

func (*UnimplementedCarrierConstantServiceServer) GetCarrierConstant(ctx context.Context, req *GetCarrierConstantRequest) (*resources.CarrierConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarrierConstant not implemented")
}

func RegisterCarrierConstantServiceServer(s *grpc.Server, srv CarrierConstantServiceServer) {
	s.RegisterService(&_CarrierConstantService_serviceDesc, srv)
}

func _CarrierConstantService_GetCarrierConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarrierConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CarrierConstantService/GetCarrierConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, req.(*GetCarrierConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarrierConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CarrierConstantService",
	HandlerType: (*CarrierConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarrierConstant",
			Handler:    _CarrierConstantService_GetCarrierConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/carrier_constant_service.proto",
}
