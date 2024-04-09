import { Schema } from '@bufbuild/protoplugin/ecmascript'
import {
  AllMessagesWithExtension,
  GetAllTypesFromSchema,
  allMessagesWithExtension,
  getAllMessages,
  getAllTypes,
} from '../utils'

import { KeySchema, key_schema } from '../gen/auth/v1/auth_pb'
import {
  DescExtension,
  DescMessage,
  getExtension,
  hasExtension,
} from '@bufbuild/protobuf'

function getKeySchemaDesc(schema: Schema) {
  for (const t of GetAllTypesFromSchema(schema)) {
    if (t.typeName === 'auth.KeySchema') {
      return t as DescMessage
    }
  }
}

// prettier-ignore
export function generateKeySchema(schema: Schema) {
  const f = schema.generateFile('key_schema.ts')

  const key_schema_desc = getKeySchemaDesc(schema)
  const key_schema_import = !key_schema_desc
    ? f.import('KeySchema', '@saacs/protos-es')
    : f.import(key_schema_desc)

  // prettier-ignore
  f.print(f.exportDecl('const', 'ItemKeySchema'), ': Record<string,', key_schema_import,'> = {',)

  for (const { message, extension } of AllMessagesWithExtension(schema, key_schema,)) {

    extension.itemType =  message.typeName

    f.print`  '${message}': ${key_schema_import}.fromJson(${extension.toJsonString({
        emitDefaultValues: true,
        enumAsInteger: true,
      })}),`
  }

  f.print`}`
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
