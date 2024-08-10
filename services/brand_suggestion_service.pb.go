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
// source: google/ads/googleads/v17/services/brand_suggestion_service.proto

package services

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

// Request message for
// [BrandSuggestionService.SuggestBrands][google.ads.googleads.v17.services.BrandSuggestionService.SuggestBrands].
type SuggestBrandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer onto which to apply the brand suggestion
	// operation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The prefix of a brand name.
	BrandPrefix *string `protobuf:"bytes,2,opt,name=brand_prefix,json=brandPrefix,proto3,oneof" json:"brand_prefix,omitempty"`
	// Optional. Ids of the brands already selected by advertisers. They will be
	// excluded in response. These are expected to be brand ids not brand names.
	SelectedBrands []string `protobuf:"bytes,3,rep,name=selected_brands,json=selectedBrands,proto3" json:"selected_brands,omitempty"`
}

func (x *SuggestBrandsRequest) Reset() {
	*x = SuggestBrandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestBrandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestBrandsRequest) ProtoMessage() {}

func (x *SuggestBrandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestBrandsRequest.ProtoReflect.Descriptor instead.
func (*SuggestBrandsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestBrandsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *SuggestBrandsRequest) GetBrandPrefix() string {
	if x != nil && x.BrandPrefix != nil {
		return *x.BrandPrefix
	}
	return ""
}

func (x *SuggestBrandsRequest) GetSelectedBrands() []string {
	if x != nil {
		return x.SelectedBrands
	}
	return nil
}

// Response message for
// [BrandSuggestionService.SuggestBrands][google.ads.googleads.v17.services.BrandSuggestionService.SuggestBrands].
type SuggestBrandsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Generated brand suggestions of verified brands for the given prefix.
	Brands []*BrandSuggestion `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *SuggestBrandsResponse) Reset() {
	*x = SuggestBrandsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestBrandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestBrandsResponse) ProtoMessage() {}

func (x *SuggestBrandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestBrandsResponse.ProtoReflect.Descriptor instead.
func (*SuggestBrandsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescGZIP(), []int{1}
}

func (x *SuggestBrandsResponse) GetBrands() []*BrandSuggestion {
	if x != nil {
		return x.Brands
	}
	return nil
}

// Information of brand suggestion.
type BrandSuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id for the brand. It would be CKG MID for verified/global scoped brands.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the brand.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Urls which uniquely identify the brand.
	Urls []string `protobuf:"bytes,3,rep,name=urls,proto3" json:"urls,omitempty"`
	// Current state of the brand.
	State enums.BrandStateEnum_BrandState `protobuf:"varint,4,opt,name=state,proto3,enum=google.ads.googleads.v17.enums.BrandStateEnum_BrandState" json:"state,omitempty"`
}

func (x *BrandSuggestion) Reset() {
	*x = BrandSuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandSuggestion) ProtoMessage() {}

func (x *BrandSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandSuggestion.ProtoReflect.Descriptor instead.
func (*BrandSuggestion) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescGZIP(), []int{2}
}

func (x *BrandSuggestion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BrandSuggestion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BrandSuggestion) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *BrandSuggestion) GetState() enums.BrandStateEnum_BrandState {
	if x != nil {
		return x.State
	}
	return enums.BrandStateEnum_BrandState(0)
}

var File_google_ads_googleads_v17_services_brand_suggestion_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa8, 0x01, 0x0a, 0x14, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x63, 0x0a, 0x15, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x22,
	0x9a, 0x01, 0x0a, 0x0f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x4f, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x32, 0xb8, 0x02, 0x0a,
	0x16, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd6, 0x01, 0x0a, 0x0d, 0x53, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0xda, 0x41,
	0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x3a,
	0x01, 0x2a, 0x22, 0x2c, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x3a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73,
	0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x87, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x42, 0x1b, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescData = file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDesc
)

func file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDescData
}

var file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_ads_googleads_v17_services_brand_suggestion_service_proto_goTypes = []any{
	(*SuggestBrandsRequest)(nil),         // 0: google.ads.googleads.v17.services.SuggestBrandsRequest
	(*SuggestBrandsResponse)(nil),        // 1: google.ads.googleads.v17.services.SuggestBrandsResponse
	(*BrandSuggestion)(nil),              // 2: google.ads.googleads.v17.services.BrandSuggestion
	(enums.BrandStateEnum_BrandState)(0), // 3: google.ads.googleads.v17.enums.BrandStateEnum.BrandState
}
var file_google_ads_googleads_v17_services_brand_suggestion_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v17.services.SuggestBrandsResponse.brands:type_name -> google.ads.googleads.v17.services.BrandSuggestion
	3, // 1: google.ads.googleads.v17.services.BrandSuggestion.state:type_name -> google.ads.googleads.v17.enums.BrandStateEnum.BrandState
	0, // 2: google.ads.googleads.v17.services.BrandSuggestionService.SuggestBrands:input_type -> google.ads.googleads.v17.services.SuggestBrandsRequest
	1, // 3: google.ads.googleads.v17.services.BrandSuggestionService.SuggestBrands:output_type -> google.ads.googleads.v17.services.SuggestBrandsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_services_brand_suggestion_service_proto_init() }
func file_google_ads_googleads_v17_services_brand_suggestion_service_proto_init() {
	if File_google_ads_googleads_v17_services_brand_suggestion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SuggestBrandsRequest); i {
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
		file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SuggestBrandsResponse); i {
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
		file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BrandSuggestion); i {
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
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v17_services_brand_suggestion_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_services_brand_suggestion_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_services_brand_suggestion_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_services_brand_suggestion_service_proto = out.File
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_rawDesc = nil
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_goTypes = nil
	file_google_ads_googleads_v17_services_brand_suggestion_service_proto_depIdxs = nil
}
