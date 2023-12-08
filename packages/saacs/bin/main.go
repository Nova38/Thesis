package main

import (
	"flag"
	"log/slog"
	"os"
	"runtime"

	"net/http"
	_ "net/http/pprof"

	"github.com/grafana/pyroscope-go"
	_ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/nova38/thesis/packages/saacs/common"
	"github.com/nova38/thesis/packages/saacs/contracts/identity"
	"github.com/nova38/thesis/packages/saacs/contracts/noauth"
	"github.com/nova38/thesis/packages/saacs/contracts/roles"
)

var (
	config common.ServerConfig
	// protoscopePushEnabled = false
	// protoscopePullEnabled = true

	protoscopeMode = "PULL"
	authMode       = "noauth"
	sm             *contractapi.ContractChaincode
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Get the chaincode server config
	config = common.ServerConfig{
		CCID:    os.Getenv("CHAINCODE_ID"),
		Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
	}

	if e := os.Getenv("PROTOSCOPE"); e == "PUSH" {
		protoscopeMode = "PUSH"
	}

	// Set the mode
	authMode = os.Getenv("AUTH_MODE")
	if authMode == "" {
		mode := flag.String(
			"auth",
			"noauth",
			"auth mode: noauth, noauth-no-sub, roles, or identity",
		)
		if mode != nil {
			authMode = *mode
		}
	}

	switch authMode {
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

	slog.Info("Using Auth Mode", "mode", authMode)
	slog.Info("Using Protoscope Mode", "mode", protoscopeMode)

}

func main() {
	if protoscopeMode == "PULL" {
		slog.Info("Starting Profiling Server")
		go func() {
			err := http.ListenAndServe(":6060", nil)
			if err != nil {
				slog.Error("Failed to start profiling server", err)
			}
		}()
	} else if protoscopeMode == "PUSH" {

		rate := 5
		runtime.SetBlockProfileRate(rate)
		runtime.SetMutexProfileFraction(rate)

		slog.Info("Starting Protoscope")
		// These 2 lines are only required if you're using mutex or block profiling
		// Read the explanation below for how to set these rates:
		runtime.SetMutexProfileFraction(5)
		runtime.SetBlockProfileRate(5)

		_, err := pyroscope.Start(pyroscope.Config{
			ApplicationName: "saacs",

			// replace this with the address of pyroscope server
			ServerAddress: "http://pyroscope:4040",

			// you can disable logging by setting this to nil
			Logger: nil,

			// you can provide static tags via a map:
			Tags: map[string]string{
				"hostname":               os.Getenv("HOSTNAME"),
				"AuthMode":               authMode,
				"ChaincodeId":            config.CCID,
				"ChaincodeServerAddress": config.Address,
			},

			ProfileTypes: []pyroscope.ProfileType{
				// these profile types are enabled by default:
				pyroscope.ProfileCPU,
				pyroscope.ProfileAllocObjects,
				pyroscope.ProfileAllocSpace,
				pyroscope.ProfileInuseObjects,
				pyroscope.ProfileInuseSpace,

				// these profile types are optional:
				pyroscope.ProfileGoroutines,
				pyroscope.ProfileMutexCount,
				pyroscope.ProfileMutexDuration,
				pyroscope.ProfileBlockCount,
				pyroscope.ProfileBlockDuration,
			},
		})
		if err != nil {
			slog.Warn("pyroscope failed to start", "error", err)
		}

	}

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
