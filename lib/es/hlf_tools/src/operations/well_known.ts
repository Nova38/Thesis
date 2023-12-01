
import * as authpb from "../gen/auth/v1/auth_pb";



const wellKnown : Record<string, authpb.Operation> = {
    "Create Collection": new authpb.Operation({
        action: authpb.Action.CREATE,
        itemType: authpb.Collection.typeName,
        collectionId: "<>",
    }),

}
