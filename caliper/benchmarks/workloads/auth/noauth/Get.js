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

"use strict";
const { Any, createRegistry } = require('@bufbuild/protobuf');

const { WorkloadModuleBase } = require("@hyperledger/caliper-core");
const ConnectorBase = require("@hyperledger/caliper-core/lib/common/core/connector-base");
const PeerGateway = require("@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway");
const hlf = require("saacs-es");
const logger = require("@hyperledger/caliper-core").CaliperUtils.getLogger(
    "my-module"
);

/**
 * Workload module for the benchmark round.
 * @type {CreateWorkload}
 * @property {string} contractId The name of the contract.
 * @property {string} contractVersion The version of the contract.
 */
class GetWorkload extends WorkloadModuleBase {
    /**
     * Initializes the workload module instance.
     */
    constructor() {
        super();
        this.contractId = "";
        this.contractVersion = "";
    }

    /**
     * Initialize the workload module with the given parameters.
     * @param {number} workerIndex The 0-based index of the worker instantiating the workload module.
     * @param {number} totalWorkers The total number of workers participating in the round.
     * @param {number} roundIndex The 0-based index of the currently executing round.
     * @param {Object} roundArguments The user-provided arguments for the round from the benchmark configuration file.
     * @param {PeerGateway} sutAdapter The adapter of the underlying SUT.
     * @param {Object} sutContext The custom context object provided by the SUT adapter.
     * @async
     */
    async initializeWorkloadModule(
        workerIndex,
        totalWorkers,
        roundIndex,
        roundArguments,
        sutAdapter,
        sutContext
    ) {
        await super.initializeWorkloadModule(
            workerIndex,
            totalWorkers,
            roundIndex,
            roundArguments,
            sutAdapter,
            sutContext
        );
        // const  i = import('../../lib')
        const args = this.roundArguments;
        this.contractId = args.contractId;
        this.numCollections = args.numCollections;
        this.contractVersion = args.contractVersion;
        this.type = args.type;
    }

    /**
     * Assemble TXs for the round.
     * @return {Promise<TxStatus[]>}
     */
    async submitTransaction() {
        let arg = null;
        if (!this.type) {
            logger.error("unknown type", this.args);

            arg = new hlf.pb.common.generic.GetRequest({
                type: hlf.pb.sample.SimpleItem.typeName,
                key: new hlf.pb.auth.ItemKey({

                itemIdParts: [
                    `item-${hlf.utils.factory.randomInt(1000)}`,
                ],
                collectionId: hlf.utils.factory.modCollectionId(
                    this.numCollections,
                    this.workerIndex
                ),
                }),
            });
        } else if (this.type == "simple") {
            arg = new hlf.pb.common.generic.GetRequest({
                key: new hlf.pb.auth.ItemKey({
                    type: hlf.pb.sample.SimpleItem.typeName,
                    itemIdParts: [
                        `item-${hlf.utils.factory.randomInt(1000)}`,
                    ],
                    collectionId: hlf.utils.factory.modCollectionId(
                        this.numCollections
                    ),
                }),
            });
        } else if (this.type == "book") {
            arg = new hlf.pb.common.generic.GetRequest({
                key: new hlf.pb.auth.ItemKey({
                    type: hlf.pb.sample.Book.typeName,
                    itemIdParts: [
                        `book-${hlf.utils.factory.randomInt(1000)}`,
                    ],
                    collectionId: hlf.utils.factory.modCollectionId(
                        this.numCollections,
                        this.workerIndex
                    ),
                }),
            });
        }

        // logger.info('this', this)
        logger.info('arg', arg);

        /** @type {PeerGateway.FabricRequestSettings}*/
        const myArgs = {
            contractId: this.contractId,

            contractFunction: "Get",
            contractArguments: [
                arg.toJsonString({
                    typeRegistry: createRegistry(...hlf.pb.sample.allMessages),
                }),
            ],
            readOnly: true,
            invokerIdentity: "User1",
            invokerMspId: "Org1MSP",
        };

        return await this.sutAdapter.sendRequests(myArgs);
    }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
    return new GetWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
