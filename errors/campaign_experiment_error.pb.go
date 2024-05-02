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
// source: google/ads/googleads/v16/errors/campaign_experiment_error.proto

package errors

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

// Enum describing possible campaign experiment errors.
type CampaignExperimentErrorEnum_CampaignExperimentError int32

const (
	// Enum unspecified.
	CampaignExperimentErrorEnum_UNSPECIFIED CampaignExperimentErrorEnum_CampaignExperimentError = 0
	// The received error code is not known in this version.
	CampaignExperimentErrorEnum_UNKNOWN CampaignExperimentErrorEnum_CampaignExperimentError = 1
	// An active campaign or experiment with this name already exists.
	CampaignExperimentErrorEnum_DUPLICATE_NAME CampaignExperimentErrorEnum_CampaignExperimentError = 2
	// Experiment cannot be updated from the current state to the
	// requested target state. For example, an experiment can only graduate
	// if its status is ENABLED.
	CampaignExperimentErrorEnum_INVALID_TRANSITION CampaignExperimentErrorEnum_CampaignExperimentError = 3
	// Cannot create an experiment from a campaign using an explicitly shared
	// budget.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET CampaignExperimentErrorEnum_CampaignExperimentError = 4
	// Cannot create an experiment for a removed base campaign.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN CampaignExperimentErrorEnum_CampaignExperimentError = 5
	// Cannot create an experiment from a draft, which has a status other than
	// proposed.
	CampaignExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT CampaignExperimentErrorEnum_CampaignExperimentError = 6
	// This customer is not allowed to create an experiment.
	CampaignExperimentErrorEnum_CUSTOMER_CANNOT_CREATE_EXPERIMENT CampaignExperimentErrorEnum_CampaignExperimentError = 7
	// This campaign is not allowed to create an experiment.
	CampaignExperimentErrorEnum_CAMPAIGN_CANNOT_CREATE_EXPERIMENT CampaignExperimentErrorEnum_CampaignExperimentError = 8
	// Trying to set an experiment duration which overlaps with another
	// experiment.
	CampaignExperimentErrorEnum_EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP CampaignExperimentErrorEnum_CampaignExperimentError = 9
	// All non-removed experiments must start and end within their campaign's
	// duration.
	CampaignExperimentErrorEnum_EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION CampaignExperimentErrorEnum_CampaignExperimentError = 10
	// The experiment cannot be modified because its status is in a terminal
	// state, such as REMOVED.
	CampaignExperimentErrorEnum_CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS CampaignExperimentErrorEnum_CampaignExperimentError = 11
)

// Enum value maps for CampaignExperimentErrorEnum_CampaignExperimentError.
var (
	CampaignExperimentErrorEnum_CampaignExperimentError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DUPLICATE_NAME",
		3:  "INVALID_TRANSITION",
		4:  "CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET",
		5:  "CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN",
		6:  "CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT",
		7:  "CUSTOMER_CANNOT_CREATE_EXPERIMENT",
		8:  "CAMPAIGN_CANNOT_CREATE_EXPERIMENT",
		9:  "EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP",
		10: "EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION",
		11: "CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS",
	}
	CampaignExperimentErrorEnum_CampaignExperimentError_value = map[string]int32{
		"UNSPECIFIED":        0,
		"UNKNOWN":            1,
		"DUPLICATE_NAME":     2,
		"INVALID_TRANSITION": 3,
		"CANNOT_CREATE_EXPERIMENT_WITH_SHARED_BUDGET":          4,
		"CANNOT_CREATE_EXPERIMENT_FOR_REMOVED_BASE_CAMPAIGN":   5,
		"CANNOT_CREATE_EXPERIMENT_FOR_NON_PROPOSED_DRAFT":      6,
		"CUSTOMER_CANNOT_CREATE_EXPERIMENT":                    7,
		"CAMPAIGN_CANNOT_CREATE_EXPERIMENT":                    8,
		"EXPERIMENT_DURATIONS_MUST_NOT_OVERLAP":                9,
		"EXPERIMENT_DURATION_MUST_BE_WITHIN_CAMPAIGN_DURATION": 10,
		"CANNOT_MUTATE_EXPERIMENT_DUE_TO_STATUS":               11,
	}
)

func (x CampaignExperimentErrorEnum_CampaignExperimentError) Enum() *CampaignExperimentErrorEnum_CampaignExperimentError {
	p := new(CampaignExperimentErrorEnum_CampaignExperimentError)
	*p = x
	return p
}

func (x CampaignExperimentErrorEnum_CampaignExperimentError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignExperimentErrorEnum_CampaignExperimentError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_enumTypes[0].Descriptor()
}

func (CampaignExperimentErrorEnum_CampaignExperimentError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_enumTypes[0]
}

func (x CampaignExperimentErrorEnum_CampaignExperimentError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignExperimentErrorEnum_CampaignExperimentError.Descriptor instead.
func (CampaignExperimentErrorEnum_CampaignExperimentError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign experiment errors.
type CampaignExperimentErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CampaignExperimentErrorEnum) Reset() {
	*x = CampaignExperimentErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignExperimentErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignExperimentErrorEnum) ProtoMessage() {}

func (x *CampaignExperimentErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignExperimentErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignExperimentErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_errors_campaign_experiment_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x22, 0x80, 0x04, 0x0a, 0x1b, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xe0, 0x03, 0x0a, 0x17, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44,
	0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x04, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52,
	0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x44, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10,
	0x05, 0x12, 0x33, 0x0a, 0x2f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x45, 0x44, 0x5f, 0x44,
	0x52, 0x41, 0x46, 0x54, 0x10, 0x06, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x25, 0x0a,
	0x21, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x08, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x4d, 0x55, 0x53,
	0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x10, 0x09, 0x12,
	0x38, 0x0a, 0x34, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x55,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x55, 0x53, 0x54, 0x5f, 0x42, 0x45, 0x5f, 0x57,
	0x49, 0x54, 0x48, 0x49, 0x4e, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x44,
	0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52,
	0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x10, 0x0b, 0x42, 0xfc, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1c, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02,
	0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescData = file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDesc
)

func file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDescData
}

var file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_goTypes = []interface{}{
	(CampaignExperimentErrorEnum_CampaignExperimentError)(0), // 0: google.ads.googleads.v16.errors.CampaignExperimentErrorEnum.CampaignExperimentError
	(*CampaignExperimentErrorEnum)(nil),                      // 1: google.ads.googleads.v16.errors.CampaignExperimentErrorEnum
}
var file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_init() }
func file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_init() {
	if File_google_ads_googleads_v16_errors_campaign_experiment_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignExperimentErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_errors_campaign_experiment_error_proto = out.File
	file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_rawDesc = nil
	file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_goTypes = nil
	file_google_ads_googleads_v16_errors_campaign_experiment_error_proto_depIdxs = nil
}
