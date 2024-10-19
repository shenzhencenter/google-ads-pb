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
// source: google/ads/googleads/v18/resources/android_privacy_shared_key_google_ad_group.proto

package resources

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// An Android privacy shared key view for Google ad group key.
type AndroidPrivacySharedKeyGoogleAdGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the Android privacy shared key.
	// Android privacy shared key resource names have the form:
	//
	// `customers/{customer_id}/androidPrivacySharedKeyGoogleAdGroups/{campaign_id}~{ad_group_id}~{android_privacy_interaction_type}~{android_privacy_network_type}~{android_privacy_interaction_date(yyyy-mm-dd)}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign ID used in the share key encoding.
	CampaignId int64 `protobuf:"varint,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Output only. The interaction type enum used in the share key encoding.
	AndroidPrivacyInteractionType enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType `protobuf:"varint,3,opt,name=android_privacy_interaction_type,json=androidPrivacyInteractionType,proto3,enum=google.ads.googleads.v18.enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType" json:"android_privacy_interaction_type,omitempty"`
	// Output only. The interaction date used in the shared key encoding in the
	// format of "YYYY-MM-DD" in UTC timezone.
	AndroidPrivacyInteractionDate string `protobuf:"bytes,4,opt,name=android_privacy_interaction_date,json=androidPrivacyInteractionDate,proto3" json:"android_privacy_interaction_date,omitempty"`
	// Output only. The network type enum used in the share key encoding.
	AndroidPrivacyNetworkType enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType `protobuf:"varint,5,opt,name=android_privacy_network_type,json=androidPrivacyNetworkType,proto3,enum=google.ads.googleads.v18.enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType" json:"android_privacy_network_type,omitempty"`
	// Output only. The ad group ID used in the share key encoding.
	AdGroupId int64 `protobuf:"varint,6,opt,name=ad_group_id,json=adGroupId,proto3" json:"ad_group_id,omitempty"`
	// Output only. 128 bit hex string of the encoded shared ad group key,
	// including a '0x' prefix. This key can be used to do a bitwise OR operator
	// with the aggregate conversion key to create a full aggregation key to
	// retrieve the Aggregate API Report in Android Privacy Sandbox.
	SharedAdGroupKey string `protobuf:"bytes,7,opt,name=shared_ad_group_key,json=sharedAdGroupKey,proto3" json:"shared_ad_group_key,omitempty"`
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) Reset() {
	*x = AndroidPrivacySharedKeyGoogleAdGroup{}
	mi := &file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidPrivacySharedKeyGoogleAdGroup) ProtoMessage() {}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidPrivacySharedKeyGoogleAdGroup.ProtoReflect.Descriptor instead.
func (*AndroidPrivacySharedKeyGoogleAdGroup) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescGZIP(), []int{0}
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetCampaignId() int64 {
	if x != nil {
		return x.CampaignId
	}
	return 0
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetAndroidPrivacyInteractionType() enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType {
	if x != nil {
		return x.AndroidPrivacyInteractionType
	}
	return enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType(0)
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetAndroidPrivacyInteractionDate() string {
	if x != nil {
		return x.AndroidPrivacyInteractionDate
	}
	return ""
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetAndroidPrivacyNetworkType() enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType {
	if x != nil {
		return x.AndroidPrivacyNetworkType
	}
	return enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType(0)
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetAdGroupId() int64 {
	if x != nil {
		return x.AdGroupId
	}
	return 0
}

func (x *AndroidPrivacySharedKeyGoogleAdGroup) GetSharedAdGroupKey() string {
	if x != nil {
		return x.SharedAdGroupKey
	}
	return ""
}

var File_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDesc = []byte{
	0x0a, 0x53, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb6, 0x07, 0x0a, 0x24, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x6a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x45, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x3f, 0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x49, 0x64, 0x12, 0xad, 0x01, 0x0a, 0x20, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x20, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x1c, 0x61, 0x6e, 0x64,
	0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x57, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x19, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0b, 0x61, 0x64, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x09, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x32, 0x0a,
	0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x10, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4b, 0x65,
	0x79, 0x3a, 0x84, 0x02, 0xea, 0x41, 0x80, 0x02, 0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0xbe, 0x01, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e, 0x7b, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x7d, 0x7e, 0x7b, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x7d, 0x42, 0x9b, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x42, 0x29, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x38, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescData = file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDesc
)

func file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDescData
}

var file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_goTypes = []any{
	(*AndroidPrivacySharedKeyGoogleAdGroup)(nil),                               // 0: google.ads.googleads.v18.resources.AndroidPrivacySharedKeyGoogleAdGroup
	(enums.AndroidPrivacyInteractionTypeEnum_AndroidPrivacyInteractionType)(0), // 1: google.ads.googleads.v18.enums.AndroidPrivacyInteractionTypeEnum.AndroidPrivacyInteractionType
	(enums.AndroidPrivacyNetworkTypeEnum_AndroidPrivacyNetworkType)(0),         // 2: google.ads.googleads.v18.enums.AndroidPrivacyNetworkTypeEnum.AndroidPrivacyNetworkType
}
var file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v18.resources.AndroidPrivacySharedKeyGoogleAdGroup.android_privacy_interaction_type:type_name -> google.ads.googleads.v18.enums.AndroidPrivacyInteractionTypeEnum.AndroidPrivacyInteractionType
	2, // 1: google.ads.googleads.v18.resources.AndroidPrivacySharedKeyGoogleAdGroup.android_privacy_network_type:type_name -> google.ads.googleads.v18.enums.AndroidPrivacyNetworkTypeEnum.AndroidPrivacyNetworkType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_init()
}
func file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_init() {
	if File_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto = out.File
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_rawDesc = nil
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_goTypes = nil
	file_google_ads_googleads_v18_resources_android_privacy_shared_key_google_ad_group_proto_depIdxs = nil
}
