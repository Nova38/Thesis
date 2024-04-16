import {
  AnyMessage,
  DescEnum,
  DescExtension,
  DescFile,
  DescMessage,
  DescService,
  Extension,
  Message,
  createRegistryFromDescriptors,
  getExtension,
  hasExtension,
} from '@bufbuild/protobuf'
import { Schema } from '@bufbuild/protoplugin/ecmascript'
import * as fs from 'fs'
import { fileURLToPath } from 'url'

import { resolve, dirname } from 'pathe'
import { y } from 'happy-dom/lib/PropertySymbol'

const __dirname = dirname(fileURLToPath(import.meta.url))
const registryBin = resolve(__dirname, 'image.bin')

export const registry = createRegistryFromDescriptors(
  fs.readFileSync(registryBin),
)

export function* GetAllTypesFromSchema(schema: Schema) {
  for (const file of schema.files) {
    for (const descType of getAllTypes(file)) {
      yield descType
    }
  }
}

export function* getAllTypes(
  desc: DescFile | DescMessage,
): Iterable<DescMessage | DescEnum | DescExtension | DescService> {
  switch (desc.kind) {
    case 'file':
      for (const message of desc.messages) {
        yield message
        yield* getAllTypes(message)
      }
      yield* desc.enums
      yield* desc.services
      yield* desc.extensions
      break
    case 'message':
      for (const message of desc.nestedMessages) {
        yield message
        yield* getAllTypes(message)
      }
      yield* desc.nestedExtensions
      yield* desc.nestedEnums
      break
  }
}
export function* getAllMessages(
  desc: DescFile | DescMessage,
): Iterable<DescMessage> {
  switch (desc.kind) {
    case 'file':
      for (const message of desc.messages) {
        yield message
        yield* getAllMessages(message)
      }
      break
    case 'message':
      for (const message of desc.nestedMessages) {
        yield message
        yield* getAllMessages(message)
      }
      break
  }
}

export function* allMessagesWithExtension<E extends Message<E>, V>(
  desc: DescFile | DescMessage,
  extensionType: Extension<any, V>,
): Iterable<{ message: DescMessage; extension: V }> {
  for (const message of getAllMessages(desc)) {
    const options = message.proto.options
    if (options && hasExtension(options, extensionType)) {
      const extension = getExtension(options, extensionType)
      yield { message, extension }
    }
  }
}

export function* AllMessagesWithExtension<E extends Message<E>, V>(
  schema: Schema,
  extensionType: Extension<any, V>,
): Iterable<{ message: DescMessage; extension: V }> {
  for (const file of schema.files) {
    yield* allMessagesWithExtension(file, extensionType)
  }
}
