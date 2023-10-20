package generators

import "google.golang.org/protobuf/compiler/protogen"

type FileGenerator interface {
	GenerateFile(gen *protogen.Plugin, file *protogen.File) (*protogen.GeneratedFile, error)
}
