export const MessageKeySchema = {
    "auth.ReferenceKey": {
        itemType: "",
        itemKind: 4,
    },
    "auth.Collection": {
        itemType: "",
        itemKind: 2,
        properties: "collectionId",
    },
    "auth.Role": {
        itemType: "",
        itemKind: 2,
        properties: "roleId",
    },
    "auth.Attribute": {
        itemType: "",
        itemKind: 2,
        properties: "mspId,oid,roleId",
    },
    "auth.UserDirectMembership": {
        itemType: "",
        itemKind: 2,
        properties: "mspId,userId",
    },
    "auth.UserEmbeddedRoles": {
        itemType: "",
        itemKind: 2,
        properties: "mspId,userId",
    },
    "auth.UserCollectionRoles": {
        itemType: "",
        itemKind: 2,
        properties: "mspId,userId",
    },
    "auth.Suggestion": {
        itemType: "",
        itemKind: 3,
        properties: "suggestionId",
    },
    "auth.HiddenTxList": {
        itemType: "",
        itemKind: 3,
        properties: "",
    },
};
