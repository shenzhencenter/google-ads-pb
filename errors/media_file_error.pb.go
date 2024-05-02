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
// source: google/ads/googleads/v16/errors/media_file_error.proto

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

// Enum describing possible media file errors.
type MediaFileErrorEnum_MediaFileError int32

const (
	// Enum unspecified.
	MediaFileErrorEnum_UNSPECIFIED MediaFileErrorEnum_MediaFileError = 0
	// The received error code is not known in this version.
	MediaFileErrorEnum_UNKNOWN MediaFileErrorEnum_MediaFileError = 1
	// Cannot create a standard icon type.
	MediaFileErrorEnum_CANNOT_CREATE_STANDARD_ICON MediaFileErrorEnum_MediaFileError = 2
	// May only select Standard Icons alone.
	MediaFileErrorEnum_CANNOT_SELECT_STANDARD_ICON_WITH_OTHER_TYPES MediaFileErrorEnum_MediaFileError = 3
	// Image contains both a media file ID and data.
	MediaFileErrorEnum_CANNOT_SPECIFY_MEDIA_FILE_ID_AND_DATA MediaFileErrorEnum_MediaFileError = 4
	// A media file with given type and reference ID already exists.
	MediaFileErrorEnum_DUPLICATE_MEDIA MediaFileErrorEnum_MediaFileError = 5
	// A required field was not specified or is an empty string.
	MediaFileErrorEnum_EMPTY_FIELD MediaFileErrorEnum_MediaFileError = 6
	// A media file may only be modified once per call.
	MediaFileErrorEnum_RESOURCE_REFERENCED_IN_MULTIPLE_OPS MediaFileErrorEnum_MediaFileError = 7
	// Field is not supported for the media sub type.
	MediaFileErrorEnum_FIELD_NOT_SUPPORTED_FOR_MEDIA_SUB_TYPE MediaFileErrorEnum_MediaFileError = 8
	// The media file ID is invalid.
	MediaFileErrorEnum_INVALID_MEDIA_FILE_ID MediaFileErrorEnum_MediaFileError = 9
	// The media subtype is invalid.
	MediaFileErrorEnum_INVALID_MEDIA_SUB_TYPE MediaFileErrorEnum_MediaFileError = 10
	// The media file type is invalid.
	MediaFileErrorEnum_INVALID_MEDIA_FILE_TYPE MediaFileErrorEnum_MediaFileError = 11
	// The mimetype is invalid.
	MediaFileErrorEnum_INVALID_MIME_TYPE MediaFileErrorEnum_MediaFileError = 12
	// The media reference ID is invalid.
	MediaFileErrorEnum_INVALID_REFERENCE_ID MediaFileErrorEnum_MediaFileError = 13
	// The YouTube video ID is invalid.
	MediaFileErrorEnum_INVALID_YOU_TUBE_ID MediaFileErrorEnum_MediaFileError = 14
	// Media file has failed transcoding
	MediaFileErrorEnum_MEDIA_FILE_FAILED_TRANSCODING MediaFileErrorEnum_MediaFileError = 15
	// Media file has not been transcoded.
	MediaFileErrorEnum_MEDIA_NOT_TRANSCODED MediaFileErrorEnum_MediaFileError = 16
	// The media type does not match the actual media file's type.
	MediaFileErrorEnum_MEDIA_TYPE_DOES_NOT_MATCH_MEDIA_FILE_TYPE MediaFileErrorEnum_MediaFileError = 17
	// None of the fields have been specified.
	MediaFileErrorEnum_NO_FIELDS_SPECIFIED MediaFileErrorEnum_MediaFileError = 18
	// One of reference ID or media file ID must be specified.
	MediaFileErrorEnum_NULL_REFERENCE_ID_AND_MEDIA_ID MediaFileErrorEnum_MediaFileError = 19
	// The string has too many characters.
	MediaFileErrorEnum_TOO_LONG MediaFileErrorEnum_MediaFileError = 20
	// The specified type is not supported.
	MediaFileErrorEnum_UNSUPPORTED_TYPE MediaFileErrorEnum_MediaFileError = 21
	// YouTube is unavailable for requesting video data.
	MediaFileErrorEnum_YOU_TUBE_SERVICE_UNAVAILABLE MediaFileErrorEnum_MediaFileError = 22
	// The YouTube video has a non positive duration.
	MediaFileErrorEnum_YOU_TUBE_VIDEO_HAS_NON_POSITIVE_DURATION MediaFileErrorEnum_MediaFileError = 23
	// The YouTube video ID is syntactically valid but the video was not found.
	MediaFileErrorEnum_YOU_TUBE_VIDEO_NOT_FOUND MediaFileErrorEnum_MediaFileError = 24
)

// Enum value maps for MediaFileErrorEnum_MediaFileError.
var (
	MediaFileErrorEnum_MediaFileError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CANNOT_CREATE_STANDARD_ICON",
		3:  "CANNOT_SELECT_STANDARD_ICON_WITH_OTHER_TYPES",
		4:  "CANNOT_SPECIFY_MEDIA_FILE_ID_AND_DATA",
		5:  "DUPLICATE_MEDIA",
		6:  "EMPTY_FIELD",
		7:  "RESOURCE_REFERENCED_IN_MULTIPLE_OPS",
		8:  "FIELD_NOT_SUPPORTED_FOR_MEDIA_SUB_TYPE",
		9:  "INVALID_MEDIA_FILE_ID",
		10: "INVALID_MEDIA_SUB_TYPE",
		11: "INVALID_MEDIA_FILE_TYPE",
		12: "INVALID_MIME_TYPE",
		13: "INVALID_REFERENCE_ID",
		14: "INVALID_YOU_TUBE_ID",
		15: "MEDIA_FILE_FAILED_TRANSCODING",
		16: "MEDIA_NOT_TRANSCODED",
		17: "MEDIA_TYPE_DOES_NOT_MATCH_MEDIA_FILE_TYPE",
		18: "NO_FIELDS_SPECIFIED",
		19: "NULL_REFERENCE_ID_AND_MEDIA_ID",
		20: "TOO_LONG",
		21: "UNSUPPORTED_TYPE",
		22: "YOU_TUBE_SERVICE_UNAVAILABLE",
		23: "YOU_TUBE_VIDEO_HAS_NON_POSITIVE_DURATION",
		24: "YOU_TUBE_VIDEO_NOT_FOUND",
	}
	MediaFileErrorEnum_MediaFileError_value = map[string]int32{
		"UNSPECIFIED":                 0,
		"UNKNOWN":                     1,
		"CANNOT_CREATE_STANDARD_ICON": 2,
		"CANNOT_SELECT_STANDARD_ICON_WITH_OTHER_TYPES": 3,
		"CANNOT_SPECIFY_MEDIA_FILE_ID_AND_DATA":        4,
		"DUPLICATE_MEDIA":                              5,
		"EMPTY_FIELD":                                  6,
		"RESOURCE_REFERENCED_IN_MULTIPLE_OPS":          7,
		"FIELD_NOT_SUPPORTED_FOR_MEDIA_SUB_TYPE":       8,
		"INVALID_MEDIA_FILE_ID":                        9,
		"INVALID_MEDIA_SUB_TYPE":                       10,
		"INVALID_MEDIA_FILE_TYPE":                      11,
		"INVALID_MIME_TYPE":                            12,
		"INVALID_REFERENCE_ID":                         13,
		"INVALID_YOU_TUBE_ID":                          14,
		"MEDIA_FILE_FAILED_TRANSCODING":                15,
		"MEDIA_NOT_TRANSCODED":                         16,
		"MEDIA_TYPE_DOES_NOT_MATCH_MEDIA_FILE_TYPE":    17,
		"NO_FIELDS_SPECIFIED":                          18,
		"NULL_REFERENCE_ID_AND_MEDIA_ID":               19,
		"TOO_LONG":                                     20,
		"UNSUPPORTED_TYPE":                             21,
		"YOU_TUBE_SERVICE_UNAVAILABLE":                 22,
		"YOU_TUBE_VIDEO_HAS_NON_POSITIVE_DURATION":     23,
		"YOU_TUBE_VIDEO_NOT_FOUND":                     24,
	}
)

func (x MediaFileErrorEnum_MediaFileError) Enum() *MediaFileErrorEnum_MediaFileError {
	p := new(MediaFileErrorEnum_MediaFileError)
	*p = x
	return p
}

func (x MediaFileErrorEnum_MediaFileError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaFileErrorEnum_MediaFileError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v16_errors_media_file_error_proto_enumTypes[0].Descriptor()
}

func (MediaFileErrorEnum_MediaFileError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v16_errors_media_file_error_proto_enumTypes[0]
}

func (x MediaFileErrorEnum_MediaFileError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaFileErrorEnum_MediaFileError.Descriptor instead.
func (MediaFileErrorEnum_MediaFileError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible media file errors.
type MediaFileErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MediaFileErrorEnum) Reset() {
	*x = MediaFileErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_errors_media_file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaFileErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaFileErrorEnum) ProtoMessage() {}

func (x *MediaFileErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_errors_media_file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaFileErrorEnum.ProtoReflect.Descriptor instead.
func (*MediaFileErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v16_errors_media_file_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_errors_media_file_error_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x97, 0x06, 0x0a, 0x12, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0x80, 0x06, 0x0a, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x49, 0x43, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x30, 0x0a, 0x2c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x5f, 0x49, 0x43, 0x4f,
	0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x53, 0x10, 0x03, 0x12, 0x29, 0x0a, 0x25, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x46, 0x49, 0x4c,
	0x45, 0x5f, 0x49, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x5f, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x10, 0x06, 0x12, 0x27, 0x0a, 0x23, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f,
	0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x45, 0x5f, 0x4f, 0x50, 0x53, 0x10, 0x07, 0x12, 0x2a,
	0x0a, 0x26, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50,
	0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f,
	0x53, 0x55, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x46, 0x49, 0x4c, 0x45,
	0x5f, 0x49, 0x44, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10,
	0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x0b, 0x12, 0x15,
	0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x49, 0x4d, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x0d, 0x12,
	0x17, 0x0a, 0x13, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x59, 0x4f, 0x55, 0x5f, 0x54,
	0x55, 0x42, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x45, 0x44, 0x49,
	0x41, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x0f, 0x12, 0x18, 0x0a, 0x14, 0x4d,
	0x45, 0x44, 0x49, 0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x43, 0x4f,
	0x44, 0x45, 0x44, 0x10, 0x10, 0x12, 0x2d, 0x0a, 0x29, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x10, 0x11, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x4f, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44,
	0x53, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x12, 0x12, 0x22, 0x0a,
	0x1e, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x5f,
	0x49, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x49, 0x44, 0x10,
	0x13, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x14, 0x12,
	0x14, 0x0a, 0x10, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x15, 0x12, 0x20, 0x0a, 0x1c, 0x59, 0x4f, 0x55, 0x5f, 0x54, 0x55, 0x42,
	0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49,
	0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x16, 0x12, 0x2c, 0x0a, 0x28, 0x59, 0x4f, 0x55, 0x5f, 0x54,
	0x55, 0x42, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x4e, 0x4f,
	0x4e, 0x5f, 0x50, 0x4f, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x55, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x17, 0x12, 0x1c, 0x0a, 0x18, 0x59, 0x4f, 0x55, 0x5f, 0x54, 0x55, 0x42,
	0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x18, 0x42, 0xf3, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x13, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x36, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescData = file_google_ads_googleads_v16_errors_media_file_error_proto_rawDesc
)

func file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_errors_media_file_error_proto_rawDescData
}

var file_google_ads_googleads_v16_errors_media_file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v16_errors_media_file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v16_errors_media_file_error_proto_goTypes = []interface{}{
	(MediaFileErrorEnum_MediaFileError)(0), // 0: google.ads.googleads.v16.errors.MediaFileErrorEnum.MediaFileError
	(*MediaFileErrorEnum)(nil),             // 1: google.ads.googleads.v16.errors.MediaFileErrorEnum
}
var file_google_ads_googleads_v16_errors_media_file_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_errors_media_file_error_proto_init() }
func file_google_ads_googleads_v16_errors_media_file_error_proto_init() {
	if File_google_ads_googleads_v16_errors_media_file_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_errors_media_file_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaFileErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v16_errors_media_file_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_errors_media_file_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_errors_media_file_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v16_errors_media_file_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v16_errors_media_file_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_errors_media_file_error_proto = out.File
	file_google_ads_googleads_v16_errors_media_file_error_proto_rawDesc = nil
	file_google_ads_googleads_v16_errors_media_file_error_proto_goTypes = nil
	file_google_ads_googleads_v16_errors_media_file_error_proto_depIdxs = nil
}
