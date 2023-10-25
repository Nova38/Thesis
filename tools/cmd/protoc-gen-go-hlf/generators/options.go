package generators

import (
	"github.com/nova38/thesis/lib/go/gen/hlf"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func DataFieldOptions(m *protogen.Message) *hlf.DataField {
	v, ok := proto.GetExtension(m.Desc.Options(), hlf.E_DataField).(*hlf.DataField)
	if !ok {
		return nil
	}
	return v
}

func TransactionTypeOptions(m *protogen.Method) *hlf.TransactionType {
	v, ok := proto.GetExtension(m.Desc.Options(), hlf.E_TransactionType).(*hlf.TransactionType)
	if !ok {
		return nil
	}
	return v
}
