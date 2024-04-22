import Ajv2020 from 'ajv/dist/2020'

const ajv = new Ajv2020()

import * as itemkey from '../gen/auth.ItemKey.schema.json'
import { ItemKey } from '../../gen'

ajv.addSchema(itemkey)
const i = new ItemKey({
  collectionId: '123',
  itemKeyParts: [],
})

console.log(i.toJson({ useProtoFieldName: true }))

const valid = ajv.validate(
  'auth.ItemKey.schema.json',
  i.toJson({ useProtoFieldName: true }),
)
if (!valid) console.log(ajv.errors)

console.log(ajv.errors)
