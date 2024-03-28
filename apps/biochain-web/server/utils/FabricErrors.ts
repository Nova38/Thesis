import { GatewayError } from '@hyperledger/fabric-gateway'

export default function HandleFabricError(error: unknown) {
  if (error instanceof GatewayError) {
  }
}
