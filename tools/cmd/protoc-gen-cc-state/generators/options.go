package generators

import (
	"github.com/nova38/thesis/lib/go/gen/key"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func KeySchemaOptions(m *protogen.Message) *key.KeySchema {
	v, ok := proto.GetExtension(m.Desc.Options(), key.E_KeySchema).(*key.KeySchema)
	if !ok {
		return nil
	}
	return v
}

func DataFieldOptions(m *protogen.Message) *key.DataField {
	v, ok := proto.GetExtension(m.Desc.Options(), key.E_DataField).(*key.DataField)
	if !ok {
		return nil
	}
	return v
}
