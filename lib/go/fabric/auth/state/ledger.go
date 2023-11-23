package state

import (
	"encoding/json"
	"log/slog"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/mennanov/fmutils"
	"github.com/nova38/thesis/lib/go/fabric/auth/common"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/samber/oops"
)

type Ledger[T common.ItemInterface] struct {
	ctx TxCtxInterface
}

// UTIL Functions

// Exists returns true if the item exists in the ledger
func (l *Ledger[T]) Exists(key string) bool {
	bytes, err := l.ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return false
	}

	return err == nil
}

// ════════════════════════════════════════════════════════
// Invoke Functions
// ════════════════════════════════════════════════════════

// Insert inserts the item into the ledger
// returns error if the item already exists
func (l *Ledger[T]) Create(obj T) (err error) {
	var (
		key   string
		bytes []byte
	)

	if key, err = MakePrimaryKey(obj); err != nil {
		return err
	}

	if l.Exists(key) {
		return oops.
			With("Key", key, "ItemType", obj.ItemType()).
			Wrap(common.AlreadyExists)
	}

	if bytes, err = json.Marshal(obj); err != nil {
		return err
	}

	return l.ctx.GetStub().PutState(key, bytes)
}

// Edit updates the item in the ledger
// returns error if the item does not exist
func (l *Ledger[T]) Update(update T, mask *fieldmaskpb.FieldMask) (err error) {
	var (
		key     string
		bytes   []byte
		current T
	)

	// Get the current item from the ledger
	if key, err = MakePrimaryKey(update); err != nil {
		return err
	}

	if current, err = l.GetFromKey(key); err != nil {
		return oops.Wrap(err)
	}

	// Apply the mask to the Updating item
	fmutils.Filter(update, mask.Paths)
	proto.Merge(current, update)

	// Put the item back into the ledger
	if bytes, err = json.Marshal(current); err != nil {
		return oops.Wrap(err)
	}

	update = current

	return l.ctx.GetStub().PutState(key, bytes)
}

// Delete deletes the item from the ledger
func (l *Ledger[T]) Delete(in T) (err error) {
	key, err := MakePrimaryKey(in)
	if err != nil {
		return err
	}

	if err = l.ctx.GetStub().DelState(key); err != nil {
		return oops.Wrap(err)
	}

	return nil
}

// ════════════════════════════════════════════════════════
// Query Functions
// ════════════════════════════════════════════════════════

func (l *Ledger[T]) GetFromKey(key string) (obj T, err error) {
	bytes, err := l.ctx.GetStub().GetState(key)
	if bytes == nil && err == nil {
		return obj, oops.
			With("Key", key, "ItemType", obj.ItemType()).
			Wrap(common.KeyNotFound)
	} else if err != nil {
		return obj, oops.Wrap(err)
	}

	if err = json.Unmarshal(bytes, obj); err != nil {
		return obj, oops.Wrap(err)
	}

	return obj, nil
}

// Get returns the item from the ledger
func (l *Ledger[T]) Get(in T) (err error) {
	var (
		key   string
		bytes []byte
	)

	itemtype := in.ItemType()
	l.ctx.GetLogger().Debug("fn: GetState", "ItemType", itemtype)

	if key, err = MakePrimaryKey(in); err != nil {
		return err
	}

	if bytes, err = l.ctx.GetStub().GetState(key); err == nil {
		return oops.
			With("Key", key, "ItemType", in.ItemType()).
			Wrap(common.AlreadyExists)
	}

	if err = json.Unmarshal(bytes, in); err != nil {
		return oops.Wrap(err)
	}
	return nil
}

// GetPartialKeyList returns a list of items of type T
// T must implement StateItem interface
// numAttr is the number of attributes in the key to search for
func (l *Ledger[T]) GetPartialKeyList(
	obj T,
	numAttr int,
	bookmark string,
) (list []T, mk string, err error) {
	// obj = []*T{}
	l.ctx.GetLogger().Info("GetPartialKeyList")

	var (
		itemtype = obj.ItemType()
		attr     = append([]string{itemtype}, obj.KeyAttr()...)
	)

	if len(attr) == 0 || len(attr) < numAttr {
		return nil, "", common.ItemInvalid
	}

	// Extract the attributes to search for
	attr = attr[:len(attr)-numAttr]

	l.ctx.GetLogger().
		Info("GetPartialKeyList",
			slog.Group(
				"Key", "ItemType", itemtype,
				slog.Int("numAttr", numAttr),
				slog.Any("attr", attr),
				slog.Group(
					"Paged",
					"Bookmark", bookmark,
					"PageSize", strconv.Itoa(int(l.ctx.GetPageSize())),
				),
			),
		)

	results, meta, err := l.ctx.GetStub().
		GetStateByPartialCompositeKeyWithPagination(
			obj.ItemKey().GetItemType(),
			attr,
			l.ctx.GetPageSize(),
			bookmark,
		)
	if err != nil {
		return nil, "", err
	}
	defer func(results shim.StateQueryIteratorInterface) {
		err := results.Close()
		if err != nil {
			l.ctx.GetLogger().Error("GetPartialKeyList", "Error", err)
		}
	}(results)

	for results.HasNext() {
		queryResponse, err := results.Next()
		if err != nil {
			return nil, "", err
		}
		obj := new(T)

		if err := json.Unmarshal(queryResponse.GetValue(), &obj); err != nil {
			return nil, "", err
		}

		list = append(list, *obj)
	}

	return list, meta.GetBookmark(), nil
}
