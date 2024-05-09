import {
  type Message,
  MethodIdempotency,
  type MethodInfo,
  type MethodInfoUnary,
  type PartialMessage,
  type ServiceType,
} from '@bufbuild/protobuf'
import type { pb } from '@saacs/saacs-pb'
import { GlobalRegistry, chaincode } from '@saacs/saacs-pb'

export type ItemKind = pb.ItemKind

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
  contractArguments: (string | Uint8Array)[]
  readOnly: boolean
  invokerMspId?: string
  invokerIdentity?: string
  contractId?: string
}

export interface CaliperContractOptions {
  invokerMspId?: string
  invokerIdentity?: string
  contractId?: string
}

interface RequestOptions {
  serializer: 'json' | 'binary'
}

export type CaliperClient<T extends ServiceType> = {
  [P in keyof T['methods']]: T['methods'][P] extends MethodInfoUnary<
    infer I,
    infer _O
  >
    ? (request: PartialMessage<I>, options?: RequestOptions) => ContractArgs
    : never
}

export interface FullCaliperClient {
  utils: CaliperClient<typeof chaincode.utils.UtilsService>
  items: CaliperClient<typeof chaincode.chaincode.ItemService>
}

export type calipeUtilsArg = CaliperClient<typeof chaincode.chaincode.ItemService>
export type caliperItem = CaliperClient<typeof chaincode.utils.UtilsService>

type ContractAgsFn<I extends Message<I>> = (
  request: PartialMessage<I>,
  options?: RequestOptions
) => ContractArgs

function createContractArgsFn<I extends Message<I>>(
  method: MethodInfo<I>,
  ContractOptions?: CaliperContractOptions,
): ContractAgsFn<I> {
  const utf8Decoder = new TextDecoder()

  return function (input, options) {
    const params
      = (options?.serializer ?? 'json') === 'json'
        ? new method.I(input).toJsonString({ emitDefaultValues: true, typeRegistry: GlobalRegistry })
        : new method.I(input).toBinary()

    const readOnly = method.idempotency === MethodIdempotency.Idempotent
      || method.idempotency === MethodIdempotency.NoSideEffects

    return {
      contractFunction: method.name,
      contractArguments: [params],
      readOnly,
      ...ContractOptions,
    }
  }
}

function buildCaliperClient<T extends ServiceType>(
  service: T,
  options?: CaliperContractOptions,
) {
  return Object.fromEntries(
    Object.entries(service.methods).map(([name, method]) => [
      name,
      createContractArgsFn(method as MethodInfoUnary<any, any>, options),
    ]),
  ) as CaliperClient<T>
}

export function CaliperItemArgsBuilder(options?: CaliperContractOptions): CaliperClient<typeof chaincode.chaincode.ItemService> {
  return Object.fromEntries([
    ...Object.entries(buildCaliperClient(chaincode.chaincode.ItemService, options)),
  ]) as CaliperClient<typeof chaincode.chaincode.ItemService>
}

export function CaliperUtilsArgsBuilder(options?: CaliperContractOptions): CaliperClient<typeof chaincode.utils.UtilsService> {
  return Object.fromEntries([
    ...Object.entries(buildCaliperClient(chaincode.utils.UtilsService, options)),
  ]) as CaliperClient<typeof chaincode.utils.UtilsService>
}

export function CaliperArgsBuilder(options?: CaliperContractOptions): FullCaliperClient {
  return {
    utils: CaliperUtilsArgsBuilder(options),
    items: CaliperItemArgsBuilder(options),
  }
}

function testing() {
  const itemArgsBuilder = CaliperItemArgsBuilder()
  const utilsArgsBuilder = CaliperUtilsArgsBuilder()

  const a = itemArgsBuilder.get({ key: { collectionId: '1' } })
}
