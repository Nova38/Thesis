import { pb } from '@saacs/saacs-pb'

const x = pb.Book.fromJson({})

console.log('Hello, Caliper!')
console.log(x.toJsonString())
