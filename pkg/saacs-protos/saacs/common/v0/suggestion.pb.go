// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: saacs/common/v0/suggestion.proto

package v0

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// Key should be
// {auth.Suggestion}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUGGESTION_ID}
type Suggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryKey   *ItemKey               `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
	SuggestionId string                 `protobuf:"bytes,2,opt,name=suggestion_id,json=suggestionId,proto3" json:"suggestion_id,omitempty"`
	Paths        *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=paths,proto3" json:"paths,omitempty"`
	Value        *anypb.Any             `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Suggestion) Reset() {
	*x = Suggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_suggestion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggestion) ProtoMessage() {}

func (x *Suggestion) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_suggestion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggestion.ProtoReflect.Descriptor instead.
func (*Suggestion) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_suggestion_proto_rawDescGZIP(), []int{0}
}

func (x *Suggestion) GetPrimaryKey() *ItemKey {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *Suggestion) GetSuggestionId() string {
	if x != nil {
		return x.SuggestionId
	}
	return ""
}

func (x *Suggestion) GetPaths() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Suggestion) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_saacs_common_v0_suggestion_proto protoreflect.FileDescriptor

var file_saacs_common_v0_suggestion_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x30, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73,
	0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x61, 0x61, 0x63, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x0a, 0x53, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x0d, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x73, 0x75, 0x67, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x17, 0xba, 0xd4, 0x1a, 0x13, 0x10, 0x03, 0x1a, 0x0f,
	0x0a, 0x0d, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42,
	0xbe, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x42, 0x0f, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x33, 0x38, 0x2f, 0x73, 0x61,
	0x61, 0x63, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x30, 0xa2, 0x02, 0x03, 0x53, 0x43, 0x56, 0xaa, 0x02, 0x0f, 0x53, 0x61, 0x61,
	0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x30, 0xca, 0x02, 0x0f, 0x53,
	0x61, 0x61, 0x63, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x30, 0xe2, 0x02,
	0x1b, 0x53, 0x61, 0x61, 0x63, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x30,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x53,
	0x61, 0x61, 0x63, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x30,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_saacs_common_v0_suggestion_proto_rawDescOnce sync.Once
	file_saacs_common_v0_suggestion_proto_rawDescData = file_saacs_common_v0_suggestion_proto_rawDesc
)

func file_saacs_common_v0_suggestion_proto_rawDescGZIP() []byte {
	file_saacs_common_v0_suggestion_proto_rawDescOnce.Do(func() {
		file_saacs_common_v0_suggestion_proto_rawDescData = protoimpl.X.CompressGZIP(file_saacs_common_v0_suggestion_proto_rawDescData)
	})
	return file_saacs_common_v0_suggestion_proto_rawDescData
}

var file_saacs_common_v0_suggestion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_saacs_common_v0_suggestion_proto_goTypes = []interface{}{
	(*Suggestion)(nil),            // 0: saacs.common.v0.Suggestion
	(*ItemKey)(nil),               // 1: saacs.common.v0.ItemKey
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
	(*anypb.Any)(nil),             // 3: google.protobuf.Any
}
var file_saacs_common_v0_suggestion_proto_depIdxs = []int32{
	1, // 0: saacs.common.v0.Suggestion.primary_key:type_name -> saacs.common.v0.ItemKey
	2, // 1: saacs.common.v0.Suggestion.paths:type_name -> google.protobuf.FieldMask
	3, // 2: saacs.common.v0.Suggestion.value:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_saacs_common_v0_suggestion_proto_init() }
func file_saacs_common_v0_suggestion_proto_init() {
	if File_saacs_common_v0_suggestion_proto != nil {
		return
	}
	file_saacs_common_v0_item_proto_init()
	file_saacs_common_v0_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_saacs_common_v0_suggestion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggestion); i {
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
			RawDescriptor: file_saacs_common_v0_suggestion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saacs_common_v0_suggestion_proto_goTypes,
		DependencyIndexes: file_saacs_common_v0_suggestion_proto_depIdxs,
		MessageInfos:      file_saacs_common_v0_suggestion_proto_msgTypes,
	}.Build()
	File_saacs_common_v0_suggestion_proto = out.File
	file_saacs_common_v0_suggestion_proto_rawDesc = nil
	file_saacs_common_v0_suggestion_proto_goTypes = nil
	file_saacs_common_v0_suggestion_proto_depIdxs = nil
}
