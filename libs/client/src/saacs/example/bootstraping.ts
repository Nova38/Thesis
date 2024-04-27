import { pb } from '@saacs/saacs-pb'

export function BiochainModel() {
  const collectionId = 'biochain'
  return new pb.AuthModel({
    name: 'Biochain Role Based Auth',
    model: {
      case: 'roles',
      value: {
        collection: {
          collectionId,
          // note: "This is a test collection",
          authType: pb.AuthType.ROLE,
          itemTypes: [
            pb.Specimen.typeName,
            pb.Collection.typeName,
            pb.UserCollectionRoles.typeName,
            pb.Role.typeName,
          ],
          default: {
            defaultPolicy: {
              actions: [
                pb.Action.VIEW,
                pb.Action.SUGGEST_VIEW,
                pb.Action.VIEW_HISTORY,
              ],
              allowSubPaths: false,
              path: '',
            },
            itemPolicies: {
              [pb.Specimen.typeName]: {
                actions: [
                  pb.Action.VIEW,
                  pb.Action.SUGGEST_VIEW,
                  pb.Action.VIEW_HISTORY,
                ],
                allowSubPaths: false,
                fullPath: '',
                path: '',
              },
              [pb.UserCollectionRoles.typeName]: {
                actions: [pb.Action.VIEW, pb.Action.SUGGEST_VIEW],
                allowSubPaths: false,
                fullPath: '',
                path: '',
              },
              [pb.Role.typeName]: {
                actions: [pb.Action.VIEW, pb.Action.SUGGEST_VIEW],
                allowSubPaths: false,
                fullPath: '',
                path: '',
              },
            },
          },
        },
        roles: [
          {
            roleId: 'admin',
            note: 'Admin Role',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {},
            },
          },
          {
            roleId: 'suggester',
            note: '',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {},
            },
          },
          {
            roleId: 'viewer',
            note: '',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {},
            },
          },
          {
            roleId: 'approver',
            note: '',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {},
            },
          },
        ],
        userCollectionRoles: [],
      },
    },
  })
}
