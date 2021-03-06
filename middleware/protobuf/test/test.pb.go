// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

/*
包名，通过protoc生成时go文件时
*/

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

// 手机类型
// 枚举类型第一个字段必须为0
type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}
var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}
func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_ae538fd892a6db9a, []int{0}
}

// 手机
type Phone struct {
	Type   PhoneType `protobuf:"varint,1,opt,name=type,enum=test.PhoneType" json:"type,omitempty"`
	Number string    `protobuf:"bytes,2,opt,name=number" json:"number,omitempty"`
	// XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	// XXX_unrecognized     []byte    `json:"-"`
	// XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ae538fd892a6db9a, []int{0}
}
func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (dst *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(dst, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

// 人
type Person struct {
	// 后面的数字表示标识号
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// repeated表示可重复
	// 可以有多个手机
	Phones               []*Phone `protobuf:"bytes,3,rep,name=phones" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ae538fd892a6db9a, []int{1}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

// 联系簿
type ContactBook struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContactBook) Reset()         { *m = ContactBook{} }
func (m *ContactBook) String() string { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()    {}
func (*ContactBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_ae538fd892a6db9a, []int{2}
}
func (m *ContactBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactBook.Unmarshal(m, b)
}
func (m *ContactBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactBook.Marshal(b, m, deterministic)
}
func (dst *ContactBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactBook.Merge(dst, src)
}
func (m *ContactBook) XXX_Size() int {
	return xxx_messageInfo_ContactBook.Size(m)
}
func (m *ContactBook) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactBook.DiscardUnknown(m)
}

var xxx_messageInfo_ContactBook proto.InternalMessageInfo

func (m *ContactBook) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func init() {
	proto.RegisterType((*Phone)(nil), "test.Phone")
	proto.RegisterType((*Person)(nil), "test.Person")
	proto.RegisterType((*ContactBook)(nil), "test.ContactBook")
	proto.RegisterEnum("test.PhoneType", PhoneType_name, PhoneType_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_ae538fd892a6db9a) }

var fileDescriptor_test_ae538fd892a6db9a = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x4d, 0x27, 0x13, 0x9d, 0x53, 0x19, 0x87, 0xb3, 0x90, 0xec, 0x2c, 0x2d, 0x48, 0x71,
	0xd1, 0x45, 0xc5, 0x17, 0xf0, 0x02, 0x82, 0x48, 0x6b, 0x10, 0x5c, 0xb7, 0x36, 0x60, 0x91, 0x26,
	0xa1, 0x89, 0x8b, 0xbe, 0xbd, 0xf4, 0x10, 0x65, 0x76, 0x7f, 0xf2, 0xfd, 0x97, 0x04, 0x20, 0x68,
	0x1f, 0x2a, 0x37, 0xdb, 0x60, 0x91, 0xaf, 0x3a, 0x7f, 0x84, 0x6d, 0xfb, 0x65, 0x8d, 0xc6, 0x02,
	0x78, 0x58, 0x9c, 0x96, 0x2c, 0x63, 0xe5, 0xbe, 0xbe, 0xa8, 0xc8, 0x49, 0xe8, 0x7d, 0x71, 0x5a,
	0x11, 0xc4, 0x4b, 0x10, 0xe6, 0x67, 0xea, 0xf5, 0x2c, 0x93, 0x8c, 0x95, 0x3b, 0x15, 0x4f, 0xf9,
	0x1b, 0x88, 0x56, 0xcf, 0xde, 0x1a, 0xdc, 0x43, 0x32, 0x0e, 0x54, 0xb2, 0x55, 0xc9, 0x38, 0x20,
	0x02, 0x37, 0xdd, 0xa4, 0xa3, 0x9f, 0x34, 0x16, 0x20, 0xdc, 0x5a, 0xec, 0xe5, 0x26, 0xdb, 0x94,
	0x69, 0x9d, 0x1e, 0x8d, 0xa9, 0x88, 0xf2, 0x3b, 0x48, 0x1f, 0xac, 0x09, 0xdd, 0x67, 0xb8, 0xb7,
	0xf6, 0x1b, 0xaf, 0xe1, 0xd4, 0xd1, 0x82, 0x97, 0x8c, 0x42, 0xe7, 0x31, 0x44, 0x97, 0xea, 0x0f,
	0xde, 0x5c, 0xc1, 0xee, 0xff, 0xd1, 0x78, 0x06, 0xfc, 0xb9, 0x79, 0x7d, 0x3a, 0x9c, 0xac, 0xea,
	0xa3, 0x51, 0x2f, 0x07, 0xd6, 0x0b, 0xfa, 0xfd, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec,
	0x73, 0xbe, 0x18, 0x0b, 0x01, 0x00, 0x00,
}
