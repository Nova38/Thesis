// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: rbac/examples/simple.proto

package examples

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *SimpleStateObject) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SimpleStateObject) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SimpleStateObject_Id) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SimpleStateObject_Id) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *SimpleStateObject_List) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *SimpleStateObject_List) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}