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
// source: google/ads/googleads/v16/enums/ad_group_primary_status_reason.proto

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

// Possible ad group status reasons.
type AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason int32

const (
	// Not specified.
	AdGroupPrimaryStatusReasonEnum_UNSPECIFIED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 0
	// Used for return value only. Represents value unknown in this version.
	AdGroupPrimaryStatusReasonEnum_UNKNOWN AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 1
	// The user-specified campaign status is removed. Contributes to
	// AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_CAMPAIGN_REMOVED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 2
	// The user-specified campaign status is paused. Contributes to
	// AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_CAMPAIGN_PAUSED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 3
	// The user-specified time for this campaign to start is in the future.
	// Contributes to AdGroupPrimaryStatus.NOT_ELIGIBLE
	AdGroupPrimaryStatusReasonEnum_CAMPAIGN_PENDING AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 4
	// The user-specified time for this campaign to end has passed. Contributes
	// to AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_CAMPAIGN_ENDED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 5
	// The user-specified ad group status is paused. Contributes to
	// AdGroupPrimaryStatus.PAUSED.
	AdGroupPrimaryStatusReasonEnum_AD_GROUP_PAUSED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 6
	// The user-specified ad group status is removed. Contributes to
	// AdGroupPrimaryStatus.REMOVED.
	AdGroupPrimaryStatusReasonEnum_AD_GROUP_REMOVED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 7
	// The construction of this ad group is not yet complete. Contributes to
	// AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_AD_GROUP_INCOMPLETE AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 8
	// The user-specified keyword statuses in this ad group are all paused.
	// Contributes to AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_KEYWORDS_PAUSED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 9
	// No eligible keywords exist in this ad group. Contributes to
	// AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_NO_KEYWORDS AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 10
	// The user-specified ad group ads statuses in this ad group are all paused.
	// Contributes to AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_AD_GROUP_ADS_PAUSED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 11
	// No eligible ad group ads exist in this ad group. Contributes to
	// AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_NO_AD_GROUP_ADS AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 12
	// Policy status reason when at least one ad is disapproved. Contributes to
	// multiple AdGroupPrimaryStatus.
	AdGroupPrimaryStatusReasonEnum_HAS_ADS_DISAPPROVED AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 13
	// Policy status reason when at least one ad is limited by policy.
	// Contributes to multiple AdGroupPrimaryStatus.
	AdGroupPrimaryStatusReasonEnum_HAS_ADS_LIMITED_BY_POLICY AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 14
	// Policy status reason when most ads are pending review. Contributes to
	// AdGroupPrimaryStatus.PENDING.
	AdGroupPrimaryStatusReasonEnum_MOST_ADS_UNDER_REVIEW AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 15
	// The AdGroup belongs to a Draft campaign. Contributes to
	// AdGroupPrimaryStatus.NOT_ELIGIBLE.
	AdGroupPrimaryStatusReasonEnum_CAMPAIGN_DRAFT AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 16
	// Ad group has been paused due to prolonged low activity in serving.
	// Contributes to AdGroupPrimaryStatus.PAUSED.
	AdGroupPrimaryStatusReasonEnum_AD_GROUP_PAUSED_DUE_TO_LOW_ACTIVITY AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason = 19
)

// Enum value maps for AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason.
var (
	AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CAMPAIGN_REMOVED",
		3:  "CAMPAIGN_PAUSED",
		4:  "CAMPAIGN_PENDING",
		5:  "CAMPAIGN_ENDED",
		6:  "AD_GROUP_PAUSED",
		7:  "AD_GROUP_REMOVED",
		8:  "AD_GROUP_INCOMPLETE",
		9:  "KEYWORDS_PAUSED",
		10: "NO_KEYWORDS",
		11: "AD_GROUP_ADS_PAUSED",
		12: "NO_AD_GROUP_ADS",
		13: "HAS_ADS_DISAPPROVED",
		14: "HAS_ADS_LIMITED_BY_POLICY",
		15: "MOST_ADS_UNDER_REVIEW",
		16: "CAMPAIGN_DRAFT",
		19: "AD_GROUP_PAUSED_DUE_TO_LOW_ACTIVITY",
	}
	AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason_value = map[string]int32{
		"UNSPECIFIED":                         0,
		"UNKNOWN":                             1,
		"CAMPAIGN_REMOVED":                    2,
		"CAMPAIGN_PAUSED":                     3,
		"CAMPAIGN_PENDING":                    4,
		"CAMPAIGN_ENDED":                      5,
		"AD_GROUP_PAUSED":                     6,
		"AD_GROUP_REMOVED":                    7,
		"AD_GROUP_INCOMPLETE":                 8,
		"KEYWORDS_PAUSED":                     9,
		"NO_KEYWORDS":                         10,
		"AD_GROUP_ADS_PAUSED":                 11,
		"NO_AD_GROUP_ADS":                     12,
		"HAS_ADS_DISAPPROVED":                 13,
		"HAS_ADS_LIMITED_BY_POLICY":           14,
		"MOST_ADS_UNDER_REVIEW":               15,
		"CAMPAIGN_DRAFT":                      16,
		"AD_GROUP_PAUSED_DUE_TO_LOW_ACTIVITY": 19,
	}
)

func (x AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason) Enum() *AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason {
	p := new(AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason)
	*p = x
	return p
}

func (x AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_enumTypes[0].Descriptor()
}

func (AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_enumTypes[0]
}

func (x AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason.Descriptor instead.
func (AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Ad Group Primary Status Reason. Provides insight into why an ad group is not
// serving or not serving optimally. These reasons are aggregated to determine
// an overall AdGroupPrimaryStatus.
type AdGroupPrimaryStatusReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdGroupPrimaryStatusReasonEnum) Reset() {
	*x = AdGroupPrimaryStatusReasonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupPrimaryStatusReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupPrimaryStatusReasonEnum) ProtoMessage() {}

func (x *AdGroupPrimaryStatusReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupPrimaryStatusReasonEnum.ProtoReflect.Descriptor instead.
func (*AdGroupPrimaryStatusReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xda, 0x03, 0x0a, 0x1e, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb7, 0x03, 0x0a, 0x1a, 0x41, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x44,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x06, 0x12,
	0x14, 0x0a, 0x10, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x08, 0x12, 0x13,
	0x0a, 0x0f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45,
	0x44, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52,
	0x44, 0x53, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x5f, 0x41, 0x44, 0x53, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x13, 0x0a,
	0x0f, 0x4e, 0x4f, 0x5f, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x41, 0x44, 0x53,
	0x10, 0x0c, 0x12, 0x17, 0x0a, 0x13, 0x48, 0x41, 0x53, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x44, 0x49,
	0x53, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x1d, 0x0a, 0x19, 0x48,
	0x41, 0x53, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42,
	0x59, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x10, 0x0e, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x4f,
	0x53, 0x54, 0x5f, 0x41, 0x44, 0x53, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x44, 0x52, 0x41, 0x46, 0x54, 0x10, 0x10, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x44, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x44, 0x55, 0x45,
	0x5f, 0x54, 0x4f, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59,
	0x10, 0x13, 0x42, 0xf9, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x1f, 0x41, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x36, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x36, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescData = file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDesc
)

func file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDescData
}

var file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_goTypes = []interface{}{
	(AdGroupPrimaryStatusReasonEnum_AdGroupPrimaryStatusReason)(0), // 0: google.ads.googleads.v16.enums.AdGroupPrimaryStatusReasonEnum.AdGroupPrimaryStatusReason
	(*AdGroupPrimaryStatusReasonEnum)(nil),                         // 1: google.ads.googleads.v16.enums.AdGroupPrimaryStatusReasonEnum
}
var file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_init() }
func file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_init() {
	if File_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdGroupPrimaryStatusReasonEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto = out.File
	file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_rawDesc = nil
	file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_goTypes = nil
	file_google_ads_googleads_v16_enums_ad_group_primary_status_reason_proto_depIdxs = nil
}
