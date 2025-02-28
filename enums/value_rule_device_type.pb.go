// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: google/ads/googleads/v19/enums/value_rule_device_type.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible device types used in conversion value rule.
type ValueRuleDeviceTypeEnum_ValueRuleDeviceType int32

const (
	// Not specified.
	ValueRuleDeviceTypeEnum_UNSPECIFIED ValueRuleDeviceTypeEnum_ValueRuleDeviceType = 0
	// Used for return value only. Represents value unknown in this version.
	ValueRuleDeviceTypeEnum_UNKNOWN ValueRuleDeviceTypeEnum_ValueRuleDeviceType = 1
	// Mobile device.
	ValueRuleDeviceTypeEnum_MOBILE ValueRuleDeviceTypeEnum_ValueRuleDeviceType = 2
	// Desktop device.
	ValueRuleDeviceTypeEnum_DESKTOP ValueRuleDeviceTypeEnum_ValueRuleDeviceType = 3
	// Tablet device.
	ValueRuleDeviceTypeEnum_TABLET ValueRuleDeviceTypeEnum_ValueRuleDeviceType = 4
)

// Enum value maps for ValueRuleDeviceTypeEnum_ValueRuleDeviceType.
var (
	ValueRuleDeviceTypeEnum_ValueRuleDeviceType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "MOBILE",
		3: "DESKTOP",
		4: "TABLET",
	}
	ValueRuleDeviceTypeEnum_ValueRuleDeviceType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"MOBILE":      2,
		"DESKTOP":     3,
		"TABLET":      4,
	}
)

func (x ValueRuleDeviceTypeEnum_ValueRuleDeviceType) Enum() *ValueRuleDeviceTypeEnum_ValueRuleDeviceType {
	p := new(ValueRuleDeviceTypeEnum_ValueRuleDeviceType)
	*p = x
	return p
}

func (x ValueRuleDeviceTypeEnum_ValueRuleDeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValueRuleDeviceTypeEnum_ValueRuleDeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_enums_value_rule_device_type_proto_enumTypes[0].Descriptor()
}

func (ValueRuleDeviceTypeEnum_ValueRuleDeviceType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_enums_value_rule_device_type_proto_enumTypes[0]
}

func (x ValueRuleDeviceTypeEnum_ValueRuleDeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValueRuleDeviceTypeEnum_ValueRuleDeviceType.Descriptor instead.
func (ValueRuleDeviceTypeEnum_ValueRuleDeviceType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible device types used in a conversion
// value rule.
type ValueRuleDeviceTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValueRuleDeviceTypeEnum) Reset() {
	*x = ValueRuleDeviceTypeEnum{}
	mi := &file_google_ads_googleads_v19_enums_value_rule_device_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValueRuleDeviceTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueRuleDeviceTypeEnum) ProtoMessage() {}

func (x *ValueRuleDeviceTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_enums_value_rule_device_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueRuleDeviceTypeEnum.ProtoReflect.Descriptor instead.
func (*ValueRuleDeviceTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_enums_value_rule_device_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDesc = string([]byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x73, 0x0a,
	0x17, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x58, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x53,
	0x4b, 0x54, 0x4f, 0x50, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54,
	0x10, 0x04, 0x42, 0xf2, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x18, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39,
	0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDesc), len(file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDescData
}

var file_google_ads_googleads_v19_enums_value_rule_device_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_enums_value_rule_device_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_enums_value_rule_device_type_proto_goTypes = []any{
	(ValueRuleDeviceTypeEnum_ValueRuleDeviceType)(0), // 0: google.ads.googleads.v19.enums.ValueRuleDeviceTypeEnum.ValueRuleDeviceType
	(*ValueRuleDeviceTypeEnum)(nil),                  // 1: google.ads.googleads.v19.enums.ValueRuleDeviceTypeEnum
}
var file_google_ads_googleads_v19_enums_value_rule_device_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_enums_value_rule_device_type_proto_init() }
func file_google_ads_googleads_v19_enums_value_rule_device_type_proto_init() {
	if File_google_ads_googleads_v19_enums_value_rule_device_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDesc), len(file_google_ads_googleads_v19_enums_value_rule_device_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_enums_value_rule_device_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_enums_value_rule_device_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_enums_value_rule_device_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_enums_value_rule_device_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_enums_value_rule_device_type_proto = out.File
	file_google_ads_googleads_v19_enums_value_rule_device_type_proto_goTypes = nil
	file_google_ads_googleads_v19_enums_value_rule_device_type_proto_depIdxs = nil
}
