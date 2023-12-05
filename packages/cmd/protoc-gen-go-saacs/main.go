package main

import (
	"github.com/nova38/thesis/packages/cmd/protoc-gen-go-saacs/generators"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	fileGenerators := []generators.FileGenerator{
		&generators.KeyGenerator{},
		// &generators.DiffGenerator{},
		&generators.ServiceGenerator{},
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
