


const {WorkloadModuleBase} = require('@hyperledger/caliper-core');
const ConnectorBase = require('@hyperledger/caliper-core/lib/common/core/connector-base');
const PeerGateway = require('@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway');
const hlf = require('hlf_tools')
const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger('workload');


/**
 * Workload module for the benchmark round.
 *
 * @type {referenceReferenceWorkload}
 * @property {string} contractId The name of the contract.
 * @property {string} contractVersion The version of the contract.
 */
class referenceReferenceWorkload extends WorkloadModuleBase {

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
    const args = this.roundArguments;
    this.contractId = args.contractId;
    this.user = args.user;
    this.userOrg = args.userOrg;
    this.item = args.item;



    this.contractVersion = args.contractVersion;
}
// workload path hlf.pb.common.reference
// Reference hlf.pb.common.reference.ReferenceRequest

/**
 * Assemble TXs for the round.
 * @return {Promise<TxStatus[]>}
 */
async submitTransaction() {
    /** @type {PeerGateway.FabricRequestSettings}*/

    const item = new hlf.pb.common.reference.ReferenceRequest({})


    const myArgs = {
        contractId: this.contractId,

        contractFunction: 'Reference',
        contractArguments: [
            item.toJsonString({typeRegistry: hlf.utils.GlobalRegistry})
        ],
        readOnly: false,
        invokerIdentity: this.user,
        invokerMspId: this.UserOrg,
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

module.exports.createWorkloadModule = createWorkloadModule;
