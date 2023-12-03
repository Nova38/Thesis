'use strict';
const hlf = require('hlf_tools')
const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger('noauth-workload');


const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

class MyWorkload extends WorkloadModuleBase {
    async submitTransaction() {
        let txArgs = {
            // TX arguments for "mycontract"
        };

        return this.sutAdapter.invokeSmartContract('', 'v1', txArgs, 30);
    }
}

function createWorkloadModule() {
    return new MyWorkload();
}

module.exports.createWorkloadModule = createWorkloadModule;
