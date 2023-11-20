package main

import (
	"fmt"
	"os"
)

type ServerConfig struct {
	CCID    string
	Address string
}

var config = ServerConfig{
	CCID:    os.Getenv("CHAINCODE_ID"),
	Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
}

func main() {
	// auth := rbac.NewAuthContract("ccbio")

	// cc, err := contractapi.NewChaincode(
	// 	auth,
	// )
	// if err != nil {
	// 	fmt.Printf("Error creating BioChain contract: %s", err)
	// 	panic(err)
	// }

	// server := &shim.ChaincodeServer{
	// 	CCID:    config.CCID,
	// 	Address: config.Address,
	// 	CC:      cc,
	// 	TLSProps: shim.TLSProperties{
	// 		Disabled: true,
	// 	},
	// }

	fmt.Print("Starting BioChain Chaincode Server")
	// fmt.Printf("config: %+v\n", config)
	// fmt.Printf("server: %+v\n", server)

	// if err := server.Start(); err != nil {
	// 	slog.Error("Failed to start", err)
	// }
}
