package state

import (
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Generic structures

// TODO: These need to be changed to utilize the new authpb.Collection
//       and authpb.CollectionItem interfaces.

type (
	FullItem[T common.ItemInterface] struct {
		Item        T                    `json:"item"`
		Suggestions []*authpb.Suggestion `json:"suggestions"`
		History     *authpb.History      `json:"history"`
	}

	ItemList[T common.ItemInterface] struct {
		Entries []T `json:"entries"`
	}

	HistoryEntry[T common.ItemInterface] struct {
		TxId      string                 `json:"txId"`
		Timestamp *timestamppb.Timestamp `json:"timestamp"`
		IsDelete  bool                   `json:"isDelete"`
		IsHidden  bool                   `json:"isHidden"`
		State     T                      `json:"item"`
	}

	HistoryList[T common.ItemInterface] struct {
		Entries []*HistoryEntry[T] `json:"entries"`
	}
)
