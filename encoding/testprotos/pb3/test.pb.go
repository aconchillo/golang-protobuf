// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb3/test.proto

package pb3

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
	Enum_TEN  Enum = 10
)

func (e Enum) Type() protoreflect.EnumType {
	return xxx_File_pb3_test_proto_enumTypes[0]
}
func (e Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

// Deprecated: Use Enum.Type.Values instead.
var Enum_name = map[int32]string{
	0:  "ZERO",
	1:  "ONE",
	2:  "TWO",
	10: "TEN",
}

// Deprecated: Use Enum.Type.Values instead.
var Enum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
	"TWO":  2,
	"TEN":  10,
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

// Deprecated: Use Enum.Type instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{0}
}

type Enums_NestedEnum int32

const (
	Enums_CERO Enums_NestedEnum = 0
	Enums_UNO  Enums_NestedEnum = 1
	Enums_DOS  Enums_NestedEnum = 2
	Enums_DIEZ Enums_NestedEnum = 10
)

func (e Enums_NestedEnum) Type() protoreflect.EnumType {
	return xxx_File_pb3_test_proto_enumTypes[1]
}
func (e Enums_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

// Deprecated: Use Enums_NestedEnum.Type.Values instead.
var Enums_NestedEnum_name = map[int32]string{
	0:  "CERO",
	1:  "UNO",
	2:  "DOS",
	10: "DIEZ",
}

// Deprecated: Use Enums_NestedEnum.Type.Values instead.
var Enums_NestedEnum_value = map[string]int32{
	"CERO": 0,
	"UNO":  1,
	"DOS":  2,
	"DIEZ": 10,
}

func (x Enums_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

// Deprecated: Use Enums_NestedEnum.Type instead.
func (Enums_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{1, 0}
}

// Scalars contains scalar field types.
type Scalars struct {
	SBool                bool     `protobuf:"varint,1,opt,name=s_bool,json=sBool,proto3" json:"s_bool,omitempty"`
	SInt32               int32    `protobuf:"varint,2,opt,name=s_int32,json=sInt32,proto3" json:"s_int32,omitempty"`
	SInt64               int64    `protobuf:"varint,3,opt,name=s_int64,json=sInt64,proto3" json:"s_int64,omitempty"`
	SUint32              uint32   `protobuf:"varint,4,opt,name=s_uint32,json=sUint32,proto3" json:"s_uint32,omitempty"`
	SUint64              uint64   `protobuf:"varint,5,opt,name=s_uint64,json=sUint64,proto3" json:"s_uint64,omitempty"`
	SSint32              int32    `protobuf:"zigzag32,6,opt,name=s_sint32,json=sSint32,proto3" json:"s_sint32,omitempty"`
	SSint64              int64    `protobuf:"zigzag64,7,opt,name=s_sint64,json=sSint64,proto3" json:"s_sint64,omitempty"`
	SFixed32             uint32   `protobuf:"fixed32,8,opt,name=s_fixed32,json=sFixed32,proto3" json:"s_fixed32,omitempty"`
	SFixed64             uint64   `protobuf:"fixed64,9,opt,name=s_fixed64,json=sFixed64,proto3" json:"s_fixed64,omitempty"`
	SSfixed32            int32    `protobuf:"fixed32,10,opt,name=s_sfixed32,json=sSfixed32,proto3" json:"s_sfixed32,omitempty"`
	SSfixed64            int64    `protobuf:"fixed64,11,opt,name=s_sfixed64,json=sSfixed64,proto3" json:"s_sfixed64,omitempty"`
	SFloat               float32  `protobuf:"fixed32,20,opt,name=s_float,json=sFloat,proto3" json:"s_float,omitempty"`
	SDouble              float64  `protobuf:"fixed64,21,opt,name=s_double,json=sDouble,proto3" json:"s_double,omitempty"`
	SBytes               []byte   `protobuf:"bytes,14,opt,name=s_bytes,json=sBytes,proto3" json:"s_bytes,omitempty"`
	SString              string   `protobuf:"bytes,13,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scalars) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[0].MessageOf(m)
}
func (m *Scalars) Reset()         { *m = Scalars{} }
func (m *Scalars) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Scalars) ProtoMessage()    {}

// Deprecated: Use Scalars.ProtoReflect.Type instead.
func (*Scalars) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{0}
}

func (m *Scalars) GetSBool() bool {
	if m != nil {
		return m.SBool
	}
	return false
}

func (m *Scalars) GetSInt32() int32 {
	if m != nil {
		return m.SInt32
	}
	return 0
}

func (m *Scalars) GetSInt64() int64 {
	if m != nil {
		return m.SInt64
	}
	return 0
}

func (m *Scalars) GetSUint32() uint32 {
	if m != nil {
		return m.SUint32
	}
	return 0
}

func (m *Scalars) GetSUint64() uint64 {
	if m != nil {
		return m.SUint64
	}
	return 0
}

func (m *Scalars) GetSSint32() int32 {
	if m != nil {
		return m.SSint32
	}
	return 0
}

func (m *Scalars) GetSSint64() int64 {
	if m != nil {
		return m.SSint64
	}
	return 0
}

func (m *Scalars) GetSFixed32() uint32 {
	if m != nil {
		return m.SFixed32
	}
	return 0
}

func (m *Scalars) GetSFixed64() uint64 {
	if m != nil {
		return m.SFixed64
	}
	return 0
}

func (m *Scalars) GetSSfixed32() int32 {
	if m != nil {
		return m.SSfixed32
	}
	return 0
}

func (m *Scalars) GetSSfixed64() int64 {
	if m != nil {
		return m.SSfixed64
	}
	return 0
}

func (m *Scalars) GetSFloat() float32 {
	if m != nil {
		return m.SFloat
	}
	return 0
}

func (m *Scalars) GetSDouble() float64 {
	if m != nil {
		return m.SDouble
	}
	return 0
}

func (m *Scalars) GetSBytes() []byte {
	if m != nil {
		return m.SBytes
	}
	return nil
}

func (m *Scalars) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

// Message contains enum fields.
type Enums struct {
	SEnum                Enum             `protobuf:"varint,1,opt,name=s_enum,json=sEnum,proto3,enum=pb3.Enum" json:"s_enum,omitempty"`
	SNestedEnum          Enums_NestedEnum `protobuf:"varint,3,opt,name=s_nested_enum,json=sNestedEnum,proto3,enum=pb3.Enums_NestedEnum" json:"s_nested_enum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Enums) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[1].MessageOf(m)
}
func (m *Enums) Reset()         { *m = Enums{} }
func (m *Enums) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Enums) ProtoMessage()    {}

// Deprecated: Use Enums.ProtoReflect.Type instead.
func (*Enums) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{1}
}

func (m *Enums) GetSEnum() Enum {
	if m != nil {
		return m.SEnum
	}
	return Enum_ZERO
}

func (m *Enums) GetSNestedEnum() Enums_NestedEnum {
	if m != nil {
		return m.SNestedEnum
	}
	return Enums_CERO
}

// Message contains nested message field.
type Nests struct {
	SNested              *Nested  `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nests) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[2].MessageOf(m)
}
func (m *Nests) Reset()         { *m = Nests{} }
func (m *Nests) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Nests) ProtoMessage()    {}

// Deprecated: Use Nests.ProtoReflect.Type instead.
func (*Nests) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{2}
}

func (m *Nests) GetSNested() *Nested {
	if m != nil {
		return m.SNested
	}
	return nil
}

// Message type used as submessage.
type Nested struct {
	SString              string   `protobuf:"bytes,1,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	SNested              *Nested  `protobuf:"bytes,2,opt,name=s_nested,json=sNested,proto3" json:"s_nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[3].MessageOf(m)
}
func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Nested) ProtoMessage()    {}

// Deprecated: Use Nested.ProtoReflect.Type instead.
func (*Nested) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{3}
}

func (m *Nested) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

func (m *Nested) GetSNested() *Nested {
	if m != nil {
		return m.SNested
	}
	return nil
}

// Message contains oneof field.
type Oneofs struct {
	// Types that are valid to be assigned to Union:
	//	*Oneofs_OneofEnum
	//	*Oneofs_OneofString
	//	*Oneofs_OneofNested
	Union                isOneofs_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Oneofs) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[4].MessageOf(m)
}
func (m *Oneofs) Reset()         { *m = Oneofs{} }
func (m *Oneofs) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Oneofs) ProtoMessage()    {}

// Deprecated: Use Oneofs.ProtoReflect.Type instead.
func (*Oneofs) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{4}
}

type isOneofs_Union interface {
	isOneofs_Union()
}

type Oneofs_OneofEnum struct {
	OneofEnum Enum `protobuf:"varint,1,opt,name=oneof_enum,json=oneofEnum,proto3,enum=pb3.Enum,oneof"`
}

type Oneofs_OneofString struct {
	OneofString string `protobuf:"bytes,2,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type Oneofs_OneofNested struct {
	OneofNested *Nested `protobuf:"bytes,3,opt,name=oneof_nested,json=oneofNested,proto3,oneof"`
}

func (*Oneofs_OneofEnum) isOneofs_Union() {}

func (*Oneofs_OneofString) isOneofs_Union() {}

func (*Oneofs_OneofNested) isOneofs_Union() {}

func (m *Oneofs) GetUnion() isOneofs_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Oneofs) GetOneofEnum() Enum {
	if x, ok := m.GetUnion().(*Oneofs_OneofEnum); ok {
		return x.OneofEnum
	}
	return Enum_ZERO
}

func (m *Oneofs) GetOneofString() string {
	if x, ok := m.GetUnion().(*Oneofs_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (m *Oneofs) GetOneofNested() *Nested {
	if x, ok := m.GetUnion().(*Oneofs_OneofNested); ok {
		return x.OneofNested
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Oneofs) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Oneofs_OneofEnum)(nil),
		(*Oneofs_OneofString)(nil),
		(*Oneofs_OneofNested)(nil),
	}
}

// Message contains map fields.
type Maps struct {
	Int32ToStr           map[int32]string   `protobuf:"bytes,1,rep,name=int32_to_str,json=int32ToStr,proto3" json:"int32_to_str,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BoolToUint32         map[bool]uint32    `protobuf:"bytes,2,rep,name=bool_to_uint32,json=boolToUint32,proto3" json:"bool_to_uint32,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Uint64ToEnum         map[uint64]Enum    `protobuf:"bytes,3,rep,name=uint64_to_enum,json=uint64ToEnum,proto3" json:"uint64_to_enum,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=pb3.Enum"`
	StrToNested          map[string]*Nested `protobuf:"bytes,4,rep,name=str_to_nested,json=strToNested,proto3" json:"str_to_nested,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StrToOneofs          map[string]*Oneofs `protobuf:"bytes,5,rep,name=str_to_oneofs,json=strToOneofs,proto3" json:"str_to_oneofs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Maps) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[5].MessageOf(m)
}
func (m *Maps) Reset()         { *m = Maps{} }
func (m *Maps) String() string { return protoimpl.X.MessageStringOf(m) }
func (*Maps) ProtoMessage()    {}

// Deprecated: Use Maps.ProtoReflect.Type instead.
func (*Maps) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{5}
}

func (m *Maps) GetInt32ToStr() map[int32]string {
	if m != nil {
		return m.Int32ToStr
	}
	return nil
}

func (m *Maps) GetBoolToUint32() map[bool]uint32 {
	if m != nil {
		return m.BoolToUint32
	}
	return nil
}

func (m *Maps) GetUint64ToEnum() map[uint64]Enum {
	if m != nil {
		return m.Uint64ToEnum
	}
	return nil
}

func (m *Maps) GetStrToNested() map[string]*Nested {
	if m != nil {
		return m.StrToNested
	}
	return nil
}

func (m *Maps) GetStrToOneofs() map[string]*Oneofs {
	if m != nil {
		return m.StrToOneofs
	}
	return nil
}

// Message for testing json_name option.
type JSONNames struct {
	SString              string   `protobuf:"bytes,1,opt,name=s_string,json=foo_bar,proto3" json:"s_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JSONNames) ProtoReflect() protoreflect.Message {
	return xxx_File_pb3_test_proto_messageTypes[6].MessageOf(m)
}
func (m *JSONNames) Reset()         { *m = JSONNames{} }
func (m *JSONNames) String() string { return protoimpl.X.MessageStringOf(m) }
func (*JSONNames) ProtoMessage()    {}

// Deprecated: Use JSONNames.ProtoReflect.Type instead.
func (*JSONNames) Descriptor() ([]byte, []int) {
	return xxx_File_pb3_test_proto_rawDescGZIP(), []int{6}
}

func (m *JSONNames) GetSString() string {
	if m != nil {
		return m.SString
	}
	return ""
}

var File_pb3_test_proto protoreflect.FileDescriptor

var xxx_File_pb3_test_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x70, 0x62, 0x33, 0x22, 0x9e, 0x03, 0x0a, 0x07, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x73, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x73, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f,
	0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x07, 0x73, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x73,
	0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x08, 0x20, 0x01, 0x28, 0x07, 0x52, 0x08, 0x73, 0x46, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x73, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x09, 0x73, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x10, 0x52, 0x09, 0x73, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x73, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x73, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x98, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x12, 0x20, 0x0a, 0x06, 0x73, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x73, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x39, 0x0a, 0x0d, 0x73, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x0b, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x32, 0x0a,
	0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x43,
	0x45, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x4e, 0x4f, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x44, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x45, 0x5a, 0x10,
	0x0a, 0x22, 0x2f, 0x0a, 0x05, 0x4e, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x73, 0x5f,
	0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x07, 0x73, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x22, 0x4b, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x08, 0x73, 0x5f, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x07, 0x73, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22,
	0x94, 0x01, 0x0a, 0x06, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x0c, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x07, 0x0a,
	0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0xaf, 0x05, 0x0a, 0x04, 0x4d, 0x61, 0x70, 0x73, 0x12,
	0x3b, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4d, 0x61, 0x70, 0x73,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x12, 0x41, 0x0a, 0x0e,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x41, 0x0a, 0x0e, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4d, 0x61,
	0x70, 0x73, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x45, 0x6e, 0x75, 0x6d, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x33, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x6c, 0x54, 0x6f, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x4a, 0x0a, 0x11, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x54, 0x6f, 0x45, 0x6e,
	0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b,
	0x0a, 0x10, 0x53, 0x74, 0x72, 0x54, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x10, 0x53,
	0x74, 0x72, 0x54, 0x6f, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x33, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x09, 0x4a, 0x53, 0x4f, 0x4e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x6f, 0x5f, 0x62, 0x61, 0x72,
	0x2a, 0x2b, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54,
	0x57, 0x4f, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x45, 0x4e, 0x10, 0x0a, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x70, 0x62, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_pb3_test_proto_rawDesc_once sync.Once
	xxx_File_pb3_test_proto_rawDesc_data = xxx_File_pb3_test_proto_rawDesc
)

func xxx_File_pb3_test_proto_rawDescGZIP() []byte {
	xxx_File_pb3_test_proto_rawDesc_once.Do(func() {
		xxx_File_pb3_test_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_pb3_test_proto_rawDesc_data)
	})
	return xxx_File_pb3_test_proto_rawDesc_data
}

var xxx_File_pb3_test_proto_enumTypes = make([]protoreflect.EnumType, 2)
var xxx_File_pb3_test_proto_messageTypes = make([]protoimpl.MessageType, 12)
var xxx_File_pb3_test_proto_goTypes = []interface{}{
	(Enum)(0),             // 0: pb3.Enum
	(Enums_NestedEnum)(0), // 1: pb3.Enums.NestedEnum
	(*Scalars)(nil),       // 2: pb3.Scalars
	(*Enums)(nil),         // 3: pb3.Enums
	(*Nests)(nil),         // 4: pb3.Nests
	(*Nested)(nil),        // 5: pb3.Nested
	(*Oneofs)(nil),        // 6: pb3.Oneofs
	(*Maps)(nil),          // 7: pb3.Maps
	(*JSONNames)(nil),     // 8: pb3.JSONNames
	nil,                   // 9: pb3.Maps.Int32ToStrEntry
	nil,                   // 10: pb3.Maps.BoolToUint32Entry
	nil,                   // 11: pb3.Maps.Uint64ToEnumEntry
	nil,                   // 12: pb3.Maps.StrToNestedEntry
	nil,                   // 13: pb3.Maps.StrToOneofsEntry
}
var xxx_File_pb3_test_proto_depIdxs = []int32{
	0,  // pb3.Enums.s_enum:type_name -> pb3.Enum
	1,  // pb3.Enums.s_nested_enum:type_name -> pb3.Enums.NestedEnum
	5,  // pb3.Nests.s_nested:type_name -> pb3.Nested
	5,  // pb3.Nested.s_nested:type_name -> pb3.Nested
	0,  // pb3.Oneofs.oneof_enum:type_name -> pb3.Enum
	5,  // pb3.Oneofs.oneof_nested:type_name -> pb3.Nested
	9,  // pb3.Maps.int32_to_str:type_name -> pb3.Maps.Int32ToStrEntry
	10, // pb3.Maps.bool_to_uint32:type_name -> pb3.Maps.BoolToUint32Entry
	11, // pb3.Maps.uint64_to_enum:type_name -> pb3.Maps.Uint64ToEnumEntry
	12, // pb3.Maps.str_to_nested:type_name -> pb3.Maps.StrToNestedEntry
	13, // pb3.Maps.str_to_oneofs:type_name -> pb3.Maps.StrToOneofsEntry
	0,  // pb3.Maps.Uint64ToEnumEntry.value:type_name -> pb3.Enum
	5,  // pb3.Maps.StrToNestedEntry.value:type_name -> pb3.Nested
	6,  // pb3.Maps.StrToOneofsEntry.value:type_name -> pb3.Oneofs
}

func init() { xxx_File_pb3_test_proto_init() }
func xxx_File_pb3_test_proto_init() {
	if File_pb3_test_proto != nil {
		return
	}
	File_pb3_test_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_pb3_test_proto_rawDesc,
		GoTypes:            xxx_File_pb3_test_proto_goTypes,
		DependencyIndexes:  xxx_File_pb3_test_proto_depIdxs,
		EnumOutputTypes:    xxx_File_pb3_test_proto_enumTypes,
		MessageOutputTypes: xxx_File_pb3_test_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_pb3_test_proto_rawDesc = nil
	xxx_File_pb3_test_proto_goTypes = nil
	xxx_File_pb3_test_proto_depIdxs = nil
}