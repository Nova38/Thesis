// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: saacs/chaincode/v0/chaincode.proto

package v0

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *GetRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetFullRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetFullRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetFullResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetFullResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListByAttrsRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListByAttrsRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ListByAttrsResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ListByAttrsResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CreateRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CreateRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CreateResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CreateResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *BatchCreateRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *BatchCreateRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *BatchCreateResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *BatchCreateResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UpdateRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UpdateRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UpdateResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UpdateResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeleteResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeleteResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetHistoryRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetHistoryResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetHistoryResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetHiddenTxRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetHiddenTxRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetHiddenTxResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetHiddenTxResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *HideTxRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *HideTxRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *HideTxResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *HideTxResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UnHideTxRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UnHideTxRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *UnHideTxResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *UnHideTxResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetSuggestionRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetSuggestionRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetSuggestionResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetSuggestionResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionListRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionListRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionListResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionListResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionListByCollectionRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionListByCollectionRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionListByCollectionResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionListByCollectionResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionListByItemRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionListByItemRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionListByItemResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionListByItemResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionByPartialKeyRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionByPartialKeyRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionByPartialKeyResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionByPartialKeyResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionCreateRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionCreateRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionCreateResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionCreateResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionDeleteRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionDeleteRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionDeleteResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionDeleteResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionApproveRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionApproveRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SuggestionApproveResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SuggestionApproveResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}
