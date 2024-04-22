module github.com/nova38/saacs/cmd/protoc-gen-saacs-go

go 1.22.2

//replace github.com/nova38/saacs/pkg/saacs-protos => ./../../pkg/saacs-protos/

//
require (
	github.com/google/go-cmp v0.6.0
	github.com/mennanov/fieldmask-utils v1.1.2
	github.com/mennanov/fmutils v0.3.0
	github.com/nova38/saacs/pkg/saacs-protos v0.0.0-20240328235441-37d593622f78
	golang.org/x/text v0.14.0
	google.golang.org/protobuf v1.33.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.33.0-20240401165935-b983156c5e99.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
)
