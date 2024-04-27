// import { auth, common } from 'saacs'
import { z } from 'zod'
import { saacs } from '@saacs/client'

const querySchema = z.object({
  collectionId: z.string().optional(),
})
export default defineEventHandler(async (event) => {
  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
)
if (!query.success) throw query.error.issues
  const cc = await useChaincode(event)
  const model = saacs.BiochainModel(query.data.collectionId ?? 'biochain')

  const bootstrapRequest = new pb.BootstrapRequest({
    collection: model.model.value?.collection,
  })


  const bootstrap = await cc.utilService.bootstrap(bootstrapRequest)

  switch (model.model.case) {
    case 'roles':
      const roleRequest = model.model.value.roles.map((role) => {
        return new pb.CreateRequest({
          collectionId: role.collectionId,
          policies: role.polices,
        })
      }
      break;

    default:
      break;
  }



  console.log(bootstrap)
  return bootstrap
})
