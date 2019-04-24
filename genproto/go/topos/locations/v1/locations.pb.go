// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topos/locations/v1/locations.proto

package locations

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type RegionType struct {
	// The name of the region type. It must have the format
	// `"regionTypes/{region_type}"`. `{region_type}` must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[a-z\d]+(-[a-z\d]+)*$`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionType) Reset()         { *m = RegionType{} }
func (m *RegionType) String() string { return proto.CompactTextString(m) }
func (*RegionType) ProtoMessage()    {}
func (*RegionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{0}
}

func (m *RegionType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionType.Unmarshal(m, b)
}
func (m *RegionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionType.Marshal(b, m, deterministic)
}
func (m *RegionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionType.Merge(m, src)
}
func (m *RegionType) XXX_Size() int {
	return xxx_messageInfo_RegionType.Size(m)
}
func (m *RegionType) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionType.DiscardUnknown(m)
}

var xxx_messageInfo_RegionType proto.InternalMessageInfo

func (m *RegionType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Region struct {
	// The name of the region. It must have the format
	// `"regionTypes/{region_type}/regions/{region}"`.
	// `{region}` must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[a-z\d]+(-[a-z\d]+)*$`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The WKB-encoded region geometry.
	Geometry             []byte   `protobuf:"bytes,2,opt,name=geometry,proto3" json:"geometry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Region) Reset()         { *m = Region{} }
func (m *Region) String() string { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()    {}
func (*Region) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{1}
}

func (m *Region) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Region.Unmarshal(m, b)
}
func (m *Region) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Region.Marshal(b, m, deterministic)
}
func (m *Region) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Region.Merge(m, src)
}
func (m *Region) XXX_Size() int {
	return xxx_messageInfo_Region.Size(m)
}
func (m *Region) XXX_DiscardUnknown() {
	xxx_messageInfo_Region.DiscardUnknown(m)
}

var xxx_messageInfo_Region proto.InternalMessageInfo

func (m *Region) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Region) GetGeometry() []byte {
	if m != nil {
		return m.Geometry
	}
	return nil
}

type GetRegionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegionRequest) Reset()         { *m = GetRegionRequest{} }
func (m *GetRegionRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegionRequest) ProtoMessage()    {}
func (*GetRegionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{2}
}

func (m *GetRegionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegionRequest.Unmarshal(m, b)
}
func (m *GetRegionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegionRequest.Marshal(b, m, deterministic)
}
func (m *GetRegionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegionRequest.Merge(m, src)
}
func (m *GetRegionRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegionRequest.Size(m)
}
func (m *GetRegionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegionRequest proto.InternalMessageInfo

func (m *GetRegionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SearchRegionsRequest struct {
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous Search request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Return regions of the given type.
	RegionType string `protobuf:"bytes,3,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	// Return regions that intersect the given WKB-encoded geometry.
	IntersectingGeometry []byte `protobuf:"bytes,4,opt,name=intersecting_geometry,json=intersectingGeometry,proto3" json:"intersecting_geometry,omitempty"`
	// Do not include region geometries in the response.
	ExcludeGeometry bool `protobuf:"varint,5,opt,name=exclude_geometry,json=excludeGeometry,proto3" json:"exclude_geometry,omitempty"`
	// Return regions included by the given region.
	IncludedByRegion     string   `protobuf:"bytes,6,opt,name=included_by_region,json=includedByRegion,proto3" json:"included_by_region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRegionsRequest) Reset()         { *m = SearchRegionsRequest{} }
func (m *SearchRegionsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRegionsRequest) ProtoMessage()    {}
func (*SearchRegionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{3}
}

func (m *SearchRegionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRegionsRequest.Unmarshal(m, b)
}
func (m *SearchRegionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRegionsRequest.Marshal(b, m, deterministic)
}
func (m *SearchRegionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRegionsRequest.Merge(m, src)
}
func (m *SearchRegionsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRegionsRequest.Size(m)
}
func (m *SearchRegionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRegionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRegionsRequest proto.InternalMessageInfo

func (m *SearchRegionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchRegionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *SearchRegionsRequest) GetRegionType() string {
	if m != nil {
		return m.RegionType
	}
	return ""
}

func (m *SearchRegionsRequest) GetIntersectingGeometry() []byte {
	if m != nil {
		return m.IntersectingGeometry
	}
	return nil
}

func (m *SearchRegionsRequest) GetExcludeGeometry() bool {
	if m != nil {
		return m.ExcludeGeometry
	}
	return false
}

func (m *SearchRegionsRequest) GetIncludedByRegion() string {
	if m != nil {
		return m.IncludedByRegion
	}
	return ""
}

type SearchRegionsResponse struct {
	Regions []*Region `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the search.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRegionsResponse) Reset()         { *m = SearchRegionsResponse{} }
func (m *SearchRegionsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchRegionsResponse) ProtoMessage()    {}
func (*SearchRegionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{4}
}

func (m *SearchRegionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRegionsResponse.Unmarshal(m, b)
}
func (m *SearchRegionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRegionsResponse.Marshal(b, m, deterministic)
}
func (m *SearchRegionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRegionsResponse.Merge(m, src)
}
func (m *SearchRegionsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchRegionsResponse.Size(m)
}
func (m *SearchRegionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRegionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRegionsResponse proto.InternalMessageInfo

func (m *SearchRegionsResponse) GetRegions() []*Region {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *SearchRegionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type GetRegionFeatureSetValuesRequest struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	FeatureSet           string   `protobuf:"bytes,2,opt,name=feature_set,json=featureSet,proto3" json:"feature_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRegionFeatureSetValuesRequest) Reset()         { *m = GetRegionFeatureSetValuesRequest{} }
func (m *GetRegionFeatureSetValuesRequest) String() string { return proto.CompactTextString(m) }
func (*GetRegionFeatureSetValuesRequest) ProtoMessage()    {}
func (*GetRegionFeatureSetValuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{5}
}

func (m *GetRegionFeatureSetValuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegionFeatureSetValuesRequest.Unmarshal(m, b)
}
func (m *GetRegionFeatureSetValuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegionFeatureSetValuesRequest.Marshal(b, m, deterministic)
}
func (m *GetRegionFeatureSetValuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegionFeatureSetValuesRequest.Merge(m, src)
}
func (m *GetRegionFeatureSetValuesRequest) XXX_Size() int {
	return xxx_messageInfo_GetRegionFeatureSetValuesRequest.Size(m)
}
func (m *GetRegionFeatureSetValuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegionFeatureSetValuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegionFeatureSetValuesRequest proto.InternalMessageInfo

func (m *GetRegionFeatureSetValuesRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *GetRegionFeatureSetValuesRequest) GetFeatureSet() string {
	if m != nil {
		return m.FeatureSet
	}
	return ""
}

type GetRegionFeatureSetValuesResponse struct {
	FeatureValues        map[string]float64 `protobuf:"bytes,1,rep,name=feature_values,json=featureValues,proto3" json:"feature_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetRegionFeatureSetValuesResponse) Reset()         { *m = GetRegionFeatureSetValuesResponse{} }
func (m *GetRegionFeatureSetValuesResponse) String() string { return proto.CompactTextString(m) }
func (*GetRegionFeatureSetValuesResponse) ProtoMessage()    {}
func (*GetRegionFeatureSetValuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{6}
}

func (m *GetRegionFeatureSetValuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRegionFeatureSetValuesResponse.Unmarshal(m, b)
}
func (m *GetRegionFeatureSetValuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRegionFeatureSetValuesResponse.Marshal(b, m, deterministic)
}
func (m *GetRegionFeatureSetValuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRegionFeatureSetValuesResponse.Merge(m, src)
}
func (m *GetRegionFeatureSetValuesResponse) XXX_Size() int {
	return xxx_messageInfo_GetRegionFeatureSetValuesResponse.Size(m)
}
func (m *GetRegionFeatureSetValuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRegionFeatureSetValuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRegionFeatureSetValuesResponse proto.InternalMessageInfo

func (m *GetRegionFeatureSetValuesResponse) GetFeatureValues() map[string]float64 {
	if m != nil {
		return m.FeatureValues
	}
	return nil
}

type LatLng struct {
	Latitude             float64  `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LatLng) Reset()         { *m = LatLng{} }
func (m *LatLng) String() string { return proto.CompactTextString(m) }
func (*LatLng) ProtoMessage()    {}
func (*LatLng) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{7}
}

func (m *LatLng) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatLng.Unmarshal(m, b)
}
func (m *LatLng) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatLng.Marshal(b, m, deterministic)
}
func (m *LatLng) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatLng.Merge(m, src)
}
func (m *LatLng) XXX_Size() int {
	return xxx_messageInfo_LatLng.Size(m)
}
func (m *LatLng) XXX_DiscardUnknown() {
	xxx_messageInfo_LatLng.DiscardUnknown(m)
}

var xxx_messageInfo_LatLng proto.InternalMessageInfo

func (m *LatLng) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *LatLng) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type LocateRegionsRequest struct {
	// Return regions of the given type.
	RegionType string `protobuf:"bytes,1,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	// Return region names that include the given location.
	Location             *LatLng  `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocateRegionsRequest) Reset()         { *m = LocateRegionsRequest{} }
func (m *LocateRegionsRequest) String() string { return proto.CompactTextString(m) }
func (*LocateRegionsRequest) ProtoMessage()    {}
func (*LocateRegionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{8}
}

func (m *LocateRegionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocateRegionsRequest.Unmarshal(m, b)
}
func (m *LocateRegionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocateRegionsRequest.Marshal(b, m, deterministic)
}
func (m *LocateRegionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocateRegionsRequest.Merge(m, src)
}
func (m *LocateRegionsRequest) XXX_Size() int {
	return xxx_messageInfo_LocateRegionsRequest.Size(m)
}
func (m *LocateRegionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LocateRegionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LocateRegionsRequest proto.InternalMessageInfo

func (m *LocateRegionsRequest) GetRegionType() string {
	if m != nil {
		return m.RegionType
	}
	return ""
}

func (m *LocateRegionsRequest) GetLocation() *LatLng {
	if m != nil {
		return m.Location
	}
	return nil
}

type LocateRegionsResponse struct {
	// The located region names.
	Regions              []string `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocateRegionsResponse) Reset()         { *m = LocateRegionsResponse{} }
func (m *LocateRegionsResponse) String() string { return proto.CompactTextString(m) }
func (*LocateRegionsResponse) ProtoMessage()    {}
func (*LocateRegionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_954338d7542bc8de, []int{9}
}

func (m *LocateRegionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocateRegionsResponse.Unmarshal(m, b)
}
func (m *LocateRegionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocateRegionsResponse.Marshal(b, m, deterministic)
}
func (m *LocateRegionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocateRegionsResponse.Merge(m, src)
}
func (m *LocateRegionsResponse) XXX_Size() int {
	return xxx_messageInfo_LocateRegionsResponse.Size(m)
}
func (m *LocateRegionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LocateRegionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LocateRegionsResponse proto.InternalMessageInfo

func (m *LocateRegionsResponse) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func init() {
	proto.RegisterType((*RegionType)(nil), "topos.locations.v1.RegionType")
	proto.RegisterType((*Region)(nil), "topos.locations.v1.Region")
	proto.RegisterType((*GetRegionRequest)(nil), "topos.locations.v1.GetRegionRequest")
	proto.RegisterType((*SearchRegionsRequest)(nil), "topos.locations.v1.SearchRegionsRequest")
	proto.RegisterType((*SearchRegionsResponse)(nil), "topos.locations.v1.SearchRegionsResponse")
	proto.RegisterType((*GetRegionFeatureSetValuesRequest)(nil), "topos.locations.v1.GetRegionFeatureSetValuesRequest")
	proto.RegisterType((*GetRegionFeatureSetValuesResponse)(nil), "topos.locations.v1.GetRegionFeatureSetValuesResponse")
	proto.RegisterMapType((map[string]float64)(nil), "topos.locations.v1.GetRegionFeatureSetValuesResponse.FeatureValuesEntry")
	proto.RegisterType((*LatLng)(nil), "topos.locations.v1.LatLng")
	proto.RegisterType((*LocateRegionsRequest)(nil), "topos.locations.v1.LocateRegionsRequest")
	proto.RegisterType((*LocateRegionsResponse)(nil), "topos.locations.v1.LocateRegionsResponse")
}

func init() { proto.RegisterFile("topos/locations/v1/locations.proto", fileDescriptor_954338d7542bc8de) }

var fileDescriptor_954338d7542bc8de = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0x4e, 0x17, 0x76, 0xdd, 0x3d, 0x88, 0x6c, 0x26, 0x8b, 0x59, 0x57, 0xd4, 0x75, 0x42, 0xc8,
	0x42, 0xcc, 0x56, 0x7e, 0x24, 0x04, 0x43, 0x62, 0x48, 0x04, 0x2e, 0x36, 0xc6, 0x14, 0xe2, 0x85,
	0x5e, 0x34, 0x65, 0x39, 0xd4, 0x86, 0x65, 0xa6, 0xb6, 0xb3, 0x84, 0x42, 0xb8, 0xd0, 0x57, 0xf0,
	0xd6, 0xc4, 0x87, 0x32, 0x31, 0x3e, 0x80, 0x0f, 0x62, 0x3a, 0x33, 0x6d, 0xf7, 0xa7, 0xa0, 0x78,
	0xd7, 0xf3, 0x33, 0xdf, 0x7c, 0xdf, 0x39, 0x5f, 0x5b, 0xa0, 0x82, 0xfb, 0x3c, 0x34, 0x7b, 0xbc,
	0xeb, 0x08, 0x8f, 0xb3, 0xd0, 0x3c, 0x5b, 0xce, 0x82, 0xb6, 0x1f, 0x70, 0xc1, 0x09, 0x91, 0x3d,
	0xed, 0x2c, 0x7d, 0xb6, 0xdc, 0x98, 0x73, 0x39, 0x77, 0x7b, 0x68, 0x3a, 0xbe, 0x67, 0x3a, 0x8c,
	0x71, 0x31, 0x78, 0x82, 0x36, 0x01, 0x2c, 0x74, 0x3d, 0xce, 0x0e, 0x22, 0x1f, 0x09, 0x81, 0x49,
	0xe6, 0x9c, 0x62, 0xdd, 0x68, 0x1a, 0xad, 0x8a, 0x25, 0x9f, 0xe9, 0x06, 0x94, 0x54, 0x47, 0x5e,
	0x95, 0x34, 0xa0, 0xec, 0x22, 0x3f, 0x45, 0x11, 0x44, 0xf5, 0x42, 0xd3, 0x68, 0xdd, 0xb5, 0xd2,
	0x98, 0x2e, 0x40, 0x75, 0x17, 0x85, 0x3a, 0x6c, 0xe1, 0xa7, 0x3e, 0x86, 0x22, 0xf7, 0x86, 0xcf,
	0x05, 0xa8, 0xed, 0xa3, 0x13, 0x74, 0x3f, 0xaa, 0xde, 0x30, 0x69, 0x7e, 0x08, 0x15, 0xdf, 0x71,
	0xd1, 0x0e, 0xbd, 0x0b, 0x75, 0xa2, 0x68, 0x95, 0xe3, 0xc4, 0xbe, 0x77, 0x81, 0xe4, 0x11, 0x80,
	0x2c, 0x0a, 0x7e, 0x82, 0x4c, 0xde, 0x5d, 0xb1, 0x64, 0xfb, 0x41, 0x9c, 0x20, 0x4f, 0x60, 0x2a,
	0x90, 0x68, 0xb6, 0x88, 0x7c, 0xac, 0x4f, 0xc8, 0x3a, 0x04, 0x99, 0xd6, 0x55, 0x98, 0xf5, 0x98,
	0xc0, 0x20, 0xc4, 0xae, 0xf0, 0x98, 0x6b, 0xa7, 0x32, 0x26, 0xa5, 0x8c, 0xda, 0x60, 0x71, 0x57,
	0xd7, 0xc8, 0x22, 0x54, 0xf1, 0xbc, 0xdb, 0xeb, 0x1f, 0x61, 0xd6, 0x5f, 0x6c, 0x1a, 0xad, 0xb2,
	0x35, 0xa3, 0xf3, 0x69, 0xeb, 0x33, 0x20, 0x1e, 0x93, 0xa9, 0x23, 0xfb, 0x30, 0xb2, 0xd5, 0xcd,
	0xf5, 0x92, 0xe4, 0x51, 0x4d, 0x2a, 0xdb, 0x91, 0x92, 0x4c, 0xfb, 0x30, 0x3b, 0x32, 0x82, 0xd0,
	0xe7, 0x2c, 0x44, 0xb2, 0x06, 0x77, 0xd4, 0xd1, 0xb0, 0x6e, 0x34, 0x27, 0x5a, 0x53, 0x2b, 0x8d,
	0xf6, 0xf8, 0x92, 0xdb, 0x7a, 0xc8, 0x49, 0x2b, 0x59, 0x80, 0x19, 0x86, 0xe7, 0xc2, 0x1e, 0x9b,
	0xd0, 0x74, 0x9c, 0x7e, 0x9b, 0x4c, 0x89, 0x7e, 0x80, 0x66, 0xba, 0xa2, 0x1d, 0x74, 0x44, 0x3f,
	0xc0, 0x7d, 0x14, 0xef, 0x9c, 0x5e, 0x1f, 0xd3, 0x2d, 0xdc, 0x87, 0x92, 0x26, 0xaf, 0x96, 0xa6,
	0xa3, 0x78, 0xc2, 0xc7, 0xea, 0x88, 0x1d, 0xa2, 0xd0, 0xf8, 0x70, 0x9c, 0xa2, 0xd0, 0x9f, 0x06,
	0x3c, 0xbd, 0x01, 0x5d, 0x0b, 0xe4, 0x70, 0x2f, 0x81, 0x39, 0x93, 0x15, 0xad, 0x73, 0x2f, 0x4f,
	0xe7, 0x5f, 0xe1, 0xda, 0xba, 0xa0, 0xb2, 0xaf, 0x99, 0x08, 0x22, 0x6b, 0xfa, 0x78, 0x30, 0xd7,
	0x78, 0x05, 0x64, 0xbc, 0x89, 0x54, 0x61, 0xe2, 0x04, 0x23, 0x2d, 0x31, 0x7e, 0x24, 0x35, 0x28,
	0x4a, 0x42, 0x52, 0x99, 0x61, 0xa9, 0x60, 0xb3, 0xb0, 0x61, 0xd0, 0x6d, 0x28, 0x75, 0x1c, 0xd1,
	0x61, 0x6e, 0x6c, 0xff, 0x9e, 0x23, 0x3c, 0xd1, 0x3f, 0x52, 0x06, 0x35, 0xac, 0x34, 0x26, 0x73,
	0x50, 0xe9, 0x71, 0xe6, 0xaa, 0xa2, 0xc2, 0xc8, 0x12, 0x94, 0x43, 0xad, 0x13, 0x2b, 0xc3, 0x11,
	0xcf, 0x8f, 0xf8, 0xd6, 0x18, 0xf3, 0xed, 0x3a, 0x94, 0x93, 0x91, 0x48, 0xd4, 0x6b, 0x1c, 0xa1,
	0x08, 0x5a, 0x69, 0x2f, 0x5d, 0x86, 0xd9, 0x91, 0x0b, 0xf5, 0x02, 0xea, 0xc3, 0x0e, 0xab, 0xa4,
	0x2e, 0x5a, 0xf9, 0x5e, 0x84, 0x4a, 0x27, 0x01, 0x25, 0x11, 0x54, 0xd2, 0xf1, 0x93, 0xf9, 0x1b,
	0xb7, 0xa3, 0xc5, 0x34, 0x6e, 0xf0, 0x2a, 0x5d, 0xfa, 0xf2, 0xe3, 0xf7, 0xd7, 0xc2, 0x3c, 0xa1,
	0xf1, 0x87, 0xec, 0x32, 0xfe, 0x10, 0x6c, 0x65, 0x2a, 0x43, 0x73, 0xc9, 0xd4, 0x2c, 0xcc, 0xa5,
	0x2b, 0xf2, 0xcb, 0x80, 0x07, 0xd7, 0xae, 0x9e, 0xac, 0xdd, 0xd2, 0x29, 0x8a, 0xdb, 0x8b, 0xff,
	0xf2, 0x17, 0x7d, 0x23, 0x69, 0xef, 0x91, 0x1d, 0x49, 0x5b, 0x71, 0xbc, 0x96, 0xb8, 0x79, 0x39,
	0xf0, 0x86, 0x6c, 0x65, 0xaf, 0x46, 0x5c, 0x52, 0x66, 0x27, 0xdf, 0x0c, 0x98, 0x1e, 0x7a, 0xf3,
	0x49, 0x2b, 0x8f, 0x58, 0xde, 0xf7, 0xb1, 0xb1, 0xf8, 0x0f, 0x9d, 0x9a, 0xf6, 0xba, 0xa4, 0xfd,
	0x9c, 0xb6, 0x07, 0x68, 0x4b, 0x87, 0x0d, 0x73, 0xbf, 0x4a, 0xc8, 0x6f, 0x86, 0x12, 0x46, 0xd2,
	0x1b, 0xb2, 0x4d, 0x3e, 0xbd, 0x3c, 0x2b, 0xe7, 0xd3, 0xcb, 0xf5, 0xe0, 0xed, 0xe9, 0x49, 0x4c,
	0xdc, 0x7e, 0xfc, 0x7e, 0x6e, 0xfc, 0xb7, 0xf8, 0x32, 0x0d, 0x0e, 0x4b, 0xf2, 0x2f, 0xb7, 0xfa,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x41, 0x0e, 0x88, 0x76, 0x3d, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LocationsClient is the client API for Locations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationsClient interface {
	// Gets a region.
	GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*Region, error)
	// Region feature values by feature.
	GetRegionFeatureSetValues(ctx context.Context, in *GetRegionFeatureSetValuesRequest, opts ...grpc.CallOption) (*GetRegionFeatureSetValuesResponse, error)
	// Search regions.
	SearchRegions(ctx context.Context, in *SearchRegionsRequest, opts ...grpc.CallOption) (*SearchRegionsResponse, error)
	// Locate regions.
	LocateRegions(ctx context.Context, in *LocateRegionsRequest, opts ...grpc.CallOption) (*LocateRegionsResponse, error)
}

type locationsClient struct {
	cc *grpc.ClientConn
}

func NewLocationsClient(cc *grpc.ClientConn) LocationsClient {
	return &locationsClient{cc}
}

func (c *locationsClient) GetRegion(ctx context.Context, in *GetRegionRequest, opts ...grpc.CallOption) (*Region, error) {
	out := new(Region)
	err := c.cc.Invoke(ctx, "/topos.locations.v1.Locations/GetRegion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) GetRegionFeatureSetValues(ctx context.Context, in *GetRegionFeatureSetValuesRequest, opts ...grpc.CallOption) (*GetRegionFeatureSetValuesResponse, error) {
	out := new(GetRegionFeatureSetValuesResponse)
	err := c.cc.Invoke(ctx, "/topos.locations.v1.Locations/GetRegionFeatureSetValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) SearchRegions(ctx context.Context, in *SearchRegionsRequest, opts ...grpc.CallOption) (*SearchRegionsResponse, error) {
	out := new(SearchRegionsResponse)
	err := c.cc.Invoke(ctx, "/topos.locations.v1.Locations/SearchRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) LocateRegions(ctx context.Context, in *LocateRegionsRequest, opts ...grpc.CallOption) (*LocateRegionsResponse, error) {
	out := new(LocateRegionsResponse)
	err := c.cc.Invoke(ctx, "/topos.locations.v1.Locations/LocateRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationsServer is the server API for Locations service.
type LocationsServer interface {
	// Gets a region.
	GetRegion(context.Context, *GetRegionRequest) (*Region, error)
	// Region feature values by feature.
	GetRegionFeatureSetValues(context.Context, *GetRegionFeatureSetValuesRequest) (*GetRegionFeatureSetValuesResponse, error)
	// Search regions.
	SearchRegions(context.Context, *SearchRegionsRequest) (*SearchRegionsResponse, error)
	// Locate regions.
	LocateRegions(context.Context, *LocateRegionsRequest) (*LocateRegionsResponse, error)
}

func RegisterLocationsServer(s *grpc.Server, srv LocationsServer) {
	s.RegisterService(&_Locations_serviceDesc, srv)
}

func _Locations_GetRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).GetRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.locations.v1.Locations/GetRegion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).GetRegion(ctx, req.(*GetRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_GetRegionFeatureSetValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegionFeatureSetValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).GetRegionFeatureSetValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.locations.v1.Locations/GetRegionFeatureSetValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).GetRegionFeatureSetValues(ctx, req.(*GetRegionFeatureSetValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_SearchRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).SearchRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.locations.v1.Locations/SearchRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).SearchRegions(ctx, req.(*SearchRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_LocateRegions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateRegionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).LocateRegions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.locations.v1.Locations/LocateRegions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).LocateRegions(ctx, req.(*LocateRegionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Locations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topos.locations.v1.Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegion",
			Handler:    _Locations_GetRegion_Handler,
		},
		{
			MethodName: "GetRegionFeatureSetValues",
			Handler:    _Locations_GetRegionFeatureSetValues_Handler,
		},
		{
			MethodName: "SearchRegions",
			Handler:    _Locations_SearchRegions_Handler,
		},
		{
			MethodName: "LocateRegions",
			Handler:    _Locations_LocateRegions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topos/locations/v1/locations.proto",
}
