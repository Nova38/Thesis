import { CommitError, CommitStatusError, EndorseError, GatewayError, SubmitError } from '@hyperledger/fabric-gateway'

export type FabricError = CommitError | GatewayError | EndorseError | SubmitError | CommitStatusError

export class AuthError extends Error {
  constructor(message: FabricError) {
    super(message.message)
    this.name = 'AuthError'
  }
}

export function isCommitError(error: Error): error is CommitError {
  return error instanceof CommitError
}
export function isGatewayError(error: Error): error is GatewayError {
  return error instanceof GatewayError
}
export function isEndorseError(error: Error): error is EndorseError {
  return error instanceof EndorseError
}
export function isSubmitError(error: Error): error is SubmitError {
  return error instanceof SubmitError
}
export function isCommitStatusError(error: Error): error is CommitStatusError {
  return error instanceof CommitStatusError
}

export function ProcessFabricError(error: Error) {
  switch (true) {
    case isCommitError(error):
      console.error('CommitError:', error)
      break
    case isGatewayError(error):
      console.error('GatewayError:', error)
      break
    case isEndorseError(error):
      console.error('EndorseError:', error)
      break
    case isSubmitError(error):
      console.error('SubmitError:', error)
      break
    case isCommitStatusError(error):
      console.error('CommitStatusError:', error)
      break

    case error instanceof Error:
      console.error('UnknownError:', error)
      break
  }
}
