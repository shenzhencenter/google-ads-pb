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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v16/services/shared_criterion_service.proto

package services

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
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

// Request message for
// [SharedCriterionService.MutateSharedCriteria][google.ads.googleads.v16.services.SharedCriterionService.MutateSharedCriteria].
type MutateSharedCriteriaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose shared criteria are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual shared criteria.
	Operations []*SharedCriterionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v16.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateSharedCriteriaRequest) Reset() {
	*x = MutateSharedCriteriaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedCriteriaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedCriteriaRequest) ProtoMessage() {}

func (x *MutateSharedCriteriaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedCriteriaRequest.ProtoReflect.Descriptor instead.
func (*MutateSharedCriteriaRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescGZIP(), []int{0}
}

func (x *MutateSharedCriteriaRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateSharedCriteriaRequest) GetOperations() []*SharedCriterionOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateSharedCriteriaRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateSharedCriteriaRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateSharedCriteriaRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove) on an shared criterion.
type SharedCriterionOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//
	//	*SharedCriterionOperation_Create
	//	*SharedCriterionOperation_Remove
	Operation isSharedCriterionOperation_Operation `protobuf_oneof:"operation"`
}

func (x *SharedCriterionOperation) Reset() {
	*x = SharedCriterionOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedCriterionOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedCriterionOperation) ProtoMessage() {}

func (x *SharedCriterionOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedCriterionOperation.ProtoReflect.Descriptor instead.
func (*SharedCriterionOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescGZIP(), []int{1}
}

func (m *SharedCriterionOperation) GetOperation() isSharedCriterionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *SharedCriterionOperation) GetCreate() *resources.SharedCriterion {
	if x, ok := x.GetOperation().(*SharedCriterionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *SharedCriterionOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*SharedCriterionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isSharedCriterionOperation_Operation interface {
	isSharedCriterionOperation_Operation()
}

type SharedCriterionOperation_Create struct {
	// Create operation: No resource name is expected for the new shared
	// criterion.
	Create *resources.SharedCriterion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type SharedCriterionOperation_Remove struct {
	// Remove operation: A resource name for the removed shared criterion is
	// expected, in this format:
	//
	// `customers/{customer_id}/sharedCriteria/{shared_set_id}~{criterion_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*SharedCriterionOperation_Create) isSharedCriterionOperation_Operation() {}

func (*SharedCriterionOperation_Remove) isSharedCriterionOperation_Operation() {}

// Response message for a shared criterion mutate.
type MutateSharedCriteriaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (for example, auth
	// errors), we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateSharedCriterionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateSharedCriteriaResponse) Reset() {
	*x = MutateSharedCriteriaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedCriteriaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedCriteriaResponse) ProtoMessage() {}

func (x *MutateSharedCriteriaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedCriteriaResponse.ProtoReflect.Descriptor instead.
func (*MutateSharedCriteriaResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescGZIP(), []int{2}
}

func (x *MutateSharedCriteriaResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateSharedCriteriaResponse) GetResults() []*MutateSharedCriterionResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the shared criterion mutate.
type MutateSharedCriterionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated shared criterion with only mutable fields after mutate. The
	// field will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	SharedCriterion *resources.SharedCriterion `protobuf:"bytes,2,opt,name=shared_criterion,json=sharedCriterion,proto3" json:"shared_criterion,omitempty"`
}

func (x *MutateSharedCriterionResult) Reset() {
	*x = MutateSharedCriterionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedCriterionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedCriterionResult) ProtoMessage() {}

func (x *MutateSharedCriterionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedCriterionResult.ProtoReflect.Descriptor instead.
func (*MutateSharedCriterionResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateSharedCriterionResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateSharedCriterionResult) GetSharedCriterion() *resources.SharedCriterion {
	if x != nil {
		return x.SharedCriterion
	}
	return nil
}

var File_google_ads_googleads_v16_services_shared_criterion_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x02, 0x0a, 0x1b, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x60,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7f,
	0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0xbf, 0x01, 0x0a, 0x18, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xfa, 0x41, 0x2a,
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x1c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x58, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x1b, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xfa, 0x41, 0x2a,
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5e, 0x0a, 0x10, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x32, 0xd3, 0x02, 0x0a, 0x16, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xf1, 0x01, 0x0a, 0x14, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x3e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0xda,
	0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01,
	0x2a, 0x22, 0x34, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a,
	0x7d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x87,
	0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x36, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x1b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x36,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e,
	0x56, 0x31, 0x36, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x36, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x36, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescData = file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDesc
)

func file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDescData
}

var file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_ads_googleads_v16_services_shared_criterion_service_proto_goTypes = []interface{}{
	(*MutateSharedCriteriaRequest)(nil),                    // 0: google.ads.googleads.v16.services.MutateSharedCriteriaRequest
	(*SharedCriterionOperation)(nil),                       // 1: google.ads.googleads.v16.services.SharedCriterionOperation
	(*MutateSharedCriteriaResponse)(nil),                   // 2: google.ads.googleads.v16.services.MutateSharedCriteriaResponse
	(*MutateSharedCriterionResult)(nil),                    // 3: google.ads.googleads.v16.services.MutateSharedCriterionResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 4: google.ads.googleads.v16.enums.ResponseContentTypeEnum.ResponseContentType
	(*resources.SharedCriterion)(nil),                      // 5: google.ads.googleads.v16.resources.SharedCriterion
	(*status.Status)(nil),                                  // 6: google.rpc.Status
}
var file_google_ads_googleads_v16_services_shared_criterion_service_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v16.services.MutateSharedCriteriaRequest.operations:type_name -> google.ads.googleads.v16.services.SharedCriterionOperation
	4, // 1: google.ads.googleads.v16.services.MutateSharedCriteriaRequest.response_content_type:type_name -> google.ads.googleads.v16.enums.ResponseContentTypeEnum.ResponseContentType
	5, // 2: google.ads.googleads.v16.services.SharedCriterionOperation.create:type_name -> google.ads.googleads.v16.resources.SharedCriterion
	6, // 3: google.ads.googleads.v16.services.MutateSharedCriteriaResponse.partial_failure_error:type_name -> google.rpc.Status
	3, // 4: google.ads.googleads.v16.services.MutateSharedCriteriaResponse.results:type_name -> google.ads.googleads.v16.services.MutateSharedCriterionResult
	5, // 5: google.ads.googleads.v16.services.MutateSharedCriterionResult.shared_criterion:type_name -> google.ads.googleads.v16.resources.SharedCriterion
	0, // 6: google.ads.googleads.v16.services.SharedCriterionService.MutateSharedCriteria:input_type -> google.ads.googleads.v16.services.MutateSharedCriteriaRequest
	2, // 7: google.ads.googleads.v16.services.SharedCriterionService.MutateSharedCriteria:output_type -> google.ads.googleads.v16.services.MutateSharedCriteriaResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v16_services_shared_criterion_service_proto_init() }
func file_google_ads_googleads_v16_services_shared_criterion_service_proto_init() {
	if File_google_ads_googleads_v16_services_shared_criterion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedCriteriaRequest); i {
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
		file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedCriterionOperation); i {
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
		file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedCriteriaResponse); i {
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
		file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedCriterionResult); i {
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
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SharedCriterionOperation_Create)(nil),
		(*SharedCriterionOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v16_services_shared_criterion_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v16_services_shared_criterion_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v16_services_shared_criterion_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v16_services_shared_criterion_service_proto = out.File
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_rawDesc = nil
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_goTypes = nil
	file_google_ads_googleads_v16_services_shared_criterion_service_proto_depIdxs = nil
}
