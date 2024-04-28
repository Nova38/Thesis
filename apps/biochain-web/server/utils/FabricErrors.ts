import {
  CommitError,
  EndorseError,
  GatewayError,
  SubmitError,
} from '@hyperledger/fabric-gateway'

import { fabricErrors as FE } from '@saacs/client'

export default function HandleFabricError(error: unknown) {
  if (
    error instanceof GatewayError ||
    error instanceof SubmitError ||
    error instanceof EndorseError
  ) {
    const parts = error.message.split('chaincode response 500, ')
    console.error(parts)
    if (parts.length === 2) {
      const msg = JSON.parse(parts[1])
      const e = JSON.parse(msg?.error)
      console.error(e)
      console.info(JSON.stringify(msg, null, 2))
    }

    return {
      type: error.constructor.name,
      message: error.message,
      code: error.code,
      details: error.details,
      cause: { details: error.cause.details, code: error.cause.code },
    }
  }
  if (error instanceof CommitError) {
    return {
      message: error.message,
      code: error.code,
      txID: error.transactionId,
    }
  } else throw error
}
