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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: google/ads/googleads/v17/enums/ad_group_criterion_primary_status_reason.proto

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

// Enum describing the possible Ad Group Criterion primary status reasons.
// Provides insight into why an Ad Group Criterion is not serving or not
// serving optimally. These reasons are aggregated to determine an overall
// Ad Group Criterion primary status.
type AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason int32

const (
	// Not specified.
	AdGroupCriterionPrimaryStatusReasonEnum_UNSPECIFIED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 0
	// Used for return value only. Represents unknown value in this version.
	AdGroupCriterionPrimaryStatusReasonEnum_UNKNOWN AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 1
	// The user-specified time for this campaign to start is in the future.
	// Contributes to AdGroupCriterionPrimaryStatus.PENDING.
	AdGroupCriterionPrimaryStatusReasonEnum_CAMPAIGN_PENDING AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 2
	// The ad group criterion is overridden by negative campaign criterion.
	// Contributes to AdGroupCriterionPrimaryStatus.NOT_ELIGIBLE.
	AdGroupCriterionPrimaryStatusReasonEnum_CAMPAIGN_CRITERION_NEGATIVE AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 3
	// The user-specified campaign status is paused. Contributes to
	// AdGroupCriterionPrimaryStatus.PAUSED.
	AdGroupCriterionPrimaryStatusReasonEnum_CAMPAIGN_PAUSED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 4
	// The user-specified campaign status is removed. Contributes to
	// AdGroupCriterionPrimaryStatus.REMOVED.
	AdGroupCriterionPrimaryStatusReasonEnum_CAMPAIGN_REMOVED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 5
	// The user-specified time for this campaign to end has passed. Contributes
	// to AdGroupCriterionPrimaryStatus.ENDED.
	AdGroupCriterionPrimaryStatusReasonEnum_CAMPAIGN_ENDED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 6
	// The user-specified ad group status is paused. Contributes to
	// AdGroupCriterionPrimaryStatus.PAUSED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_PAUSED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 7
	// The user-specified ad group status is removed. Contributes to
	// AdGroupCriterionPrimaryStatus.REMOVED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_REMOVED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 8
	// The ad group criterion is disapproved by the ads approval system.
	// Contributes to AdGroupCriterionPrimaryStatus.NOT_ELIGIBLE.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_DISAPPROVED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 9
	// The ad group criterion is rarely served. Contributes to
	// AdGroupCriterionPrimaryStatus.NOT_ELIGIBLE.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_RARELY_SERVED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 10
	// The ad group criterion has a low quality score. Contributes to
	// AdGroupCriterionPrimaryStatus.LIMITED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_LOW_QUALITY AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 11
	// The ad group criterion is under review. Contributes to
	// AdGroupCriterionPrimaryStatus.PENDING.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_UNDER_REVIEW AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 12
	// The ad group criterion is pending review. Contributes to
	// AdGroupCriterionPrimaryStatus.NOT_ELIGIBLE.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_PENDING_REVIEW AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 13
	// The ad group criterion's bid is below the value necessary to serve on the
	// first page. Contributes to AdGroupCriterionPrimaryStatus.LIMITED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_BELOW_FIRST_PAGE_BID AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 14
	// The ad group criterion is negative. Contributes to
	// AdGroupCriterionPrimaryStatus.NOT_ELIGIBLE.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_NEGATIVE AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 15
	// The ad group criterion is restricted. Contributes to
	// AdGroupCriterionPrimaryStatus.NOT_ELIGIBLE.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_RESTRICTED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 16
	// The user-specified ad group criterion status is paused. Contributes to
	// AdGroupCriterionPrimaryStatus.PAUSED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_PAUSED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 17
	// The ad group criterion has been paused due to prolonged low activity in
	// serving. Contributes to AdGroupCriterionPrimaryStatus.PAUSED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_PAUSED_DUE_TO_LOW_ACTIVITY AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 18
	// The user-specified ad group criterion status is removed. Contributes to
	// AdGroupCriterionPrimaryStatus.REMOVED.
	AdGroupCriterionPrimaryStatusReasonEnum_AD_GROUP_CRITERION_REMOVED AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason = 19
)

// Enum value maps for AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason.
var (
	AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CAMPAIGN_PENDING",
		3:  "CAMPAIGN_CRITERION_NEGATIVE",
		4:  "CAMPAIGN_PAUSED",
		5:  "CAMPAIGN_REMOVED",
		6:  "CAMPAIGN_ENDED",
		7:  "AD_GROUP_PAUSED",
		8:  "AD_GROUP_REMOVED",
		9:  "AD_GROUP_CRITERION_DISAPPROVED",
		10: "AD_GROUP_CRITERION_RARELY_SERVED",
		11: "AD_GROUP_CRITERION_LOW_QUALITY",
		12: "AD_GROUP_CRITERION_UNDER_REVIEW",
		13: "AD_GROUP_CRITERION_PENDING_REVIEW",
		14: "AD_GROUP_CRITERION_BELOW_FIRST_PAGE_BID",
		15: "AD_GROUP_CRITERION_NEGATIVE",
		16: "AD_GROUP_CRITERION_RESTRICTED",
		17: "AD_GROUP_CRITERION_PAUSED",
		18: "AD_GROUP_CRITERION_PAUSED_DUE_TO_LOW_ACTIVITY",
		19: "AD_GROUP_CRITERION_REMOVED",
	}
	AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason_value = map[string]int32{
		"UNSPECIFIED":                                   0,
		"UNKNOWN":                                       1,
		"CAMPAIGN_PENDING":                              2,
		"CAMPAIGN_CRITERION_NEGATIVE":                   3,
		"CAMPAIGN_PAUSED":                               4,
		"CAMPAIGN_REMOVED":                              5,
		"CAMPAIGN_ENDED":                                6,
		"AD_GROUP_PAUSED":                               7,
		"AD_GROUP_REMOVED":                              8,
		"AD_GROUP_CRITERION_DISAPPROVED":                9,
		"AD_GROUP_CRITERION_RARELY_SERVED":              10,
		"AD_GROUP_CRITERION_LOW_QUALITY":                11,
		"AD_GROUP_CRITERION_UNDER_REVIEW":               12,
		"AD_GROUP_CRITERION_PENDING_REVIEW":             13,
		"AD_GROUP_CRITERION_BELOW_FIRST_PAGE_BID":       14,
		"AD_GROUP_CRITERION_NEGATIVE":                   15,
		"AD_GROUP_CRITERION_RESTRICTED":                 16,
		"AD_GROUP_CRITERION_PAUSED":                     17,
		"AD_GROUP_CRITERION_PAUSED_DUE_TO_LOW_ACTIVITY": 18,
		"AD_GROUP_CRITERION_REMOVED":                    19,
	}
)

func (x AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason) Enum() *AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason {
	p := new(AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason)
	*p = x
	return p
}

func (x AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_enumTypes[0].Descriptor()
}

func (AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_enumTypes[0]
}

func (x AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason.Descriptor instead.
func (AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ad group criterion primary status
// reasons.
type AdGroupCriterionPrimaryStatusReasonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdGroupCriterionPrimaryStatusReasonEnum) Reset() {
	*x = AdGroupCriterionPrimaryStatusReasonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdGroupCriterionPrimaryStatusReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupCriterionPrimaryStatusReasonEnum) ProtoMessage() {}

func (x *AdGroupCriterionPrimaryStatusReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupCriterionPrimaryStatusReasonEnum.ProtoReflect.Descriptor instead.
func (*AdGroupCriterionPrimaryStatusReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22,
	0xad, 0x05, 0x0a, 0x27, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x81, 0x05, 0x0a, 0x23,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4d, 0x50, 0x41,
	0x49, 0x47, 0x4e, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x45,
	0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41, 0x4d, 0x50,
	0x41, 0x49, 0x47, 0x4e, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a,
	0x10, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x44, 0x5f, 0x47, 0x52,
	0x4f, 0x55, 0x50, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44,
	0x10, 0x08, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x50, 0x50, 0x52,
	0x4f, 0x56, 0x45, 0x44, 0x10, 0x09, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x41, 0x52,
	0x45, 0x4c, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e,
	0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x0b,
	0x12, 0x23, 0x0a, 0x1f, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49,
	0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x0c, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x0d, 0x12, 0x2b, 0x0a, 0x27,
	0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49,
	0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x4c, 0x4f, 0x57, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x5f, 0x50,
	0x41, 0x47, 0x45, 0x5f, 0x42, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x44, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f,
	0x4e, 0x45, 0x47, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x0f, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x44,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x1d, 0x0a,
	0x19, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52,
	0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x11, 0x12, 0x31, 0x0a, 0x2d,
	0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49,
	0x4f, 0x4e, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x44, 0x55, 0x45, 0x5f, 0x54, 0x4f,
	0x5f, 0x4c, 0x4f, 0x57, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x12, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43, 0x52, 0x49, 0x54,
	0x45, 0x52, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x13, 0x42,
	0x82, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x28, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea,
	0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescData = file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDesc
)

func file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDescData
}

var file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_goTypes = []any{
	(AdGroupCriterionPrimaryStatusReasonEnum_AdGroupCriterionPrimaryStatusReason)(0), // 0: google.ads.googleads.v17.enums.AdGroupCriterionPrimaryStatusReasonEnum.AdGroupCriterionPrimaryStatusReason
	(*AdGroupCriterionPrimaryStatusReasonEnum)(nil),                                  // 1: google.ads.googleads.v17.enums.AdGroupCriterionPrimaryStatusReasonEnum
}
var file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_init()
}
func file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_init() {
	if File_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AdGroupCriterionPrimaryStatusReasonEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto = out.File
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_rawDesc = nil
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_goTypes = nil
	file_google_ads_googleads_v17_enums_ad_group_criterion_primary_status_reason_proto_depIdxs = nil
}
