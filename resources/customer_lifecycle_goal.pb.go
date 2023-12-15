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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: google/ads/googleads/v15/resources/customer_lifecycle_goal.proto

package resources

import (
	common "github.com/shenzhencenter/google-ads-pb/common"
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

// Account level customer lifecycle goal settings.
type CustomerLifecycleGoal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the customer lifecycle goal.
	// Customer lifecycle resource names have the form:
	//
	// `customers/{customer_id}/customerLifecycleGoal`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Common lifecycle goal settings shared among different types of
	// lifecycle goals.
	LifecycleGoalCustomerDefinitionSettings *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings `protobuf:"bytes,2,opt,name=lifecycle_goal_customer_definition_settings,json=lifecycleGoalCustomerDefinitionSettings,proto3" json:"lifecycle_goal_customer_definition_settings,omitempty"`
	// Output only. Customer acquisition goal customer level value settings.
	CustomerAcquisitionGoalValueSettings *common.LifecycleGoalValueSettings `protobuf:"bytes,3,opt,name=customer_acquisition_goal_value_settings,json=customerAcquisitionGoalValueSettings,proto3" json:"customer_acquisition_goal_value_settings,omitempty"`
}

func (x *CustomerLifecycleGoal) Reset() {
	*x = CustomerLifecycleGoal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerLifecycleGoal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerLifecycleGoal) ProtoMessage() {}

func (x *CustomerLifecycleGoal) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerLifecycleGoal.ProtoReflect.Descriptor instead.
func (*CustomerLifecycleGoal) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerLifecycleGoal) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CustomerLifecycleGoal) GetLifecycleGoalCustomerDefinitionSettings() *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings {
	if x != nil {
		return x.LifecycleGoalCustomerDefinitionSettings
	}
	return nil
}

func (x *CustomerLifecycleGoal) GetCustomerAcquisitionGoalValueSettings() *common.LifecycleGoalValueSettings {
	if x != nil {
		return x.CustomerAcquisitionGoalValueSettings
	}
	return nil
}

// Lifecycle goal common settings, including existing user lists and existing
// high lifetime value user lists, shared among different types of lifecycle
// goals.
type CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. User lists which represent existing customers.
	ExistingUserLists []string `protobuf:"bytes,1,rep,name=existing_user_lists,json=existingUserLists,proto3" json:"existing_user_lists,omitempty"`
	// Output only. User lists which represent customers of high lifetime value.
	// In current stage, high lifetime value feature is in beta and this field
	// is read-only.
	HighLifetimeValueUserLists []string `protobuf:"bytes,2,rep,name=high_lifetime_value_user_lists,json=highLifetimeValueUserLists,proto3" json:"high_lifetime_value_user_lists,omitempty"`
}

func (x *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) Reset() {
	*x = CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) ProtoMessage() {}

func (x *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings.ProtoReflect.Descriptor instead.
func (*CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) GetExistingUserLists() []string {
	if x != nil {
		return x.ExistingUserLists
	}
	return nil
}

func (x *CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings) GetHighLifetimeValueUserLists() []string {
	if x != nil {
		return x.HighLifetimeValueUserLists
	}
	return nil
}

var File_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x06, 0x0a, 0x15, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47,
	0x6f, 0x61, 0x6c, 0x12, 0x5b, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xe0, 0x41, 0x05, 0xfa,
	0x41, 0x30, 0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f,
	0x61, 0x6c, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0xc4, 0x01, 0x0a, 0x2b, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x67,
	0x6f, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x61, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61,
	0x6c, 0x2e, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x27,
	0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x28, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x24, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x6f, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x1a, 0xf3, 0x01, 0x0a, 0x27, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x47, 0x6f, 0x61, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x59,
	0x0a, 0x13, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03,
	0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x11, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x6d, 0x0a, 0x1e, 0x68, 0x69, 0x67,
	0x68, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x1a, 0x68, 0x69,
	0x67, 0x68, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x3a, 0x63, 0xea, 0x41, 0x60, 0x0a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x12, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x42, 0x8c, 0x02,
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x47, 0x6f, 0x61, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x35, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x35, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x35, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescData = file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDesc
)

func file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescData)
	})
	return file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDescData
}

var file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_goTypes = []interface{}{
	(*CustomerLifecycleGoal)(nil),                                         // 0: google.ads.googleads.v15.resources.CustomerLifecycleGoal
	(*CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings)(nil), // 1: google.ads.googleads.v15.resources.CustomerLifecycleGoal.LifecycleGoalCustomerDefinitionSettings
	(*common.LifecycleGoalValueSettings)(nil),                             // 2: google.ads.googleads.v15.common.LifecycleGoalValueSettings
}
var file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v15.resources.CustomerLifecycleGoal.lifecycle_goal_customer_definition_settings:type_name -> google.ads.googleads.v15.resources.CustomerLifecycleGoal.LifecycleGoalCustomerDefinitionSettings
	2, // 1: google.ads.googleads.v15.resources.CustomerLifecycleGoal.customer_acquisition_goal_value_settings:type_name -> google.ads.googleads.v15.common.LifecycleGoalValueSettings
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_init() }
func file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_init() {
	if File_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerLifecycleGoal); i {
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
		file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerLifecycleGoal_LifecycleGoalCustomerDefinitionSettings); i {
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
			RawDescriptor: file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto = out.File
	file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_rawDesc = nil
	file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_goTypes = nil
	file_google_ads_googleads_v15_resources_customer_lifecycle_goal_proto_depIdxs = nil
}
