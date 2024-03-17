import { createRegistry, Any } from '@bufbuild/protobuf'
import { pb, utils } from 'saacs'

let i = new pb.common.generic.AuthorizeOperationRequest()

let v = new pb.sample.Book({
    isbn: "1234567890",
})
let reg = createRegistry(pb.sample.Book, pb.sample.SimpleItem)
console.log(reg)
let a = Any.pack(v)
console.log(a)


let item = new pb.auth.objects.Item({
    value: a,
})
let arg = new pb.common.generic.CreateRequest({
    item: item,
})
let j = a.toJsonString({ typeRegistry: reg })


console.log(j)
