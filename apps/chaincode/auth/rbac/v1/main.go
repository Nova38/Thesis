package main

// type ServerConfig struct {
//	CCID    string
//	Address string
//}

// func runChaincode() {
//	fmt.Println("Starting BioChain")

// config := ServerConfig{
// 	CCID:    os.Getenv("CHAINCODE_ID"),
// 	Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
// }

// bioContract := makeBioContract()
// authContract := makeAuthContract()

// contracts.SpecimenContract

// authContract := new(contracts.RbacContractImpl)
// authContract.TransactionContextHandler = &contracts.AuthTxCtx{}

// sm, err := contractapi.NewChaincode(authContract)
// if err != nil {
// 	fmt.Printf("Error creating BioChain contract: %s", err)
// 	panic(err)
// }

// server := &shim.ChaincodeServer{
// 	CCID:    config.CCID,
// 	Address: config.Address,
// 	CC:      sm,
// 	TLSProps: shim.TLSProperties{
// 		Disabled: true,
// 	},
// }

// fmt.Print("Starting Auth Chaincode Server")
// fmt.Printf("config: %+v\n", config)
// fmt.Printf("server: %+v\n", server)

// if err := server.Start(); err != nil {
// 	slog.Error("Failed to start", err)
// }
// }
