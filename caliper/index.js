'use strict'
// check ts

const hlf = require("saacs-es");
const { default: createJITI } = require("jiti");
const { Any, createRegistry } = require('@bufbuild/protobuf');

// const jiti = require("jiti")(__filename);
// const hlf = await import("hlftools");



// const item =new hlf.pb.common.generic.AuthorizeOperationRequest({})

// item.toJsonString({typeRegistry: hlf.utils.GlobalRegistry})




const obj = new hlf.pb.sample.SimpleItem({
    name: "item-1-1-1",
})


console.log(hlf.utils.GlobalRegistry.findMessage("auth.AuthorizeOperationRequest"))

const arg = new hlf.pb.common.generic.CreateRequest()
arg.item = Any.pack(obj)


// console.log(arg.toJsonString({typeRegistry: createRegistry(...hlf.pb.sample.allMessages)}))
console.log(`${hlf.utils.factory.modCollectionId(3,3)}`)
