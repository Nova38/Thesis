// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: saacs/common/v0/operation.proto

package v0

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action       Action                 `protobuf:"varint,1,opt,name=action,proto3,enum=saacs.common.v0.Action" json:"action,omitempty"`
	CollectionId string                 `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	ItemType     string                 `protobuf:"bytes,3,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	Paths        *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=paths,proto3" json:"paths,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

func (x *Operation) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *Operation) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *Operation) GetPaths() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Paths
	}
	return nil
}

var File_saacs_common_v0_operation_proto protoreflect.FileDescriptor

var file_saacs_common_v0_operation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x30, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcf, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x82, 0x01, 0x05, 0x10,
	0x01, 0x22, 0x01, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x42, 0xbd, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x42, 0x0e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x33, 0x38, 0x2f,
	0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x56, 0xaa, 0x02, 0x0f, 0x53,
	0x61, 0x61, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x30, 0xca, 0x02,
	0x0f, 0x53, 0x61, 0x61, 0x63, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x30,
	0xe2, 0x02, 0x1b, 0x53, 0x61, 0x61, 0x63, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c,
	0x56, 0x30, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x53, 0x61, 0x61, 0x63, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a,
	0x56, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saacs_common_v0_operation_proto_rawDescOnce sync.Once
	file_saacs_common_v0_operation_proto_rawDescData = file_saacs_common_v0_operation_proto_rawDesc
)

func file_saacs_common_v0_operation_proto_rawDescGZIP() []byte {
	file_saacs_common_v0_operation_proto_rawDescOnce.Do(func() {
		file_saacs_common_v0_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_saacs_common_v0_operation_proto_rawDescData)
	})
	return file_saacs_common_v0_operation_proto_rawDescData
}

var file_saacs_common_v0_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_saacs_common_v0_operation_proto_goTypes = []interface{}{
	(*Operation)(nil),             // 0: saacs.common.v0.Operation
	(Action)(0),                   // 1: saacs.common.v0.Action
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
}
var file_saacs_common_v0_operation_proto_depIdxs = []int32{
	1, // 0: saacs.common.v0.Operation.action:type_name -> saacs.common.v0.Action
	2, // 1: saacs.common.v0.Operation.paths:type_name -> google.protobuf.FieldMask
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_saacs_common_v0_operation_proto_init() }
func file_saacs_common_v0_operation_proto_init() {
	if File_saacs_common_v0_operation_proto != nil {
		return
	}
	file_saacs_common_v0_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_saacs_common_v0_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
			RawDescriptor: file_saacs_common_v0_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saacs_common_v0_operation_proto_goTypes,
		DependencyIndexes: file_saacs_common_v0_operation_proto_depIdxs,
		MessageInfos:      file_saacs_common_v0_operation_proto_msgTypes,
	}.Build()
	File_saacs_common_v0_operation_proto = out.File
	file_saacs_common_v0_operation_proto_rawDesc = nil
	file_saacs_common_v0_operation_proto_goTypes = nil
	file_saacs_common_v0_operation_proto_depIdxs = nil
}
