package main

import (
	"github.com/nova38/thesis/tools/cmd/protoc-gen-go-hlf/generators"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {

	fileGenerators := []generators.FileGenerator{
		&generators.KeyGenerator{},
		&generators.DiffGenerator{},
	}

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {

		for _, generator := range fileGenerators {
			for _, f := range gen.Files {
				if f.Generate {
					_, err := generator.GenerateFile(gen, f)
					if err != nil {
						return err
					}

				}
			}
		}

		return nil
	})
}
