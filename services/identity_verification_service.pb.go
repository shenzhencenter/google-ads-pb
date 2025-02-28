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
// source: google/ads/googleads/v19/services/identity_verification_service.proto

package services

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Request message for
//
//	[IdentityVerificationService.StartIdentityVerification].
type StartIdentityVerificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The Id of the customer for whom we are creating this
	// verification.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The verification program type for which we want to start the
	// verification.
	VerificationProgram enums.IdentityVerificationProgramEnum_IdentityVerificationProgram `protobuf:"varint,2,opt,name=verification_program,json=verificationProgram,proto3,enum=google.ads.googleads.v19.enums.IdentityVerificationProgramEnum_IdentityVerificationProgram" json:"verification_program,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *StartIdentityVerificationRequest) Reset() {
	*x = StartIdentityVerificationRequest{}
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartIdentityVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartIdentityVerificationRequest) ProtoMessage() {}

func (x *StartIdentityVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartIdentityVerificationRequest.ProtoReflect.Descriptor instead.
func (*StartIdentityVerificationRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP(), []int{0}
}

func (x *StartIdentityVerificationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *StartIdentityVerificationRequest) GetVerificationProgram() enums.IdentityVerificationProgramEnum_IdentityVerificationProgram {
	if x != nil {
		return x.VerificationProgram
	}
	return enums.IdentityVerificationProgramEnum_IdentityVerificationProgram(0)
}

// Request message for
//
//	[IdentityVerificationService.GetIdentityVerification].
type GetIdentityVerificationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required.  The ID of the customer for whom we are requesting verification
	//
	//	information.
	CustomerId    string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetIdentityVerificationRequest) Reset() {
	*x = GetIdentityVerificationRequest{}
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdentityVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdentityVerificationRequest) ProtoMessage() {}

func (x *GetIdentityVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdentityVerificationRequest.ProtoReflect.Descriptor instead.
func (*GetIdentityVerificationRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetIdentityVerificationRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

// Response message for
//
//	[IdentityVerificationService.GetIdentityVerification].
type GetIdentityVerificationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of identity verifications for the customer.
	IdentityVerification []*IdentityVerification `protobuf:"bytes,1,rep,name=identity_verification,json=identityVerification,proto3" json:"identity_verification,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *GetIdentityVerificationResponse) Reset() {
	*x = GetIdentityVerificationResponse{}
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetIdentityVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdentityVerificationResponse) ProtoMessage() {}

func (x *GetIdentityVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdentityVerificationResponse.ProtoReflect.Descriptor instead.
func (*GetIdentityVerificationResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetIdentityVerificationResponse) GetIdentityVerification() []*IdentityVerification {
	if x != nil {
		return x.IdentityVerification
	}
	return nil
}

// An identity verification for a customer.
type IdentityVerification struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The verification program type.
	VerificationProgram enums.IdentityVerificationProgramEnum_IdentityVerificationProgram `protobuf:"varint,1,opt,name=verification_program,json=verificationProgram,proto3,enum=google.ads.googleads.v19.enums.IdentityVerificationProgramEnum_IdentityVerificationProgram" json:"verification_program,omitempty"`
	// The verification requirement for this verification program for this
	// customer.
	IdentityVerificationRequirement *IdentityVerificationRequirement `protobuf:"bytes,2,opt,name=identity_verification_requirement,json=identityVerificationRequirement,proto3,oneof" json:"identity_verification_requirement,omitempty"`
	// Information regarding progress for this verification program for this
	// customer.
	VerificationProgress *IdentityVerificationProgress `protobuf:"bytes,3,opt,name=verification_progress,json=verificationProgress,proto3,oneof" json:"verification_progress,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *IdentityVerification) Reset() {
	*x = IdentityVerification{}
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityVerification) ProtoMessage() {}

func (x *IdentityVerification) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityVerification.ProtoReflect.Descriptor instead.
func (*IdentityVerification) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP(), []int{3}
}

func (x *IdentityVerification) GetVerificationProgram() enums.IdentityVerificationProgramEnum_IdentityVerificationProgram {
	if x != nil {
		return x.VerificationProgram
	}
	return enums.IdentityVerificationProgramEnum_IdentityVerificationProgram(0)
}

func (x *IdentityVerification) GetIdentityVerificationRequirement() *IdentityVerificationRequirement {
	if x != nil {
		return x.IdentityVerificationRequirement
	}
	return nil
}

func (x *IdentityVerification) GetVerificationProgress() *IdentityVerificationProgress {
	if x != nil {
		return x.VerificationProgress
	}
	return nil
}

// Information regarding the verification progress for a verification program
// type.
type IdentityVerificationProgress struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current Status (PENDING_USER_ACTION, SUCCESS, FAILURE etc)
	ProgramStatus enums.IdentityVerificationProgramStatusEnum_IdentityVerificationProgramStatus `protobuf:"varint,1,opt,name=program_status,json=programStatus,proto3,enum=google.ads.googleads.v19.enums.IdentityVerificationProgramStatusEnum_IdentityVerificationProgramStatus" json:"program_status,omitempty"`
	// The timestamp when the action url will expire in "yyyy-MM-dd HH:mm:ss"
	// format.
	InvitationLinkExpirationTime string `protobuf:"bytes,2,opt,name=invitation_link_expiration_time,json=invitationLinkExpirationTime,proto3" json:"invitation_link_expiration_time,omitempty"`
	// Action URL for user to complete verification for the given verification
	// program type.
	ActionUrl     string `protobuf:"bytes,3,opt,name=action_url,json=actionUrl,proto3" json:"action_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IdentityVerificationProgress) Reset() {
	*x = IdentityVerificationProgress{}
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityVerificationProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityVerificationProgress) ProtoMessage() {}

func (x *IdentityVerificationProgress) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityVerificationProgress.ProtoReflect.Descriptor instead.
func (*IdentityVerificationProgress) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP(), []int{4}
}

func (x *IdentityVerificationProgress) GetProgramStatus() enums.IdentityVerificationProgramStatusEnum_IdentityVerificationProgramStatus {
	if x != nil {
		return x.ProgramStatus
	}
	return enums.IdentityVerificationProgramStatusEnum_IdentityVerificationProgramStatus(0)
}

func (x *IdentityVerificationProgress) GetInvitationLinkExpirationTime() string {
	if x != nil {
		return x.InvitationLinkExpirationTime
	}
	return ""
}

func (x *IdentityVerificationProgress) GetActionUrl() string {
	if x != nil {
		return x.ActionUrl
	}
	return ""
}

// Information regarding the verification requirement for a verification program
// type.
type IdentityVerificationRequirement struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The deadline to start verification in "yyyy-MM-dd HH:mm:ss" format.
	VerificationStartDeadlineTime string `protobuf:"bytes,1,opt,name=verification_start_deadline_time,json=verificationStartDeadlineTime,proto3" json:"verification_start_deadline_time,omitempty"`
	// The deadline to submit verification.
	VerificationCompletionDeadlineTime string `protobuf:"bytes,2,opt,name=verification_completion_deadline_time,json=verificationCompletionDeadlineTime,proto3" json:"verification_completion_deadline_time,omitempty"`
	unknownFields                      protoimpl.UnknownFields
	sizeCache                          protoimpl.SizeCache
}

func (x *IdentityVerificationRequirement) Reset() {
	*x = IdentityVerificationRequirement{}
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityVerificationRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityVerificationRequirement) ProtoMessage() {}

func (x *IdentityVerificationRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityVerificationRequirement.ProtoReflect.Descriptor instead.
func (*IdentityVerificationRequirement) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP(), []int{5}
}

func (x *IdentityVerificationRequirement) GetVerificationStartDeadlineTime() string {
	if x != nil {
		return x.VerificationStartDeadlineTime
	}
	return ""
}

func (x *IdentityVerificationRequirement) GetVerificationCompletionDeadlineTime() string {
	if x != nil {
		return x.VerificationCompletionDeadlineTime
	}
	return ""
}

var File_google_ads_googleads_v19_services_identity_verification_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDesc = string([]byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x42, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x49,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde,
	0x01, 0x0a, 0x20, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x93, 0x01, 0x0a, 0x14, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x13, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x22,
	0x46, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x15, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf8, 0x03, 0x0a, 0x14, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x8e, 0x01, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x5b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x13,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x12, 0x93, 0x01, 0x0a, 0x21, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x1f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x79, 0x0a, 0x15, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x48, 0x01, 0x52, 0x14, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x88, 0x01, 0x01, 0x42, 0x24, 0x0a, 0x22, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x95, 0x02, 0x0a, 0x1c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x67,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x45, 0x0a, 0x1f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1c, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0xbd, 0x01, 0x0a,
	0x1f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x47, 0x0a, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x65, 0x61,
	0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x25, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x22, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xb8, 0x04, 0x0a,
	0x1b, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe0, 0x01, 0x0a,
	0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x66, 0xda, 0x41, 0x20, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x3d, 0x3a, 0x01, 0x2a, 0x22, 0x38, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x3a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0xee, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4c, 0xda, 0x41, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12, 0x36, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x67, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x45, 0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x64, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x8c, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x42, 0x20, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x39, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x31, 0x39, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02,
	0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x39, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescData []byte
)

func file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDesc), len(file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDesc)))
	})
	return file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDescData
}

var file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_ads_googleads_v19_services_identity_verification_service_proto_goTypes = []any{
	(*StartIdentityVerificationRequest)(nil),                                           // 0: google.ads.googleads.v19.services.StartIdentityVerificationRequest
	(*GetIdentityVerificationRequest)(nil),                                             // 1: google.ads.googleads.v19.services.GetIdentityVerificationRequest
	(*GetIdentityVerificationResponse)(nil),                                            // 2: google.ads.googleads.v19.services.GetIdentityVerificationResponse
	(*IdentityVerification)(nil),                                                       // 3: google.ads.googleads.v19.services.IdentityVerification
	(*IdentityVerificationProgress)(nil),                                               // 4: google.ads.googleads.v19.services.IdentityVerificationProgress
	(*IdentityVerificationRequirement)(nil),                                            // 5: google.ads.googleads.v19.services.IdentityVerificationRequirement
	(enums.IdentityVerificationProgramEnum_IdentityVerificationProgram)(0),             // 6: google.ads.googleads.v19.enums.IdentityVerificationProgramEnum.IdentityVerificationProgram
	(enums.IdentityVerificationProgramStatusEnum_IdentityVerificationProgramStatus)(0), // 7: google.ads.googleads.v19.enums.IdentityVerificationProgramStatusEnum.IdentityVerificationProgramStatus
	(*emptypb.Empty)(nil),                                                              // 8: google.protobuf.Empty
}
var file_google_ads_googleads_v19_services_identity_verification_service_proto_depIdxs = []int32{
	6, // 0: google.ads.googleads.v19.services.StartIdentityVerificationRequest.verification_program:type_name -> google.ads.googleads.v19.enums.IdentityVerificationProgramEnum.IdentityVerificationProgram
	3, // 1: google.ads.googleads.v19.services.GetIdentityVerificationResponse.identity_verification:type_name -> google.ads.googleads.v19.services.IdentityVerification
	6, // 2: google.ads.googleads.v19.services.IdentityVerification.verification_program:type_name -> google.ads.googleads.v19.enums.IdentityVerificationProgramEnum.IdentityVerificationProgram
	5, // 3: google.ads.googleads.v19.services.IdentityVerification.identity_verification_requirement:type_name -> google.ads.googleads.v19.services.IdentityVerificationRequirement
	4, // 4: google.ads.googleads.v19.services.IdentityVerification.verification_progress:type_name -> google.ads.googleads.v19.services.IdentityVerificationProgress
	7, // 5: google.ads.googleads.v19.services.IdentityVerificationProgress.program_status:type_name -> google.ads.googleads.v19.enums.IdentityVerificationProgramStatusEnum.IdentityVerificationProgramStatus
	0, // 6: google.ads.googleads.v19.services.IdentityVerificationService.StartIdentityVerification:input_type -> google.ads.googleads.v19.services.StartIdentityVerificationRequest
	1, // 7: google.ads.googleads.v19.services.IdentityVerificationService.GetIdentityVerification:input_type -> google.ads.googleads.v19.services.GetIdentityVerificationRequest
	8, // 8: google.ads.googleads.v19.services.IdentityVerificationService.StartIdentityVerification:output_type -> google.protobuf.Empty
	2, // 9: google.ads.googleads.v19.services.IdentityVerificationService.GetIdentityVerification:output_type -> google.ads.googleads.v19.services.GetIdentityVerificationResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v19_services_identity_verification_service_proto_init() }
func file_google_ads_googleads_v19_services_identity_verification_service_proto_init() {
	if File_google_ads_googleads_v19_services_identity_verification_service_proto != nil {
		return
	}
	file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDesc), len(file_google_ads_googleads_v19_services_identity_verification_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v19_services_identity_verification_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v19_services_identity_verification_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v19_services_identity_verification_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v19_services_identity_verification_service_proto = out.File
	file_google_ads_googleads_v19_services_identity_verification_service_proto_goTypes = nil
	file_google_ads_googleads_v19_services_identity_verification_service_proto_depIdxs = nil
}
