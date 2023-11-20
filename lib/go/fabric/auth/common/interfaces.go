package common

import (
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/proto"
)

type (
	ServerConfig struct {
		CCID    string
		Address string
	}

	ObjectInterface interface {
		KeyAttr() (attr []string)
		ObjectKey() *authpb.ObjectKey
		Namespace() string
		proto.Message
	}

	PrimaryObjectInterface interface {
		ObjectInterface
		IsPrimary() bool
	}

	SubObjectInterface interface {
		ObjectInterface
		IsSecondary() bool
	}

	GlobalObjectInterface interface {
		ObjectInterface
		IsGlobal() bool
	}
)
