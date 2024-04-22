module github.com/nova38/saacs/cmd/protoc-gen-saacs-go

go 1.22.2

//replace github.com/nova38/saacs/pkg/saacs-protos => ./../../pkg/saacs-protos/

//
require (
	github.com/google/go-cmp v0.6.0
	github.com/mennanov/fieldmask-utils v1.1.2
	github.com/mennanov/fmutils v0.3.0
	golang.org/x/text v0.14.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	google.golang.org/genproto v0.0.0-20231016165738-49dd2c1f3d0b // indirect
)
