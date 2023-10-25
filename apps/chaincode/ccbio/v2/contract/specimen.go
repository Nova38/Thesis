package contract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

// SpecimenContract contract for handling BasicAssets
type SpecimenContract struct {
	contractapi.Contract
}

func BuildSpecimenContract() *SpecimenContract {
	return &SpecimenContract{
		Contract: contractapi.Contract{
			Name: "ccbio.Specimen",
			// BeforeTransaction: ,
			Info: metadata.InfoMetadata{
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
			},
		},
	}
}

// ────────────────────────────────────────────────────────────
