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
// source: google/ads/googleads/v17/resources/experiment_arm.proto

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

// A Google ads experiment for users to experiment changes on multiple
// campaigns, compare the performance, and apply the effective changes.
type ExperimentArm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the experiment arm.
	// Experiment arm resource names have the form:
	//
	// `customers/{customer_id}/experimentArms/{TrialArm.trial_id}~{TrialArm.trial_arm_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The experiment to which the ExperimentArm belongs.
	Experiment string `protobuf:"bytes,8,opt,name=experiment,proto3" json:"experiment,omitempty"`
	// Required. The name of the experiment arm. It must have a minimum length of
	// 1 and maximum length of 1024. It must be unique under an experiment.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether this arm is a control arm. A control arm is the arm against
	// which the other arms are compared.
	Control bool `protobuf:"varint,4,opt,name=control,proto3" json:"control,omitempty"`
	// Traffic split of the trial arm. The value should be between 1 and 100
	// and must total 100 between the two trial arms.
	TrafficSplit int64 `protobuf:"varint,5,opt,name=traffic_split,json=trafficSplit,proto3" json:"traffic_split,omitempty"`
	// List of campaigns in the trial arm. The max length is one.
	Campaigns []string `protobuf:"bytes,6,rep,name=campaigns,proto3" json:"campaigns,omitempty"`
	// Output only. The in design campaigns in the treatment experiment arm.
	InDesignCampaigns []string `protobuf:"bytes,7,rep,name=in_design_campaigns,json=inDesignCampaigns,proto3" json:"in_design_campaigns,omitempty"`
}

func (x *ExperimentArm) Reset() {
	*x = ExperimentArm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_resources_experiment_arm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentArm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentArm) ProtoMessage() {}

func (x *ExperimentArm) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_resources_experiment_arm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentArm.ProtoReflect.Descriptor instead.
func (*ExperimentArm) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescGZIP(), []int{0}
}

func (x *ExperimentArm) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *ExperimentArm) GetExperiment() string {
	if x != nil {
		return x.Experiment
	}
	return ""
}

func (x *ExperimentArm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExperimentArm) GetControl() bool {
	if x != nil {
		return x.Control
	}
	return false
}

func (x *ExperimentArm) GetTrafficSplit() int64 {
	if x != nil {
		return x.TrafficSplit
	}
	return 0
}

func (x *ExperimentArm) GetCampaigns() []string {
	if x != nil {
		return x.Campaigns
	}
	return nil
}

func (x *ExperimentArm) GetInDesignCampaigns() []string {
	if x != nil {
		return x.InDesignCampaigns
	}
	return nil
}

var File_google_ads_googleads_v17_resources_experiment_arm_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a, 0x0d, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x6d, 0x12, 0x53, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x2e, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x28, 0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41,
	0x72, 0x6d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4b, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x25, 0x0a, 0x23, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x73, 0x70, 0x6c, 0x69,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x44, 0x0a, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x52, 0x09, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x12, 0x59, 0x0a, 0x13, 0x69,
	0x6e, 0x5f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x29, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x23,
	0x0a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61,
	0x69, 0x67, 0x6e, 0x52, 0x11, 0x69, 0x6e, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x73, 0x3a, 0x6d, 0xea, 0x41, 0x6a, 0x0a, 0x26, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x72, 0x6d, 0x12, 0x40, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x6d, 0x73, 0x2f, 0x7b, 0x74, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x7e, 0x7b, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x72,
	0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x84, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x42, 0x12, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x6d, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73,
	0x2e, 0x56, 0x31, 0x37, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02,
	0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0xea, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x37, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescData = file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDesc
)

func file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDescData
}

var file_google_ads_googleads_v17_resources_experiment_arm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_resources_experiment_arm_proto_goTypes = []any{
	(*ExperimentArm)(nil), // 0: google.ads.googleads.v17.resources.ExperimentArm
}
var file_google_ads_googleads_v17_resources_experiment_arm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_resources_experiment_arm_proto_init() }
func file_google_ads_googleads_v17_resources_experiment_arm_proto_init() {
	if File_google_ads_googleads_v17_resources_experiment_arm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_resources_experiment_arm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExperimentArm); i {
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
			RawDescriptor: file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_resources_experiment_arm_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_resources_experiment_arm_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v17_resources_experiment_arm_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_resources_experiment_arm_proto = out.File
	file_google_ads_googleads_v17_resources_experiment_arm_proto_rawDesc = nil
	file_google_ads_googleads_v17_resources_experiment_arm_proto_goTypes = nil
	file_google_ads_googleads_v17_resources_experiment_arm_proto_depIdxs = nil
}
