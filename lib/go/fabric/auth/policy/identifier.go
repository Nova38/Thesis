package policy

import (
	"fmt"

	auth_pb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
)

const (
	fmtAttribute = "Attribute[<Name:%s><Value:%s><MspId:%s>]"
	fmtRole      = "Role[<Id:%s>]"
	fmtIdentity  = "Identity[<MspId:%s><Id:%s>]"
)

func GetACLKey(i *auth_pb.Identifier) (key string, err error) {
	switch x := i.GetId().(type) {
	case nil:
		return "", oops.Errorf("Identifier.Id is nil")
	case *auth_pb.Identifier_Attribute_:
		if x.Attribute != nil {
			return fmt.Sprintf(
				fmtAttribute,
				x.Attribute.Name,
				x.Attribute.Value,
				x.Attribute.MspId,
			), nil
		}
	case *auth_pb.Identifier_Role_:
		if x.Role != nil {
			return fmt.Sprintf(fmtRole, x.Role.Id), nil
		}

	case *auth_pb.Identifier_Identity_:
		if x.Identity != nil {
			return fmt.Sprintf(fmtIdentity, x.Identity.MspId, x.Identity.Id), nil
		}
	}

	return "", oops.Errorf("Failed to make ACL key")
}
