package contracts

import (
	"log/slog"

	"github.com/nova38/biochain/chaincode/ccbio/context"
	"github.com/nova38/biochain/chaincode/ccbio/schema"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/samber/oops"
	"google.golang.org/protobuf/proto"

	"google.golang.org/genproto/protobuf/field_mask"
)

type TimestampedProto interface {
	proto.Message
	GetLastModified() *schema.LastModified
}

type MappedProto interface {
	schema.Specimen_Image | schema.Specimen_HiddenTx
}

func UpdateMappedFiled[T TimestampedProto](ctx context.LastModifiedTxContext, oldMap map[string]T, newMap map[string]T) (updatedMap map[string]T, modified bool, err error) {
	// Make a cleared last modified
	clearedLastModified := ctx.MakeClearedLastModified()
	// Make last modified for this transaction
	updatedAt, err := ctx.MakeLastModified()
	if err != nil {
		return nil, false, errors.Wrap(err, "Failed to make last modified")
	}
	// Get the keys for the old and new maps
	oKeys := lo.Keys(oldMap)
	nKeys := lo.Keys(newMap)
	// If both are empty, then nothing to do
	if len(oKeys) == 0 && len(nKeys) == 0 {
		removed, added := lo.Difference(oKeys, nKeys)
		shared := lo.Union(oKeys, nKeys)
		updatedImages := lo.PickByKeys(newMap, added)
		// any updated?
		modified := (len(removed) > 0) || (len(added) > 0)
		for _, key := range shared {
			newItem := newMap[key]
			oldItem := oldMap[key]
			if !proto.Equal(clearedLastModified, newItem.GetLastModified()) {
				if !proto.Equal(oldItem, newItem) {
					updatedImages[key] = newItem
					modified = true
					// Update the last modified
					modifiedAt, ok := proto.Clone(updatedAt).(*schema.LastModified)
					if !ok {
						return nil, true, errors.New("Failed to clone updatedAt")
					}
					// Update the last modified
					ts := updatedImages[key].GetLastModified()
					if ts == nil {
						return nil, true, errors.New("Failed to get last modified")
					}
					ts.TxId = modifiedAt.TxId
					ts.UpdatedAt = modifiedAt.UpdatedAt
					ts.UserId = modifiedAt.UserId
					ts.UserName = modifiedAt.UserName
				}
				updatedImages[key] = oldItem
			} else {
				updatedImages[key] = oldItem
			}
		}
		return updatedMap, modified, nil
	}
	return oldMap, false, nil
}

var SpecimenFeildDomainMap map[string]schema.Action_Domain = map[string]schema.Action_Domain{
	"primary":      schema.Action_DOMAIN_PRIMARY,
	"secondary":    schema.Action_DOMAIN_SECONDARY,
	"taxon":        schema.Action_DOMAIN_TAXON,
	"georeference": schema.Action_DOMAIN_GEOREFERENCE,
	"images":       schema.Action_DOMAIN_IMAGES,
	"loans":        schema.Action_DOMAIN_LOANS,
	"grants":       schema.Action_DOMAIN_GRANTS,
}

func GetDomainsFromMask(mask *field_mask.FieldMask) (domains []schema.Action_Domain, err error) {
	// Get the domains from the mask

	domains = []schema.Action_Domain{}
	if mask == nil {
		return domains, oops.Errorf("Mask is nil")
	}
	if !mask.IsValid(&schema.Specimen{}) {
		return domains, oops.Errorf("Invalid mask: %v", mask)
	}

	for _, path := range mask.Paths {

		domain, ok := SpecimenFeildDomainMap[path]
		if !ok {
			return nil, errors.Wrap(err, "Failed to get domain from mask")
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

// /**
//  * UpdateViaFeildMask updates the current object with the updated object
//  **/
// func UpdateViaFeildMask(
// 	current *schema.Specimen,
// 	updated *schema.Specimen,
// 	mask *field_mask.FieldMask,

// ) (domains []schema.Action_Domain, err error) {

// 	if !mask.IsValid(&schema.Specimen{}) {
// 		return domains, oops.Errorf("Invalid mask: %v", mask)
// 	}

// 	ref_current := current.ProtoReflect()
// 	ref_updated := updated.ProtoReflect()

// 	// Filter the updated object to only the fields in the mask
// 	masked, err := fmutils.Filter(ref_updated, mask.GetPaths())
// 	if err != nil {
// 		return nil, oops.Wrap(err)
// 	}

// 	return domains, nil
// }

type ClearableFeild interface {
	GetLastModified() *schema.LastModified
}

func IsCleared[T ClearableFeild](obj T) bool {
	last_modified := obj.GetLastModified()

	if last_modified == nil {
		return true
	}
	return proto.Equal(obj.GetLastModified(), &schema.LastModified{})
}

type MessageFeild interface {
	proto.Message
	GetLastModified() *schema.LastModified
}

func UpdateMessageField(field string, cur, req MessageFeild) (bool, error) {
	fn := "UpdateMessageField"

	if req == nil {
		slog.Warn(fn, field, " on req is nil")
	} else if IsCleared(req) {
		slog.Info(fn, field, "skipping => last modified cleared")
	} else if proto.Equal(cur, req) {
		slog.Info(fn, field, "skipping => no change")
	} else {
		slog.Info(fn, field, "updating =>", slog.Any("primary", req))

		// if err := ctx.AddActionDomain(schema.Action_DOMAIN_PRIMARY); err != nil {
		// 	return nil, errors.Wrap(err, "SpecimenUpdate failed to add action domain")
		// }

		proto.Merge(cur, req)

		return true, nil
		// // Update the last modified
		// modified, ok := proto.Clone(updatedAt).(*schema.LastModified)
		// if !ok {
		// 	return false, errors.New("Failed to clone updatedAt")
		// }
		// cur.LastModified = modified
	}

	return false, nil
}

// // Separate function
// func updateImages(ctx context.TxContext, currentImages, newImages map[string]*schema.Specimen_Image) (map[string]*schema.Specimen_Image, bool) {

// 	updatedAt :=

// 	cKeys := lo.Keys(currentImages)
// 	nKeys := lo.Keys(newImages)

// 	removed, added := lo.Difference(cKeys, nKeys)
// 	shared := lo.Union(cKeys, nKeys)

// 	updatedImages := lo.PickByKeys(newImages, added)
// 	// any updated?
// 	modified := (len(removed) > 0) || (len(added) > 0)

// 	for _, key := range shared {
// 		newItem := newImages[key]
// 		oldItem := currentImages[key]
// 		if !proto.Equal(ctx.MakeClearedLastModified(), newItem.GetLastModified()) {
// 			if !proto.Equal(oldItem, newItem) {
// 				updatedImages[key] = newItem
// 				modified = true
// 				// Update the last modified
// 				modified, ok := proto.Clone(updatedAt).(*schema.LastModified)
// 				if !ok {
// 					return nil, errors.New("Failed to clone updatedAt")
// 				}
// 				updatedImages[key].LastModified = modified
// 			}
// 			updatedImages[key] = oldItem
// 		} else {
// 			updatedImages[key] = oldItem
// 		}
// 	}

// 	return updatedImages, modified
// }
