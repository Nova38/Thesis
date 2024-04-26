import {
  GeneratedFile,
  Schema,
  localName,
} from '@bufbuild/protoplugin/ecmascript'
import type {
  DescFile,
  DescMessage,
  DescMethod,
  DescService,
  FileDescriptorSet,
} from '@bufbuild/protobuf'

import {
  Empty,
  MethodKind,
  createDescriptorSet,
  createRegistryFromDescriptors,
  getExtension,
  hasExtension,
} from '@bufbuild/protobuf'

import * as fs from 'fs'
import { TransactionType } from '../gen/saacs/common/v0/enums_pb'
import { transaction_type } from '../gen/saacs/common/v0/options_pb'

// Modified from the original protoc-gen-ts plugin
export function generateGateway(schema: Schema) {
  for (const file of schema.files) {
    generateFile(schema, file)
  }
}

function generateFile(schema: Schema, file: DescFile) {
  if (file.services.length == 0) return

  const f = schema.generateFile(file.name + '_pb_gateway.ts')

  const contract = f.import('Contract', '@hyperledger/fabric-gateway')
  const registry = f.import('IMessageTypeRegistry', '@bufbuild/protobuf')
  const jsonWriteOptions = f.import(
    'JsonWriteStringOptions',
    '@bufbuild/protobuf',
  )
  // const jsonReadOptions = f.import(
  //     "JsonReadOptions",
  //     "@bufbuild/protobuf"
  // );

  f.print('const utf8Decoder = new TextDecoder();')
  // f.import("TextDecoder")
  // f.preamble(file);
  const { Message, JsonValue, IMessageTypeRegistry } = schema.runtime
  // Convert the Message ImportSymbol to a type-only ImportSymbol
  const MessageAsType = Message.toTypeOnly()
  for (const service of file.services) {
    const localServiceName = localName(service)
    f.print(f.jsDoc(service))

    f.print`${f.exportDecl('class', `${localServiceName}Client`)}  {`
    f.print`    private contract: ${contract};`
    f.print`    private jsonWriteOptions:Partial<${jsonWriteOptions}> = {};`
    f.print`    registry: ${registry};`
    f.print()
    f.print`    constructor(contract: ${contract}, registry: ${registry}) {`
    f.print`        this.contract = contract;`
    f.print`        this.registry = registry;`
    f.print`        this.jsonWriteOptions.typeRegistry = registry`
    f.print`    }`
    f.print()
    for (const method of service.methods) {
      genMethod(schema, method, f)
    }
    f.print`}`
    f.print``
  }
}

function genMethod(schema: Schema, method: DescMethod, f: GeneratedFile) {
  if (method.methodKind === MethodKind.Unary) {
    f.print()
    f.print(f.jsDoc(method, '    '))

    const { PartialMessage } = schema.runtime

    const tt =
      method.proto.options &&
      hasExtension(method.proto.options, transaction_type)
        ? getExtension(method.proto.options, transaction_type)
        : undefined

    const callMethod =
      tt === TransactionType.INVOKE
        ? 'submitTransaction'
        : 'evaluateTransaction'

    // prettier-ignore
    if (method.input.typeName == "google.protobuf.Empty") {
      f.print`  async ${localName(method)}(): Promise<${method.output}> {`
      f.print`    const results = utf8Decoder.decode(`
      f.print`      await this.contract.${callMethod}(${f.string(method.name)})`
      f.print`    )`
      f.print`    return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`
      f.print`  }`
      f.print``
    } else {
          f.print`  async ${localName(method)}(req: ${PartialMessage}<${method.input}>): Promise<${method.output}> {`
          f.print`    const msg = req instanceof ${method.input} ? req : new ${method.input}(req);`
          f.print`    const results = utf8Decoder.decode(`
          f.print`      await this.contract.${callMethod}(`
          f.print`        ${f.string(method.name)},`
          f.print`        msg.toJsonString(this.jsonWriteOptions)`
          f.print`      ))`
          f.print`    return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`
          f.print`    }`
          f.print``
        }
  }
}
