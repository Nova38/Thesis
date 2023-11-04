// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: hlf/utils.proto

package hlf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_TRANSACTION_TYPE_INVOKE      TransactionType = 1
	TransactionType_TRANSACTION_TYPE_QUERY       TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSACTION_TYPE_UNSPECIFIED",
		1: "TRANSACTION_TYPE_INVOKE",
		2: "TRANSACTION_TYPE_QUERY",
	}
	TransactionType_value = map[string]int32{
		"TRANSACTION_TYPE_UNSPECIFIED": 0,
		"TRANSACTION_TYPE_INVOKE":      1,
		"TRANSACTION_TYPE_QUERY":       2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_hlf_utils_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_hlf_utils_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_hlf_utils_proto_rawDescGZIP(), []int{0}
}

var file_hlf_utils_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*TransactionType)(nil),
		Field:         50556,
		Name:          "hlf.transaction_type",
		Tag:           "varint,50556,opt,name=transaction_type,enum=hlf.TransactionType",
		Filename:      "hlf/utils.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional hlf.TransactionType transaction_type = 50556;
	E_TransactionType = &file_hlf_utils_proto_extTypes[0]
)

var File_hlf_utils_proto protoreflect.FileDescriptor

var file_hlf_utils_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x6c, 0x66, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x68, 0x6c, 0x66, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x6c, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a,
	0x17, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x10, 0x02, 0x3a, 0x64, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfc, 0x8a, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x68, 0x6c, 0x66, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x6a, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x6c, 0x66, 0x42, 0x0a, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x33, 0x38, 0x2f, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x2f,
	0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x6c, 0x66, 0xa2, 0x02,
	0x03, 0x48, 0x58, 0x58, 0xaa, 0x02, 0x03, 0x48, 0x6c, 0x66, 0xca, 0x02, 0x03, 0x48, 0x6c, 0x66,
	0xe2, 0x02, 0x0f, 0x48, 0x6c, 0x66, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x03, 0x48, 0x6c, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hlf_utils_proto_rawDescOnce sync.Once
	file_hlf_utils_proto_rawDescData = file_hlf_utils_proto_rawDesc
)

func file_hlf_utils_proto_rawDescGZIP() []byte {
	file_hlf_utils_proto_rawDescOnce.Do(func() {
		file_hlf_utils_proto_rawDescData = protoimpl.X.CompressGZIP(file_hlf_utils_proto_rawDescData)
	})
	return file_hlf_utils_proto_rawDescData
}

var file_hlf_utils_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hlf_utils_proto_goTypes = []interface{}{
	(TransactionType)(0),               // 0: hlf.TransactionType
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_hlf_utils_proto_depIdxs = []int32{
	1, // 0: hlf.transaction_type:extendee -> google.protobuf.MethodOptions
	0, // 1: hlf.transaction_type:type_name -> hlf.TransactionType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hlf_utils_proto_init() }
func file_hlf_utils_proto_init() {
	if File_hlf_utils_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hlf_utils_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_hlf_utils_proto_goTypes,
		DependencyIndexes: file_hlf_utils_proto_depIdxs,
		EnumInfos:         file_hlf_utils_proto_enumTypes,
		ExtensionInfos:    file_hlf_utils_proto_extTypes,
	}.Build()
	File_hlf_utils_proto = out.File
	file_hlf_utils_proto_rawDesc = nil
	file_hlf_utils_proto_goTypes = nil
	file_hlf_utils_proto_depIdxs = nil
}
