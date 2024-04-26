package generators

import (
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func TransactionTypeOptions(m *protogen.Method) *pb.TransactionType {
	v, ok := proto.GetExtension(m.Desc.Options(), pb.E_TransactionType).(*pb.TransactionType)
	if !ok {
		return nil
	}
	return v
}
