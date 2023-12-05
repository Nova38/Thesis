package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/packages/saacs/common"
	"github.com/nova38/thesis/packages/saacs/contracts/identity"
	"github.com/nova38/thesis/packages/saacs/contracts/noauth"
	"github.com/nova38/thesis/packages/saacs/contracts/roles"
)

var (
	config common.ServerConfig
	sm     *contractapi.ContractChaincode
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	config = common.ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	mode := os.Getenv("AUTH_MODE")
	if mode == "" {
		authMode := flag.String(
			"auth",
			"noauth",
			"auth mode: noauth, noauth-no-sub, roles, or identity",
		)
		if authMode == nil {
			mode = "noauth"
		} else {
			mode = *authMode
		}
	}

	switch mode {
	case "noauth":
		slog.Info("Using NoAuth Contract")
		sm = noauth.BuildContract()
	case "noauth-no-sub":
		slog.Info("Using NoAuthNoSub Contract")
		sm = noauth.NoSubBuildContract()
	case "roles":
		slog.Info("Using Roles Contract")
		sm = roles.BuildContract()
	case "identity":
		slog.Info("Using Identity Contract")
		sm = identity.BuildContract()
	default:

	}

}

func main() {

	slog.Info("Starting Chaincode")

	if name := os.Getenv("CHAINCODE_NAME"); name != "" {
		sm.Info.Contact.Name = name
	}

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      sm,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	slog.Info("Starting Chaincode Server",
		"config", config,
		"server", server,
	)

	if err := server.Start(); err != nil {
		slog.Error("Failed to start", err)
	}
}
