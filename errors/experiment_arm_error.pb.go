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
// source: google/ads/googleads/v17/errors/experiment_arm_error.proto

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

// Enum describing possible experiment arm errors.
type ExperimentArmErrorEnum_ExperimentArmError int32

const (
	// Enum unspecified.
	ExperimentArmErrorEnum_UNSPECIFIED ExperimentArmErrorEnum_ExperimentArmError = 0
	// The received error code is not known in this version.
	ExperimentArmErrorEnum_UNKNOWN ExperimentArmErrorEnum_ExperimentArmError = 1
	// Number of experiment arms is above limit.
	ExperimentArmErrorEnum_EXPERIMENT_ARM_COUNT_LIMIT_EXCEEDED ExperimentArmErrorEnum_ExperimentArmError = 2
	// Cannot add campaign with invalid status to the experiment arm.
	ExperimentArmErrorEnum_INVALID_CAMPAIGN_STATUS ExperimentArmErrorEnum_ExperimentArmError = 3
	// Cannot add duplicate experiment arm name in one experiment.
	ExperimentArmErrorEnum_DUPLICATE_EXPERIMENT_ARM_NAME ExperimentArmErrorEnum_ExperimentArmError = 4
	// Cannot set campaigns of treatment experiment arm.
	ExperimentArmErrorEnum_CANNOT_SET_TREATMENT_ARM_CAMPAIGN ExperimentArmErrorEnum_ExperimentArmError = 5
	// Cannot edit campaign ids in trial arms in non SETUP experiment.
	ExperimentArmErrorEnum_CANNOT_MODIFY_CAMPAIGN_IDS ExperimentArmErrorEnum_ExperimentArmError = 6
	// Cannot modify the campaigns in the control arm
	// if there is not a suffix set in the trial.
	ExperimentArmErrorEnum_CANNOT_MODIFY_CAMPAIGN_WITHOUT_SUFFIX_SET ExperimentArmErrorEnum_ExperimentArmError = 7
	// Traffic split related settings (like traffic share bounds) can't be
	// modified after the trial has started.
	ExperimentArmErrorEnum_CANNOT_MUTATE_TRAFFIC_SPLIT_AFTER_START ExperimentArmErrorEnum_ExperimentArmError = 8
	// Cannot use shared budget on experiment's control campaign.
	ExperimentArmErrorEnum_CANNOT_ADD_CAMPAIGN_WITH_SHARED_BUDGET ExperimentArmErrorEnum_ExperimentArmError = 9
	// Cannot use custom budget on experiment's control campaigns.
	ExperimentArmErrorEnum_CANNOT_ADD_CAMPAIGN_WITH_CUSTOM_BUDGET ExperimentArmErrorEnum_ExperimentArmError = 10
	// Cannot have enable_dynamic_assets turned on in experiment's campaigns.
	ExperimentArmErrorEnum_CANNOT_ADD_CAMPAIGNS_WITH_DYNAMIC_ASSETS_ENABLED ExperimentArmErrorEnum_ExperimentArmError = 11
	// Cannot use campaign's advertising channel sub type in experiment.
	ExperimentArmErrorEnum_UNSUPPORTED_CAMPAIGN_ADVERTISING_CHANNEL_SUB_TYPE ExperimentArmErrorEnum_ExperimentArmError = 12
	// Experiment date range must be within base campaign's date range.
	ExperimentArmErrorEnum_CANNOT_ADD_BASE_CAMPAIGN_WITH_DATE_RANGE ExperimentArmErrorEnum_ExperimentArmError = 13
	// Bidding strategy is not supported in experiments.
	ExperimentArmErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_IN_EXPERIMENTS ExperimentArmErrorEnum_ExperimentArmError = 14
	// Traffic split is not supported for some channel types.
	ExperimentArmErrorEnum_TRAFFIC_SPLIT_NOT_SUPPORTED_FOR_CHANNEL_TYPE ExperimentArmErrorEnum_ExperimentArmError = 15
)

// Enum value maps for ExperimentArmErrorEnum_ExperimentArmError.
var (
	ExperimentArmErrorEnum_ExperimentArmError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "EXPERIMENT_ARM_COUNT_LIMIT_EXCEEDED",
		3:  "INVALID_CAMPAIGN_STATUS",
		4:  "DUPLICATE_EXPERIMENT_ARM_NAME",
		5:  "CANNOT_SET_TREATMENT_ARM_CAMPAIGN",
		6:  "CANNOT_MODIFY_CAMPAIGN_IDS",
		7:  "CANNOT_MODIFY_CAMPAIGN_WITHOUT_SUFFIX_SET",
		8:  "CANNOT_MUTATE_TRAFFIC_SPLIT_AFTER_START",
		9:  "CANNOT_ADD_CAMPAIGN_WITH_SHARED_BUDGET",
		10: "CANNOT_ADD_CAMPAIGN_WITH_CUSTOM_BUDGET",
		11: "CANNOT_ADD_CAMPAIGNS_WITH_DYNAMIC_ASSETS_ENABLED",
		12: "UNSUPPORTED_CAMPAIGN_ADVERTISING_CHANNEL_SUB_TYPE",
		13: "CANNOT_ADD_BASE_CAMPAIGN_WITH_DATE_RANGE",
		14: "BIDDING_STRATEGY_NOT_SUPPORTED_IN_EXPERIMENTS",
		15: "TRAFFIC_SPLIT_NOT_SUPPORTED_FOR_CHANNEL_TYPE",
	}
	ExperimentArmErrorEnum_ExperimentArmError_value = map[string]int32{
		"UNSPECIFIED":                                       0,
		"UNKNOWN":                                           1,
		"EXPERIMENT_ARM_COUNT_LIMIT_EXCEEDED":               2,
		"INVALID_CAMPAIGN_STATUS":                           3,
		"DUPLICATE_EXPERIMENT_ARM_NAME":                     4,
		"CANNOT_SET_TREATMENT_ARM_CAMPAIGN":                 5,
		"CANNOT_MODIFY_CAMPAIGN_IDS":                        6,
		"CANNOT_MODIFY_CAMPAIGN_WITHOUT_SUFFIX_SET":         7,
		"CANNOT_MUTATE_TRAFFIC_SPLIT_AFTER_START":           8,
		"CANNOT_ADD_CAMPAIGN_WITH_SHARED_BUDGET":            9,
		"CANNOT_ADD_CAMPAIGN_WITH_CUSTOM_BUDGET":            10,
		"CANNOT_ADD_CAMPAIGNS_WITH_DYNAMIC_ASSETS_ENABLED":  11,
		"UNSUPPORTED_CAMPAIGN_ADVERTISING_CHANNEL_SUB_TYPE": 12,
		"CANNOT_ADD_BASE_CAMPAIGN_WITH_DATE_RANGE":          13,
		"BIDDING_STRATEGY_NOT_SUPPORTED_IN_EXPERIMENTS":     14,
		"TRAFFIC_SPLIT_NOT_SUPPORTED_FOR_CHANNEL_TYPE":      15,
	}
)

func (x ExperimentArmErrorEnum_ExperimentArmError) Enum() *ExperimentArmErrorEnum_ExperimentArmError {
	p := new(ExperimentArmErrorEnum_ExperimentArmError)
	*p = x
	return p
}

func (x ExperimentArmErrorEnum_ExperimentArmError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExperimentArmErrorEnum_ExperimentArmError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_experiment_arm_error_proto_enumTypes[0].Descriptor()
}

func (ExperimentArmErrorEnum_ExperimentArmError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_experiment_arm_error_proto_enumTypes[0]
}

func (x ExperimentArmErrorEnum_ExperimentArmError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExperimentArmErrorEnum_ExperimentArmError.Descriptor instead.
func (ExperimentArmErrorEnum_ExperimentArmError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible experiment arm error.
type ExperimentArmErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExperimentArmErrorEnum) Reset() {
	*x = ExperimentArmErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_experiment_arm_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentArmErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentArmErrorEnum) ProtoMessage() {}

func (x *ExperimentArmErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_experiment_arm_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentArmErrorEnum.ProtoReflect.Descriptor instead.
func (*ExperimentArmErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_experiment_arm_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x6d,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xb1, 0x05,
	0x0a, 0x16, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x6d, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x96, 0x05, 0x0a, 0x12, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x27, 0x0a,
	0x23, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x52, 0x4d, 0x5f,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58, 0x43, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x52, 0x4d, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x53, 0x45, 0x54, 0x5f, 0x54, 0x52, 0x45, 0x41, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x52, 0x4d, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x10, 0x05, 0x12, 0x1e, 0x0a,
	0x1a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x06, 0x12, 0x2d, 0x0a,
	0x29, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x43,
	0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f,
	0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x07, 0x12, 0x2b, 0x0a, 0x27,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x52,
	0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x5f, 0x41, 0x46, 0x54, 0x45,
	0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x08, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x10, 0x09, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x44, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10,
	0x0a, 0x12, 0x34, 0x0a, 0x30, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x53, 0x5f, 0x45, 0x4e,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x35, 0x0a, 0x31, 0x55, 0x4e, 0x53, 0x55, 0x50,
	0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f,
	0x41, 0x44, 0x56, 0x45, 0x52, 0x54, 0x49, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0c, 0x12, 0x2c,
	0x0a, 0x28, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x42, 0x41, 0x53,
	0x45, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0d, 0x12, 0x31, 0x0a, 0x2d,
	0x42, 0x49, 0x44, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x49,
	0x4e, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0e, 0x12,
	0x30, 0x0a, 0x2c, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x0f, 0x42, 0xf7, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x17, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescData = file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_experiment_arm_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_experiment_arm_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_experiment_arm_error_proto_goTypes = []any{
	(ExperimentArmErrorEnum_ExperimentArmError)(0), // 0: google.ads.googleads.v17.errors.ExperimentArmErrorEnum.ExperimentArmError
	(*ExperimentArmErrorEnum)(nil),                 // 1: google.ads.googleads.v17.errors.ExperimentArmErrorEnum
}
var file_google_ads_googleads_v17_errors_experiment_arm_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_experiment_arm_error_proto_init() }
func file_google_ads_googleads_v17_errors_experiment_arm_error_proto_init() {
	if File_google_ads_googleads_v17_errors_experiment_arm_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_experiment_arm_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExperimentArmErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_experiment_arm_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_experiment_arm_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_experiment_arm_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_experiment_arm_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_experiment_arm_error_proto = out.File
	file_google_ads_googleads_v17_errors_experiment_arm_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_experiment_arm_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_experiment_arm_error_proto_depIdxs = nil
}
