import { pb } from '@saacs/saacs-pb'
import {
  WorkloadModuleBase,
} from '@hyperledger/caliper-core'

const x = pb.Book.fromJson({})

console.log('Hello, Caliper!')
console.log(x.toJsonString())

export function test() {
  console.log('Hello, Caliper!')
}

WorkloadModuleBase
// ^?

class MyWorkload extends WorkloadModuleBase {
  async submitTransaction() {
    const txArgs = {
      // TX arguments for "mycontract"
    }

    return this.sutAdapter.invokeSmartContract('mycontract', 'v1', txArgs, 30)
  }
}
