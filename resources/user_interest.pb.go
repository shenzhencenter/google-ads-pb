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
// source: google/ads/googleads/v18/resources/user_interest.proto

package resources

import (
	common "github.com/shenzhencenter/google-ads-pb/common"
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

// A user interest: a particular interest-based vertical to be targeted.
type UserInterest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the user interest.
	// User interest resource names have the form:
	//
	// `customers/{customer_id}/userInterests/{user_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Taxonomy type of the user interest.
	TaxonomyType enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType `protobuf:"varint,2,opt,name=taxonomy_type,json=taxonomyType,proto3,enum=google.ads.googleads.v18.enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType" json:"taxonomy_type,omitempty"`
	// Output only. The ID of the user interest.
	UserInterestId *int64 `protobuf:"varint,8,opt,name=user_interest_id,json=userInterestId,proto3,oneof" json:"user_interest_id,omitempty"`
	// Output only. The name of the user interest.
	Name *string `protobuf:"bytes,9,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The parent of the user interest.
	UserInterestParent *string `protobuf:"bytes,10,opt,name=user_interest_parent,json=userInterestParent,proto3,oneof" json:"user_interest_parent,omitempty"`
	// Output only. True if the user interest is launched to all channels and
	// locales.
	LaunchedToAll *bool `protobuf:"varint,11,opt,name=launched_to_all,json=launchedToAll,proto3,oneof" json:"launched_to_all,omitempty"`
	// Output only. Availability information of the user interest.
	Availabilities []*common.CriterionCategoryAvailability `protobuf:"bytes,7,rep,name=availabilities,proto3" json:"availabilities,omitempty"`
}

func (x *UserInterest) Reset() {
	*x = UserInterest{}
	mi := &file_google_ads_googleads_v18_resources_user_interest_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInterest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInterest) ProtoMessage() {}

func (x *UserInterest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_resources_user_interest_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInterest.ProtoReflect.Descriptor instead.
func (*UserInterest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_resources_user_interest_proto_rawDescGZIP(), []int{0}
}

func (x *UserInterest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *UserInterest) GetTaxonomyType() enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType {
	if x != nil {
		return x.TaxonomyType
	}
	return enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType(0)
}

func (x *UserInterest) GetUserInterestId() int64 {
	if x != nil && x.UserInterestId != nil {
		return *x.UserInterestId
	}
	return 0
}

func (x *UserInterest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UserInterest) GetUserInterestParent() string {
	if x != nil && x.UserInterestParent != nil {
		return *x.UserInterestParent
	}
	return ""
}

func (x *UserInterest) GetLaunchedToAll() bool {
	if x != nil && x.LaunchedToAll != nil {
		return *x.LaunchedToAll
	}
	return false
}

func (x *UserInterest) GetAvailabilities() []*common.CriterionCategoryAvailability {
	if x != nil {
		return x.Availabilities
	}
	return nil
}

var File_google_ads_googleads_v18_resources_user_interest_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_resources_user_interest_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xeb, 0x05, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x03, 0xfa, 0x41,
	0x27, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7f, 0x0a, 0x0d, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x55, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x54, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x78, 0x6f, 0x6e,
	0x6f, 0x6d, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x64, 0x0a, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x27, 0x0a,
	0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x48, 0x02, 0x52, 0x12, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x30, 0x0a, 0x0f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x61,
	0x6c, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x03, 0x52,
	0x0d, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x64, 0x54, 0x6f, 0x41, 0x6c, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x6b, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x3a, 0x64,
	0xea, 0x41, 0x61, 0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x12, 0x38, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x6c, 0x42,
	0x83, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x11, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x38, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_resources_user_interest_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_resources_user_interest_proto_rawDescData = file_google_ads_googleads_v18_resources_user_interest_proto_rawDesc
)

func file_google_ads_googleads_v18_resources_user_interest_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_resources_user_interest_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_resources_user_interest_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_resources_user_interest_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_resources_user_interest_proto_rawDescData
}

var file_google_ads_googleads_v18_resources_user_interest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_resources_user_interest_proto_goTypes = []any{
	(*UserInterest)(nil), // 0: google.ads.googleads.v18.resources.UserInterest
	(enums.UserInterestTaxonomyTypeEnum_UserInterestTaxonomyType)(0), // 1: google.ads.googleads.v18.enums.UserInterestTaxonomyTypeEnum.UserInterestTaxonomyType
	(*common.CriterionCategoryAvailability)(nil),                     // 2: google.ads.googleads.v18.common.CriterionCategoryAvailability
}
var file_google_ads_googleads_v18_resources_user_interest_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v18.resources.UserInterest.taxonomy_type:type_name -> google.ads.googleads.v18.enums.UserInterestTaxonomyTypeEnum.UserInterestTaxonomyType
	2, // 1: google.ads.googleads.v18.resources.UserInterest.availabilities:type_name -> google.ads.googleads.v18.common.CriterionCategoryAvailability
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_resources_user_interest_proto_init() }
func file_google_ads_googleads_v18_resources_user_interest_proto_init() {
	if File_google_ads_googleads_v18_resources_user_interest_proto != nil {
		return
	}
	file_google_ads_googleads_v18_resources_user_interest_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_resources_user_interest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_resources_user_interest_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_resources_user_interest_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_resources_user_interest_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_resources_user_interest_proto = out.File
	file_google_ads_googleads_v18_resources_user_interest_proto_rawDesc = nil
	file_google_ads_googleads_v18_resources_user_interest_proto_goTypes = nil
	file_google_ads_googleads_v18_resources_user_interest_proto_depIdxs = nil
}
