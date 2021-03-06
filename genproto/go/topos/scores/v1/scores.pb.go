// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topos/scores/v1/scores.proto

package scores

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Score struct {
	// The [resource name](https://cloud.google.com/apis/design/resource_names) of
	// the origin vertex. It must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[\w-]+(/[\w-]+)*$`.
	VertexA string `protobuf:"bytes,1,opt,name=vertex_a,json=vertexA,proto3" json:"vertex_a,omitempty"`
	// The [resource name](https://cloud.google.com/apis/design/resource_names) of
	// the destination vertex. It must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[\w-]+(/[\w-]+)*$`.
	VertexB string `protobuf:"bytes,2,opt,name=vertex_b,json=vertexB,proto3" json:"vertex_b,omitempty"`
	// The score value.
	Score                float64  `protobuf:"fixed64,3,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Score) Reset()         { *m = Score{} }
func (m *Score) String() string { return proto.CompactTextString(m) }
func (*Score) ProtoMessage()    {}
func (*Score) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{0}
}

func (m *Score) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Score.Unmarshal(m, b)
}
func (m *Score) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Score.Marshal(b, m, deterministic)
}
func (m *Score) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Score.Merge(m, src)
}
func (m *Score) XXX_Size() int {
	return xxx_messageInfo_Score.Size(m)
}
func (m *Score) XXX_DiscardUnknown() {
	xxx_messageInfo_Score.DiscardUnknown(m)
}

var xxx_messageInfo_Score proto.InternalMessageInfo

func (m *Score) GetVertexA() string {
	if m != nil {
		return m.VertexA
	}
	return ""
}

func (m *Score) GetVertexB() string {
	if m != nil {
		return m.VertexB
	}
	return ""
}

func (m *Score) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type GetGraphScoreRequest struct {
	// The name of the graph.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource name of the origin vertex.
	VertexA string `protobuf:"bytes,2,opt,name=vertex_a,json=vertexA,proto3" json:"vertex_a,omitempty"`
	// The resource name of the destination vertex.
	VertexB              string   `protobuf:"bytes,3,opt,name=vertex_b,json=vertexB,proto3" json:"vertex_b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGraphScoreRequest) Reset()         { *m = GetGraphScoreRequest{} }
func (m *GetGraphScoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetGraphScoreRequest) ProtoMessage()    {}
func (*GetGraphScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{1}
}

func (m *GetGraphScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGraphScoreRequest.Unmarshal(m, b)
}
func (m *GetGraphScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGraphScoreRequest.Marshal(b, m, deterministic)
}
func (m *GetGraphScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGraphScoreRequest.Merge(m, src)
}
func (m *GetGraphScoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetGraphScoreRequest.Size(m)
}
func (m *GetGraphScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGraphScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGraphScoreRequest proto.InternalMessageInfo

func (m *GetGraphScoreRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetGraphScoreRequest) GetVertexA() string {
	if m != nil {
		return m.VertexA
	}
	return ""
}

func (m *GetGraphScoreRequest) GetVertexB() string {
	if m != nil {
		return m.VertexB
	}
	return ""
}

type TopGraphScoresRequest struct {
	// The name of the graph. It must have the format `"graphs/{graph}"`.
	// `{graphs}` must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[a-z\d]+([.-][a-z\d]+)*$`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The resource name of the origin vertex.
	VertexA string `protobuf:"bytes,3,opt,name=vertex_a,json=vertexA,proto3" json:"vertex_a,omitempty"`
	// The resource name of the destination vertex.
	VertexB string `protobuf:"bytes,4,opt,name=vertex_b,json=vertexB,proto3" json:"vertex_b,omitempty"`
	// Return the scores in ascending order by score value.
	AscendingOrder bool `protobuf:"varint,5,opt,name=ascending_order,json=ascendingOrder,proto3" json:"ascending_order,omitempty"`
	// If any are specified, only return scores where the name of vertex
	// unspecified in the query is one of the filter values.
	VertexFilters        []string `protobuf:"bytes,6,rep,name=vertex_filters,json=vertexFilters,proto3" json:"vertex_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopGraphScoresRequest) Reset()         { *m = TopGraphScoresRequest{} }
func (m *TopGraphScoresRequest) String() string { return proto.CompactTextString(m) }
func (*TopGraphScoresRequest) ProtoMessage()    {}
func (*TopGraphScoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{2}
}

func (m *TopGraphScoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopGraphScoresRequest.Unmarshal(m, b)
}
func (m *TopGraphScoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopGraphScoresRequest.Marshal(b, m, deterministic)
}
func (m *TopGraphScoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopGraphScoresRequest.Merge(m, src)
}
func (m *TopGraphScoresRequest) XXX_Size() int {
	return xxx_messageInfo_TopGraphScoresRequest.Size(m)
}
func (m *TopGraphScoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopGraphScoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopGraphScoresRequest proto.InternalMessageInfo

func (m *TopGraphScoresRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopGraphScoresRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *TopGraphScoresRequest) GetVertexA() string {
	if m != nil {
		return m.VertexA
	}
	return ""
}

func (m *TopGraphScoresRequest) GetVertexB() string {
	if m != nil {
		return m.VertexB
	}
	return ""
}

func (m *TopGraphScoresRequest) GetAscendingOrder() bool {
	if m != nil {
		return m.AscendingOrder
	}
	return false
}

func (m *TopGraphScoresRequest) GetVertexFilters() []string {
	if m != nil {
		return m.VertexFilters
	}
	return nil
}

type TopGraphScoresResponse struct {
	Scores               []*Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopGraphScoresResponse) Reset()         { *m = TopGraphScoresResponse{} }
func (m *TopGraphScoresResponse) String() string { return proto.CompactTextString(m) }
func (*TopGraphScoresResponse) ProtoMessage()    {}
func (*TopGraphScoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{3}
}

func (m *TopGraphScoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopGraphScoresResponse.Unmarshal(m, b)
}
func (m *TopGraphScoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopGraphScoresResponse.Marshal(b, m, deterministic)
}
func (m *TopGraphScoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopGraphScoresResponse.Merge(m, src)
}
func (m *TopGraphScoresResponse) XXX_Size() int {
	return xxx_messageInfo_TopGraphScoresResponse.Size(m)
}
func (m *TopGraphScoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopGraphScoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopGraphScoresResponse proto.InternalMessageInfo

func (m *TopGraphScoresResponse) GetScores() []*Score {
	if m != nil {
		return m.Scores
	}
	return nil
}

func init() {
	proto.RegisterType((*Score)(nil), "topos.scores.v1.Score")
	proto.RegisterType((*GetGraphScoreRequest)(nil), "topos.scores.v1.GetGraphScoreRequest")
	proto.RegisterType((*TopGraphScoresRequest)(nil), "topos.scores.v1.TopGraphScoresRequest")
	proto.RegisterType((*TopGraphScoresResponse)(nil), "topos.scores.v1.TopGraphScoresResponse")
}

func init() {
	proto.RegisterFile("topos/scores/v1/scores.proto", fileDescriptor_05589be8939bdcb5)
}

var fileDescriptor_05589be8939bdcb5 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x96, 0x93, 0x25, 0xb4, 0x46, 0xeb, 0x24, 0x6b, 0x4c, 0x21, 0x54, 0x28, 0x8a, 0xb4, 0x2d,
	0x42, 0x22, 0xd6, 0xc6, 0x0d, 0xc4, 0x61, 0x43, 0x62, 0xdc, 0x90, 0x52, 0x4e, 0x5c, 0x8a, 0xdb,
	0x19, 0xcf, 0x52, 0x6b, 0x1b, 0xdb, 0xad, 0x50, 0x11, 0x17, 0x4e, 0x70, 0xe6, 0x7f, 0x71, 0xe1,
	0x2f, 0xf0, 0x43, 0x50, 0xec, 0x00, 0x4d, 0xdb, 0xb0, 0xdb, 0xf3, 0xf7, 0x9e, 0xbf, 0xef, 0xbd,
	0xf7, 0xd9, 0x70, 0x68, 0xa5, 0x92, 0x06, 0x9b, 0xa9, 0xd4, 0xd4, 0xe0, 0xe5, 0x59, 0x13, 0x95,
	0x4a, 0x4b, 0x2b, 0xd1, 0x81, 0xcb, 0x96, 0x0d, 0xb6, 0x3c, 0x4b, 0x87, 0x4c, 0x4a, 0x36, 0xa3,
	0x98, 0x28, 0x8e, 0x89, 0x10, 0xd2, 0x12, 0xcb, 0xa5, 0x68, 0xca, 0xf3, 0x11, 0x8c, 0x46, 0x75,
	0x29, 0xba, 0x0f, 0x7b, 0x4b, 0xaa, 0x2d, 0xfd, 0x38, 0x26, 0x09, 0xc8, 0x40, 0xd1, 0xaf, 0xee,
	0xf8, 0xf3, 0xc5, 0x5a, 0x6a, 0x92, 0x04, 0xeb, 0xa9, 0x4b, 0x74, 0x08, 0x23, 0xa7, 0x94, 0x84,
	0x19, 0x28, 0x40, 0xe5, 0x0f, 0xf9, 0x3b, 0x78, 0x78, 0x45, 0xed, 0x95, 0x26, 0xea, 0xc6, 0x91,
	0x57, 0xf4, 0xc3, 0x82, 0x1a, 0x8b, 0x10, 0xdc, 0x13, 0x64, 0x4e, 0x1b, 0x7e, 0x17, 0xb7, 0x74,
	0x83, 0x6e, 0xdd, 0xb0, 0xa5, 0x9b, 0xff, 0x00, 0xf0, 0xde, 0x1b, 0xa9, 0xfe, 0x49, 0x98, 0xff,
	0x69, 0x3c, 0x80, 0x7d, 0x45, 0x18, 0x1d, 0x1b, 0xbe, 0xa2, 0x4e, 0x24, 0xaa, 0x7a, 0x35, 0x30,
	0xe2, 0xab, 0x76, 0x03, 0x61, 0x77, 0x03, 0x7b, 0xed, 0xc1, 0x4f, 0xe1, 0x01, 0x31, 0x53, 0x2a,
	0xae, 0xb9, 0x60, 0x63, 0xa9, 0xaf, 0xa9, 0x4e, 0xa2, 0x0c, 0x14, 0xbd, 0x6a, 0xf0, 0x17, 0x7e,
	0x5d, 0xa3, 0xe8, 0x18, 0x0e, 0x1a, 0x8e, 0xf7, 0x7c, 0x66, 0xa9, 0x36, 0x49, 0x9c, 0x85, 0x45,
	0xbf, 0xda, 0xf7, 0xe8, 0x4b, 0x0f, 0xe6, 0xaf, 0xe0, 0xd1, 0xe6, 0x3c, 0x46, 0x49, 0x61, 0x28,
	0x2a, 0x61, 0xec, 0xcd, 0x4c, 0x40, 0x16, 0x16, 0x77, 0xcf, 0x8f, 0xca, 0x0d, 0x87, 0x4b, 0xbf,
	0xe3, 0xa6, 0xea, 0xfc, 0x6b, 0x00, 0x63, 0x4f, 0x81, 0x34, 0xdc, 0x6f, 0xf9, 0x80, 0x8e, 0xb7,
	0xee, 0xee, 0xf2, 0x29, 0xed, 0x90, 0xc8, 0xf3, 0x2f, 0x3f, 0x7f, 0x7d, 0x0f, 0x86, 0x28, 0xad,
	0x5f, 0xdd, 0xa7, 0x7a, 0xb5, 0xcf, 0x59, 0x7d, 0xd9, 0xe0, 0x47, 0x9f, 0x9b, 0x57, 0x88, 0xbe,
	0x01, 0x38, 0x68, 0x4f, 0x82, 0x4e, 0xb6, 0xe8, 0x76, 0x5a, 0x97, 0x9e, 0xde, 0x5a, 0xe7, 0x57,
	0x92, 0x9f, 0xb8, 0x3e, 0x32, 0xf4, 0xb0, 0xbb, 0x8f, 0xa7, 0x56, 0xaa, 0xcb, 0x17, 0x6f, 0x2f,
	0x18, 0xb7, 0x37, 0x8b, 0x49, 0x39, 0x95, 0x73, 0xec, 0xc8, 0x1f, 0x13, 0xfe, 0x27, 0x50, 0xdc,
	0x60, 0x46, 0x85, 0xfb, 0x07, 0x98, 0x49, 0xbc, 0xf1, 0xad, 0x9e, 0xf9, 0x68, 0x12, 0xbb, 0x82,
	0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x3a, 0xaa, 0x40, 0x77, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ScoresClient is the client API for Scores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScoresClient interface {
	// Gets a graph score.
	GetGraphScore(ctx context.Context, in *GetGraphScoreRequest, opts ...grpc.CallOption) (*Score, error)
	// Gets graph scores in descending order by score value.
	TopGraphScores(ctx context.Context, in *TopGraphScoresRequest, opts ...grpc.CallOption) (*TopGraphScoresResponse, error)
}

type scoresClient struct {
	cc grpc.ClientConnInterface
}

func NewScoresClient(cc grpc.ClientConnInterface) ScoresClient {
	return &scoresClient{cc}
}

func (c *scoresClient) GetGraphScore(ctx context.Context, in *GetGraphScoreRequest, opts ...grpc.CallOption) (*Score, error) {
	out := new(Score)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/GetGraphScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoresClient) TopGraphScores(ctx context.Context, in *TopGraphScoresRequest, opts ...grpc.CallOption) (*TopGraphScoresResponse, error) {
	out := new(TopGraphScoresResponse)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/TopGraphScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoresServer is the server API for Scores service.
type ScoresServer interface {
	// Gets a graph score.
	GetGraphScore(context.Context, *GetGraphScoreRequest) (*Score, error)
	// Gets graph scores in descending order by score value.
	TopGraphScores(context.Context, *TopGraphScoresRequest) (*TopGraphScoresResponse, error)
}

// UnimplementedScoresServer can be embedded to have forward compatible implementations.
type UnimplementedScoresServer struct {
}

func (*UnimplementedScoresServer) GetGraphScore(ctx context.Context, req *GetGraphScoreRequest) (*Score, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraphScore not implemented")
}
func (*UnimplementedScoresServer) TopGraphScores(ctx context.Context, req *TopGraphScoresRequest) (*TopGraphScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopGraphScores not implemented")
}

func RegisterScoresServer(s *grpc.Server, srv ScoresServer) {
	s.RegisterService(&_Scores_serviceDesc, srv)
}

func _Scores_GetGraphScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGraphScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).GetGraphScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/GetGraphScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).GetGraphScore(ctx, req.(*GetGraphScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scores_TopGraphScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopGraphScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).TopGraphScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/TopGraphScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).TopGraphScores(ctx, req.(*TopGraphScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scores_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topos.scores.v1.Scores",
	HandlerType: (*ScoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGraphScore",
			Handler:    _Scores_GetGraphScore_Handler,
		},
		{
			MethodName: "TopGraphScores",
			Handler:    _Scores_TopGraphScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topos/scores/v1/scores.proto",
}
