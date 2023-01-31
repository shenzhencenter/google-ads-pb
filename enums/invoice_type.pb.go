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
// 	protoc        v3.21.9
// source: google/ads/googleads/v12/enums/invoice_type.proto

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

// The possible type of invoices.
type InvoiceTypeEnum_InvoiceType int32

const (
	// Not specified.
	InvoiceTypeEnum_UNSPECIFIED InvoiceTypeEnum_InvoiceType = 0
	// Used for return value only. Represents value unknown in this version.
	InvoiceTypeEnum_UNKNOWN InvoiceTypeEnum_InvoiceType = 1
	// An invoice with a negative amount. The account receives a credit.
	InvoiceTypeEnum_CREDIT_MEMO InvoiceTypeEnum_InvoiceType = 2
	// An invoice with a positive amount. The account owes a balance.
	InvoiceTypeEnum_INVOICE InvoiceTypeEnum_InvoiceType = 3
)

// Enum value maps for InvoiceTypeEnum_InvoiceType.
var (
	InvoiceTypeEnum_InvoiceType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CREDIT_MEMO",
		3: "INVOICE",
	}
	InvoiceTypeEnum_InvoiceType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"CREDIT_MEMO": 2,
		"INVOICE":     3,
	}
)

func (x InvoiceTypeEnum_InvoiceType) Enum() *InvoiceTypeEnum_InvoiceType {
	p := new(InvoiceTypeEnum_InvoiceType)
	*p = x
	return p
}

func (x InvoiceTypeEnum_InvoiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvoiceTypeEnum_InvoiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v12_enums_invoice_type_proto_enumTypes[0].Descriptor()
}

func (InvoiceTypeEnum_InvoiceType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v12_enums_invoice_type_proto_enumTypes[0]
}

func (x InvoiceTypeEnum_InvoiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceTypeEnum_InvoiceType.Descriptor instead.
func (InvoiceTypeEnum_InvoiceType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the type of invoices.
type InvoiceTypeEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InvoiceTypeEnum) Reset() {
	*x = InvoiceTypeEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v12_enums_invoice_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceTypeEnum) ProtoMessage() {}

func (x *InvoiceTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v12_enums_invoice_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceTypeEnum.ProtoReflect.Descriptor instead.
func (*InvoiceTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v12_enums_invoice_type_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v12_enums_invoice_type_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x32, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x22, 0x5c, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x49, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x5f, 0x4d, 0x45,
	0x4d, 0x4f, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x10,
	0x03, 0x42, 0xea, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x42, 0x10, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x31, 0x32, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x31, 0x32, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x32, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescData = file_google_ads_googleads_v12_enums_invoice_type_proto_rawDesc
)

func file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescData)
	})
	return file_google_ads_googleads_v12_enums_invoice_type_proto_rawDescData
}

var file_google_ads_googleads_v12_enums_invoice_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v12_enums_invoice_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v12_enums_invoice_type_proto_goTypes = []interface{}{
	(InvoiceTypeEnum_InvoiceType)(0), // 0: google.ads.googleads.v12.enums.InvoiceTypeEnum.InvoiceType
	(*InvoiceTypeEnum)(nil),          // 1: google.ads.googleads.v12.enums.InvoiceTypeEnum
}
var file_google_ads_googleads_v12_enums_invoice_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v12_enums_invoice_type_proto_init() }
func file_google_ads_googleads_v12_enums_invoice_type_proto_init() {
	if File_google_ads_googleads_v12_enums_invoice_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v12_enums_invoice_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceTypeEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v12_enums_invoice_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v12_enums_invoice_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v12_enums_invoice_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v12_enums_invoice_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v12_enums_invoice_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v12_enums_invoice_type_proto = out.File
	file_google_ads_googleads_v12_enums_invoice_type_proto_rawDesc = nil
	file_google_ads_googleads_v12_enums_invoice_type_proto_goTypes = nil
	file_google_ads_googleads_v12_enums_invoice_type_proto_depIdxs = nil
}
