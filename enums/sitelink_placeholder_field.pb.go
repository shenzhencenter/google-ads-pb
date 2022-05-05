// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: google/ads/googleads/v10/enums/sitelink_placeholder_field.proto

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

// Possible values for Sitelink placeholder fields.
type SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField int32

const (
	// Not specified.
	SitelinkPlaceholderFieldEnum_UNSPECIFIED SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	SitelinkPlaceholderFieldEnum_UNKNOWN SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 1
	// Data Type: STRING. The link text for your sitelink.
	SitelinkPlaceholderFieldEnum_TEXT SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 2
	// Data Type: STRING. First line of the sitelink description.
	SitelinkPlaceholderFieldEnum_LINE_1 SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 3
	// Data Type: STRING. Second line of the sitelink description.
	SitelinkPlaceholderFieldEnum_LINE_2 SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 4
	// Data Type: URL_LIST. Final URLs for the sitelink when using Upgraded
	// URLs.
	SitelinkPlaceholderFieldEnum_FINAL_URLS SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 5
	// Data Type: URL_LIST. Final Mobile URLs for the sitelink when using
	// Upgraded URLs.
	SitelinkPlaceholderFieldEnum_FINAL_MOBILE_URLS SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 6
	// Data Type: URL. Tracking template for the sitelink when using Upgraded
	// URLs.
	SitelinkPlaceholderFieldEnum_TRACKING_URL SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 7
	// Data Type: STRING. Final URL suffix for sitelink when using parallel
	// tracking.
	SitelinkPlaceholderFieldEnum_FINAL_URL_SUFFIX SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField = 8
)

// Enum value maps for SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField.
var (
	SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "TEXT",
		3: "LINE_1",
		4: "LINE_2",
		5: "FINAL_URLS",
		6: "FINAL_MOBILE_URLS",
		7: "TRACKING_URL",
		8: "FINAL_URL_SUFFIX",
	}
	SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField_value = map[string]int32{
		"UNSPECIFIED":       0,
		"UNKNOWN":           1,
		"TEXT":              2,
		"LINE_1":            3,
		"LINE_2":            4,
		"FINAL_URLS":        5,
		"FINAL_MOBILE_URLS": 6,
		"TRACKING_URL":      7,
		"FINAL_URL_SUFFIX":  8,
	}
)

func (x SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) Enum() *SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField {
	p := new(SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField)
	*p = x
	return p
}

func (x SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_enumTypes[0]
}

func (x SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField.Descriptor instead.
func (SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Sitelink placeholder fields.
type SitelinkPlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SitelinkPlaceholderFieldEnum) Reset() {
	*x = SitelinkPlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SitelinkPlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SitelinkPlaceholderFieldEnum) ProtoMessage() {}

func (x *SitelinkPlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SitelinkPlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*SitelinkPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x73, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x22, 0xca, 0x01, 0x0a, 0x1c, 0x53, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xa9, 0x01, 0x0a, 0x18, 0x53, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49, 0x4e, 0x45, 0x5f,
	0x31, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x32, 0x10, 0x04, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x05, 0x12,
	0x15, 0x0a, 0x11, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f,
	0x55, 0x52, 0x4c, 0x53, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4e, 0x41,
	0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x10, 0x08, 0x42, 0xf7,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1d, 0x53, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescData = file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_goTypes = []interface{}{
	(SitelinkPlaceholderFieldEnum_SitelinkPlaceholderField)(0), // 0: google.ads.googleads.v10.enums.SitelinkPlaceholderFieldEnum.SitelinkPlaceholderField
	(*SitelinkPlaceholderFieldEnum)(nil),                       // 1: google.ads.googleads.v10.enums.SitelinkPlaceholderFieldEnum
}
var file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_init() }
func file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_init() {
	if File_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SitelinkPlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto = out.File
	file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_sitelink_placeholder_field_proto_depIdxs = nil
}
