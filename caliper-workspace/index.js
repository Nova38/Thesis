import { FieldMask } from "@bufbuild/protobuf"
import { Operation } from "./lib/gen/auth/v1/auth_pb.js"

const op = new  Operation({
    paths: new FieldMask({
        paths: ["a", "b"]
    }), 
})

console.log(op)
console.log(op.toJson())
