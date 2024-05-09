'use strict'
const hlf = require('saacs-es')
const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger(
  'dos',
)
const { Any, createRegistry, FieldMask } = require('@bufbuild/protobuf')

const { WorkloadModuleBase } = require('@hyperledger/caliper-core')
const { randomInt } = require('node:crypto')

const users = ['User1', 'User2', 'User3', 'User4', 'User5', 'Admin']

function randomTime() {
  return `${Date.now()}-${process.hrtime.bigint()}-${randomInt(0, 1000)}`
}

function randomUpdateArgs() {
  const args = new hlf.pb.pb.UpdateRequest({
    updateMask: new FieldMask({
      paths: ['name'],
    }),
    item: new hlf.pb.pb.Item({
      key: new hlf.pb.pb.ItemKey({
        collectionId: 'collection0',
        itemType: 'auth.Collection',
        itemKeyParts: ['collection0'],
      }),
      value: Any.pack(
        new hlf.pb.pb.Collection({
          collectionId: 'collection0',
          name: `write-${randomInt(0, 100)}`,
        }),
      ),
    }),
  })

  return args.toJsonString({
    typeRegistry: createRegistry(
      ...hlf.pb.sample.allMessages,
      ...hlf.pb.pb.allMessages,
    ),
  })
}

function randomNewBookArgs() {
  Date.now().toString()

  const args = new hlf.pb.pb.CreateRequest({
    item: new hlf.pb.pb.Item({
      value: Any.pack(
        new hlf.pb.sample.Book({
          collectionId: 'collection0',
          isbn: `isbn-${randomTime()}`,
        }),
      ),
    }),
  })
  return args.toJsonString({
    typeRegistry: createRegistry(
      ...hlf.pb.sample.allMessages,
      ...hlf.pb.pb.allMessages,
    ),
  })
}

class NoAuthTestWorkload extends WorkloadModuleBase {
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

    this.chaincodeID = args.chaincodeID
    this.updateAuthChance = args.updateAuthChance

    this.key = new hlf.pb.pb.ItemKey({
      collectionId: 'collection0',
      itemType: 'auth.Collection',
      itemKeyParts: ['collection0'],
    })

    const ReadArg = new hlf.pb.pb.GetRequest({
      key: this.key,
    })

    this.ReadArg = ReadArg.toJsonString({
      typeRegistry: createRegistry(...hlf.pb.sample.allMessages),
    })
  }

  async submitTransaction() {
    /** @type {PeerGateway.FabricRequestSettings} */

    // logger.info(this.updateChance)

    const toUpdateAuth = Math.random() < this.updateAuthChance

    const user = users[randomInt(0, 5)]

    if (toUpdateAuth)
      logger.info(
        `Updating Auth: round ${this.roundIndex}; auth update chances ${this.updateAuthChance}`,
      )

    const args = {
      contractId: this.chaincodeID,
      invokerIdentity: user,
      invokerMspId: 'Org1MSP',

      readOnly: false,
      contractFunction: toUpdateAuth ? 'Update' : 'Create',
      contractArguments: toUpdateAuth
        ? [randomUpdateArgs()]
        : [randomNewBookArgs()],
    }

    return this.sutAdapter.sendRequests(args)
  }
}

function createWorkloadModule() {
  return new NoAuthTestWorkload()
}

module.exports.createWorkloadModule = createWorkloadModule
