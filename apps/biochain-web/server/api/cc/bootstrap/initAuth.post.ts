// import { auth, common } from 'saacs'
import { z } from 'zod'
import { saacs } from '@saacs/client'

// const querySchema =
export default defineEventHandler(async (event) => {
  const query = await getValidatedQuery(
    event,
    z.object({
      collectionId: z.string(),
    }).parse,
  )

  await requireAuthSession(event)

  const cc = await useChaincode(event)

  const requests = saacs.BootstrapBiochainRequests(query.collectionId)
  console.log(
    'Requests',
    requests.bootstrap.toJsonString({ typeRegistry: cc.service.registry }),
  )
  // // Make the collection request

  const replies = []

  try {
    console.log('Bootstrap Biochain')

    const bootstrap = await cc.utilService.bootstrap(requests.bootstrap)
    console.log('Bootstrap Biochain successful')
    replies.push({
      key: 'bootstrap',
      reply: bootstrap.toJson({ typeRegistry: cc.service.registry }),
    })
  } catch (error) {
    console.error('Failed to create', JSON.stringify(error))
    replies.push({ key: 'bootstrap', error })
  }

  for (const r of requests.create) {
    const k = r?.item?.key?.itemKeyParts?.join(':')
    console.log('Creating', k)
    try {
      const reply = await cc.service.create(r)
      replies.push({
        key: k,
        reply: reply.toJson({ typeRegistry: cc.service.registry }),
      })
    } catch (error) {
      console.error('Failed to create', k, JSON.stringify(error))
      replies.push({ key: k, error })
    }
  }

  return { replies }
})
