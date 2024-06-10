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
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: google/ads/googleads/v17/enums/simulation_modification_method.proto

package enums

import (
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

// Enum describing the method by which a simulation modifies a field.
type SimulationModificationMethodEnum_SimulationModificationMethod int32

const (
	// Not specified.
	SimulationModificationMethodEnum_UNSPECIFIED SimulationModificationMethodEnum_SimulationModificationMethod = 0
	// Used for return value only. Represents value unknown in this version.
	SimulationModificationMethodEnum_UNKNOWN SimulationModificationMethodEnum_SimulationModificationMethod = 1
	// The values in a simulation were applied to all children of a given
	// resource uniformly. Overrides on child resources were not respected.
	SimulationModificationMethodEnum_UNIFORM SimulationModificationMethodEnum_SimulationModificationMethod = 2
	// The values in a simulation were applied to the given resource.
	// Overrides on child resources were respected, and traffic estimates
	// do not include these resources.
	SimulationModificationMethodEnum_DEFAULT SimulationModificationMethodEnum_SimulationModificationMethod = 3
	// The values in a simulation were all scaled by the same factor.
	// For example, in a simulated TargetCpa campaign, the campaign target and
	// all ad group targets were scaled by a factor of X.
	SimulationModificationMethodEnum_SCALING SimulationModificationMethodEnum_SimulationModificationMethod = 4
)

// Enum value maps for SimulationModificationMethodEnum_SimulationModificationMethod.
var (
	SimulationModificationMethodEnum_SimulationModificationMethod_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "UNIFORM",
		3: "DEFAULT",
		4: "SCALING",
	}
	SimulationModificationMethodEnum_SimulationModificationMethod_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"UNIFORM":     2,
		"DEFAULT":     3,
		"SCALING":     4,
	}
)

func (x SimulationModificationMethodEnum_SimulationModificationMethod) Enum() *SimulationModificationMethodEnum_SimulationModificationMethod {
	p := new(SimulationModificationMethodEnum_SimulationModificationMethod)
	*p = x
	return p
}

func (x SimulationModificationMethodEnum_SimulationModificationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SimulationModificationMethodEnum_SimulationModificationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_enums_simulation_modification_method_proto_enumTypes[0].Descriptor()
}

func (SimulationModificationMethodEnum_SimulationModificationMethod) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_enums_simulation_modification_method_proto_enumTypes[0]
}

func (x SimulationModificationMethodEnum_SimulationModificationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SimulationModificationMethodEnum_SimulationModificationMethod.Descriptor instead.
func (SimulationModificationMethodEnum_SimulationModificationMethod) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the method by which a simulation modifies
// a field.
type SimulationModificationMethodEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SimulationModificationMethodEnum) Reset() {
	*x = SimulationModificationMethodEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_enums_simulation_modification_method_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulationModificationMethodEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulationModificationMethodEnum) ProtoMessage() {}

func (x *SimulationModificationMethodEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_enums_simulation_modification_method_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulationModificationMethodEnum.ProtoReflect.Descriptor instead.
func (*SimulationModificationMethodEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_enums_simulation_modification_method_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x20, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x63, 0x0a, 0x1c, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x49, 0x46,
	0x4f, 0x52, 0x4d, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x43, 0x41, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42,
	0xfb, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x21, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0x37, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56,
	0x31, 0x37, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescData = file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDesc
)

func file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDescData
}

var file_google_ads_googleads_v17_enums_simulation_modification_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_enums_simulation_modification_method_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_enums_simulation_modification_method_proto_goTypes = []interface{}{
	(SimulationModificationMethodEnum_SimulationModificationMethod)(0), // 0: google.ads.googleads.v17.enums.SimulationModificationMethodEnum.SimulationModificationMethod
	(*SimulationModificationMethodEnum)(nil),                           // 1: google.ads.googleads.v17.enums.SimulationModificationMethodEnum
}
var file_google_ads_googleads_v17_enums_simulation_modification_method_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_enums_simulation_modification_method_proto_init() }
func file_google_ads_googleads_v17_enums_simulation_modification_method_proto_init() {
	if File_google_ads_googleads_v17_enums_simulation_modification_method_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_enums_simulation_modification_method_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulationModificationMethodEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_enums_simulation_modification_method_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_enums_simulation_modification_method_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_enums_simulation_modification_method_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_enums_simulation_modification_method_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_enums_simulation_modification_method_proto = out.File
	file_google_ads_googleads_v17_enums_simulation_modification_method_proto_rawDesc = nil
	file_google_ads_googleads_v17_enums_simulation_modification_method_proto_goTypes = nil
	file_google_ads_googleads_v17_enums_simulation_modification_method_proto_depIdxs = nil
}
