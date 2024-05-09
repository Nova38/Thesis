'use strict'


const pb = require('@saacs/saacs-pb')
const { CaliperUtils, saacs } = require('@saacs/client')
const { WorkloadModuleBase } = require('@hyperledger/caliper-core')
const PeerGateway = require('@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway')
const helper = require('./helper')

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
    const args = roundArguments
    this.contractId = args.contractId

    helper.logger.info('args', {
      args,
      contractId: args.contractId,
      tContractId: this.contractId,
      workerIndex,
      roundIndex,
      roundArguments,
      sutContext,
    })

    this.ArgsBuilder = CaliperUtils.CaliperArgsBuilder({
      contractId: this.contractId,
      invokerMspId: 'Org1MSP',
      invokerIdentity: 'User1',
    })
  }

  /**
   * Assemble TXs for the round.
   * @return {Promise<TxStatus[]>}
   */
  async submitTransaction() {
    helper.logger.info(this.ArgsBuilder)

    const txStatus = await this.sutAdapter.sendRequests(
      this.ArgsBuilder.utils.getCurrentUser(),
    )

    helper.logger.info('txStatus', txStatus)

    return { txStatus }
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
