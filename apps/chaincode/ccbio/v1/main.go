package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/nova38/thesis/apps/chaincode/ccbio/v1/context"
	schema "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v1"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hyperledger/fabric-chaincode-go/shim"

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

func runChaincodeCCAS() {
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

// func example() {
// 	item := &schema.Specimen{
// 		Id: &schema.Specimen_Id{
// 			CollectionId: "",
// 			Id:           "",
// 		},
// 		Primary:      nil,
// 		Secondary:    nil,
// 		Taxon:        nil,
// 		Georeference: nil,
// 		Images:       nil,
// 		Loans:        "",
// 		Grants:       "",
// 		HiddenTxs:    nil,
// 		LastModified: nil,
// 	}

// 	fmt.Printf("%+v\n", item)

// 	// name := &schema.E_Namespace{}

// 	// Print out the value of the protobuf message option called namespace for the type specimen (which is a message)
// 	// in the file schema/defs.proto
// 	d := schema.E_Namespace.TypeDescriptor()
// 	name := item.ProtoReflect().Descriptor().Options().ProtoReflect().Get(d)

// 	fmt.Printf("%+v\n", name)

// }

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

func runChaincode() {
	fmt.Println("Starting BioChain")

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

	if err := sm.Start(); err != nil {
		slog.Error("Failed to start", err)
	}

	// server := &shim.ChaincodeServer{
	// 	CCID:    config.CCID,
	// 	Address: config.Address,
	// 	CC:      sm,
	// 	TLSProps: shim.TLSProperties{
	// 		Disabled: true,
	// 	},
	// }

	// fmt.Print("Starting BioChain Chaincode Server")
	// fmt.Printf("config: %+v\n", config)
	// fmt.Printf("server: %+v\n", server)

	// if err := server.Start(); err != nil {
	// 	slog.Error("Failed to start", err)
	// }
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
	// opts.ReplaceAttr([]string{"time"})
	// handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)

	oops.SourceFragmentsHidden = false
	// zerolog.SetGlobalLevel(zerolog.DebugLevel)
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
	// 	short := file
	// 	for i := len(file) - 1; i > 0; i-- {
	// 		if file[i] == '/' {
	// 			short = file[i+1:]
	// 			break
	// 		}
	// 	}
	// 	file = short
	// 	return file + ":" + strconv.Itoa(line)
	// }
	// log.Logger = log.With().Caller().Logger()

	// example()
	runChaincode()
}
