import { AnyMessage, isMessage } from '@bufbuild/protobuf'
import { ItemKeySchema } from '../gen/key_schema'
import { ItemKey } from '../gen/types_pb'

export function ToKeySchema(message: AnyMessage) {
  if (isMessage(message) === false)
    throw new Error('message is not a protobuf message')
  else if (!ItemKeySchema) throw new Error('ItemKeySchema is not defined')

  return ItemKeySchema?.[message.getType().typeName]
}

export function ItemKeyExtractor(message: AnyMessage, paths: string[]) {}

export function ToKey(message: AnyMessage) {
  const keySchema = ToKeySchema(message)
  if (!keySchema) throw new Error('key schema not found')

  const key = new ItemKey({
    collectionId: message['collectionId'],
    itemKind: keySchema.itemKind,
    itemType: keySchema.itemType,
    itemKeyParts: [],
  })

  for (const field of keySchema.properties?.paths || []) {
  }
}
