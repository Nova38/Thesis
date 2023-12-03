#!/usr/bin/env -S npx tsx

// See here how to run this plugin: https://github.com/bufbuild/protobuf-es/tree/main/packages/protoplugin-example

import { createEcmaScriptPlugin, runNodeJs } from "@bufbuild/protoplugin";
import { ImportSymbol, Schema, findCustomMessageOption } from "@bufbuild/protoplugin/ecmascript";
import type {
    AnyMessage,
    DescEnum,
    DescExtension,
    DescFile,
    DescMessage,
    DescService,
    FileDescriptorSet,


} from "@bufbuild/protobuf";
import {
    literalString,
    makeJsDoc,
    localName,
} from "@bufbuild/protoplugin/ecmascript";

import { Empty, MethodKind, createDescriptorSet, createRegistryFromDescriptors } from "@bufbuild/protobuf";


import * as fs from 'fs';


const registry = createRegistryFromDescriptors(
    fs.readFileSync("image.bin")
  );

const protocGenReg = createEcmaScriptPlugin({
    name: "protoc-gen-reg",
    version: `v1`,
    generateTs,
    //   parseOptions:
});




runNodeJs(protocGenReg);


function generateTs(schema: Schema) {
    generateFabricTest(schema)

}




// const inputType = registry.findMessage(method.input.typeName)

const header = `

`
function generateFabricTest(schema: Schema) {
    for (const file of schema.files) {
        // const f = schema.generateFile(file.name + ".fabric");
        // const registry = createRegistryFromDescriptors(createDescriptorSet(schema.proto.sourceFileDescriptors) )


        // f.preamble(file);
        const {
            Message,
            JsonValue
        } = schema.runtime;
        // Convert the Message ImportSymbol to a type-only ImportSymbol

        for (const service of file.services) {
            const localServiceName = localName(service);

            const p = file.name.split("/")
            const base = p.pop() || ""
            const pbase = p.pop() || ""



            for (const method of service.methods) {
                localName(method)

                const mFile  = schema.generateFile(file.name + "/" + method.name+ ".js");


                const hlf_path = `hlf.pb.${pbase}.${base}`
                const inputType = `${hlf_path}.${method.input.name}`
                // const inputType = registry.findMessage(method.input.typeName)

                mFile.print`


const {WorkloadModuleBase} = require('@hyperledger/caliper-core');
const ConnectorBase = require('@hyperledger/caliper-core/lib/common/core/connector-base');
const PeerGateway = require('@hyperledger/caliper-fabric/lib/connector-versions/peer-gateway/PeerGateway');
const hlf = require('hlf_tools')
const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger('workload');


/**
 * Workload module for the benchmark round.
 *
 * @type {${base}${method.name}Workload}
 * @property {string} contractId The name of the contract.
 * @property {string} contractVersion The version of the contract.
 */
class ${base}${method.name}Workload extends WorkloadModuleBase {

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
// workload path ${hlf_path}
// ${method.name} ${inputType}

/**
 * Assemble TXs for the round.
 * @return {Promise<TxStatus[]>}
 */
async submitTransaction() {

    const item = new ${inputType}({arg.item})

    /** @type {PeerGateway.FabricRequestSettings}*/
    const myArgs = {
        contractId: this.contractId,

        contractFunction: '${method.name}',
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
                `








            }
        }
    }
}



function generateWorkload(schema: Schema) {
    for (const file of schema.files) {
        const f = schema.generateFile(file.name + "_reg");
        f.preamble(file);

        for (const service of file.services) {
            const localServiceName = localName(service);

            // const f = schema.generateFile(file.name +localServiceName+ "_reg.ts");
            // f.preamble(file);
            // for (const method of service.methods) {
                // const f = schema.generateFile(file.name + localServiceName + "_" + method + ".workload.ts");
                // f.preamble(file);

                // f.print(header)

                // console.log(method.input.typeName)
                // f.print`module.exports.createWorkloadModule = createWorkloadModule;`
            }
        }
    }

// }
        // const registry = createRegistryFromDescriptors(createDescriptorSet(schema.proto.sourceFileDescriptors) )
