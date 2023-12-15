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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v15/common/asset_policy.proto

package common

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
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

// Contains policy information for an asset inside an ad.
type AdAssetPolicySummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of policy findings for this asset.
	PolicyTopicEntries []*PolicyTopicEntry `protobuf:"bytes,1,rep,name=policy_topic_entries,json=policyTopicEntries,proto3" json:"policy_topic_entries,omitempty"`
	// Where in the review process this asset.
	ReviewStatus enums.PolicyReviewStatusEnum_PolicyReviewStatus `protobuf:"varint,2,opt,name=review_status,json=reviewStatus,proto3,enum=google.ads.googleads.v15.enums.PolicyReviewStatusEnum_PolicyReviewStatus" json:"review_status,omitempty"`
	// The overall approval status of this asset, which is calculated based on
	// the status of its individual policy topic entries.
	ApprovalStatus enums.PolicyApprovalStatusEnum_PolicyApprovalStatus `protobuf:"varint,3,opt,name=approval_status,json=approvalStatus,proto3,enum=google.ads.googleads.v15.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus" json:"approval_status,omitempty"`
}

func (x *AdAssetPolicySummary) Reset() {
	*x = AdAssetPolicySummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdAssetPolicySummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdAssetPolicySummary) ProtoMessage() {}

func (x *AdAssetPolicySummary) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdAssetPolicySummary.ProtoReflect.Descriptor instead.
func (*AdAssetPolicySummary) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_common_asset_policy_proto_rawDescGZIP(), []int{0}
}

func (x *AdAssetPolicySummary) GetPolicyTopicEntries() []*PolicyTopicEntry {
	if x != nil {
		return x.PolicyTopicEntries
	}
	return nil
}

func (x *AdAssetPolicySummary) GetReviewStatus() enums.PolicyReviewStatusEnum_PolicyReviewStatus {
	if x != nil {
		return x.ReviewStatus
	}
	return enums.PolicyReviewStatusEnum_PolicyReviewStatus(0)
}

func (x *AdAssetPolicySummary) GetApprovalStatus() enums.PolicyApprovalStatusEnum_PolicyApprovalStatus {
	if x != nil {
		return x.ApprovalStatus
	}
	return enums.PolicyApprovalStatusEnum_PolicyApprovalStatus(0)
}

// Provides the detail of a PrimaryStatus.
// Each asset link has a PrimaryStatus value (e.g. NOT_ELIGIBLE, meaning not
// serving), and list of corroborating PrimaryStatusReasons (e.g.
// [ASSET_DISAPPROVED]). Each reason may have some additional details
// annotated with it.  For instance, when the reason is ASSET_DISAPPROVED, the
// details field will contain additional information about the offline
// evaluation errors which led to the asset being disapproved.
type AssetLinkPrimaryStatusDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provides the reason of this PrimaryStatus.
	Reason *enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason `protobuf:"varint,1,opt,name=reason,proto3,enum=google.ads.googleads.v15.enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason,oneof" json:"reason,omitempty"`
	// Provides the PrimaryStatus of this status detail.
	Status *enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v15.enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus,oneof" json:"status,omitempty"`
	// Provides the details associated with the asset link primary status.
	//
	// Types that are assignable to Details:
	//
	//	*AssetLinkPrimaryStatusDetails_AssetDisapproved
	Details isAssetLinkPrimaryStatusDetails_Details `protobuf_oneof:"details"`
}

func (x *AssetLinkPrimaryStatusDetails) Reset() {
	*x = AssetLinkPrimaryStatusDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetLinkPrimaryStatusDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetLinkPrimaryStatusDetails) ProtoMessage() {}

func (x *AssetLinkPrimaryStatusDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetLinkPrimaryStatusDetails.ProtoReflect.Descriptor instead.
func (*AssetLinkPrimaryStatusDetails) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_common_asset_policy_proto_rawDescGZIP(), []int{1}
}

func (x *AssetLinkPrimaryStatusDetails) GetReason() enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason(0)
}

func (x *AssetLinkPrimaryStatusDetails) GetStatus() enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus(0)
}

func (m *AssetLinkPrimaryStatusDetails) GetDetails() isAssetLinkPrimaryStatusDetails_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (x *AssetLinkPrimaryStatusDetails) GetAssetDisapproved() *AssetDisapproved {
	if x, ok := x.GetDetails().(*AssetLinkPrimaryStatusDetails_AssetDisapproved); ok {
		return x.AssetDisapproved
	}
	return nil
}

type isAssetLinkPrimaryStatusDetails_Details interface {
	isAssetLinkPrimaryStatusDetails_Details()
}

type AssetLinkPrimaryStatusDetails_AssetDisapproved struct {
	// Provides the details for AssetLinkPrimaryStatusReason.ASSET_DISAPPROVED
	AssetDisapproved *AssetDisapproved `protobuf:"bytes,3,opt,name=asset_disapproved,json=assetDisapproved,proto3,oneof"`
}

func (*AssetLinkPrimaryStatusDetails_AssetDisapproved) isAssetLinkPrimaryStatusDetails_Details() {}

// Details related to AssetLinkPrimaryStatusReasonPB.ASSET_DISAPPROVED
type AssetDisapproved struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provides the quality evaluation disapproval reason of an asset.
	OfflineEvaluationErrorReasons []enums.AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons `protobuf:"varint,1,rep,packed,name=offline_evaluation_error_reasons,json=offlineEvaluationErrorReasons,proto3,enum=google.ads.googleads.v15.enums.AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons" json:"offline_evaluation_error_reasons,omitempty"`
}

func (x *AssetDisapproved) Reset() {
	*x = AssetDisapproved{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetDisapproved) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetDisapproved) ProtoMessage() {}

func (x *AssetDisapproved) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetDisapproved.ProtoReflect.Descriptor instead.
func (*AssetDisapproved) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_common_asset_policy_proto_rawDescGZIP(), []int{2}
}

func (x *AssetDisapproved) GetOfflineEvaluationErrorReasons() []enums.AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons {
	if x != nil {
		return x.OfflineEvaluationErrorReasons
	}
	return nil
}

var File_google_ads_googleads_v15_common_asset_policy_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_common_asset_policy_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x35, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe3, 0x02, 0x0a, 0x14, 0x41, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x63, 0x0a, 0x14, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x6e, 0x0a,
	0x0d, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x49, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x76, 0x0a,
	0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8e, 0x03, 0x0a, 0x1d, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x7a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x88, 0x01, 0x01, 0x12, 0x6e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x02, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x60, 0x0a, 0x11, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x69, 0x73,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x10, 0x61, 0x73, 0x73, 0x65, 0x74, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0xb2, 0x01, 0x0a, 0x20,
	0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x52, 0x1d, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x42, 0xf0, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x35, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_common_asset_policy_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_common_asset_policy_proto_rawDescData = file_google_ads_googleads_v15_common_asset_policy_proto_rawDesc
)

func file_google_ads_googleads_v15_common_asset_policy_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_common_asset_policy_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_common_asset_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_common_asset_policy_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_common_asset_policy_proto_rawDescData
}

var file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v15_common_asset_policy_proto_goTypes = []interface{}{
	(*AdAssetPolicySummary)(nil),                                                         // 0: google.ads.googleads.v15.common.AdAssetPolicySummary
	(*AssetLinkPrimaryStatusDetails)(nil),                                                // 1: google.ads.googleads.v15.common.AssetLinkPrimaryStatusDetails
	(*AssetDisapproved)(nil),                                                             // 2: google.ads.googleads.v15.common.AssetDisapproved
	(*PolicyTopicEntry)(nil),                                                             // 3: google.ads.googleads.v15.common.PolicyTopicEntry
	(enums.PolicyReviewStatusEnum_PolicyReviewStatus)(0),                                 // 4: google.ads.googleads.v15.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	(enums.PolicyApprovalStatusEnum_PolicyApprovalStatus)(0),                             // 5: google.ads.googleads.v15.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
	(enums.AssetLinkPrimaryStatusReasonEnum_AssetLinkPrimaryStatusReason)(0),             // 6: google.ads.googleads.v15.enums.AssetLinkPrimaryStatusReasonEnum.AssetLinkPrimaryStatusReason
	(enums.AssetLinkPrimaryStatusEnum_AssetLinkPrimaryStatus)(0),                         // 7: google.ads.googleads.v15.enums.AssetLinkPrimaryStatusEnum.AssetLinkPrimaryStatus
	(enums.AssetOfflineEvaluationErrorReasonsEnum_AssetOfflineEvaluationErrorReasons)(0), // 8: google.ads.googleads.v15.enums.AssetOfflineEvaluationErrorReasonsEnum.AssetOfflineEvaluationErrorReasons
}
var file_google_ads_googleads_v15_common_asset_policy_proto_depIdxs = []int32{
	3, // 0: google.ads.googleads.v15.common.AdAssetPolicySummary.policy_topic_entries:type_name -> google.ads.googleads.v15.common.PolicyTopicEntry
	4, // 1: google.ads.googleads.v15.common.AdAssetPolicySummary.review_status:type_name -> google.ads.googleads.v15.enums.PolicyReviewStatusEnum.PolicyReviewStatus
	5, // 2: google.ads.googleads.v15.common.AdAssetPolicySummary.approval_status:type_name -> google.ads.googleads.v15.enums.PolicyApprovalStatusEnum.PolicyApprovalStatus
	6, // 3: google.ads.googleads.v15.common.AssetLinkPrimaryStatusDetails.reason:type_name -> google.ads.googleads.v15.enums.AssetLinkPrimaryStatusReasonEnum.AssetLinkPrimaryStatusReason
	7, // 4: google.ads.googleads.v15.common.AssetLinkPrimaryStatusDetails.status:type_name -> google.ads.googleads.v15.enums.AssetLinkPrimaryStatusEnum.AssetLinkPrimaryStatus
	2, // 5: google.ads.googleads.v15.common.AssetLinkPrimaryStatusDetails.asset_disapproved:type_name -> google.ads.googleads.v15.common.AssetDisapproved
	8, // 6: google.ads.googleads.v15.common.AssetDisapproved.offline_evaluation_error_reasons:type_name -> google.ads.googleads.v15.enums.AssetOfflineEvaluationErrorReasonsEnum.AssetOfflineEvaluationErrorReasons
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_common_asset_policy_proto_init() }
func file_google_ads_googleads_v15_common_asset_policy_proto_init() {
	if File_google_ads_googleads_v15_common_asset_policy_proto != nil {
		return
	}
	file_google_ads_googleads_v15_common_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdAssetPolicySummary); i {
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
		file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetLinkPrimaryStatusDetails); i {
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
		file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetDisapproved); i {
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
	file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AssetLinkPrimaryStatusDetails_AssetDisapproved)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v15_common_asset_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_common_asset_policy_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_common_asset_policy_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_common_asset_policy_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_common_asset_policy_proto = out.File
	file_google_ads_googleads_v15_common_asset_policy_proto_rawDesc = nil
	file_google_ads_googleads_v15_common_asset_policy_proto_goTypes = nil
	file_google_ads_googleads_v15_common_asset_policy_proto_depIdxs = nil
}
