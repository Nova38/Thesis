export const MessageKeySchema = {
  "auth.ReferenceKey" : {
    "itemType": "auth.ReferenceKey",
    "itemKind": 4
},
  "auth.Collection" : {
    "itemType": "auth.Collection",
    "itemKind": 2,
    "properties": "collectionId"
},
  "auth.Role" : {
    "itemType": "auth.Role",
    "itemKind": 2,
    "properties": "roleId"
},
  "auth.Attribute" : {
    "itemType": "auth.Attribute",
    "itemKind": 2,
    "properties": "mspId,oid,roleId"
},
  "auth.UserDirectMembership" : {
    "itemType": "auth.UserDirectMembership",
    "itemKind": 2,
    "properties": "mspId,userId"
},
  "auth.UserEmbeddedRoles" : {
    "itemType": "auth.UserEmbeddedRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
},
  "auth.UserCollectionRoles" : {
    "itemType": "auth.UserCollectionRoles",
    "itemKind": 2,
    "properties": "mspId,userId"
},
  "auth.Suggestion" : {
    "itemType": "auth.Suggestion",
    "itemKind": 3,
    "properties": "suggestionId"
},
  "auth.HiddenTxList" : {
    "itemType": "auth.HiddenTxList",
    "itemKind": 3,
    "properties": ""
},
};
