module github.com/nova38/saacs/cmd/protoc-gen-saacs-go

go 1.21

replace github.com/nova38/saacs/lib/saacs-protos-go => ../../lib/saacs-protos-go/

require (
	github.com/google/go-cmp v0.6.0
	github.com/mennanov/fieldmask-utils v1.1.2
	github.com/mennanov/fmutils v0.3.0
	github.com/nova38/saacs/lib/saacs-protos-go v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.14.0
	google.golang.org/protobuf v1.33.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.33.0-20240221180331-f05a6f4403ce.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
)
