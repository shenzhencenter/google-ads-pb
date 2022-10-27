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
// source: google/ads/googleads/v12/services/customer_manager_link_service.proto

package services

import (
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for [CustomerManagerLinkService.MutateCustomerManagerLink][google.ads.googleads.v12.services.CustomerManagerLinkService.MutateCustomerManagerLink].
type MutateCustomerManagerLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose customer manager links are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual customer manager links.
	Operations []*CustomerManagerLinkOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,3,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MutateCustomerManagerLinkRequest) Reset() {
	*x = MutateCustomerManagerLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerManagerLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerManagerLinkRequest) ProtoMessage() {}

func (x *MutateCustomerManagerLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerManagerLinkRequest.ProtoReflect.Descriptor instead.
func (*MutateCustomerManagerLinkRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateCustomerManagerLinkRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateCustomerManagerLinkRequest) GetOperations() []*CustomerManagerLinkOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateCustomerManagerLinkRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Request message for [CustomerManagerLinkService.MoveManagerLink][google.ads.googleads.v12.services.CustomerManagerLinkService.MoveManagerLink].
type MoveManagerLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the client customer that is being moved.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The resource name of the previous CustomerManagerLink.
	// The resource name has the form:
	// `customers/{customer_id}/customerManagerLinks/{manager_customer_id}~{manager_link_id}`
	PreviousCustomerManagerLink string `protobuf:"bytes,2,opt,name=previous_customer_manager_link,json=previousCustomerManagerLink,proto3" json:"previous_customer_manager_link,omitempty"`
	// Required. The resource name of the new manager customer that the client wants to move
	// to. Customer resource names have the format: "customers/{customer_id}"
	NewManager string `protobuf:"bytes,3,opt,name=new_manager,json=newManager,proto3" json:"new_manager,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
}

func (x *MoveManagerLinkRequest) Reset() {
	*x = MoveManagerLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveManagerLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveManagerLinkRequest) ProtoMessage() {}

func (x *MoveManagerLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveManagerLinkRequest.ProtoReflect.Descriptor instead.
func (*MoveManagerLinkRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP(), []int{1}
}

func (x *MoveManagerLinkRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MoveManagerLinkRequest) GetPreviousCustomerManagerLink() string {
	if x != nil {
		return x.PreviousCustomerManagerLink
	}
	return ""
}

func (x *MoveManagerLinkRequest) GetNewManager() string {
	if x != nil {
		return x.NewManager
	}
	return ""
}

func (x *MoveManagerLinkRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

// Updates the status of a CustomerManagerLink.
// The following actions are possible:
// 1. Update operation with status ACTIVE accepts a pending invitation.
// 2. Update operation with status REFUSED declines a pending invitation.
// 3. Update operation with status INACTIVE terminates link to manager.
type CustomerManagerLinkOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*CustomerManagerLinkOperation_Update
	Operation isCustomerManagerLinkOperation_Operation `protobuf_oneof:"operation"`
}

func (x *CustomerManagerLinkOperation) Reset() {
	*x = CustomerManagerLinkOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerManagerLinkOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerManagerLinkOperation) ProtoMessage() {}

func (x *CustomerManagerLinkOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerManagerLinkOperation.ProtoReflect.Descriptor instead.
func (*CustomerManagerLinkOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerManagerLinkOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *CustomerManagerLinkOperation) GetOperation() isCustomerManagerLinkOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *CustomerManagerLinkOperation) GetUpdate() *resources.CustomerManagerLink {
	if x, ok := x.GetOperation().(*CustomerManagerLinkOperation_Update); ok {
		return x.Update
	}
	return nil
}

type isCustomerManagerLinkOperation_Operation interface {
	isCustomerManagerLinkOperation_Operation()
}

type CustomerManagerLinkOperation_Update struct {
	// Update operation: The link is expected to have a valid resource name.
	Update *resources.CustomerManagerLink `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

func (*CustomerManagerLinkOperation_Update) isCustomerManagerLinkOperation_Operation() {}

// Response message for a CustomerManagerLink mutate.
type MutateCustomerManagerLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A result that identifies the resource affected by the mutate request.
	Results []*MutateCustomerManagerLinkResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateCustomerManagerLinkResponse) Reset() {
	*x = MutateCustomerManagerLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerManagerLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerManagerLinkResponse) ProtoMessage() {}

func (x *MutateCustomerManagerLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerManagerLinkResponse.ProtoReflect.Descriptor instead.
func (*MutateCustomerManagerLinkResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateCustomerManagerLinkResponse) GetResults() []*MutateCustomerManagerLinkResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// Response message for a CustomerManagerLink moveManagerLink.
type MoveManagerLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations. Represents a CustomerManagerLink
	// resource of the newly created link between client customer and new manager
	// customer.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MoveManagerLinkResponse) Reset() {
	*x = MoveManagerLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveManagerLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveManagerLinkResponse) ProtoMessage() {}

func (x *MoveManagerLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveManagerLinkResponse.ProtoReflect.Descriptor instead.
func (*MoveManagerLinkResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP(), []int{4}
}

func (x *MoveManagerLinkResponse) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// The result for the customer manager link mutate.
type MutateCustomerManagerLinkResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *MutateCustomerManagerLinkResult) Reset() {
	*x = MutateCustomerManagerLinkResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateCustomerManagerLinkResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateCustomerManagerLinkResult) ProtoMessage() {}

func (x *MutateCustomerManagerLinkResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateCustomerManagerLinkResult.ProtoReflect.Descriptor instead.
func (*MutateCustomerManagerLinkResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP(), []int{5}
}

func (x *MutateCustomerManagerLinkResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

var File_google_ads_googleads_v12_services_customer_manager_link_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd3, 0x01, 0x0a, 0x20, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x64, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xd3, 0x01, 0x0a, 0x16, 0x4d, 0x6f, 0x76, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x48, 0x0a, 0x1e, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x1b, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x24, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xbb, 0x01, 0x0a, 0x1c,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c,
	0x69, 0x6e, 0x6b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x51, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x0b, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x01, 0x0a, 0x21, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x71, 0x0a,
	0x17, 0x4d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x31, 0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x79, 0x0a, 0x1f, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x56, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x41, 0x2e, 0x0a,
	0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x81, 0x05, 0x0a, 0x1a,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c,
	0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x02, 0x0a, 0x19, 0x4d,
	0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x5e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x22, 0x3a, 0x2f, 0x76, 0x31,
	0x32, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x92, 0x02, 0x0a, 0x0f, 0x4d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x48, 0x22, 0x43, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x3a, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x3a, 0x01, 0x2a, 0xda,
	0x41, 0x36, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x2c, 0x6e, 0x65, 0x77,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42,
	0x8b, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1f, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31, 0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x32, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x32, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescData = file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDesc
)

func file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDescData
}

var file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_ads_googleads_v12_services_customer_manager_link_service_proto_goTypes = []interface{}{
	(*MutateCustomerManagerLinkRequest)(nil),  // 0: google.ads.googleads.v12.services.MutateCustomerManagerLinkRequest
	(*MoveManagerLinkRequest)(nil),            // 1: google.ads.googleads.v12.services.MoveManagerLinkRequest
	(*CustomerManagerLinkOperation)(nil),      // 2: google.ads.googleads.v12.services.CustomerManagerLinkOperation
	(*MutateCustomerManagerLinkResponse)(nil), // 3: google.ads.googleads.v12.services.MutateCustomerManagerLinkResponse
	(*MoveManagerLinkResponse)(nil),           // 4: google.ads.googleads.v12.services.MoveManagerLinkResponse
	(*MutateCustomerManagerLinkResult)(nil),   // 5: google.ads.googleads.v12.services.MutateCustomerManagerLinkResult
	(*fieldmaskpb.FieldMask)(nil),             // 6: google.protobuf.FieldMask
	(*resources.CustomerManagerLink)(nil),     // 7: google.ads.googleads.v12.resources.CustomerManagerLink
}
var file_google_ads_googleads_v12_services_customer_manager_link_service_proto_depIdxs = []int32{
	2, // 0: google.ads.googleads.v12.services.MutateCustomerManagerLinkRequest.operations:type_name -> google.ads.googleads.v12.services.CustomerManagerLinkOperation
	6, // 1: google.ads.googleads.v12.services.CustomerManagerLinkOperation.update_mask:type_name -> google.protobuf.FieldMask
	7, // 2: google.ads.googleads.v12.services.CustomerManagerLinkOperation.update:type_name -> google.ads.googleads.v12.resources.CustomerManagerLink
	5, // 3: google.ads.googleads.v12.services.MutateCustomerManagerLinkResponse.results:type_name -> google.ads.googleads.v12.services.MutateCustomerManagerLinkResult
	0, // 4: google.ads.googleads.v12.services.CustomerManagerLinkService.MutateCustomerManagerLink:input_type -> google.ads.googleads.v12.services.MutateCustomerManagerLinkRequest
	1, // 5: google.ads.googleads.v12.services.CustomerManagerLinkService.MoveManagerLink:input_type -> google.ads.googleads.v12.services.MoveManagerLinkRequest
	3, // 6: google.ads.googleads.v12.services.CustomerManagerLinkService.MutateCustomerManagerLink:output_type -> google.ads.googleads.v12.services.MutateCustomerManagerLinkResponse
	4, // 7: google.ads.googleads.v12.services.CustomerManagerLinkService.MoveManagerLink:output_type -> google.ads.googleads.v12.services.MoveManagerLinkResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_services_customer_manager_link_service_proto_init() }
func file_google_ads_googleads_v12_services_customer_manager_link_service_proto_init() {
	if File_google_ads_googleads_v12_services_customer_manager_link_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerManagerLinkRequest); i {
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
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveManagerLinkRequest); i {
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
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerManagerLinkOperation); i {
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
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerManagerLinkResponse); i {
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
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveManagerLinkResponse); i {
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
		file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateCustomerManagerLinkResult); i {
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
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CustomerManagerLinkOperation_Update)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v12_services_customer_manager_link_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_services_customer_manager_link_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v12_services_customer_manager_link_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_services_customer_manager_link_service_proto = out.File
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_rawDesc = nil
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_goTypes = nil
	file_google_ads_googleads_v12_services_customer_manager_link_service_proto_depIdxs = nil
}
