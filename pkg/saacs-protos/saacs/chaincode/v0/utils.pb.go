// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: saacs/chaincode/v0/utils.proto

package v0

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v01 "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
	v0 "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCurrentUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty bool `protobuf:"varint,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *GetCurrentUserRequest) Reset() {
	*x = GetCurrentUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentUserRequest) ProtoMessage() {}

func (x *GetCurrentUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentUserRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentUserRequest) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrentUserRequest) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

type GetCurrentUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       *v0.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Registered bool     `protobuf:"varint,2,opt,name=registered,proto3" json:"registered,omitempty"`
}

func (x *GetCurrentUserResponse) Reset() {
	*x = GetCurrentUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentUserResponse) ProtoMessage() {}

func (x *GetCurrentUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentUserResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentUserResponse) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentUserResponse) GetUser() *v0.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetCurrentUserResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

type GetCurrentFullUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User                *v0.User                    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Registered          bool                        `protobuf:"varint,2,opt,name=registered,proto3" json:"registered,omitempty"`
	UserCollectionRoles []*v01.UserCollectionRoles  `protobuf:"bytes,3,rep,name=user_collection_roles,json=userCollectionRoles,proto3" json:"user_collection_roles,omitempty"`
	UserMemberships     []*v01.UserDirectMembership `protobuf:"bytes,4,rep,name=user_memberships,json=userMemberships,proto3" json:"user_memberships,omitempty"`
}

func (x *GetCurrentFullUserResponse) Reset() {
	*x = GetCurrentFullUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentFullUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentFullUserResponse) ProtoMessage() {}

func (x *GetCurrentFullUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentFullUserResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentFullUserResponse) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrentFullUserResponse) GetUser() *v0.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *GetCurrentFullUserResponse) GetRegistered() bool {
	if x != nil {
		return x.Registered
	}
	return false
}

func (x *GetCurrentFullUserResponse) GetUserCollectionRoles() []*v01.UserCollectionRoles {
	if x != nil {
		return x.UserCollectionRoles
	}
	return nil
}

func (x *GetCurrentFullUserResponse) GetUserMemberships() []*v01.UserDirectMembership {
	if x != nil {
		return x.UserMemberships
	}
	return nil
}

type AuthorizeOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *v0.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *AuthorizeOperationRequest) Reset() {
	*x = AuthorizeOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeOperationRequest) ProtoMessage() {}

func (x *AuthorizeOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeOperationRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeOperationRequest) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizeOperationRequest) GetOperation() *v0.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AuthorizeOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorized bool `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
}

func (x *AuthorizeOperationResponse) Reset() {
	*x = AuthorizeOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeOperationResponse) ProtoMessage() {}

func (x *AuthorizeOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeOperationResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeOperationResponse) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorizeOperationResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

type GetCollectionsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty bool `protobuf:"varint,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *GetCollectionsListRequest) Reset() {
	*x = GetCollectionsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionsListRequest) ProtoMessage() {}

func (x *GetCollectionsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionsListRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionsListRequest) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{5}
}

func (x *GetCollectionsListRequest) GetEmpty() bool {
	if x != nil {
		return x.Empty
	}
	return false
}

type GetCollectionsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*v01.Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
}

func (x *GetCollectionsListResponse) Reset() {
	*x = GetCollectionsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionsListResponse) ProtoMessage() {}

func (x *GetCollectionsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionsListResponse.ProtoReflect.Descriptor instead.
func (*GetCollectionsListResponse) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{6}
}

func (x *GetCollectionsListResponse) GetCollections() []*v01.Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type BootstrapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection *v01.Collection `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *BootstrapRequest) Reset() {
	*x = BootstrapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapRequest) ProtoMessage() {}

func (x *BootstrapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapRequest.ProtoReflect.Descriptor instead.
func (*BootstrapRequest) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{7}
}

func (x *BootstrapRequest) GetCollection() *v01.Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

type BootstrapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Collection *v01.Collection `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
}

func (x *BootstrapResponse) Reset() {
	*x = BootstrapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapResponse) ProtoMessage() {}

func (x *BootstrapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapResponse.ProtoReflect.Descriptor instead.
func (*BootstrapResponse) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{8}
}

func (x *BootstrapResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BootstrapResponse) GetCollection() *v01.Collection {
	if x != nil {
		return x.Collection
	}
	return nil
}

type GetCollectionAuthModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
}

func (x *GetCollectionAuthModelRequest) Reset() {
	*x = GetCollectionAuthModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionAuthModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionAuthModelRequest) ProtoMessage() {}

func (x *GetCollectionAuthModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_chaincode_v0_utils_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionAuthModelRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionAuthModelRequest) Descriptor() ([]byte, []int) {
	return file_saacs_chaincode_v0_utils_proto_rawDescGZIP(), []int{9}
}

func (x *GetCollectionAuthModelRequest) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

var File_saacs_chaincode_v0_utils_proto protoreflect.FileDescriptor

var file_saacs_chaincode_v0_utils_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x2f, 0x76, 0x30, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x30, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x30, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x30,
	0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x73, 0x61,
	0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73,
	0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x61,
	0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x63, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x8f, 0x02, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x61, 0x61, 0x63,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x13, 0x75, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x4e,
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0x55,
	0x0a, 0x19, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x22, 0x31, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x59, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x61, 0x61, 0x63,
	0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x54, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x61, 0x61, 0x63,
	0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x68, 0x0a, 0x11, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x61, 0x61,
	0x63, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x44, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xf6, 0x03, 0x0a, 0x0c, 0x55, 0x74, 0x69, 0x6c,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x73, 0x61, 0x61,
	0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x30, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0d, 0xe0, 0xd7, 0x18, 0x02, 0xfa, 0x9a, 0x1c, 0x02, 0x08, 0x0a, 0x90, 0x02, 0x01,
	0x12, 0x64, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x24, 0x2e,
	0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x30, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0a, 0xe0, 0xd7, 0x18, 0x01,
	0xfa, 0x9a, 0x1c, 0x02, 0x08, 0x01, 0x12, 0x82, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e,
	0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x76, 0x30, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73,
	0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x30, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0xe0, 0xd7,
	0x18, 0x01, 0xfa, 0x9a, 0x1c, 0x02, 0x08, 0x01, 0x90, 0x02, 0x01, 0x12, 0x82, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x0d, 0xe0, 0xd7, 0x18, 0x02, 0xfa, 0x9a, 0x1c, 0x02, 0x08, 0x0a, 0x90, 0x02, 0x01,
	0x42, 0xcb, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x30, 0x42, 0x0a, 0x55, 0x74, 0x69,
	0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x33, 0x38, 0x2f, 0x73, 0x61, 0x61,
	0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x2f, 0x76, 0x30, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x56, 0xaa, 0x02, 0x12, 0x53,
	0x61, 0x61, 0x63, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x56,
	0x30, 0xca, 0x02, 0x12, 0x53, 0x61, 0x61, 0x63, 0x73, 0x5c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x5c, 0x56, 0x30, 0xe2, 0x02, 0x1e, 0x53, 0x61, 0x61, 0x63, 0x73, 0x5c, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5c, 0x56, 0x30, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x61, 0x61, 0x63, 0x73, 0x3a,
	0x3a, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x30, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saacs_chaincode_v0_utils_proto_rawDescOnce sync.Once
	file_saacs_chaincode_v0_utils_proto_rawDescData = file_saacs_chaincode_v0_utils_proto_rawDesc
)

func file_saacs_chaincode_v0_utils_proto_rawDescGZIP() []byte {
	file_saacs_chaincode_v0_utils_proto_rawDescOnce.Do(func() {
		file_saacs_chaincode_v0_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_saacs_chaincode_v0_utils_proto_rawDescData)
	})
	return file_saacs_chaincode_v0_utils_proto_rawDescData
}

var file_saacs_chaincode_v0_utils_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_saacs_chaincode_v0_utils_proto_goTypes = []interface{}{
	(*GetCurrentUserRequest)(nil),         // 0: saacs.chaincode.v0.GetCurrentUserRequest
	(*GetCurrentUserResponse)(nil),        // 1: saacs.chaincode.v0.GetCurrentUserResponse
	(*GetCurrentFullUserResponse)(nil),    // 2: saacs.chaincode.v0.GetCurrentFullUserResponse
	(*AuthorizeOperationRequest)(nil),     // 3: saacs.chaincode.v0.AuthorizeOperationRequest
	(*AuthorizeOperationResponse)(nil),    // 4: saacs.chaincode.v0.AuthorizeOperationResponse
	(*GetCollectionsListRequest)(nil),     // 5: saacs.chaincode.v0.GetCollectionsListRequest
	(*GetCollectionsListResponse)(nil),    // 6: saacs.chaincode.v0.GetCollectionsListResponse
	(*BootstrapRequest)(nil),              // 7: saacs.chaincode.v0.BootstrapRequest
	(*BootstrapResponse)(nil),             // 8: saacs.chaincode.v0.BootstrapResponse
	(*GetCollectionAuthModelRequest)(nil), // 9: saacs.chaincode.v0.GetCollectionAuthModelRequest
	(*v0.User)(nil),                       // 10: saacs.common.v0.User
	(*v01.UserCollectionRoles)(nil),       // 11: saacs.auth.v0.UserCollectionRoles
	(*v01.UserDirectMembership)(nil),      // 12: saacs.auth.v0.UserDirectMembership
	(*v0.Operation)(nil),                  // 13: saacs.common.v0.Operation
	(*v01.Collection)(nil),                // 14: saacs.auth.v0.Collection
}
var file_saacs_chaincode_v0_utils_proto_depIdxs = []int32{
	10, // 0: saacs.chaincode.v0.GetCurrentUserResponse.user:type_name -> saacs.common.v0.User
	10, // 1: saacs.chaincode.v0.GetCurrentFullUserResponse.user:type_name -> saacs.common.v0.User
	11, // 2: saacs.chaincode.v0.GetCurrentFullUserResponse.user_collection_roles:type_name -> saacs.auth.v0.UserCollectionRoles
	12, // 3: saacs.chaincode.v0.GetCurrentFullUserResponse.user_memberships:type_name -> saacs.auth.v0.UserDirectMembership
	13, // 4: saacs.chaincode.v0.AuthorizeOperationRequest.operation:type_name -> saacs.common.v0.Operation
	14, // 5: saacs.chaincode.v0.GetCollectionsListResponse.collections:type_name -> saacs.auth.v0.Collection
	14, // 6: saacs.chaincode.v0.BootstrapRequest.collection:type_name -> saacs.auth.v0.Collection
	14, // 7: saacs.chaincode.v0.BootstrapResponse.collection:type_name -> saacs.auth.v0.Collection
	0,  // 8: saacs.chaincode.v0.UtilsService.GetCurrentUser:input_type -> saacs.chaincode.v0.GetCurrentUserRequest
	7,  // 9: saacs.chaincode.v0.UtilsService.Bootstrap:input_type -> saacs.chaincode.v0.BootstrapRequest
	3,  // 10: saacs.chaincode.v0.UtilsService.AuthorizeOperation:input_type -> saacs.chaincode.v0.AuthorizeOperationRequest
	5,  // 11: saacs.chaincode.v0.UtilsService.GetCollectionsList:input_type -> saacs.chaincode.v0.GetCollectionsListRequest
	1,  // 12: saacs.chaincode.v0.UtilsService.GetCurrentUser:output_type -> saacs.chaincode.v0.GetCurrentUserResponse
	8,  // 13: saacs.chaincode.v0.UtilsService.Bootstrap:output_type -> saacs.chaincode.v0.BootstrapResponse
	4,  // 14: saacs.chaincode.v0.UtilsService.AuthorizeOperation:output_type -> saacs.chaincode.v0.AuthorizeOperationResponse
	6,  // 15: saacs.chaincode.v0.UtilsService.GetCollectionsList:output_type -> saacs.chaincode.v0.GetCollectionsListResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_saacs_chaincode_v0_utils_proto_init() }
func file_saacs_chaincode_v0_utils_proto_init() {
	if File_saacs_chaincode_v0_utils_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_saacs_chaincode_v0_utils_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentUserRequest); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentUserResponse); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentFullUserResponse); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeOperationRequest); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeOperationResponse); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionsListRequest); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionsListResponse); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapRequest); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapResponse); i {
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
		file_saacs_chaincode_v0_utils_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionAuthModelRequest); i {
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
			RawDescriptor: file_saacs_chaincode_v0_utils_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_saacs_chaincode_v0_utils_proto_goTypes,
		DependencyIndexes: file_saacs_chaincode_v0_utils_proto_depIdxs,
		MessageInfos:      file_saacs_chaincode_v0_utils_proto_msgTypes,
	}.Build()
	File_saacs_chaincode_v0_utils_proto = out.File
	file_saacs_chaincode_v0_utils_proto_rawDesc = nil
	file_saacs_chaincode_v0_utils_proto_goTypes = nil
	file_saacs_chaincode_v0_utils_proto_depIdxs = nil
}
