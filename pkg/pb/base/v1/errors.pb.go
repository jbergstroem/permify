// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: errors.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type ErrorCode int32

const (
	ErrorCode_no_error ErrorCode = 0
	// authn
	ErrorCode_missing_bearer_token ErrorCode = 1001
	ErrorCode_unauthenticated      ErrorCode = 1002
	// validation
	ErrorCode_validation_error                                  ErrorCode = 2000
	ErrorCode_undefined_child_type                              ErrorCode = 2002
	ErrorCode_undefined_child_kind                              ErrorCode = 2003
	ErrorCode_undefined_relation_reference                      ErrorCode = 2006
	ErrorCode_not_supported_relation_walk                       ErrorCode = 2007
	ErrorCode_entity_and_subject_cannot_be_equal                ErrorCode = 2008
	ErrorCode_depth_not_enough                                  ErrorCode = 2009
	ErrorCode_relation_reference_not_found_in_entity_references ErrorCode = 2010
	ErrorCode_relation_reference_must_have_one_entity_reference ErrorCode = 2011
	ErrorCode_duplicated_entity_reference                       ErrorCode = 2012
	ErrorCode_duplicated_relation_reference                     ErrorCode = 2013
	ErrorCode_duplicated_action_reference                       ErrorCode = 2014
	ErrorCode_schema_parse                                      ErrorCode = 2015
	ErrorCode_schema_compile                                    ErrorCode = 2016
	// not found
	ErrorCode_not_found_error               ErrorCode = 4000
	ErrorCode_entity_type_can_not_found     ErrorCode = 4001
	ErrorCode_action_can_not_found          ErrorCode = 4002
	ErrorCode_schema_not_found              ErrorCode = 4003
	ErrorCode_subject_type_not_found        ErrorCode = 4004
	ErrorCode_entity_definition_not_found   ErrorCode = 4005
	ErrorCode_action_definition_not_found   ErrorCode = 4006
	ErrorCode_relation_definition_not_found ErrorCode = 4007
	ErrorCode_record_not_found              ErrorCode = 4008
	// internal
	ErrorCode_internal_error        ErrorCode = 5000
	ErrorCode_cancelled             ErrorCode = 5001
	ErrorCode_sql_builder_error     ErrorCode = 5002
	ErrorCode_circuit_breaker_error ErrorCode = 5003
	ErrorCode_unique_constraint     ErrorCode = 5004
	ErrorCode_execution             ErrorCode = 5005
	ErrorCode_scan                  ErrorCode = 5006
	ErrorCode_migration             ErrorCode = 5007
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "no_error",
		1001: "missing_bearer_token",
		1002: "unauthenticated",
		2000: "validation_error",
		2002: "undefined_child_type",
		2003: "undefined_child_kind",
		2006: "undefined_relation_reference",
		2007: "not_supported_relation_walk",
		2008: "entity_and_subject_cannot_be_equal",
		2009: "depth_not_enough",
		2010: "relation_reference_not_found_in_entity_references",
		2011: "relation_reference_must_have_one_entity_reference",
		2012: "duplicated_entity_reference",
		2013: "duplicated_relation_reference",
		2014: "duplicated_action_reference",
		2015: "schema_parse",
		2016: "schema_compile",
		4000: "not_found_error",
		4001: "entity_type_can_not_found",
		4002: "action_can_not_found",
		4003: "schema_not_found",
		4004: "subject_type_not_found",
		4005: "entity_definition_not_found",
		4006: "action_definition_not_found",
		4007: "relation_definition_not_found",
		4008: "record_not_found",
		5000: "internal_error",
		5001: "cancelled",
		5002: "sql_builder_error",
		5003: "circuit_breaker_error",
		5004: "unique_constraint",
		5005: "execution",
		5006: "scan",
		5007: "migration",
	}
	ErrorCode_value = map[string]int32{
		"no_error":                                          0,
		"missing_bearer_token":                              1001,
		"unauthenticated":                                   1002,
		"validation_error":                                  2000,
		"undefined_child_type":                              2002,
		"undefined_child_kind":                              2003,
		"undefined_relation_reference":                      2006,
		"not_supported_relation_walk":                       2007,
		"entity_and_subject_cannot_be_equal":                2008,
		"depth_not_enough":                                  2009,
		"relation_reference_not_found_in_entity_references": 2010,
		"relation_reference_must_have_one_entity_reference": 2011,
		"duplicated_entity_reference":                       2012,
		"duplicated_relation_reference":                     2013,
		"duplicated_action_reference":                       2014,
		"schema_parse":                                      2015,
		"schema_compile":                                    2016,
		"not_found_error":                                   4000,
		"entity_type_can_not_found":                         4001,
		"action_can_not_found":                              4002,
		"schema_not_found":                                  4003,
		"subject_type_not_found":                            4004,
		"entity_definition_not_found":                       4005,
		"action_definition_not_found":                       4006,
		"relation_definition_not_found":                     4007,
		"record_not_found":                                  4008,
		"internal_error":                                    5000,
		"cancelled":                                         5001,
		"sql_builder_error":                                 5002,
		"circuit_breaker_error":                             5003,
		"unique_constraint":                                 5004,
		"execution":                                         5005,
		"scan":                                              5006,
		"migration":                                         5007,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{0}
}

// ErrorResponse
type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=ErrorCode" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorResponse) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_no_error
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_errors_proto protoreflect.FileDescriptor

var file_errors_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0xbc, 0x07, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x14, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0xe9, 0x07,
	0x12, 0x14, 0x0a, 0x0f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x10, 0xea, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xd0, 0x0f, 0x12, 0x19, 0x0a,
	0x14, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x10, 0xd2, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x75, 0x6e, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x10, 0xd3, 0x0f, 0x12, 0x21, 0x0a, 0x1c, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x10, 0xd6, 0x0f, 0x12, 0x20, 0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x77, 0x61, 0x6c, 0x6b, 0x10, 0xd7, 0x0f, 0x12, 0x27, 0x0a, 0x22, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x10, 0xd8,
	0x0f, 0x12, 0x15, 0x0a, 0x10, 0x64, 0x65, 0x70, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x65,
	0x6e, 0x6f, 0x75, 0x67, 0x68, 0x10, 0xd9, 0x0f, 0x12, 0x36, 0x0a, 0x31, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x10, 0xda, 0x0f,
	0x12, 0x36, 0x0a, 0x31, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x76, 0x65,
	0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x10, 0xdb, 0x0f, 0x12, 0x20, 0x0a, 0x1b, 0x64, 0x75, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x10, 0xdc, 0x0f, 0x12, 0x22, 0x0a, 0x1d, 0x64, 0x75,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x10, 0xdd, 0x0f, 0x12, 0x20,
	0x0a, 0x1b, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x10, 0xde, 0x0f,
	0x12, 0x11, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x10, 0xdf, 0x0f, 0x12, 0x13, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x10, 0xe0, 0x0f, 0x12, 0x14, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x5f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0xa0, 0x1f, 0x12, 0x1e,
	0x0a, 0x19, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x61,
	0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa1, 0x1f, 0x12, 0x19,
	0x0a, 0x14, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6e, 0x5f, 0x6e, 0x6f, 0x74,
	0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa2, 0x1f, 0x12, 0x15, 0x0a, 0x10, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa3, 0x1f,
	0x12, 0x1b, 0x0a, 0x16, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa4, 0x1f, 0x12, 0x20, 0x0a,
	0x1b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa5, 0x1f, 0x12,
	0x20, 0x0a, 0x1b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa6,
	0x1f, 0x12, 0x22, 0x0a, 0x1d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0xa7, 0x1f, 0x12, 0x15, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0xa8, 0x1f, 0x12, 0x13, 0x0a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x88,
	0x27, 0x12, 0x0e, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x89,
	0x27, 0x12, 0x16, 0x0a, 0x11, 0x73, 0x71, 0x6c, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x8a, 0x27, 0x12, 0x1a, 0x0a, 0x15, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x8b, 0x27, 0x12, 0x16, 0x0a, 0x11, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x10, 0x8c, 0x27, 0x12, 0x0e, 0x0a,
	0x09, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x8d, 0x27, 0x12, 0x09, 0x0a,
	0x04, 0x73, 0x63, 0x61, 0x6e, 0x10, 0x8e, 0x27, 0x12, 0x0e, 0x0a, 0x09, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x8f, 0x27, 0x42, 0x3a, 0x42, 0x0b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_proto_rawDescOnce sync.Once
	file_errors_proto_rawDescData = file_errors_proto_rawDesc
)

func file_errors_proto_rawDescGZIP() []byte {
	file_errors_proto_rawDescOnce.Do(func() {
		file_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_proto_rawDescData)
	})
	return file_errors_proto_rawDescData
}

var file_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_proto_goTypes = []interface{}{
	(ErrorCode)(0),        // 0: ErrorCode
	(*ErrorResponse)(nil), // 1: ErrorResponse
}
var file_errors_proto_depIdxs = []int32{
	0, // 0: ErrorResponse.code:type_name -> ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_proto_init() }
func file_errors_proto_init() {
	if File_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_proto_goTypes,
		DependencyIndexes: file_errors_proto_depIdxs,
		EnumInfos:         file_errors_proto_enumTypes,
		MessageInfos:      file_errors_proto_msgTypes,
	}.Build()
	File_errors_proto = out.File
	file_errors_proto_rawDesc = nil
	file_errors_proto_goTypes = nil
	file_errors_proto_depIdxs = nil
}
