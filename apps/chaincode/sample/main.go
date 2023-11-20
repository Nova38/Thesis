package main

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials"
)

const (
	mspID        = "Org1MSP"
	basePath     = "C:/Users/devcs/.vscode/extensions/spydra.hyperledger-fabric-debugger-0.1.3/"
	cryptoPath   = basePath + "fabric/local/organizations/peerOrganizations/org1.debugger.com"
	certPath     = cryptoPath + "/users/User1/msp/signcerts/cert.pem"
	keyPath      = cryptoPath + "/users/User1/msp/keystore/"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.debugger.com/tls/ca.crt"
	peerEndpoint = "localhost:5051"
	gatewayPeer  = "peer0.org1.debugger.com"
)

func main() {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	fmt.Printf("gw: %+v\n", gw)

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "Thesis-caas"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "default"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	contract.SubmitTransaction("auth.users:UserGetCurrentId")
	UserGetCurrentId(contract)
	// initLedger(contract)
	// getAllAssets(contract)
	// createAsset(contract)
	// readAssetByID(contract)
	// transferAssetAsync(contract)
	// exampleErrorHandling(contract)
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	// transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := os.ReadDir(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func UserGetCurrentId(contract *client.Contract) (string, error) {
	fmt.Println(
		"\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger",
	)

	bytes := []byte(`{"msp_id":"Org1MSP","user_id":"User1"}`)

	evaluateResult, err := contract.EvaluateTransaction("auth.generic:GetAllTypes", string(bytes))
	if err != nil {
		fmt.Printf("Error, failed to evaluate transaction: %v", err)
		return "", fmt.Errorf("failed to evaluate transaction: %w", err)
	}
	x := map[string]string{}
	// v := &struct {
	// 	msp_id  string
	// 	user_id string
	// }{}
	json.Unmarshal(evaluateResult, &x)

	fmt.Println(x)

	s, err := json.MarshalIndent(string(evaluateResult), "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal result: %w", err)
	}

	fmt.Printf("Result: %s\n", s)
	// fmt.Println(s)
	// fmt.Println(evaluateResult)

	fmt.Println("End of UserGetCurrentId")
	return string(evaluateResult), nil
}
