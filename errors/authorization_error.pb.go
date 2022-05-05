// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: google/ads/googleads/v10/errors/authorization_error.proto

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

// Enum describing possible authorization errors.
type AuthorizationErrorEnum_AuthorizationError int32

const (
	// Enum unspecified.
	AuthorizationErrorEnum_UNSPECIFIED AuthorizationErrorEnum_AuthorizationError = 0
	// The received error code is not known in this version.
	AuthorizationErrorEnum_UNKNOWN AuthorizationErrorEnum_AuthorizationError = 1
	// User doesn't have permission to access customer. Note: If you're
	// accessing a client customer, the manager's customer ID must be set in the
	// `login-customer-id` header. Learn more at
	// https://developers.google.com/google-ads/api/docs/concepts/call-structure#cid
	AuthorizationErrorEnum_USER_PERMISSION_DENIED AuthorizationErrorEnum_AuthorizationError = 2
	// The developer token is not on the allow-list.
	AuthorizationErrorEnum_DEVELOPER_TOKEN_NOT_ON_ALLOWLIST AuthorizationErrorEnum_AuthorizationError = 13
	// The developer token is not allowed with the project sent in the request.
	AuthorizationErrorEnum_DEVELOPER_TOKEN_PROHIBITED AuthorizationErrorEnum_AuthorizationError = 4
	// The Google Cloud project sent in the request does not have permission to
	// access the api.
	AuthorizationErrorEnum_PROJECT_DISABLED AuthorizationErrorEnum_AuthorizationError = 5
	// Authorization of the client failed.
	AuthorizationErrorEnum_AUTHORIZATION_ERROR AuthorizationErrorEnum_AuthorizationError = 6
	// The user does not have permission to perform this action
	// (e.g., ADD, UPDATE, REMOVE) on the resource or call a method.
	AuthorizationErrorEnum_ACTION_NOT_PERMITTED AuthorizationErrorEnum_AuthorizationError = 7
	// Signup not complete.
	AuthorizationErrorEnum_INCOMPLETE_SIGNUP AuthorizationErrorEnum_AuthorizationError = 8
	// The customer can't be used because it isn't enabled.
	AuthorizationErrorEnum_CUSTOMER_NOT_ENABLED AuthorizationErrorEnum_AuthorizationError = 24
	// The developer must sign the terms of service. They can be found here:
	// ads.google.com/aw/apicenter
	AuthorizationErrorEnum_MISSING_TOS AuthorizationErrorEnum_AuthorizationError = 9
	// The developer token is not approved. Non-approved developer tokens can
	// only be used with test accounts.
	AuthorizationErrorEnum_DEVELOPER_TOKEN_NOT_APPROVED AuthorizationErrorEnum_AuthorizationError = 10
	// The login customer specified does not have access to the account
	// specified, so the request is invalid.
	AuthorizationErrorEnum_INVALID_LOGIN_CUSTOMER_ID_SERVING_CUSTOMER_ID_COMBINATION AuthorizationErrorEnum_AuthorizationError = 11
	// The developer specified does not have access to the service.
	AuthorizationErrorEnum_SERVICE_ACCESS_DENIED AuthorizationErrorEnum_AuthorizationError = 12
	// The customer (or login customer) isn't in Google Ads. It belongs to
	// another ads system.
	AuthorizationErrorEnum_ACCESS_DENIED_FOR_ACCOUNT_TYPE AuthorizationErrorEnum_AuthorizationError = 25
	// The developer does not have access to the metrics queried.
	AuthorizationErrorEnum_METRIC_ACCESS_DENIED AuthorizationErrorEnum_AuthorizationError = 26
)

// Enum value maps for AuthorizationErrorEnum_AuthorizationError.
var (
	AuthorizationErrorEnum_AuthorizationError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "USER_PERMISSION_DENIED",
		13: "DEVELOPER_TOKEN_NOT_ON_ALLOWLIST",
		4:  "DEVELOPER_TOKEN_PROHIBITED",
		5:  "PROJECT_DISABLED",
		6:  "AUTHORIZATION_ERROR",
		7:  "ACTION_NOT_PERMITTED",
		8:  "INCOMPLETE_SIGNUP",
		24: "CUSTOMER_NOT_ENABLED",
		9:  "MISSING_TOS",
		10: "DEVELOPER_TOKEN_NOT_APPROVED",
		11: "INVALID_LOGIN_CUSTOMER_ID_SERVING_CUSTOMER_ID_COMBINATION",
		12: "SERVICE_ACCESS_DENIED",
		25: "ACCESS_DENIED_FOR_ACCOUNT_TYPE",
		26: "METRIC_ACCESS_DENIED",
	}
	AuthorizationErrorEnum_AuthorizationError_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"USER_PERMISSION_DENIED":           2,
		"DEVELOPER_TOKEN_NOT_ON_ALLOWLIST": 13,
		"DEVELOPER_TOKEN_PROHIBITED":       4,
		"PROJECT_DISABLED":                 5,
		"AUTHORIZATION_ERROR":              6,
		"ACTION_NOT_PERMITTED":             7,
		"INCOMPLETE_SIGNUP":                8,
		"CUSTOMER_NOT_ENABLED":             24,
		"MISSING_TOS":                      9,
		"DEVELOPER_TOKEN_NOT_APPROVED":     10,
		"INVALID_LOGIN_CUSTOMER_ID_SERVING_CUSTOMER_ID_COMBINATION": 11,
		"SERVICE_ACCESS_DENIED":          12,
		"ACCESS_DENIED_FOR_ACCOUNT_TYPE": 25,
		"METRIC_ACCESS_DENIED":           26,
	}
)

func (x AuthorizationErrorEnum_AuthorizationError) Enum() *AuthorizationErrorEnum_AuthorizationError {
	p := new(AuthorizationErrorEnum_AuthorizationError)
	*p = x
	return p
}

func (x AuthorizationErrorEnum_AuthorizationError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthorizationErrorEnum_AuthorizationError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v10_errors_authorization_error_proto_enumTypes[0].Descriptor()
}

func (AuthorizationErrorEnum_AuthorizationError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v10_errors_authorization_error_proto_enumTypes[0]
}

func (x AuthorizationErrorEnum_AuthorizationError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthorizationErrorEnum_AuthorizationError.Descriptor instead.
func (AuthorizationErrorEnum_AuthorizationError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible authorization errors.
type AuthorizationErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthorizationErrorEnum) Reset() {
	*x = AuthorizationErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_errors_authorization_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationErrorEnum) ProtoMessage() {}

func (x *AuthorizationErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_errors_authorization_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationErrorEnum.ProtoReflect.Descriptor instead.
func (*AuthorizationErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v10_errors_authorization_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_errors_authorization_error_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xf4, 0x03, 0x0a,
	0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd9, 0x03, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x56, 0x45,
	0x4c, 0x4f, 0x50, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0d, 0x12, 0x1e,
	0x0a, 0x1a, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x48, 0x49, 0x42, 0x49, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x18, 0x0a,
	0x14, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x54, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x10, 0x08, 0x12, 0x18,
	0x0a, 0x14, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x53, 0x10, 0x09, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45, 0x56,
	0x45, 0x4c, 0x4f, 0x50, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x3d, 0x0a, 0x39, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x43, 0x55, 0x53,
	0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4d,
	0x42, 0x49, 0x4e, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e,
	0x49, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x19, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45,
	0x44, 0x10, 0x1a, 0x42, 0xf7, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x17, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x30, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30,
	0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x30, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescData = file_google_ads_googleads_v10_errors_authorization_error_proto_rawDesc
)

func file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_errors_authorization_error_proto_rawDescData
}

var file_google_ads_googleads_v10_errors_authorization_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v10_errors_authorization_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_errors_authorization_error_proto_goTypes = []interface{}{
	(AuthorizationErrorEnum_AuthorizationError)(0), // 0: google.ads.googleads.v10.errors.AuthorizationErrorEnum.AuthorizationError
	(*AuthorizationErrorEnum)(nil),                 // 1: google.ads.googleads.v10.errors.AuthorizationErrorEnum
}
var file_google_ads_googleads_v10_errors_authorization_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_errors_authorization_error_proto_init() }
func file_google_ads_googleads_v10_errors_authorization_error_proto_init() {
	if File_google_ads_googleads_v10_errors_authorization_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_errors_authorization_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v10_errors_authorization_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_errors_authorization_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_errors_authorization_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v10_errors_authorization_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v10_errors_authorization_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_errors_authorization_error_proto = out.File
	file_google_ads_googleads_v10_errors_authorization_error_proto_rawDesc = nil
	file_google_ads_googleads_v10_errors_authorization_error_proto_goTypes = nil
	file_google_ads_googleads_v10_errors_authorization_error_proto_depIdxs = nil
}
