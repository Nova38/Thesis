package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	base "github.com/nova38/thesis/lib/go/fabric/auth/common/base"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/oops"
)

// =============================================
// Transaction Context
// =============================================
var _ state.TxCtxInterface = (*NoAuthTxCtx)(nil)

type NoAuthTxCtx struct {
	state.BaseTxCtx
}

func Authenticate(ctx state.TxCtxInterface, ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())
	}

	return true, nil
}

// =============================================
// Contract
// =============================================

// NoAuthContract is the contract for the NoAuth chaincode

type NoAuthContract struct {
	contractapi.Contract
}

func (c *NoAuthContract) Test(ctx state.TxCtxInterface) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Test")

	return true, nil
}

func BeforeTransaction(ctx state.TxCtxInterface) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	// Set the authenticator handler
	ctx.SetAuthenticator(Authenticate)

	return nil
}

// =============================================
// Before Transaction
// =============================================

// =============================================
// Main
// =============================================

func main() {
	fmt.Println("Starting BioChain")

	config := common.ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	auth := new(NoAuthContract)
	users := new(base.UserImpl)
	collections := new(base.CollectionImpl)

	auth.Name = "auth.none"
	users.Name = "auth.users"
	collections.Name = "auth.collections"

	auth.TransactionContextHandler = &NoAuthTxCtx{}
	users.TransactionContextHandler = &NoAuthTxCtx{}
	collections.TransactionContextHandler = &NoAuthTxCtx{}

	auth.BeforeTransaction = BeforeTransaction
	users.BeforeTransaction = BeforeTransaction
	collections.BeforeTransaction = BeforeTransaction

	sm, err := contractapi.NewChaincode(auth, users, collections)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      sm,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	fmt.Print("Starting Auth Chaincode Server")
	fmt.Printf("config: %+v\n", config)
	fmt.Printf("server: %+v\n", server)

	if err := server.Start(); err != nil {
		slog.Error("Failed to start", err)
	}
}
