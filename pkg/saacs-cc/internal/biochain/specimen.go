// biochain.v1
package biochain

import (
	"strings"

	"github.com/nova38/saacs/pkg/saacs-cc/common"
	pb "github.com/nova38/saacs/pkg/saacs-protos/saacs/common/v0"

	biochainpb "github.com/nova38/saacs/pkg/saacs-protos/saacs/biochain/v0"
	"github.com/samber/lo"
	"github.com/samber/oops"
)

func SpecimenPostActionProcessing(
	ctx common.TxCtxInterface,
	item common.ItemInterface,
	ops []*pb.Operation,
) error {

	specimen, ok := item.(*biochainpb.Specimen)
	if ok {
		activity, err := ctx.MakeLastModified()
		if err != nil {
			return oops.Wrap(err)
		}

		op := ops[0]
		switch op.GetAction() {

		case pb.Action_ACTION_CREATE:
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

		case pb.Action_ACTION_UPDATE:
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
