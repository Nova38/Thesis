package rbac

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/nova38/thesis/lib/go/gen/rbac"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestValidations(t *testing.T) {
	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	testCases := []struct {
		msg         *rbac.ACL_Operation
		valid       bool
		wantMessage string
		desc        string
	}{
		{
			msg: &rbac.ACL_Operation{
				Domain: 0,
				Action: 0,
				Paths:  &fieldmaskpb.FieldMask{},
			},
			valid:       false,
			wantMessage: "Domain and action must not be set to unspecified",
			desc:        "Validate ACL_Operation with unspecified domain and action",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err = v.Validate(tC.msg)
			t.Log("err", err, strings.Contains(err.Error(), tC.wantMessage))
			switch {
			case tC.valid && err != nil:
				t.Errorf("Validate() got unexpected error: %v", err)
			case !tC.valid && err == nil:
				t.Errorf("Validate() got unexpected success")
			case !tC.valid && err != nil && !strings.Contains(err.Error(), tC.wantMessage):
				t.Errorf("Validate() got unexpected error: %v, want: %v", err, tC.wantMessage)
			case tC.valid && err == nil:
				return
			}
		})
	}
}
