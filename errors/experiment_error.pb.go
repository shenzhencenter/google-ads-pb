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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: google/ads/googleads/v17/errors/experiment_error.proto

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

// Enum describing possible experiment errors.
type ExperimentErrorEnum_ExperimentError int32

const (
	// Enum unspecified.
	ExperimentErrorEnum_UNSPECIFIED ExperimentErrorEnum_ExperimentError = 0
	// The received error code is not known in this version.
	ExperimentErrorEnum_UNKNOWN ExperimentErrorEnum_ExperimentError = 1
	// The start date of an experiment cannot be set in the past.
	// Use a start date in the future.
	ExperimentErrorEnum_CANNOT_SET_START_DATE_IN_PAST ExperimentErrorEnum_ExperimentError = 2
	// The end date of an experiment is before its start date.
	// Use an end date after the start date.
	ExperimentErrorEnum_END_DATE_BEFORE_START_DATE ExperimentErrorEnum_ExperimentError = 3
	// The start date of an experiment is too far in the future.
	// Use a start date no more than 1 year in the future.
	ExperimentErrorEnum_START_DATE_TOO_FAR_IN_FUTURE ExperimentErrorEnum_ExperimentError = 4
	// The experiment has the same name as an existing active experiment.
	ExperimentErrorEnum_DUPLICATE_EXPERIMENT_NAME ExperimentErrorEnum_ExperimentError = 5
	// Experiments can only be modified when they are ENABLED.
	ExperimentErrorEnum_CANNOT_MODIFY_REMOVED_EXPERIMENT ExperimentErrorEnum_ExperimentError = 6
	// The start date of an experiment cannot be modified if the existing start
	// date has already passed.
	ExperimentErrorEnum_START_DATE_ALREADY_PASSED ExperimentErrorEnum_ExperimentError = 7
	// The end date of an experiment cannot be set in the past.
	ExperimentErrorEnum_CANNOT_SET_END_DATE_IN_PAST ExperimentErrorEnum_ExperimentError = 8
	// The status of an experiment cannot be set to REMOVED.
	ExperimentErrorEnum_CANNOT_SET_STATUS_TO_REMOVED ExperimentErrorEnum_ExperimentError = 9
	// The end date of an expired experiment cannot be modified.
	ExperimentErrorEnum_CANNOT_MODIFY_PAST_END_DATE ExperimentErrorEnum_ExperimentError = 10
	// The status is invalid.
	ExperimentErrorEnum_INVALID_STATUS ExperimentErrorEnum_ExperimentError = 11
	// Experiment arm contains campaigns with invalid advertising channel type.
	ExperimentErrorEnum_INVALID_CAMPAIGN_CHANNEL_TYPE ExperimentErrorEnum_ExperimentError = 12
	// A pair of trials share members and have overlapping date ranges.
	ExperimentErrorEnum_OVERLAPPING_MEMBERS_AND_DATE_RANGE ExperimentErrorEnum_ExperimentError = 13
	// Experiment arm contains invalid traffic split.
	ExperimentErrorEnum_INVALID_TRIAL_ARM_TRAFFIC_SPLIT ExperimentErrorEnum_ExperimentError = 14
	// Experiment contains trial arms with overlapping traffic split.
	ExperimentErrorEnum_TRAFFIC_SPLIT_OVERLAPPING ExperimentErrorEnum_ExperimentError = 15
	// The total traffic split of trial arms is not equal to 100.
	ExperimentErrorEnum_SUM_TRIAL_ARM_TRAFFIC_UNEQUALS_TO_TRIAL_TRAFFIC_SPLIT_DENOMINATOR ExperimentErrorEnum_ExperimentError = 16
	// Traffic split related settings (like traffic share bounds) can't be
	// modified after the experiment has started.
	ExperimentErrorEnum_CANNOT_MODIFY_TRAFFIC_SPLIT_AFTER_START ExperimentErrorEnum_ExperimentError = 17
	// The experiment could not be found.
	ExperimentErrorEnum_EXPERIMENT_NOT_FOUND ExperimentErrorEnum_ExperimentError = 18
	// Experiment has not begun.
	ExperimentErrorEnum_EXPERIMENT_NOT_YET_STARTED ExperimentErrorEnum_ExperimentError = 19
	// The experiment cannot have more than one control arm.
	ExperimentErrorEnum_CANNOT_HAVE_MULTIPLE_CONTROL_ARMS ExperimentErrorEnum_ExperimentError = 20
	// The experiment doesn't set in-design campaigns.
	ExperimentErrorEnum_IN_DESIGN_CAMPAIGNS_NOT_SET ExperimentErrorEnum_ExperimentError = 21
	// Clients must use the graduate action to graduate experiments and cannot
	// set the status to GRADUATED directly.
	ExperimentErrorEnum_CANNOT_SET_STATUS_TO_GRADUATED ExperimentErrorEnum_ExperimentError = 22
	// Cannot use shared budget on base campaign when scheduling an experiment.
	ExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_CAMPAIGN_WITH_SHARED_BUDGET ExperimentErrorEnum_ExperimentError = 23
	// Cannot use custom budget on base campaign when scheduling an experiment.
	ExperimentErrorEnum_CANNOT_CREATE_EXPERIMENT_CAMPAIGN_WITH_CUSTOM_BUDGET ExperimentErrorEnum_ExperimentError = 24
	// Invalid status transition.
	ExperimentErrorEnum_STATUS_TRANSITION_INVALID ExperimentErrorEnum_ExperimentError = 25
	// The experiment campaign name conflicts with a pre-existing campaign.
	ExperimentErrorEnum_DUPLICATE_EXPERIMENT_CAMPAIGN_NAME ExperimentErrorEnum_ExperimentError = 26
	// Cannot remove in creation experiments.
	ExperimentErrorEnum_CANNOT_REMOVE_IN_CREATION_EXPERIMENT ExperimentErrorEnum_ExperimentError = 27
	// Cannot add campaign with deprecated ad types. Deprecated ad types:
	// ENHANCED_DISPLAY, GALLERY, GMAIL, KEYWORDLESS, TEXT.
	ExperimentErrorEnum_CANNOT_ADD_CAMPAIGN_WITH_DEPRECATED_AD_TYPES ExperimentErrorEnum_ExperimentError = 28
	// Sync can only be enabled for supported experiment types. Supported
	// experiment types: SEARCH_CUSTOM, DISPLAY_CUSTOM,
	// DISPLAY_AUTOMATED_BIDDING_STRATEGY, SEARCH_AUTOMATED_BIDDING_STRATEGY.
	ExperimentErrorEnum_CANNOT_ENABLE_SYNC_FOR_UNSUPPORTED_EXPERIMENT_TYPE ExperimentErrorEnum_ExperimentError = 29
	// Experiment length cannot be longer than max length.
	ExperimentErrorEnum_INVALID_DURATION_FOR_AN_EXPERIMENT ExperimentErrorEnum_ExperimentError = 30
)

// Enum value maps for ExperimentErrorEnum_ExperimentError.
var (
	ExperimentErrorEnum_ExperimentError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CANNOT_SET_START_DATE_IN_PAST",
		3:  "END_DATE_BEFORE_START_DATE",
		4:  "START_DATE_TOO_FAR_IN_FUTURE",
		5:  "DUPLICATE_EXPERIMENT_NAME",
		6:  "CANNOT_MODIFY_REMOVED_EXPERIMENT",
		7:  "START_DATE_ALREADY_PASSED",
		8:  "CANNOT_SET_END_DATE_IN_PAST",
		9:  "CANNOT_SET_STATUS_TO_REMOVED",
		10: "CANNOT_MODIFY_PAST_END_DATE",
		11: "INVALID_STATUS",
		12: "INVALID_CAMPAIGN_CHANNEL_TYPE",
		13: "OVERLAPPING_MEMBERS_AND_DATE_RANGE",
		14: "INVALID_TRIAL_ARM_TRAFFIC_SPLIT",
		15: "TRAFFIC_SPLIT_OVERLAPPING",
		16: "SUM_TRIAL_ARM_TRAFFIC_UNEQUALS_TO_TRIAL_TRAFFIC_SPLIT_DENOMINATOR",
		17: "CANNOT_MODIFY_TRAFFIC_SPLIT_AFTER_START",
		18: "EXPERIMENT_NOT_FOUND",
		19: "EXPERIMENT_NOT_YET_STARTED",
		20: "CANNOT_HAVE_MULTIPLE_CONTROL_ARMS",
		21: "IN_DESIGN_CAMPAIGNS_NOT_SET",
		22: "CANNOT_SET_STATUS_TO_GRADUATED",
		23: "CANNOT_CREATE_EXPERIMENT_CAMPAIGN_WITH_SHARED_BUDGET",
		24: "CANNOT_CREATE_EXPERIMENT_CAMPAIGN_WITH_CUSTOM_BUDGET",
		25: "STATUS_TRANSITION_INVALID",
		26: "DUPLICATE_EXPERIMENT_CAMPAIGN_NAME",
		27: "CANNOT_REMOVE_IN_CREATION_EXPERIMENT",
		28: "CANNOT_ADD_CAMPAIGN_WITH_DEPRECATED_AD_TYPES",
		29: "CANNOT_ENABLE_SYNC_FOR_UNSUPPORTED_EXPERIMENT_TYPE",
		30: "INVALID_DURATION_FOR_AN_EXPERIMENT",
	}
	ExperimentErrorEnum_ExperimentError_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"CANNOT_SET_START_DATE_IN_PAST":      2,
		"END_DATE_BEFORE_START_DATE":         3,
		"START_DATE_TOO_FAR_IN_FUTURE":       4,
		"DUPLICATE_EXPERIMENT_NAME":          5,
		"CANNOT_MODIFY_REMOVED_EXPERIMENT":   6,
		"START_DATE_ALREADY_PASSED":          7,
		"CANNOT_SET_END_DATE_IN_PAST":        8,
		"CANNOT_SET_STATUS_TO_REMOVED":       9,
		"CANNOT_MODIFY_PAST_END_DATE":        10,
		"INVALID_STATUS":                     11,
		"INVALID_CAMPAIGN_CHANNEL_TYPE":      12,
		"OVERLAPPING_MEMBERS_AND_DATE_RANGE": 13,
		"INVALID_TRIAL_ARM_TRAFFIC_SPLIT":    14,
		"TRAFFIC_SPLIT_OVERLAPPING":          15,
		"SUM_TRIAL_ARM_TRAFFIC_UNEQUALS_TO_TRIAL_TRAFFIC_SPLIT_DENOMINATOR": 16,
		"CANNOT_MODIFY_TRAFFIC_SPLIT_AFTER_START":                           17,
		"EXPERIMENT_NOT_FOUND":                                              18,
		"EXPERIMENT_NOT_YET_STARTED":                                        19,
		"CANNOT_HAVE_MULTIPLE_CONTROL_ARMS":                                 20,
		"IN_DESIGN_CAMPAIGNS_NOT_SET":                                       21,
		"CANNOT_SET_STATUS_TO_GRADUATED":                                    22,
		"CANNOT_CREATE_EXPERIMENT_CAMPAIGN_WITH_SHARED_BUDGET":              23,
		"CANNOT_CREATE_EXPERIMENT_CAMPAIGN_WITH_CUSTOM_BUDGET":              24,
		"STATUS_TRANSITION_INVALID":                                         25,
		"DUPLICATE_EXPERIMENT_CAMPAIGN_NAME":                                26,
		"CANNOT_REMOVE_IN_CREATION_EXPERIMENT":                              27,
		"CANNOT_ADD_CAMPAIGN_WITH_DEPRECATED_AD_TYPES":                      28,
		"CANNOT_ENABLE_SYNC_FOR_UNSUPPORTED_EXPERIMENT_TYPE":                29,
		"INVALID_DURATION_FOR_AN_EXPERIMENT":                                30,
	}
)

func (x ExperimentErrorEnum_ExperimentError) Enum() *ExperimentErrorEnum_ExperimentError {
	p := new(ExperimentErrorEnum_ExperimentError)
	*p = x
	return p
}

func (x ExperimentErrorEnum_ExperimentError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExperimentErrorEnum_ExperimentError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_experiment_error_proto_enumTypes[0].Descriptor()
}

func (ExperimentErrorEnum_ExperimentError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_experiment_error_proto_enumTypes[0]
}

func (x ExperimentErrorEnum_ExperimentError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExperimentErrorEnum_ExperimentError.Descriptor instead.
func (ExperimentErrorEnum_ExperimentError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible experiment error.
type ExperimentErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExperimentErrorEnum) Reset() {
	*x = ExperimentErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_experiment_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentErrorEnum) ProtoMessage() {}

func (x *ExperimentErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_experiment_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentErrorEnum.ProtoReflect.Descriptor instead.
func (*ExperimentErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_experiment_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_experiment_error_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xa8, 0x09, 0x0a, 0x13, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x90, 0x09, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f,
	0x50, 0x41, 0x53, 0x54, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x4e, 0x44, 0x5f, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x42, 0x45, 0x46, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x46, 0x41, 0x52, 0x5f, 0x49, 0x4e, 0x5f,
	0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x55, 0x50, 0x4c,
	0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x05, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44,
	0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x1d, 0x0a,
	0x19, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x45, 0x44, 0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x41, 0x53, 0x54, 0x10, 0x08, 0x12, 0x20, 0x0a,
	0x1c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x54, 0x4f, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59,
	0x5f, 0x50, 0x41, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x0a,
	0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x10, 0x0b, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0c, 0x12, 0x26, 0x0a, 0x22, 0x4f, 0x56, 0x45, 0x52, 0x4c,
	0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x5f, 0x41,
	0x4e, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0d, 0x12,
	0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c,
	0x5f, 0x41, 0x52, 0x4d, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x53, 0x50, 0x4c,
	0x49, 0x54, 0x10, 0x0e, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f,
	0x53, 0x50, 0x4c, 0x49, 0x54, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x50, 0x49, 0x4e,
	0x47, 0x10, 0x0f, 0x12, 0x45, 0x0a, 0x41, 0x53, 0x55, 0x4d, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c,
	0x5f, 0x41, 0x52, 0x4d, 0x5f, 0x54, 0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x55, 0x4e, 0x45,
	0x51, 0x55, 0x41, 0x4c, 0x53, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x54,
	0x52, 0x41, 0x46, 0x46, 0x49, 0x43, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x5f, 0x44, 0x45, 0x4e,
	0x4f, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x4f, 0x52, 0x10, 0x10, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x41,
	0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x54, 0x52, 0x41, 0x46,
	0x46, 0x49, 0x43, 0x5f, 0x53, 0x50, 0x4c, 0x49, 0x54, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x11, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x50, 0x45, 0x52,
	0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x12, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x59, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x13, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x48, 0x41, 0x56, 0x45,
	0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f,
	0x4c, 0x5f, 0x41, 0x52, 0x4d, 0x53, 0x10, 0x14, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e, 0x5f, 0x44,
	0x45, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x15, 0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54,
	0x4f, 0x5f, 0x47, 0x52, 0x41, 0x44, 0x55, 0x41, 0x54, 0x45, 0x44, 0x10, 0x16, 0x12, 0x38, 0x0a,
	0x34, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45,
	0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49,
	0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x42,
	0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x17, 0x12, 0x38, 0x0a, 0x34, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10,
	0x18, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x19,
	0x12, 0x26, 0x0a, 0x22, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58,
	0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47,
	0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x1a, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x1b, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44,
	0x5f, 0x43, 0x41, 0x4d, 0x50, 0x41, 0x49, 0x47, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x44,
	0x45, 0x50, 0x52, 0x45, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x53, 0x10, 0x1c, 0x12, 0x36, 0x0a, 0x32, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x55,
	0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52,
	0x49, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x1d, 0x12, 0x26, 0x0a, 0x22,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x45, 0x52, 0x49, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x1e, 0x42, 0xf4, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x14, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
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
	file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescData = file_google_ads_googleads_v17_errors_experiment_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_experiment_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_experiment_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_experiment_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_experiment_error_proto_goTypes = []interface{}{
	(ExperimentErrorEnum_ExperimentError)(0), // 0: google.ads.googleads.v17.errors.ExperimentErrorEnum.ExperimentError
	(*ExperimentErrorEnum)(nil),              // 1: google.ads.googleads.v17.errors.ExperimentErrorEnum
}
var file_google_ads_googleads_v17_errors_experiment_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_experiment_error_proto_init() }
func file_google_ads_googleads_v17_errors_experiment_error_proto_init() {
	if File_google_ads_googleads_v17_errors_experiment_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_experiment_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_experiment_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_experiment_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_experiment_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_experiment_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_experiment_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_experiment_error_proto = out.File
	file_google_ads_googleads_v17_errors_experiment_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_experiment_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_experiment_error_proto_depIdxs = nil
}
