// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: struct.proto

package tablestore

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/gogo/protobuf/types"
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

// `NullValue` is a singleton enumeration to represent the null value for the
// `Value` type union.
//
//  The JSON representation for `NullValue` is JSON `null`.
type NullValue int32

const (
	// Null value.
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}

var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}

func (NullValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0605f6bcb0ae6db1, []int{0}
}

// `Struct` represents a structured data value, consisting of fields
// which map to dynamically typed values. In some languages, `Struct`
// might be supported by a native representation. For example, in
// scripting languages like JS a struct is represented as an
// object. The details of that representation are described together
// with the proto support for the language.
//
// The JSON representation for `Struct` is JSON object.
type Struct struct {
	// Unordered map of dynamically typed values.
	Fields               map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Struct) Reset()         { *m = Struct{} }
func (m *Struct) String() string { return proto.CompactTextString(m) }
func (*Struct) ProtoMessage()    {}
func (*Struct) Descriptor() ([]byte, []int) {
	return fileDescriptor_0605f6bcb0ae6db1, []int{0}
}
func (m *Struct) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Struct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Struct.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Struct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Struct.Merge(m, src)
}
func (m *Struct) XXX_Size() int {
	return m.Size()
}
func (m *Struct) XXX_DiscardUnknown() {
	xxx_messageInfo_Struct.DiscardUnknown(m)
}

var xxx_messageInfo_Struct proto.InternalMessageInfo

func (m *Struct) GetFields() map[string]*Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

// `Value` represents a dynamically typed value which can be either
// null, a number, a string, a boolean, a recursive struct value, or a
// list of values. A producer of value is expected to set one of that
// variants, absence of any variant indicates an error.
//
// The JSON representation for `Value` is JSON value.
type Value struct {
	// The kind of value.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_StructValue
	//	*Value_ListValue
	//	*Value_IntegerValue
	//	*Value_BytesValue
	//	*Value_TimestampValue
	Kind                 isValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_0605f6bcb0ae6db1, []int{1}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Value.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return m.Size()
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Kind interface {
	isValue_Kind()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=tablestore.NullValue,oneof" json:"null_value,omitempty"`
}
type Value_NumberValue struct {
	NumberValue float64 `protobuf:"fixed64,2,opt,name=number_value,json=numberValue,proto3,oneof" json:"number_value,omitempty"`
}
type Value_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof" json:"string_value,omitempty"`
}
type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof" json:"bool_value,omitempty"`
}
type Value_StructValue struct {
	StructValue *Struct `protobuf:"bytes,5,opt,name=struct_value,json=structValue,proto3,oneof" json:"struct_value,omitempty"`
}
type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,6,opt,name=list_value,json=listValue,proto3,oneof" json:"list_value,omitempty"`
}
type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,7,opt,name=integer_value,json=integerValue,proto3,oneof" json:"integer_value,omitempty"`
}
type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,8,opt,name=bytes_value,json=bytesValue,proto3,oneof" json:"bytes_value,omitempty"`
}
type Value_TimestampValue struct {
	TimestampValue *types.Timestamp `protobuf:"bytes,9,opt,name=timestamp_value,json=timestampValue,proto3,oneof" json:"timestamp_value,omitempty"`
}

func (*Value_NullValue) isValue_Kind()      {}
func (*Value_NumberValue) isValue_Kind()    {}
func (*Value_StringValue) isValue_Kind()    {}
func (*Value_BoolValue) isValue_Kind()      {}
func (*Value_StructValue) isValue_Kind()    {}
func (*Value_ListValue) isValue_Kind()      {}
func (*Value_IntegerValue) isValue_Kind()   {}
func (*Value_BytesValue) isValue_Kind()     {}
func (*Value_TimestampValue) isValue_Kind() {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (m *Value) GetNumberValue() float64 {
	if x, ok := m.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetStructValue() *Struct {
	if x, ok := m.GetKind().(*Value_StructValue); ok {
		return x.StructValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (m *Value) GetIntegerValue() int64 {
	if x, ok := m.GetKind().(*Value_IntegerValue); ok {
		return x.IntegerValue
	}
	return 0
}

func (m *Value) GetBytesValue() []byte {
	if x, ok := m.GetKind().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Value) GetTimestampValue() *types.Timestamp {
	if x, ok := m.GetKind().(*Value_TimestampValue); ok {
		return x.TimestampValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_TimestampValue)(nil),
	}
}

// `ListValue` is a wrapper around a repeated field of values.
//
// The JSON representation for `ListValue` is JSON array.
type ListValue struct {
	// Repeated field of dynamically typed values.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0605f6bcb0ae6db1, []int{2}
}
func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(m, src)
}
func (m *ListValue) XXX_Size() int {
	return m.Size()
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("tablestore.NullValue", NullValue_name, NullValue_value)
	proto.RegisterType((*Struct)(nil), "tablestore.Struct")
	proto.RegisterMapType((map[string]*Value)(nil), "tablestore.Struct.FieldsEntry")
	proto.RegisterType((*Value)(nil), "tablestore.Value")
	proto.RegisterType((*ListValue)(nil), "tablestore.ListValue")
}

func init() { proto.RegisterFile("struct.proto", fileDescriptor_0605f6bcb0ae6db1) }

var fileDescriptor_0605f6bcb0ae6db1 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x33, 0x9b, 0x6d, 0xdc, 0x9c, 0xc4, 0x5a, 0x07, 0x84, 0x52, 0x21, 0x1b, 0x2b, 0x62,
	0xf4, 0x22, 0x0b, 0x15, 0xaa, 0x78, 0xe7, 0x42, 0x25, 0x17, 0x61, 0x2f, 0x46, 0x77, 0x6f, 0x4b,
	0xb3, 0x3b, 0x1b, 0x42, 0xa7, 0x49, 0xc9, 0x4c, 0x84, 0xbe, 0x84, 0xd7, 0x3e, 0x88, 0x0f, 0xe1,
	0xa5, 0x8f, 0x20, 0xf5, 0x45, 0x24, 0xf3, 0xaf, 0x15, 0xbd, 0x4b, 0xbe, 0xf3, 0xfb, 0xbe, 0x39,
	0x67, 0xe6, 0x40, 0xc8, 0x45, 0xdb, 0xdd, 0x8a, 0x74, 0xdb, 0x36, 0xa2, 0xc1, 0x20, 0x56, 0x05,
	0xa3, 0x5c, 0x34, 0x2d, 0x9d, 0x9c, 0x97, 0x4d, 0x53, 0x32, 0x7a, 0x21, 0x2b, 0x45, 0x77, 0x7f,
	0x21, 0xaa, 0x0d, 0xe5, 0x62, 0xb5, 0xd9, 0x2a, 0x78, 0xfa, 0x15, 0x81, 0xf7, 0x49, 0xba, 0xf1,
	0x1c, 0xbc, 0xfb, 0x8a, 0xb2, 0x3b, 0x3e, 0x46, 0xb1, 0x9b, 0x04, 0xb3, 0x28, 0x3d, 0x04, 0xa5,
	0x8a, 0x49, 0x3f, 0x4a, 0x60, 0x51, 0x8b, 0x76, 0x47, 0x34, 0x3d, 0xc9, 0x21, 0x38, 0x92, 0xf1,
	0x08, 0xdc, 0x35, 0xdd, 0x8d, 0x51, 0x8c, 0x12, 0x9f, 0xf4, 0x9f, 0xf8, 0x25, 0x0c, 0xbe, 0xac,
	0x58, 0x47, 0xc7, 0x27, 0x31, 0x4a, 0x82, 0xd9, 0xe3, 0xe3, 0xdc, 0x9b, 0xbe, 0x40, 0x54, 0xfd,
	0xfd, 0xc9, 0x3b, 0x34, 0xfd, 0xee, 0xc2, 0x40, 0x8a, 0x78, 0x0e, 0x50, 0x77, 0x8c, 0x2d, 0x95,
	0xb7, 0xcf, 0x1b, 0xce, 0x9e, 0x1c, 0x7b, 0xaf, 0x3a, 0xc6, 0x24, 0x9a, 0x39, 0xc4, 0xaf, 0xcd,
	0x0f, 0x7e, 0x0e, 0x61, 0xdd, 0x6d, 0x0a, 0xda, 0x2e, 0x0f, 0xa7, 0xa2, 0xcc, 0x21, 0x81, 0x52,
	0x2d, 0xc4, 0x45, 0x5b, 0xd5, 0xa5, 0x86, 0xdc, 0xbe, 0xdd, 0x1e, 0x52, 0xaa, 0x82, 0xce, 0x01,
	0x8a, 0xa6, 0x31, 0x1d, 0x9c, 0xc6, 0x28, 0x39, 0xeb, 0x8f, 0xea, 0x35, 0x05, 0xbc, 0x35, 0x57,
	0xaf, 0x91, 0x81, 0x1c, 0x10, 0xff, 0x7b, 0x71, 0x3a, 0xb9, 0xbb, 0x15, 0x76, 0x36, 0x56, 0x71,
	0x63, 0xf3, 0xa4, 0xed, 0xaf, 0xd9, 0xf2, 0x8a, 0x0b, 0x3b, 0x1b, 0x33, 0x3f, 0xf8, 0x05, 0x3c,
	0xac, 0x6a, 0x41, 0x4b, 0x3b, 0xdc, 0x83, 0x18, 0x25, 0x6e, 0xe6, 0x90, 0x50, 0xcb, 0x0a, 0x7b,
	0x06, 0x41, 0xb1, 0x13, 0x94, 0x6b, 0xe8, 0x2c, 0x46, 0x49, 0x98, 0x39, 0x04, 0xa4, 0xa8, 0x90,
	0x05, 0x3c, 0xb2, 0xbb, 0xa0, 0x31, 0x5f, 0xb6, 0x31, 0x49, 0xd5, 0xce, 0xa4, 0x66, 0x67, 0xd2,
	0xcf, 0x86, 0xcb, 0x1c, 0x32, 0xb4, 0x26, 0x19, 0x73, 0xe9, 0xc1, 0xe9, 0xba, 0xaa, 0xef, 0xa6,
	0x73, 0xf0, 0x6d, 0xcb, 0xf8, 0x15, 0x78, 0x32, 0xd1, 0x6c, 0xd2, 0x7f, 0x5e, 0x5c, 0x03, 0xaf,
	0x9f, 0x82, 0x6f, 0x9f, 0x11, 0x0f, 0x01, 0xae, 0xae, 0xf3, 0x7c, 0x79, 0xf3, 0x21, 0xbf, 0x5e,
	0x8c, 0x9c, 0xcb, 0xd1, 0x8f, 0x7d, 0x84, 0x7e, 0xee, 0x23, 0xf4, 0x6b, 0x1f, 0xa1, 0x6f, 0xbf,
	0x23, 0xa7, 0xf0, 0x64, 0x53, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x51, 0xb0, 0xdf,
	0xf2, 0x02, 0x00, 0x00,
}

func (m *Struct) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Struct) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Struct) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Fields) > 0 {
		for k := range m.Fields {
			v := m.Fields[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintStruct(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintStruct(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintStruct(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Value) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Kind != nil {
		{
			size := m.Kind.Size()
			i -= size
			if _, err := m.Kind.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Value_NullValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_NullValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintStruct(dAtA, i, uint64(m.NullValue))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *Value_NumberValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_NumberValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.NumberValue))))
	i--
	dAtA[i] = 0x11
	return len(dAtA) - i, nil
}
func (m *Value_StringValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_StringValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.StringValue)
	copy(dAtA[i:], m.StringValue)
	i = encodeVarintStruct(dAtA, i, uint64(len(m.StringValue)))
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}
func (m *Value_BoolValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_BoolValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i--
	if m.BoolValue {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x20
	return len(dAtA) - i, nil
}
func (m *Value_StructValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_StructValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StructValue != nil {
		{
			size, err := m.StructValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStruct(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *Value_ListValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_ListValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ListValue != nil {
		{
			size, err := m.ListValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStruct(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *Value_IntegerValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_IntegerValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarintStruct(dAtA, i, uint64(m.IntegerValue))
	i--
	dAtA[i] = 0x38
	return len(dAtA) - i, nil
}
func (m *Value_BytesValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_BytesValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BytesValue != nil {
		i -= len(m.BytesValue)
		copy(dAtA[i:], m.BytesValue)
		i = encodeVarintStruct(dAtA, i, uint64(len(m.BytesValue)))
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *Value_TimestampValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Value_TimestampValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TimestampValue != nil {
		{
			size, err := m.TimestampValue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStruct(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *ListValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListValue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListValue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Values[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStruct(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintStruct(dAtA []byte, offset int, v uint64) int {
	offset -= sovStruct(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Struct) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fields) > 0 {
		for k, v := range m.Fields {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovStruct(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovStruct(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovStruct(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Value) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Kind != nil {
		n += m.Kind.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Value_NullValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStruct(uint64(m.NullValue))
	return n
}
func (m *Value_NumberValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}
func (m *Value_StringValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringValue)
	n += 1 + l + sovStruct(uint64(l))
	return n
}
func (m *Value_BoolValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 2
	return n
}
func (m *Value_StructValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StructValue != nil {
		l = m.StructValue.Size()
		n += 1 + l + sovStruct(uint64(l))
	}
	return n
}
func (m *Value_ListValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListValue != nil {
		l = m.ListValue.Size()
		n += 1 + l + sovStruct(uint64(l))
	}
	return n
}
func (m *Value_IntegerValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovStruct(uint64(m.IntegerValue))
	return n
}
func (m *Value_BytesValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BytesValue != nil {
		l = len(m.BytesValue)
		n += 1 + l + sovStruct(uint64(l))
	}
	return n
}
func (m *Value_TimestampValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimestampValue != nil {
		l = m.TimestampValue.Size()
		n += 1 + l + sovStruct(uint64(l))
	}
	return n
}
func (m *ListValue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovStruct(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStruct(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStruct(x uint64) (n int) {
	return sovStruct(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Struct) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Struct: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Struct: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fields == nil {
				m.Fields = make(map[string]*Value)
			}
			var mapkey string
			var mapvalue *Value
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStruct
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStruct
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthStruct
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthStruct
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStruct
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthStruct
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthStruct
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Value{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipStruct(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthStruct
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Fields[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStruct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStruct
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStruct
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Value) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NullValue", wireType)
			}
			var v NullValue
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= NullValue(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Kind = &Value_NullValue{v}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Kind = &Value_NumberValue{float64(math.Float64frombits(v))}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = &Value_StringValue{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Kind = &Value_BoolValue{b}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StructValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Struct{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &Value_StructValue{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ListValue{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &Value_ListValue{v}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntegerValue", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Kind = &Value_IntegerValue{v}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Kind = &Value_BytesValue{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimestampValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Timestamp{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Kind = &Value_TimestampValue{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStruct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStruct
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStruct
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStruct
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStruct
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &Value{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStruct(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStruct
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStruct
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStruct(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStruct
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStruct
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStruct
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStruct
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStruct
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStruct        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStruct          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStruct = fmt.Errorf("proto: unexpected end of group")
)