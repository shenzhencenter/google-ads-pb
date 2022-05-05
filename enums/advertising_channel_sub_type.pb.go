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
// source: google/ads/googleads/v10/enums/advertising_channel_sub_type.proto

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

// Enum describing the different channel subtypes.
type AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType int32

const (
	// Not specified.
	AdvertisingChannelSubTypeEnum_UNSPECIFIED AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 0
	// Used as a return value only. Represents value unknown in this version.
	AdvertisingChannelSubTypeEnum_UNKNOWN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 1
	// Mobile app campaigns for Search.
	AdvertisingChannelSubTypeEnum_SEARCH_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 2
	// Mobile app campaigns for Display.
	AdvertisingChannelSubTypeEnum_DISPLAY_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 3
	// AdWords express campaigns for search.
	AdvertisingChannelSubTypeEnum_SEARCH_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 4
	// AdWords Express campaigns for display.
	AdvertisingChannelSubTypeEnum_DISPLAY_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 5
	// Smart Shopping campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_SMART_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 6
	// Gmail Ad campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_GMAIL_AD AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 7
	// Smart display campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 8
	// Video Outstream campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_OUTSTREAM AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 9
	// Video TrueView for Action campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_ACTION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 10
	// Video campaigns with non-skippable video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_NON_SKIPPABLE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 11
	// App Campaign that allows you to easily promote your Android or iOS app
	// across Google's top properties including Search, Play, YouTube, and the
	// Google Display Network.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 12
	// App Campaign for engagement, focused on driving re-engagement with the
	// app across several of Google's top properties including Search, YouTube,
	// and the Google Display Network.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN_FOR_ENGAGEMENT AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 13
	// Campaigns specialized for local advertising.
	AdvertisingChannelSubTypeEnum_LOCAL_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 14
	// Shopping Comparison Listing campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_COMPARISON_LISTING_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 15
	// Standard Smart campaigns.
	AdvertisingChannelSubTypeEnum_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 16
	// Video campaigns with sequence video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_SEQUENCE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 17
	// App Campaign for pre registration, specialized for advertising mobile
	// app pre-registration, that targets multiple advertising channels across
	// Google Play, YouTube and Display Network. See
	// https://support.google.com/google-ads/answer/9441344 to learn more.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN_FOR_PRE_REGISTRATION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 18
)

// Enum value maps for AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType.
var (
	AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "SEARCH_MOBILE_APP",
		3:  "DISPLAY_MOBILE_APP",
		4:  "SEARCH_EXPRESS",
		5:  "DISPLAY_EXPRESS",
		6:  "SHOPPING_SMART_ADS",
		7:  "DISPLAY_GMAIL_AD",
		8:  "DISPLAY_SMART_CAMPAIGN",
		9:  "VIDEO_OUTSTREAM",
		10: "VIDEO_ACTION",
		11: "VIDEO_NON_SKIPPABLE",
		12: "APP_CAMPAIGN",
		13: "APP_CAMPAIGN_FOR_ENGAGEMENT",
		14: "LOCAL_CAMPAIGN",
		15: "SHOPPING_COMPARISON_LISTING_ADS",
		16: "SMART_CAMPAIGN",
		17: "VIDEO_SEQUENCE",
		18: "APP_CAMPAIGN_FOR_PRE_REGISTRATION",
	}
	AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value = map[string]int32{
		"UNSPECIFIED":                       0,
		"UNKNOWN":                           1,
		"SEARCH_MOBILE_APP":                 2,
		"DISPLAY_MOBILE_APP":                3,
		"SEARCH_EXPRESS":                    4,
		"DISPLAY_EXPRESS":                   5,
		"SHOPPING_SMART_ADS":                6,
		"DISPLAY_GMAIL_AD":                  7,
		"DISPLAY_SMART_CAMPAIGN":            8,
		"VIDEO_OUTSTREAM":                   9,
		"VIDEO_ACTION":                      10,
		"VIDEO_NON_SKIPPABLE":               11,
		"APP_CAMPAIGN":                      12,
		"APP_CAMPAIGN_FOR_ENGAGEMENT":       13,
		"LOCAL_CAMPAIGN":                    14,
		"SHOPPING_COMPARISON_LISTING_ADS":   15,
		"SMART_CAMPAIGN":                    16,
		"VIDEO_SEQUENCE":                    17,
		"APP_CAMPAIGN_FOR_PRE_REGISTRATION": 18,
	}
)

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Enum() *AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType {
	p := new(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType)
	*p = x
	return p
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_enumTypes[0].Descriptor()
}

func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_enumTypes[0]
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType.Descriptor instead.
func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescGZIP(), []int{0, 0}
}

// An immutable specialization of an Advertising Channel.
type AdvertisingChannelSubTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdvertisingChannelSubTypeEnum) Reset() {
	*x = AdvertisingChannelSubTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvertisingChannelSubTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertisingChannelSubTypeEnum) ProtoMessage() {}

func (x *AdvertisingChannelSubTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertisingChannelSubTypeEnum.ProtoReflect.Descriptor instead.
func (*AdvertisingChannelSubTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x22, 0xf8, 0x03, 0x0a, 0x1d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd6, 0x03, 0x0a, 0x19, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4d, 0x4f, 0x42, 0x49,
	0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x49, 0x53, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x48, 0x4f,
	0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x41, 0x44, 0x53, 0x10,
	0x06, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x47, 0x4d, 0x41,
	0x49, 0x4c, 0x5f, 0x41, 0x44, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x49, 0x53, 0x50, 0x4c,
	0x41, 0x59, 0x5f, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x10, 0x08, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x4f, 0x55, 0x54,
	0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x49, 0x44, 0x45,
	0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x56, 0x49,
	0x44, 0x45, 0x4f, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x50, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x50, 0x50, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x10, 0x0c, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x50, 0x50, 0x5f, 0x43, 0x41, 0x4d,
	0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x45, 0x4e, 0x47, 0x41, 0x47, 0x45,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x0e, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x48,
	0x4f, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f,
	0x4e, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x44, 0x53, 0x10, 0x0f, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x53, 0x45, 0x51,
	0x55, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x11, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x50, 0x50, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x5f,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x12, 0x42, 0xf8,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1e, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69,
	0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x30, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescData = file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDesc
)

func file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDescData
}

var file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_goTypes = []interface{}{
	(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType)(0), // 0: google.ads.googleads.v10.enums.AdvertisingChannelSubTypeEnum.AdvertisingChannelSubType
	(*AdvertisingChannelSubTypeEnum)(nil),                        // 1: google.ads.googleads.v10.enums.AdvertisingChannelSubTypeEnum
}
var file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_init() }
func file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_init() {
	if File_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvertisingChannelSubTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto = out.File
	file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_rawDesc = nil
	file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_goTypes = nil
	file_google_ads_googleads_v10_enums_advertising_channel_sub_type_proto_depIdxs = nil
}
