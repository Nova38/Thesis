package rbac

import (
	"github.com/samber/lo"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Write protobuf helpers to get all the fields of a protobuf message

func getFields(m proto.Message) []string {
	return GetFields(m.ProtoReflect().Descriptor())
}

func GetFields(m protoreflect.MessageDescriptor) []string {
	fields := []string{}

	fd0 := m.Fields()

	// iterate over the fields
	for i := range lo.Range(fd0.Len()) {
		// get the field descriptor
		fd := fd0.Get(i)

		switch {
		case fd.IsList():
		case fd.IsMap():
		case fd.Message() != nil:
			fields = lo.Union(fields, []string{fd.JSONName()}, GetFields(fd.Message()))
		case fd.Enum() != nil:

		}

		// get the field name
		name := string(fd.Name())
		// add the field name to the list
		fields = append(fields, name)
	}

	return fields
}
