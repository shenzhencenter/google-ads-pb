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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v16/enums/asset_offline_evaluation_error_reasons.proto

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

// Enum describing the quality evaluation disapproval reason of an asset.
type AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons int32

const (
	// Not specified.
	AssetOfflineEvaluationErrorReasonsEnum_UNSPECIFIED AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 0
	// Used for return value only. Represents value unknown in this version.
	AssetOfflineEvaluationErrorReasonsEnum_UNKNOWN AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 1
	// One or more descriptions repeats its corresponding row header.
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_DESCRIPTION_REPEATS_ROW_HEADER AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 2
	// Price asset contains repetitive headers.
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_REPETITIVE_HEADERS AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 3
	// Price item header is not relevant to the price type.
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_HEADER_INCOMPATIBLE_WITH_PRICE_TYPE AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 4
	// Price item description is not relevant to the item header.
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_DESCRIPTION_INCOMPATIBLE_WITH_ITEM_HEADER AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 5
	// Price asset has a price qualifier in a description.
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_DESCRIPTION_HAS_PRICE_QUALIFIER AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 6
	// Unsupported language for price assets
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_UNSUPPORTED_LANGUAGE AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 7
	// Human raters identified an issue with the price asset that isn't captured
	// by other error reasons. The primary purpose of this value is to represent
	// legacy FeedItem disapprovals that are no longer produced.
	AssetOfflineEvaluationErrorReasonsEnum_PRICE_ASSET_OTHER_ERROR AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons = 8
)

// Enum value maps for AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons.
var (
	AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "PRICE_ASSET_DESCRIPTION_REPEATS_ROW_HEADER",
		3: "PRICE_ASSET_REPETITIVE_HEADERS",
		4: "PRICE_ASSET_HEADER_INCOMPATIBLE_WITH_PRICE_TYPE",
		5: "PRICE_ASSET_DESCRIPTION_INCOMPATIBLE_WITH_ITEM_HEADER",
		6: "PRICE_ASSET_DESCRIPTION_HAS_PRICE_QUALIFIER",
		7: "PRICE_ASSET_UNSUPPORTED_LANGUAGE",
		8: "PRICE_ASSET_OTHER_ERROR",
	}
	AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"PRICE_ASSET_DESCRIPTION_REPEATS_ROW_HEADER":            2,
		"PRICE_ASSET_REPETITIVE_HEADERS":                        3,
		"PRICE_ASSET_HEADER_INCOMPATIBLE_WITH_PRICE_TYPE":       4,
		"PRICE_ASSET_DESCRIPTION_INCOMPATIBLE_WITH_ITEM_HEADER": 5,
		"PRICE_ASSET_DESCRIPTION_HAS_PRICE_QUALIFIER":           6,
		"PRICE_ASSET_UNSUPPORTED_LANGUAGE":                      7,
		"PRICE_ASSET_OTHER_ERROR":                               8,
	}
)

func (x AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons) Enum() *AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons {
	p := new(AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons)
	*p = x
	return p
}

func (x AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_enumTypes[0].Descriptor()
}

func (AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_enumTypes[0]
}

func (x AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons.Descriptor instead.
func (AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescGZIP(), []int{0, 0}
}

// Provides the quality evaluation disapproval reason of an asset.
type AssetOfflineEvaluationErrorReasonsEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AssetOfflineEvaluationErrorReasonsEnum) Reset() {
	*x = AssetOfflineEvaluationErrorReasonsEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetOfflineEvaluationErrorReasonsEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetOfflineEvaluationErrorReasonsEnum) ProtoMessage() {}

func (x *AssetOfflineEvaluationErrorReasonsEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetOfflineEvaluationErrorReasonsEnum.ProtoReflect.Descriptor instead.
func (*AssetOfflineEvaluationErrorReasonsEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xa5, 0x03,
	0x0a, 0x26, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfa, 0x02, 0x0a, 0x22, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x2e, 0x0a,
	0x2a, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x53,
	0x5f, 0x52, 0x4f, 0x57, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x22, 0x0a,
	0x1e, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x50,
	0x45, 0x54, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x53, 0x10,
	0x03, 0x12, 0x33, 0x0a, 0x2f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54,
	0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x10, 0x04, 0x12, 0x39, 0x0a, 0x35, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f,
	0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57,
	0x49, 0x54, 0x48, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10,
	0x05, 0x12, 0x2f, 0x0a, 0x2b, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x41, 0x53,
	0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x46, 0x49, 0x45, 0x52,
	0x10, 0x06, 0x12, 0x24, 0x0a, 0x20, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4c, 0x41,
	0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x49, 0x43,
	0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x08, 0x42, 0x81, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x27, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x36, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescData = file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDesc
)

func file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDescData
}

var file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_goTypes = []interface{}{
	(AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons)(0), // 0: google.ads.googleads.v16.enums.AssetOfflineEvaluationErrorReasonsEnum.AssetOfflineEvaluationErrorReasons
	(*AssetOfflineEvaluationErrorReasonsEnum)(nil),                                 // 1: google.ads.googleads.v16.enums.AssetOfflineEvaluationErrorReasonsEnum
}
var file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_init() }
func file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_init() {
	if File_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetOfflineEvaluationErrorReasonsEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto = out.File
	file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_rawDesc = nil
	file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_goTypes = nil
	file_google_ads_googleads_v16_enums_asset_offline_evaluation_error_reasons_proto_depIdxs = nil
}
