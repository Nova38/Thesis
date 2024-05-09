/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

'use strict'

const { ConfigUtil } = require('@hyperledger/caliper-core')

// const {saacs} = jiti('@saacs/client')
const { pb, GlobalRegistry } = require('@saacs/saacs-pb')
const { CaliperUtils, saacs } = require('@saacs/client')

const { WorkloadModuleBase } = require('@hyperledger/caliper-core')
const ConnectorBase = require('@hyperledger/caliper-core/lib/common/core/connector-base')
const PeerGateway = require('@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway')

const helper = require('../helper')

/**
 * Workload module for the benchmark round.
 * @type {CreateWorkload}
 * @property {string} contractId The name of the contract.
 * @property {string} contractVersion The version of the contract.
 */
class BootstrapWorkload extends WorkloadModuleBase {
  /**
   * Initializes the workload module instance.
   */
  constructor() {
    super()
    this.contractId = ''
    this.contractVersion = ''
  }

  /**
   * Initialize the workload module with the given parameters.
   * @param {number} workerIndex The 0-based index of the worker instantiating the workload module.
   * @param {number} totalWorkers The total number of workers participating in the round.
   * @param {number} roundIndex The 0-based index of the currently executing round.
   * @param {object} roundArguments The user-provided arguments for the round from the benchmark configuration file.
   * @param {PeerGateway} sutAdapter The adapter of the underlying SUT.
   * @param {object} sutContext The custom context object provided by the SUT adapter.
   * @async
   */
  async initializeWorkloadModule(
    workerIndex,
    totalWorkers,
    roundIndex,
    roundArguments,
    sutAdapter,
    sutContext,
  ) {
    await super.initializeWorkloadModule(
      workerIndex,
      totalWorkers,
      roundIndex,
      roundArguments,
      sutAdapter,
      sutContext,
    )
    const args = this.roundArguments
    this.collectionId = args.collectionID
    this.contractId = args.contractId
    // this.collectionId = helper.GetCollectionId()

    this.mode = this.contractId === 'caliper-saacs-binary' ? 'binary' : 'json'

    this.ArgsBuilder = CaliperUtils.CaliperArgsBuilder({
      contractId: this.contractId,
      invokerMspId: 'Org1MSP',
      invokerIdentity: 'User1',
    })

    this.bootstrap = this.ArgsBuilder.utils.bootstrap({
      collection: {
        authType: pb.AuthType.IDENTITY,
        name: this.collectionId,
        collectionId: this.collectionId,
        itemTypes: [
          pb.Book.typeName,
          pb.SimpleItem.typeName,
          pb.Collection.typeName,
          pb.UserDirectMembership.typeName,
        ],
        default: {
          defaultPolicy: {
            actions: saacs.ActionsAlias.all.full,
          },
        },
      },
    })
  }

  /**
   * Assemble TXs for the round.
   * @return {Promise<TxStatus[]>}
   */
  async submitTransaction() {
    this.bootstrap.readOnly = false
    const txStatus = await this.sutAdapter.sendRequests(this.bootstrap)
    helper.logger.info('txStatus', txStatus)

    return txStatus
  }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
  return new BootstrapWorkload()
}

module.exports.createWorkloadModule = createWorkloadModule
