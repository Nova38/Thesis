package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/apps/chaincode/auth/v1/contracts"
)

type ServerConfig struct {
	CCID    string
	Address string
}

func runChaincode() {
	fmt.Println("Starting BioChain")

	config := ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	// bioContract := makeBioContract()
	// authContract := makeAuthContract()

	// contracts.SpecimenContract

	authContract := new(contracts.AuthContractImpl)
	authContract.TransactionContextHandler = &contracts.AuthTxCtx{}

	sm, err := contractapi.NewChaincode(authContract)
	if err != nil {
		fmt.Printf("Error creating BioChain contract: %s", err)
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
