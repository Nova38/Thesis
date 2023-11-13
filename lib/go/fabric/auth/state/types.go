package state

import (
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
