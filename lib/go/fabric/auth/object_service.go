package auth

import (
	"github.com/golang/protobuf/proto"
	"github.com/nova38/thesis/lib/go/fabric/state"
)

type ProtoStateObject interface {
	state.Object
	proto.Message
}

type IStateObjectService interface{}
