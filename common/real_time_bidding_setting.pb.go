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
// source: google/ads/googleads/v17/common/real_time_bidding_setting.proto

package common

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

// Settings for Real-Time Bidding, a feature only available for campaigns
// targeting the Ad Exchange network.
type RealTimeBiddingSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the campaign is opted in to real-time bidding.
	OptIn *bool `protobuf:"varint,2,opt,name=opt_in,json=optIn,proto3,oneof" json:"opt_in,omitempty"`
}

func (x *RealTimeBiddingSetting) Reset() {
	*x = RealTimeBiddingSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RealTimeBiddingSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RealTimeBiddingSetting) ProtoMessage() {}

func (x *RealTimeBiddingSetting) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RealTimeBiddingSetting.ProtoReflect.Descriptor instead.
func (*RealTimeBiddingSetting) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescGZIP(), []int{0}
}

func (x *RealTimeBiddingSetting) GetOptIn() bool {
	if x != nil && x.OptIn != nil {
		return *x.OptIn
	}
	return false
}

var File_google_ads_googleads_v17_common_real_time_bidding_setting_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x62, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x06,
	0x6f, 0x70, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x05,
	0x6f, 0x70, 0x74, 0x49, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x70, 0x74,
	0x5f, 0x69, 0x6e, 0x42, 0xfb, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x1b, 0x52, 0x65, 0x61,
	0x6c, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x37, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x37, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x5c, 0x56, 0x31, 0x37, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xea, 0x02, 0x23, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescData = file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDesc
)

func file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDescData
}

var file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_goTypes = []any{
	(*RealTimeBiddingSetting)(nil), // 0: google.ads.googleads.v17.common.RealTimeBiddingSetting
}
var file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_init() }
func file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_init() {
	if File_google_ads_googleads_v17_common_real_time_bidding_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RealTimeBiddingSetting); i {
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
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_common_real_time_bidding_setting_proto = out.File
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_rawDesc = nil
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_goTypes = nil
	file_google_ads_googleads_v17_common_real_time_bidding_setting_proto_depIdxs = nil
}
