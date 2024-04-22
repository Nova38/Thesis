package identity

import (
	"fmt"
	"log/slog"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	"github.com/nova38/saacs/pkg/saacs-cc/serializer"
	"github.com/nova38/saacs/pkg/saacs-cc/state"
	v1 "github.com/nova38/saacs/pkg/saacs-protos/auth/v1"
	cc "github.com/nova38/saacs/pkg/saacs-protos/chaincode/common"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func BeforeTransaction(ctx common.TxCtxInterface) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	// Set the authenticator handler

	return nil
}

func BuildContract() *contractapi.ContractChaincode {
	contract := new(IdentiyContract)
	contract.BeforeTransaction = BeforeTransaction
	contract.TransactionContextHandler = new(TxCtx)

	sm, err := contractapi.NewChaincode(contract)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}
	sm.TransactionSerializer = &serializer.TxSerializer{}

	return sm
}

func (c *IdentiyContract) Bootstrap(
	ctx common.TxCtxInterface,
	req *cc.BootstrapRequest,
) (res *cc.BootstrapResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	ctx.GetLogger().Info("NoAuthContract.Boostrap")
	if err = ctx.Validate(req); err != nil {
		ctx.LogError(err)
		return nil, oops.Wrap(err)
	}

	// Check to see if the bootstrap has already been done
	// If so, return an error
	if bootstraped, err := ctx.CheckBootstrap(); err != nil {
		return nil, oops.Wrap(err)
	} else if bootstraped {
		ctx.GetLogger().Warn("Bootstrap already done")
		return nil, oops.Errorf("Bootstrap already done")
	}

	for _, col := range req.GetCollections() {
		lo.Must0(c.CreateCollection(ctx, &cc.CreateCollectionRequest{Collection: col}))

	}

	return &cc.BootstrapResponse{Success: true}, nil
}

func (c *IdentiyContract) CreateCollection(
	ctx common.TxCtxInterface,
	req *cc.CreateCollectionRequest,
) (res *cc.CreateCollectionResponse, err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()
	ctx.GetLogger().Info("RoleContract.CreateCollection")
	col := req.GetCollection()
	user, err := ctx.GetUserId()
	if err != nil {
		return nil, oops.Wrap(err)
	}

	if state.Exists(ctx, col) {
		return nil, oops.Errorf("Collection already exists")
	}

	col.AuthType = v1.AuthType_AUTH_TYPE_IDENTITY

	// Make Sure the auth types are in the collection
	authTypes := []string{
		common.CollectionItemType,
		common.UserDirectMembershipItemType,
	}
	col.ItemTypes = append(col.GetItemTypes(), authTypes...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes()) // Deduplicate the item types

	// Exclude the auth types from the default policy
	// col.Default.DefaultExcludedTypes = append(
	// 	col.GetDefault().GetDefaultExcludedTypes(),
	// 	authTypes...)
	col.Default.DefaultExcludedTypes = lo.Uniq(col.GetDefault().GetDefaultExcludedTypes())

	// Add the auth types to the collection
	membership := &v1.UserDirectMembership{
		CollectionId: col.GetCollectionId(),
		MspId:        user.GetMspId(),
		UserId:       user.GetUserId(),
		Polices: &v1.Polices{
			ItemPolicies: map[string]*v1.PathPolicy{},
			DefaultPolicy: &v1.PathPolicy{
				Path:          "",
				FullPath:      "",
				AllowSubPaths: false,
				Actions: []v1.Action{
					v1.Action_ACTION_UTILITY,
					v1.Action_ACTION_VIEW,
					v1.Action_ACTION_CREATE,
					v1.Action_ACTION_UPDATE,
					v1.Action_ACTION_DELETE,
					v1.Action_ACTION_SUGGEST_VIEW,
					v1.Action_ACTION_SUGGEST_CREATE,
					v1.Action_ACTION_SUGGEST_DELETE,
					v1.Action_ACTION_SUGGEST_APPROVE,
					v1.Action_ACTION_VIEW_HISTORY,
					v1.Action_ACTION_VIEW_HIDDEN_TXS,
					v1.Action_ACTION_HIDE_TX,
					v1.Action_ACTION_REFERENCE_CREATE,
					v1.Action_ACTION_REFERENCE_DELETE,
					v1.Action_ACTION_REFERENCE_VIEW,
				},
			},
		},
	}

	col.ItemTypes = append(col.GetItemTypes(), []string{membership.ItemType()}...)
	col.ItemTypes = lo.Uniq(col.GetItemTypes())

	// state.Put(ctx, col)
	if err = (state.Ledger[*v1.Collection]{}.PrimaryCreate(ctx, col)); err != nil {
		return nil, oops.Wrap(err)
	}
	if err = (state.Ledger[*v1.UserDirectMembership]{}.PrimaryCreate(ctx, membership)); err != nil {
		return nil, oops.Wrap(err)
	}

	ctx.GetLogger().Info("Bootstraping", slog.Any("membership", membership))

	return &cc.CreateCollectionResponse{Collection: col}, nil

}
