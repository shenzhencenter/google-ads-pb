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
// source: google/ads/googleads/v18/errors/customer_lifecycle_goal_error.proto

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

// Enum describing possible customer lifecycle goal errors.
type CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError int32

const (
	// Enum unspecified.
	CustomerLifecycleGoalErrorEnum_UNSPECIFIED CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 0
	// The received error code is not known in this version.
	CustomerLifecycleGoalErrorEnum_UNKNOWN CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 1
	// CustomerLifecycleGoal.customer_acquisition_goal_value_settings.value must
	// be set.
	CustomerLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_VALUE_MISSING CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 2
	// CustomerLifecycleGoal.customer_acquisition_goal_value_settings.value must
	// be no less than 0.01.
	CustomerLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_INVALID_VALUE CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 3
	// CustomerLifecycleGoal.customer_acquisition_goal_value_settings.high_lifetime_value
	// must be no less than 0.01. Also, to set this field,
	// CustomerLifecycleGoal.customer_acquisition_goal_value_settings.value must
	// also be present, and high_lifetime_value must be greater than value.
	CustomerLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_INVALID_HIGH_LIFETIME_VALUE CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 4
	// CustomerLifecycleGoal.customer_acquisition_goal_value_settings.value
	// cannot be cleared. This value would have no effect as long as none of
	// your campaigns adopt the customer acquisitiong goal.
	CustomerLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_VALUE_CANNOT_BE_CLEARED CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 5
	// CustomerLifecycleGoal.customer_acquisition_goal_value_settings.high_lifetime_value
	// cannot be cleared. This value would have no effect as long as none of
	// your campaigns adopt the high value optimization of customer acquisitiong
	// goal.
	CustomerLifecycleGoalErrorEnum_CUSTOMER_ACQUISITION_HIGH_LIFETIME_VALUE_CANNOT_BE_CLEARED CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 6
	// Found invalid value in
	// CustomerLifecycleGoal.lifecycle_goal_customer_definition_settings.existing_user_lists.
	// The userlist must be accessible, active and belong to one of the
	// following types: CRM_BASED, RULE_BASED, REMARKETING.
	CustomerLifecycleGoalErrorEnum_INVALID_EXISTING_USER_LIST CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 7
	// Found invalid value in
	// CustomerLifecycleGoal.lifecycle_goal_customer_definition_settings.high_lifetime_value_user_lists.
	// The userlist must be accessible, active and belong to one of the
	// following types: CRM_BASED, RULE_BASED, REMARKETING.
	CustomerLifecycleGoalErrorEnum_INVALID_HIGH_LIFETIME_VALUE_USER_LIST CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError = 8
)

// Enum value maps for CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError.
var (
	CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CUSTOMER_ACQUISITION_VALUE_MISSING",
		3: "CUSTOMER_ACQUISITION_INVALID_VALUE",
		4: "CUSTOMER_ACQUISITION_INVALID_HIGH_LIFETIME_VALUE",
		5: "CUSTOMER_ACQUISITION_VALUE_CANNOT_BE_CLEARED",
		6: "CUSTOMER_ACQUISITION_HIGH_LIFETIME_VALUE_CANNOT_BE_CLEARED",
		7: "INVALID_EXISTING_USER_LIST",
		8: "INVALID_HIGH_LIFETIME_VALUE_USER_LIST",
	}
	CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"CUSTOMER_ACQUISITION_VALUE_MISSING": 2,
		"CUSTOMER_ACQUISITION_INVALID_VALUE": 3,
		"CUSTOMER_ACQUISITION_INVALID_HIGH_LIFETIME_VALUE":           4,
		"CUSTOMER_ACQUISITION_VALUE_CANNOT_BE_CLEARED":               5,
		"CUSTOMER_ACQUISITION_HIGH_LIFETIME_VALUE_CANNOT_BE_CLEARED": 6,
		"INVALID_EXISTING_USER_LIST":                                 7,
		"INVALID_HIGH_LIFETIME_VALUE_USER_LIST":                      8,
	}
)

func (x CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError) Enum() *CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError {
	p := new(CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError)
	*p = x
	return p
}

func (x CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_enumTypes[0].Descriptor()
}

func (CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_enumTypes[0]
}

func (x CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError.Descriptor instead.
func (CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible customer lifecycle goal errors.
type CustomerLifecycleGoalErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomerLifecycleGoalErrorEnum) Reset() {
	*x = CustomerLifecycleGoalErrorEnum{}
	mi := &file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerLifecycleGoalErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerLifecycleGoalErrorEnum) ProtoMessage() {}

func (x *CustomerLifecycleGoalErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerLifecycleGoalErrorEnum.ProtoReflect.Descriptor instead.
func (*CustomerLifecycleGoalErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xa0, 0x03, 0x0a, 0x1e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xfd, 0x02, 0x0a, 0x1a, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47,
	0x6f, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x26,
	0x0a, 0x22, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x10, 0x03, 0x12, 0x34, 0x0a, 0x30, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x4c, 0x49, 0x46, 0x45,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x04, 0x12, 0x30, 0x0a, 0x2c,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49, 0x53, 0x49,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x42, 0x45, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x45, 0x44, 0x10, 0x05, 0x12, 0x3e,
	0x0a, 0x3a, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x51, 0x55, 0x49,
	0x53, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x4c, 0x49, 0x46, 0x45,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x42, 0x45, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x45, 0x44, 0x10, 0x06, 0x12, 0x1e,
	0x0a, 0x1a, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x07, 0x12, 0x29,
	0x0a, 0x25, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x5f, 0x4c,
	0x49, 0x46, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x08, 0x42, 0xff, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x42, 0x1f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x38, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescData = file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDesc
)

func file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDescData
}

var file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_goTypes = []any{
	(CustomerLifecycleGoalErrorEnum_CustomerLifecycleGoalError)(0), // 0: google.ads.googleads.v18.errors.CustomerLifecycleGoalErrorEnum.CustomerLifecycleGoalError
	(*CustomerLifecycleGoalErrorEnum)(nil),                         // 1: google.ads.googleads.v18.errors.CustomerLifecycleGoalErrorEnum
}
var file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_init() }
func file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_init() {
	if File_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto = out.File
	file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_rawDesc = nil
	file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_goTypes = nil
	file_google_ads_googleads_v18_errors_customer_lifecycle_goal_error_proto_depIdxs = nil
}
