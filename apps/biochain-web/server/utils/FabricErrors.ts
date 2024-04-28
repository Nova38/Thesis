import {
  CommitError,
  EndorseError,
  GatewayError,
  SubmitError,
} from '@hyperledger/fabric-gateway'

import { fabricErrors as FE } from '@saacs/client'

function isCommitError(error: Error): error is CommitError {
  return error instanceof CommitError
}

export default function HandleFabricError(error: unknown) {
  if (
    error instanceof GatewayError ||
    error instanceof SubmitError ||
    error instanceof EndorseError
  ) {
    throw {
      type: error.constructor.name,
      message: error.message,
      code: error.code,
      details: error.details,
      cause: { details: error.cause.details, code: error.cause.code },
    }
  }
  if (error instanceof CommitError) {
    throw {
      message: error.message,
      code: error.code,
      txID: error.transactionId,
    }
  } else throw error
}
