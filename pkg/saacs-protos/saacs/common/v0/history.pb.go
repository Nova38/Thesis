// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: saacs/common/v0/history.proto

package v0

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HiddenTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId      string                 `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	MspId     string                 `protobuf:"bytes,2,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	UserId    string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Note      string                 `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *HiddenTx) Reset() {
	*x = HiddenTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiddenTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiddenTx) ProtoMessage() {}

func (x *HiddenTx) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiddenTx.ProtoReflect.Descriptor instead.
func (*HiddenTx) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_history_proto_rawDescGZIP(), []int{0}
}

func (x *HiddenTx) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *HiddenTx) GetMspId() string {
	if x != nil {
		return x.MspId
	}
	return ""
}

func (x *HiddenTx) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HiddenTx) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *HiddenTx) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

// Key should be {COLLECTION_ID}{auth.HiddenTxList}{?msp_id}{ITEM_TYPE}{...ITEM_ID}
type HiddenTxList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrimaryKey *ItemKey    `protobuf:"bytes,1,opt,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"` // The key that is used to store the item
	MspId      string      `protobuf:"bytes,2,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	Txs        []*HiddenTx `protobuf:"bytes,4,rep,name=txs,proto3" json:"txs,omitempty"` // The list of hidden txs by tx_id
}

func (x *HiddenTxList) Reset() {
	*x = HiddenTxList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiddenTxList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiddenTxList) ProtoMessage() {}

func (x *HiddenTxList) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiddenTxList.ProtoReflect.Descriptor instead.
func (*HiddenTxList) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_history_proto_rawDescGZIP(), []int{1}
}

func (x *HiddenTxList) GetPrimaryKey() *ItemKey {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

func (x *HiddenTxList) GetMspId() string {
	if x != nil {
		return x.MspId
	}
	return ""
}

func (x *HiddenTxList) GetTxs() []*HiddenTx {
	if x != nil {
		return x.Txs
	}
	return nil
}

type HistoryEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction id that caused the change
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// Whether the item was deleted
	IsDelete bool `protobuf:"varint,2,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	// Whether the transaction was hidden
	IsHidden bool `protobuf:"varint,3,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	// The timestamp of the change
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A note about the change
	Note string `protobuf:"bytes,5,opt,name=note,proto3" json:"note,omitempty"`
	// The value of the item
	Value *anypb.Any `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HistoryEntry) Reset() {
	*x = HistoryEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryEntry) ProtoMessage() {}

func (x *HistoryEntry) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryEntry.ProtoReflect.Descriptor instead.
func (*HistoryEntry) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_history_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryEntry) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *HistoryEntry) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

func (x *HistoryEntry) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *HistoryEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *HistoryEntry) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *HistoryEntry) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries   []*HistoryEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	HiddenTxs *HiddenTxList   `protobuf:"bytes,2,opt,name=hidden_txs,json=hiddenTxs,proto3" json:"hidden_txs,omitempty"`
	// The key is the msp_id of the group that is hiding the tx
	HiddenTxsByMspId map[string]*HiddenTxList `protobuf:"bytes,3,rep,name=hidden_txs_by_msp_id,json=hiddenTxsByMspId,proto3" json:"hidden_txs_by_msp_id,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *History) Reset() {
	*x = History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_history_proto_rawDescGZIP(), []int{3}
}

func (x *History) GetEntries() []*HistoryEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *History) GetHiddenTxs() *HiddenTxList {
	if x != nil {
		return x.HiddenTxs
	}
	return nil
}

func (x *History) GetHiddenTxsByMspId() map[string]*HiddenTxList {
	if x != nil {
		return x.HiddenTxsByMspId
	}
	return nil
}

type HistoryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Include bool           `protobuf:"varint,1,opt,name=include,proto3" json:"include,omitempty"`
	Hidden  *HiddenOptions `protobuf:"bytes,2,opt,name=hidden,proto3" json:"hidden,omitempty"`
}

func (x *HistoryOptions) Reset() {
	*x = HistoryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryOptions) ProtoMessage() {}

func (x *HistoryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryOptions.ProtoReflect.Descriptor instead.
func (*HistoryOptions) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_history_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryOptions) GetInclude() bool {
	if x != nil {
		return x.Include
	}
	return false
}

func (x *HistoryOptions) GetHidden() *HiddenOptions {
	if x != nil {
		return x.Hidden
	}
	return nil
}

type HiddenOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Include bool     `protobuf:"varint,1,opt,name=include,proto3" json:"include,omitempty"`
	MspIds  []string `protobuf:"bytes,3,rep,name=msp_ids,json=mspIds,proto3" json:"msp_ids,omitempty"`
}

func (x *HiddenOptions) Reset() {
	*x = HiddenOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_saacs_common_v0_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiddenOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiddenOptions) ProtoMessage() {}

func (x *HiddenOptions) ProtoReflect() protoreflect.Message {
	mi := &file_saacs_common_v0_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiddenOptions.ProtoReflect.Descriptor instead.
func (*HiddenOptions) Descriptor() ([]byte, []int) {
	return file_saacs_common_v0_history_proto_rawDescGZIP(), []int{5}
}

func (x *HiddenOptions) GetInclude() bool {
	if x != nil {
		return x.Include
	}
	return false
}

func (x *HiddenOptions) GetMspIds() []string {
	if x != nil {
		return x.MspIds
	}
	return nil
}

var File_saacs_common_v0_history_proto protoreflect.FileDescriptor

var file_saacs_common_v0_history_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x30, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x61, 0x61,
	0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x48, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x54, 0x78, 0x12, 0x1c, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x78, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6d, 0x73, 0x70, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x0c, 0x48, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4b, 0x65, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x15, 0x0a,
	0x06, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x73, 0x70, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x03, 0x74, 0x78, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x30, 0x2e, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x54, 0x78, 0x52, 0x03, 0x74, 0x78,
	0x73, 0x3a, 0x10, 0xba, 0xd4, 0x1a, 0x0c, 0x10, 0x03, 0x1a, 0x08, 0x0a, 0x06, 0x6d, 0x73, 0x70,
	0x5f, 0x69, 0x64, 0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x78,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x40, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f,
	0x74, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc4,
	0x02, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x61,
	0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x74, 0x78,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x54, 0x78,
	0x73, 0x12, 0x5e, 0x0a, 0x14, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x74, 0x78, 0x73, 0x5f,
	0x62, 0x79, 0x5f, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x30, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x54, 0x78, 0x73, 0x42, 0x79, 0x4d, 0x73, 0x70, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x10, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x54, 0x78, 0x73, 0x42, 0x79, 0x4d, 0x73, 0x70, 0x49,
	0x64, 0x1a, 0x62, 0x0a, 0x15, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x54, 0x78, 0x73, 0x42, 0x79,
	0x4d, 0x73, 0x70, 0x49, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x61,
	0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x2e, 0x48, 0x69,
	0x64, 0x64, 0x65, 0x6e, 0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x62, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x30, 0x2e, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x0d, 0x48, 0x69, 0x64,
	0x64, 0x65, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x73, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x73, 0x70, 0x49, 0x64, 0x73, 0x42, 0xbb, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x30, 0x42, 0x0c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x6f, 0x76, 0x61, 0x33, 0x38, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x73, 0x61, 0x61, 0x63, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x73, 0x61, 0x61, 0x63, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x30, 0xa2,
	0x02, 0x03, 0x53, 0x43, 0x56, 0xaa, 0x02, 0x0f, 0x53, 0x61, 0x61, 0x63, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x30, 0xca, 0x02, 0x0f, 0x53, 0x61, 0x61, 0x63, 0x73, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x30, 0xe2, 0x02, 0x1b, 0x53, 0x61, 0x61, 0x63,
	0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x30, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x53, 0x61, 0x61, 0x63, 0x73, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_saacs_common_v0_history_proto_rawDescOnce sync.Once
	file_saacs_common_v0_history_proto_rawDescData = file_saacs_common_v0_history_proto_rawDesc
)

func file_saacs_common_v0_history_proto_rawDescGZIP() []byte {
	file_saacs_common_v0_history_proto_rawDescOnce.Do(func() {
		file_saacs_common_v0_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_saacs_common_v0_history_proto_rawDescData)
	})
	return file_saacs_common_v0_history_proto_rawDescData
}

var file_saacs_common_v0_history_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_saacs_common_v0_history_proto_goTypes = []interface{}{
	(*HiddenTx)(nil),              // 0: saacs.common.v0.HiddenTx
	(*HiddenTxList)(nil),          // 1: saacs.common.v0.HiddenTxList
	(*HistoryEntry)(nil),          // 2: saacs.common.v0.HistoryEntry
	(*History)(nil),               // 3: saacs.common.v0.History
	(*HistoryOptions)(nil),        // 4: saacs.common.v0.HistoryOptions
	(*HiddenOptions)(nil),         // 5: saacs.common.v0.HiddenOptions
	nil,                           // 6: saacs.common.v0.History.HiddenTxsByMspIdEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*ItemKey)(nil),               // 8: saacs.common.v0.ItemKey
	(*anypb.Any)(nil),             // 9: google.protobuf.Any
}
var file_saacs_common_v0_history_proto_depIdxs = []int32{
	7,  // 0: saacs.common.v0.HiddenTx.timestamp:type_name -> google.protobuf.Timestamp
	8,  // 1: saacs.common.v0.HiddenTxList.primary_key:type_name -> saacs.common.v0.ItemKey
	0,  // 2: saacs.common.v0.HiddenTxList.txs:type_name -> saacs.common.v0.HiddenTx
	7,  // 3: saacs.common.v0.HistoryEntry.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 4: saacs.common.v0.HistoryEntry.value:type_name -> google.protobuf.Any
	2,  // 5: saacs.common.v0.History.entries:type_name -> saacs.common.v0.HistoryEntry
	1,  // 6: saacs.common.v0.History.hidden_txs:type_name -> saacs.common.v0.HiddenTxList
	6,  // 7: saacs.common.v0.History.hidden_txs_by_msp_id:type_name -> saacs.common.v0.History.HiddenTxsByMspIdEntry
	5,  // 8: saacs.common.v0.HistoryOptions.hidden:type_name -> saacs.common.v0.HiddenOptions
	1,  // 9: saacs.common.v0.History.HiddenTxsByMspIdEntry.value:type_name -> saacs.common.v0.HiddenTxList
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_saacs_common_v0_history_proto_init() }
func file_saacs_common_v0_history_proto_init() {
	if File_saacs_common_v0_history_proto != nil {
		return
	}
	file_saacs_common_v0_item_proto_init()
	file_saacs_common_v0_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_saacs_common_v0_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiddenTx); i {
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
		file_saacs_common_v0_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiddenTxList); i {
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
		file_saacs_common_v0_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryEntry); i {
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
		file_saacs_common_v0_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*History); i {
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
		file_saacs_common_v0_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryOptions); i {
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
		file_saacs_common_v0_history_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiddenOptions); i {
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
			RawDescriptor: file_saacs_common_v0_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saacs_common_v0_history_proto_goTypes,
		DependencyIndexes: file_saacs_common_v0_history_proto_depIdxs,
		MessageInfos:      file_saacs_common_v0_history_proto_msgTypes,
	}.Build()
	File_saacs_common_v0_history_proto = out.File
	file_saacs_common_v0_history_proto_rawDesc = nil
	file_saacs_common_v0_history_proto_goTypes = nil
	file_saacs_common_v0_history_proto_depIdxs = nil
}