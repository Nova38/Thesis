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
const { pb, GlobalRegistry } = require('@saacs/saacs-pb')
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
    this.collectionId = args.collectionID
    this.mode = this.contractId === 'caliper-saacs-binary' ? 'binary' : 'json'
    this.requestOptions = { serializer: this.mode }

    this.ArgsBuilder = CaliperUtils.CaliperArgsBuilder({
      contractId: this.contractId,
      invokerMspId: 'Org1MSP',
      invokerIdentity: 'User1',
    })

    // logger.info(JSON.stringify(saacsConfig));
  }

  /**
   * Assemble TXs for the round.
   * @return {Promise<TxStatus[]>}
   */
  async submitTransaction() {
    // /** @type {PeerGateway.} */

    const utf8Decoder = new TextDecoder()

    const spec = new pb.Specimen({
      collectionId: this.collectionId,
    })
    const item = saacs.PrimaryToItem(spec)

    let bookmark = ''
    let lastBookmark = '--'
    while (bookmark !== lastBookmark) {
      helper.logger.info(
        'Still more to fetch',
        'lastBookmark',
        lastBookmark,
        'bookmark',
        bookmark,
        'reqOptions',
        this.requestOptions,
      )

      const req = this.ArgsBuilder.items.listByAttrs(
        {
          key: {
            collectionId: this.collectionId,
            itemKeyParts: [this.collectionId],
            itemType: pb.Book.typeName,
          },
          pagination: { bookmark, pageSize: this.batchSize },
        },
        this.requestOptions,
      )

      // const result = await cc.service.listByAttrs(
      //   new pb.ListByAttrsRequest({
      //     pagination: {
      //       bookmark: bookmark ?? '',
      //       pageSize: query.data.limit ?? 1000,
      //     },
      //     key: new pb.ItemKey({
      //       collectionId: query.data.collectionId,
      //       itemKeyParts: [query.data.collectionId],
      //       itemType: ccbio.Specimen.typeName,
      //     }),
      //     numAttrs: 0,
      //   }),
      // )
      const raw = await this.sutAdapter.sendRequests(req)
      const r = raw.GetResult()

      const result = new pb.ListRequest()

      if (this.mode === 'binary') {
        result.fromBinary(r)
      } else {
        result.fromJsonString(utf8Decoder.decode(reply), {
          typeRegistry: GlobalRegistry,
        })
      }

      helper.logger.info('result pagination', result.pagination.toJsonString())

      lastBookmark = bookmark
      bookmark = result.pagination.bookmark ?? ''
    }
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
