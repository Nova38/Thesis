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
    this.Users = args.users
    this.UserMap = {}
    if (!this.Users) {
      helper.logger.info(
        'Users not found in the configuration, setting new Users value',
      )
      this.Users = []
    }

    this.ArgsBuilder = CaliperUtils.CaliperArgsBuilder({
      contractId: this.contractId,
      invokerMspId: 'Org1MSP',
      invokerIdentity: 'User1',
    })

    this.requests = this.Users.map((user) => {
      const util = CaliperUtils.CaliperArgsBuilder({
        contractId: this.contractId,
        invokerMspId: 'Org1MSP',
        invokerIdentity: user,
      })
      return util.utils.getCurrentUser()
    })
  }

  /**
   * Assemble TXs for the round.
   * @return {Promise<TxStatus[]>}
   */
  async submitTransaction() {
    helper.logger.info(this.ArgsBuilder, this.Users)
    const utf8Decoder = new TextDecoder()

    const results = {}
    for (const user of this.Users) {
      const util = CaliperUtils.CaliperArgsBuilder({
        contractId: this.contractId,
        invokerMspId: 'Org1MSP',
        invokerIdentity: user,
      })
      const req = util.utils.getCurrentUser()

      const txStatus = await await this.sutAdapter.sendRequests(req)
      const txResponse = utf8Decoder.decode(txStatus.GetResult())
      helper.logger.info('txResponse', txResponse)

      const result = new pb.chaincode.utils.GetCurrentUserResponse()
      result.fromJsonString(txResponse)

      results[user] = result.user
    }
    helper.logger.info('Results', results)
    helper.SetUserIds(results)
    // const txStatus = await this.sutAdapter.sendRequests(this.requests)
    // helper.logger.info('txStatus', txStatus)

    // const res = await txStatus.map((tx) => {
    //   const txResponse = tx.GetResult()
    //   helper.logger.info('txResponse', txResponse)
    //   const result = new pb.chaincode.utils.GetCurrentUserResponse()
    //   result.fromJsonString(utf8Decoder.decode(tx.GetResult()))
    //   return result
    // })
    // helper.logger.info('Responses', res)

    // this.UserMap = res.reduce((acc, user) => {
    //   acc[user.getId()] = user
    //   return acc
    // }, {})
    // helper.logger.info('this.UserMap', this.UserMap)

    return txStatus
  }

  // cleanupWorkloadModule() {
  //   helper.logger.info('cleanupWorkloadModule', 'setting', this.UserMap)
  //   helper.SetUserIds(this.UserMap)
  // }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
  return new BootstrapWorkload()
}

module.exports.createWorkloadModule = createWorkloadModule
