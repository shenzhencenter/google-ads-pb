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
// source: google/ads/googleads/v18/resources/asset_set_type_view.proto

package resources

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// An asset set type view.
// This view reports non-overcounted metrics for each asset set type. Child
// asset set types are not included in this report. Their stats are aggregated
// under the parent asset set type.
type AssetSetTypeView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the asset set type view.
	// Asset set type view resource names have the form:
	//
	// `customers/{customer_id}/assetSetTypeViews/{asset_set_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The asset set type of the asset set type view.
	AssetSetType enums.AssetSetTypeEnum_AssetSetType `protobuf:"varint,3,opt,name=asset_set_type,json=assetSetType,proto3,enum=google.ads.googleads.v18.enums.AssetSetTypeEnum_AssetSetType" json:"asset_set_type,omitempty"`
}

func (x *AssetSetTypeView) Reset() {
	*x = AssetSetTypeView{}
	mi := &file_google_ads_googleads_v18_resources_asset_set_type_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetSetTypeView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetTypeView) ProtoMessage() {}

func (x *AssetSetTypeView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v18_resources_asset_set_type_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetTypeView.ProtoReflect.Descriptor instead.
func (*AssetSetTypeView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescGZIP(), []int{0}
}

func (x *AssetSetTypeView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetSetTypeView) GetAssetSetType() enums.AssetSetTypeEnum_AssetSetType {
	if x != nil {
		return x.AssetSetType
	}
	return enums.AssetSetTypeEnum_AssetSetType(0)
}

var File_google_ads_googleads_v18_resources_asset_set_type_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x1a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x56, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x31, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x2b, 0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x56, 0x69,
	0x65, 0x77, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x68, 0x0a, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x38, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x6a, 0xea, 0x41, 0x67, 0x0a,
	0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x3a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x7d, 0x42, 0x87, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x42, 0x15, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x56,
	0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x38, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x38, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x38, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescData = file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDesc
)

func file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDescData
}

var file_google_ads_googleads_v18_resources_asset_set_type_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v18_resources_asset_set_type_view_proto_goTypes = []any{
	(*AssetSetTypeView)(nil),                 // 0: google.ads.googleads.v18.resources.AssetSetTypeView
	(enums.AssetSetTypeEnum_AssetSetType)(0), // 1: google.ads.googleads.v18.enums.AssetSetTypeEnum.AssetSetType
}
var file_google_ads_googleads_v18_resources_asset_set_type_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v18.resources.AssetSetTypeView.asset_set_type:type_name -> google.ads.googleads.v18.enums.AssetSetTypeEnum.AssetSetType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v18_resources_asset_set_type_view_proto_init() }
func file_google_ads_googleads_v18_resources_asset_set_type_view_proto_init() {
	if File_google_ads_googleads_v18_resources_asset_set_type_view_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v18_resources_asset_set_type_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v18_resources_asset_set_type_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v18_resources_asset_set_type_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v18_resources_asset_set_type_view_proto = out.File
	file_google_ads_googleads_v18_resources_asset_set_type_view_proto_rawDesc = nil
	file_google_ads_googleads_v18_resources_asset_set_type_view_proto_goTypes = nil
	file_google_ads_googleads_v18_resources_asset_set_type_view_proto_depIdxs = nil
}
