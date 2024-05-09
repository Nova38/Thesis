import fs from 'node:fs'
import type { PlainMessage } from '@bufbuild/protobuf'
import { key, pb } from '@saacs/saacs-pb'
import { Any } from '@bufbuild/protobuf'
import { NormalizeSpecimen } from '../src/utils/normalize'
import orn from './f_ku_orn_cov.json' assert { type: 'json' }

export function ToItem(m: key.PrimaryItemTypeMessage) {
  return new pb.Item(
    {
      key: key.PrimaryToKeySchema(m),
      value: Any.pack(m),
    },
  )
}
const l = orn as unkown as { entries: PlainMessage<pb.Specimen>[] }

const specimens = l.entries
  .map(s => new pb.Specimen().fromJsonString(JSON.stringify(s)))
  .map(s => NormalizeSpecimen(s))

const list = new pb.SpecimenList({ specimens })

console.log('Specimens:', specimens)

fs.writeFileSync('./specimen.json', list.toJsonString())
fs.writeFileSync('./specimen.bin', list.toBinary())

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
