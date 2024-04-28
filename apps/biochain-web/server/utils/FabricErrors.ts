import { CommitError, GatewayError } from '@hyperledger/fabric-gateway'

import { fabricErrors as FE } from '@saacs/client'


function isCommitError(error: Error): error is CommitError {
  return error instanceof CommitError
}

export default function HandleFabricError(error: unknown) {

  switch (true) {
    case :
      console.error('CommitError:', error)

      break
    case FE.isGatewayError(error):
      console.error('GatewayError:', error)
      break
    case FE.isEndorseError(error):
      console.error('EndorseError:', error)
      break
    case FE.isSubmitError(error):
      console.error('SubmitError:', error)
      break
    case FE.isCommitStatusError(error):
      console.error('CommitStatusError:', error)
      break

    case error instanceof Error:
      console.error('UnknownError:', error)
      break
  }
}
