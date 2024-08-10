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
// source: google/ads/googleads/v17/services/customer_sk_ad_network_conversion_value_schema_service.proto

package services

import (
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// A single update operation for a CustomerSkAdNetworkConversionValueSchema.
type CustomerSkAdNetworkConversionValueSchemaOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Update operation: The schema is expected to have a valid resource name.
	Update *resources.CustomerSkAdNetworkConversionValueSchema `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) Reset() {
	*x = CustomerSkAdNetworkConversionValueSchemaOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerSkAdNetworkConversionValueSchemaOperation) ProtoMessage() {}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerSkAdNetworkConversionValueSchemaOperation.ProtoReflect.Descriptor instead.
func (*CustomerSkAdNetworkConversionValueSchemaOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerSkAdNetworkConversionValueSchemaOperation) GetUpdate() *resources.CustomerSkAdNetworkConversionValueSchema {
	if x != nil {
		return x.Update
	}
	return nil
}

// Request message for
// [CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema][google.ads.googleads.v17.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema].
type MutateCustomerSkAdNetworkConversionValueSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the customer whose shared sets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The operation to perform.
	Operation *CustomerSkAdNetworkConversionValueSchemaOperation `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// Optional. If true, enables returning warnings. Warnings return error
	// messages and error codes without blocking the execution of the mutate
	// operation.
	EnableWarnings bool `protobuf:"varint,4,opt,name=enable_warnings,json=enableWarnings,proto3" json:"enable_warnings,omitempty"`
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaRequest) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetOperation() *CustomerSkAdNetworkConversionValueSchemaOperation {
	if x != nil {
		return x.Operation
	}
	return nil
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaRequest) GetEnableWarnings() bool {
	if x != nil {
		return x.EnableWarnings
	}
	return false
}

// The result for the CustomerSkAdNetworkConversionValueSchema mutate.
type MutateCustomerSkAdNetworkConversionValueSchemaResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the customer that was modified.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// App ID of the SkanConversionValue modified.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaResult) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResult) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// Response message for MutateCustomerSkAdNetworkConversionValueSchema.
type MutateCustomerSkAdNetworkConversionValueSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All results for the mutate.
	Result *MutateCustomerSkAdNetworkConversionValueSchemaResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// Non blocking errors that provides schema validation failure details.
	// Returned only when enable_warnings = true.
	Warning *status.Status `protobuf:"bytes,2,opt,name=warning,proto3" json:"warning,omitempty"`
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) Reset() {
	*x = MutateCustomerSkAdNetworkConversionValueSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerSkAdNetworkConversionValueSchemaResponse) ProtoMessage() {}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerSkAdNetworkConversionValueSchemaResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerSkAdNetworkConversionValueSchemaResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) GetResult() *MutateCustomerSkAdNetworkConversionValueSchemaResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *MutateCustomerSkAdNetworkConversionValueSchemaResponse) GetWarning() *status.Status {
	if x != nil {
		return x.Warning
	}
	return nil
}

var File_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc = []byte{
	0x0a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x6b, 0x5f,
	0x61, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x57, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x73, 0x6b, 0x5f, 0x61, 0x64, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x31, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x64,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x35, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x72, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x54, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53,
	0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x34, 0x4d, 0x75, 0x74, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x6b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x46, 0xfa, 0x41, 0x43, 0x0a, 0x41, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b,
	0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x36, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x57,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x2c, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x32, 0xbc, 0x03,
	0x0a, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xc1, 0x02, 0x0a, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x58, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x59,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x54, 0x3a, 0x01, 0x2a, 0x22, 0x4f, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6b, 0x41,
	0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x3a, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0xa0, 0x02, 0x0a,
	0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x34, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x53, 0x6b, 0x41, 0x64, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData = file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc
)

func file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDescData
}

var file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes = []any{
	(*CustomerSkAdNetworkConversionValueSchemaOperation)(nil),      // 0: google.ads.googleads.v17.services.CustomerSkAdNetworkConversionValueSchemaOperation
	(*MutateCustomerSkAdNetworkConversionValueSchemaRequest)(nil),  // 1: google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest
	(*MutateCustomerSkAdNetworkConversionValueSchemaResult)(nil),   // 2: google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaResult
	(*MutateCustomerSkAdNetworkConversionValueSchemaResponse)(nil), // 3: google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse
	(*resources.CustomerSkAdNetworkConversionValueSchema)(nil),     // 4: google.ads.googleads.v17.resources.CustomerSkAdNetworkConversionValueSchema
	(*status.Status)(nil), // 5: google.rpc.Status
}
var file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs = []int32{
	4, // 0: google.ads.googleads.v17.services.CustomerSkAdNetworkConversionValueSchemaOperation.update:type_name -> google.ads.googleads.v17.resources.CustomerSkAdNetworkConversionValueSchema
	0, // 1: google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest.operation:type_name -> google.ads.googleads.v17.services.CustomerSkAdNetworkConversionValueSchemaOperation
	2, // 2: google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse.result:type_name -> google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaResult
	5, // 3: google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse.warning:type_name -> google.rpc.Status
	1, // 4: google.ads.googleads.v17.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema:input_type -> google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaRequest
	3, // 5: google.ads.googleads.v17.services.CustomerSkAdNetworkConversionValueSchemaService.MutateCustomerSkAdNetworkConversionValueSchema:output_type -> google.ads.googleads.v17.services.MutateCustomerSkAdNetworkConversionValueSchemaResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_init()
}
func file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_init() {
	if File_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CustomerSkAdNetworkConversionValueSchemaOperation); i {
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
		file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MutateCustomerSkAdNetworkConversionValueSchemaRequest); i {
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
		file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MutateCustomerSkAdNetworkConversionValueSchemaResult); i {
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
		file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MutateCustomerSkAdNetworkConversionValueSchemaResponse); i {
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
			RawDescriptor: file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto = out.File
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_rawDesc = nil
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_goTypes = nil
	file_google_ads_googleads_v17_services_customer_sk_ad_network_conversion_value_schema_service_proto_depIdxs = nil
}
