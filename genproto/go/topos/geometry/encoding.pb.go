// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topos/geometry/encoding.proto

package geometry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Encoding int32

const (
	Encoding_WKB     Encoding = 0
	Encoding_GEOJSON Encoding = 1
)

var Encoding_name = map[int32]string{
	0: "WKB",
	1: "GEOJSON",
}

var Encoding_value = map[string]int32{
	"WKB":     0,
	"GEOJSON": 1,
}

func (x Encoding) String() string {
	return proto.EnumName(Encoding_name, int32(x))
}

func (Encoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_326db63c67ecb653, []int{0}
}

func init() {
	proto.RegisterEnum("topos.geometry.Encoding", Encoding_name, Encoding_value)
}

func init() { proto.RegisterFile("topos/geometry/encoding.proto", fileDescriptor_326db63c67ecb653) }

var fileDescriptor_326db63c67ecb653 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xc9, 0x2f, 0xc8,
	0x2f, 0xd6, 0x4f, 0x4f, 0xcd, 0xcf, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x4f, 0xcd, 0x4b, 0xce, 0x4f,
	0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x4b, 0xeb, 0xc1, 0xa4,
	0xb5, 0x14, 0xb8, 0x38, 0x5c, 0xa1, 0x2a, 0x84, 0xd8, 0xb9, 0x98, 0xc3, 0xbd, 0x9d, 0x04, 0x18,
	0x84, 0xb8, 0xb9, 0xd8, 0xdd, 0x5d, 0xfd, 0xbd, 0x82, 0xfd, 0xfd, 0x04, 0x18, 0x9d, 0xac, 0xa3,
	0x2c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xc1, 0xda, 0x75, 0x13,
	0x33, 0x61, 0x8c, 0x82, 0x4c, 0x90, 0x5d, 0x79, 0x60, 0xc3, 0xf5, 0xd3, 0xf3, 0xf5, 0x51, 0x6d,
	0x4f, 0x62, 0x03, 0x4b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0xd0, 0x07, 0x27, 0x96,
	0x00, 0x00, 0x00,
}
