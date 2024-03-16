/* eslint-disable ts/await-thenable */
import { GatewayError } from '@hyperledger/fabric-gateway'
import type { EventHandler, EventHandlerRequest } from 'h3'

export function defineFabricWrappedResponseHandler<
  T extends EventHandlerRequest,
  D,
>(handler: EventHandler<T, D>): EventHandler<T, D> {
  return defineEventHandler<T>(async (event) => {
    console.log('defineFabricWrappedResponseHandler', event)
    try {
      const response = await handler(event)
      return { response: response as D }
    }
    catch (err) {
      if (err instanceof GatewayError) {
        return createError({
          cause: err,
        })
      }

      return { err }
    }
  })
}
