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
const fs = require('node:fs')
const { resolve } = require('node:path')

const pathe = require('pathe')
const { pb } = require('@saacs/saacs-pb')
const { CaliperUtils, saacs } = require('@saacs/client')

const { WorkloadModuleBase } = require('@hyperledger/caliper-core')
const PeerGateway = require('@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway')

const chunk = require('lodash.chunk')

const assetDir = pathe.resolve(__dirname, '../assets')

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
    // const  i = import('../../lib')
    const args = this.roundArguments
    // this.numCollections = args.numCollections
    // this.contractVersion = args.contractVersion
    // this.collectionId = args.collectionId

    this.contractId = args.contractId
    this.collectionId = helper.GetCollectionId()
    this.batchSize = this.args?.batchSize || 1000
    this.mode = this.contractId === 'caliper-saacs-binary' ? 'binary' : 'json'
    this.requestOptions = { serializer: this.mode }

    helper.logger.info(
      'Contract ID',
      this.contractId,
      'Contract Version',
      this.contractVersion,
      'Collection ID',
      this.collectionId,
      'Batch Size',
      this.batchSize,
      'Mode',
      this.mode,
      'Asset Dir',
      assetDir,
    )

    this.ArgsBuilder = CaliperUtils.CaliperArgsBuilder({
      contractId: this.contractId,
      invokerMspId: 'Org1MSP',
      invokerIdentity: 'User1',
    })

    const list = new pb.SpecimenList()
    if (this.contractId === 'caliper-saacs-binary') {
      this.mode = 'binary'
      const file = fs.readFileSync(`${assetDir}/specimen.bin`)
      list.fromBinary(file)
    } else {
      this.mode = 'json'
      const file = fs.readFileSync(`${assetDir}/specimen.json`)
      list.fromJsonString(file.toString())
    }

    const req = list.specimens
      .map((specimen) => {
        specimen.collectionId = this.collectionId
        return specimen
      })
      .map((specimen) => {
        return saacs.PrimaryToItem(specimen)
      })
      .map((item) => {
        return this.ArgsBuilder.items.create({ item }, this.requestOptions)
      })

    this.requests = chunk(req, this.batchSize)
    helper.logger.info(
      'Batch size',
      this.batchSize,
      'Number of requests',
      this.requests.length,
    )

    // logger.info(JSON.stringify(saacsConfig));
  }

  /**
   * Assemble TXs for the round.
   * @return {Promise<TxStatus[]>}
   */
  async submitTransaction() {
    /** @type {PeerGateway.FabricRequestSettings} */

    const status = []

    for (const batch of this.requests) {
      const res = await this.sutAdapter.sendRequests(batch)
      helper.logger.info('Batch processed')
      status.push(res)
    }

    helper.logger.info('txStatus')

    return status
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
