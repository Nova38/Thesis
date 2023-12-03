package generators

import (
	auth_pb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func TransactionTypeOptions(m *protogen.Method) *auth_pb.TransactionType {
	v, ok := proto.GetExtension(m.Desc.Options(), auth_pb.E_TransactionType).(*auth_pb.TransactionType)
	if !ok {
		return nil
	}
	return v
}
