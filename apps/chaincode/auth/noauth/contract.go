package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	// _ "net/http/pprof"

	// "github.com/grafana/pyroscope-go"
	// _ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof" // add this line as well

	"github.com/charmbracelet/log"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	common "github.com/nova38/thesis/lib/go/fabric/auth/common"
	"github.com/nova38/thesis/lib/go/fabric/auth/contracts"
	"github.com/nova38/thesis/lib/go/fabric/auth/serializer"
	"github.com/nova38/thesis/lib/go/fabric/auth/state"
	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	"github.com/samber/lo"

	"github.com/samber/oops"
)

// ═════════════════════════════════════════════
// Transaction Context
// ═════════════════════════════════════════════
var _ common.TxCtxInterface = (*NoAuthTxCtx)(nil)

type NoAuthTxCtx struct {
	state.BaseTxCtx
}

func Authenticate(ctx common.TxCtxInterface, ops []*authpb.Operation) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Authenticate")

	for _, op := range ops {
		ctx.GetLogger().Info(op.String())

		// policy.ValidateOperation()
	}

	return true, nil
}

// ═════════════════════════════════════════════
// Contract
// ═════════════════════════════════════════════

func iError() error {
	return oops.Errorf("Error")
}

// NoAuthContract is the contract for the NoAuth chaincode

type NoAuthContract struct {
	contractapi.Contract
}

func (c *NoAuthContract) Test(ctx common.TxCtxInterface) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.Test")

	return true, nil
}

func (c *NoAuthContract) TestFail(ctx common.TxCtxInterface) (bool, error) {
	ctx.GetLogger().Info("NoAuthContract.TestFail")
	e := iError()
	b := lo.Must(json.MarshalIndent(e, "", "  "))
	ctx.GetLogger().Info(string(b))

	return false, e
}

func BeforeTransaction(ctx common.TxCtxInterface) (err error) {
	defer func() { ctx.HandleFnError(&err, recover()) }()

	if err = ctx.HandelBefore(); err != nil {
		return oops.Wrap(err)
	}

	// Set the authenticator handler
	ctx.SetAuthenticator(Authenticate)

	return nil
}

// ═════════════════════════════════════════════
// Before Transaction
// ═════════════════════════════════════════════

// ═════════════════════════════════════════════
// Main
// ═════════════════════════════════════════════

func startPyrosope() {
}

func main() {
	// fmt.Println("Starting http for pprof")
	// go func() {
	// 	fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	// }()
	// appName := fmt.Sprintf("%s-%s", ccid, "noauth")

	// v, err := pyroscope.Start(pyroscope.Config{
	// 	ApplicationName: "noauth",

	// 	// replace this with the address of pyroscope server
	// 	ServerAddress: "http://pyroscope:4040",

	// 	// you can disable logging by setting this to nil
	// 	Logger: pyroscope.StandardLogger,

	// 	// you can provide static tags via a map:
	// 	Tags: map[string]string{"hostname": os.Getenv("HOSTNAME")},

	// 	ProfileTypes: []pyroscope.ProfileType{
	// 		// these profile types are enabled by default:
	// 		pyroscope.ProfileCPU,
	// 		pyroscope.ProfileAllocObjects,
	// 		pyroscope.ProfileAllocSpace,
	// 		pyroscope.ProfileInuseObjects,
	// 		pyroscope.ProfileInuseSpace,

	// 		// these profile types are optional:
	// 		pyroscope.ProfileGoroutines,
	// 		pyroscope.ProfileMutexCount,
	// 		pyroscope.ProfileMutexDuration,
	// 		pyroscope.ProfileBlockCount,
	// 		pyroscope.ProfileBlockDuration,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println("Error starting Pyroscope", err)
	// }
	// defer v.Stop()

	fmt.Println("Starting BioChain")

	handler := log.New(os.Stderr)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	config := common.ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	auth := new(NoAuthContract)

	generic := new(contracts.ItemContractImpl)

	auth.Name = "auth.none"

	generic.Name = "auth.generic"

	auth.TransactionContextHandler = &NoAuthTxCtx{}

	generic.TransactionContextHandler = &NoAuthTxCtx{}

	auth.BeforeTransaction = BeforeTransaction

	generic.BeforeTransaction = BeforeTransaction

	sm, err := contractapi.NewChaincode(auth, generic)
	if err != nil {
		fmt.Printf("Error creating No Auth contract: %s", err)
		panic(err)
	}

	sm.TransactionSerializer = &serializer.TxSerializer{}

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
