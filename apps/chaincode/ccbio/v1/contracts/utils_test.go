package contracts

import (
	"reflect"
	"testing"

	"github.com/nova38/thesis/apps/chaincode/ccbio/v1/context"
	schema "github.com/nova38/thesis/lib/gen/go/ccbio/schema/v1"

	"google.golang.org/genproto/protobuf/field_mask"
)

func TestUpdateMappedFiled(t *testing.T) {
	type args[T TimestampedProto] struct {
		ctx    context.LastModifiedTxContext
		oldMap map[string]T
		newMap map[string]T
	}
	type testCase[T TimestampedProto] struct {
		name           string
		args           args[T]
		wantUpdatedMap map[string]T
		wantModified   bool
		wantErr        bool
	}
	tests := []testCase[*schema.Specimen_Image]{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args[*schema.Specimen_Image]{
				ctx:    nil,
				oldMap: nil,
				newMap: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUpdatedMap, gotModified, err := UpdateMappedFiled(tt.args.ctx, tt.args.oldMap, tt.args.newMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMappedFiled() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotUpdatedMap, tt.wantUpdatedMap) {
				t.Errorf("UpdateMappedFiled() gotUpdatedMap = %v, want %v", gotUpdatedMap, tt.wantUpdatedMap)
			}
			if gotModified != tt.wantModified {
				t.Errorf("UpdateMappedFiled() gotModified = %v, want %v", gotModified, tt.wantModified)
			}
		})
	}
}

func TestGetDomainsFromMask(t *testing.T) {
	type args struct {
		mask *field_mask.FieldMask
	}
	tests := []struct {
		name        string
		args        args
		wantDomains []schema.Action_Domain
		wantErr     bool
	}{
		{
			name: "Nil mask should fail",
			args: args{
				mask: nil,
			},
			wantDomains: []schema.Action_Domain{},
			wantErr:     true,
		},
		{
			name: "Invalid mask should fail",
			args: args{
				mask: &field_mask.FieldMask{
					Paths: []string{"invalid"},
				},
			},
			wantDomains: []schema.Action_Domain{},
			wantErr:     true,
		},

		{
			name: "primary is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"primary"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_PRIMARY,
			},
			wantErr: false,
		},
		{
			name: "secondary is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"secondary"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_SECONDARY,
			},
			wantErr: false,
		},
		{
			name: "taxon is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"taxon"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_TAXON,
			},
			wantErr: false,
		},
		{
			name: "georeference is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"georeference"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_GEOREFERENCE,
			},
			wantErr: false,
		},
		{
			name: "images is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"images"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_IMAGES,
			},
			wantErr: false,
		},
		{
			name: "loans is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"loans"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_LOANS,
			},
			wantErr: false,
		},
		{
			name: "grants is a valid mask",
			args: args{
				mask: &field_mask.FieldMask{Paths: []string{"grants"}},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_GRANTS,
			},
			wantErr: false,
		},
		{
			name: "Valid mask with multiple paths",
			args: args{
				mask: &field_mask.FieldMask{
					Paths: []string{
						"primary",
						"secondary",
						"taxon",
					},
				},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_PRIMARY,
				schema.Action_DOMAIN_SECONDARY,
				schema.Action_DOMAIN_TAXON,
			},
			wantErr: false,
		},
		{
			name: "All Valid mask",
			args: args{
				mask: &field_mask.FieldMask{
					Paths: []string{
						"primary",
						"secondary",
						"taxon",
						"georeference",
						"images",
						"loans",
						"grants",
					},
				},
			},
			wantDomains: []schema.Action_Domain{
				schema.Action_DOMAIN_PRIMARY,
				schema.Action_DOMAIN_SECONDARY,
				schema.Action_DOMAIN_TAXON,
				schema.Action_DOMAIN_GEOREFERENCE,
				schema.Action_DOMAIN_IMAGES,
				schema.Action_DOMAIN_LOANS,
				schema.Action_DOMAIN_GRANTS,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDomains, err := GetDomainsFromMask(tt.args.mask)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDomainsFromMask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotDomains, tt.wantDomains) {
				t.Errorf("GetDomainsFromMask() = %v, want %v", gotDomains, tt.wantDomains)
			}
		})
	}
}
