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
// source: google/ads/googleads/v17/errors/account_budget_proposal_error.proto

package errors

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

// Enum describing possible account budget proposal errors.
type AccountBudgetProposalErrorEnum_AccountBudgetProposalError int32

const (
	// Enum unspecified.
	AccountBudgetProposalErrorEnum_UNSPECIFIED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 0
	// The received error code is not known in this version.
	AccountBudgetProposalErrorEnum_UNKNOWN AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 1
	// The field mask must be empty for create/end/remove proposals.
	AccountBudgetProposalErrorEnum_FIELD_MASK_NOT_ALLOWED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 2
	// The field cannot be set because of the proposal type.
	AccountBudgetProposalErrorEnum_IMMUTABLE_FIELD AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 3
	// The field is required because of the proposal type.
	AccountBudgetProposalErrorEnum_REQUIRED_FIELD_MISSING AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 4
	// Proposals that have been approved cannot be cancelled.
	AccountBudgetProposalErrorEnum_CANNOT_CANCEL_APPROVED_PROPOSAL AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 5
	// Budgets that haven't been approved cannot be removed.
	AccountBudgetProposalErrorEnum_CANNOT_REMOVE_UNAPPROVED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 6
	// Budgets that are currently running cannot be removed.
	AccountBudgetProposalErrorEnum_CANNOT_REMOVE_RUNNING_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 7
	// Budgets that haven't been approved cannot be truncated.
	AccountBudgetProposalErrorEnum_CANNOT_END_UNAPPROVED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 8
	// Only budgets that are currently running can be truncated.
	AccountBudgetProposalErrorEnum_CANNOT_END_INACTIVE_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 9
	// All budgets must have names.
	AccountBudgetProposalErrorEnum_BUDGET_NAME_REQUIRED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 10
	// Expired budgets cannot be edited after a sufficient amount of time has
	// passed.
	AccountBudgetProposalErrorEnum_CANNOT_UPDATE_OLD_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 11
	// It is not permissible a propose a new budget that ends in the past.
	AccountBudgetProposalErrorEnum_CANNOT_END_IN_PAST AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 12
	// An expired budget cannot be extended to overlap with the running budget.
	AccountBudgetProposalErrorEnum_CANNOT_EXTEND_END_TIME AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 13
	// A purchase order number is required.
	AccountBudgetProposalErrorEnum_PURCHASE_ORDER_NUMBER_REQUIRED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 14
	// Budgets that have a pending update cannot be updated.
	AccountBudgetProposalErrorEnum_PENDING_UPDATE_PROPOSAL_EXISTS AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 15
	// Cannot propose more than one budget when the corresponding billing setup
	// hasn't been approved.
	AccountBudgetProposalErrorEnum_MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 16
	// Cannot update the start time of a budget that has already started.
	AccountBudgetProposalErrorEnum_CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 17
	// Cannot update the spending limit of a budget with an amount lower than
	// what has already been spent.
	AccountBudgetProposalErrorEnum_SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 18
	// Cannot propose a budget update without actually changing any fields.
	AccountBudgetProposalErrorEnum_UPDATE_IS_NO_OP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 19
	// The end time must come after the start time.
	AccountBudgetProposalErrorEnum_END_TIME_MUST_FOLLOW_START_TIME AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 20
	// The budget's date range must fall within the date range of its billing
	// setup.
	AccountBudgetProposalErrorEnum_BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 21
	// The user is not authorized to mutate budgets for the given billing setup.
	AccountBudgetProposalErrorEnum_NOT_AUTHORIZED AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 22
	// Mutates are not allowed for the given billing setup.
	AccountBudgetProposalErrorEnum_INVALID_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 23
	// Budget creation failed as it overlaps with a pending budget proposal
	// or an approved budget.
	AccountBudgetProposalErrorEnum_OVERLAPS_EXISTING_BUDGET AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 24
	// The control setting in user's payments profile doesn't allow budget
	// creation through API. Log in to Google Ads to create budget.
	AccountBudgetProposalErrorEnum_CANNOT_CREATE_BUDGET_THROUGH_API AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 25
	// Master service agreement has not been signed yet for the Payments
	// Profile.
	AccountBudgetProposalErrorEnum_INVALID_MASTER_SERVICE_AGREEMENT AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 26
	// Budget mutates are not allowed because the given billing setup is
	// canceled.
	AccountBudgetProposalErrorEnum_CANCELED_BILLING_SETUP AccountBudgetProposalErrorEnum_AccountBudgetProposalError = 27
)

// Enum value maps for AccountBudgetProposalErrorEnum_AccountBudgetProposalError.
var (
	AccountBudgetProposalErrorEnum_AccountBudgetProposalError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "FIELD_MASK_NOT_ALLOWED",
		3:  "IMMUTABLE_FIELD",
		4:  "REQUIRED_FIELD_MISSING",
		5:  "CANNOT_CANCEL_APPROVED_PROPOSAL",
		6:  "CANNOT_REMOVE_UNAPPROVED_BUDGET",
		7:  "CANNOT_REMOVE_RUNNING_BUDGET",
		8:  "CANNOT_END_UNAPPROVED_BUDGET",
		9:  "CANNOT_END_INACTIVE_BUDGET",
		10: "BUDGET_NAME_REQUIRED",
		11: "CANNOT_UPDATE_OLD_BUDGET",
		12: "CANNOT_END_IN_PAST",
		13: "CANNOT_EXTEND_END_TIME",
		14: "PURCHASE_ORDER_NUMBER_REQUIRED",
		15: "PENDING_UPDATE_PROPOSAL_EXISTS",
		16: "MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP",
		17: "CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET",
		18: "SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED",
		19: "UPDATE_IS_NO_OP",
		20: "END_TIME_MUST_FOLLOW_START_TIME",
		21: "BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP",
		22: "NOT_AUTHORIZED",
		23: "INVALID_BILLING_SETUP",
		24: "OVERLAPS_EXISTING_BUDGET",
		25: "CANNOT_CREATE_BUDGET_THROUGH_API",
		26: "INVALID_MASTER_SERVICE_AGREEMENT",
		27: "CANCELED_BILLING_SETUP",
	}
	AccountBudgetProposalErrorEnum_AccountBudgetProposalError_value = map[string]int32{
		"UNSPECIFIED":                     0,
		"UNKNOWN":                         1,
		"FIELD_MASK_NOT_ALLOWED":          2,
		"IMMUTABLE_FIELD":                 3,
		"REQUIRED_FIELD_MISSING":          4,
		"CANNOT_CANCEL_APPROVED_PROPOSAL": 5,
		"CANNOT_REMOVE_UNAPPROVED_BUDGET": 6,
		"CANNOT_REMOVE_RUNNING_BUDGET":    7,
		"CANNOT_END_UNAPPROVED_BUDGET":    8,
		"CANNOT_END_INACTIVE_BUDGET":      9,
		"BUDGET_NAME_REQUIRED":            10,
		"CANNOT_UPDATE_OLD_BUDGET":        11,
		"CANNOT_END_IN_PAST":              12,
		"CANNOT_EXTEND_END_TIME":          13,
		"PURCHASE_ORDER_NUMBER_REQUIRED":  14,
		"PENDING_UPDATE_PROPOSAL_EXISTS":  15,
		"MULTIPLE_BUDGETS_NOT_ALLOWED_FOR_UNAPPROVED_BILLING_SETUP": 16,
		"CANNOT_UPDATE_START_TIME_FOR_STARTED_BUDGET":               17,
		"SPENDING_LIMIT_LOWER_THAN_ACCRUED_COST_NOT_ALLOWED":        18,
		"UPDATE_IS_NO_OP":                                   19,
		"END_TIME_MUST_FOLLOW_START_TIME":                   20,
		"BUDGET_DATE_RANGE_INCOMPATIBLE_WITH_BILLING_SETUP": 21,
		"NOT_AUTHORIZED":                                    22,
		"INVALID_BILLING_SETUP":                             23,
		"OVERLAPS_EXISTING_BUDGET":                          24,
		"CANNOT_CREATE_BUDGET_THROUGH_API":                  25,
		"INVALID_MASTER_SERVICE_AGREEMENT":                  26,
		"CANCELED_BILLING_SETUP":                            27,
	}
)

func (x AccountBudgetProposalErrorEnum_AccountBudgetProposalError) Enum() *AccountBudgetProposalErrorEnum_AccountBudgetProposalError {
	p := new(AccountBudgetProposalErrorEnum_AccountBudgetProposalError)
	*p = x
	return p
}

func (x AccountBudgetProposalErrorEnum_AccountBudgetProposalError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountBudgetProposalErrorEnum_AccountBudgetProposalError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_enumTypes[0].Descriptor()
}

func (AccountBudgetProposalErrorEnum_AccountBudgetProposalError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_enumTypes[0]
}

func (x AccountBudgetProposalErrorEnum_AccountBudgetProposalError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountBudgetProposalErrorEnum_AccountBudgetProposalError.Descriptor instead.
func (AccountBudgetProposalErrorEnum_AccountBudgetProposalError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible account budget proposal errors.
type AccountBudgetProposalErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AccountBudgetProposalErrorEnum) Reset() {
	*x = AccountBudgetProposalErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountBudgetProposalErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountBudgetProposalErrorEnum) ProtoMessage() {}

func (x *AccountBudgetProposalErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountBudgetProposalErrorEnum.ProtoReflect.Descriptor instead.
func (*AccountBudgetProposalErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v17_errors_account_budget_proposal_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x37, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0xda, 0x07, 0x0a, 0x1e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xb7, 0x07, 0x0a, 0x1a, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x4d, 0x41, 0x53, 0x4b, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x51, 0x55, 0x49,
	0x52, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x50, 0x52,
	0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x50, 0x50, 0x52,
	0x4f, 0x56, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x06, 0x12, 0x20, 0x0a,
	0x1c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x52,
	0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x07, 0x12,
	0x20, 0x0a, 0x1c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x55, 0x4e,
	0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10,
	0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x5f,
	0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10,
	0x09, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x4e, 0x41, 0x4d, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4c, 0x44,
	0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x41, 0x53, 0x54, 0x10,
	0x0c, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x54, 0x45,
	0x4e, 0x44, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x0d, 0x12, 0x22, 0x0a,
	0x1e, 0x50, 0x55, 0x52, 0x43, 0x48, 0x41, 0x53, 0x45, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x0e, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x53, 0x10, 0x0f, 0x12, 0x3d, 0x0a, 0x39, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c,
	0x45, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x41, 0x50, 0x50, 0x52,
	0x4f, 0x56, 0x45, 0x44, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54,
	0x55, 0x50, 0x10, 0x10, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45,
	0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x10, 0x11, 0x12, 0x36, 0x0a, 0x32, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x5f, 0x54, 0x48,
	0x41, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x52, 0x55, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x53, 0x54, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x12, 0x12, 0x13, 0x0a,
	0x0f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x4f, 0x5f, 0x4f, 0x50,
	0x10, 0x13, 0x12, 0x23, 0x0a, 0x1f, 0x45, 0x4e, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4d,
	0x55, 0x53, 0x54, 0x5f, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x14, 0x12, 0x35, 0x0a, 0x31, 0x42, 0x55, 0x44, 0x47, 0x45,
	0x54, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x42,
	0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x10, 0x15, 0x12, 0x12,
	0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44,
	0x10, 0x16, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x42, 0x49,
	0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x10, 0x17, 0x12, 0x1c, 0x0a,
	0x18, 0x4f, 0x56, 0x45, 0x52, 0x4c, 0x41, 0x50, 0x53, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x42, 0x55, 0x44, 0x47, 0x45, 0x54, 0x10, 0x18, 0x12, 0x24, 0x0a, 0x20, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x55, 0x44,
	0x47, 0x45, 0x54, 0x5f, 0x54, 0x48, 0x52, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x41, 0x50, 0x49, 0x10,
	0x19, 0x12, 0x24, 0x0a, 0x20, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x41, 0x53,
	0x54, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x47, 0x52, 0x45,
	0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x1a, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x45, 0x44, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x55,
	0x50, 0x10, 0x1b, 0x42, 0xff, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x37, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x37, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x31, 0x37, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x37, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea,
	0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x37, 0x3a, 0x3a, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescData = file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDesc
)

func file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDescData
}

var file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_goTypes = []interface{}{
	(AccountBudgetProposalErrorEnum_AccountBudgetProposalError)(0), // 0: google.ads.googleads.v17.errors.AccountBudgetProposalErrorEnum.AccountBudgetProposalError
	(*AccountBudgetProposalErrorEnum)(nil),                         // 1: google.ads.googleads.v17.errors.AccountBudgetProposalErrorEnum
}
var file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_init() }
func file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_init() {
	if File_google_ads_googleads_v17_errors_account_budget_proposal_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountBudgetProposalErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v17_errors_account_budget_proposal_error_proto = out.File
	file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_rawDesc = nil
	file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_goTypes = nil
	file_google_ads_googleads_v17_errors_account_budget_proposal_error_proto_depIdxs = nil
}
