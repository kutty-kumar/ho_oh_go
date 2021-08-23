// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: core_v1/core.proto

package core_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	Gender_unknown_gender Gender = 0
	Gender_male           Gender = 1
	Gender_female         Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "unknown_gender",
		1: "male",
		2: "female",
	}
	Gender_value = map[string]int32{
		"unknown_gender": 0,
		"male":           1,
		"female":         2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_core_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_core_v1_core_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_unknown_status Status = 0
	Status_active         Status = 1
	Status_inactive       Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "unknown_status",
		1: "active",
		2: "inactive",
	}
	Status_value = map[string]int32{
		"unknown_status": 0,
		"active":         1,
		"inactive":       2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_core_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_core_v1_core_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{1}
}

type TimeUnit int32

const (
	TimeUnit_unknown_time_unit TimeUnit = 0
	TimeUnit_second            TimeUnit = 2
	TimeUnit_millisecond       TimeUnit = 1
	TimeUnit_minute            TimeUnit = 3
	TimeUnit_hour              TimeUnit = 4
	TimeUnit_day               TimeUnit = 5
	TimeUnit_week              TimeUnit = 6
	TimeUnit_month             TimeUnit = 7
	TimeUnit_year              TimeUnit = 8
)

// Enum value maps for TimeUnit.
var (
	TimeUnit_name = map[int32]string{
		0: "unknown_time_unit",
		2: "second",
		1: "millisecond",
		3: "minute",
		4: "hour",
		5: "day",
		6: "week",
		7: "month",
		8: "year",
	}
	TimeUnit_value = map[string]int32{
		"unknown_time_unit": 0,
		"second":            2,
		"millisecond":       1,
		"minute":            3,
		"hour":              4,
		"day":               5,
		"week":              6,
		"month":             7,
		"year":              8,
	}
)

func (x TimeUnit) Enum() *TimeUnit {
	p := new(TimeUnit)
	*p = x
	return p
}

func (x TimeUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_core_proto_enumTypes[2].Descriptor()
}

func (TimeUnit) Type() protoreflect.EnumType {
	return &file_core_v1_core_proto_enumTypes[2]
}

func (x TimeUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeUnit.Descriptor instead.
func (TimeUnit) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{2}
}

type IdentityType int32

const (
	IdentityType_unknown_identity_type IdentityType = 0
	IdentityType_email                 IdentityType = 2
	IdentityType_phone                 IdentityType = 1
)

// Enum value maps for IdentityType.
var (
	IdentityType_name = map[int32]string{
		0: "unknown_identity_type",
		2: "email",
		1: "phone",
	}
	IdentityType_value = map[string]int32{
		"unknown_identity_type": 0,
		"email":                 2,
		"phone":                 1,
	}
)

func (x IdentityType) Enum() *IdentityType {
	p := new(IdentityType)
	*p = x
	return p
}

func (x IdentityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_core_proto_enumTypes[3].Descriptor()
}

func (IdentityType) Type() protoreflect.EnumType {
	return &file_core_v1_core_proto_enumTypes[3]
}

func (x IdentityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityType.Descriptor instead.
func (IdentityType) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{3}
}

type Relation int32

const (
	Relation_unknown_relation Relation = 0
	Relation_husband          Relation = 6
	Relation_father           Relation = 1
	Relation_mother           Relation = 2
	Relation_son              Relation = 3
	Relation_daughter         Relation = 4
	Relation_wife             Relation = 5
)

// Enum value maps for Relation.
var (
	Relation_name = map[int32]string{
		0: "unknown_relation",
		6: "husband",
		1: "father",
		2: "mother",
		3: "son",
		4: "daughter",
		5: "wife",
	}
	Relation_value = map[string]int32{
		"unknown_relation": 0,
		"husband":          6,
		"father":           1,
		"mother":           2,
		"son":              3,
		"daughter":         4,
		"wife":             5,
	}
)

func (x Relation) Enum() *Relation {
	p := new(Relation)
	*p = x
	return p
}

func (x Relation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Relation) Descriptor() protoreflect.EnumDescriptor {
	return file_core_v1_core_proto_enumTypes[4].Descriptor()
}

func (Relation) Type() protoreflect.EnumType {
	return &file_core_v1_core_proto_enumTypes[4]
}

func (x Relation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Relation.Descriptor instead.
func (Relation) EnumDescriptor() ([]byte, []int) {
	return file_core_v1_core_proto_rawDescGZIP(), []int{4}
}

var File_core_v1_core_proto protoreflect.FileDescriptor

var file_core_v1_core_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31, 0x2a, 0x32, 0x0a,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x6d,
	0x61, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x65, 0x6d, 0x61, 0x6c, 0x65, 0x10,
	0x02, 0x2a, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x69,
	0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x02, 0x2a, 0x7c, 0x0a, 0x08, 0x54, 0x69, 0x6d,
	0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x10, 0x04, 0x12,
	0x07, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x65, 0x65, 0x6b,
	0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x10, 0x07, 0x12, 0x08, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x10, 0x08, 0x2a, 0x3f, 0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x75, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x2a, 0x66, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x10, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x68, 0x75,
	0x73, 0x62, 0x61, 0x6e, 0x64, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x66, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x73, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x64, 0x61, 0x75, 0x67,
	0x68, 0x74, 0x65, 0x72, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x69, 0x66, 0x65, 0x10, 0x05,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x75, 0x74, 0x74, 0x79, 0x2d, 0x6b, 0x75, 0x6d, 0x61, 0x72, 0x2f, 0x68, 0x6f, 0x5f, 0x6f, 0x68,
	0x5f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_v1_core_proto_rawDescOnce sync.Once
	file_core_v1_core_proto_rawDescData = file_core_v1_core_proto_rawDesc
)

func file_core_v1_core_proto_rawDescGZIP() []byte {
	file_core_v1_core_proto_rawDescOnce.Do(func() {
		file_core_v1_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_v1_core_proto_rawDescData)
	})
	return file_core_v1_core_proto_rawDescData
}

var file_core_v1_core_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_core_v1_core_proto_goTypes = []interface{}{
	(Gender)(0),       // 0: core_v1.Gender
	(Status)(0),       // 1: core_v1.Status
	(TimeUnit)(0),     // 2: core_v1.TimeUnit
	(IdentityType)(0), // 3: core_v1.IdentityType
	(Relation)(0),     // 4: core_v1.Relation
}
var file_core_v1_core_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_v1_core_proto_init() }
func file_core_v1_core_proto_init() {
	if File_core_v1_core_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_v1_core_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_v1_core_proto_goTypes,
		DependencyIndexes: file_core_v1_core_proto_depIdxs,
		EnumInfos:         file_core_v1_core_proto_enumTypes,
	}.Build()
	File_core_v1_core_proto = out.File
	file_core_v1_core_proto_rawDesc = nil
	file_core_v1_core_proto_goTypes = nil
	file_core_v1_core_proto_depIdxs = nil
}
