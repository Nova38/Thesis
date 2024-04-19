import { GeneratedFile, Schema } from '@bufbuild/protoplugin/ecmascript'
import {
  AllMessagesWithExtension,
  GetAllTypesFromSchema,
  allMessagesWithExtension,
  getAllMessages,
  getAllTypes,
} from '../utils'
import { camelCase } from 'scule'

import { KeySchema, key_schema } from '../gen/auth/v1/auth_pb'
import {
  DescExtension,
  DescMessage,
  getExtension,
  hasExtension,
} from '@bufbuild/protobuf'

function getDesc(schema: Schema, name: string) {
  for (const t of GetAllTypesFromSchema(schema)) {
    if (t.typeName === name) {
      return t as DescMessage
    }
  }
  return undefined
}
function getMessageImport(f: GeneratedFile, schema: Schema, name: string) {
  const desc = getDesc(schema, name)
  return !desc ? f.import(name, '@saacs/protos-es') : f.import(desc)
}

function getKeySchemaDesc(schema: Schema) {
  return getDesc(schema, 'auth.KeySchema')
}

// prettier-ignore
export function generateKeySchema(schema: Schema) {
  const f = schema.generateFile('key_schema.ts')

  const { Message, MessageType, proto3 } = schema.runtime
  const isMessageImport = f.import('isMessage ', '@bufbuild/protobuf')



  const KeySchemaImport =  getMessageImport(f, schema, 'auth.KeySchema')
  const ItemKeyImport = getMessageImport(f, schema, 'auth.ItemKey')
  const ItemImport = getMessageImport(f, schema, 'auth.Item')

    const Items = [...AllMessagesWithExtension(schema, key_schema)].map(
      (item) => {
        item.extension.itemType = item.message.typeName
        return item
      },
    )
    const PrimaryItems = Items.filter((item) => item.extension.itemKind === 2)
    const SecondaryItems = Items.filter((item) => item.extension.itemKind === 3)



  const exportFunction = (prefix: string,items: { message: DescMessage, extension: KeySchema }[]) => {
    f.print`${f.exportDecl('type', prefix + 'ItemType')} = `

    items.forEach(({ message }) => {f.print`  | ${message}`    })
    f.print('')

    // prettier-ignore
    f.print`${f.exportDecl('const', prefix + 'ItemKeySchema')} : Record<string, ${KeySchemaImport}> = {`
    items.forEach(({ message, extension }) => {
      f.print`  '${message.typeName}': ${KeySchemaImport}.fromJson(${extension
        .toJsonString({emitDefaultValues: true, enumAsInteger: true, prettySpaces: 4})}),`
    })
    f.print`}`

  }

  exportFunction('', Items)
  exportFunction('Primary', PrimaryItems)
  exportFunction('Secondary', SecondaryItems)


  f.print`${f.exportDecl('function', 'PrimaryToKeySchema')} (item: PrimaryItemType){`
  f.print`  switch (true) {`
  PrimaryItems.map(({ message, extension }) => {
    f.print`// ${extension.properties?.paths.join(', ')?? ''}`
    const paths = extension.properties?.paths?.map((path) =>
      path
        .split('.')
        .map((s) => camelCase(s))
        .join('?.'),
    )?.map((path) => `item.${path}`)

    f.print`// ${paths?.join(', ') ?? ''}`

    f.print`    case ${isMessageImport}(item, ${message}):`
    f.print`      return new ${ItemKeyImport}({`
    f.print`        itemType: '${message.typeName}',`
    f.print`        itemKind: ${extension.itemKind},`
    f.print`        collectionId: item?.collectionId,`
    f.print`        itemKeyParts: [ ${paths?.join(', ') ?? ''} ]`
    f.print`      })`
  })
  f.print`    default:`
  f.print`      throw new Error('Unknown item type')`
  f.print`  }`
  f.print`}`


  //   const ToKeySchemaExport = f.exportDecl('const', 'ToKeySchema')
  // f.print`${ToKeySchemaExport} = (item: ItemType): ${KeySchemaImport} => {`
  // f.print` return ItemKeySchema[item.getType().typeName]`
  // f.print`}`



  // ${f.exportDecl('export const ItemKeySchema = new Map<string, KeySchema>([')}`

  // for (const file of schema.files) {
  //   const f = schema.generateFile(file.name + '_pb_key.ts')
  //   // const createRegistry = f.import("createRegistry", "@bufbuild/protobuf");
  //   const { Message, MessageType, JsonValue } = schema.runtime
  //   const MessageAsType = Message.toTypeOnly()

  //   f.print`export const MessageKeySchema = {`
  //   for (const descType of getAllMessages(file)) {
  //     const options = descType.proto.options

  //     if (options && hasExtension(options, key_schema)) {
  //       const key = getExtension(options, key_schema)
  //       key.itemType = descType.typeName
  //       f.print`  "${descType.typeName}" : ${key.toJsonString({
  //         enumAsInteger: true,
  //         prettySpaces: 4,
  //         emitDefaultValues: true,
  //       })},`
  //     }
  //   }
  //   f.print`};`
  // }
}
