
import { Any, createRegistry, createRegistryFromDescriptors } from "@bufbuild/protobuf";
// import * as items_pb from "../gen/chaincode/sample/v0/items_pb";
import * as auth_pb from "../gen/auth/v1/auth_pb";

export const Registry = createRegistry(



    auth_pb.Polices,
    auth_pb.Attribute,
    auth_pb.Collection,
    auth_pb.HiddenTx,
    auth_pb.HiddenTxList,
    auth_pb.History,
    auth_pb.Item,
    auth_pb.ItemKey,
    auth_pb.KeySchema,
    auth_pb.UserMembership,
    auth_pb.Operation,
    auth_pb.PathPolicy,
    auth_pb.Reference,
    auth_pb.Role,
    auth_pb.StateActivity,
    auth_pb.Suggestion,
)
