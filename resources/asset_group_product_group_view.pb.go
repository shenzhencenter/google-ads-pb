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
// 	protoc        v3.21.9
// source: google/ads/googleads/v12/resources/asset_group_product_group_view.proto

package resources

import (
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

// An asset group product group view.
type AssetGroupProductGroupView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the asset group product group view.
	// Asset group product group view resource names have the form:
	//
	// `customers/{customer_id}/assetGroupProductGroupViews/{asset_group_id}~{listing_group_filter_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The asset group associated with the listing group filter.
	AssetGroup string `protobuf:"bytes,2,opt,name=asset_group,json=assetGroup,proto3" json:"asset_group,omitempty"`
	// Output only. The resource name of the asset group listing group filter.
	AssetGroupListingGroupFilter string `protobuf:"bytes,4,opt,name=asset_group_listing_group_filter,json=assetGroupListingGroupFilter,proto3" json:"asset_group_listing_group_filter,omitempty"`
}

func (x *AssetGroupProductGroupView) Reset() {
	*x = AssetGroupProductGroupView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetGroupProductGroupView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetGroupProductGroupView) ProtoMessage() {}

func (x *AssetGroupProductGroupView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetGroupProductGroupView.ProtoReflect.Descriptor instead.
func (*AssetGroupProductGroupView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescGZIP(), []int{0}
}

func (x *AssetGroupProductGroupView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AssetGroupProductGroupView) GetAssetGroup() string {
	if x != nil {
		return x.AssetGroup
	}
	return ""
}

func (x *AssetGroupProductGroupView) GetAssetGroupListingGroupFilter() string {
	if x != nil {
		return x.AssetGroupListingGroupFilter
	}
	return ""
}

var File_google_ads_googleads_v12_resources_asset_group_product_group_view_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDesc = []byte{
	0x0a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x1a, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x65, 0x77, 0x12, 0x60, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x3b, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x35, 0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x65, 0x77, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2b, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x61, 0x73,
	0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x85, 0x01, 0x0a, 0x20, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x3d, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x37, 0x0a, 0x35, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x1c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x3a, 0x99, 0x01, 0xea, 0x41, 0x95, 0x01, 0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x65, 0x77, 0x12, 0x5e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x69, 0x65, 0x77,
	0x73, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x7d, 0x7e, 0x7b, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x91, 0x02, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1f, 0x41, 0x73, 0x73, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56,
	0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescData = file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDesc
)

func file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDescData
}

var file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_goTypes = []interface{}{
	(*AssetGroupProductGroupView)(nil), // 0: google.ads.googleads.v12.resources.AssetGroupProductGroupView
}
var file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_init() }
func file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_init() {
	if File_google_ads_googleads_v12_resources_asset_group_product_group_view_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetGroupProductGroupView); i {
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
			RawDescriptor: file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_resources_asset_group_product_group_view_proto = out.File
	file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_rawDesc = nil
	file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_goTypes = nil
	file_google_ads_googleads_v12_resources_asset_group_product_group_view_proto_depIdxs = nil
}
