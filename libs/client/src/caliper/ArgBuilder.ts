import {
  type Message,
  MethodIdempotency,
  type MethodInfo,
  type MethodInfoUnary,
  type PartialMessage,
  type ServiceType,
} from '@bufbuild/protobuf'
import { GlobalRegistry, chaincode } from '@saacs/saacs-pb'

// export interface CaliperArgs {
//   contractId: string
//   contractFunction: string
//   contractArguments: string[]
//   invokerIdentity?: string
//   mspId?: string
//   readonly?: boolean
// }

export interface ContractArgs {
  contractFunction: string
  contractArguments: string[]
}

export type CaliperClient<T extends ServiceType> = {
  [P in keyof T['methods']]: T['methods'][P] extends MethodInfoUnary<
    infer I,
    infer _O
  >
    ? (request: PartialMessage<I>) => ContractArgs
    : never
}

interface Options {
  serializer: 'json' | 'binary'
}
type ContractAgsFn<I extends Message<I>> = (
  request: PartialMessage<I>,
  options?: Options
) => ContractArgs

function createContractArgsFn<I extends Message<I>>(
  method: MethodInfo<I>,
): ContractAgsFn<I> {
  const utf8Decoder = new TextDecoder()

  return function (input, options) {
    const params
      = (options?.serializer ?? 'json') === 'json'
        ? new method.I(input).toJsonString({ emitDefaultValues: true, typeRegistry: GlobalRegistry })
        : utf8Decoder.decode(new method.I(input).toBinary())

    return {
      contractFunction: method.name,
      contractArguments: [params],
    }
  }
}

function buildCaliperClient<T extends ServiceType>(
  service: T,
) {
  return Object.fromEntries(
    Object.entries(service.methods).map(([name, method]) => [
      name,
      createContractArgsFn(method as MethodInfoUnary<any, any>),
    ]),
  ) as CaliperClient<T>
}

export function CaliperItemArgsBuilder() {
  return Object.fromEntries([
    ...Object.entries(buildCaliperClient(chaincode.chaincode.ItemService)),
  ]) as CaliperClient<typeof chaincode.chaincode.ItemService>
}

export function CaliperUtilsArgsBuilder() {
  return Object.fromEntries([
    ...Object.entries(buildCaliperClient(chaincode.utils.UtilsService)),
  ]) as CaliperClient<typeof chaincode.utils.UtilsService>
}

function testing() {
  const itemArgsBuilder = CaliperItemArgsBuilder()
  const utilsArgsBuilder = CaliperUtilsArgsBuilder()

  itemArgsBuilder.get({ key: { collectionId: '1', itemId: '1' } })
}
