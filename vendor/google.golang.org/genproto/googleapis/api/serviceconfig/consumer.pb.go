// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/consumer.proto

package serviceconfig

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// Supported data type of the property values
type Property_PropertyType int32

const (
	// The type is unspecified, and will result in an error.
	Property_UNSPECIFIED Property_PropertyType = 0
	// The type is `int64`.
	Property_INT64 Property_PropertyType = 1
	// The type is `bool`.
	Property_BOOL Property_PropertyType = 2
	// The type is `string`.
	Property_STRING Property_PropertyType = 3
	// The type is 'double'.
	Property_DOUBLE Property_PropertyType = 4
)

var Property_PropertyType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "INT64",
	2: "BOOL",
	3: "STRING",
	4: "DOUBLE",
}

var Property_PropertyType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"INT64":       1,
	"BOOL":        2,
	"STRING":      3,
	"DOUBLE":      4,
}

func (x Property_PropertyType) String() string {
	return proto.EnumName(Property_PropertyType_name, int32(x))
}

func (Property_PropertyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bec5e69370b3104e, []int{1, 0}
}

// A descriptor for defining project properties for a service. One service may
// have many consumer projects, and the service may want to behave differently
// depending on some properties on the project. For example, a project may be
// associated with a school, or a business, or a government agency, a business
// type property on the project may affect how a service responds to the client.
// This descriptor defines which properties are allowed to be set on a project.
//
// Example:
//
//    project_properties:
//      properties:
//      - name: NO_WATERMARK
//        type: BOOL
//        description: Allows usage of the API without watermarks.
//      - name: EXTENDED_TILE_CACHE_PERIOD
//        type: INT64
type ProjectProperties struct {
	// List of per consumer project-specific properties.
	Properties           []*Property `protobuf:"bytes,1,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProjectProperties) Reset()         { *m = ProjectProperties{} }
func (m *ProjectProperties) String() string { return proto.CompactTextString(m) }
func (*ProjectProperties) ProtoMessage()    {}
func (*ProjectProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_bec5e69370b3104e, []int{0}
}

func (m *ProjectProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectProperties.Unmarshal(m, b)
}
func (m *ProjectProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectProperties.Marshal(b, m, deterministic)
}
func (m *ProjectProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectProperties.Merge(m, src)
}
func (m *ProjectProperties) XXX_Size() int {
	return xxx_messageInfo_ProjectProperties.Size(m)
}
func (m *ProjectProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectProperties.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectProperties proto.InternalMessageInfo

func (m *ProjectProperties) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

// Defines project properties.
//
// API services can define properties that can be assigned to consumer projects
// so that backends can perform response customization without having to make
// additional calls or maintain additional storage. For example, Maps API
// defines properties that controls map tile cache period, or whether to embed a
// watermark in a result.
//
// These values can be set via API producer console. Only API providers can
// define and set these properties.
type Property struct {
	// The name of the property (a.k.a key).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of this property.
	Type Property_PropertyType `protobuf:"varint,2,opt,name=type,proto3,enum=google.api.Property_PropertyType" json:"type,omitempty"`
	// The description of the property
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_bec5e69370b3104e, []int{1}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetType() Property_PropertyType {
	if m != nil {
		return m.Type
	}
	return Property_UNSPECIFIED
}

func (m *Property) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.api.Property_PropertyType", Property_PropertyType_name, Property_PropertyType_value)
	proto.RegisterType((*ProjectProperties)(nil), "google.api.ProjectProperties")
	proto.RegisterType((*Property)(nil), "google.api.Property")
}

func init() {
	proto.RegisterFile("google/api/consumer.proto", fileDescriptor_bec5e69370b3104e)
}

var fileDescriptor_bec5e69370b3104e = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0x85, 0xbe, 0x04, 0x06, 0xc5, 0xba, 0xf1, 0x50, 0x6f, 0x95, 0x13, 0xa7, 0x36,
	0x41, 0xf4, 0xe2, 0xad, 0x50, 0x4d, 0x13, 0x02, 0x4d, 0x81, 0x8b, 0xb7, 0x5a, 0xc7, 0x75, 0x0d,
	0xec, 0x6c, 0xb6, 0xd5, 0x84, 0x0f, 0xe8, 0xf7, 0x32, 0x2c, 0x88, 0x35, 0xf1, 0xf6, 0xcc, 0x3e,
	0x7f, 0xb2, 0xf9, 0xc1, 0xa5, 0x20, 0x12, 0x6b, 0x0c, 0x73, 0x2d, 0xc3, 0x82, 0x54, 0xf9, 0xbe,
	0x41, 0x13, 0x68, 0x43, 0x15, 0x71, 0xd8, 0x5b, 0x41, 0xae, 0x65, 0x3f, 0x81, 0xf3, 0xd4, 0xd0,
	0x1b, 0x16, 0x55, 0x6a, 0x48, 0xa3, 0xa9, 0x24, 0x96, 0x7c, 0x04, 0xa0, 0x8f, 0x97, 0xc7, 0xfc,
	0xe6, 0xa0, 0x3b, 0xbc, 0x08, 0x7e, 0x5a, 0xc1, 0x21, 0xbb, 0xcd, 0x6a, 0xb9, 0xfe, 0x27, 0x83,
	0xf6, 0xb7, 0xc1, 0x39, 0x38, 0x2a, 0xdf, 0xa0, 0xc7, 0x7c, 0x36, 0xe8, 0x64, 0x56, 0xf3, 0x1b,
	0x70, 0xaa, 0xad, 0x46, 0xaf, 0xe1, 0xb3, 0x41, 0x6f, 0x78, 0xf5, 0xd7, 0xe0, 0x51, 0x2c, 0xb7,
	0x1a, 0x33, 0x1b, 0xe7, 0x3e, 0x74, 0x9f, 0xb1, 0x2c, 0x8c, 0xd4, 0x95, 0x24, 0xe5, 0x35, 0xed,
	0x62, 0xfd, 0xa9, 0x3f, 0x85, 0x93, 0x7a, 0x8f, 0x9f, 0x41, 0x77, 0x35, 0x5b, 0xa4, 0xf1, 0x38,
	0xb9, 0x4f, 0xe2, 0x89, 0xfb, 0x8f, 0x77, 0xe0, 0x7f, 0x32, 0x5b, 0xde, 0x8e, 0x5c, 0xc6, 0xdb,
	0xe0, 0x44, 0xf3, 0xf9, 0xd4, 0x6d, 0x70, 0x80, 0xd6, 0x62, 0x99, 0x25, 0xb3, 0x07, 0xb7, 0xb9,
	0xd3, 0x93, 0xf9, 0x2a, 0x9a, 0xc6, 0xae, 0x13, 0xbd, 0x42, 0xaf, 0xa0, 0x4d, 0xed, 0x77, 0xd1,
	0xe9, 0xf8, 0x00, 0x30, 0xdd, 0xf1, 0x4b, 0xd9, 0x63, 0x7c, 0x30, 0x05, 0xad, 0x73, 0x25, 0x02,
	0x32, 0x22, 0x14, 0xa8, 0x2c, 0xdd, 0x70, 0x6f, 0xe5, 0x5a, 0x96, 0x96, 0x7d, 0x89, 0xe6, 0x43,
	0x16, 0x58, 0x90, 0x7a, 0x91, 0xe2, 0xee, 0xd7, 0xf5, 0xd4, 0xb2, 0x8d, 0xeb, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0xa4, 0x04, 0x2c, 0xac, 0x01, 0x00, 0x00,
}
