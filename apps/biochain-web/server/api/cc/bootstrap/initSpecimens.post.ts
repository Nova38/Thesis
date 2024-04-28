// import { auth, common } from 'saacs'
import { z } from 'zod'
import { saacs } from '@saacs/client'
import { ZSpecimen } from '~/utils/cc/proto/Specimen'
import { destr } from 'destr'

// const querySchema =
export default defineEventHandler(async (event) => {
  const query = await getValidatedQuery(
    event,
    z.object({
      collectionId: z.string(),
    }).parse,
  )
  //await requireAuthSession(event)

  const body = await readValidatedBody(
    event,
    z.object({ entries: z.array(ZSpecimen) }).parse,
  )

  const cc = await useChaincode(event)

  const collectionId = query.collectionId ?? 'default'

  const txs = body.entries
    .map((s) => new pb.Specimen({ ...s, collectionId }))
    .map((s) => saacs.PrimaryToItem(s))
    .map((i) => new pb.CreateRequest({ item: i }))

  // // Make the collection request
  console.log('Bootstrap Biochain')

  const replies = await Promise.allSettled(
    txs.map(async (tx) => {
      const key = tx?.item?.key?.itemKeyParts?.join(':')
      console.log('Attempting To Make', key)

      try {
        const r = await cc.service.create(tx)
        console.log(`Creation of  ${key} successful`)
        return { key, failed: false }
      } catch (e) {
        console.error('Failed to create', key, destr(e))
        console.warn(`Creation of  ${key} Failed`)
        return { key, failed: true }
      }
    }),
  )

  const res = useStorage(`import:log`).setItem(`${Date.now()}`, replies)
  return replies

  // for (const tx of txs) {
  //   const key = tx?.item?.key?.itemKeyParts?.join(':')
  //   console.log('Creating', key)
  //   try {
  //     const reply = await cc.service.create(tx)
  //     replies.push({
  //       key: key,
  //       reply: reply.toJson({ typeRegistry: cc.service.registry }),
  //     })
  //     console.log(`Creation of  ${key} successful`)
  //     replies.push({
  //       key,
  //       success: true,
  //       reply: reply.toJson({ typeRegistry: cc.service.registry }),
  //     })
  //   } catch (error) {
  //     console.error('Failed to create', key, JSON.stringify(error))
  //     console.warn(`Creation of  ${key} Failed`)
  //     replies.push({ key, error })
  //   }
  // }

  return { replies }
})
