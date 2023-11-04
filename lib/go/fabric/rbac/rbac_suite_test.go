package rbac_test

import (
	"testing"

	"github.com/nova38/thesis/lib/go/fabric/rbac"
	rbac_pb "github.com/nova38/thesis/lib/go/gen/rbac"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRbac(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rbac Suite")
}

var _ = Describe("rbac", func() {
	var collection rbac_pb.Collection

	Describe("Object Paths", func() {
		BeforeEach(func() {
			collection = rbac_pb.Collection{
				Id:              &rbac_pb.Collection_Id{},
				ObjectNamespace: "",
				ObjectType:      "",
				Roles:           map[string]string{},
				Acl: map[string]*rbac_pb.ACL{
					"0": {
						RoleDefs: &rbac_pb.ACL_Policy_Roles{
							View:   false,
							Create: false,
							Edit:   false,
							Delete: false,
						},
						RolePermissions: &rbac_pb.ACL_Policy_Roles{
							View:   false,
							Create: false,
							Edit:   false,
							Delete: false,
						},
						Memberships: &rbac_pb.ACL_Policy_Roles{
							View:   false,
							Create: false,
							Edit:   false,
							Delete: false,
						},
						Object: &rbac_pb.ACL_Policy_Object{
							View:        false,
							Create:      false,
							Delete:      false,
							ViewHistory: false,
							HiddenTx:    false,
						},
						ObjectPaths: &rbac_pb.ACL_PathRolePermission{},
					},
					"1": {
						RoleDefs: &rbac_pb.ACL_Policy_Roles{
							View:   false,
							Create: false,
							Edit:   false,
							Delete: false,
						},
						RolePermissions: &rbac_pb.ACL_Policy_Roles{
							View:   false,
							Create: false,
							Edit:   false,
							Delete: false,
						},
						Memberships: &rbac_pb.ACL_Policy_Roles{
							View:   false,
							Create: false,
							Edit:   false,
							Delete: false,
						},
						Object: &rbac_pb.ACL_Policy_Object{
							View:        false,
							Create:      false,
							Delete:      false,
							ViewHistory: false,
							HiddenTx:    false,
						},
						ObjectPaths: &rbac_pb.ACL_PathRolePermission{
							Path:          "",
							AllowSubPaths: true,
							SubPaths:      map[string]*rbac_pb.ACL_PathRolePermission{},
							Policy:        &rbac_pb.ACL_Policy_ObjectField{},
						},
					},
				},
			}
		})

		It("List all paths", func() {
			paths := rbac.GetAllPaths(&collection)

			Expect(paths).NotTo(BeEmpty())
		})
	})
})
