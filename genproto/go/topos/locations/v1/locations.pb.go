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
	ExcludeGeometry      bool     `protobuf:"varint,5,opt,name=exclude_geometry,json=excludeGeometry,proto3" json:"exclude_geometry,omitempty"`
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
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xcd, 0x00, 0xbb, 0x6e, 0x2f, 0x22, 0x64, 0xb2, 0x98, 0x75, 0x45, 0x5d, 0x27, 0x84, 0x2c,
	0x3c, 0x6c, 0xe5, 0x43, 0x42, 0x30, 0x24, 0x86, 0x44, 0xe0, 0x61, 0x63, 0x4c, 0x21, 0x3e, 0xe8,
	0xc3, 0xa6, 0x2e, 0x97, 0xda, 0xb0, 0xcc, 0xd4, 0x76, 0x96, 0xb0, 0x10, 0x5e, 0xfc, 0x0b, 0xbe,
	0x9a, 0xf8, 0xa3, 0x48, 0x8c, 0x3f, 0xc0, 0x1f, 0x62, 0x3a, 0x33, 0x6d, 0xf7, 0xa3, 0xac, 0xe2,
	0x5b, 0xe7, 0xde, 0x33, 0xa7, 0xe7, 0xdc, 0x39, 0xd3, 0x02, 0x93, 0x22, 0x10, 0x91, 0xdd, 0x11,
	0x6d, 0x57, 0xfa, 0x82, 0x47, 0xf6, 0xf9, 0x6a, 0xb6, 0x68, 0x04, 0xa1, 0x90, 0x82, 0x52, 0x85,
	0x69, 0x64, 0xe5, 0xf3, 0xd5, 0xea, 0x82, 0x27, 0x84, 0xd7, 0x41, 0xdb, 0x0d, 0x7c, 0xdb, 0xe5,
	0x5c, 0xc8, 0xfe, 0x1d, 0xac, 0x06, 0xe0, 0xa0, 0xe7, 0x0b, 0x7e, 0xd4, 0x0b, 0x90, 0x52, 0x98,
	0xe2, 0xee, 0x19, 0x56, 0x48, 0x8d, 0xd4, 0x2d, 0x47, 0x3d, 0xb3, 0x2d, 0x28, 0x6a, 0x44, 0x5e,
	0x97, 0x56, 0xa1, 0xe4, 0xa1, 0x38, 0x43, 0x19, 0xf6, 0x2a, 0x13, 0x35, 0x52, 0xbf, 0xef, 0xa4,
	0x6b, 0xb6, 0x04, 0x73, 0xfb, 0x28, 0xf5, 0x66, 0x07, 0xbf, 0x74, 0x31, 0x92, 0xb9, 0x6f, 0xb8,
	0x21, 0x50, 0x3e, 0x44, 0x37, 0x6c, 0x7f, 0xd6, 0xd8, 0x28, 0x01, 0x3f, 0x06, 0x2b, 0x70, 0x3d,
	0x6c, 0x45, 0xfe, 0xa5, 0xde, 0x51, 0x70, 0x4a, 0x71, 0xe1, 0xd0, 0xbf, 0x44, 0xfa, 0x04, 0x40,
	0x35, 0xa5, 0x38, 0x45, 0xae, 0xde, 0x6d, 0x39, 0x0a, 0x7e, 0x14, 0x17, 0xe8, 0x33, 0x98, 0x0e,
	0x15, 0x5b, 0x4b, 0xf6, 0x02, 0xac, 0x4c, 0xaa, 0x3e, 0x84, 0x99, 0xd7, 0x75, 0x98, 0xf7, 0xb9,
	0xc4, 0x30, 0xc2, 0xb6, 0xf4, 0xb9, 0xd7, 0x4a, 0x6d, 0x4c, 0x29, 0x1b, 0xe5, 0xfe, 0xe6, 0xbe,
	0xe9, 0xd1, 0x65, 0x98, 0xc3, 0x8b, 0x76, 0xa7, 0x7b, 0x8c, 0x19, 0xbe, 0x50, 0x23, 0xf5, 0x92,
	0x33, 0x6b, 0xea, 0x09, 0x94, 0x75, 0x61, 0x7e, 0xc8, 0x54, 0x14, 0x08, 0x1e, 0x21, 0xdd, 0x80,
	0x7b, 0x5a, 0x46, 0x54, 0x21, 0xb5, 0xc9, 0xfa, 0xf4, 0x5a, 0xb5, 0x31, 0x7a, 0x6c, 0x0d, 0x33,
	0xb6, 0x04, 0x4a, 0x97, 0x60, 0x96, 0xe3, 0x85, 0x6c, 0x8d, 0x78, 0x9e, 0x89, 0xcb, 0xef, 0x12,
	0xdf, 0xec, 0x23, 0xd4, 0xd2, 0xa1, 0xef, 0xa1, 0x2b, 0xbb, 0x21, 0x1e, 0xa2, 0x7c, 0xef, 0x76,
	0xba, 0x98, 0xce, 0xf5, 0x21, 0x14, 0x35, 0xad, 0x39, 0x06, 0xb3, 0x8a, 0x67, 0x76, 0xa2, 0xb7,
	0xb4, 0x22, 0x94, 0x86, 0x1f, 0x4e, 0x52, 0x16, 0xf6, 0x93, 0xc0, 0xf3, 0x31, 0xec, 0xc6, 0xa0,
	0x80, 0x07, 0x09, 0xcd, 0xb9, 0xea, 0x18, 0x9f, 0x07, 0x79, 0x3e, 0xff, 0x4a, 0xd7, 0x30, 0x0d,
	0x5d, 0x7d, 0xc3, 0x65, 0xd8, 0x73, 0x66, 0x4e, 0xfa, 0x6b, 0xd5, 0xd7, 0x40, 0x47, 0x41, 0x74,
	0x0e, 0x26, 0x4f, 0xb1, 0x67, 0x2c, 0xc6, 0x8f, 0xb4, 0x0c, 0x05, 0x25, 0x48, 0x39, 0x23, 0x8e,
	0x5e, 0x6c, 0x4f, 0x6c, 0x11, 0xb6, 0x0b, 0xc5, 0xa6, 0x2b, 0x9b, 0xdc, 0x8b, 0x03, 0xdd, 0x71,
	0xa5, 0x2f, 0xbb, 0xc7, 0x3a, 0x72, 0xc4, 0x49, 0xd7, 0x74, 0x01, 0xac, 0x8e, 0xe0, 0x9e, 0x6e,
	0x6a, 0x8e, 0xac, 0xc0, 0x04, 0x94, 0x9b, 0xb1, 0x33, 0x1c, 0x4a, 0xf1, 0x50, 0x12, 0xc9, 0x48,
	0x12, 0x37, 0xa1, 0x94, 0x8c, 0x44, 0xb1, 0xde, 0x92, 0x08, 0x2d, 0xd0, 0x49, 0xb1, 0x6c, 0x15,
	0xe6, 0x87, 0x5e, 0x68, 0x0e, 0xa0, 0x32, 0x98, 0x30, 0x2b, 0x4d, 0xd1, 0xda, 0x8f, 0x02, 0x58,
	0xcd, 0x84, 0x94, 0xf6, 0xc0, 0x4a, 0xc7, 0x4f, 0x17, 0xc7, 0x9e, 0x8e, 0x31, 0x53, 0x1d, 0x93,
	0x55, 0xb6, 0xf2, 0xf5, 0xe6, 0xf7, 0xb7, 0x89, 0x45, 0xca, 0xe2, 0x4f, 0xd3, 0x55, 0x7c, 0xb5,
	0x77, 0x32, 0x97, 0x91, 0xbd, 0x62, 0x1b, 0x15, 0xf6, 0xca, 0x35, 0xfd, 0x45, 0xe0, 0xd1, 0xad,
	0x47, 0x4f, 0x37, 0xee, 0x98, 0x14, 0xad, 0xed, 0xe5, 0x7f, 0xe5, 0x8b, 0xbd, 0x55, 0xb2, 0x0f,
	0xe8, 0x9e, 0x92, 0xad, 0x35, 0xde, 0x2a, 0xdc, 0xbe, 0xea, 0xbb, 0x21, 0x3b, 0xd9, 0xd5, 0x88,
	0x5b, 0x3a, 0xec, 0xf4, 0x3b, 0x81, 0x99, 0x81, 0x9b, 0x4f, 0xeb, 0x79, 0xc2, 0xf2, 0xbe, 0x78,
	0xd5, 0xe5, 0x7f, 0x40, 0x1a, 0xd9, 0x9b, 0x4a, 0xf6, 0x0b, 0xd6, 0xe8, 0x93, 0xad, 0x12, 0x36,
	0xa8, 0xfd, 0x3a, 0x11, 0xbf, 0x1d, 0x29, 0x1a, 0x25, 0x6f, 0x20, 0x36, 0xf9, 0xf2, 0xf2, 0xa2,
	0x9c, 0x2f, 0x2f, 0x37, 0x83, 0x77, 0x97, 0xa7, 0x38, 0x71, 0xf7, 0xe9, 0x87, 0x85, 0xd1, 0x1f,
	0xdd, 0xab, 0x74, 0xf1, 0xa9, 0xa8, 0xfe, 0x5b, 0xeb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x78,
	0xf2, 0x7e, 0x1c, 0x0f, 0x07, 0x00, 0x00,
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
