package state

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Generic structures

type (
	Object interface {
		Key() (attr []string, err error)
		Namespace() (namespace string)
	}
	ProtoObject interface {
		Object
		proto.Message
		GetHiddenTxs()
	}
)

type (
	ObjectList[T Object] struct {
		Entries []T `json:"entries"`
	}

	HistoryEntry[T Object] struct {
		TxId      string                 `json:"txId"`
		Timestamp *timestamppb.Timestamp `json:"timestamp"`
		IsDelete  bool                   `json:"isDelete"`
		IsHidden  bool                   `json:"isHidden"`
		State     T                      `json:"object"`
	}

	HistoryList[T Object] struct {
		Entries []*HistoryEntry[T] `json:"entries"`
	}
)
