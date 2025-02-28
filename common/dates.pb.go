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
// source: google/ads/googleads/v19/common/dates.proto

package common

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
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

// A date range.
type DateRange struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The start date, in yyyy-mm-dd format. This date is inclusive.
	StartDate *string `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3,oneof" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format. This date is inclusive.
	EndDate       *string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3,oneof" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	mi := &file_google_ads_googleads_v19_common_dates_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_common_dates_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_common_dates_proto_rawDescGZIP(), []int{0}
}

func (x *DateRange) GetStartDate() string {
	if x != nil && x.StartDate != nil {
		return *x.StartDate
	}
	return ""
}

func (x *DateRange) GetEndDate() string {
	if x != nil && x.EndDate != nil {
		return *x.EndDate
	}
	return ""
}

// The year month range inclusive of the start and end months.
// Eg: A year month range to represent Jan 2020 would be: (Jan 2020, Jan 2020).
type YearMonthRange struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The inclusive start year month.
	Start *YearMonth `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// The inclusive end year month.
	End           *YearMonth `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YearMonthRange) Reset() {
	*x = YearMonthRange{}
	mi := &file_google_ads_googleads_v19_common_dates_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YearMonthRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YearMonthRange) ProtoMessage() {}

func (x *YearMonthRange) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_common_dates_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YearMonthRange.ProtoReflect.Descriptor instead.
func (*YearMonthRange) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_common_dates_proto_rawDescGZIP(), []int{1}
}

func (x *YearMonthRange) GetStart() *YearMonth {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *YearMonthRange) GetEnd() *YearMonth {
	if x != nil {
		return x.End
	}
	return nil
}

// Year month.
type YearMonth struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The year (for example, 2020).
	Year int64 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// The month of the year. (for example, FEBRUARY).
	Month         enums.MonthOfYearEnum_MonthOfYear `protobuf:"varint,2,opt,name=month,proto3,enum=google.ads.googleads.v19.enums.MonthOfYearEnum_MonthOfYear" json:"month,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YearMonth) Reset() {
	*x = YearMonth{}
	mi := &file_google_ads_googleads_v19_common_dates_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YearMonth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YearMonth) ProtoMessage() {}

func (x *YearMonth) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_common_dates_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YearMonth.ProtoReflect.Descriptor instead.
func (*YearMonth) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_common_dates_proto_rawDescGZIP(), []int{2}
}

func (x *YearMonth) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *YearMonth) GetMonth() enums.MonthOfYearEnum_MonthOfYear {
	if x != nil {
		return x.Month
	}
	return enums.MonthOfYearEnum_MonthOfYear(0)
}

var File_google_ads_googleads_v19_common_dates_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_common_dates_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x32,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x5f, 0x6f, 0x66, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x90, 0x01, 0x0a, 0x0e, 0x59, 0x65, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x65, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x3c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x59, 0x65, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x72, 0x0a, 0x09, 0x59, 0x65, 0x61, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x51, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x4f, 0x66, 0x59, 0x65, 0x61, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x4f, 0x66, 0x59, 0x65, 0x61, 0x72, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x42, 0xea, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0a,
	0x44, 0x61, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x39, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_google_ads_googleads_v19_common_dates_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_common_dates_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_common_dates_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_common_dates_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_common_dates_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_common_dates_proto_rawDesc), len(file_google_ads_googleads_v19_common_dates_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_common_dates_proto_rawDescData
}

var file_google_ads_googleads_v19_common_dates_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v19_common_dates_proto_goTypes = []any{
	(*DateRange)(nil),                      // 0: google.ads.googleads.v19.common.DateRange
	(*YearMonthRange)(nil),                 // 1: google.ads.googleads.v19.common.YearMonthRange
	(*YearMonth)(nil),                      // 2: google.ads.googleads.v19.common.YearMonth
	(enums.MonthOfYearEnum_MonthOfYear)(0), // 3: google.ads.googleads.v19.enums.MonthOfYearEnum.MonthOfYear
}
var file_google_ads_googleads_v19_common_dates_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v19.common.YearMonthRange.start:type_name -> google.ads.googleads.v19.common.YearMonth
	2, // 1: google.ads.googleads.v19.common.YearMonthRange.end:type_name -> google.ads.googleads.v19.common.YearMonth
	3, // 2: google.ads.googleads.v19.common.YearMonth.month:type_name -> google.ads.googleads.v19.enums.MonthOfYearEnum.MonthOfYear
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_common_dates_proto_init() }
func file_google_ads_googleads_v19_common_dates_proto_init() {
	if File_google_ads_googleads_v19_common_dates_proto != nil {
		return
	}
	file_google_ads_googleads_v19_common_dates_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_common_dates_proto_rawDesc), len(file_google_ads_googleads_v19_common_dates_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v19_common_dates_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_common_dates_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v19_common_dates_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_common_dates_proto = out.File
	file_google_ads_googleads_v19_common_dates_proto_goTypes = nil
	file_google_ads_googleads_v19_common_dates_proto_depIdxs = nil
}
