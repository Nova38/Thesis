package generators

import (
	authpb "github.com/nova38/thesis/pkg/saacs/gen/auth/v1"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func TransactionTypeOptions(m *protogen.Method) *authpb.TransactionType {
	v, ok := proto.GetExtension(m.Desc.Options(), authpb.E_TransactionType).(*authpb.TransactionType)
	if !ok {
		return nil
	}
	return v
}
