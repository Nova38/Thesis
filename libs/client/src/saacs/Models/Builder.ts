import { pb } from '@saacs/saacs-pb'
import { ActionsAlias } from '../actions'

enum UserAuth {
  Admin,
  Contributor,
  Viewer,
}

export function AuthModelBuilder(
  collection: string,
  type: pb.AuthType,
  itemTypes: string[],
  users: map<UserAuth, string>,
) {
  switch (type) {
    case pb.AuthType.NONE :
      return new pb.AuthModel({})

    case pb.AuthType.IDENTITY:
      return new pb.AuthModel({})

    case pb.AuthType.GLOBAL_ROLE:
      return new pb.AuthModel({})

    case pb.AuthType.ROLE:
      return new pb.AuthModel({})
  }
}
