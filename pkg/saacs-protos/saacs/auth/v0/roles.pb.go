// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: saacs/auth/v0/roles.proto

package v0

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
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

// Shared Auth Object for Role Based Authentication
type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId  string   `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	RoleId        string   `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Polices       *Polices `protobuf:"bytes,4,opt,name=polices,proto3" json:"polices,omitempty"`
	Note          string   `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	ParentRoleIds []string `protobuf:"bytes,6,rep,name=parent_role_ids,json=parentRoleIds,proto3" json:"parent_role_ids,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_auth_v0_roles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_auth_v0_roles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_saacs_auth_v0_roles_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *Role) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *Role) GetPolices() *Polices {
	if x != nil {
		return x.Polices
	}
	return nil
}

func (x *Role) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Role) GetParentRoleIds() []string {
	if x != nil {
		return x.ParentRoleIds
	}
	return nil
}

type RoleIDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId []string `protobuf:"bytes,1,rep,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *RoleIDList) Reset() {
	*x = RoleIDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_auth_v0_roles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleIDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleIDList) ProtoMessage() {}

func (x *RoleIDList) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_auth_v0_roles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleIDList.ProtoReflect.Descriptor instead.
func (*RoleIDList) Descriptor() ([]byte, []int) {
	return file_saacs_auth_v0_roles_proto_rawDescGZIP(), []int{1}
}

func (x *RoleIDList) GetRoleId() []string {
	if x != nil {
		return x.RoleId
	}
	return nil
}

// Auth Object For RBAC
type UserCollectionRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection that the user is a member of
	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// The msp of the organization that the user's certificate is from
	MspId string `protobuf:"bytes,2,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	// The id of the user from the certificate
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The roles that the user has in the collection
	RoleIds []string `protobuf:"bytes,4,rep,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`
	Note    string   `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *UserCollectionRoles) Reset() {
	*x = UserCollectionRoles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_auth_v0_roles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCollectionRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCollectionRoles) ProtoMessage() {}

func (x *UserCollectionRoles) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_auth_v0_roles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCollectionRoles.ProtoReflect.Descriptor instead.
func (*UserCollectionRoles) Descriptor() ([]byte, []int) {
	return file_saacs_auth_v0_roles_proto_rawDescGZIP(), []int{2}
}

func (x *UserCollectionRoles) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *UserCollectionRoles) GetMspId() string {
	if x != nil {
		return x.MspId
	}
	return ""
}

func (x *UserCollectionRoles) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserCollectionRoles) GetRoleIds() []string {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

func (x *UserCollectionRoles) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

// Auth Object For Embedded RBAC
type UserGlobalRoles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// The msp of the organization that the user's certificate is from
	MspId string `protobuf:"bytes,2,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	// The id of the user from the certificate
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The roles that the user has in the collection
	// key is the collection id
	// value is the list of rolesIds
	Roles map[string]*RoleIDList `protobuf:"bytes,4,rep,name=roles,proto3" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserGlobalRoles) Reset() {
	*x = UserGlobalRoles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_auth_v0_roles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGlobalRoles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGlobalRoles) ProtoMessage() {}

func (x *UserGlobalRoles) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_auth_v0_roles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGlobalRoles.ProtoReflect.Descriptor instead.
func (*UserGlobalRoles) Descriptor() ([]byte, []int) {
	return file_saacs_auth_v0_roles_proto_rawDescGZIP(), []int{3}
}

func (x *UserGlobalRoles) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *UserGlobalRoles) GetMspId() string {
	if x != nil {
		return x.MspId
	}
	return ""
}

func (x *UserGlobalRoles) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserGlobalRoles) GetRoles() map[string]*RoleIDList {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_saacs_auth_v0_roles_proto protoreflect.FileDescriptor

var file_saacs_auth_v0_roles_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x30, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x61, 0x61,
	0x63, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x61, 0x61, 0x63, 0x73,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x30, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x30, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x65, 0x73,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x3a, 0x11, 0xba,
	0xd4, 0x1a, 0x0d, 0x10, 0x02, 0x1a, 0x09, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x22, 0x25, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xdf, 0x01, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6d, 0x73, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x92, 0x01, 0x04, 0x08, 0x01, 0x18, 0x01, 0x52, 0x07, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x3a, 0x1d, 0xba, 0xd4, 0x1a, 0x15,
	0x10, 0x02, 0x1a, 0x11, 0x0a, 0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x80, 0xa5, 0x1b, 0x03, 0x22, 0xbb, 0x02, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xba, 0x48, 0x09, 0x72, 0x07, 0x0a, 0x05, 0x55, 0x53, 0x45,
	0x52, 0x53, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6d, 0x73, 0x70, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x3f, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x30, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x1a, 0x53, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x30, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x1d, 0xba, 0xd4, 0x1a, 0x15, 0x10, 0x02,
	0x1a, 0x11, 0x0a, 0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x80, 0xa5, 0x1b, 0x04, 0x42, 0xad, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x30, 0x42, 0x0a, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x33, 0x38, 0x2f, 0x73,
	0x61, 0x61, 0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x30, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x56, 0xaa, 0x02, 0x0d, 0x53, 0x61, 0x61, 0x63,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x30, 0xca, 0x02, 0x0d, 0x53, 0x61, 0x61, 0x63,
	0x73, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x30, 0xe2, 0x02, 0x19, 0x53, 0x61, 0x61, 0x63,
	0x73, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x30, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x53, 0x61, 0x61, 0x63, 0x73, 0x3a, 0x3a, 0x41,
	0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saacs_auth_v0_roles_proto_rawDescOnce sync.Once
	file_saacs_auth_v0_roles_proto_rawDescData = file_saacs_auth_v0_roles_proto_rawDesc
)

func file_saacs_auth_v0_roles_proto_rawDescGZIP() []byte {
	file_saacs_auth_v0_roles_proto_rawDescOnce.Do(func() {
		file_saacs_auth_v0_roles_proto_rawDescData = protoimpl.X.CompressGZIP(file_saacs_auth_v0_roles_proto_rawDescData)
	})
	return file_saacs_auth_v0_roles_proto_rawDescData
}

var file_saacs_auth_v0_roles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_saacs_auth_v0_roles_proto_goTypes = []interface{}{
	(*Role)(nil),                // 0: saacs.auth.v0.Role
	(*RoleIDList)(nil),          // 1: saacs.auth.v0.RoleIDList
	(*UserCollectionRoles)(nil), // 2: saacs.auth.v0.UserCollectionRoles
	(*UserGlobalRoles)(nil),     // 3: saacs.auth.v0.UserGlobalRoles
	nil,                         // 4: saacs.auth.v0.UserGlobalRoles.RolesEntry
	(*Polices)(nil),             // 5: saacs.auth.v0.Polices
}
var file_saacs_auth_v0_roles_proto_depIdxs = []int32{
	5, // 0: saacs.auth.v0.Role.polices:type_name -> saacs.auth.v0.Polices
	4, // 1: saacs.auth.v0.UserGlobalRoles.roles:type_name -> saacs.auth.v0.UserGlobalRoles.RolesEntry
	1, // 2: saacs.auth.v0.UserGlobalRoles.RolesEntry.value:type_name -> saacs.auth.v0.RoleIDList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_saacs_auth_v0_roles_proto_init() }
func file_saacs_auth_v0_roles_proto_init() {
	if File_saacs_auth_v0_roles_proto != nil {
		return
	}
	file_saacs_auth_v0_policy_proto_init()
	file_saacs_auth_v0_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_saacs_auth_v0_roles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_saacs_auth_v0_roles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleIDList); i {
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
		file_saacs_auth_v0_roles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCollectionRoles); i {
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
		file_saacs_auth_v0_roles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGlobalRoles); i {
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
			RawDescriptor: file_saacs_auth_v0_roles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saacs_auth_v0_roles_proto_goTypes,
		DependencyIndexes: file_saacs_auth_v0_roles_proto_depIdxs,
		MessageInfos:      file_saacs_auth_v0_roles_proto_msgTypes,
	}.Build()
	File_saacs_auth_v0_roles_proto = out.File
	file_saacs_auth_v0_roles_proto_rawDesc = nil
	file_saacs_auth_v0_roles_proto_goTypes = nil
	file_saacs_auth_v0_roles_proto_depIdxs = nil
}
