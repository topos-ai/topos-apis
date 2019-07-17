// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topos/scores/v1/scores.proto

package scores

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ListGraphScoresRequest struct {
	// The name of the graph.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGraphScoresRequest) Reset()         { *m = ListGraphScoresRequest{} }
func (m *ListGraphScoresRequest) String() string { return proto.CompactTextString(m) }
func (*ListGraphScoresRequest) ProtoMessage()    {}
func (*ListGraphScoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{1}
}

func (m *ListGraphScoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGraphScoresRequest.Unmarshal(m, b)
}
func (m *ListGraphScoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGraphScoresRequest.Marshal(b, m, deterministic)
}
func (m *ListGraphScoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGraphScoresRequest.Merge(m, src)
}
func (m *ListGraphScoresRequest) XXX_Size() int {
	return xxx_messageInfo_ListGraphScoresRequest.Size(m)
}
func (m *ListGraphScoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGraphScoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGraphScoresRequest proto.InternalMessageInfo

func (m *ListGraphScoresRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListGraphScoresRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListGraphScoresRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListGraphScoresResponse struct {
	// The name of the score.
	Scores []*Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no more
	// results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGraphScoresResponse) Reset()         { *m = ListGraphScoresResponse{} }
func (m *ListGraphScoresResponse) String() string { return proto.CompactTextString(m) }
func (*ListGraphScoresResponse) ProtoMessage()    {}
func (*ListGraphScoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{2}
}

func (m *ListGraphScoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGraphScoresResponse.Unmarshal(m, b)
}
func (m *ListGraphScoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGraphScoresResponse.Marshal(b, m, deterministic)
}
func (m *ListGraphScoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGraphScoresResponse.Merge(m, src)
}
func (m *ListGraphScoresResponse) XXX_Size() int {
	return xxx_messageInfo_ListGraphScoresResponse.Size(m)
}
func (m *ListGraphScoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGraphScoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGraphScoresResponse proto.InternalMessageInfo

func (m *ListGraphScoresResponse) GetScores() []*Score {
	if m != nil {
		return m.Scores
	}
	return nil
}

func (m *ListGraphScoresResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
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
	return fileDescriptor_05589be8939bdcb5, []int{3}
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

type SetGraphScoreRequest struct {
	// The name of the graph. It must have the format `"graphs/{graph}"`.
	// `{graphs}` must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[a-z\d]+(-[a-z\d]+)*$`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The score to set
	Score                *Score   `protobuf:"bytes,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGraphScoreRequest) Reset()         { *m = SetGraphScoreRequest{} }
func (m *SetGraphScoreRequest) String() string { return proto.CompactTextString(m) }
func (*SetGraphScoreRequest) ProtoMessage()    {}
func (*SetGraphScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{4}
}

func (m *SetGraphScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGraphScoreRequest.Unmarshal(m, b)
}
func (m *SetGraphScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGraphScoreRequest.Marshal(b, m, deterministic)
}
func (m *SetGraphScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGraphScoreRequest.Merge(m, src)
}
func (m *SetGraphScoreRequest) XXX_Size() int {
	return xxx_messageInfo_SetGraphScoreRequest.Size(m)
}
func (m *SetGraphScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGraphScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGraphScoreRequest proto.InternalMessageInfo

func (m *SetGraphScoreRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetGraphScoreRequest) GetScore() *Score {
	if m != nil {
		return m.Score
	}
	return nil
}

type BatchSetGraphScoresRequest struct {
	// The name of the graph. It must have the format `"graphs/{graph}"`.
	// `{graphs}` must match the [regular
	// expression](https://github.com/google/re2/wiki/Syntax)
	// `^[a-z\d]+(-[a-z\d]+)*$`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The scores to set.
	Scores               []*Score `protobuf:"bytes,2,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchSetGraphScoresRequest) Reset()         { *m = BatchSetGraphScoresRequest{} }
func (m *BatchSetGraphScoresRequest) String() string { return proto.CompactTextString(m) }
func (*BatchSetGraphScoresRequest) ProtoMessage()    {}
func (*BatchSetGraphScoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{5}
}

func (m *BatchSetGraphScoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchSetGraphScoresRequest.Unmarshal(m, b)
}
func (m *BatchSetGraphScoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchSetGraphScoresRequest.Marshal(b, m, deterministic)
}
func (m *BatchSetGraphScoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSetGraphScoresRequest.Merge(m, src)
}
func (m *BatchSetGraphScoresRequest) XXX_Size() int {
	return xxx_messageInfo_BatchSetGraphScoresRequest.Size(m)
}
func (m *BatchSetGraphScoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSetGraphScoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSetGraphScoresRequest proto.InternalMessageInfo

func (m *BatchSetGraphScoresRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BatchSetGraphScoresRequest) GetScores() []*Score {
	if m != nil {
		return m.Scores
	}
	return nil
}

type BatchSetGraphScoresResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchSetGraphScoresResponse) Reset()         { *m = BatchSetGraphScoresResponse{} }
func (m *BatchSetGraphScoresResponse) String() string { return proto.CompactTextString(m) }
func (*BatchSetGraphScoresResponse) ProtoMessage()    {}
func (*BatchSetGraphScoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{6}
}

func (m *BatchSetGraphScoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchSetGraphScoresResponse.Unmarshal(m, b)
}
func (m *BatchSetGraphScoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchSetGraphScoresResponse.Marshal(b, m, deterministic)
}
func (m *BatchSetGraphScoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchSetGraphScoresResponse.Merge(m, src)
}
func (m *BatchSetGraphScoresResponse) XXX_Size() int {
	return xxx_messageInfo_BatchSetGraphScoresResponse.Size(m)
}
func (m *BatchSetGraphScoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchSetGraphScoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchSetGraphScoresResponse proto.InternalMessageInfo

type DeleteGraphRequest struct {
	// The name of the graph.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGraphRequest) Reset()         { *m = DeleteGraphRequest{} }
func (m *DeleteGraphRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGraphRequest) ProtoMessage()    {}
func (*DeleteGraphRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{7}
}

func (m *DeleteGraphRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGraphRequest.Unmarshal(m, b)
}
func (m *DeleteGraphRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGraphRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGraphRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGraphRequest.Merge(m, src)
}
func (m *DeleteGraphRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGraphRequest.Size(m)
}
func (m *DeleteGraphRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGraphRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGraphRequest proto.InternalMessageInfo

func (m *DeleteGraphRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteGraphScoreRequest struct {
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

func (m *DeleteGraphScoreRequest) Reset()         { *m = DeleteGraphScoreRequest{} }
func (m *DeleteGraphScoreRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGraphScoreRequest) ProtoMessage()    {}
func (*DeleteGraphScoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{8}
}

func (m *DeleteGraphScoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGraphScoreRequest.Unmarshal(m, b)
}
func (m *DeleteGraphScoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGraphScoreRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGraphScoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGraphScoreRequest.Merge(m, src)
}
func (m *DeleteGraphScoreRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGraphScoreRequest.Size(m)
}
func (m *DeleteGraphScoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGraphScoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGraphScoreRequest proto.InternalMessageInfo

func (m *DeleteGraphScoreRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteGraphScoreRequest) GetVertexA() string {
	if m != nil {
		return m.VertexA
	}
	return ""
}

func (m *DeleteGraphScoreRequest) GetVertexB() string {
	if m != nil {
		return m.VertexB
	}
	return ""
}

type TopGraphScoresRequest struct {
	// The name of the graph.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The resource name of the origin vertex.
	VertexA string `protobuf:"bytes,3,opt,name=vertex_a,json=vertexA,proto3" json:"vertex_a,omitempty"`
	// The resource name of the destination vertex.
	VertexB              string   `protobuf:"bytes,4,opt,name=vertex_b,json=vertexB,proto3" json:"vertex_b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopGraphScoresRequest) Reset()         { *m = TopGraphScoresRequest{} }
func (m *TopGraphScoresRequest) String() string { return proto.CompactTextString(m) }
func (*TopGraphScoresRequest) ProtoMessage()    {}
func (*TopGraphScoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{9}
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

type TopGraphScoresResponse struct {
	// Scores in decending order by score value.
	Scores               []*Score `protobuf:"bytes,1,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopGraphScoresResponse) Reset()         { *m = TopGraphScoresResponse{} }
func (m *TopGraphScoresResponse) String() string { return proto.CompactTextString(m) }
func (*TopGraphScoresResponse) ProtoMessage()    {}
func (*TopGraphScoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05589be8939bdcb5, []int{10}
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
	proto.RegisterType((*ListGraphScoresRequest)(nil), "topos.scores.v1.ListGraphScoresRequest")
	proto.RegisterType((*ListGraphScoresResponse)(nil), "topos.scores.v1.ListGraphScoresResponse")
	proto.RegisterType((*GetGraphScoreRequest)(nil), "topos.scores.v1.GetGraphScoreRequest")
	proto.RegisterType((*SetGraphScoreRequest)(nil), "topos.scores.v1.SetGraphScoreRequest")
	proto.RegisterType((*BatchSetGraphScoresRequest)(nil), "topos.scores.v1.BatchSetGraphScoresRequest")
	proto.RegisterType((*BatchSetGraphScoresResponse)(nil), "topos.scores.v1.BatchSetGraphScoresResponse")
	proto.RegisterType((*DeleteGraphRequest)(nil), "topos.scores.v1.DeleteGraphRequest")
	proto.RegisterType((*DeleteGraphScoreRequest)(nil), "topos.scores.v1.DeleteGraphScoreRequest")
	proto.RegisterType((*TopGraphScoresRequest)(nil), "topos.scores.v1.TopGraphScoresRequest")
	proto.RegisterType((*TopGraphScoresResponse)(nil), "topos.scores.v1.TopGraphScoresResponse")
}

func init() { proto.RegisterFile("topos/scores/v1/scores.proto", fileDescriptor_05589be8939bdcb5) }

var fileDescriptor_05589be8939bdcb5 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe5, 0xa4, 0x4d, 0x9b, 0xa9, 0x42, 0xd0, 0x34, 0x4d, 0x83, 0x93, 0xa2, 0xc8, 0x15,
	0x89, 0x15, 0x8a, 0xad, 0x96, 0x5b, 0x10, 0x07, 0x22, 0x50, 0x39, 0x70, 0x40, 0x76, 0x0f, 0x88,
	0x4b, 0xea, 0x44, 0x4b, 0x12, 0xd1, 0x7a, 0xdd, 0xec, 0x36, 0x0a, 0x45, 0xe5, 0xc0, 0x09, 0xce,
	0x5c, 0x38, 0xf2, 0x4e, 0xbc, 0x02, 0x0f, 0x82, 0xbc, 0xde, 0x82, 0xe3, 0x7f, 0x09, 0x08, 0x6e,
	0xde, 0x9d, 0xd9, 0xf9, 0xed, 0x7e, 0xb3, 0x9f, 0x17, 0x1a, 0x9c, 0x7a, 0x94, 0x99, 0x6c, 0x48,
	0xa7, 0x84, 0x99, 0xb3, 0x43, 0xf9, 0x65, 0x78, 0x53, 0xca, 0x29, 0x96, 0x45, 0xd4, 0x90, 0x73,
	0xb3, 0x43, 0xb5, 0x31, 0xa2, 0x74, 0x74, 0x46, 0x4c, 0xc7, 0x9b, 0x98, 0x8e, 0xeb, 0x52, 0xee,
	0xf0, 0x09, 0x75, 0x65, 0xba, 0x5a, 0x97, 0x51, 0x31, 0x1a, 0x5c, 0xbe, 0x31, 0xc9, 0xb9, 0xc7,
	0xdf, 0x05, 0x41, 0xcd, 0x86, 0x75, 0xdb, 0xaf, 0x83, 0x77, 0x60, 0x73, 0x46, 0xa6, 0x9c, 0xcc,
	0xfb, 0x4e, 0x4d, 0x69, 0x2a, 0x7a, 0xd1, 0xda, 0x08, 0xc6, 0x4f, 0x42, 0xa1, 0x41, 0x2d, 0x17,
	0x0e, 0xf5, 0xb0, 0x02, 0xeb, 0x62, 0x1b, 0xb5, 0x7c, 0x53, 0xd1, 0x15, 0x2b, 0x18, 0x68, 0x63,
	0xa8, 0xbe, 0x98, 0x30, 0x7e, 0x3c, 0x75, 0xbc, 0xb1, 0xa8, 0xce, 0x2c, 0x72, 0x71, 0x49, 0x18,
	0x47, 0x84, 0x35, 0xd7, 0x39, 0x27, 0x92, 0x20, 0xbe, 0xb1, 0x0e, 0x45, 0xcf, 0x19, 0x91, 0x3e,
	0x9b, 0x5c, 0x11, 0x51, 0x7f, 0xdd, 0xda, 0xf4, 0x27, 0xec, 0xc9, 0x15, 0xc1, 0x3d, 0x00, 0x11,
	0xe4, 0xf4, 0x2d, 0x71, 0x05, 0xa5, 0x68, 0x89, 0xf4, 0x13, 0x7f, 0x42, 0xbb, 0x80, 0xdd, 0x18,
	0x89, 0x79, 0xd4, 0x65, 0x04, 0x0d, 0x28, 0x04, 0x0a, 0xd5, 0x94, 0x66, 0x5e, 0xdf, 0x3a, 0xaa,
	0x1a, 0x11, 0xd9, 0x0c, 0xb1, 0xc0, 0x92, 0x59, 0xd8, 0x82, 0xb2, 0x4b, 0xe6, 0xbc, 0x1f, 0xc2,
	0x05, 0x87, 0x2d, 0xf9, 0xd3, 0x2f, 0x7f, 0x21, 0x4f, 0xa1, 0x72, 0x4c, 0x42, 0xc4, 0xac, 0xa3,
	0x85, 0x45, 0xcd, 0xa5, 0x8b, 0x9a, 0x5f, 0x10, 0x55, 0x7b, 0x05, 0x15, 0x7b, 0x55, 0xc2, 0xc1,
	0x4d, 0x03, 0xfc, 0xf2, 0xe9, 0x87, 0x94, 0x8d, 0x39, 0x05, 0xb5, 0xe7, 0xf0, 0xe1, 0x78, 0xa1,
	0x7c, 0x66, 0x73, 0x7e, 0xab, 0x98, 0x5b, 0x45, 0x45, 0x6d, 0x0f, 0xea, 0x89, 0x84, 0xa0, 0x29,
	0x9a, 0x0e, 0xf8, 0x94, 0x9c, 0x11, 0x4e, 0x44, 0x30, 0x03, 0xac, 0x0d, 0x61, 0x37, 0x94, 0xf9,
	0x9f, 0x94, 0xfe, 0x00, 0x3b, 0x27, 0xd4, 0xfb, 0x17, 0xf7, 0x34, 0xcc, 0xcf, 0xa7, 0xf3, 0xd7,
	0x16, 0xf9, 0xcf, 0xa1, 0x1a, 0xe5, 0xff, 0xdd, 0xed, 0x3d, 0xfa, 0xba, 0x01, 0x85, 0xa0, 0x04,
	0x7e, 0x52, 0xa0, 0x1c, 0x31, 0x05, 0xb6, 0x63, 0xcb, 0x93, 0x0d, 0xaa, 0xea, 0xcb, 0x13, 0x65,
	0x2b, 0xb5, 0x8f, 0xdf, 0x7f, 0x7c, 0xc9, 0x35, 0x50, 0xf5, 0xff, 0x4f, 0xef, 0x7d, 0x85, 0x1e,
	0x8f, 0xfc, 0x34, 0x66, 0x76, 0xae, 0xe5, 0xff, 0x0a, 0xa7, 0x50, 0x5a, 0xf0, 0x0a, 0xde, 0x8b,
	0x95, 0x4f, 0xf2, 0x92, 0x9a, 0x72, 0xda, 0x95, 0x98, 0xd7, 0x50, 0xb2, 0x97, 0x30, 0xed, 0x3f,
	0x61, 0x3e, 0x10, 0xcc, 0xb6, 0x76, 0x37, 0x9d, 0xd9, 0x65, 0x84, 0x77, 0x03, 0x8b, 0xe1, 0x37,
	0x05, 0xb6, 0x13, 0x1c, 0x80, 0xf7, 0x63, 0xe5, 0xd3, 0x9d, 0xa8, 0x1e, 0xac, 0x96, 0x2c, 0x3b,
	0x61, 0x88, 0x1d, 0xea, 0xda, 0x7e, 0xc6, 0x0e, 0x07, 0x72, 0x7d, 0x57, 0xe9, 0xe0, 0x08, 0xb6,
	0x42, 0xd6, 0xc2, 0xfd, 0x18, 0x2c, 0x6e, 0x51, 0xb5, 0x6a, 0x04, 0xaf, 0x88, 0x71, 0xf3, 0x8a,
	0x18, 0xcf, 0xfc, 0x57, 0x44, 0xab, 0x0b, 0xf6, 0x4e, 0x67, 0x3b, 0x81, 0x8d, 0x73, 0xb8, 0x1d,
	0xf5, 0x30, 0xea, 0x59, 0xb4, 0x48, 0x43, 0x92, 0x91, 0xf2, 0x12, 0x74, 0xb2, 0x2e, 0xc1, 0x67,
	0x05, 0x6e, 0x2d, 0x3a, 0x0b, 0x5b, 0x31, 0x70, 0xa2, 0xf5, 0xd5, 0xf6, 0xd2, 0x3c, 0x29, 0x7b,
	0x4b, 0xec, 0xa3, 0x89, 0x59, 0x17, 0x83, 0x53, 0xaf, 0x57, 0x7b, 0x5d, 0x8d, 0x3c, 0xe7, 0x8f,
	0x82, 0xaf, 0x41, 0x41, 0x9c, 0xec, 0xe1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x18, 0xf7,
	0x6a, 0xef, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScoresClient is the client API for Scores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScoresClient interface {
	// Lists a graph's scores.
	ListGraphScores(ctx context.Context, in *ListGraphScoresRequest, opts ...grpc.CallOption) (*ListGraphScoresResponse, error)
	// Gets a graph score.
	GetGraphScore(ctx context.Context, in *GetGraphScoreRequest, opts ...grpc.CallOption) (*Score, error)
	// Sets a graph score.
	SetGraphScore(ctx context.Context, in *SetGraphScoreRequest, opts ...grpc.CallOption) (*Score, error)
	// Sets a batch of graph scores.
	BatchSetGraphScores(ctx context.Context, in *BatchSetGraphScoresRequest, opts ...grpc.CallOption) (*BatchSetGraphScoresResponse, error)
	// Deletes a graph.
	DeleteGraph(ctx context.Context, in *DeleteGraphRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes a score.
	DeleteGraphScore(ctx context.Context, in *DeleteGraphScoreRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets graph scores in descending order by score value.
	TopGraphScores(ctx context.Context, in *TopGraphScoresRequest, opts ...grpc.CallOption) (*TopGraphScoresResponse, error)
}

type scoresClient struct {
	cc *grpc.ClientConn
}

func NewScoresClient(cc *grpc.ClientConn) ScoresClient {
	return &scoresClient{cc}
}

func (c *scoresClient) ListGraphScores(ctx context.Context, in *ListGraphScoresRequest, opts ...grpc.CallOption) (*ListGraphScoresResponse, error) {
	out := new(ListGraphScoresResponse)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/ListGraphScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoresClient) GetGraphScore(ctx context.Context, in *GetGraphScoreRequest, opts ...grpc.CallOption) (*Score, error) {
	out := new(Score)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/GetGraphScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoresClient) SetGraphScore(ctx context.Context, in *SetGraphScoreRequest, opts ...grpc.CallOption) (*Score, error) {
	out := new(Score)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/SetGraphScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoresClient) BatchSetGraphScores(ctx context.Context, in *BatchSetGraphScoresRequest, opts ...grpc.CallOption) (*BatchSetGraphScoresResponse, error) {
	out := new(BatchSetGraphScoresResponse)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/BatchSetGraphScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoresClient) DeleteGraph(ctx context.Context, in *DeleteGraphRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/DeleteGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoresClient) DeleteGraphScore(ctx context.Context, in *DeleteGraphScoreRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/topos.scores.v1.Scores/DeleteGraphScore", in, out, opts...)
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
	// Lists a graph's scores.
	ListGraphScores(context.Context, *ListGraphScoresRequest) (*ListGraphScoresResponse, error)
	// Gets a graph score.
	GetGraphScore(context.Context, *GetGraphScoreRequest) (*Score, error)
	// Sets a graph score.
	SetGraphScore(context.Context, *SetGraphScoreRequest) (*Score, error)
	// Sets a batch of graph scores.
	BatchSetGraphScores(context.Context, *BatchSetGraphScoresRequest) (*BatchSetGraphScoresResponse, error)
	// Deletes a graph.
	DeleteGraph(context.Context, *DeleteGraphRequest) (*empty.Empty, error)
	// Deletes a score.
	DeleteGraphScore(context.Context, *DeleteGraphScoreRequest) (*empty.Empty, error)
	// Gets graph scores in descending order by score value.
	TopGraphScores(context.Context, *TopGraphScoresRequest) (*TopGraphScoresResponse, error)
}

// UnimplementedScoresServer can be embedded to have forward compatible implementations.
type UnimplementedScoresServer struct {
}

func (*UnimplementedScoresServer) ListGraphScores(ctx context.Context, req *ListGraphScoresRequest) (*ListGraphScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGraphScores not implemented")
}
func (*UnimplementedScoresServer) GetGraphScore(ctx context.Context, req *GetGraphScoreRequest) (*Score, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraphScore not implemented")
}
func (*UnimplementedScoresServer) SetGraphScore(ctx context.Context, req *SetGraphScoreRequest) (*Score, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGraphScore not implemented")
}
func (*UnimplementedScoresServer) BatchSetGraphScores(ctx context.Context, req *BatchSetGraphScoresRequest) (*BatchSetGraphScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSetGraphScores not implemented")
}
func (*UnimplementedScoresServer) DeleteGraph(ctx context.Context, req *DeleteGraphRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGraph not implemented")
}
func (*UnimplementedScoresServer) DeleteGraphScore(ctx context.Context, req *DeleteGraphScoreRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGraphScore not implemented")
}
func (*UnimplementedScoresServer) TopGraphScores(ctx context.Context, req *TopGraphScoresRequest) (*TopGraphScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopGraphScores not implemented")
}

func RegisterScoresServer(s *grpc.Server, srv ScoresServer) {
	s.RegisterService(&_Scores_serviceDesc, srv)
}

func _Scores_ListGraphScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGraphScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).ListGraphScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/ListGraphScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).ListGraphScores(ctx, req.(*ListGraphScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _Scores_SetGraphScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGraphScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).SetGraphScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/SetGraphScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).SetGraphScore(ctx, req.(*SetGraphScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scores_BatchSetGraphScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSetGraphScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).BatchSetGraphScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/BatchSetGraphScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).BatchSetGraphScores(ctx, req.(*BatchSetGraphScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scores_DeleteGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).DeleteGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/DeleteGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).DeleteGraph(ctx, req.(*DeleteGraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scores_DeleteGraphScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGraphScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoresServer).DeleteGraphScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topos.scores.v1.Scores/DeleteGraphScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoresServer).DeleteGraphScore(ctx, req.(*DeleteGraphScoreRequest))
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
			MethodName: "ListGraphScores",
			Handler:    _Scores_ListGraphScores_Handler,
		},
		{
			MethodName: "GetGraphScore",
			Handler:    _Scores_GetGraphScore_Handler,
		},
		{
			MethodName: "SetGraphScore",
			Handler:    _Scores_SetGraphScore_Handler,
		},
		{
			MethodName: "BatchSetGraphScores",
			Handler:    _Scores_BatchSetGraphScores_Handler,
		},
		{
			MethodName: "DeleteGraph",
			Handler:    _Scores_DeleteGraph_Handler,
		},
		{
			MethodName: "DeleteGraphScore",
			Handler:    _Scores_DeleteGraphScore_Handler,
		},
		{
			MethodName: "TopGraphScores",
			Handler:    _Scores_TopGraphScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topos/scores/v1/scores.proto",
}
