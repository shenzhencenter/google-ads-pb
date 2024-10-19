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
// 	protoc-gen-go v1.35.1
// 	protoc        v4.24.4
// source: google/ads/googleads/v18/enums/conversion_action_type.proto

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

// Possible types of a conversion action.
type ConversionActionTypeEnum_ConversionActionType int32

const (
	// Not specified.
	ConversionActionTypeEnum_UNSPECIFIED ConversionActionTypeEnum_ConversionActionType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionTypeEnum_UNKNOWN ConversionActionTypeEnum_ConversionActionType = 1
	// Conversions that occur when a user clicks on an ad's call extension.
	ConversionActionTypeEnum_AD_CALL ConversionActionTypeEnum_ConversionActionType = 2
	// Conversions that occur when a user on a mobile device clicks a phone
	// number.
	ConversionActionTypeEnum_CLICK_TO_CALL ConversionActionTypeEnum_ConversionActionType = 3
	// Conversions that occur when a user downloads a mobile app from the Google
	// Play Store.
	ConversionActionTypeEnum_GOOGLE_PLAY_DOWNLOAD ConversionActionTypeEnum_ConversionActionType = 4
	// Conversions that occur when a user makes a purchase in an app through
	// Android billing.
	ConversionActionTypeEnum_GOOGLE_PLAY_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 5
	// Call conversions that are tracked by the advertiser and uploaded.
	ConversionActionTypeEnum_UPLOAD_CALLS ConversionActionTypeEnum_ConversionActionType = 6
	// Conversions that are tracked by the advertiser and uploaded with
	// attributed clicks.
	ConversionActionTypeEnum_UPLOAD_CLICKS ConversionActionTypeEnum_ConversionActionType = 7
	// Conversions that occur on a webpage.
	ConversionActionTypeEnum_WEBPAGE ConversionActionTypeEnum_ConversionActionType = 8
	// Conversions that occur when a user calls a dynamically-generated phone
	// number from an advertiser's website.
	ConversionActionTypeEnum_WEBSITE_CALL ConversionActionTypeEnum_ConversionActionType = 9
	// Store Sales conversion based on first-party or third-party merchant
	// data uploads.
	// Only customers on the allowlist can use store sales direct upload types.
	ConversionActionTypeEnum_STORE_SALES_DIRECT_UPLOAD ConversionActionTypeEnum_ConversionActionType = 10
	// Store Sales conversion based on first-party or third-party merchant
	// data uploads and/or from in-store purchases using cards from payment
	// networks.
	// Only customers on the allowlist can use store sales types.
	// Read only.
	ConversionActionTypeEnum_STORE_SALES ConversionActionTypeEnum_ConversionActionType = 11
	// Android app first open conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_ANDROID_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 12
	// Android app in app purchase conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_ANDROID_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 13
	// Android app custom conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_ANDROID_CUSTOM ConversionActionTypeEnum_ConversionActionType = 14
	// iOS app first open conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_IOS_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 15
	// iOS app in app purchase conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_IOS_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 16
	// iOS app custom conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_IOS_CUSTOM ConversionActionTypeEnum_ConversionActionType = 17
	// Android app first open conversions tracked through Third Party App
	// Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 18
	// Android app in app purchase conversions tracked through Third Party App
	// Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 19
	// Android app custom conversions tracked through Third Party App Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM ConversionActionTypeEnum_ConversionActionType = 20
	// iOS app first open conversions tracked through Third Party App Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 21
	// iOS app in app purchase conversions tracked through Third Party App
	// Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 22
	// iOS app custom conversions tracked through Third Party App Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM ConversionActionTypeEnum_ConversionActionType = 23
	// Conversions that occur when a user pre-registers a mobile app from the
	// Google Play Store. Read only.
	ConversionActionTypeEnum_ANDROID_APP_PRE_REGISTRATION ConversionActionTypeEnum_ConversionActionType = 24
	// Conversions that track all Google Play downloads which aren't tracked
	// by an app-specific type. Read only.
	ConversionActionTypeEnum_ANDROID_INSTALLS_ALL_OTHER_APPS ConversionActionTypeEnum_ConversionActionType = 25
	// Floodlight activity that counts the number of times that users have
	// visited a particular webpage after seeing or clicking on one of
	// an advertiser's ads. Read only.
	ConversionActionTypeEnum_FLOODLIGHT_ACTION ConversionActionTypeEnum_ConversionActionType = 26
	// Floodlight activity that tracks the number of sales made or the number
	// of items purchased. Can also capture the total value of each sale.
	// Read only.
	ConversionActionTypeEnum_FLOODLIGHT_TRANSACTION ConversionActionTypeEnum_ConversionActionType = 27
	// Conversions that track local actions from Google's products and
	// services after interacting with an ad. Read only.
	ConversionActionTypeEnum_GOOGLE_HOSTED ConversionActionTypeEnum_ConversionActionType = 28
	// Conversions reported when a user submits a lead form. Read only.
	ConversionActionTypeEnum_LEAD_FORM_SUBMIT ConversionActionTypeEnum_ConversionActionType = 29
	// Conversions that come from Salesforce. Read only.
	ConversionActionTypeEnum_SALESFORCE ConversionActionTypeEnum_ConversionActionType = 30
	// Conversions imported from Search Ads 360 Floodlight data. Read only.
	ConversionActionTypeEnum_SEARCH_ADS_360 ConversionActionTypeEnum_ConversionActionType = 31
	// Call conversions that occur on Smart campaign Ads without call tracking
	// setup, using Smart campaign custom criteria. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_AD_CLICKS_TO_CALL ConversionActionTypeEnum_ConversionActionType = 32
	// The user clicks on a call element within Google Maps. Smart campaign
	// only. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_MAP_CLICKS_TO_CALL ConversionActionTypeEnum_ConversionActionType = 33
	// The user requests directions to a business location within Google Maps.
	// Smart campaign only. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_MAP_DIRECTIONS ConversionActionTypeEnum_ConversionActionType = 34
	// Call conversions that occur on Smart campaign Ads with call tracking
	// setup, using Smart campaign custom criteria. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_TRACKED_CALLS ConversionActionTypeEnum_ConversionActionType = 35
	// Conversions that occur when a user visits an advertiser's retail store.
	// Read only.
	ConversionActionTypeEnum_STORE_VISITS ConversionActionTypeEnum_ConversionActionType = 36
	// Conversions created from website events (such as form submissions or page
	// loads), that don't use individually coded event snippets. Read only.
	ConversionActionTypeEnum_WEBPAGE_CODELESS ConversionActionTypeEnum_ConversionActionType = 37
	// Conversions that come from linked Universal Analytics goals.
	ConversionActionTypeEnum_UNIVERSAL_ANALYTICS_GOAL ConversionActionTypeEnum_ConversionActionType = 38
	// Conversions that come from linked Universal Analytics transactions.
	ConversionActionTypeEnum_UNIVERSAL_ANALYTICS_TRANSACTION ConversionActionTypeEnum_ConversionActionType = 39
	// Conversions that come from linked Google Analytics 4 custom event
	// conversions.
	ConversionActionTypeEnum_GOOGLE_ANALYTICS_4_CUSTOM ConversionActionTypeEnum_ConversionActionType = 40
	// Conversions that come from linked Google Analytics 4 purchase
	// conversions.
	ConversionActionTypeEnum_GOOGLE_ANALYTICS_4_PURCHASE ConversionActionTypeEnum_ConversionActionType = 41
)

// Enum value maps for ConversionActionTypeEnum_ConversionActionType.
var (
	ConversionActionTypeEnum_ConversionActionType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "AD_CALL",
		3:  "CLICK_TO_CALL",
		4:  "GOOGLE_PLAY_DOWNLOAD",
		5:  "GOOGLE_PLAY_IN_APP_PURCHASE",
		6:  "UPLOAD_CALLS",
		7:  "UPLOAD_CLICKS",
		8:  "WEBPAGE",
		9:  "WEBSITE_CALL",
		10: "STORE_SALES_DIRECT_UPLOAD",
		11: "STORE_SALES",
		12: "FIREBASE_ANDROID_FIRST_OPEN",
		13: "FIREBASE_ANDROID_IN_APP_PURCHASE",
		14: "FIREBASE_ANDROID_CUSTOM",
		15: "FIREBASE_IOS_FIRST_OPEN",
		16: "FIREBASE_IOS_IN_APP_PURCHASE",
		17: "FIREBASE_IOS_CUSTOM",
		18: "THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN",
		19: "THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE",
		20: "THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM",
		21: "THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN",
		22: "THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE",
		23: "THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM",
		24: "ANDROID_APP_PRE_REGISTRATION",
		25: "ANDROID_INSTALLS_ALL_OTHER_APPS",
		26: "FLOODLIGHT_ACTION",
		27: "FLOODLIGHT_TRANSACTION",
		28: "GOOGLE_HOSTED",
		29: "LEAD_FORM_SUBMIT",
		30: "SALESFORCE",
		31: "SEARCH_ADS_360",
		32: "SMART_CAMPAIGN_AD_CLICKS_TO_CALL",
		33: "SMART_CAMPAIGN_MAP_CLICKS_TO_CALL",
		34: "SMART_CAMPAIGN_MAP_DIRECTIONS",
		35: "SMART_CAMPAIGN_TRACKED_CALLS",
		36: "STORE_VISITS",
		37: "WEBPAGE_CODELESS",
		38: "UNIVERSAL_ANALYTICS_GOAL",
		39: "UNIVERSAL_ANALYTICS_TRANSACTION",
		40: "GOOGLE_ANALYTICS_4_CUSTOM",
		41: "GOOGLE_ANALYTICS_4_PURCHASE",
	}
	ConversionActionTypeEnum_ConversionActionType_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"AD_CALL":                          2,
		"CLICK_TO_CALL":                    3,
		"GOOGLE_PLAY_DOWNLOAD":             4,
		"GOOGLE_PLAY_IN_APP_PURCHASE":      5,
		"UPLOAD_CALLS":                     6,
		"UPLOAD_CLICKS":                    7,
		"WEBPAGE":                          8,
		"WEBSITE_CALL":                     9,
		"STORE_SALES_DIRECT_UPLOAD":        10,
		"STORE_SALES":                      11,
		"FIREBASE_ANDROID_FIRST_OPEN":      12,
		"FIREBASE_ANDROID_IN_APP_PURCHASE": 13,
		"FIREBASE_ANDROID_CUSTOM":          14,
		"FIREBASE_IOS_FIRST_OPEN":          15,
		"FIREBASE_IOS_IN_APP_PURCHASE":     16,
		"FIREBASE_IOS_CUSTOM":              17,
		"THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN":      18,
		"THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE": 19,
		"THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM":          20,
		"THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN":          21,
		"THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE":     22,
		"THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM":              23,
		"ANDROID_APP_PRE_REGISTRATION":                      24,
		"ANDROID_INSTALLS_ALL_OTHER_APPS":                   25,
		"FLOODLIGHT_ACTION":                                 26,
		"FLOODLIGHT_TRANSACTION":                            27,
		"GOOGLE_HOSTED":                                     28,
		"LEAD_FORM_SUBMIT":                                  29,
		"SALESFORCE":                                        30,
		"SEARCH_ADS_360":                                    31,
		"SMART_CAMPAIGN_AD_CLICKS_TO_CALL":                  32,
		"SMART_CAMPAIGN_MAP_CLICKS_TO_CALL":                 33,
		"SMART_CAMPAIGN_MAP_DIRECTIONS":                     34,
		"SMART_CAMPAIGN_TRACKED_CALLS":                      35,
		"STORE_VISITS":                                      36,
		"WEBPAGE_CODELESS":                                  37,
		"UNIVERSAL_ANALYTICS_GOAL":                          38,
		"UNIVERSAL_ANALYTICS_TRANSACTION":                   39,
		"GOOGLE_ANALYTICS_4_CUSTOM":                         40,
		"GOOGLE_ANALYTICS_4_PURCHASE":                       41,
	}
)

func (x ConversionActionTypeEnum_ConversionActionType) Enum() *ConversionActionTypeEnum_ConversionActionType {
	p := new(ConversionActionTypeEnum_ConversionActionType)
	*p = x
	return p
}

func (x ConversionActionTypeEnum_ConversionActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionActionTypeEnum_ConversionActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_enums_conversion_action_type_proto_enumTypes[0].Descriptor()
}

func (ConversionActionTypeEnum_ConversionActionType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_enums_conversion_action_type_proto_enumTypes[0]
}

func (x ConversionActionTypeEnum_ConversionActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionActionTypeEnum_ConversionActionType.Descriptor instead.
func (ConversionActionTypeEnum_ConversionActionType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible types of a conversion action.
type ConversionActionTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionActionTypeEnum) Reset() {
	*x = ConversionActionTypeEnum{}
	mi := &file_google_ads_googleads_v18_enums_conversion_action_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionActionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionActionTypeEnum) ProtoMessage() {}

func (x *ConversionActionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_enums_conversion_action_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionActionTypeEnum.ProtoReflect.Descriptor instead.
func (*ConversionActionTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_enums_conversion_action_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x82, 0x0a,
	0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xe5, 0x09, 0x0a, 0x14, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10,
	0x03, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59,
	0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x47,
	0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x06, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x53, 0x10,
	0x07, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x45, 0x42, 0x50, 0x41, 0x47, 0x45, 0x10, 0x08, 0x12, 0x10,
	0x0a, 0x0c, 0x57, 0x45, 0x42, 0x53, 0x49, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x09,
	0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x5f,
	0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x0a, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x0b,
	0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10,
	0x0c, 0x12, 0x24, 0x0a, 0x20, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x41, 0x4e,
	0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x55, 0x52,
	0x43, 0x48, 0x41, 0x53, 0x45, 0x10, 0x0d, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x52, 0x45, 0x42,
	0x41, 0x53, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54,
	0x4f, 0x4d, 0x10, 0x0e, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x49, 0x4f, 0x53, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10,
	0x0f, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4f,
	0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53,
	0x45, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x49, 0x52, 0x45, 0x42, 0x41, 0x53, 0x45, 0x5f,
	0x49, 0x4f, 0x53, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x11, 0x12, 0x30, 0x0a, 0x2c,
	0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x5f,
	0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49,
	0x44, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x12, 0x12, 0x35,
	0x0a, 0x31, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x52,
	0x4f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48,
	0x41, 0x53, 0x45, 0x10, 0x13, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50,
	0x41, 0x52, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49,
	0x43, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f,
	0x4d, 0x10, 0x14, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52,
	0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53,
	0x5f, 0x49, 0x4f, 0x53, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10,
	0x15, 0x12, 0x31, 0x0a, 0x2d, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59,
	0x5f, 0x41, 0x50, 0x50, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x49,
	0x4f, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41,
	0x53, 0x45, 0x10, 0x16, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41,
	0x52, 0x54, 0x59, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43,
	0x53, 0x5f, 0x49, 0x4f, 0x53, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10, 0x17, 0x12, 0x20,
	0x0a, 0x1c, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x50, 0x52,
	0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x18,
	0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x49, 0x4e, 0x53, 0x54,
	0x41, 0x4c, 0x4c, 0x53, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x41,
	0x50, 0x50, 0x53, 0x10, 0x19, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x4c, 0x4f, 0x4f, 0x44, 0x4c, 0x49,
	0x47, 0x48, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1a, 0x12, 0x1a, 0x0a, 0x16,
	0x46, 0x4c, 0x4f, 0x4f, 0x44, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1b, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x4f, 0x4f, 0x47,
	0x4c, 0x45, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x45, 0x44, 0x10, 0x1c, 0x12, 0x14, 0x0a, 0x10, 0x4c,
	0x45, 0x41, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x10,
	0x1d, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x41, 0x4c, 0x45, 0x53, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x10,
	0x1e, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x44, 0x53, 0x5f,
	0x33, 0x36, 0x30, 0x10, 0x1f, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x41, 0x44, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b,
	0x53, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x20, 0x12, 0x25, 0x0a, 0x21, 0x53,
	0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x41,
	0x50, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x53, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x41, 0x4c, 0x4c,
	0x10, 0x21, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x53, 0x10, 0x22, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x45, 0x44, 0x5f,
	0x43, 0x41, 0x4c, 0x4c, 0x53, 0x10, 0x23, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x5f, 0x56, 0x49, 0x53, 0x49, 0x54, 0x53, 0x10, 0x24, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x45, 0x42,
	0x50, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x25, 0x12,
	0x1c, 0x0a, 0x18, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x5f, 0x41, 0x4e, 0x41,
	0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x47, 0x4f, 0x41, 0x4c, 0x10, 0x26, 0x12, 0x23, 0x0a,
	0x1f, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59,
	0x54, 0x49, 0x43, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x27, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x4e, 0x41,
	0x4c, 0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x34, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10,
	0x28, 0x12, 0x1f, 0x0a, 0x1b, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x41, 0x4e, 0x41, 0x4c,
	0x59, 0x54, 0x49, 0x43, 0x53, 0x5f, 0x34, 0x5f, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45,
	0x10, 0x29, 0x42, 0xf3, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x19, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x38, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescData = file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDesc
)

func file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDescData
}

var file_google_ads_googleads_v18_enums_conversion_action_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_enums_conversion_action_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_enums_conversion_action_type_proto_goTypes = []any{
	(ConversionActionTypeEnum_ConversionActionType)(0), // 0: google.ads.googleads.v18.enums.ConversionActionTypeEnum.ConversionActionType
	(*ConversionActionTypeEnum)(nil),                   // 1: google.ads.googleads.v18.enums.ConversionActionTypeEnum
}
var file_google_ads_googleads_v18_enums_conversion_action_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_enums_conversion_action_type_proto_init() }
func file_google_ads_googleads_v18_enums_conversion_action_type_proto_init() {
	if File_google_ads_googleads_v18_enums_conversion_action_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_enums_conversion_action_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_enums_conversion_action_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_enums_conversion_action_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_enums_conversion_action_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_enums_conversion_action_type_proto = out.File
	file_google_ads_googleads_v18_enums_conversion_action_type_proto_rawDesc = nil
	file_google_ads_googleads_v18_enums_conversion_action_type_proto_goTypes = nil
	file_google_ads_googleads_v18_enums_conversion_action_type_proto_depIdxs = nil
}
