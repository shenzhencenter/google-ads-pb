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
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: google/ads/googleads/v19/errors/image_error.proto

package errors

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum describing possible image errors.
type ImageErrorEnum_ImageError int32

const (
	// Enum unspecified.
	ImageErrorEnum_UNSPECIFIED ImageErrorEnum_ImageError = 0
	// The received error code is not known in this version.
	ImageErrorEnum_UNKNOWN ImageErrorEnum_ImageError = 1
	// The image is not valid.
	ImageErrorEnum_INVALID_IMAGE ImageErrorEnum_ImageError = 2
	// The image could not be stored.
	ImageErrorEnum_STORAGE_ERROR ImageErrorEnum_ImageError = 3
	// There was a problem with the request.
	ImageErrorEnum_BAD_REQUEST ImageErrorEnum_ImageError = 4
	// The image is not of legal dimensions.
	ImageErrorEnum_UNEXPECTED_SIZE ImageErrorEnum_ImageError = 5
	// Animated image are not permitted.
	ImageErrorEnum_ANIMATED_NOT_ALLOWED ImageErrorEnum_ImageError = 6
	// Animation is too long.
	ImageErrorEnum_ANIMATION_TOO_LONG ImageErrorEnum_ImageError = 7
	// There was an error on the server.
	ImageErrorEnum_SERVER_ERROR ImageErrorEnum_ImageError = 8
	// Image cannot be in CMYK color format.
	ImageErrorEnum_CMYK_JPEG_NOT_ALLOWED ImageErrorEnum_ImageError = 9
	// Flash images are not permitted.
	ImageErrorEnum_FLASH_NOT_ALLOWED ImageErrorEnum_ImageError = 10
	// Flash images must support clickTag.
	ImageErrorEnum_FLASH_WITHOUT_CLICKTAG ImageErrorEnum_ImageError = 11
	// A flash error has occurred after fixing the click tag.
	ImageErrorEnum_FLASH_ERROR_AFTER_FIXING_CLICK_TAG ImageErrorEnum_ImageError = 12
	// Unacceptable visual effects.
	ImageErrorEnum_ANIMATED_VISUAL_EFFECT ImageErrorEnum_ImageError = 13
	// There was a problem with the flash image.
	ImageErrorEnum_FLASH_ERROR ImageErrorEnum_ImageError = 14
	// Incorrect image layout.
	ImageErrorEnum_LAYOUT_PROBLEM ImageErrorEnum_ImageError = 15
	// There was a problem reading the image file.
	ImageErrorEnum_PROBLEM_READING_IMAGE_FILE ImageErrorEnum_ImageError = 16
	// There was an error storing the image.
	ImageErrorEnum_ERROR_STORING_IMAGE ImageErrorEnum_ImageError = 17
	// The aspect ratio of the image is not allowed.
	ImageErrorEnum_ASPECT_RATIO_NOT_ALLOWED ImageErrorEnum_ImageError = 18
	// Flash cannot have network objects.
	ImageErrorEnum_FLASH_HAS_NETWORK_OBJECTS ImageErrorEnum_ImageError = 19
	// Flash cannot have network methods.
	ImageErrorEnum_FLASH_HAS_NETWORK_METHODS ImageErrorEnum_ImageError = 20
	// Flash cannot have a Url.
	ImageErrorEnum_FLASH_HAS_URL ImageErrorEnum_ImageError = 21
	// Flash cannot use mouse tracking.
	ImageErrorEnum_FLASH_HAS_MOUSE_TRACKING ImageErrorEnum_ImageError = 22
	// Flash cannot have a random number.
	ImageErrorEnum_FLASH_HAS_RANDOM_NUM ImageErrorEnum_ImageError = 23
	// Ad click target cannot be '_self'.
	ImageErrorEnum_FLASH_SELF_TARGETS ImageErrorEnum_ImageError = 24
	// GetUrl method should only use '_blank'.
	ImageErrorEnum_FLASH_BAD_GETURL_TARGET ImageErrorEnum_ImageError = 25
	// Flash version is not supported.
	ImageErrorEnum_FLASH_VERSION_NOT_SUPPORTED ImageErrorEnum_ImageError = 26
	// Flash movies need to have hard coded click URL or clickTAG
	ImageErrorEnum_FLASH_WITHOUT_HARD_CODED_CLICK_URL ImageErrorEnum_ImageError = 27
	// Uploaded flash file is corrupted.
	ImageErrorEnum_INVALID_FLASH_FILE ImageErrorEnum_ImageError = 28
	// Uploaded flash file can be parsed, but the click tag can not be fixed
	// properly.
	ImageErrorEnum_FAILED_TO_FIX_CLICK_TAG_IN_FLASH ImageErrorEnum_ImageError = 29
	// Flash movie accesses network resources
	ImageErrorEnum_FLASH_ACCESSES_NETWORK_RESOURCES ImageErrorEnum_ImageError = 30
	// Flash movie attempts to call external javascript code
	ImageErrorEnum_FLASH_EXTERNAL_JS_CALL ImageErrorEnum_ImageError = 31
	// Flash movie attempts to call flash system commands
	ImageErrorEnum_FLASH_EXTERNAL_FS_CALL ImageErrorEnum_ImageError = 32
	// Image file is too large.
	ImageErrorEnum_FILE_TOO_LARGE ImageErrorEnum_ImageError = 33
	// Image data is too large.
	ImageErrorEnum_IMAGE_DATA_TOO_LARGE ImageErrorEnum_ImageError = 34
	// Error while processing the image.
	ImageErrorEnum_IMAGE_PROCESSING_ERROR ImageErrorEnum_ImageError = 35
	// Image is too small.
	ImageErrorEnum_IMAGE_TOO_SMALL ImageErrorEnum_ImageError = 36
	// Input was invalid.
	ImageErrorEnum_INVALID_INPUT ImageErrorEnum_ImageError = 37
	// There was a problem reading the image file.
	ImageErrorEnum_PROBLEM_READING_FILE ImageErrorEnum_ImageError = 38
	// Image constraints are violated, but details like ASPECT_RATIO_NOT_ALLOWED
	// can't be provided. This happens when asset spec contains more than one
	// constraint and different criteria of different constraints are violated.
	ImageErrorEnum_IMAGE_CONSTRAINTS_VIOLATED ImageErrorEnum_ImageError = 39
	// Image format is not allowed.
	ImageErrorEnum_FORMAT_NOT_ALLOWED ImageErrorEnum_ImageError = 40
)

// Enum value maps for ImageErrorEnum_ImageError.
var (
	ImageErrorEnum_ImageError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "INVALID_IMAGE",
		3:  "STORAGE_ERROR",
		4:  "BAD_REQUEST",
		5:  "UNEXPECTED_SIZE",
		6:  "ANIMATED_NOT_ALLOWED",
		7:  "ANIMATION_TOO_LONG",
		8:  "SERVER_ERROR",
		9:  "CMYK_JPEG_NOT_ALLOWED",
		10: "FLASH_NOT_ALLOWED",
		11: "FLASH_WITHOUT_CLICKTAG",
		12: "FLASH_ERROR_AFTER_FIXING_CLICK_TAG",
		13: "ANIMATED_VISUAL_EFFECT",
		14: "FLASH_ERROR",
		15: "LAYOUT_PROBLEM",
		16: "PROBLEM_READING_IMAGE_FILE",
		17: "ERROR_STORING_IMAGE",
		18: "ASPECT_RATIO_NOT_ALLOWED",
		19: "FLASH_HAS_NETWORK_OBJECTS",
		20: "FLASH_HAS_NETWORK_METHODS",
		21: "FLASH_HAS_URL",
		22: "FLASH_HAS_MOUSE_TRACKING",
		23: "FLASH_HAS_RANDOM_NUM",
		24: "FLASH_SELF_TARGETS",
		25: "FLASH_BAD_GETURL_TARGET",
		26: "FLASH_VERSION_NOT_SUPPORTED",
		27: "FLASH_WITHOUT_HARD_CODED_CLICK_URL",
		28: "INVALID_FLASH_FILE",
		29: "FAILED_TO_FIX_CLICK_TAG_IN_FLASH",
		30: "FLASH_ACCESSES_NETWORK_RESOURCES",
		31: "FLASH_EXTERNAL_JS_CALL",
		32: "FLASH_EXTERNAL_FS_CALL",
		33: "FILE_TOO_LARGE",
		34: "IMAGE_DATA_TOO_LARGE",
		35: "IMAGE_PROCESSING_ERROR",
		36: "IMAGE_TOO_SMALL",
		37: "INVALID_INPUT",
		38: "PROBLEM_READING_FILE",
		39: "IMAGE_CONSTRAINTS_VIOLATED",
		40: "FORMAT_NOT_ALLOWED",
	}
	ImageErrorEnum_ImageError_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"INVALID_IMAGE":                      2,
		"STORAGE_ERROR":                      3,
		"BAD_REQUEST":                        4,
		"UNEXPECTED_SIZE":                    5,
		"ANIMATED_NOT_ALLOWED":               6,
		"ANIMATION_TOO_LONG":                 7,
		"SERVER_ERROR":                       8,
		"CMYK_JPEG_NOT_ALLOWED":              9,
		"FLASH_NOT_ALLOWED":                  10,
		"FLASH_WITHOUT_CLICKTAG":             11,
		"FLASH_ERROR_AFTER_FIXING_CLICK_TAG": 12,
		"ANIMATED_VISUAL_EFFECT":             13,
		"FLASH_ERROR":                        14,
		"LAYOUT_PROBLEM":                     15,
		"PROBLEM_READING_IMAGE_FILE":         16,
		"ERROR_STORING_IMAGE":                17,
		"ASPECT_RATIO_NOT_ALLOWED":           18,
		"FLASH_HAS_NETWORK_OBJECTS":          19,
		"FLASH_HAS_NETWORK_METHODS":          20,
		"FLASH_HAS_URL":                      21,
		"FLASH_HAS_MOUSE_TRACKING":           22,
		"FLASH_HAS_RANDOM_NUM":               23,
		"FLASH_SELF_TARGETS":                 24,
		"FLASH_BAD_GETURL_TARGET":            25,
		"FLASH_VERSION_NOT_SUPPORTED":        26,
		"FLASH_WITHOUT_HARD_CODED_CLICK_URL": 27,
		"INVALID_FLASH_FILE":                 28,
		"FAILED_TO_FIX_CLICK_TAG_IN_FLASH":   29,
		"FLASH_ACCESSES_NETWORK_RESOURCES":   30,
		"FLASH_EXTERNAL_JS_CALL":             31,
		"FLASH_EXTERNAL_FS_CALL":             32,
		"FILE_TOO_LARGE":                     33,
		"IMAGE_DATA_TOO_LARGE":               34,
		"IMAGE_PROCESSING_ERROR":             35,
		"IMAGE_TOO_SMALL":                    36,
		"INVALID_INPUT":                      37,
		"PROBLEM_READING_FILE":               38,
		"IMAGE_CONSTRAINTS_VIOLATED":         39,
		"FORMAT_NOT_ALLOWED":                 40,
	}
)

func (x ImageErrorEnum_ImageError) Enum() *ImageErrorEnum_ImageError {
	p := new(ImageErrorEnum_ImageError)
	*p = x
	return p
}

func (x ImageErrorEnum_ImageError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageErrorEnum_ImageError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v19_errors_image_error_proto_enumTypes[0].Descriptor()
}

func (ImageErrorEnum_ImageError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v19_errors_image_error_proto_enumTypes[0]
}

func (x ImageErrorEnum_ImageError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageErrorEnum_ImageError.Descriptor instead.
func (ImageErrorEnum_ImageError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_errors_image_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible image errors.
type ImageErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageErrorEnum) Reset() {
	*x = ImageErrorEnum{}
	mi := &file_google_ads_googleads_v19_errors_image_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageErrorEnum) ProtoMessage() {}

func (x *ImageErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_errors_image_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageErrorEnum.ProtoReflect.Descriptor instead.
func (*ImageErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_errors_image_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v19_errors_image_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_errors_image_error_proto_rawDesc = string([]byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x22, 0xc2, 0x08, 0x0a, 0x0e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xaf, 0x08, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x4f, 0x52, 0x41,
	0x47, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41,
	0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x55,
	0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x10, 0x05,
	0x12, 0x18, 0x0a, 0x14, 0x41, 0x4e, 0x49, 0x4d, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x06, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x4e,
	0x49, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x4e, 0x47,
	0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4d, 0x59, 0x4b, 0x5f, 0x4a, 0x50, 0x45,
	0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x15, 0x0a, 0x11, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c,
	0x4f, 0x57, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f,
	0x57, 0x49, 0x54, 0x48, 0x4f, 0x55, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x54, 0x41, 0x47,
	0x10, 0x0b, 0x12, 0x26, 0x0a, 0x22, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x58, 0x49, 0x4e, 0x47, 0x5f, 0x43,
	0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x0c, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x4e,
	0x49, 0x4d, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x56, 0x49, 0x53, 0x55, 0x41, 0x4c, 0x5f, 0x45, 0x46,
	0x46, 0x45, 0x43, 0x54, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x41, 0x59, 0x4f, 0x55,
	0x54, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x10, 0x0f, 0x12, 0x1e, 0x0a, 0x1a, 0x50,
	0x52, 0x4f, 0x42, 0x4c, 0x45, 0x4d, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x49,
	0x4d, 0x41, 0x47, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x11, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x53, 0x50, 0x45, 0x43, 0x54, 0x5f, 0x52,
	0x41, 0x54, 0x49, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x12, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x48, 0x41, 0x53, 0x5f,
	0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x53, 0x10,
	0x13, 0x12, 0x1d, 0x0a, 0x19, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x53, 0x10, 0x14,
	0x12, 0x11, 0x0a, 0x0d, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x55, 0x52,
	0x4c, 0x10, 0x15, 0x12, 0x1c, 0x0a, 0x18, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x48, 0x41, 0x53,
	0x5f, 0x4d, 0x4f, 0x55, 0x53, 0x45, 0x5f, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10,
	0x16, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x52,
	0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x5f, 0x4e, 0x55, 0x4d, 0x10, 0x17, 0x12, 0x16, 0x0a, 0x12, 0x46,
	0x4c, 0x41, 0x53, 0x48, 0x5f, 0x53, 0x45, 0x4c, 0x46, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54,
	0x53, 0x10, 0x18, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x42, 0x41, 0x44,
	0x5f, 0x47, 0x45, 0x54, 0x55, 0x52, 0x4c, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x19,
	0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x1a, 0x12, 0x26, 0x0a, 0x22, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x4f,
	0x55, 0x54, 0x5f, 0x48, 0x41, 0x52, 0x44, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x44, 0x5f, 0x43, 0x4c,
	0x49, 0x43, 0x4b, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x1b, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10,
	0x1c, 0x12, 0x24, 0x0a, 0x20, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x46,
	0x49, 0x58, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x5f,
	0x46, 0x4c, 0x41, 0x53, 0x48, 0x10, 0x1d, 0x12, 0x24, 0x0a, 0x20, 0x46, 0x4c, 0x41, 0x53, 0x48,
	0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x45, 0x53, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x53, 0x10, 0x1e, 0x12, 0x1a, 0x0a,
	0x16, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x4a, 0x53, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x1f, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x4c, 0x41,
	0x53, 0x48, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x46, 0x53, 0x5f, 0x43,
	0x41, 0x4c, 0x4c, 0x10, 0x20, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x4f,
	0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x21, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c, 0x41, 0x52, 0x47,
	0x45, 0x10, 0x22, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x4f,
	0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x23, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x53, 0x4d, 0x41,
	0x4c, 0x4c, 0x10, 0x24, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x49, 0x4e, 0x50, 0x55, 0x54, 0x10, 0x25, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x42, 0x4c,
	0x45, 0x4d, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10,
	0x26, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x54, 0x53, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x27, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x28, 0x42, 0xef, 0x01, 0x0a, 0x23, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x42, 0x0f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x39, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_google_ads_googleads_v19_errors_image_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_errors_image_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_errors_image_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_errors_image_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_errors_image_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_errors_image_error_proto_rawDesc), len(file_google_ads_googleads_v19_errors_image_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_errors_image_error_proto_rawDescData
}

var file_google_ads_googleads_v19_errors_image_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v19_errors_image_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v19_errors_image_error_proto_goTypes = []any{
	(ImageErrorEnum_ImageError)(0), // 0: google.ads.googleads.v19.errors.ImageErrorEnum.ImageError
	(*ImageErrorEnum)(nil),         // 1: google.ads.googleads.v19.errors.ImageErrorEnum
}
var file_google_ads_googleads_v19_errors_image_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_errors_image_error_proto_init() }
func file_google_ads_googleads_v19_errors_image_error_proto_init() {
	if File_google_ads_googleads_v19_errors_image_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_errors_image_error_proto_rawDesc), len(file_google_ads_googleads_v19_errors_image_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_errors_image_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_errors_image_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v19_errors_image_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v19_errors_image_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_errors_image_error_proto = out.File
	file_google_ads_googleads_v19_errors_image_error_proto_goTypes = nil
	file_google_ads_googleads_v19_errors_image_error_proto_depIdxs = nil
}
