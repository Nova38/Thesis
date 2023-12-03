'use strict';
const hlf = require('hlf_tools')
const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger('noauth-workload');


const { WorkloadModuleBase } = require('@hyperledger/caliper-core');
const { randomInt } = require('crypto');
const users = ['User1', 'User2', 'User3', 'User4', 'User5', 'Admin']


class NoAuthTestWorkload extends WorkloadModuleBase {
    async submitTransaction() {

        if (!this.roundArguments.user){
            let i = randomInt(0, 5)
            this.roundArguments.user = users[i]
        }

         /** @type {PeerGateway.FabricRequestSettings}*/
         const myArgs = {
            contractId: "noauth",

            contractFunction: 'Test',
            contractArguments: [],
            readOnly: true,
            invokerIdentity: this.roundArguments.user ,
            invokerMspId: "Org1MSP",
        };


        if (randomInt(0, 100) > this.roundArguments.failRate){

            myArgs.contractFunction = 'TestFail'

        }
        return this.sutAdapter.sendRequests(myArgs);
    }
}

function createWorkloadModule() {
    return new NoAuthTestWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
