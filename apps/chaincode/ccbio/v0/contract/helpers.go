package contract

import (
	"strings"

	authpb "github.com/nova38/thesis/lib/go/gen/auth/v1"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v0"
	lop "github.com/samber/lo/parallel"
	"github.com/samber/oops"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Update Specimen with mask

func SetLastModByMask(
	specimen *pb.Specimen,
	mask *fieldmaskpb.FieldMask,
	mod *authpb.StateActivity,
) error {
	if mask == nil {
		return nil
	}

	grouped := lop.GroupBy(mask.Paths, func(i string) string {
		return strings.Split(i, ".")[0]
	})

	for _, i := range grouped {
		f := strings.Split(i[0], ".")[0]
		switch f {
		case "Primary", "primary":
			if specimen.Primary == nil {
				specimen.Primary = &pb.Specimen_Primary{}
			}
			specimen.Primary.LastModified = mod
		case "Secondary":
			if specimen.Secondary == nil {
				specimen.Secondary = &pb.Specimen_Secondary{}
			}
			specimen.Secondary.LastModified = mod
		case "Taxon":
			if specimen.Taxon == nil {
				specimen.Taxon = &pb.Specimen_Taxon{}
			}
			specimen.Taxon.LastModified = mod
		case "Georeference":
			if specimen.Georeference == nil {
				specimen.Georeference = &pb.Specimen_Georeference{}
			}
			specimen.Georeference.LastModified = mod

		case "Image":
			for _, j := range i {
				// check to see if the image is in the array
				parts := strings.Split(j, ".")
				if len(parts) < 2 {
					continue
				}
				fm := parts[1]
				_, ok := specimen.Images[fm]
				if !ok {
					return oops.
						With("image", fm).
						Errorf("Image not found")
				}
				if specimen.Images[fm] == nil {
					specimen.Images[fm] = &pb.Specimen_Image{}
				}
				specimen.Images[fm].LastModified = mod
			}
		case "Loans":
			for _, j := range i {
				// check to see if the image is in the array
				parts := strings.Split(j, ".")
				if len(parts) < 2 {
					continue
				}
				fm := parts[1]
				_, ok := specimen.Loans[fm]
				if !ok {
					return oops.
						With("Loans", fm).
						Errorf("Image not found")
				}
				if specimen.Loans[fm] == nil {
					specimen.Loans[fm] = &pb.Specimen_Loan{}
				}
				specimen.Loans[fm].LastModified = mod
			}
		case "Grants":
			for _, j := range i {
				// check to see if the image is in the array
				parts := strings.Split(j, ".")
				if len(parts) < 2 {
					continue
				}
				fm := parts[1]
				_, ok := specimen.Grants[fm]
				if !ok {
					return oops.
						With("Loans", fm).
						Errorf("Image not found")
				}
				if specimen.Grants[fm] == nil {
					specimen.Grants[fm] = &pb.Specimen_Grant{}
				}
				specimen.Grants[fm].LastModified = mod
			}
		case "HiddenTx", "Id":
			return oops.Errorf("Cannot set LastModified on %s", f)
		default:
			return oops.
				With("field", f).
				Errorf("Unknown field")

		}
	}

	return nil
}
