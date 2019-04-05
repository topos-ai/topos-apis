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
	IntersectingGeometry []byte   `protobuf:"bytes,4,opt,name=intersecting_geometry,json=intersectingGeometry,proto3" json:"intersecting_geometry,omitempty"`
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

func init() {
	proto.RegisterType((*RegionType)(nil), "topos.locations.v1.RegionType")
	proto.RegisterType((*Region)(nil), "topos.locations.v1.Region")
	proto.RegisterType((*GetRegionRequest)(nil), "topos.locations.v1.GetRegionRequest")
	proto.RegisterType((*SearchRegionsRequest)(nil), "topos.locations.v1.SearchRegionsRequest")
	proto.RegisterType((*SearchRegionsResponse)(nil), "topos.locations.v1.SearchRegionsResponse")
}

func init() { proto.RegisterFile("topos/locations/v1/locations.proto", fileDescriptor_954338d7542bc8de) }

var fileDescriptor_954338d7542bc8de = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0xae, 0xd2, 0x50,
	0x10, 0x4e, 0x01, 0x91, 0x0e, 0x12, 0xcd, 0x09, 0x24, 0x4d, 0x45, 0x6d, 0x1a, 0x42, 0x2a, 0x8b,
	0x56, 0xc0, 0x18, 0xa3, 0x71, 0xe3, 0x86, 0x8d, 0x0b, 0x53, 0x58, 0xb9, 0x69, 0x2a, 0x99, 0xd4,
	0x46, 0x38, 0xa7, 0xf6, 0x1c, 0x88, 0xc5, 0xb0, 0xf1, 0x15, 0xdc, 0xba, 0xbd, 0xc9, 0x7d, 0x9f,
	0xfb, 0x0a, 0xf7, 0x41, 0x6e, 0x7a, 0x4a, 0x5b, 0x7e, 0x9a, 0x9b, 0xbb, 0xeb, 0xcc, 0x7c, 0x33,
	0xf3, 0x7d, 0xdf, 0x9c, 0x82, 0x29, 0x58, 0xc4, 0xb8, 0xb3, 0x62, 0x4b, 0x5f, 0x84, 0x8c, 0x72,
	0x67, 0x3b, 0x2e, 0x03, 0x3b, 0x8a, 0x99, 0x60, 0x84, 0x48, 0x8c, 0x5d, 0xa6, 0xb7, 0x63, 0xbd,
	0x1f, 0x30, 0x16, 0xac, 0xd0, 0xf1, 0xa3, 0xd0, 0xf1, 0x29, 0x65, 0xe2, 0xb8, 0xc3, 0x34, 0x00,
	0x5c, 0x0c, 0x42, 0x46, 0x17, 0x49, 0x84, 0x84, 0x40, 0x83, 0xfa, 0x6b, 0xd4, 0x14, 0x43, 0xb1,
	0x54, 0x57, 0x7e, 0x9b, 0xef, 0xa1, 0x99, 0x21, 0xaa, 0xaa, 0x44, 0x87, 0x56, 0x80, 0x6c, 0x8d,
	0x22, 0x4e, 0xb4, 0x9a, 0xa1, 0x58, 0x4f, 0xdc, 0x22, 0x36, 0x87, 0xf0, 0x6c, 0x86, 0x22, 0x6b,
	0x76, 0xf1, 0xd7, 0x06, 0xb9, 0xa8, 0xdc, 0x70, 0xad, 0x40, 0x77, 0x8e, 0x7e, 0xbc, 0xfc, 0x91,
	0x61, 0x79, 0x0e, 0x7e, 0x0e, 0x6a, 0xe4, 0x07, 0xe8, 0xf1, 0x70, 0x97, 0x75, 0x3c, 0x72, 0x5b,
	0x69, 0x62, 0x1e, 0xee, 0x90, 0xbc, 0x00, 0x90, 0x45, 0xc1, 0x7e, 0x22, 0x95, 0xbb, 0x55, 0x57,
	0xc2, 0x17, 0x69, 0x82, 0xbc, 0x82, 0x76, 0x2c, 0xa7, 0x79, 0x22, 0x89, 0x50, 0xab, 0xcb, 0x3a,
	0xc4, 0xa5, 0xd6, 0x29, 0xf4, 0x42, 0x2a, 0x30, 0xe6, 0xb8, 0x14, 0x21, 0x0d, 0xbc, 0x42, 0x46,
	0x43, 0xca, 0xe8, 0x1e, 0x17, 0x67, 0xb9, 0xa4, 0x0d, 0xf4, 0xce, 0x98, 0xf2, 0x88, 0x51, 0x8e,
	0xe4, 0x2d, 0x3c, 0xce, 0x66, 0x73, 0x4d, 0x31, 0xea, 0x56, 0x7b, 0xa2, 0xdb, 0x97, 0xb7, 0xb0,
	0x0f, 0x5e, 0xe4, 0x50, 0x32, 0x84, 0xa7, 0x14, 0x7f, 0x0b, 0xef, 0x42, 0x48, 0x27, 0x4d, 0x7f,
	0xcd, 0xc5, 0x4c, 0xae, 0x6a, 0xa0, 0x7e, 0xc9, 0x07, 0x91, 0x04, 0xd4, 0xc2, 0x57, 0x32, 0xa8,
	0xda, 0x73, 0x6e, 0xbb, 0x7e, 0x0f, 0x1b, 0x73, 0xf4, 0xf7, 0xe6, 0xf6, 0x5f, 0x6d, 0x40, 0xcc,
	0xf4, 0x45, 0xfd, 0x49, 0x2f, 0xf2, 0xa9, 0xb4, 0x89, 0x3b, 0x23, 0xe7, 0xc0, 0xd6, 0x19, 0xed,
	0xc9, 0x7f, 0x05, 0x3a, 0x27, 0x06, 0x10, 0xab, 0x6a, 0x72, 0xd5, 0x35, 0xf5, 0xd7, 0x0f, 0x40,
	0x66, 0x6e, 0x9a, 0xef, 0x24, 0xa5, 0x37, 0xc4, 0x96, 0x94, 0x8e, 0xee, 0x78, 0xca, 0x6c, 0x9f,
	0x53, 0xfb, 0xc0, 0xe5, 0x98, 0xcf, 0x2f, 0xbf, 0xf5, 0x2f, 0xff, 0x92, 0x8f, 0x45, 0xf0, 0xbd,
	0x29, 0x1f, 0xfd, 0xf4, 0x2e, 0x00, 0x00, 0xff, 0xff, 0xac, 0xa5, 0x15, 0xea, 0x4c, 0x03, 0x00,
	0x00,
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
	// Search regions.
	SearchRegions(ctx context.Context, in *SearchRegionsRequest, opts ...grpc.CallOption) (*SearchRegionsResponse, error)
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

func (c *locationsClient) SearchRegions(ctx context.Context, in *SearchRegionsRequest, opts ...grpc.CallOption) (*SearchRegionsResponse, error) {
	out := new(SearchRegionsResponse)
	err := c.cc.Invoke(ctx, "/topos.locations.v1.Locations/SearchRegions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationsServer is the server API for Locations service.
type LocationsServer interface {
	// Gets a region.
	GetRegion(context.Context, *GetRegionRequest) (*Region, error)
	// Search regions.
	SearchRegions(context.Context, *SearchRegionsRequest) (*SearchRegionsResponse, error)
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

var _Locations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topos.locations.v1.Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegion",
			Handler:    _Locations_GetRegion_Handler,
		},
		{
			MethodName: "SearchRegions",
			Handler:    _Locations_SearchRegions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topos/locations/v1/locations.proto",
}
