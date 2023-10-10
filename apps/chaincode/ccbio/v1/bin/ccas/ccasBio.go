package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/nova38/thesis/apps/chaincode/ccbio/v1/context"
	schema "github.com/nova38/thesis/lib/gen/go/ccbio/schema/v1"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hyperledger/fabric-chaincode-go/shim"

	_ "github.com/samber/lo"
	"github.com/samber/oops"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
	"github.com/nova38/thesis/apps/chaincode/ccbio/v1/contracts"
)

type ServerConfig struct {
	CCID    string
	Address string
}

func makeBioContract() *contracts.SpecimenContract {
	contract := new(contracts.SpecimenContract)
	contract.TransactionContextHandler = &context.AuthTxContext{}
	contract.BeforeTransaction = context.HandelBefore

	contract.Info = metadata.InfoMetadata{
		Description: "",
		Title:       "Biochain Chaincode",
		Contact: &metadata.ContactMetadata{
			Name:  "Thomas Atkins",
			URL:   "https://biochain.iitc.ku.edu",
			Email: "tom@ku.edu",
		},
		License: &metadata.LicenseMetadata{
			Name: "MIT",
			URL:  "https://example.com",
		},
		Version: "latest",
	}

	contract.Name = "thomas.BioContract"

	return contract
}

func makeAuthContract() *contracts.AuthContract {
	contract := new(contracts.AuthContract)
	contract.TransactionContextHandler = &context.AuthTxContext{}
	contract.BeforeTransaction = context.HandelBefore

	contract.Info = metadata.InfoMetadata{
		Description: "",
		Title:       "Biochain Chaincode",
		Contact: &metadata.ContactMetadata{
			Name:  "Thomas Atkins",
			URL:   "https://biochain.iitc.ku.edu",
			Email: "tom@ku.edu",
		},
		License: &metadata.LicenseMetadata{
			Name: "MIT",
			URL:  "https://example.com",
		},
		Version: "latest",
	}

	contract.Name = "thomas.AuthContract"

	return contract
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

	specimenContract := contracts.BuildSpecimenContract()
	specimenContract.TransactionContextHandler = &context.AuthTxContext{}
	authContract := contracts.BuildAuthContract()
	authContract.TransactionContextHandler = &context.AuthTxContext{}

	sm, err := contractapi.NewChaincode(specimenContract, authContract)
	if err != nil {
		fmt.Printf("Error creating BioChain contract: %s", err)
		panic(err)
	}
	sm.Info.Title = "CCBIO"
	sm.DefaultContract = specimenContract.Name

	server := &shim.ChaincodeServer{
		CCID:    config.CCID,
		Address: config.Address,
		CC:      sm,
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}

	fmt.Print("Starting BioChain Chaincode Server")
	fmt.Printf("config: %+v\n", config)
	fmt.Printf("server: %+v\n", server)

	if err := server.Start(); err != nil {
		slog.Error("Failed to start", err)
	}
}

func print_empty() {
	out := &schema.LastModified{
		UserId: &schema.User_Id{
			MspId: "",
			Id:    "",
		},
		UserName: "",
		TxId:     "",
		UpdatedAt: &timestamppb.Timestamp{
			Seconds: 0,
			Nanos:   0,
		},
	}
	// encode it in json
	json, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}

func FormatTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		a.Value = slog.StringValue(a.Value.Time().Format(time.Kitchen))
		return a
	}
	return a
}

func main() {
	// logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	// slog.SetDefault(logger)

	print_empty()
	// attr := &map[string]string{"time": ""}
	opts := slog.HandlerOptions{
		Level:       slog.LevelDebug,
		AddSource:   true,
		ReplaceAttr: FormatTime,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)

	oops.SourceFragmentsHidden = false

	runChaincode()
}
