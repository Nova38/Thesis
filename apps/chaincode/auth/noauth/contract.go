package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/nova38/thesis/packages/saacs/auth/common"
	"github.com/nova38/thesis/packages/saacs/contracts/noauth"
)

// _ "net/http/pprof"

// "github.com/grafana/pyroscope-go"
// _ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof" // add this line as well

// ═════════════════════════════════════════════
// Contract
// ═════════════════════════════════════════════

// ═════════════════════════════════════════════
// Before Transaction
// ═════════════════════════════════════════════

// ═════════════════════════════════════════════
// Main
// ═════════════════════════════════════════════

func startPyrosope() {
}

func mainPro() {
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

	// handler := log.New(os.Stderr)
	// logger := slog.New(handler)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	config := common.ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	sm := noauth.BuildContract()

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
	// auth := new(NoAuthContract)

	// generic := new(contracts.ItemContractImpl)

	// auth.Name = "auth.none"

	// generic.Name = "auth.generic"

	// auth.TransactionContextHandler = &NoAuthTxCtx{}

	// generic.TransactionContextHandler = &NoAuthTxCtx{}

	// auth.BeforeTransaction = BeforeTransaction

	// generic.BeforeTransaction = BeforeTransaction

	// sm, err := contractapi.NewChaincode(auth, generic)
	// if err != nil {
	// 	fmt.Printf("Error creating No Auth contract: %s", err)
	// 	panic(err)
	// }

	// sm.TransactionSerializer = &serializer.TxSerializer{}

}
