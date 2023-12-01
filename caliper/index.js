'use strict'
const l = require("hlf_tools");
const { default: createJITI } = require("jiti");

const jiti = require("jiti")(__filename);
const hlf = import("hlf_tools");


let x = new  l.gen.auth.v1.auth_pb.Attribute()




console.log(x.toJson())
