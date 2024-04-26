// import { auth, common } from 'saacs'
import { z } from 'zod'

const querySchema = z.object({
  collectionId: z.string().optional(),
})
export default defineEventHandler(async (event) => {
  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  )
  if (!query.success) throw query.error.issues

  const bootstrapRequest = new pb.BootstrapRequest({
    collection: new pb.Collection({
      authType: pb.AuthType.ROLE,
      collectionId: 'orn',
      default: {
        defaultPolicy: new pb.PathPolicy({
          actions: [
            pb.Action.VIEW,
            pb.Action.VIEW_HISTORY,
            pb.Action.VIEW_HISTORY,
          ],
        }),
      },
      itemTypes: ['saacs.biochain.v0.Specimen'],
      name: query.data.collectionId ?? 'orn',
    }),
  })

  const cc = await useChaincode(event)
  const bootstrap = await cc.utilService.bootstrap(bootstrapRequest)
  console.log(bootstrap)
  return bootstrap
})
