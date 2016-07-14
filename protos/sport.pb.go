// Code generated by protoc-gen-go.
// source: protos/sport.proto
// DO NOT EDIT!

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/sport.proto

It has these top-level messages:
	Sport
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Sport_DayOfWeek int32

const (
	Sport_SUNDAY    Sport_DayOfWeek = 0
	Sport_MONDAY    Sport_DayOfWeek = 1
	Sport_TUESDAY   Sport_DayOfWeek = 2
	Sport_WEDNESDAY Sport_DayOfWeek = 3
	Sport_THURSDAY  Sport_DayOfWeek = 4
	Sport_FRIDAY    Sport_DayOfWeek = 5
	Sport_SATURDAY  Sport_DayOfWeek = 6
)

var Sport_DayOfWeek_name = map[int32]string{
	0: "SUNDAY",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
}
var Sport_DayOfWeek_value = map[string]int32{
	"SUNDAY":    0,
	"MONDAY":    1,
	"TUESDAY":   2,
	"WEDNESDAY": 3,
	"THURSDAY":  4,
	"FRIDAY":    5,
	"SATURDAY":  6,
}

func (x Sport_DayOfWeek) String() string {
	return proto.EnumName(Sport_DayOfWeek_name, int32(x))
}
func (Sport_DayOfWeek) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Sport struct {
	Id        uint64          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string          `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	DayOfWeek Sport_DayOfWeek `protobuf:"varint,3,opt,name=dayOfWeek,enum=protos.Sport_DayOfWeek" json:"dayOfWeek,omitempty"`
	TimeOfDay uint32          `protobuf:"varint,4,opt,name=timeOfDay" json:"timeOfDay,omitempty"`
	Location  string          `protobuf:"bytes,5,opt,name=location" json:"location,omitempty"`
}

func (m *Sport) Reset()                    { *m = Sport{} }
func (m *Sport) String() string            { return proto.CompactTextString(m) }
func (*Sport) ProtoMessage()               {}
func (*Sport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Sport)(nil), "protos.Sport")
	proto.RegisterEnum("protos.Sport_DayOfWeek", Sport_DayOfWeek_name, Sport_DayOfWeek_value)
}

func init() { proto.RegisterFile("protos/sport.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x62, 0x4a,
	0xad, 0x4c, 0x5c, 0xac, 0xc1, 0x20, 0x71, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x96, 0x20, 0x20, 0x4b, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09,
	0x28, 0xc2, 0x19, 0x04, 0x66, 0x0b, 0x99, 0x72, 0x71, 0xa6, 0x24, 0x56, 0xfa, 0xa7, 0x85, 0xa7,
	0xa6, 0x66, 0x4b, 0x30, 0x03, 0x25, 0xf8, 0x8c, 0xc4, 0x21, 0x06, 0x16, 0xeb, 0x81, 0x4d, 0xd1,
	0x73, 0x81, 0x49, 0x07, 0x21, 0x54, 0x0a, 0xc9, 0x70, 0x71, 0x96, 0x64, 0xe6, 0xa6, 0xfa, 0xa7,
	0x01, 0x65, 0x25, 0x58, 0x80, 0xda, 0x78, 0x83, 0x10, 0x02, 0x42, 0x52, 0x5c, 0x1c, 0x39, 0xf9,
	0xc9, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0xac, 0x60, 0xcb, 0xe0, 0x7c, 0xa5, 0x74, 0x2e, 0x4e,
	0xb8, 0x89, 0x42, 0x5c, 0x5c, 0x6c, 0xc1, 0xa1, 0x7e, 0x2e, 0x8e, 0x91, 0x02, 0x0c, 0x20, 0xb6,
	0xaf, 0x3f, 0x98, 0xcd, 0x28, 0xc4, 0xcd, 0xc5, 0x1e, 0x12, 0xea, 0x1a, 0x0c, 0xe2, 0x30, 0x09,
	0xf1, 0x72, 0x71, 0x86, 0xbb, 0xba, 0xf8, 0x41, 0xb8, 0xcc, 0x42, 0x3c, 0x5c, 0x1c, 0x21, 0x1e,
	0xa1, 0x41, 0x60, 0x1e, 0x0b, 0x48, 0x97, 0x5b, 0x90, 0x27, 0x88, 0xcd, 0x0a, 0x92, 0x09, 0x76,
	0x0c, 0x09, 0x0d, 0x02, 0xf1, 0xd8, 0x92, 0x20, 0xe1, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x86, 0x37, 0x22, 0x2b, 0x2c, 0x01, 0x00, 0x00,
}