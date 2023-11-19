package state

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Generic structures

// TODO: These need to be changed to utilize the new authpb.Collection
//       and authpb.CollectionObject interfaces.

type (

	//CollectionObjectKey interface {
	//}
	//
	//CollectionObject interface {
	//	GetId() *CollectionObjectKey
	//}

	FullObject[T common.ObjectInterface] struct {
		Object      T                    `json:"object"`
		Suggestions []*authpb.Suggestion `json:"suggestions"`
		History     *authpb.History      `json:"history"`
	}

	ObjectList[T common.ObjectInterface] struct {
		Entries []T `json:"entries"`
	}

	HistoryEntry[T common.ObjectInterface] struct {
		TxId      string                 `json:"txId"`
		Timestamp *timestamppb.Timestamp `json:"timestamp"`
		IsDelete  bool                   `json:"isDelete"`
		IsHidden  bool                   `json:"isHidden"`
		State     T                      `json:"object"`
	}

	HistoryList[T common.ObjectInterface] struct {
		Entries []*HistoryEntry[T] `json:"entries"`
	}
)
