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
// source: google/ads/googleads/v16/resources/product_link.proto

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

// Represents the data sharing connection between  a Google
// Ads customer and another product.
type ProductLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Resource name of the product link.
	// ProductLink resource names have the form:
	//
	// `customers/{customer_id}/productLinks/{product_link_id} `
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the link.
	// This field is read only.
	ProductLinkId *int64 `protobuf:"varint,2,opt,name=product_link_id,json=productLinkId,proto3,oneof" json:"product_link_id,omitempty"`
	// Output only. The type of the linked product.
	Type enums.LinkedProductTypeEnum_LinkedProductType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v16.enums.LinkedProductTypeEnum_LinkedProductType" json:"type,omitempty"`
	// A product linked to this account.
	//
	// Types that are assignable to LinkedProduct:
	//
	//	*ProductLink_DataPartner
	//	*ProductLink_GoogleAds
	//	*ProductLink_MerchantCenter
	//	*ProductLink_AdvertisingPartner
	LinkedProduct isProductLink_LinkedProduct `protobuf_oneof:"linked_product"`
}

func (x *ProductLink) Reset() {
	*x = ProductLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductLink) ProtoMessage() {}

func (x *ProductLink) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductLink.ProtoReflect.Descriptor instead.
func (*ProductLink) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_product_link_proto_rawDescGZIP(), []int{0}
}

func (x *ProductLink) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ProductLink) GetProductLinkId() int64 {
	if x != nil && x.ProductLinkId != nil {
		return *x.ProductLinkId
	}
	return 0
}

func (x *ProductLink) GetType() enums.LinkedProductTypeEnum_LinkedProductType {
	if x != nil {
		return x.Type
	}
	return enums.LinkedProductTypeEnum_LinkedProductType(0)
}

func (m *ProductLink) GetLinkedProduct() isProductLink_LinkedProduct {
	if m != nil {
		return m.LinkedProduct
	}
	return nil
}

func (x *ProductLink) GetDataPartner() *DataPartnerIdentifier {
	if x, ok := x.GetLinkedProduct().(*ProductLink_DataPartner); ok {
		return x.DataPartner
	}
	return nil
}

func (x *ProductLink) GetGoogleAds() *GoogleAdsIdentifier {
	if x, ok := x.GetLinkedProduct().(*ProductLink_GoogleAds); ok {
		return x.GoogleAds
	}
	return nil
}

func (x *ProductLink) GetMerchantCenter() *MerchantCenterIdentifier {
	if x, ok := x.GetLinkedProduct().(*ProductLink_MerchantCenter); ok {
		return x.MerchantCenter
	}
	return nil
}

func (x *ProductLink) GetAdvertisingPartner() *AdvertisingPartnerIdentifier {
	if x, ok := x.GetLinkedProduct().(*ProductLink_AdvertisingPartner); ok {
		return x.AdvertisingPartner
	}
	return nil
}

type isProductLink_LinkedProduct interface {
	isProductLink_LinkedProduct()
}

type ProductLink_DataPartner struct {
	// Immutable. Data partner link.
	DataPartner *DataPartnerIdentifier `protobuf:"bytes,4,opt,name=data_partner,json=dataPartner,proto3,oneof"`
}

type ProductLink_GoogleAds struct {
	// Immutable. Google Ads link.
	GoogleAds *GoogleAdsIdentifier `protobuf:"bytes,5,opt,name=google_ads,json=googleAds,proto3,oneof"`
}

type ProductLink_MerchantCenter struct {
	// Immutable. Google Merchant Center link.
	MerchantCenter *MerchantCenterIdentifier `protobuf:"bytes,12,opt,name=merchant_center,json=merchantCenter,proto3,oneof"`
}

type ProductLink_AdvertisingPartner struct {
	// Output only. Advertising Partner link.
	AdvertisingPartner *AdvertisingPartnerIdentifier `protobuf:"bytes,13,opt,name=advertising_partner,json=advertisingPartner,proto3,oneof"`
}

func (*ProductLink_DataPartner) isProductLink_LinkedProduct() {}

func (*ProductLink_GoogleAds) isProductLink_LinkedProduct() {}

func (*ProductLink_MerchantCenter) isProductLink_LinkedProduct() {}

func (*ProductLink_AdvertisingPartner) isProductLink_LinkedProduct() {}

// The identifier for Data Partner account.
type DataPartnerIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The customer ID of the Data partner account.
	// This field is required and should not be empty when creating a new
	// data partner link. It is unable to be modified after the creation of
	// the link.
	DataPartnerId *int64 `protobuf:"varint,1,opt,name=data_partner_id,json=dataPartnerId,proto3,oneof" json:"data_partner_id,omitempty"`
}

func (x *DataPartnerIdentifier) Reset() {
	*x = DataPartnerIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPartnerIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPartnerIdentifier) ProtoMessage() {}

func (x *DataPartnerIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPartnerIdentifier.ProtoReflect.Descriptor instead.
func (*DataPartnerIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_product_link_proto_rawDescGZIP(), []int{1}
}

func (x *DataPartnerIdentifier) GetDataPartnerId() int64 {
	if x != nil && x.DataPartnerId != nil {
		return *x.DataPartnerId
	}
	return 0
}

// The identifier for Google Ads account.
type GoogleAdsIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Google Ads account.
	// This field is required and should not be empty when creating a new
	// Google Ads link. It is unable to be modified after the creation of
	// the link.
	Customer *string `protobuf:"bytes,1,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
}

func (x *GoogleAdsIdentifier) Reset() {
	*x = GoogleAdsIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleAdsIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleAdsIdentifier) ProtoMessage() {}

func (x *GoogleAdsIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleAdsIdentifier.ProtoReflect.Descriptor instead.
func (*GoogleAdsIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_product_link_proto_rawDescGZIP(), []int{2}
}

func (x *GoogleAdsIdentifier) GetCustomer() string {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return ""
}

// The identifier for Google Merchant Center account
type MerchantCenterIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The customer ID of the Google Merchant Center account.
	// This field is required and should not be empty when creating a new
	// Merchant Center link. It is unable to be modified after the creation of
	// the link.
	MerchantCenterId *int64 `protobuf:"varint,1,opt,name=merchant_center_id,json=merchantCenterId,proto3,oneof" json:"merchant_center_id,omitempty"`
}

func (x *MerchantCenterIdentifier) Reset() {
	*x = MerchantCenterIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantCenterIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantCenterIdentifier) ProtoMessage() {}

func (x *MerchantCenterIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantCenterIdentifier.ProtoReflect.Descriptor instead.
func (*MerchantCenterIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_product_link_proto_rawDescGZIP(), []int{3}
}

func (x *MerchantCenterIdentifier) GetMerchantCenterId() int64 {
	if x != nil && x.MerchantCenterId != nil {
		return *x.MerchantCenterId
	}
	return 0
}

// The identifier for the Advertising Partner Google Ads account.
type AdvertisingPartnerIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The resource name of the advertising partner Google Ads
	// account. This field is required and should not be empty when creating a new
	// Advertising Partner link. It is unable to be modified after the creation of
	// the link.
	Customer *string `protobuf:"bytes,1,opt,name=customer,proto3,oneof" json:"customer,omitempty"`
}

func (x *AdvertisingPartnerIdentifier) Reset() {
	*x = AdvertisingPartnerIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvertisingPartnerIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertisingPartnerIdentifier) ProtoMessage() {}

func (x *AdvertisingPartnerIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertisingPartnerIdentifier.ProtoReflect.Descriptor instead.
func (*AdvertisingPartnerIdentifier) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_resources_product_link_proto_rawDescGZIP(), []int{4}
}

func (x *AdvertisingPartnerIdentifier) GetCustomer() string {
	if x != nil && x.Customer != nil {
		return *x.Customer
	}
	return ""
}

var File_google_ads_googleads_v16_resources_product_link_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_resources_product_link_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x38, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa9, 0x06, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x51, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2c, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x26,
	0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x60, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x63, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x5d, 0x0a,
	0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48,
	0x00, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x12, 0x6c, 0x0a, 0x0f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x78, 0x0a, 0x13, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x41, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x00,
	0x52, 0x12, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x3a, 0x61, 0xea, 0x41, 0x5e, 0x0a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x36, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c,
	0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x10, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x22, 0x5d, 0x0a,
	0x15, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x13,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x69, 0x0a, 0x18,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x12, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x00, 0x52, 0x10, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x1c, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41,
	0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x82, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x10, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x36, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x36, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_resources_product_link_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_resources_product_link_proto_rawDescData = file_google_ads_googleads_v16_resources_product_link_proto_rawDesc
)

func file_google_ads_googleads_v16_resources_product_link_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_resources_product_link_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_resources_product_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_resources_product_link_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_resources_product_link_proto_rawDescData
}

var file_google_ads_googleads_v16_resources_product_link_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v16_resources_product_link_proto_goTypes = []interface{}{
	(*ProductLink)(nil),                                // 0: google.ads.googleads.v16.resources.ProductLink
	(*DataPartnerIdentifier)(nil),                      // 1: google.ads.googleads.v16.resources.DataPartnerIdentifier
	(*GoogleAdsIdentifier)(nil),                        // 2: google.ads.googleads.v16.resources.GoogleAdsIdentifier
	(*MerchantCenterIdentifier)(nil),                   // 3: google.ads.googleads.v16.resources.MerchantCenterIdentifier
	(*AdvertisingPartnerIdentifier)(nil),               // 4: google.ads.googleads.v16.resources.AdvertisingPartnerIdentifier
	(enums.LinkedProductTypeEnum_LinkedProductType)(0), // 5: google.ads.googleads.v16.enums.LinkedProductTypeEnum.LinkedProductType
}
var file_google_ads_googleads_v16_resources_product_link_proto_depIdxs = []int32{
	5, // 0: google.ads.googleads.v16.resources.ProductLink.type:type_name -> google.ads.googleads.v16.enums.LinkedProductTypeEnum.LinkedProductType
	1, // 1: google.ads.googleads.v16.resources.ProductLink.data_partner:type_name -> google.ads.googleads.v16.resources.DataPartnerIdentifier
	2, // 2: google.ads.googleads.v16.resources.ProductLink.google_ads:type_name -> google.ads.googleads.v16.resources.GoogleAdsIdentifier
	3, // 3: google.ads.googleads.v16.resources.ProductLink.merchant_center:type_name -> google.ads.googleads.v16.resources.MerchantCenterIdentifier
	4, // 4: google.ads.googleads.v16.resources.ProductLink.advertising_partner:type_name -> google.ads.googleads.v16.resources.AdvertisingPartnerIdentifier
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_resources_product_link_proto_init() }
func file_google_ads_googleads_v16_resources_product_link_proto_init() {
	if File_google_ads_googleads_v16_resources_product_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductLink); i {
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
		file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPartnerIdentifier); i {
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
		file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleAdsIdentifier); i {
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
		file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantCenterIdentifier); i {
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
		file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvertisingPartnerIdentifier); i {
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
	file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProductLink_DataPartner)(nil),
		(*ProductLink_GoogleAds)(nil),
		(*ProductLink_MerchantCenter)(nil),
		(*ProductLink_AdvertisingPartner)(nil),
	}
	file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_google_ads_googleads_v16_resources_product_link_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v16_resources_product_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v16_resources_product_link_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_resources_product_link_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v16_resources_product_link_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_resources_product_link_proto = out.File
	file_google_ads_googleads_v16_resources_product_link_proto_rawDesc = nil
	file_google_ads_googleads_v16_resources_product_link_proto_goTypes = nil
	file_google_ads_googleads_v16_resources_product_link_proto_depIdxs = nil
}
