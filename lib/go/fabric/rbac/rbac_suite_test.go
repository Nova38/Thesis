package rbac_test

import (
	"fmt"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	pb "github.com/nova38/thesis/lib/go/gen/chaincode/ccbio/schema/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRbac(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rbac Suite")
}

var _ = Describe("RBAC", func() {
	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	Describe("Collection CEL validation", func() {
		Context("Missing Parts", func() {
			It("Empty Collection Should Fail", func() {
				msg := &pb.Collection{
					Roles: map[int32]string{},
				}
				Expect(v.Validate(msg)).To(BeFalse())
			})
		})

		// Context("Valid Collection", func() {
		// 	It("Should Pass", func() {
		// 		msg := &pb.Collection{
		// 			Id:               &pb.Collection_Id{},
		// 			ObjectNamespace:  "",
		// 			ObjectDescriptor: &descriptorpb.DescriptorProto{},
		// 			Roles:            map[int32]string{},
		// 			Acl:              map[int32]*pb.Collection_ACL{},
		// 		}
		// 		Expect(v.Validate(msg)).To(Succeed())
		// 	})
		// })
	})
})
