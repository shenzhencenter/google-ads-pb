// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v16/enums/google_ads_field_data_type.proto

package enums

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

// These are the various types a GoogleAdsService artifact may take on.
type GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType int32

const (
	// Unspecified
	GoogleAdsFieldDataTypeEnum_UNSPECIFIED GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 0
	// Unknown
	GoogleAdsFieldDataTypeEnum_UNKNOWN GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 1
	// Maps to google.protobuf.BoolValue
	//
	// Applicable operators:  =, !=
	GoogleAdsFieldDataTypeEnum_BOOLEAN GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 2
	// Maps to google.protobuf.StringValue. It can be compared using the set of
	// operators specific to dates however.
	//
	// Applicable operators:  =, <, >, <=, >=, BETWEEN, DURING, and IN
	GoogleAdsFieldDataTypeEnum_DATE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 3
	// Maps to google.protobuf.DoubleValue
	//
	// Applicable operators:  =, !=, <, >, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_DOUBLE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 4
	// Maps to an enum. It's specific definition can be found at type_url.
	//
	// Applicable operators:  =, !=, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_ENUM GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 5
	// Maps to google.protobuf.FloatValue
	//
	// Applicable operators:  =, !=, <, >, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_FLOAT GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 6
	// Maps to google.protobuf.Int32Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_INT32 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 7
	// Maps to google.protobuf.Int64Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_INT64 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 8
	// Maps to a protocol buffer message type. The data type's details can be
	// found in type_url.
	//
	// No operators work with MESSAGE fields.
	GoogleAdsFieldDataTypeEnum_MESSAGE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 9
	// Maps to google.protobuf.StringValue. Represents the resource name
	// (unique id) of a resource or one of its foreign keys.
	//
	// No operators work with RESOURCE_NAME fields.
	GoogleAdsFieldDataTypeEnum_RESOURCE_NAME GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 10
	// Maps to google.protobuf.StringValue.
	//
	// Applicable operators:  =, !=, LIKE, NOT LIKE, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_STRING GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 11
	// Maps to google.protobuf.UInt64Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_UINT64 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 12
)

// Enum value maps for GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType.
var (
	GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "BOOLEAN",
		3:  "DATE",
		4:  "DOUBLE",
		5:  "ENUM",
		6:  "FLOAT",
		7:  "INT32",
		8:  "INT64",
		9:  "MESSAGE",
		10: "RESOURCE_NAME",
		11: "STRING",
		12: "UINT64",
	}
	GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_value = map[string]int32{
		"UNSPECIFIED":   0,
		"UNKNOWN":       1,
		"BOOLEAN":       2,
		"DATE":          3,
		"DOUBLE":        4,
		"ENUM":          5,
		"FLOAT":         6,
		"INT32":         7,
		"INT64":         8,
		"MESSAGE":       9,
		"RESOURCE_NAME": 10,
		"STRING":        11,
		"UINT64":        12,
	}
)

func (x GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) Enum() *GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType {
	p := new(GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType)
	*p = x
	return p
}

func (x GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_enumTypes[0].Descriptor()
}

func (GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_enumTypes[0]
}

func (x GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType.Descriptor instead.
func (GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container holding the various data types.
type GoogleAdsFieldDataTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GoogleAdsFieldDataTypeEnum) Reset() {
	*x = GoogleAdsFieldDataTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsFieldDataTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsFieldDataTypeEnum) ProtoMessage() {}

func (x *GoogleAdsFieldDataTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsFieldDataTypeEnum.ProtoReflect.Descriptor instead.
func (*GoogleAdsFieldDataTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_enums_google_ads_field_data_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x22, 0xdb, 0x01, 0x0a, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0xbc, 0x01, 0x0a, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f,
	0x4c, 0x45, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04,
	0x45, 0x4e, 0x55, 0x4d, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10,
	0x06, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05,
	0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x0c, 0x42,
	0xf5, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescData = file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDesc
)

func file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDescData
}

var file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_goTypes = []interface{}{
	(GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType)(0), // 0: google.ads.googleads.v16.enums.GoogleAdsFieldDataTypeEnum.GoogleAdsFieldDataType
	(*GoogleAdsFieldDataTypeEnum)(nil),                     // 1: google.ads.googleads.v16.enums.GoogleAdsFieldDataTypeEnum
}
var file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_init() }
func file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_init() {
	if File_google_ads_googleads_v16_enums_google_ads_field_data_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsFieldDataTypeEnum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_enums_google_ads_field_data_type_proto = out.File
	file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_rawDesc = nil
	file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_goTypes = nil
	file_google_ads_googleads_v16_enums_google_ads_field_data_type_proto_depIdxs = nil
}
