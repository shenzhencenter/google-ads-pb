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
// source: google/ads/googleads/v19/enums/ad_strength.proto

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

// Enum listing the possible ad strengths.
type AdStrengthEnum_AdStrength int32

const (
	// Not specified.
	AdStrengthEnum_UNSPECIFIED AdStrengthEnum_AdStrength = 0
	// Used for return value only. Represents value unknown in this version.
	AdStrengthEnum_UNKNOWN AdStrengthEnum_AdStrength = 1
	// The ad strength is currently pending.
	AdStrengthEnum_PENDING AdStrengthEnum_AdStrength = 2
	// No ads could be generated.
	AdStrengthEnum_NO_ADS AdStrengthEnum_AdStrength = 3
	// Poor strength.
	AdStrengthEnum_POOR AdStrengthEnum_AdStrength = 4
	// Average strength.
	AdStrengthEnum_AVERAGE AdStrengthEnum_AdStrength = 5
	// Good strength.
	AdStrengthEnum_GOOD AdStrengthEnum_AdStrength = 6
	// Excellent strength.
	AdStrengthEnum_EXCELLENT AdStrengthEnum_AdStrength = 7
)

// Enum value maps for AdStrengthEnum_AdStrength.
var (
	AdStrengthEnum_AdStrength_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "PENDING",
		3: "NO_ADS",
		4: "POOR",
		5: "AVERAGE",
		6: "GOOD",
		7: "EXCELLENT",
	}
	AdStrengthEnum_AdStrength_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"PENDING":     2,
		"NO_ADS":      3,
		"POOR":        4,
		"AVERAGE":     5,
		"GOOD":        6,
		"EXCELLENT":   7,
	}
)

func (x AdStrengthEnum_AdStrength) Enum() *AdStrengthEnum_AdStrength {
	p := new(AdStrengthEnum_AdStrength)
	*p = x
	return p
}

func (x AdStrengthEnum_AdStrength) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdStrengthEnum_AdStrength) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_enums_ad_strength_proto_enumTypes[0].Descriptor()
}

func (AdStrengthEnum_AdStrength) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_enums_ad_strength_proto_enumTypes[0]
}

func (x AdStrengthEnum_AdStrength) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdStrengthEnum_AdStrength.Descriptor instead.
func (AdStrengthEnum_AdStrength) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ad strengths.
type AdStrengthEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdStrengthEnum) Reset() {
	*x = AdStrengthEnum{}
	mi := &file_google_ads_googleads_v19_enums_ad_strength_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdStrengthEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdStrengthEnum) ProtoMessage() {}

func (x *AdStrengthEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_enums_ad_strength_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdStrengthEnum.ProtoReflect.Descriptor instead.
func (*AdStrengthEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_enums_ad_strength_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_enums_ad_strength_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x73, 0x0a, 0x0a, 0x41, 0x64, 0x53, 0x74, 0x72, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x4e, 0x4f, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f,
	0x4f, 0x52, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x4f, 0x4f, 0x44, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x45,
	0x58, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x42, 0xe9, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x42, 0x0f, 0x41, 0x64, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a,
	0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a,
	0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_ad_strength_proto_rawDesc), len(file_google_ads_googleads_v19_enums_ad_strength_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_enums_ad_strength_proto_rawDescData
}

var file_google_ads_googleads_v19_enums_ad_strength_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_enums_ad_strength_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_enums_ad_strength_proto_goTypes = []any{
	(AdStrengthEnum_AdStrength)(0), // 0: google.ads.googleads.v19.enums.AdStrengthEnum.AdStrength
	(*AdStrengthEnum)(nil),         // 1: google.ads.googleads.v19.enums.AdStrengthEnum
}
var file_google_ads_googleads_v19_enums_ad_strength_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_enums_ad_strength_proto_init() }
func file_google_ads_googleads_v19_enums_ad_strength_proto_init() {
	if File_google_ads_googleads_v19_enums_ad_strength_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_enums_ad_strength_proto_rawDesc), len(file_google_ads_googleads_v19_enums_ad_strength_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_enums_ad_strength_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_enums_ad_strength_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_enums_ad_strength_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_enums_ad_strength_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_enums_ad_strength_proto = out.File
	file_google_ads_googleads_v19_enums_ad_strength_proto_goTypes = nil
	file_google_ads_googleads_v19_enums_ad_strength_proto_depIdxs = nil
}
