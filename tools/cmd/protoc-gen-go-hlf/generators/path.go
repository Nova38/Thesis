package generators

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type DiffGenerator struct{}

func (d *DiffGenerator) GenerateFile(
	gen *protogen.Plugin,
	file *protogen.File,
) (*protogen.GeneratedFile, error) {
	filename := file.GeneratedFilenamePrefix + ".cc.path.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	// Generate the header
	g.P("// Code generated by protoc-gen-go-hlf. DO NOT EDIT.")
	g.P("// versions:")

	g.P("// - protoc-gen-cckey v0.0.1")
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, msg := range file.Messages {
		d.GenerateMessage(g, msg)
	}

	return g, nil
}

// Handle the messages
func (d *DiffGenerator) GenerateMessage(
	g *protogen.GeneratedFile,
	msg *protogen.Message,
) {
	// Skip map entries as we can't directly compare them
	if msg.Desc.IsMapEntry() {
		return
	}

	for _, m := range msg.Messages {
		d.GenerateMessage(g, m)
	}

	// Start the function
	g.P(
		"func (m *", msg.GoIdent.GoName, ") ",
		"DiffPath(other *", msg.GoIdent.GoName, ")",
		"(updated []string, all bool) {",
	)
	g.P()
	g.P("all = true")
	defer func() { g.P("return updated, all"); g.P("}"); g.P() }()

	for _, f := range msg.Fields {
		d.handelField(g, f)
	}

	g.P()

	g.P()
}

// https://github.com/planetscale/vtprotobuf/blob/5a1a54cdaaaeaeed916c09300c7cfdbb9b9bb051/features/equal/equal.go#L167
func (d *DiffGenerator) handelField(
	g *protogen.GeneratedFile,
	f *protogen.Field,
) {
	g.P("// ", f.Desc.Name(), ": is a ", f.Desc.Kind())

	// Check if the field is a scalar

	// check if the field is a list
	if f.Desc.IsList() {
		g.P("// TODO: Handle lists")
	} else if f.Desc.IsMap() {
		// TODO: Handle maps
		g.P("// TODO: Handle maps")
	} else if isScalar(f.Desc.Kind()) {
		// Check if the field is the same
		g.P("if m.", f.GoName, " != other.", f.GoName, " {")
		g.P("\tupdated = append(updated, \"", f.Desc.Name(), "\")")
		// TODO: Update the message
		g.P("} else { all = false }")

	} else if f.Desc.Kind() == protoreflect.MessageKind {

		// Check if the field is a well known type
		if d.wktLookup(f.Desc.Message().FullName()) {
			return
		}
		// Compute the diff then add the prefix to the updated list
		// check if the fields are not nil
		g.P("if m.", f.GoName, " != nil || other.", f.GoName, " != nil {")

		g.P("updated_", f.GoName, ", all_", f.GoName, " := m.", f.GoName, ".DiffPath(other.", f.GoName, ")")
		// Check if the number of updated fields is greater than 0
		g.P("if len(updated_", f.GoName, ") > 0 {")
		// check if all the fields are updated
		g.P("if all_", f.GoName, " {")
		g.P("\tupdated = append(updated, \"", f.Desc.Name(), "\")")
		g.P("} else {")
		g.P("\tfor _, u := range updated_", f.GoName, " {")
		g.P("\t\tupdated = append(updated, \"", f.Desc.Name(), ".", "\" + u)")
		g.P("\t}")
		g.P("}")
		g.P("} else { all = false }")

		g.P("} else { all = false }")
	}
}

// https://github.com/planetscale/vtprotobuf/blob/5a1a54cdaaaeaeed916c09300c7cfdbb9b9bb051/features/equal/equal.go#L272C1-L285C2
func isScalar(kind protoreflect.Kind) bool {
	switch kind {
	case
		protoreflect.BoolKind,
		protoreflect.StringKind,
		protoreflect.DoubleKind, protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind,
		protoreflect.FloatKind, protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Uint64Kind, protoreflect.Sint64Kind,
		protoreflect.Int32Kind, protoreflect.Uint32Kind, protoreflect.Sint32Kind,
		protoreflect.EnumKind:
		return true
	}
	return false
}

func (d *DiffGenerator) wktLookup(t protoreflect.FullName) bool {
	switch t {
	case
		"google.protobuf.Timestamp",
		"google.protobuf.Duration",
		"google.protobuf.Any",
		"google.protobuf.Struct",
		"google.protobuf.Value",
		"google.protobuf.FieldMask",
		"google.protobuf.Type",
		"google.protobuf.ListValue":
		return true
	}
	return false
}
