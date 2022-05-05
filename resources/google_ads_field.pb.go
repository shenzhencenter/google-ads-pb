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
// source: google/ads/googleads/v10/resources/google_ads_field.proto

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

// A field or resource (artifact) used by GoogleAdsService.
type GoogleAdsField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the artifact.
	// Artifact resource names have the form:
	//
	// `googleAdsFields/{name}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The name of the artifact.
	Name *string `protobuf:"bytes,21,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Output only. The category of the artifact.
	Category enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory `protobuf:"varint,3,opt,name=category,proto3,enum=google.ads.googleads.v10.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory" json:"category,omitempty"`
	// Output only. Whether the artifact can be used in a SELECT clause in search
	// queries.
	Selectable *bool `protobuf:"varint,22,opt,name=selectable,proto3,oneof" json:"selectable,omitempty"`
	// Output only. Whether the artifact can be used in a WHERE clause in search
	// queries.
	Filterable *bool `protobuf:"varint,23,opt,name=filterable,proto3,oneof" json:"filterable,omitempty"`
	// Output only. Whether the artifact can be used in a ORDER BY clause in search
	// queries.
	Sortable *bool `protobuf:"varint,24,opt,name=sortable,proto3,oneof" json:"sortable,omitempty"`
	// Output only. The names of all resources, segments, and metrics that are selectable with
	// the described artifact.
	SelectableWith []string `protobuf:"bytes,25,rep,name=selectable_with,json=selectableWith,proto3" json:"selectable_with,omitempty"`
	// Output only. The names of all resources that are selectable with the described
	// artifact. Fields from these resources do not segment metrics when included
	// in search queries.
	//
	// This field is only set for artifacts whose category is RESOURCE.
	AttributeResources []string `protobuf:"bytes,26,rep,name=attribute_resources,json=attributeResources,proto3" json:"attribute_resources,omitempty"`
	// Output only. This field lists the names of all metrics that are selectable with the
	// described artifact when it is used in the FROM clause.
	// It is only set for artifacts whose category is RESOURCE.
	Metrics []string `protobuf:"bytes,27,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Output only. This field lists the names of all artifacts, whether a segment or another
	// resource, that segment metrics when included in search queries and when the
	// described artifact is used in the FROM clause. It is only set for artifacts
	// whose category is RESOURCE.
	Segments []string `protobuf:"bytes,28,rep,name=segments,proto3" json:"segments,omitempty"`
	// Output only. Values the artifact can assume if it is a field of type ENUM.
	//
	// This field is only set for artifacts of category SEGMENT or ATTRIBUTE.
	EnumValues []string `protobuf:"bytes,29,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"`
	// Output only. This field determines the operators that can be used with the artifact
	// in WHERE clauses.
	DataType enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType `protobuf:"varint,12,opt,name=data_type,json=dataType,proto3,enum=google.ads.googleads.v10.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType" json:"data_type,omitempty"`
	// Output only. The URL of proto describing the artifact's data type.
	TypeUrl *string `protobuf:"bytes,30,opt,name=type_url,json=typeUrl,proto3,oneof" json:"type_url,omitempty"`
	// Output only. Whether the field artifact is repeated.
	IsRepeated *bool `protobuf:"varint,31,opt,name=is_repeated,json=isRepeated,proto3,oneof" json:"is_repeated,omitempty"`
}

func (x *GoogleAdsField) Reset() {
	*x = GoogleAdsField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v10_resources_google_ads_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsField) ProtoMessage() {}

func (x *GoogleAdsField) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v10_resources_google_ads_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsField.ProtoReflect.Descriptor instead.
func (*GoogleAdsField) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleAdsField) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *GoogleAdsField) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GoogleAdsField) GetCategory() enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory {
	if x != nil {
		return x.Category
	}
	return enums.GoogleAdsFieldCategoryEnum_UNSPECIFIED
}

func (x *GoogleAdsField) GetSelectable() bool {
	if x != nil && x.Selectable != nil {
		return *x.Selectable
	}
	return false
}

func (x *GoogleAdsField) GetFilterable() bool {
	if x != nil && x.Filterable != nil {
		return *x.Filterable
	}
	return false
}

func (x *GoogleAdsField) GetSortable() bool {
	if x != nil && x.Sortable != nil {
		return *x.Sortable
	}
	return false
}

func (x *GoogleAdsField) GetSelectableWith() []string {
	if x != nil {
		return x.SelectableWith
	}
	return nil
}

func (x *GoogleAdsField) GetAttributeResources() []string {
	if x != nil {
		return x.AttributeResources
	}
	return nil
}

func (x *GoogleAdsField) GetMetrics() []string {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *GoogleAdsField) GetSegments() []string {
	if x != nil {
		return x.Segments
	}
	return nil
}

func (x *GoogleAdsField) GetEnumValues() []string {
	if x != nil {
		return x.EnumValues
	}
	return nil
}

func (x *GoogleAdsField) GetDataType() enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType {
	if x != nil {
		return x.DataType
	}
	return enums.GoogleAdsFieldDataTypeEnum_UNSPECIFIED
}

func (x *GoogleAdsField) GetTypeUrl() string {
	if x != nil && x.TypeUrl != nil {
		return *x.TypeUrl
	}
	return ""
}

func (x *GoogleAdsField) GetIsRepeated() bool {
	if x != nil && x.IsRepeated != nil {
		return *x.IsRepeated
	}
	return false
}

var File_google_ads_googleads_v10_resources_google_ads_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a,
	0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x07, 0x0a,
	0x0e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x54, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x29, 0x0a, 0x27,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x72, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x48, 0x01, 0x52, 0x0a, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x28, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x17, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x02, 0x52, 0x0a, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x08, 0x73,
	0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x03, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x18, 0x19, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x12,
	0x34, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x1b, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x1c, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x51,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x30, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x23, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x04, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55,
	0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48,
	0x05, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01,
	0x3a, 0x50, 0xea, 0x41, 0x4d, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2f,
	0x7b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x7d, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x6f,
	0x72, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x85, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x13,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x30, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x30, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x30,
	0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescData = file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDesc
)

func file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDescData
}

var file_google_ads_googleads_v10_resources_google_ads_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v10_resources_google_ads_field_proto_goTypes = []interface{}{
	(*GoogleAdsField)(nil), // 0: google.ads.googleads.v10.resources.GoogleAdsField
	(enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory)(0), // 1: google.ads.googleads.v10.enums.GoogleAdsFieldCategoryEnum.GoogleAdsFieldCategory
	(enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType)(0), // 2: google.ads.googleads.v10.enums.GoogleAdsFieldDataTypeEnum.GoogleAdsFieldDataType
}
var file_google_ads_googleads_v10_resources_google_ads_field_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v10.resources.GoogleAdsField.category:type_name -> google.ads.googleads.v10.enums.GoogleAdsFieldCategoryEnum.GoogleAdsFieldCategory
	2, // 1: google.ads.googleads.v10.resources.GoogleAdsField.data_type:type_name -> google.ads.googleads.v10.enums.GoogleAdsFieldDataTypeEnum.GoogleAdsFieldDataType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v10_resources_google_ads_field_proto_init() }
func file_google_ads_googleads_v10_resources_google_ads_field_proto_init() {
	if File_google_ads_googleads_v10_resources_google_ads_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v10_resources_google_ads_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsField); i {
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
	file_google_ads_googleads_v10_resources_google_ads_field_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v10_resources_google_ads_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v10_resources_google_ads_field_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v10_resources_google_ads_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v10_resources_google_ads_field_proto = out.File
	file_google_ads_googleads_v10_resources_google_ads_field_proto_rawDesc = nil
	file_google_ads_googleads_v10_resources_google_ads_field_proto_goTypes = nil
	file_google_ads_googleads_v10_resources_google_ads_field_proto_depIdxs = nil
}
