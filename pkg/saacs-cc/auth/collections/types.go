package collections

import (
	"github.com/nova38/saacs/pkg/saacs-cc/common"
	authpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/auth/v0"
)

type Builder interface {
	CreateCollection(
		ctx common.TxCtxInterface,
		col *authpb.Collection,
	) (res *authpb.Collection, err error)
}
