package context

import (
	"reflect"
	"testing"

	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v2"
)

func TestExtractSpecimenId(t *testing.T) {
	type args struct {
		req interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantId  *pb.Specimen_Id
		wantErr bool
	}{
		{
			name: "Specimen Create Request Extraction",
			args: args{
				req: &pb.SpecimenCreateRequest{
					Id: &pb.Specimen_Id{
						CollectionId: "CollectionIDValue",
						Id:           "IDValue",
					},
					Taxon:        &pb.Specimen_Taxon{},
					Primary:      &pb.Specimen_Primary{},
					Secondary:    &pb.Specimen_Secondary{},
					Georeference: &pb.Specimen_Georeference{},
					Images:       map[string]*pb.Specimen_Image{},
					Loans:        "",
					Grants:       "",
				},
			},
			wantId: &pb.Specimen_Id{
				CollectionId: "CollectionIDValue",
				Id:           "IDValue",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := ExtractSpecimenId(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractSpecimenId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotId, tt.wantId) {
				t.Errorf("ExtractSpecimenId() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}
