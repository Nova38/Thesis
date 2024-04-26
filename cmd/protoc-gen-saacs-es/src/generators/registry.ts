import { Schema } from '@bufbuild/protoplugin/ecmascript'
import {
  GetAllTypesFromSchema,
  getAllEnum,
  getAllMessages,
  getAllTypes,
} from '../utils'

export function GenTypes(schema: Schema) {
  let f = schema.generateFile('types_pb.ts')

  f.print`export default {`
  for (const file of schema.files) {
    for (const descType of getAllMessages(file)) {
      f.print`${descType},`
    }
    for (const descType of getAllEnum(file)) {
      f.print`${descType},`
    }
  }
  f.print`}`
  f.print`export {`
  for (const file of schema.files) {
    for (const descType of getAllMessages(file)) {
      f.print`${descType},`
    }
    for (const descType of getAllEnum(file)) {
      f.print`${descType},`
    }
  }
  f.print`}`
  f.print``
}
export function generateRegistry(schema: Schema) {
  let f = schema.generateFile('global_reg.ts')

  let ExpReg = f.exportDecl('const', 'GlobalRegistry')

  f.print`import { createRegistry } from "@bufbuild/protobuf";`
  f.print`${ExpReg}: ${schema.runtime.IMessageTypeRegistry} = createRegistry(`
  for (const descType of GetAllTypesFromSchema(schema)) {
    if (descType.kind == 'message') {
      f.print`  ${descType}, `
    }
  }
  f.print`);`
  f.print``
}
