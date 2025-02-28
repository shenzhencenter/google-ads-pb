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
// source: google/ads/googleads/v19/enums/hotel_asset_suggestion_status.proto

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

// Possible statuses of a hotel asset suggestion.
type HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus int32

const (
	// Enum unspecified.
	HotelAssetSuggestionStatusEnum_UNSPECIFIED HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus = 0
	// The received error code is not known in this version.
	HotelAssetSuggestionStatusEnum_UNKNOWN HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus = 1
	// The hotel asset suggestion was successfully retrieved.
	HotelAssetSuggestionStatusEnum_SUCCESS HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus = 2
	// A hotel look up returns nothing.
	HotelAssetSuggestionStatusEnum_HOTEL_NOT_FOUND HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus = 3
	// A Google Places ID is invalid and cannot be decoded.
	HotelAssetSuggestionStatusEnum_INVALID_PLACE_ID HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus = 4
)

// Enum value maps for HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus.
var (
	HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "SUCCESS",
		3: "HOTEL_NOT_FOUND",
		4: "INVALID_PLACE_ID",
	}
	HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus_value = map[string]int32{
		"UNSPECIFIED":      0,
		"UNKNOWN":          1,
		"SUCCESS":          2,
		"HOTEL_NOT_FOUND":  3,
		"INVALID_PLACE_ID": 4,
	}
)

func (x HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus) Enum() *HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus {
	p := new(HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus)
	*p = x
	return p
}

func (x HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_enumTypes[0].Descriptor()
}

func (HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_enumTypes[0]
}

func (x HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus.Descriptor instead.
func (HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible statuses of a hotel asset suggestion.
type HotelAssetSuggestionStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HotelAssetSuggestionStatusEnum) Reset() {
	*x = HotelAssetSuggestionStatusEnum{}
	mi := &file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HotelAssetSuggestionStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotelAssetSuggestionStatusEnum) ProtoMessage() {}

func (x *HotelAssetSuggestionStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotelAssetSuggestionStatusEnum.ProtoReflect.Descriptor instead.
func (*HotelAssetSuggestionStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDesc = string([]byte{
	0x0a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x1e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x72, 0x0a, 0x1a, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x48, 0x4f, 0x54, 0x45, 0x4c, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x04, 0x42, 0xf9, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x1f, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
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
	file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDesc), len(file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDescData
}

var file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_goTypes = []any{
	(HotelAssetSuggestionStatusEnum_HotelAssetSuggestionStatus)(0), // 0: google.ads.googleads.v19.enums.HotelAssetSuggestionStatusEnum.HotelAssetSuggestionStatus
	(*HotelAssetSuggestionStatusEnum)(nil),                         // 1: google.ads.googleads.v19.enums.HotelAssetSuggestionStatusEnum
}
var file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_init() }
func file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_init() {
	if File_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDesc), len(file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto = out.File
	file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_goTypes = nil
	file_google_ads_googleads_v19_enums_hotel_asset_suggestion_status_proto_depIdxs = nil
}
