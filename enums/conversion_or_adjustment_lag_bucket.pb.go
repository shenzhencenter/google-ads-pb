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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: google/ads/googleads/v12/enums/conversion_or_adjustment_lag_bucket.proto

package enums

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

// Enum representing the number of days between the impression and the
// conversion or between the impression and adjustments to the conversion.
type ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket int32

const (
	// Not specified.
	ConversionOrAdjustmentLagBucketEnum_UNSPECIFIED ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionOrAdjustmentLagBucketEnum_UNKNOWN ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 1
	// Conversion lag bucket from 0 to 1 day. 0 day is included, 1 day is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_LESS_THAN_ONE_DAY ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 2
	// Conversion lag bucket from 1 to 2 days. 1 day is included, 2 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_ONE_TO_TWO_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 3
	// Conversion lag bucket from 2 to 3 days. 2 days is included,
	// 3 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TWO_TO_THREE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 4
	// Conversion lag bucket from 3 to 4 days. 3 days is included,
	// 4 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_THREE_TO_FOUR_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 5
	// Conversion lag bucket from 4 to 5 days. 4 days is included,
	// 5 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FOUR_TO_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 6
	// Conversion lag bucket from 5 to 6 days. 5 days is included,
	// 6 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FIVE_TO_SIX_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 7
	// Conversion lag bucket from 6 to 7 days. 6 days is included,
	// 7 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_SIX_TO_SEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 8
	// Conversion lag bucket from 7 to 8 days. 7 days is included,
	// 8 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_SEVEN_TO_EIGHT_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 9
	// Conversion lag bucket from 8 to 9 days. 8 days is included,
	// 9 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_EIGHT_TO_NINE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 10
	// Conversion lag bucket from 9 to 10 days. 9 days is included,
	// 10 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_NINE_TO_TEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 11
	// Conversion lag bucket from 10 to 11 days. 10 days is included,
	// 11 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TEN_TO_ELEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 12
	// Conversion lag bucket from 11 to 12 days. 11 days is included,
	// 12 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_ELEVEN_TO_TWELVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 13
	// Conversion lag bucket from 12 to 13 days. 12 days is included,
	// 13 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TWELVE_TO_THIRTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 14
	// Conversion lag bucket from 13 to 14 days. 13 days is included,
	// 14 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_THIRTEEN_TO_FOURTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 15
	// Conversion lag bucket from 14 to 21 days. 14 days is included,
	// 21 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FOURTEEN_TO_TWENTY_ONE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 16
	// Conversion lag bucket from 21 to 30 days. 21 days is included,
	// 30 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_TWENTY_ONE_TO_THIRTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 17
	// Conversion lag bucket from 30 to 45 days. 30 days is included,
	// 45 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_THIRTY_TO_FORTY_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 18
	// Conversion lag bucket from 45 to 60 days. 45 days is included,
	// 60 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_FORTY_FIVE_TO_SIXTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 19
	// Conversion lag bucket from 60 to 90 days. 60 days is included,
	// 90 days is not.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_SIXTY_TO_NINETY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 20
	// Conversion adjustment lag bucket from 0 to 1 day. 0 day is included,
	// 1 day is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_LESS_THAN_ONE_DAY ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 21
	// Conversion adjustment lag bucket from 1 to 2 days. 1 day is included,
	// 2 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_ONE_TO_TWO_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 22
	// Conversion adjustment lag bucket from 2 to 3 days. 2 days is included,
	// 3 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TWO_TO_THREE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 23
	// Conversion adjustment lag bucket from 3 to 4 days. 3 days is included,
	// 4 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_THREE_TO_FOUR_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 24
	// Conversion adjustment lag bucket from 4 to 5 days. 4 days is included,
	// 5 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FOUR_TO_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 25
	// Conversion adjustment lag bucket from 5 to 6 days. 5 days is included,
	// 6 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FIVE_TO_SIX_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 26
	// Conversion adjustment lag bucket from 6 to 7 days. 6 days is included,
	// 7 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_SIX_TO_SEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 27
	// Conversion adjustment lag bucket from 7 to 8 days. 7 days is included,
	// 8 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_SEVEN_TO_EIGHT_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 28
	// Conversion adjustment lag bucket from 8 to 9 days. 8 days is included,
	// 9 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_EIGHT_TO_NINE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 29
	// Conversion adjustment lag bucket from 9 to 10 days. 9 days is included,
	// 10 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_NINE_TO_TEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 30
	// Conversion adjustment lag bucket from 10 to 11 days. 10 days is included,
	// 11 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TEN_TO_ELEVEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 31
	// Conversion adjustment lag bucket from 11 to 12 days. 11 days is included,
	// 12 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_ELEVEN_TO_TWELVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 32
	// Conversion adjustment lag bucket from 12 to 13 days. 12 days is included,
	// 13 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TWELVE_TO_THIRTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 33
	// Conversion adjustment lag bucket from 13 to 14 days. 13 days is included,
	// 14 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_THIRTEEN_TO_FOURTEEN_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 34
	// Conversion adjustment lag bucket from 14 to 21 days. 14 days is included,
	// 21 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FOURTEEN_TO_TWENTY_ONE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 35
	// Conversion adjustment lag bucket from 21 to 30 days. 21 days is included,
	// 30 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_TWENTY_ONE_TO_THIRTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 36
	// Conversion adjustment lag bucket from 30 to 45 days. 30 days is included,
	// 45 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_THIRTY_TO_FORTY_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 37
	// Conversion adjustment lag bucket from 45 to 60 days. 45 days is included,
	// 60 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_FORTY_FIVE_TO_SIXTY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 38
	// Conversion adjustment lag bucket from 60 to 90 days. 60 days is included,
	// 90 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_SIXTY_TO_NINETY_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 39
	// Conversion adjustment lag bucket from 90 to 145 days. 90 days is
	// included, 145 days is not.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_NINETY_TO_ONE_HUNDRED_AND_FORTY_FIVE_DAYS ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 40
	// Conversion lag bucket UNKNOWN. This is for dates before conversion lag
	// bucket was available in Google Ads.
	ConversionOrAdjustmentLagBucketEnum_CONVERSION_UNKNOWN ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 41
	// Conversion adjustment lag bucket UNKNOWN. This is for dates before
	// conversion adjustment lag bucket was available in Google Ads.
	ConversionOrAdjustmentLagBucketEnum_ADJUSTMENT_UNKNOWN ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket = 42
)

// Enum value maps for ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket.
var (
	ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CONVERSION_LESS_THAN_ONE_DAY",
		3:  "CONVERSION_ONE_TO_TWO_DAYS",
		4:  "CONVERSION_TWO_TO_THREE_DAYS",
		5:  "CONVERSION_THREE_TO_FOUR_DAYS",
		6:  "CONVERSION_FOUR_TO_FIVE_DAYS",
		7:  "CONVERSION_FIVE_TO_SIX_DAYS",
		8:  "CONVERSION_SIX_TO_SEVEN_DAYS",
		9:  "CONVERSION_SEVEN_TO_EIGHT_DAYS",
		10: "CONVERSION_EIGHT_TO_NINE_DAYS",
		11: "CONVERSION_NINE_TO_TEN_DAYS",
		12: "CONVERSION_TEN_TO_ELEVEN_DAYS",
		13: "CONVERSION_ELEVEN_TO_TWELVE_DAYS",
		14: "CONVERSION_TWELVE_TO_THIRTEEN_DAYS",
		15: "CONVERSION_THIRTEEN_TO_FOURTEEN_DAYS",
		16: "CONVERSION_FOURTEEN_TO_TWENTY_ONE_DAYS",
		17: "CONVERSION_TWENTY_ONE_TO_THIRTY_DAYS",
		18: "CONVERSION_THIRTY_TO_FORTY_FIVE_DAYS",
		19: "CONVERSION_FORTY_FIVE_TO_SIXTY_DAYS",
		20: "CONVERSION_SIXTY_TO_NINETY_DAYS",
		21: "ADJUSTMENT_LESS_THAN_ONE_DAY",
		22: "ADJUSTMENT_ONE_TO_TWO_DAYS",
		23: "ADJUSTMENT_TWO_TO_THREE_DAYS",
		24: "ADJUSTMENT_THREE_TO_FOUR_DAYS",
		25: "ADJUSTMENT_FOUR_TO_FIVE_DAYS",
		26: "ADJUSTMENT_FIVE_TO_SIX_DAYS",
		27: "ADJUSTMENT_SIX_TO_SEVEN_DAYS",
		28: "ADJUSTMENT_SEVEN_TO_EIGHT_DAYS",
		29: "ADJUSTMENT_EIGHT_TO_NINE_DAYS",
		30: "ADJUSTMENT_NINE_TO_TEN_DAYS",
		31: "ADJUSTMENT_TEN_TO_ELEVEN_DAYS",
		32: "ADJUSTMENT_ELEVEN_TO_TWELVE_DAYS",
		33: "ADJUSTMENT_TWELVE_TO_THIRTEEN_DAYS",
		34: "ADJUSTMENT_THIRTEEN_TO_FOURTEEN_DAYS",
		35: "ADJUSTMENT_FOURTEEN_TO_TWENTY_ONE_DAYS",
		36: "ADJUSTMENT_TWENTY_ONE_TO_THIRTY_DAYS",
		37: "ADJUSTMENT_THIRTY_TO_FORTY_FIVE_DAYS",
		38: "ADJUSTMENT_FORTY_FIVE_TO_SIXTY_DAYS",
		39: "ADJUSTMENT_SIXTY_TO_NINETY_DAYS",
		40: "ADJUSTMENT_NINETY_TO_ONE_HUNDRED_AND_FORTY_FIVE_DAYS",
		41: "CONVERSION_UNKNOWN",
		42: "ADJUSTMENT_UNKNOWN",
	}
	ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket_value = map[string]int32{
		"UNSPECIFIED":                                          0,
		"UNKNOWN":                                              1,
		"CONVERSION_LESS_THAN_ONE_DAY":                         2,
		"CONVERSION_ONE_TO_TWO_DAYS":                           3,
		"CONVERSION_TWO_TO_THREE_DAYS":                         4,
		"CONVERSION_THREE_TO_FOUR_DAYS":                        5,
		"CONVERSION_FOUR_TO_FIVE_DAYS":                         6,
		"CONVERSION_FIVE_TO_SIX_DAYS":                          7,
		"CONVERSION_SIX_TO_SEVEN_DAYS":                         8,
		"CONVERSION_SEVEN_TO_EIGHT_DAYS":                       9,
		"CONVERSION_EIGHT_TO_NINE_DAYS":                        10,
		"CONVERSION_NINE_TO_TEN_DAYS":                          11,
		"CONVERSION_TEN_TO_ELEVEN_DAYS":                        12,
		"CONVERSION_ELEVEN_TO_TWELVE_DAYS":                     13,
		"CONVERSION_TWELVE_TO_THIRTEEN_DAYS":                   14,
		"CONVERSION_THIRTEEN_TO_FOURTEEN_DAYS":                 15,
		"CONVERSION_FOURTEEN_TO_TWENTY_ONE_DAYS":               16,
		"CONVERSION_TWENTY_ONE_TO_THIRTY_DAYS":                 17,
		"CONVERSION_THIRTY_TO_FORTY_FIVE_DAYS":                 18,
		"CONVERSION_FORTY_FIVE_TO_SIXTY_DAYS":                  19,
		"CONVERSION_SIXTY_TO_NINETY_DAYS":                      20,
		"ADJUSTMENT_LESS_THAN_ONE_DAY":                         21,
		"ADJUSTMENT_ONE_TO_TWO_DAYS":                           22,
		"ADJUSTMENT_TWO_TO_THREE_DAYS":                         23,
		"ADJUSTMENT_THREE_TO_FOUR_DAYS":                        24,
		"ADJUSTMENT_FOUR_TO_FIVE_DAYS":                         25,
		"ADJUSTMENT_FIVE_TO_SIX_DAYS":                          26,
		"ADJUSTMENT_SIX_TO_SEVEN_DAYS":                         27,
		"ADJUSTMENT_SEVEN_TO_EIGHT_DAYS":                       28,
		"ADJUSTMENT_EIGHT_TO_NINE_DAYS":                        29,
		"ADJUSTMENT_NINE_TO_TEN_DAYS":                          30,
		"ADJUSTMENT_TEN_TO_ELEVEN_DAYS":                        31,
		"ADJUSTMENT_ELEVEN_TO_TWELVE_DAYS":                     32,
		"ADJUSTMENT_TWELVE_TO_THIRTEEN_DAYS":                   33,
		"ADJUSTMENT_THIRTEEN_TO_FOURTEEN_DAYS":                 34,
		"ADJUSTMENT_FOURTEEN_TO_TWENTY_ONE_DAYS":               35,
		"ADJUSTMENT_TWENTY_ONE_TO_THIRTY_DAYS":                 36,
		"ADJUSTMENT_THIRTY_TO_FORTY_FIVE_DAYS":                 37,
		"ADJUSTMENT_FORTY_FIVE_TO_SIXTY_DAYS":                  38,
		"ADJUSTMENT_SIXTY_TO_NINETY_DAYS":                      39,
		"ADJUSTMENT_NINETY_TO_ONE_HUNDRED_AND_FORTY_FIVE_DAYS": 40,
		"CONVERSION_UNKNOWN":                                   41,
		"ADJUSTMENT_UNKNOWN":                                   42,
	}
)

func (x ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) Enum() *ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket {
	p := new(ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket)
	*p = x
	return p
}

func (x ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_enumTypes[0].Descriptor()
}

func (ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_enumTypes[0]
}

func (x ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket.Descriptor instead.
func (ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum representing the number of days between the impression and
// the conversion or between the impression and adjustments to the conversion.
type ConversionOrAdjustmentLagBucketEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConversionOrAdjustmentLagBucketEnum) Reset() {
	*x = ConversionOrAdjustmentLagBucketEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionOrAdjustmentLagBucketEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionOrAdjustmentLagBucketEnum) ProtoMessage() {}

func (x *ConversionOrAdjustmentLagBucketEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionOrAdjustmentLagBucketEnum.ProtoReflect.Descriptor instead.
func (*ConversionOrAdjustmentLagBucketEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDesc = []byte{
	0x0a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x5f, 0x61,
	0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x67, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0xcb, 0x0c, 0x0a, 0x23, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x41, 0x64, 0x6a, 0x75, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xa3, 0x0c, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x4f, 0x72, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x67,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x45,
	0x5f, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57, 0x4f, 0x5f,
	0x44, 0x41, 0x59, 0x53, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x52, 0x45,
	0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x54, 0x4f, 0x5f,
	0x46, 0x4f, 0x55, 0x52, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x5f, 0x54,
	0x4f, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x06, 0x12, 0x1f, 0x0a,
	0x1b, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x56, 0x45,
	0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x49, 0x58, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x07, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x58,
	0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x08,
	0x12, 0x22, 0x0a, 0x1e, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x45, 0x56, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x09, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x49, 0x4e, 0x45,
	0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0a, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x56, 0x45,
	0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x45,
	0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0b, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x4c,
	0x45, 0x56, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0c, 0x12, 0x24, 0x0a, 0x20, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x4c, 0x45, 0x56, 0x45, 0x4e,
	0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57, 0x45, 0x4c, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10,
	0x0d, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x57, 0x45, 0x4c, 0x56, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x45,
	0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x0e, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x45, 0x45, 0x4e,
	0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59,
	0x53, 0x10, 0x0f, 0x12, 0x2a, 0x0a, 0x26, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57,
	0x45, 0x4e, 0x54, 0x59, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x10, 0x12,
	0x28, 0x0a, 0x24, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x57,
	0x45, 0x4e, 0x54, 0x59, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52,
	0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x11, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x4f, 0x4e,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f, 0x54,
	0x4f, 0x5f, 0x46, 0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59,
	0x53, 0x10, 0x12, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x4f, 0x5f,
	0x53, 0x49, 0x58, 0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x13, 0x12, 0x23, 0x0a, 0x1f,
	0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x49, 0x58, 0x54, 0x59,
	0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x49, 0x4e, 0x45, 0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10,
	0x14, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x44, 0x41,
	0x59, 0x10, 0x15, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x44, 0x41, 0x59,
	0x53, 0x10, 0x16, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x57, 0x4f, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x44,
	0x41, 0x59, 0x53, 0x10, 0x17, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x4f, 0x55,
	0x52, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x18, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x44, 0x4a, 0x55,
	0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x52, 0x5f, 0x54, 0x4f, 0x5f, 0x46,
	0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x19, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x44,
	0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x4f,
	0x5f, 0x53, 0x49, 0x58, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x1a, 0x12, 0x20, 0x0a, 0x1c, 0x41,
	0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x58, 0x5f, 0x54, 0x4f,
	0x5f, 0x53, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x1b, 0x12, 0x22, 0x0a,
	0x1e, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x56, 0x45,
	0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10,
	0x1c, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x45, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x49, 0x4e, 0x45, 0x5f, 0x44, 0x41,
	0x59, 0x53, 0x10, 0x1d, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x4e, 0x49, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x45, 0x4e, 0x5f, 0x44,
	0x41, 0x59, 0x53, 0x10, 0x1e, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x45, 0x4c, 0x45, 0x56, 0x45,
	0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x1f, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x44, 0x4a, 0x55,
	0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x4c, 0x45, 0x56, 0x45, 0x4e, 0x5f, 0x54, 0x4f,
	0x5f, 0x54, 0x57, 0x45, 0x4c, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x20, 0x12, 0x26,
	0x0a, 0x22, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x57, 0x45,
	0x4c, 0x56, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f,
	0x44, 0x41, 0x59, 0x53, 0x10, 0x21, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x54, 0x4f,
	0x5f, 0x46, 0x4f, 0x55, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x22,
	0x12, 0x2a, 0x0a, 0x26, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x52, 0x54, 0x45, 0x45, 0x4e, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x57, 0x45, 0x4e, 0x54,
	0x59, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x23, 0x12, 0x28, 0x0a, 0x24,
	0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x57, 0x45, 0x4e, 0x54,
	0x59, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f,
	0x44, 0x41, 0x59, 0x53, 0x10, 0x24, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x54, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x46,
	0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x25,
	0x12, 0x27, 0x0a, 0x23, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x46,
	0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x49, 0x58,
	0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x26, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x44, 0x4a,
	0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x58, 0x54, 0x59, 0x5f, 0x54, 0x4f,
	0x5f, 0x4e, 0x49, 0x4e, 0x45, 0x54, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x27, 0x12, 0x38,
	0x0a, 0x34, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x49, 0x4e,
	0x45, 0x54, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4f, 0x4e, 0x45, 0x5f, 0x48, 0x55, 0x4e, 0x44, 0x52,
	0x45, 0x44, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x54, 0x59, 0x5f, 0x46, 0x49, 0x56,
	0x45, 0x5f, 0x44, 0x41, 0x59, 0x53, 0x10, 0x28, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x4e, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x29,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x44, 0x4a, 0x55, 0x53, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x2a, 0x42, 0xfe, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42,
	0x24, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x41, 0x64, 0x6a,
	0x75, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x45,
	0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescData = file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDesc
)

func file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDescData
}

var file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_goTypes = []interface{}{
	(ConversionOrAdjustmentLagBucketEnum_ConversionOrAdjustmentLagBucket)(0), // 0: google.ads.googleads.v12.enums.ConversionOrAdjustmentLagBucketEnum.ConversionOrAdjustmentLagBucket
	(*ConversionOrAdjustmentLagBucketEnum)(nil),                              // 1: google.ads.googleads.v12.enums.ConversionOrAdjustmentLagBucketEnum
}
var file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_init() }
func file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_init() {
	if File_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionOrAdjustmentLagBucketEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto = out.File
	file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_rawDesc = nil
	file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_goTypes = nil
	file_google_ads_googleads_v12_enums_conversion_or_adjustment_lag_bucket_proto_depIdxs = nil
}
