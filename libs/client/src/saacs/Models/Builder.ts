import { pb } from '@saacs/saacs-pb'
import { ActionsAlias } from '../actions'

enum UserAuth {
  Admin,
  Contributor,
  Viewer,
}

export function AuthModelBuilder(
  name: string,
  collectionId: string,
  authType: pb.AuthType,
  itemTypes: string[],
  users: Record<UserAuth, string>,
) {
  const collection = new pb.Collection({
    authType,
    collectionId,
    name,
    itemTypes,
  })

  switch (authType) {
    case pb.AuthType.IDENTITY:

      return new pb.AuthModel({
        name,
        model: {
          case: 'identity',
          value: {
            collection,
            userDirectMembership: [
              {
                collectionId,
              },
            ],
          },
        },
      })

    case pb.AuthType.GLOBAL_ROLE:
      return new pb.AuthModel({})

    case pb.AuthType.ROLE:
      return new pb.AuthModel({})
  }
}
