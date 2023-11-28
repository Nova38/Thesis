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

'use strict';

// import { WorkloadModuleBase } from '@hyperledger/caliper-core';
// import { Operation } from '../../lib';

//
const {WorkloadModuleBase} = require('@hyperledger/caliper-core');
const ConnectorBase = require('@hyperledger/caliper-core/lib/common/core/connector-base');
const PeerGateway = require('@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway');
const lib = require('es')

const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger('my-module');



/**
 *  @typedef {GetWorkload} SpecialType
 * 
 * @property {PeerGateway} othe
 */


/**
 * Workload module for the benchmark round.
 * @type {GetWorkload}
 * @property {string} contractId The name of the contract.
 * @property {string} contractVersion The version of the contract.
 */
class GetWorkload extends WorkloadModuleBase {

    /**
     * Initializes the workload module instance.
     */
    constructor() {
        super();
        this.contractId = '';
        this.contractVersion = '';
        
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
    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);
        // const  i = import('../../lib')
        const args = this.roundArguments;
        this.contractId = args.contractId;

        lib.gen.auth.v1.auth_pb.Action.CREATE

        this.contractVersion = args.contractVersion;
    }

    /**
     * Assemble TXs for the round.
     * @return {Promise<TxStatus[]>}
     */
    async submitTransaction() {
        /** @type {PeerGateway.FabricRequestSettings}*/ 
        const myArgs = {
            contractId: this.contractId,
            
            contractFunction: 'Test',
            contractArguments: [],
            readOnly: false,
            invokerIdentity: "User1",
            invokerMspId: "Org1MSP",
        };


        const txStatus = await this.sutAdapter.sendRequests(myArgs);
        
        logger.info('txStatus', txStatus);

        return txStatus;
    }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
    return new GetWorkload();
}

// const _createWorkloadModule = createWorkloadModule;
// export { _createWorkloadModule as createWorkloadModule };


module.exports.createWorkloadModule = createWorkloadModule;
