package state

import (
	"encoding/json"

	"github.com/nova38/thesis/packages/saacs/common"
	authpb "github.com/nova38/thesis/packages/saacs/gen/auth/v1"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"
)

func (l RawLedger[T]) PrimaryCreate(ctx common.TxCtxInterface, obj T) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	key, err := common.MakePrimaryKey(obj)
	if err != nil {
		return oops.Wrap(err)
	}

	if bytes, err := json.Marshal(obj); err != nil {
		return oops.Hint("Failed To Marshal").Wrap(err)
	} else {
		if err := ctx.GetStub().PutState(key, bytes); err != nil {
			return oops.With("key", key).Wrap(err)
		}
	}

	if ctx.EnabledHidden() {
		hiddenKey, err := common.MakeHiddenKey(obj)
		if err != nil {
			return oops.Wrap(err)
		}

		hiddenBytes, err := json.Marshal(&authpb.HiddenTxList{
			PrimaryKey: obj.ItemKey(),
			Txs:        []*authpb.HiddenTx{},
		})

		if err != nil {
			return oops.Wrap(err)
		}

		if err = ctx.GetStub().PutState(hiddenKey, hiddenBytes); err != nil {
			return oops.With(
				"Key", hiddenKey,
				"ItemType", obj.ItemType(),
			).Wrap(err)
		}
	}

	return nil
}

func UnmarshalPrimary[T common.ItemInterface](bytes []byte, obj T) (err error) {
	return json.Unmarshal(bytes, obj)
}
func UnmarshalNewPrimary[T common.ItemInterface](bytes []byte, base T) (item T, err error) {
	item, ok := proto.Clone(base).(T)
	if !ok {
		return item, oops.Errorf("Failed to clone")
	}
	proto.Reset(item)

	if err = json.Unmarshal(bytes, item); err != nil {
		return item, oops.Wrap(err)
	}
	return item, nil
}

// PrimaryExists returns true if the item exists in the ledger
func PrimaryExists[T common.ItemInterface](ctx common.TxCtxInterface, obj T) bool {

	key := lo.Must(common.MakePrimaryKey(obj))
	return common.KeyExists(ctx, key)
}
