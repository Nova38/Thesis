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

	ItemInterface interface {
		KeyAttr() (attr []string)
		ItemKey() *authpb.ItemKey
		SetKeyAttr(attr []string)
		ItemType() string

		proto.Message
	}

	PrimaryItemInterface interface {
		ItemInterface
		IsPrimary() bool
	}

	SubItemInterface interface {
		ItemInterface
		IsSecondary() bool
	}

	GlobalItemInterface interface {
		ItemInterface
		IsGlobal() bool
	}
)
