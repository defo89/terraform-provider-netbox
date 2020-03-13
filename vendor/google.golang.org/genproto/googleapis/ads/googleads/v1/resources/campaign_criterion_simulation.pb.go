// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_criterion_simulation.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// A campaign criterion simulation. Supported combinations of advertising
// channel type, criterion ids, simulation type and simulation modification
// method is detailed below respectively.
//
// 1. SEARCH - 30000,30001,30002 - BID_MODIFIER - UNIFORM
// 2. SHOPPING - 30000,30001,30002 - BID_MODIFIER - UNIFORM
// 3. DISPLAY - 30001 - BID_MODIFIER - UNIFORM
type CampaignCriterionSimulation struct {
	// Immutable. The resource name of the campaign criterion simulation.
	// Campaign criterion simulation resource names have the form:
	//
	// `customers/{customer_id}/campaignCriterionSimulations/{campaign_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Campaign ID of the simulation.
	CampaignId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Output only. Criterion ID of the simulation.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// Output only. The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// Output only. How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v1.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// Output only. First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Output only. Last day on which the simulation is based, in YYYY-MM-DD format.
	EndDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*CampaignCriterionSimulation_BidModifierPointList
	PointList            isCampaignCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignCriterionSimulation) Reset()         { *m = CampaignCriterionSimulation{} }
func (m *CampaignCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionSimulation) ProtoMessage()    {}
func (*CampaignCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eae0e5c50ed969c, []int{0}
}

func (m *CampaignCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionSimulation.Unmarshal(m, b)
}
func (m *CampaignCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionSimulation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionSimulation.Merge(m, src)
}
func (m *CampaignCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionSimulation.Size(m)
}
func (m *CampaignCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionSimulation proto.InternalMessageInfo

func (m *CampaignCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterionSimulation) GetCampaignId() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isCampaignCriterionSimulation_PointList interface {
	isCampaignCriterionSimulation_PointList()
}

type CampaignCriterionSimulation_BidModifierPointList struct {
	BidModifierPointList *common.BidModifierSimulationPointList `protobuf:"bytes,8,opt,name=bid_modifier_point_list,json=bidModifierPointList,proto3,oneof"`
}

func (*CampaignCriterionSimulation_BidModifierPointList) isCampaignCriterionSimulation_PointList() {}

func (m *CampaignCriterionSimulation) GetPointList() isCampaignCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetBidModifierPointList() *common.BidModifierSimulationPointList {
	if x, ok := m.GetPointList().(*CampaignCriterionSimulation_BidModifierPointList); ok {
		return x.BidModifierPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionSimulation_BidModifierPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterionSimulation)(nil), "google.ads.googleads.v1.resources.CampaignCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_criterion_simulation.proto", fileDescriptor_4eae0e5c50ed969c)
}

var fileDescriptor_4eae0e5c50ed969c = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x27, 0xed, 0xfe, 0x7a, 0x83, 0x43, 0x40, 0x22, 0x6c, 0x13, 0x74, 0x48, 0x93, 0x76, 0x72,
	0xd4, 0x6d, 0xda, 0x21, 0x4c, 0x68, 0xc9, 0x98, 0xc6, 0x10, 0x43, 0xa3, 0x43, 0x95, 0x40, 0x95,
	0x22, 0x37, 0xf6, 0x32, 0x4b, 0x8d, 0x1d, 0xd9, 0xce, 0xd0, 0x04, 0x7b, 0x00, 0x0e, 0xbb, 0x20,
	0x9e, 0x80, 0x23, 0x8f, 0xc2, 0x53, 0xec, 0xbc, 0x47, 0xe0, 0x84, 0xea, 0x38, 0x49, 0xb5, 0xae,
	0xa5, 0xe2, 0xf6, 0x25, 0xdf, 0xef, 0x8f, 0xfd, 0xfb, 0x6c, 0x83, 0xfd, 0x98, 0xf3, 0xb8, 0x47,
	0x5c, 0x84, 0xa5, 0x9b, 0x97, 0xfd, 0xea, 0xbc, 0xe9, 0x0a, 0x22, 0x79, 0x26, 0x22, 0x22, 0xdd,
	0x08, 0x25, 0x29, 0xa2, 0x31, 0x0b, 0x23, 0x41, 0x15, 0x11, 0x94, 0xb3, 0x50, 0xd2, 0x24, 0xeb,
	0x21, 0x45, 0x39, 0x83, 0xa9, 0xe0, 0x8a, 0xdb, 0xab, 0x39, 0x17, 0x22, 0x2c, 0x61, 0x29, 0x03,
	0xcf, 0x9b, 0xb0, 0x94, 0x59, 0x72, 0x47, 0x39, 0x45, 0x3c, 0x49, 0x38, 0x73, 0x6f, 0x6b, 0x2e,
	0x05, 0xa3, 0x08, 0x84, 0x65, 0x89, 0x1c, 0xc0, 0x87, 0x09, 0xc7, 0xf4, 0x94, 0x46, 0xe6, 0x83,
	0xa8, 0x33, 0x8e, 0x8d, 0xc6, 0xe6, 0xc4, 0x1a, 0xea, 0x22, 0x25, 0x86, 0xf4, 0xac, 0x20, 0xa5,
	0xd4, 0x3d, 0xa5, 0xa4, 0x87, 0xc3, 0x2e, 0x39, 0x43, 0xe7, 0x94, 0x0b, 0x03, 0x78, 0x32, 0x00,
	0x28, 0x36, 0x68, 0x5a, 0x4f, 0x4d, 0x4b, 0x7f, 0x75, 0xb3, 0x53, 0xf7, 0xb3, 0x40, 0x69, 0x4a,
	0x84, 0x34, 0xfd, 0x95, 0x01, 0x2a, 0x62, 0x8c, 0x2b, 0xed, 0x6e, 0xba, 0xcf, 0x7f, 0xcc, 0x82,
	0xe5, 0x3d, 0x13, 0xf7, 0x5e, 0x91, 0xf6, 0x49, 0xb9, 0x48, 0x1b, 0x81, 0xfb, 0x85, 0x5f, 0xc8,
	0x50, 0x42, 0x1c, 0xab, 0x61, 0xad, 0xcf, 0x07, 0x3b, 0xd7, 0xfe, 0xf4, 0x1f, 0x7f, 0x1b, 0x6c,
	0x55, 0xd1, 0x9b, 0x2a, 0xa5, 0x12, 0x46, 0x3c, 0x71, 0xc7, 0x88, 0xb6, 0x16, 0x0b, 0xc9, 0x77,
	0x28, 0x21, 0xb6, 0x0f, 0x16, 0xca, 0x81, 0x53, 0xec, 0xd4, 0x1a, 0xd6, 0xfa, 0xc2, 0xc6, 0xb2,
	0xd1, 0x83, 0xc5, 0xb6, 0xe0, 0x21, 0x53, 0xdb, 0x5b, 0x6d, 0xd4, 0xcb, 0x48, 0x50, 0xbf, 0xf6,
	0xeb, 0x2d, 0x50, 0x90, 0x0e, 0xb1, 0xbd, 0x07, 0x16, 0xab, 0xa3, 0x42, 0xb1, 0x53, 0x9f, 0x50,
	0x63, 0xa1, 0x64, 0x1d, 0x62, 0xfb, 0x23, 0x98, 0xea, 0x8f, 0xc4, 0x99, 0x6a, 0x58, 0xeb, 0x0f,
	0x36, 0x76, 0xe1, 0xa8, 0x03, 0xa6, 0x07, 0x09, 0xab, 0xed, 0x7c, 0xb8, 0x48, 0xc9, 0x3e, 0xcb,
	0x92, 0x5b, 0xbf, 0x72, 0x07, 0x2d, 0x69, 0x7f, 0xb7, 0xc0, 0xc3, 0x3b, 0x8e, 0x8c, 0x33, 0xad,
	0xad, 0x3a, 0x13, 0x5b, 0x1d, 0x0d, 0x68, 0x1c, 0x69, 0x89, 0x5b, 0xc6, 0xc3, 0x80, 0x7c, 0x19,
	0x76, 0x32, 0xd4, 0xb0, 0x77, 0x01, 0x90, 0x0a, 0x09, 0x15, 0x62, 0xa4, 0x88, 0x33, 0xa3, 0x23,
	0x5b, 0x19, 0x8a, 0xec, 0x44, 0x09, 0xca, 0xe2, 0x81, 0xcc, 0xe6, 0x35, 0xe9, 0x15, 0x52, 0xc4,
	0xde, 0x01, 0x73, 0x84, 0xe1, 0x9c, 0x3f, 0x3b, 0x29, 0x7f, 0x96, 0x30, 0xac, 0xd9, 0x5f, 0xc1,
	0xe3, 0x2e, 0xc5, 0xe6, 0x2a, 0x11, 0x11, 0xa6, 0x9c, 0x32, 0x15, 0xf6, 0xa8, 0x54, 0xce, 0x9c,
	0x16, 0x7b, 0x39, 0x32, 0x97, 0xfc, 0x02, 0xc3, 0x80, 0xe2, 0x23, 0xc3, 0xae, 0x22, 0x38, 0xee,
	0xcb, 0xbc, 0xa5, 0x52, 0x69, 0xbb, 0xd7, 0xf7, 0x5a, 0x8f, 0xba, 0x15, 0xac, 0x6c, 0x7a, 0x57,
	0xd6, 0x8d, 0xff, 0xcd, 0xfa, 0xbf, 0x03, 0x6c, 0xbf, 0x8f, 0x32, 0xa9, 0x78, 0x42, 0x84, 0x74,
	0xbf, 0x14, 0xe5, 0x65, 0xf9, 0x6c, 0xdd, 0xc1, 0xe8, 0xe3, 0xc6, 0x3d, 0x6a, 0x97, 0xc1, 0x22,
	0x00, 0x55, 0x00, 0xc1, 0x55, 0x0d, 0xac, 0x45, 0x3c, 0x81, 0xff, 0x7c, 0xe4, 0x82, 0xc6, 0x98,
	0x75, 0x1e, 0xf7, 0x87, 0x70, 0x6c, 0x7d, 0x7a, 0x63, 0x64, 0x62, 0xde, 0x43, 0x2c, 0x86, 0x5c,
	0xc4, 0x6e, 0x4c, 0x98, 0x1e, 0x91, 0x5b, 0xed, 0x7a, 0xcc, 0x8b, 0xfc, 0xa2, 0xac, 0x7e, 0xd6,
	0xea, 0x07, 0xbe, 0xff, 0xab, 0xb6, 0x7a, 0x90, 0x4b, 0xfa, 0x58, 0xc2, 0xbc, 0xec, 0x57, 0xed,
	0x26, 0x6c, 0x15, 0xc8, 0xdf, 0x05, 0xa6, 0xe3, 0x63, 0xd9, 0x29, 0x31, 0x9d, 0x76, 0xb3, 0x53,
	0x62, 0x6e, 0x6a, 0x6b, 0x79, 0xc3, 0xf3, 0x7c, 0x2c, 0x3d, 0xaf, 0x44, 0x79, 0x5e, 0xbb, 0xe9,
	0x79, 0x25, 0xae, 0x3b, 0xa3, 0x17, 0xbb, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x02, 0xbb, 0xf5,
	0x9f, 0x3d, 0x06, 0x00, 0x00,
}
