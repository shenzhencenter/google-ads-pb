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
// source: google/ads/googleads/v19/enums/listing_type.proto

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

// Possible listing types.
type ListingTypeEnum_ListingType int32

const (
	// Not specified.
	ListingTypeEnum_UNSPECIFIED ListingTypeEnum_ListingType = 0
	// Used for return value only. Represents value unknown in this version.
	ListingTypeEnum_UNKNOWN ListingTypeEnum_ListingType = 1
	// This campaign serves vehicle ads.
	ListingTypeEnum_VEHICLES ListingTypeEnum_ListingType = 2
)

// Enum value maps for ListingTypeEnum_ListingType.
var (
	ListingTypeEnum_ListingType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "VEHICLES",
	}
	ListingTypeEnum_ListingType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"VEHICLES":    2,
	}
)

func (x ListingTypeEnum_ListingType) Enum() *ListingTypeEnum_ListingType {
	p := new(ListingTypeEnum_ListingType)
	*p = x
	return p
}

func (x ListingTypeEnum_ListingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListingTypeEnum_ListingType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_enums_listing_type_proto_enumTypes[0].Descriptor()
}

func (ListingTypeEnum_ListingType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_enums_listing_type_proto_enumTypes[0]
}

func (x ListingTypeEnum_ListingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListingTypeEnum_ListingType.Descriptor instead.
func (ListingTypeEnum_ListingType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_listing_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible listing types.
type ListingTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListingTypeEnum) Reset() {
	*x = ListingTypeEnum{}
	mi := &file_google_ads_googleads_v19_enums_listing_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListingTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListingTypeEnum) ProtoMessage() {}

func (x *ListingTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_enums_listing_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListingTypeEnum.ProtoReflect.Descriptor instead.
func (*ListingTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_listing_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_enums_listing_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_enums_listing_type_proto_rawDesc = string([]byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x22, 0x4c, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x39, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x53, 0x10,
	0x02, 0x42, 0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_google_ads_googleads_v19_enums_listing_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_enums_listing_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_enums_listing_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_enums_listing_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_enums_listing_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_listing_type_proto_rawDesc), len(file_google_ads_googleads_v19_enums_listing_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_enums_listing_type_proto_rawDescData
}

var file_google_ads_googleads_v19_enums_listing_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_enums_listing_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_enums_listing_type_proto_goTypes = []any{
	(ListingTypeEnum_ListingType)(0), // 0: google.ads.googleads.v19.enums.ListingTypeEnum.ListingType
	(*ListingTypeEnum)(nil),          // 1: google.ads.googleads.v19.enums.ListingTypeEnum
}
var file_google_ads_googleads_v19_enums_listing_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_enums_listing_type_proto_init() }
func file_google_ads_googleads_v19_enums_listing_type_proto_init() {
	if File_google_ads_googleads_v19_enums_listing_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_listing_type_proto_rawDesc), len(file_google_ads_googleads_v19_enums_listing_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_enums_listing_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_enums_listing_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_enums_listing_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_enums_listing_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_enums_listing_type_proto = out.File
	file_google_ads_googleads_v19_enums_listing_type_proto_goTypes = nil
	file_google_ads_googleads_v19_enums_listing_type_proto_depIdxs = nil
}
