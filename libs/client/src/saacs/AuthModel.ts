import { pb } from '@saacs/saacs-pb'
import type { PartialMessage } from '@bufbuild/protobuf'
import { PrimaryToItem } from './Item'

// export interface AuthModel =

export function AuthModelToRequests(authModel: PartialMessage<pb.AuthModel>): {
  bootstrap: pb.BootstrapRequest
  create: pb.CreateRequest[]
} {
  const model = new pb.AuthModel(authModel).model
  if (!model)
    throw new Error('AuthModel model is not defined')

  const bootstrapRequest = new pb.BootstrapRequest({
    collection: model.value?.collection,
  })

  switch (model.case) {
    case 'roles':
      return {
        bootstrap: bootstrapRequest,
        create: model.value.roles.map((role) => {
          return new pb.CreateRequest({
            item: PrimaryToItem(role),
          })
        }),
      }

    case 'identity':
      return {
        bootstrap: bootstrapRequest,
        create: model.value.userDirectMembership.map((identity) => {
          return new pb.CreateRequest({
            item: PrimaryToItem(identity),
          })
        }),
      }

    case 'globalRoles':
      return {
        bootstrap: bootstrapRequest,
        create: model.value.roles.map((globalRole) => {
          return new pb.CreateRequest({
            item: PrimaryToItem(globalRole),
          })
        }),
      }

    case 'attribute':
      return {
        bootstrap: bootstrapRequest,
        create: model.value.attribute.map((attribute) => {
          return new pb.CreateRequest({
            item: PrimaryToItem(attribute),
          })
        }),
      }

    default:
      throw new Error('AuthModel case not recognized')
      break
  }
}
