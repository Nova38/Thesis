import type { PlainMessage } from '@bufbuild/protobuf'
import { key, pb } from '@saacs/saacs-pb'
import { Any } from '@bufbuild/protobuf'
import orn from './f_ku_orn_cov.json' assert { type: 'json' }

export function ToItem(m: key.PrimaryItemTypeMessage) {
  return new pb.Item(
    {
      key: key.PrimaryToKeySchema(m),
      value: Any.pack(m),
    },
  )
}
const l = orn as { entries: PlainMessage<pb.Specimen>[] }

const specimens = l.entries
  .map(s => new pb.Specimen({ ...s }))
  .map(s => ToItem(s))
  .map(i => new pb.CreateRequest({ item: i }))

console.log('Specimens:', specimens)

// const l = orn as { entries: PlainMessage<pb.Specimen>[] }

// console.log('Specimens:', specimens)
// // const batch = specimens.splice(0, 100)

// const { client, gateway } = await BuildFromBaseDir('../../../infra/network/')

// const utils = createUtilGateway(contract)
// const BiochainClient = createBiochainGateway(contract)
// try {
//   const u = await utils.getCurrentUser({})
//   console.log(u)

//   // const r = await utils.bootstrap({ collection: { collectionId: 'Testing', name: 'Testing', authType: pb.AuthType.ROLE, itemTypes: [
//   //   'saacs.biochain.v0.Specimen',
//   // ] } })

//   // console.log(r)

//   const r2 = await utils.getCollectionsList({})
//   console.log(r2)
// }
// catch (error) {
//   console.error(error)
// }
