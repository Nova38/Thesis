import { pb } from '@saacs/saacs-pb'
import { ActionsAlias } from '../actions'

export function BiochainModel(collectionId?: string): pb.AuthModel {
  collectionId ??= 'biochain'
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
              actions: [...ActionsAlias.all.level.view],
              allowSubPaths: false,
              path: '',
            },
            itemPolicies: {},
          },
        },
        roles: [
          {
            roleId: 'PermissionManager',
            note: 'Role to manage permissions',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {
                [pb.Role.typeName]: {
                  actions: [...ActionsAlias.all.full],
                  allowSubPaths: false,
                },
                [pb.UserCollectionRoles.typeName]: {
                  actions: [...ActionsAlias.all.full],
                  allowSubPaths: false,
                },
              },
            },
          },
          {
            roleId: 'SpecimenManager',
            note: '',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {
                [pb.Specimen.typeName]: {
                  actions: [...ActionsAlias.all.full],
                  allowSubPaths: false,
                },
              },
            },
          },
          {
            roleId: 'Contributor',
            note: 'I am a note',
            parentRoleIds: [],
            collectionId,
            polices: {
              defaultPolicy: {},
              defaultExcludedTypes: [],
              itemPolicies: {
                [pb.Specimen.typeName]: {
                  actions: [
                    ...ActionsAlias.all.level.view,
                    ...ActionsAlias.all.level.suggest,
                    ...ActionsAlias.all.level.create,
                  ],
                  allowSubPaths: true,
                  subPaths: {
                    primary: {
                      actions: [...ActionsAlias.all.full],
                      allowSubPaths: false,
                      fullPath: 'primary',
                      path: 'primary',
                    },
                    secondary: {
                      actions: [...ActionsAlias.all.full],
                      allowSubPaths: false,
                      fullPath: 'secondary',
                      path: 'secondary',
                    },
                    georeference: {
                      actions: [...ActionsAlias.all.level.view],
                      allowSubPaths: true,
                      fullPath: 'georeference',
                      path: 'georeference',
                      subPaths: {
                        country: {
                          actions: [...ActionsAlias.all.level.view],
                          allowSubPaths: false,
                          fullPath: 'georeference.country',
                          path: 'country',
                        },
                      },
                    },

                    taxon: {
                      actions: [...ActionsAlias.all.full],
                      allowSubPaths: false,
                      fullPath: 'taxon',
                      path: 'taxon',
                    },
                    loans: {
                      actions: [...ActionsAlias.all.full],
                      allowSubPaths: false,
                      fullPath: 'loans',
                      path: 'loans',
                    },
                    grants: {
                      actions: [...ActionsAlias.all.full],
                      allowSubPaths: false,
                      fullPath: 'grants',
                      path: 'grants',
                    },
                    images: {
                      actions: [...ActionsAlias.all.full],
                      allowSubPaths: false,
                      fullPath: 'images',
                      path: 'images',
                    },

                  },
                },
              },
            },
          },

          // {
          //   roleId: 'Suggester',
          //   note: 'Suggest Updates to Specimens',
          //   parentRoleIds: [],
          //   collectionId,
          //   polices: {
          //     defaultPolicy: {
          //       actions: [...ActionsAlias.all.level.suggest],
          //     },
          //     defaultExcludedTypes: [
          //       pb.Collection.typeName,
          //       pb.Role.typeName,
          //       pb.UserCollectionRoles.typeName,
          //     ],
          //     itemPolicies: {},
          //   },
          // },
        ],
      },
    },
  })
}
