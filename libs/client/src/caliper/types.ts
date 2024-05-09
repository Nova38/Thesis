import type { Buffer } from 'node:buffer'

/**
 * @typedef {object} FabricRequestSettings
 *
 * @property {string} channel Required. The name of the channel whose contract should be invoked.
 * @property {string} contractId Required. The name/ID of the contract whose function
 *           should be invoked.
 * @property {string} contractFunction Required. The name of the function that should be
 *           invoked in the contract.
 * @property {string[]} [contractArguments] Optional. The list of {string} arguments that should
 *           be passed to the contract.
 * @property {Map<string, Buffer>} [transientMap] Optional. The transient map that should be
 *           passed to the contract.
 * @property {string} [invokerMspId] Optional. The MspId of the invoker. Required if there are more than
 *           1 organisation defined in the network configuration file
 * @property {string} invokerIdentity Required. The identity name of the invoker
 * @property {boolean} [readOnly] Optional. Indicates whether the request is a submit or evaluation.
 *           contract. If an admin is needed, use the organization name prefixed with a _ symbol.
 */

export interface FabricRequestSettings {
  channel: string
  contractId: string
  contractFunction: string
  contractArguments: string[]
  transientMap: Map<string, Buffer>
  invokerMspId: string
  invokerIdentity: string
  readOnly: boolean
}
