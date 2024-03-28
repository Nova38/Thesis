// biochain.v1
package biochain

import (
	"strings"

	authpb "github.com/nova38/saacs/libs/saacs-protos-go/auth/v1"
	biochainpb "github.com/nova38/saacs/libs/saacs-protos-go/biochain/v1"
	"github.com/nova38/saacs/pkg/chaincode/common"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func SpecimenPostActionProcessing(
	ctx common.TxCtxInterface,
	item common.ItemInterface,
	ops []*authpb.Operation,
) error {

	specimen, ok := item.(*biochainpb.Specimen)
	if ok {
		activity, err := ctx.MakeLastModified()
		if err != nil {
			return oops.Wrap(err)
		}

		op := ops[0]
		switch op.GetAction() {

		case authpb.Action_ACTION_CREATE:
			// Make sure all of the fields are there
			if specimen.GetPrimary() == nil {
				specimen.Primary = &biochainpb.Specimen_Primary{}
			}
			if specimen.GetSecondary() == nil {
				specimen.Secondary = &biochainpb.Specimen_Secondary{}
			}
			if specimen.GetTaxon() == nil {
				specimen.Taxon = &biochainpb.Specimen_Taxon{}
			}
			if specimen.GetGeoreference() == nil {
				specimen.Georeference = &biochainpb.Specimen_Georeference{}
			}

			specimen.Primary.LastModified = activity
			specimen.Secondary.LastModified = activity
			specimen.Taxon.LastModified = activity
			specimen.Georeference.LastModified = activity

		case authpb.Action_ACTION_UPDATE:
			paths := op.GetPaths().GetPaths()

			rootPaths := lo.UniqBy(paths, func(p string) string {
				return strings.Split(p, ".")[0]
			})

			for _, rootPath := range rootPaths {
				switch rootPath {
				case "primary":
					if specimen.GetPrimary() == nil {
						specimen.Primary = &biochainpb.Specimen_Primary{}
					}
					specimen.Primary.LastModified = activity
				case "secondary":
					if specimen.GetSecondary() == nil {
						specimen.Secondary = &biochainpb.Specimen_Secondary{}
					}
					specimen.Secondary.LastModified = activity
				case "taxon":
					if specimen.GetTaxon() == nil {
						specimen.Taxon = &biochainpb.Specimen_Taxon{}
					}
					specimen.Taxon.LastModified = activity
				case "georeference":
					if specimen.GetGeoreference() == nil {
						specimen.Georeference = &biochainpb.Specimen_Georeference{}
					}
					specimen.Georeference.LastModified = activity
				}
			}
			specimen.LastModified = activity

		}
	}
	return nil
}
