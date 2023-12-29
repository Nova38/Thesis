import {
    Any,
    EnumType,
    IMessageTypeRegistry,
    MessageType,
    ServiceType,
    createRegistry,
    createRegistryFromDescriptors,
} from "@bufbuild/protobuf";
// import * as items_pb from "../gen/chaincode/sample/v0/items_pb";
// import * as auth_pb from "@auth/auth_pb.js"

import * as auth_pb from "../gen/auth/v1/index.js";
import * as common_pb from "../gen/chaincode/common/index.js";
import * as ccbio from "../gen/biochain/v1/index.js";
import { auth, sample } from "../gen/index.js";

export type Registry = IMessageTypeRegistry;

export const GlobalRegistry: IMessageTypeRegistry = createRegistry(
    ...auth_pb.auth.allMessages,
    ...auth_pb.objects.allMessages,
    ...common_pb.generic.allMessages,
    ...common_pb.reference.allMessages,
    ...ccbio.allMessages,
    ...sample.allMessages,
);
// auth_pb.allTypes.forEach((item) => {
//     GlobalTypes.push(item)
// }
// GlobalTypes.push(common_pb.helper.allTypes)

// // reg.push(common.reference.allTypes)

// GlobalTypes.push(ccbio.service.allTypes)
// GlobalTypes.push(ccbio.state.allTypes)

// export const Registry = createRegistry(...GlobalTypes)

export function build() {
    // let item = new authpb.Collection({
    // })
}

// export const Registry = createRegistry(

//     auth_pb.Polices,
//     auth_pb.Attribute,
//     auth_pb.Collection,
//     auth_pb.HiddenTx,
//     auth_pb.HiddenTxList,
//     auth_pb.History,
//     auth_pb.Item,
//     auth_pb.ItemKey,
//     auth_pb.KeySchema,
//     auth_pb.UserDirectMembership,
//     auth_pb.Operation,
//     auth_pb.PathPolicy,
//     auth_pb.Reference,
//     auth_pb.Role,
//     auth_pb.StateActivity,
//     auth_pb.Suggestion,
// )
