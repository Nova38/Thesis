import {
  type Message,
  MethodIdempotency,
  type MethodInfo,
  type MethodInfoUnary,
  type PartialMessage,
  type ServiceType,
} from '@bufbuild/protobuf'
import { GlobalRegistry, chaincode, pb } from '@saacs/saacs-pb'
import type { Contract } from '@hyperledger/fabric-gateway'

// based on connect promise client
export type GatewayClient<T extends ServiceType> = {
  [P in keyof T['methods']]: T['methods'][P] extends MethodInfoUnary<
    infer I,
    infer O
  >
    ? (request: PartialMessage<I>) => Promise<O>
    : never
}

export type BiochainGateway = GatewayClient<
  typeof chaincode.chaincode.ItemService
>
export function createBiochainGateway(contract: Contract): GatewayClient<typeof chaincode.chaincode.ItemService> {
  return createPromiseClient<typeof chaincode.chaincode.ItemService>(chaincode.chaincode.ItemService, contract)
}
export function createUtilGateway(contract: Contract): GatewayClient<typeof chaincode.utils.UtilsService> {
  return createPromiseClient<typeof chaincode.utils.UtilsService>(chaincode.utils.UtilsService, contract)
}
/**
 * Create a PromiseClient for the given service, invoking RPCs through the
 * given transport.
 */
export function createPromiseClient<T extends ServiceType>(
  service: T,
  contract: Contract,
) {
  return Object.fromEntries(
    Object.entries(service.methods).map(([name, method]) => [
      name,
      createContractFn(contract, method as MethodInfoUnary<any, any>),
    ]),
  ) as GatewayClient<T>
}

interface Options {
  serializer: 'json' | 'binary'
}

type ContractFn<I extends Message<I>, O extends Message<O>> = (
  request: PartialMessage<I>,
  options?: Options,
) => Promise<O>

function createContractFn<I extends Message<I>, O extends Message<O>>(
  contract: Contract,
  method: MethodInfo<I, O>,
): ContractFn<I, O> {
  const utf8Decoder = new TextDecoder()

  return async function (input, options) {
    const params
      = (options?.serializer ?? 'json') === 'json'
        ? new method.I(input).toJsonString({ emitDefaultValues: true, typeRegistry: GlobalRegistry })
        : new method.I(input).toBinary()

    const decode = (reply: Uint8Array) =>
      (options?.serializer ?? 'json') === 'json'
        ? method.O.fromJsonString(utf8Decoder.decode(reply), { typeRegistry: GlobalRegistry })
        : method.O.fromBinary(reply)

    console.log('params', params)

    try {
      if (method.idempotency === MethodIdempotency.Idempotent
        || method.idempotency === MethodIdempotency.NoSideEffects) {
        return contract.evaluateTransaction(method.name, params).then((e) => {
          console.log(e)
          return decode(e)
        })
      }

      return contract.submitTransaction(method.name, params).then(decode)
    }
    catch (error) {
      console.error('Error:', error)
      throw error
    }
  }
}
