"use strict";
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);
var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};

// index.ts
var es_exports = {};
__export(es_exports, {
  gen: () => gen,
  operations: () => operations,
  utils: () => utils
});
module.exports = __toCommonJS(es_exports);

// gen/auth/v1/auth_key.ts
var auth_key_exports = {};
__export(auth_key_exports, {
  AttributeKey: () => AttributeKey,
  RoleKey: () => RoleKey,
  UserCollectionRolesKey: () => UserCollectionRolesKey,
  UserMembershipKey: () => UserMembershipKey
});
function RoleKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.roleId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.roleId);
  return attr;
}
function AttributeKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.mspId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.mspId);
  if (!(item == null ? void 0 : item.oid)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.oid);
  return attr;
}
function UserMembershipKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.mspId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.mspId);
  if (!(item == null ? void 0 : item.userId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.userId);
  return attr;
}
function UserCollectionRolesKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.mspId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.mspId);
  if (!(item == null ? void 0 : item.userId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.userId);
  return attr;
}

// gen/auth/v1/auth_pb.ts
var auth_pb_exports = {};
__export(auth_pb_exports, {
  Action: () => Action,
  Attribute: () => Attribute,
  AuthType: () => AuthType,
  Collection: () => Collection,
  FullItem: () => FullItem,
  HiddenTx: () => HiddenTx,
  HiddenTxList: () => HiddenTxList,
  History: () => History,
  HistoryEntry: () => HistoryEntry,
  Item: () => Item,
  ItemKey: () => ItemKey,
  ItemKind: () => ItemKind,
  KeySchema: () => KeySchema,
  Operation: () => Operation,
  PathPolicy: () => PathPolicy,
  Polices: () => Polices,
  Reference: () => Reference,
  ReferenceKey: () => ReferenceKey,
  Role: () => Role,
  StateActivity: () => StateActivity,
  Suggestion: () => Suggestion,
  TransactionType: () => TransactionType,
  TxError: () => TxError,
  User: () => User,
  UserCollectionRoles: () => UserCollectionRoles,
  UserMembership: () => UserMembership
});
var import_protobuf = require("@bufbuild/protobuf");
var TransactionType = /* @__PURE__ */ ((TransactionType2) => {
  TransactionType2[TransactionType2["UNSPECIFIED"] = 0] = "UNSPECIFIED";
  TransactionType2[TransactionType2["INVOKE"] = 1] = "INVOKE";
  TransactionType2[TransactionType2["QUERY"] = 2] = "QUERY";
  return TransactionType2;
})(TransactionType || {});
import_protobuf.proto3.util.setEnumType(TransactionType, "auth.TransactionType", [
  { no: 0, name: "TRANSACTION_TYPE_UNSPECIFIED" },
  { no: 1, name: "TRANSACTION_TYPE_INVOKE" },
  { no: 2, name: "TRANSACTION_TYPE_QUERY" }
]);
var AuthType = /* @__PURE__ */ ((AuthType2) => {
  AuthType2[AuthType2["UNSPECIFIED"] = 0] = "UNSPECIFIED";
  AuthType2[AuthType2["NONE"] = 1] = "NONE";
  AuthType2[AuthType2["ROLE"] = 2] = "ROLE";
  AuthType2[AuthType2["IDENTITY"] = 3] = "IDENTITY";
  return AuthType2;
})(AuthType || {});
import_protobuf.proto3.util.setEnumType(AuthType, "auth.AuthType", [
  { no: 0, name: "AUTH_TYPE_UNSPECIFIED" },
  { no: 1, name: "AUTH_TYPE_NONE" },
  { no: 2, name: "AUTH_TYPE_ROLE" },
  { no: 3, name: "AUTH_TYPE_IDENTITY" }
]);
var ItemKind = /* @__PURE__ */ ((ItemKind2) => {
  ItemKind2[ItemKind2["UNSPECIFIED"] = 0] = "UNSPECIFIED";
  ItemKind2[ItemKind2["GLOBAL_ITEM"] = 1] = "GLOBAL_ITEM";
  ItemKind2[ItemKind2["PRIMARY_ITEM"] = 2] = "PRIMARY_ITEM";
  ItemKind2[ItemKind2["SUB_ITEM"] = 3] = "SUB_ITEM";
  ItemKind2[ItemKind2["REFERENCE"] = 4] = "REFERENCE";
  return ItemKind2;
})(ItemKind || {});
import_protobuf.proto3.util.setEnumType(ItemKind, "auth.ItemKind", [
  { no: 0, name: "ITEM_KIND_UNSPECIFIED" },
  { no: 1, name: "ITEM_KIND_GLOBAL_ITEM" },
  { no: 2, name: "ITEM_KIND_PRIMARY_ITEM" },
  { no: 3, name: "ITEM_KIND_SUB_ITEM" },
  { no: 4, name: "ITEM_KIND_REFERENCE" }
]);
var Action = /* @__PURE__ */ ((Action2) => {
  Action2[Action2["UNSPECIFIED"] = 0] = "UNSPECIFIED";
  Action2[Action2["UTILITY"] = 1] = "UTILITY";
  Action2[Action2["VIEW"] = 10] = "VIEW";
  Action2[Action2["CREATE"] = 11] = "CREATE";
  Action2[Action2["UPDATE"] = 12] = "UPDATE";
  Action2[Action2["DELETE"] = 13] = "DELETE";
  Action2[Action2["SUGGEST_VIEW"] = 14] = "SUGGEST_VIEW";
  Action2[Action2["SUGGEST_CREATE"] = 15] = "SUGGEST_CREATE";
  Action2[Action2["SUGGEST_DELETE"] = 16] = "SUGGEST_DELETE";
  Action2[Action2["SUGGEST_APPROVE"] = 17] = "SUGGEST_APPROVE";
  Action2[Action2["VIEW_HISTORY"] = 18] = "VIEW_HISTORY";
  Action2[Action2["VIEW_HIDDEN_TXS"] = 19] = "VIEW_HIDDEN_TXS";
  Action2[Action2["HIDE_TX"] = 20] = "HIDE_TX";
  Action2[Action2["REFERENCE_CREATE"] = 21] = "REFERENCE_CREATE";
  Action2[Action2["REFERENCE_DELETE"] = 22] = "REFERENCE_DELETE";
  Action2[Action2["REFERENCE_VIEW"] = 23] = "REFERENCE_VIEW";
  return Action2;
})(Action || {});
import_protobuf.proto3.util.setEnumType(Action, "auth.Action", [
  { no: 0, name: "ACTION_UNSPECIFIED" },
  { no: 1, name: "ACTION_UTILITY" },
  { no: 10, name: "ACTION_VIEW" },
  { no: 11, name: "ACTION_CREATE" },
  { no: 12, name: "ACTION_UPDATE" },
  { no: 13, name: "ACTION_DELETE" },
  { no: 14, name: "ACTION_SUGGEST_VIEW" },
  { no: 15, name: "ACTION_SUGGEST_CREATE" },
  { no: 16, name: "ACTION_SUGGEST_DELETE" },
  { no: 17, name: "ACTION_SUGGEST_APPROVE" },
  { no: 18, name: "ACTION_VIEW_HISTORY" },
  { no: 19, name: "ACTION_VIEW_HIDDEN_TXS" },
  { no: 20, name: "ACTION_HIDE_TX" },
  { no: 21, name: "ACTION_REFERENCE_CREATE" },
  { no: 22, name: "ACTION_REFERENCE_DELETE" },
  { no: 23, name: "ACTION_REFERENCE_VIEW" }
]);
var TxError = /* @__PURE__ */ ((TxError2) => {
  TxError2[TxError2["UNSPECIFIED"] = 0] = "UNSPECIFIED";
  TxError2[TxError2["REQUEST_INVALID"] = 1] = "REQUEST_INVALID";
  TxError2[TxError2["RUNTIME"] = 2] = "RUNTIME";
  TxError2[TxError2["RUNTIME_BAD_OPS"] = 3] = "RUNTIME_BAD_OPS";
  TxError2[TxError2["KEY_NOT_FOUND"] = 4] = "KEY_NOT_FOUND";
  TxError2[TxError2["KEY_ALREADY_EXISTS"] = 5] = "KEY_ALREADY_EXISTS";
  TxError2[TxError2["COLLECTION_INVALID_ID"] = 11] = "COLLECTION_INVALID_ID";
  TxError2[TxError2["COLLECTION_UNREGISTERED"] = 12] = "COLLECTION_UNREGISTERED";
  TxError2[TxError2["COLLECTION_ALREADY_REGISTERED"] = 13] = "COLLECTION_ALREADY_REGISTERED";
  TxError2[TxError2["COLLECTION_INVALID"] = 14] = "COLLECTION_INVALID";
  TxError2[TxError2["COLLECTION_INVALID_ITEM_TYPE"] = 15] = "COLLECTION_INVALID_ITEM_TYPE";
  TxError2[TxError2["COLLECTION_INVALID_ROLE_ID"] = 16] = "COLLECTION_INVALID_ROLE_ID";
  TxError2[TxError2["USER_INVALID_ID"] = 20] = "USER_INVALID_ID";
  TxError2[TxError2["USER_UNREGISTERED"] = 21] = "USER_UNREGISTERED";
  TxError2[TxError2["USER_ALREADY_REGISTERED"] = 22] = "USER_ALREADY_REGISTERED";
  TxError2[TxError2["USER_INVALID"] = 23] = "USER_INVALID";
  TxError2[TxError2["USER_NO_ROLE"] = 24] = "USER_NO_ROLE";
  TxError2[TxError2["USER_PERMISSION_DENIED"] = 26] = "USER_PERMISSION_DENIED";
  TxError2[TxError2["ITEM_INVALID_ID"] = 31] = "ITEM_INVALID_ID";
  TxError2[TxError2["ITEM_UNREGISTERED"] = 32] = "ITEM_UNREGISTERED";
  TxError2[TxError2["ITEM_ALREADY_REGISTERED"] = 33] = "ITEM_ALREADY_REGISTERED";
  TxError2[TxError2["ITEM_INVALID"] = 34] = "ITEM_INVALID";
  TxError2[TxError2["INVALID_ITEM_FIELD_PATH"] = 35] = "INVALID_ITEM_FIELD_PATH";
  TxError2[TxError2["INVALID_ITEM_FIELD_VALUE"] = 36] = "INVALID_ITEM_FIELD_VALUE";
  return TxError2;
})(TxError || {});
import_protobuf.proto3.util.setEnumType(TxError, "auth.TxError", [
  { no: 0, name: "UNSPECIFIED" },
  { no: 1, name: "REQUEST_INVALID" },
  { no: 2, name: "RUNTIME" },
  { no: 3, name: "RUNTIME_BAD_OPS" },
  { no: 4, name: "KEY_NOT_FOUND" },
  { no: 5, name: "KEY_ALREADY_EXISTS" },
  { no: 11, name: "COLLECTION_INVALID_ID" },
  { no: 12, name: "COLLECTION_UNREGISTERED" },
  { no: 13, name: "COLLECTION_ALREADY_REGISTERED" },
  { no: 14, name: "COLLECTION_INVALID" },
  { no: 15, name: "COLLECTION_INVALID_ITEM_TYPE" },
  { no: 16, name: "COLLECTION_INVALID_ROLE_ID" },
  { no: 20, name: "USER_INVALID_ID" },
  { no: 21, name: "USER_UNREGISTERED" },
  { no: 22, name: "USER_ALREADY_REGISTERED" },
  { no: 23, name: "USER_INVALID" },
  { no: 24, name: "USER_NO_ROLE" },
  { no: 26, name: "USER_PERMISSION_DENIED" },
  { no: 31, name: "ITEM_INVALID_ID" },
  { no: 32, name: "ITEM_UNREGISTERED" },
  { no: 33, name: "ITEM_ALREADY_REGISTERED" },
  { no: 34, name: "ITEM_INVALID" },
  { no: 35, name: "INVALID_ITEM_FIELD_PATH" },
  { no: 36, name: "INVALID_ITEM_FIELD_VALUE" }
]);
var _KeySchema = class _KeySchema extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The item type of the key
     *
     * @generated from field: string item_type = 1;
     */
    this.itemType = "";
    /**
     * The kind of item that the key is for
     *
     * @generated from field: auth.ItemKind item_kind = 2;
     */
    this.itemKind = 0 /* UNSPECIFIED */;
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _KeySchema().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _KeySchema().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _KeySchema().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_KeySchema, a, b);
  }
};
_KeySchema.runtime = import_protobuf.proto3;
_KeySchema.typeName = "auth.KeySchema";
_KeySchema.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "item_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "item_kind", kind: "enum", T: import_protobuf.proto3.getEnumType(ItemKind) },
  { no: 3, name: "keys", kind: "message", T: import_protobuf.FieldMask }
]);
var KeySchema = _KeySchema;
var _StateActivity = class _StateActivity extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The transaction id that caused the change
     *
     * @generated from field: string tx_id = 1;
     */
    this.txId = "";
    /**
     * The msp of the user that caused the change
     *
     * @generated from field: string msp_id = 2;
     */
    this.mspId = "";
    /**
     * The id of the user that caused the change
     *
     * @generated from field: string user_id = 3;
     */
    this.userId = "";
    /**
     * A note about the change
     *
     * @generated from field: string note = 5;
     */
    this.note = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _StateActivity().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _StateActivity().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _StateActivity().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_StateActivity, a, b);
  }
};
_StateActivity.runtime = import_protobuf.proto3;
_StateActivity.typeName = "auth.StateActivity";
_StateActivity.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "tx_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "msp_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "user_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "timestamp", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 5,
    name: "note",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var StateActivity = _StateActivity;
var _Operation = class _Operation extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: auth.Action action = 1;
     */
    this.action = 0 /* UNSPECIFIED */;
    /**
     * @generated from field: string collection_id = 2;
     */
    this.collectionId = "";
    /**
     * @generated from field: string item_type = 3;
     */
    this.itemType = "";
    /**
     * @generated from field: string secondary_item_type = 4;
     */
    this.secondaryItemType = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Operation().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Operation().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Operation().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Operation, a, b);
  }
};
_Operation.runtime = import_protobuf.proto3;
_Operation.typeName = "auth.Operation";
_Operation.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "action", kind: "enum", T: import_protobuf.proto3.getEnumType(Action) },
  {
    no: 2,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "item_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "secondary_item_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "paths", kind: "message", T: import_protobuf.FieldMask }
]);
var Operation = _Operation;
var _PathPolicy = class _PathPolicy extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The path is a sub path of a field mask
     *
     * @generated from field: string path = 1;
     */
    this.path = "";
    /**
     * @generated from field: string full_path = 2;
     */
    this.fullPath = "";
    /**
     * @generated from field: bool allow_sub_paths = 3;
     */
    this.allowSubPaths = false;
    /**
     * The key is a valid sub path in the type of state item
     *
     * @generated from field: map<string, auth.PathPolicy> sub_paths = 4;
     */
    this.subPaths = {};
    /**
     * If the policy is not set than use a parent policy unless nested policy is set
     *
     * @generated from field: repeated auth.Action actions = 5;
     */
    this.actions = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _PathPolicy().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _PathPolicy().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _PathPolicy().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_PathPolicy, a, b);
  }
};
_PathPolicy.runtime = import_protobuf.proto3;
_PathPolicy.typeName = "auth.PathPolicy";
_PathPolicy.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "path",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "full_path",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "allow_sub_paths",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 4, name: "sub_paths", kind: "map", K: 9, V: { kind: "message", T: _PathPolicy } },
  { no: 5, name: "actions", kind: "enum", T: import_protobuf.proto3.getEnumType(Action), repeated: true }
]);
var PathPolicy = _PathPolicy;
var _Polices = class _Polices extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * key is the item type
     *
     * @generated from field: map<string, auth.PathPolicy> item_policies = 1;
     */
    this.itemPolicies = {};
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Polices().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Polices().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Polices().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Polices, a, b);
  }
};
_Polices.runtime = import_protobuf.proto3;
_Polices.typeName = "auth.Polices";
_Polices.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "item_policies", kind: "map", K: 9, V: { kind: "message", T: PathPolicy } }
]);
var Polices = _Polices;
var _Item = class _Item extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Item().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Item().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Item().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Item, a, b);
  }
};
_Item.runtime = import_protobuf.proto3;
_Item.typeName = "auth.Item";
_Item.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "key", kind: "message", T: ItemKey },
  { no: 2, name: "value", kind: "message", T: import_protobuf.Any }
]);
var Item = _Item;
var _FullItem = class _FullItem extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated auth.Suggestion suggestions = 3;
     */
    this.suggestions = [];
    /**
     * @generated from field: repeated auth.Reference references = 4;
     */
    this.references = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _FullItem().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _FullItem().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _FullItem().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_FullItem, a, b);
  }
};
_FullItem.runtime = import_protobuf.proto3;
_FullItem.typeName = "auth.FullItem";
_FullItem.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "history", kind: "message", T: History },
  { no: 3, name: "suggestions", kind: "message", T: Suggestion, repeated: true },
  { no: 4, name: "references", kind: "message", T: Reference, repeated: true }
]);
var FullItem = _FullItem;
var _HistoryEntry = class _HistoryEntry extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The transaction id that caused the change
     *
     * @generated from field: string tx_id = 1;
     */
    this.txId = "";
    /**
     * Whether the item was deleted
     *
     * @generated from field: bool is_delete = 2;
     */
    this.isDelete = false;
    /**
     * Whether the transaction was hidden
     *
     * @generated from field: bool is_hidden = 3;
     */
    this.isHidden = false;
    /**
     * A note about the change
     *
     * @generated from field: string note = 5;
     */
    this.note = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HistoryEntry().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HistoryEntry().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HistoryEntry().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_HistoryEntry, a, b);
  }
};
_HistoryEntry.runtime = import_protobuf.proto3;
_HistoryEntry.typeName = "auth.HistoryEntry";
_HistoryEntry.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "tx_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "is_delete",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  {
    no: 3,
    name: "is_hidden",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 4, name: "timestamp", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 5,
    name: "note",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "value", kind: "message", T: import_protobuf.Any }
]);
var HistoryEntry = _HistoryEntry;
var _History = class _History extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated auth.HistoryEntry entries = 1;
     */
    this.entries = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _History().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _History().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _History().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_History, a, b);
  }
};
_History.runtime = import_protobuf.proto3;
_History.typeName = "auth.History";
_History.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "entries", kind: "message", T: HistoryEntry, repeated: true },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTxList }
]);
var History = _History;
var _ItemKey = class _ItemKey extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string item_type = 2;
     */
    this.itemType = "";
    /**
     * @generated from field: repeated string item_id_parts = 3;
     */
    this.itemIdParts = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ItemKey().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ItemKey().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ItemKey().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ItemKey, a, b);
  }
};
_ItemKey.runtime = import_protobuf.proto3;
_ItemKey.typeName = "auth.ItemKey";
_ItemKey.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "item_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "item_id_parts", kind: "scalar", T: 9, repeated: true }
]);
var ItemKey = _ItemKey;
var _ReferenceKey = class _ReferenceKey extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceKey().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceKey().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceKey().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_ReferenceKey, a, b);
  }
};
_ReferenceKey.runtime = import_protobuf.proto3;
_ReferenceKey.typeName = "auth.ReferenceKey";
_ReferenceKey.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "key1", kind: "message", T: ItemKey },
  { no: 2, name: "key2", kind: "message", T: ItemKey }
]);
var ReferenceKey = _ReferenceKey;
var _Reference = class _Reference extends import_protobuf.Message {
  constructor(data) {
    super();
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Reference().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Reference().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Reference().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Reference, a, b);
  }
};
_Reference.runtime = import_protobuf.proto3;
_Reference.typeName = "auth.Reference";
_Reference.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "reference", kind: "message", T: ReferenceKey },
  { no: 2, name: "item1", kind: "message", T: Item },
  { no: 3, name: "item2", kind: "message", T: Item }
]);
var Reference = _Reference;
var _Collection = class _Collection extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The key for the ledger
     *
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: string description = 3;
     */
    this.description = "";
    /**
     * @generated from field: auth.AuthType auth_type = 4;
     */
    this.authType = 0 /* UNSPECIFIED */;
    /**
     *  [(buf.validate.field).repeated.items = {
     *   string: {prefix: "type.googleapis.com/"}
     * }];
     *
     * @generated from field: repeated string item_types = 5;
     */
    this.itemTypes = [];
    /**
     * @generated from field: repeated string reference_types = 6;
     */
    this.referenceTypes = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Collection().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Collection().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Collection().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Collection, a, b);
  }
};
_Collection.runtime = import_protobuf.proto3;
_Collection.typeName = "auth.Collection";
_Collection.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "auth_type", kind: "enum", T: import_protobuf.proto3.getEnumType(AuthType) },
  { no: 5, name: "item_types", kind: "scalar", T: 9, repeated: true },
  { no: 6, name: "reference_types", kind: "scalar", T: 9, repeated: true },
  { no: 7, name: "default", kind: "message", T: Polices }
]);
var Collection = _Collection;
var _User = class _User extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string msp_id = 2;
     */
    this.mspId = "";
    /**
     * @generated from field: string user_id = 3;
     */
    this.userId = "";
    /**
     * @generated from field: string name = 4;
     */
    this.name = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _User().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _User().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _User().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_User, a, b);
  }
};
_User.runtime = import_protobuf.proto3;
_User.typeName = "auth.User";
_User.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "msp_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "user_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var User = _User;
var _Suggestion = class _Suggestion extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string suggestion_id = 2;
     */
    this.suggestionId = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Suggestion().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Suggestion().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Suggestion().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Suggestion, a, b);
  }
};
_Suggestion.runtime = import_protobuf.proto3;
_Suggestion.typeName = "auth.Suggestion";
_Suggestion.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "primary_key", kind: "message", T: ItemKey },
  {
    no: 2,
    name: "suggestion_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "paths", kind: "message", T: import_protobuf.FieldMask },
  { no: 6, name: "value", kind: "message", T: import_protobuf.Any }
]);
var Suggestion = _Suggestion;
var _HiddenTx = class _HiddenTx extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string tx_id = 1;
     */
    this.txId = "";
    /**
     * @generated from field: string msp_id = 2;
     */
    this.mspId = "";
    /**
     * @generated from field: string user_id = 3;
     */
    this.userId = "";
    /**
     * @generated from field: string note = 5;
     */
    this.note = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HiddenTx().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HiddenTx().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HiddenTx().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_HiddenTx, a, b);
  }
};
_HiddenTx.runtime = import_protobuf.proto3;
_HiddenTx.typeName = "auth.HiddenTx";
_HiddenTx.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "tx_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "msp_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "user_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "timestamp", kind: "message", T: import_protobuf.Timestamp },
  {
    no: 5,
    name: "note",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var HiddenTx = _HiddenTx;
var _HiddenTxList = class _HiddenTxList extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The list of hidden txs by tx_id
     *
     * @generated from field: repeated auth.HiddenTx txs = 4;
     */
    this.txs = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HiddenTxList().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HiddenTxList().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HiddenTxList().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_HiddenTxList, a, b);
  }
};
_HiddenTxList.runtime = import_protobuf.proto3;
_HiddenTxList.typeName = "auth.HiddenTxList";
_HiddenTxList.fields = import_protobuf.proto3.util.newFieldList(() => [
  { no: 1, name: "primary_key", kind: "message", T: ItemKey },
  { no: 4, name: "txs", kind: "message", T: HiddenTx, repeated: true }
]);
var HiddenTxList = _HiddenTxList;
var _Role = class _Role extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string role_id = 2;
     */
    this.roleId = "";
    /**
     * @generated from field: string description = 5;
     */
    this.description = "";
    /**
     * @generated from field: repeated string parent_role_ids = 6;
     */
    this.parentRoleIds = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Role().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Role().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Role().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Role, a, b);
  }
};
_Role.runtime = import_protobuf.proto3;
_Role.typeName = "auth.Role";
_Role.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "role_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "polices", kind: "message", T: Polices },
  {
    no: 5,
    name: "description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "parent_role_ids", kind: "scalar", T: 9, repeated: true }
]);
var Role = _Role;
var _Attribute = class _Attribute extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * The msp of the organization that this attribute applies to
     *
     * @generated from field: string msp_id = 2;
     */
    this.mspId = "";
    /**
     * The oid of the attribute
     *
     * @generated from field: string oid = 3;
     */
    this.oid = "";
    /**
     * The value of the attribute required to be satisfied by the user to have the
     * role
     *
     * @generated from field: string value = 4;
     */
    this.value = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Attribute().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Attribute().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Attribute().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_Attribute, a, b);
  }
};
_Attribute.runtime = import_protobuf.proto3;
_Attribute.typeName = "auth.Attribute";
_Attribute.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "msp_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "oid",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "value",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "polices", kind: "message", T: Polices }
]);
var Attribute = _Attribute;
var _UserMembership = class _UserMembership extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The collection that the user is a member of
     *
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * The msp of the organization that the user's certificate is from
     *
     * @generated from field: string msp_id = 2;
     */
    this.mspId = "";
    /**
     * The id of the user from the certificate
     *
     * @generated from field: string user_id = 3;
     */
    this.userId = "";
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UserMembership().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UserMembership().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UserMembership().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UserMembership, a, b);
  }
};
_UserMembership.runtime = import_protobuf.proto3;
_UserMembership.typeName = "auth.UserMembership";
_UserMembership.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "msp_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "user_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "polices", kind: "message", T: Polices }
]);
var UserMembership = _UserMembership;
var _UserCollectionRoles = class _UserCollectionRoles extends import_protobuf.Message {
  constructor(data) {
    super();
    /**
     * The collection that the user is a member of
     *
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * The msp of the organization that the user's certificate is from
     *
     * @generated from field: string msp_id = 2;
     */
    this.mspId = "";
    /**
     * The id of the user from the certificate
     *
     * @generated from field: string user_id = 3;
     */
    this.userId = "";
    /**
     * The roles that the user has in the collection
     *
     * @generated from field: repeated string role_ids = 4;
     */
    this.roleIds = [];
    import_protobuf.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UserCollectionRoles().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UserCollectionRoles().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UserCollectionRoles().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf.proto3.util.equals(_UserCollectionRoles, a, b);
  }
};
_UserCollectionRoles.runtime = import_protobuf.proto3;
_UserCollectionRoles.typeName = "auth.UserCollectionRoles";
_UserCollectionRoles.fields = import_protobuf.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "msp_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "user_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "role_ids", kind: "scalar", T: 9, repeated: true }
]);
var UserCollectionRoles = _UserCollectionRoles;

// gen/auth/v1/auth_reg.ts
var auth_reg_exports = {};
__export(auth_reg_exports, {
  allTypes: () => allTypes
});
var allTypes = [
  KeySchema,
  StateActivity,
  Operation,
  PathPolicy,
  Polices,
  Item,
  FullItem,
  HistoryEntry,
  History,
  ItemKey,
  ReferenceKey,
  Reference,
  Collection,
  User,
  Suggestion,
  HiddenTx,
  HiddenTxList,
  Role,
  Attribute,
  UserMembership,
  UserCollectionRoles
];

// gen/chaincode/auth/common/generic_gateway.ts
var generic_gateway_exports = {};
__export(generic_gateway_exports, {
  GenericServiceClient: () => GenericServiceClient
});

// gen/chaincode/auth/common/generic_pb.ts
var generic_pb_exports = {};
__export(generic_pb_exports, {
  AuthorizeOperationRequest: () => AuthorizeOperationRequest,
  AuthorizeOperationResponse: () => AuthorizeOperationResponse,
  BootstrapRequest: () => BootstrapRequest,
  BootstrapResponse: () => BootstrapResponse,
  CreateCollectionRequest: () => CreateCollectionRequest,
  CreateCollectionResponse: () => CreateCollectionResponse,
  CreateRequest: () => CreateRequest,
  CreateResponse: () => CreateResponse,
  CreateUserResponse: () => CreateUserResponse,
  DeleteRequest: () => DeleteRequest,
  DeleteResponse: () => DeleteResponse,
  GetCurrentFullUserResponse: () => GetCurrentFullUserResponse,
  GetCurrentUserResponse: () => GetCurrentUserResponse,
  GetRequest: () => GetRequest,
  GetResponse: () => GetResponse,
  HiddenTxRequest: () => HiddenTxRequest,
  HiddenTxResponse: () => HiddenTxResponse,
  HideTxRequest: () => HideTxRequest,
  HideTxResponse: () => HideTxResponse,
  HistoryRequest: () => HistoryRequest,
  HistoryResponse: () => HistoryResponse,
  ListByAttrsRequest: () => ListByAttrsRequest,
  ListByAttrsResponse: () => ListByAttrsResponse,
  ListByCollectionRequest: () => ListByCollectionRequest,
  ListByCollectionResponse: () => ListByCollectionResponse,
  ListRequest: () => ListRequest,
  ListResponse: () => ListResponse,
  ReferenceByCollectionRequest: () => ReferenceByCollectionRequest,
  ReferenceByCollectionResponse: () => ReferenceByCollectionResponse,
  ReferenceByItemRequest: () => ReferenceByItemRequest,
  ReferenceByItemResponse: () => ReferenceByItemResponse,
  ReferenceByPartialKeyRequest: () => ReferenceByPartialKeyRequest,
  ReferenceByPartialKeyResponse: () => ReferenceByPartialKeyResponse,
  ReferenceCreateRequest: () => ReferenceCreateRequest,
  ReferenceCreateResponse: () => ReferenceCreateResponse,
  ReferenceDeleteRequest: () => ReferenceDeleteRequest,
  ReferenceDeleteResponse: () => ReferenceDeleteResponse,
  ReferenceRequest: () => ReferenceRequest,
  ReferenceResponse: () => ReferenceResponse,
  SuggestionApproveRequest: () => SuggestionApproveRequest,
  SuggestionApproveResponse: () => SuggestionApproveResponse,
  SuggestionByPartialKeyRequest: () => SuggestionByPartialKeyRequest,
  SuggestionByPartialKeyResponse: () => SuggestionByPartialKeyResponse,
  SuggestionCreateRequest: () => SuggestionCreateRequest,
  SuggestionCreateResponse: () => SuggestionCreateResponse,
  SuggestionDeleteRequest: () => SuggestionDeleteRequest,
  SuggestionDeleteResponse: () => SuggestionDeleteResponse,
  SuggestionListByCollectionRequest: () => SuggestionListByCollectionRequest,
  SuggestionListByCollectionResponse: () => SuggestionListByCollectionResponse,
  SuggestionListByItemRequest: () => SuggestionListByItemRequest,
  SuggestionListByItemResponse: () => SuggestionListByItemResponse,
  SuggestionListRequest: () => SuggestionListRequest,
  SuggestionListResponse: () => SuggestionListResponse,
  SuggestionRequest: () => SuggestionRequest,
  SuggestionResponse: () => SuggestionResponse,
  UnHideTxRequest: () => UnHideTxRequest,
  UnHideTxResponse: () => UnHideTxResponse,
  UpdateRequest: () => UpdateRequest,
  UpdateResponse: () => UpdateResponse
});
var import_protobuf2 = require("@bufbuild/protobuf");
var _GetCurrentUserResponse = class _GetCurrentUserResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool registerd = 2;
     */
    this.registerd = false;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetCurrentUserResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetCurrentUserResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetCurrentUserResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_GetCurrentUserResponse, a, b);
  }
};
_GetCurrentUserResponse.runtime = import_protobuf2.proto3;
_GetCurrentUserResponse.typeName = "auth.common.GetCurrentUserResponse";
_GetCurrentUserResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "user", kind: "message", T: User },
  {
    no: 2,
    name: "registerd",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var GetCurrentUserResponse = _GetCurrentUserResponse;
var _GetCurrentFullUserResponse = class _GetCurrentFullUserResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool registerd = 2;
     */
    this.registerd = false;
    /**
     * @generated from field: repeated auth.UserCollectionRoles user_collection_roles = 3;
     */
    this.userCollectionRoles = [];
    /**
     * @generated from field: repeated auth.UserMembership user_memberships = 4;
     */
    this.userMemberships = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetCurrentFullUserResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetCurrentFullUserResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetCurrentFullUserResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_GetCurrentFullUserResponse, a, b);
  }
};
_GetCurrentFullUserResponse.runtime = import_protobuf2.proto3;
_GetCurrentFullUserResponse.typeName = "auth.common.GetCurrentFullUserResponse";
_GetCurrentFullUserResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "user", kind: "message", T: User },
  {
    no: 2,
    name: "registerd",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 3, name: "user_collection_roles", kind: "message", T: UserCollectionRoles, repeated: true },
  { no: 4, name: "user_memberships", kind: "message", T: UserMembership, repeated: true }
]);
var GetCurrentFullUserResponse = _GetCurrentFullUserResponse;
var _AuthorizeOperationRequest = class _AuthorizeOperationRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AuthorizeOperationRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AuthorizeOperationRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AuthorizeOperationRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_AuthorizeOperationRequest, a, b);
  }
};
_AuthorizeOperationRequest.runtime = import_protobuf2.proto3;
_AuthorizeOperationRequest.typeName = "auth.common.AuthorizeOperationRequest";
_AuthorizeOperationRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "operation", kind: "message", T: Operation }
]);
var AuthorizeOperationRequest = _AuthorizeOperationRequest;
var _AuthorizeOperationResponse = class _AuthorizeOperationResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool authorized = 1;
     */
    this.authorized = false;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _AuthorizeOperationResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _AuthorizeOperationResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _AuthorizeOperationResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_AuthorizeOperationResponse, a, b);
  }
};
_AuthorizeOperationResponse.runtime = import_protobuf2.proto3;
_AuthorizeOperationResponse.typeName = "auth.common.AuthorizeOperationResponse";
_AuthorizeOperationResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "authorized",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var AuthorizeOperationResponse = _AuthorizeOperationResponse;
var _BootstrapRequest = class _BootstrapRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated string default_types = 1;
     */
    this.defaultTypes = [];
    /**
     * @generated from field: bool add_default_setup = 2;
     */
    this.addDefaultSetup = false;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _BootstrapRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _BootstrapRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _BootstrapRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_BootstrapRequest, a, b);
  }
};
_BootstrapRequest.runtime = import_protobuf2.proto3;
_BootstrapRequest.typeName = "auth.common.BootstrapRequest";
_BootstrapRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "default_types", kind: "scalar", T: 9, repeated: true },
  {
    no: 2,
    name: "add_default_setup",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var BootstrapRequest = _BootstrapRequest;
var _BootstrapResponse = class _BootstrapResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool success = 1;
     */
    this.success = false;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _BootstrapResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _BootstrapResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _BootstrapResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_BootstrapResponse, a, b);
  }
};
_BootstrapResponse.runtime = import_protobuf2.proto3;
_BootstrapResponse.typeName = "auth.common.BootstrapResponse";
_BootstrapResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "success",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var BootstrapResponse = _BootstrapResponse;
var _CreateCollectionRequest = class _CreateCollectionRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateCollectionRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateCollectionRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateCollectionRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_CreateCollectionRequest, a, b);
  }
};
_CreateCollectionRequest.runtime = import_protobuf2.proto3;
_CreateCollectionRequest.typeName = "auth.common.CreateCollectionRequest";
_CreateCollectionRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "collection", kind: "message", T: Collection }
]);
var CreateCollectionRequest = _CreateCollectionRequest;
var _CreateCollectionResponse = class _CreateCollectionResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateCollectionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateCollectionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateCollectionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_CreateCollectionResponse, a, b);
  }
};
_CreateCollectionResponse.runtime = import_protobuf2.proto3;
_CreateCollectionResponse.typeName = "auth.common.CreateCollectionResponse";
_CreateCollectionResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "collection", kind: "message", T: Collection }
]);
var CreateCollectionResponse = _CreateCollectionResponse;
var _CreateUserResponse = class _CreateUserResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateUserResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateUserResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateUserResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_CreateUserResponse, a, b);
  }
};
_CreateUserResponse.runtime = import_protobuf2.proto3;
_CreateUserResponse.typeName = "auth.common.CreateUserResponse";
_CreateUserResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "user", kind: "message", T: User }
]);
var CreateUserResponse = _CreateUserResponse;
var _GetRequest = class _GetRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_GetRequest, a, b);
  }
};
_GetRequest.runtime = import_protobuf2.proto3;
_GetRequest.typeName = "auth.common.GetRequest";
_GetRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 3, name: "item", kind: "message", T: Item }
]);
var GetRequest = _GetRequest;
var _GetResponse = class _GetResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _GetResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _GetResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _GetResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_GetResponse, a, b);
  }
};
_GetResponse.runtime = import_protobuf2.proto3;
_GetResponse.typeName = "auth.common.GetResponse";
_GetResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item }
]);
var GetResponse = _GetResponse;
var _ListRequest = class _ListRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ListRequest, a, b);
  }
};
_ListRequest.runtime = import_protobuf2.proto3;
_ListRequest.typeName = "auth.common.ListRequest";
_ListRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 3, name: "item", kind: "message", T: Item }
]);
var ListRequest = _ListRequest;
var _ListResponse = class _ListResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.Item items = 2;
     */
    this.items = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ListResponse, a, b);
  }
};
_ListResponse.runtime = import_protobuf2.proto3;
_ListResponse.typeName = "auth.common.ListResponse";
_ListResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "items", kind: "message", T: Item, repeated: true }
]);
var ListResponse = _ListResponse;
var _ListByCollectionRequest = class _ListByCollectionRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListByCollectionRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListByCollectionRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListByCollectionRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ListByCollectionRequest, a, b);
  }
};
_ListByCollectionRequest.runtime = import_protobuf2.proto3;
_ListByCollectionRequest.typeName = "auth.common.ListByCollectionRequest";
_ListByCollectionRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 3, name: "item", kind: "message", T: Item }
]);
var ListByCollectionRequest = _ListByCollectionRequest;
var _ListByCollectionResponse = class _ListByCollectionResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.Item items = 2;
     */
    this.items = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListByCollectionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListByCollectionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListByCollectionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ListByCollectionResponse, a, b);
  }
};
_ListByCollectionResponse.runtime = import_protobuf2.proto3;
_ListByCollectionResponse.typeName = "auth.common.ListByCollectionResponse";
_ListByCollectionResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "items", kind: "message", T: Item, repeated: true }
]);
var ListByCollectionResponse = _ListByCollectionResponse;
var _ListByAttrsRequest = class _ListByAttrsRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    /**
     * @generated from field: int32 num_attrs = 4;
     */
    this.numAttrs = 0;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListByAttrsRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListByAttrsRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListByAttrsRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ListByAttrsRequest, a, b);
  }
};
_ListByAttrsRequest.runtime = import_protobuf2.proto3;
_ListByAttrsRequest.typeName = "auth.common.ListByAttrsRequest";
_ListByAttrsRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 3, name: "item", kind: "message", T: Item },
  {
    no: 4,
    name: "num_attrs",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);
var ListByAttrsRequest = _ListByAttrsRequest;
var _ListByAttrsResponse = class _ListByAttrsResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.Item items = 2;
     */
    this.items = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ListByAttrsResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ListByAttrsResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ListByAttrsResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ListByAttrsResponse, a, b);
  }
};
_ListByAttrsResponse.runtime = import_protobuf2.proto3;
_ListByAttrsResponse.typeName = "auth.common.ListByAttrsResponse";
_ListByAttrsResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "items", kind: "message", T: Item, repeated: true }
]);
var ListByAttrsResponse = _ListByAttrsResponse;
var _CreateRequest = class _CreateRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_CreateRequest, a, b);
  }
};
_CreateRequest.runtime = import_protobuf2.proto3;
_CreateRequest.typeName = "auth.common.CreateRequest";
_CreateRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var CreateRequest = _CreateRequest;
var _CreateResponse = class _CreateResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _CreateResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _CreateResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _CreateResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_CreateResponse, a, b);
  }
};
_CreateResponse.runtime = import_protobuf2.proto3;
_CreateResponse.typeName = "auth.common.CreateResponse";
_CreateResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var CreateResponse = _CreateResponse;
var _UpdateRequest = class _UpdateRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_UpdateRequest, a, b);
  }
};
_UpdateRequest.runtime = import_protobuf2.proto3;
_UpdateRequest.typeName = "auth.common.UpdateRequest";
_UpdateRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item },
  { no: 3, name: "update_mask", kind: "message", T: import_protobuf2.FieldMask }
]);
var UpdateRequest = _UpdateRequest;
var _UpdateResponse = class _UpdateResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UpdateResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UpdateResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UpdateResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_UpdateResponse, a, b);
  }
};
_UpdateResponse.runtime = import_protobuf2.proto3;
_UpdateResponse.typeName = "auth.common.UpdateResponse";
_UpdateResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var UpdateResponse = _UpdateResponse;
var _DeleteRequest = class _DeleteRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string reason = 4;
     */
    this.reason = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_DeleteRequest, a, b);
  }
};
_DeleteRequest.runtime = import_protobuf2.proto3;
_DeleteRequest.typeName = "auth.common.DeleteRequest";
_DeleteRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  {
    no: 4,
    name: "reason",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var DeleteRequest = _DeleteRequest;
var _DeleteResponse = class _DeleteResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _DeleteResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _DeleteResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _DeleteResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_DeleteResponse, a, b);
  }
};
_DeleteResponse.runtime = import_protobuf2.proto3;
_DeleteResponse.typeName = "auth.common.DeleteResponse";
_DeleteResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var DeleteResponse = _DeleteResponse;
var _HistoryRequest = class _HistoryRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HistoryRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HistoryRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HistoryRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_HistoryRequest, a, b);
  }
};
_HistoryRequest.runtime = import_protobuf2.proto3;
_HistoryRequest.typeName = "auth.common.HistoryRequest";
_HistoryRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item }
]);
var HistoryRequest = _HistoryRequest;
var _HistoryResponse = class _HistoryResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HistoryResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HistoryResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HistoryResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_HistoryResponse, a, b);
  }
};
_HistoryResponse.runtime = import_protobuf2.proto3;
_HistoryResponse.typeName = "auth.common.HistoryResponse";
_HistoryResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 2, name: "history", kind: "message", T: History }
]);
var HistoryResponse = _HistoryResponse;
var _HiddenTxRequest = class _HiddenTxRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HiddenTxRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HiddenTxRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HiddenTxRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_HiddenTxRequest, a, b);
  }
};
_HiddenTxRequest.runtime = import_protobuf2.proto3;
_HiddenTxRequest.typeName = "auth.common.HiddenTxRequest";
_HiddenTxRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item }
]);
var HiddenTxRequest = _HiddenTxRequest;
var _HiddenTxResponse = class _HiddenTxResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: repeated auth.HiddenTx hidden_txs = 2;
     */
    this.hiddenTxs = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HiddenTxResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HiddenTxResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HiddenTxResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_HiddenTxResponse, a, b);
  }
};
_HiddenTxResponse.runtime = import_protobuf2.proto3;
_HiddenTxResponse.typeName = "auth.common.HiddenTxResponse";
_HiddenTxResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTx, repeated: true }
]);
var HiddenTxResponse = _HiddenTxResponse;
var _HideTxRequest = class _HideTxRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HideTxRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HideTxRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HideTxRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_HideTxRequest, a, b);
  }
};
_HideTxRequest.runtime = import_protobuf2.proto3;
_HideTxRequest.typeName = "auth.common.HideTxRequest";
_HideTxRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "hidden_tx", kind: "message", T: HiddenTx }
]);
var HideTxRequest = _HideTxRequest;
var _HideTxResponse = class _HideTxResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _HideTxResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _HideTxResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _HideTxResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_HideTxResponse, a, b);
  }
};
_HideTxResponse.runtime = import_protobuf2.proto3;
_HideTxResponse.typeName = "auth.common.HideTxResponse";
_HideTxResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTxList }
]);
var HideTxResponse = _HideTxResponse;
var _UnHideTxRequest = class _UnHideTxRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string tx_id = 2;
     */
    this.txId = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UnHideTxRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UnHideTxRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UnHideTxRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_UnHideTxRequest, a, b);
  }
};
_UnHideTxRequest.runtime = import_protobuf2.proto3;
_UnHideTxRequest.typeName = "auth.common.UnHideTxRequest";
_UnHideTxRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  {
    no: 2,
    name: "tx_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var UnHideTxRequest = _UnHideTxRequest;
var _UnHideTxResponse = class _UnHideTxResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _UnHideTxResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _UnHideTxResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _UnHideTxResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_UnHideTxResponse, a, b);
  }
};
_UnHideTxResponse.runtime = import_protobuf2.proto3;
_UnHideTxResponse.typeName = "auth.common.UnHideTxResponse";
_UnHideTxResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTxList }
]);
var UnHideTxResponse = _UnHideTxResponse;
var _ReferenceRequest = class _ReferenceRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceRequest, a, b);
  }
};
_ReferenceRequest.runtime = import_protobuf2.proto3;
_ReferenceRequest.typeName = "auth.common.ReferenceRequest";
_ReferenceRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "reference", kind: "message", T: ReferenceKey }
]);
var ReferenceRequest = _ReferenceRequest;
var _ReferenceResponse = class _ReferenceResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool exists = 1;
     */
    this.exists = false;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceResponse, a, b);
  }
};
_ReferenceResponse.runtime = import_protobuf2.proto3;
_ReferenceResponse.typeName = "auth.common.ReferenceResponse";
_ReferenceResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "exists",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  },
  { no: 2, name: "reference", kind: "message", T: Reference }
]);
var ReferenceResponse = _ReferenceResponse;
var _ReferenceByCollectionRequest = class _ReferenceByCollectionRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    /**
     * @generated from field: string collection_id = 3;
     */
    this.collectionId = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceByCollectionRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceByCollectionRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceByCollectionRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceByCollectionRequest, a, b);
  }
};
_ReferenceByCollectionRequest.runtime = import_protobuf2.proto3;
_ReferenceByCollectionRequest.typeName = "auth.common.ReferenceByCollectionRequest";
_ReferenceByCollectionRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var ReferenceByCollectionRequest = _ReferenceByCollectionRequest;
var _ReferenceByCollectionResponse = class _ReferenceByCollectionResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.ReferenceKey references = 2;
     */
    this.references = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceByCollectionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceByCollectionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceByCollectionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceByCollectionResponse, a, b);
  }
};
_ReferenceByCollectionResponse.runtime = import_protobuf2.proto3;
_ReferenceByCollectionResponse.typeName = "auth.common.ReferenceByCollectionResponse";
_ReferenceByCollectionResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "references", kind: "message", T: ReferenceKey, repeated: true }
]);
var ReferenceByCollectionResponse = _ReferenceByCollectionResponse;
var _ReferenceByPartialKeyRequest = class _ReferenceByPartialKeyRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceByPartialKeyRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceByPartialKeyRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceByPartialKeyRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceByPartialKeyRequest, a, b);
  }
};
_ReferenceByPartialKeyRequest.runtime = import_protobuf2.proto3;
_ReferenceByPartialKeyRequest.typeName = "auth.common.ReferenceByPartialKeyRequest";
_ReferenceByPartialKeyRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  { no: 3, name: "reference", kind: "message", T: ReferenceKey }
]);
var ReferenceByPartialKeyRequest = _ReferenceByPartialKeyRequest;
var _ReferenceByPartialKeyResponse = class _ReferenceByPartialKeyResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.ReferenceKey references = 2;
     */
    this.references = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceByPartialKeyResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceByPartialKeyResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceByPartialKeyResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceByPartialKeyResponse, a, b);
  }
};
_ReferenceByPartialKeyResponse.runtime = import_protobuf2.proto3;
_ReferenceByPartialKeyResponse.typeName = "auth.common.ReferenceByPartialKeyResponse";
_ReferenceByPartialKeyResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "references", kind: "message", T: ReferenceKey, repeated: true }
]);
var ReferenceByPartialKeyResponse = _ReferenceByPartialKeyResponse;
var _ReferenceByItemRequest = class _ReferenceByItemRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    /**
     * @generated from field: string collection_id = 3;
     */
    this.collectionId = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceByItemRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceByItemRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceByItemRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceByItemRequest, a, b);
  }
};
_ReferenceByItemRequest.runtime = import_protobuf2.proto3;
_ReferenceByItemRequest.typeName = "auth.common.ReferenceByItemRequest";
_ReferenceByItemRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "item_key", kind: "message", T: ItemKey }
]);
var ReferenceByItemRequest = _ReferenceByItemRequest;
var _ReferenceByItemResponse = class _ReferenceByItemResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.ReferenceKey references = 2;
     */
    this.references = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceByItemResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceByItemResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceByItemResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceByItemResponse, a, b);
  }
};
_ReferenceByItemResponse.runtime = import_protobuf2.proto3;
_ReferenceByItemResponse.typeName = "auth.common.ReferenceByItemResponse";
_ReferenceByItemResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "references", kind: "message", T: ReferenceKey, repeated: true }
]);
var ReferenceByItemResponse = _ReferenceByItemResponse;
var _ReferenceCreateRequest = class _ReferenceCreateRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceCreateRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceCreateRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceCreateRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceCreateRequest, a, b);
  }
};
_ReferenceCreateRequest.runtime = import_protobuf2.proto3;
_ReferenceCreateRequest.typeName = "auth.common.ReferenceCreateRequest";
_ReferenceCreateRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceCreateRequest = _ReferenceCreateRequest;
var _ReferenceCreateResponse = class _ReferenceCreateResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceCreateResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceCreateResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceCreateResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceCreateResponse, a, b);
  }
};
_ReferenceCreateResponse.runtime = import_protobuf2.proto3;
_ReferenceCreateResponse.typeName = "auth.common.ReferenceCreateResponse";
_ReferenceCreateResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceCreateResponse = _ReferenceCreateResponse;
var _ReferenceDeleteRequest = class _ReferenceDeleteRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceDeleteRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceDeleteRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceDeleteRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceDeleteRequest, a, b);
  }
};
_ReferenceDeleteRequest.runtime = import_protobuf2.proto3;
_ReferenceDeleteRequest.typeName = "auth.common.ReferenceDeleteRequest";
_ReferenceDeleteRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceDeleteRequest = _ReferenceDeleteRequest;
var _ReferenceDeleteResponse = class _ReferenceDeleteResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _ReferenceDeleteResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _ReferenceDeleteResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _ReferenceDeleteResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_ReferenceDeleteResponse, a, b);
  }
};
_ReferenceDeleteResponse.runtime = import_protobuf2.proto3;
_ReferenceDeleteResponse.typeName = "auth.common.ReferenceDeleteResponse";
_ReferenceDeleteResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceDeleteResponse = _ReferenceDeleteResponse;
var _SuggestionRequest = class _SuggestionRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string suggestion_id = 2;
     */
    this.suggestionId = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionRequest, a, b);
  }
};
_SuggestionRequest.runtime = import_protobuf2.proto3;
_SuggestionRequest.typeName = "auth.common.SuggestionRequest";
_SuggestionRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item_key", kind: "message", T: ItemKey },
  {
    no: 2,
    name: "suggestion_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SuggestionRequest = _SuggestionRequest;
var _SuggestionResponse = class _SuggestionResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionResponse, a, b);
  }
};
_SuggestionResponse.runtime = import_protobuf2.proto3;
_SuggestionResponse.typeName = "auth.common.SuggestionResponse";
_SuggestionResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionResponse = _SuggestionResponse;
var _SuggestionListRequest = class _SuggestionListRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * auth.Item item = 3;
     *
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionListRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionListRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionListRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionListRequest, a, b);
  }
};
_SuggestionListRequest.runtime = import_protobuf2.proto3;
_SuggestionListRequest.typeName = "auth.common.SuggestionListRequest";
_SuggestionListRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  }
]);
var SuggestionListRequest = _SuggestionListRequest;
var _SuggestionListResponse = class _SuggestionListResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.Suggestion suggestions = 2;
     */
    this.suggestions = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionListResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionListResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionListResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionListResponse, a, b);
  }
};
_SuggestionListResponse.runtime = import_protobuf2.proto3;
_SuggestionListResponse.typeName = "auth.common.SuggestionListResponse";
_SuggestionListResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "suggestions", kind: "message", T: Suggestion, repeated: true }
]);
var SuggestionListResponse = _SuggestionListResponse;
var _SuggestionListByCollectionRequest = class _SuggestionListByCollectionRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    /**
     * @generated from field: string collection_id = 3;
     */
    this.collectionId = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionListByCollectionRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionListByCollectionRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionListByCollectionRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionListByCollectionRequest, a, b);
  }
};
_SuggestionListByCollectionRequest.runtime = import_protobuf2.proto3;
_SuggestionListByCollectionRequest.typeName = "auth.common.SuggestionListByCollectionRequest";
_SuggestionListByCollectionRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SuggestionListByCollectionRequest = _SuggestionListByCollectionRequest;
var _SuggestionListByCollectionResponse = class _SuggestionListByCollectionResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.Suggestion suggestions = 2;
     */
    this.suggestions = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionListByCollectionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionListByCollectionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionListByCollectionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionListByCollectionResponse, a, b);
  }
};
_SuggestionListByCollectionResponse.runtime = import_protobuf2.proto3;
_SuggestionListByCollectionResponse.typeName = "auth.common.SuggestionListByCollectionResponse";
_SuggestionListByCollectionResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "suggestions", kind: "message", T: Suggestion, repeated: true }
]);
var SuggestionListByCollectionResponse = _SuggestionListByCollectionResponse;
var _SuggestionListByItemRequest = class _SuggestionListByItemRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionListByItemRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionListByItemRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionListByItemRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionListByItemRequest, a, b);
  }
};
_SuggestionListByItemRequest.runtime = import_protobuf2.proto3;
_SuggestionListByItemRequest.typeName = "auth.common.SuggestionListByItemRequest";
_SuggestionListByItemRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item_key", kind: "message", T: ItemKey }
]);
var SuggestionListByItemRequest = _SuggestionListByItemRequest;
var _SuggestionListByItemResponse = class _SuggestionListByItemResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated auth.Suggestion suggestions = 1;
     */
    this.suggestions = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionListByItemResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionListByItemResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionListByItemResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionListByItemResponse, a, b);
  }
};
_SuggestionListByItemResponse.runtime = import_protobuf2.proto3;
_SuggestionListByItemResponse.typeName = "auth.common.SuggestionListByItemResponse";
_SuggestionListByItemResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "suggestions", kind: "message", T: Suggestion, repeated: true }
]);
var SuggestionListByItemResponse = _SuggestionListByItemResponse;
var _SuggestionByPartialKeyRequest = class _SuggestionByPartialKeyRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: uint32 limit = 2;
     */
    this.limit = 0;
    /**
     * @generated from field: int32 num_attrs = 3;
     */
    this.numAttrs = 0;
    /**
     * @generated from field: string suggestion_id = 5;
     */
    this.suggestionId = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionByPartialKeyRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionByPartialKeyRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionByPartialKeyRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionByPartialKeyRequest, a, b);
  }
};
_SuggestionByPartialKeyRequest.runtime = import_protobuf2.proto3;
_SuggestionByPartialKeyRequest.typeName = "auth.common.SuggestionByPartialKeyRequest";
_SuggestionByPartialKeyRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "limit",
    kind: "scalar",
    T: 13
    /* ScalarType.UINT32 */
  },
  {
    no: 3,
    name: "num_attrs",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 4, name: "item_key", kind: "message", T: ItemKey },
  {
    no: 5,
    name: "suggestion_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SuggestionByPartialKeyRequest = _SuggestionByPartialKeyRequest;
var _SuggestionByPartialKeyResponse = class _SuggestionByPartialKeyResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated auth.Suggestion suggestions = 2;
     */
    this.suggestions = [];
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionByPartialKeyResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionByPartialKeyResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionByPartialKeyResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionByPartialKeyResponse, a, b);
  }
};
_SuggestionByPartialKeyResponse.runtime = import_protobuf2.proto3;
_SuggestionByPartialKeyResponse.typeName = "auth.common.SuggestionByPartialKeyResponse";
_SuggestionByPartialKeyResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "suggestions", kind: "message", T: Suggestion, repeated: true }
]);
var SuggestionByPartialKeyResponse = _SuggestionByPartialKeyResponse;
var _SuggestionCreateRequest = class _SuggestionCreateRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionCreateRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionCreateRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionCreateRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionCreateRequest, a, b);
  }
};
_SuggestionCreateRequest.runtime = import_protobuf2.proto3;
_SuggestionCreateRequest.typeName = "auth.common.SuggestionCreateRequest";
_SuggestionCreateRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionCreateRequest = _SuggestionCreateRequest;
var _SuggestionCreateResponse = class _SuggestionCreateResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionCreateResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionCreateResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionCreateResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionCreateResponse, a, b);
  }
};
_SuggestionCreateResponse.runtime = import_protobuf2.proto3;
_SuggestionCreateResponse.typeName = "auth.common.SuggestionCreateResponse";
_SuggestionCreateResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionCreateResponse = _SuggestionCreateResponse;
var _SuggestionDeleteRequest = class _SuggestionDeleteRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string suggestion_id = 2;
     */
    this.suggestionId = "";
    /**
     * @generated from field: string reason = 3;
     */
    this.reason = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionDeleteRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionDeleteRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionDeleteRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionDeleteRequest, a, b);
  }
};
_SuggestionDeleteRequest.runtime = import_protobuf2.proto3;
_SuggestionDeleteRequest.typeName = "auth.common.SuggestionDeleteRequest";
_SuggestionDeleteRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item_key", kind: "message", T: ItemKey },
  {
    no: 2,
    name: "suggestion_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "reason",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SuggestionDeleteRequest = _SuggestionDeleteRequest;
var _SuggestionDeleteResponse = class _SuggestionDeleteResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionDeleteResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionDeleteResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionDeleteResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionDeleteResponse, a, b);
  }
};
_SuggestionDeleteResponse.runtime = import_protobuf2.proto3;
_SuggestionDeleteResponse.typeName = "auth.common.SuggestionDeleteResponse";
_SuggestionDeleteResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 4, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionDeleteResponse = _SuggestionDeleteResponse;
var _SuggestionApproveRequest = class _SuggestionApproveRequest extends import_protobuf2.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string suggestion_id = 2;
     */
    this.suggestionId = "";
    /**
     * @generated from field: string reason = 3;
     */
    this.reason = "";
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionApproveRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionApproveRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionApproveRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionApproveRequest, a, b);
  }
};
_SuggestionApproveRequest.runtime = import_protobuf2.proto3;
_SuggestionApproveRequest.typeName = "auth.common.SuggestionApproveRequest";
_SuggestionApproveRequest.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "item_key", kind: "message", T: ItemKey },
  {
    no: 2,
    name: "suggestion_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "reason",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SuggestionApproveRequest = _SuggestionApproveRequest;
var _SuggestionApproveResponse = class _SuggestionApproveResponse extends import_protobuf2.Message {
  constructor(data) {
    super();
    import_protobuf2.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SuggestionApproveResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SuggestionApproveResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SuggestionApproveResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf2.proto3.util.equals(_SuggestionApproveResponse, a, b);
  }
};
_SuggestionApproveResponse.runtime = import_protobuf2.proto3;
_SuggestionApproveResponse.typeName = "auth.common.SuggestionApproveResponse";
_SuggestionApproveResponse.fields = import_protobuf2.proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion },
  { no: 2, name: "item", kind: "message", T: Item }
]);
var SuggestionApproveResponse = _SuggestionApproveResponse;

// gen/chaincode/auth/common/generic_gateway.ts
var GenericServiceClient = class {
  constructor(contract, registry) {
    this.registry = "";
    this.contract = contract;
  }
  call(service, method, data) {
    return __async(this, null, function* () {
      const headers = new Headers([]);
      headers.set("content-type", contentType);
      const response = yield fetch(
        `${this.baseUrl}/${service}/${method}`,
        {
          method: "POST",
          headers,
          body: data.toJsonString()
        }
      );
      if (response.status === 200) {
        if (contentType === "application/json") {
          return yield response.json();
        }
        return new Uint8Array(yield response.arrayBuffer());
      }
      throw Error(`HTTP ${response.status} ${response.statusText}`);
    });
  }
  /**
   * ══════════════════════════════════ Helper ═════════════════════════════════════
   * ────────────────────────────────── Query ──────────────────────────────────────
   * rpc GetAllTypes(google.protobuf.Empty) returns (GetAllTypesResponse) {
   *   option (auth.transaction_type) = TRANSACTION_TYPE_QUERY;
   *   option (auth.operation) = {action: ACTION_UTILITY};
   * }
   *
   * @generated from rpc auth.common.GenericService.GetCurrentUser
   */
  getCurrentUser(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "GetCurrentUser",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return GetCurrentUserResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * ──────────────────────────────── Invoke ───────────────────────────────────────
   *
   * @generated from rpc auth.common.GenericService.Bootstrap
   */
  bootstrap(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Bootstrap",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return BootstrapResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.AuthorizeOperation
   */
  authorizeOperation(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "AuthorizeOperation",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return AuthorizeOperationResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.CreateUser
   */
  createUser(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "CreateUser",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return CreateUserResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.CreateCollection
   */
  createCollection(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "CreateCollection",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return CreateCollectionResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.Get
   */
  get(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Get",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return GetResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.List
   */
  list(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "List",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ListResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.ListByCollection
   */
  listByCollection(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "ListByCollection",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ListByCollectionResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.ListByAttrs
   */
  listByAttrs(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "ListByAttrs",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ListByAttrsResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.Create
   */
  create(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Create",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return CreateResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.Update
   */
  update(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Update",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return UpdateResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.Delete
   */
  delete(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Delete",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return DeleteResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.History
   */
  history(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "History",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return HistoryResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.HiddenTx
   */
  hiddenTx(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "HiddenTx",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return HiddenTxResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.HideTx
   */
  hideTx(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "HideTx",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return HideTxResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.UnHideTx
   */
  unHideTx(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "UnHideTx",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return UnHideTxResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.Reference
   */
  reference(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Reference",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ReferenceResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceByItem
   */
  referenceByItem(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "ReferenceByItem",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ReferenceByItemResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceByPartialKey
   */
  referenceByPartialKey(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "ReferenceByPartialKey",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ReferenceByPartialKeyResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceCreate
   */
  referenceCreate(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "ReferenceCreate",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ReferenceCreateResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceDelete
   */
  referenceDelete(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "ReferenceDelete",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return ReferenceDeleteResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.Suggestion
   */
  suggestion(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "Suggestion",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SuggestionResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionListByCollection
   */
  suggestionListByCollection(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "SuggestionListByCollection",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SuggestionListByCollectionResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionByPartialKey
   */
  suggestionByPartialKey(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "SuggestionByPartialKey",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SuggestionByPartialKeyResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * ──────────────────────────────── Invoke ───────────────────────────────────────
   *
   * @generated from rpc auth.common.GenericService.SuggestionCreate
   */
  suggestionCreate(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "SuggestionCreate",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SuggestionCreateResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionDelete
   */
  suggestionDelete(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "SuggestionDelete",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SuggestionDeleteResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionApprove
   */
  suggestionApprove(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "auth.common.GenericService",
        "SuggestionApprove",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SuggestionApproveResponse.fromJson(data);
        })
      );
    });
  }
};

// gen/chaincode/auth/common/generic_key.ts
var generic_key_exports = {};

// gen/chaincode/auth/common/generic_reg.ts
var generic_reg_exports = {};
__export(generic_reg_exports, {
  allTypes: () => allTypes2
});
var allTypes2 = [
  GetCurrentUserResponse,
  GetCurrentFullUserResponse,
  AuthorizeOperationRequest,
  AuthorizeOperationResponse,
  BootstrapRequest,
  BootstrapResponse,
  CreateCollectionRequest,
  CreateCollectionResponse,
  CreateUserResponse,
  GetRequest,
  GetResponse,
  ListRequest,
  ListResponse,
  ListByCollectionRequest,
  ListByCollectionResponse,
  ListByAttrsRequest,
  ListByAttrsResponse,
  CreateRequest,
  CreateResponse,
  UpdateRequest,
  UpdateResponse,
  DeleteRequest,
  DeleteResponse,
  HistoryRequest,
  HistoryResponse,
  HiddenTxRequest,
  HiddenTxResponse,
  HideTxRequest,
  HideTxResponse,
  UnHideTxRequest,
  UnHideTxResponse,
  ReferenceRequest,
  ReferenceResponse,
  ReferenceByCollectionRequest,
  ReferenceByCollectionResponse,
  ReferenceByPartialKeyRequest,
  ReferenceByPartialKeyResponse,
  ReferenceByItemRequest,
  ReferenceByItemResponse,
  ReferenceCreateRequest,
  ReferenceCreateResponse,
  ReferenceDeleteRequest,
  ReferenceDeleteResponse,
  SuggestionRequest,
  SuggestionResponse,
  SuggestionListRequest,
  SuggestionListResponse,
  SuggestionListByCollectionRequest,
  SuggestionListByCollectionResponse,
  SuggestionListByItemRequest,
  SuggestionListByItemResponse,
  SuggestionByPartialKeyRequest,
  SuggestionByPartialKeyResponse,
  SuggestionCreateRequest,
  SuggestionCreateResponse,
  SuggestionDeleteRequest,
  SuggestionDeleteResponse,
  SuggestionApproveRequest,
  SuggestionApproveResponse
];

// gen/chaincode/auth/common/helper_reg.ts
var helper_reg_exports = {};
__export(helper_reg_exports, {
  allTypes: () => allTypes3
});
var allTypes3 = [];

// gen/chaincode/auth/rbac/schema/v1/rbac_reg.ts
var rbac_reg_exports = {};
__export(rbac_reg_exports, {
  allTypes: () => allTypes4
});
var allTypes4 = [];

// gen/chaincode/ccbio/schema/v0/service_gateway.ts
var service_gateway_exports = {};
__export(service_gateway_exports, {
  SpecimenServiceClient: () => SpecimenServiceClient
});

// gen/chaincode/ccbio/schema/v0/service_pb.ts
var service_pb_exports = {};
__export(service_pb_exports, {
  SpecimenCreateRequest: () => SpecimenCreateRequest,
  SpecimenCreateResponse: () => SpecimenCreateResponse,
  SpecimenDeleteRequest: () => SpecimenDeleteRequest,
  SpecimenDeleteResponse: () => SpecimenDeleteResponse,
  SpecimenGetByCollectionRequest: () => SpecimenGetByCollectionRequest,
  SpecimenGetByCollectionResponse: () => SpecimenGetByCollectionResponse,
  SpecimenGetHistoryRequest: () => SpecimenGetHistoryRequest,
  SpecimenGetHistoryResponse: () => SpecimenGetHistoryResponse,
  SpecimenGetListRequest: () => SpecimenGetListRequest,
  SpecimenGetListResponse: () => SpecimenGetListResponse,
  SpecimenGetRequest: () => SpecimenGetRequest,
  SpecimenGetResponse: () => SpecimenGetResponse,
  SpecimenHideTxRequest: () => SpecimenHideTxRequest,
  SpecimenHideTxResponse: () => SpecimenHideTxResponse,
  SpecimenUnHideTxRequest: () => SpecimenUnHideTxRequest,
  SpecimenUnHideTxResponse: () => SpecimenUnHideTxResponse,
  SpecimenUpdateRequest: () => SpecimenUpdateRequest,
  SpecimenUpdateResponse: () => SpecimenUpdateResponse
});
var import_protobuf4 = require("@bufbuild/protobuf");

// gen/chaincode/ccbio/schema/v0/state_pb.ts
var state_pb_exports = {};
__export(state_pb_exports, {
  Specimen: () => Specimen,
  Specimen_Georeference: () => Specimen_Georeference,
  Specimen_Grant: () => Specimen_Grant,
  Specimen_Image: () => Specimen_Image,
  Specimen_Loan: () => Specimen_Loan,
  Specimen_Primary: () => Specimen_Primary,
  Specimen_Secondary: () => Specimen_Secondary,
  Specimen_Taxon: () => Specimen_Taxon
});
var import_protobuf3 = require("@bufbuild/protobuf");
var _Specimen = class _Specimen extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string specimen_id = 2;
     */
    this.specimenId = "";
    /**
     * @generated from field: map<string, ccbio.schema.v0.Specimen.Image> images = 7;
     */
    this.images = {};
    /**
     * @generated from field: map<string, ccbio.schema.v0.Specimen.Loan> loans = 10;
     */
    this.loans = {};
    /**
     * @generated from field: map<string, ccbio.schema.v0.Specimen.Grant> grants = 11;
     */
    this.grants = {};
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen, a, b);
  }
};
_Specimen.runtime = import_protobuf3.proto3;
_Specimen.typeName = "ccbio.schema.v0.Specimen";
_Specimen.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "specimen_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "primary", kind: "message", T: Specimen_Primary },
  { no: 4, name: "secondary", kind: "message", T: Specimen_Secondary },
  { no: 5, name: "taxon", kind: "message", T: Specimen_Taxon },
  { no: 6, name: "georeference", kind: "message", T: Specimen_Georeference },
  { no: 7, name: "images", kind: "map", K: 9, V: { kind: "message", T: Specimen_Image } },
  { no: 10, name: "loans", kind: "map", K: 9, V: { kind: "message", T: Specimen_Loan } },
  { no: 11, name: "grants", kind: "map", K: 9, V: { kind: "message", T: Specimen_Grant } }
]);
var Specimen = _Specimen;
var _Specimen_Primary = class _Specimen_Primary extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string catalog_number = 1;
     */
    this.catalogNumber = "";
    /**
     * @generated from field: string accession_number = 2;
     */
    this.accessionNumber = "";
    /**
     * @generated from field: string field_number = 3;
     */
    this.fieldNumber = "";
    /**
     * @generated from field: string tissue_number = 4;
     */
    this.tissueNumber = "";
    /**
     * @generated from field: string cataloger = 5;
     */
    this.cataloger = "";
    /**
     * @generated from field: string collector = 6;
     */
    this.collector = "";
    /**
     * @generated from field: string determiner = 7;
     */
    this.determiner = "";
    /**
     * @generated from field: string determined_reason = 11;
     */
    this.determinedReason = "";
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Primary().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Primary().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Primary().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Primary, a, b);
  }
};
_Specimen_Primary.runtime = import_protobuf3.proto3;
_Specimen_Primary.typeName = "ccbio.schema.v0.Specimen.Primary";
_Specimen_Primary.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "catalog_number",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "accession_number",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "field_number",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "tissue_number",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "cataloger",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "collector",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "determiner",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 8, name: "field_date", kind: "message", T: import_protobuf3.Timestamp },
  { no: 9, name: "catalog_date", kind: "message", T: import_protobuf3.Timestamp },
  { no: 10, name: "determined_date", kind: "message", T: import_protobuf3.Timestamp },
  {
    no: 11,
    name: "determined_reason",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Primary = _Specimen_Primary;
var _Specimen_Secondary = class _Specimen_Secondary extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string preparation = 3;
     */
    this.preparation = "";
    /**
     * @generated from field: string condition = 4;
     */
    this.condition = "";
    /**
     * @generated from field: string notes = 5;
     */
    this.notes = "";
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Secondary().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Secondary().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Secondary().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Secondary, a, b);
  }
};
_Specimen_Secondary.runtime = import_protobuf3.proto3;
_Specimen_Secondary.typeName = "ccbio.schema.v0.Specimen.Secondary";
_Specimen_Secondary.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 3,
    name: "preparation",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "condition",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "notes",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Secondary = _Specimen_Secondary;
var _Specimen_Taxon = class _Specimen_Taxon extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string kingdom = 1;
     */
    this.kingdom = "";
    /**
     * @generated from field: string phylum = 2;
     */
    this.phylum = "";
    /**
     * @generated from field: string class = 3;
     */
    this.class = "";
    /**
     * @generated from field: string order = 4;
     */
    this.order = "";
    /**
     * @generated from field: string family = 5;
     */
    this.family = "";
    /**
     * @generated from field: string genus = 6;
     */
    this.genus = "";
    /**
     * @generated from field: string species = 7;
     */
    this.species = "";
    /**
     * @generated from field: string subspecies = 8;
     */
    this.subspecies = "";
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Taxon().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Taxon().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Taxon().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Taxon, a, b);
  }
};
_Specimen_Taxon.runtime = import_protobuf3.proto3;
_Specimen_Taxon.typeName = "ccbio.schema.v0.Specimen.Taxon";
_Specimen_Taxon.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "kingdom",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "phylum",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "class",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "order",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "family",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "genus",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "species",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 8,
    name: "subspecies",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Taxon = _Specimen_Taxon;
var _Specimen_Georeference = class _Specimen_Georeference extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string country = 1;
     */
    this.country = "";
    /**
     * @generated from field: string state_province = 2;
     */
    this.stateProvince = "";
    /**
     * @generated from field: string county = 3;
     */
    this.county = "";
    /**
     * @generated from field: string locality = 4;
     */
    this.locality = "";
    /**
     * @generated from field: string latitude = 5;
     */
    this.latitude = "";
    /**
     * @generated from field: string longitude = 6;
     */
    this.longitude = "";
    /**
     * @generated from field: string habitat = 7;
     */
    this.habitat = "";
    /**
     * @generated from field: repeated string notes = 8;
     */
    this.notes = [];
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Georeference().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Georeference().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Georeference().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Georeference, a, b);
  }
};
_Specimen_Georeference.runtime = import_protobuf3.proto3;
_Specimen_Georeference.typeName = "ccbio.schema.v0.Specimen.Georeference";
_Specimen_Georeference.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "country",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "state_province",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "county",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "locality",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 5,
    name: "latitude",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "longitude",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "habitat",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 8, name: "notes", kind: "scalar", T: 9, repeated: true },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Georeference = _Specimen_Georeference;
var _Specimen_Image = class _Specimen_Image extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string url = 2;
     */
    this.url = "";
    /**
     * @generated from field: string notes = 3;
     */
    this.notes = "";
    /**
     * @generated from field: string hash = 4;
     */
    this.hash = "";
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Image().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Image().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Image().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Image, a, b);
  }
};
_Specimen_Image.runtime = import_protobuf3.proto3;
_Specimen_Image.typeName = "ccbio.schema.v0.Specimen.Image";
_Specimen_Image.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "notes",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "hash",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Image = _Specimen_Image;
var _Specimen_Loan = class _Specimen_Loan extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string description = 2;
     */
    this.description = "";
    /**
     * @generated from field: string loaned_by = 3;
     */
    this.loanedBy = "";
    /**
     * @generated from field: string loaned_to = 4;
     */
    this.loanedTo = "";
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Loan().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Loan().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Loan().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Loan, a, b);
  }
};
_Specimen_Loan.runtime = import_protobuf3.proto3;
_Specimen_Loan.typeName = "ccbio.schema.v0.Specimen.Loan";
_Specimen_Loan.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "loaned_by",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "loaned_to",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "loaned_date", kind: "message", T: import_protobuf3.Timestamp },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Loan = _Specimen_Loan;
var _Specimen_Grant = class _Specimen_Grant extends import_protobuf3.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string description = 2;
     */
    this.description = "";
    /**
     * @generated from field: string granted_by = 3;
     */
    this.grantedBy = "";
    /**
     * @generated from field: string granted_to = 4;
     */
    this.grantedTo = "";
    import_protobuf3.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Specimen_Grant().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Specimen_Grant().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Specimen_Grant().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf3.proto3.util.equals(_Specimen_Grant, a, b);
  }
};
_Specimen_Grant.runtime = import_protobuf3.proto3;
_Specimen_Grant.typeName = "ccbio.schema.v0.Specimen.Grant";
_Specimen_Grant.fields = import_protobuf3.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "granted_by",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "granted_to",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "granted_date", kind: "message", T: import_protobuf3.Timestamp },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Grant = _Specimen_Grant;

// gen/chaincode/ccbio/schema/v0/service_pb.ts
var _SpecimenGetRequest = class _SpecimenGetRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * Specimen.Id id = 1 [(buf.validate.field).required = true];
     *
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string specimen_id = 2;
     */
    this.specimenId = "";
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetRequest, a, b);
  }
};
_SpecimenGetRequest.runtime = import_protobuf4.proto3;
_SpecimenGetRequest.typeName = "ccbio.schema.v0.SpecimenGetRequest";
_SpecimenGetRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "specimen_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SpecimenGetRequest = _SpecimenGetRequest;
var _SpecimenGetResponse = class _SpecimenGetResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetResponse, a, b);
  }
};
_SpecimenGetResponse.runtime = import_protobuf4.proto3;
_SpecimenGetResponse.typeName = "ccbio.schema.v0.SpecimenGetResponse";
_SpecimenGetResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenGetResponse = _SpecimenGetResponse;
var _SpecimenGetListRequest = class _SpecimenGetListRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: int32 page_size = 2;
     */
    this.pageSize = 0;
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetListRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetListRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetListRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetListRequest, a, b);
  }
};
_SpecimenGetListRequest.runtime = import_protobuf4.proto3;
_SpecimenGetListRequest.typeName = "ccbio.schema.v0.SpecimenGetListRequest";
_SpecimenGetListRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "page_size",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);
var SpecimenGetListRequest = _SpecimenGetListRequest;
var _SpecimenGetListResponse = class _SpecimenGetListResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string bookmark = 1;
     */
    this.bookmark = "";
    /**
     * @generated from field: repeated ccbio.schema.v0.Specimen specimens = 2;
     */
    this.specimens = [];
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetListResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetListResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetListResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetListResponse, a, b);
  }
};
_SpecimenGetListResponse.runtime = import_protobuf4.proto3;
_SpecimenGetListResponse.typeName = "ccbio.schema.v0.SpecimenGetListResponse";
_SpecimenGetListResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "bookmark",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "specimens", kind: "message", T: Specimen, repeated: true }
]);
var SpecimenGetListResponse = _SpecimenGetListResponse;
var _SpecimenGetByCollectionRequest = class _SpecimenGetByCollectionRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetByCollectionRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetByCollectionRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetByCollectionRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetByCollectionRequest, a, b);
  }
};
_SpecimenGetByCollectionRequest.runtime = import_protobuf4.proto3;
_SpecimenGetByCollectionRequest.typeName = "ccbio.schema.v0.SpecimenGetByCollectionRequest";
_SpecimenGetByCollectionRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SpecimenGetByCollectionRequest = _SpecimenGetByCollectionRequest;
var _SpecimenGetByCollectionResponse = class _SpecimenGetByCollectionResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated ccbio.schema.v0.Specimen specimens = 1;
     */
    this.specimens = [];
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetByCollectionResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetByCollectionResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetByCollectionResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetByCollectionResponse, a, b);
  }
};
_SpecimenGetByCollectionResponse.runtime = import_protobuf4.proto3;
_SpecimenGetByCollectionResponse.typeName = "ccbio.schema.v0.SpecimenGetByCollectionResponse";
_SpecimenGetByCollectionResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimens", kind: "message", T: Specimen, repeated: true }
]);
var SpecimenGetByCollectionResponse = _SpecimenGetByCollectionResponse;
var _SpecimenGetHistoryRequest = class _SpecimenGetHistoryRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string specimen_id = 2;
     */
    this.specimenId = "";
    /**
     * @generated from field: bool include_hidden = 3;
     */
    this.includeHidden = false;
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetHistoryRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetHistoryRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetHistoryRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetHistoryRequest, a, b);
  }
};
_SpecimenGetHistoryRequest.runtime = import_protobuf4.proto3;
_SpecimenGetHistoryRequest.typeName = "ccbio.schema.v0.SpecimenGetHistoryRequest";
_SpecimenGetHistoryRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "specimen_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "include_hidden",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var SpecimenGetHistoryRequest = _SpecimenGetHistoryRequest;
var _SpecimenGetHistoryResponse = class _SpecimenGetHistoryResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenGetHistoryResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenGetHistoryResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenGetHistoryResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenGetHistoryResponse, a, b);
  }
};
_SpecimenGetHistoryResponse.runtime = import_protobuf4.proto3;
_SpecimenGetHistoryResponse.typeName = "ccbio.schema.v0.SpecimenGetHistoryResponse";
_SpecimenGetHistoryResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "history", kind: "message", T: StateActivity }
]);
var SpecimenGetHistoryResponse = _SpecimenGetHistoryResponse;
var _SpecimenCreateRequest = class _SpecimenCreateRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenCreateRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenCreateRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenCreateRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenCreateRequest, a, b);
  }
};
_SpecimenCreateRequest.runtime = import_protobuf4.proto3;
_SpecimenCreateRequest.typeName = "ccbio.schema.v0.SpecimenCreateRequest";
_SpecimenCreateRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenCreateRequest = _SpecimenCreateRequest;
var _SpecimenCreateResponse = class _SpecimenCreateResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenCreateResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenCreateResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenCreateResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenCreateResponse, a, b);
  }
};
_SpecimenCreateResponse.runtime = import_protobuf4.proto3;
_SpecimenCreateResponse.typeName = "ccbio.schema.v0.SpecimenCreateResponse";
_SpecimenCreateResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenCreateResponse = _SpecimenCreateResponse;
var _SpecimenUpdateRequest = class _SpecimenUpdateRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenUpdateRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenUpdateRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenUpdateRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenUpdateRequest, a, b);
  }
};
_SpecimenUpdateRequest.runtime = import_protobuf4.proto3;
_SpecimenUpdateRequest.typeName = "ccbio.schema.v0.SpecimenUpdateRequest";
_SpecimenUpdateRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen },
  { no: 2, name: "update_mask", kind: "message", T: import_protobuf4.FieldMask }
]);
var SpecimenUpdateRequest = _SpecimenUpdateRequest;
var _SpecimenUpdateResponse = class _SpecimenUpdateResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenUpdateResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenUpdateResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenUpdateResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenUpdateResponse, a, b);
  }
};
_SpecimenUpdateResponse.runtime = import_protobuf4.proto3;
_SpecimenUpdateResponse.typeName = "ccbio.schema.v0.SpecimenUpdateResponse";
_SpecimenUpdateResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen },
  { no: 2, name: "update_mask", kind: "message", T: import_protobuf4.FieldMask }
]);
var SpecimenUpdateResponse = _SpecimenUpdateResponse;
var _SpecimenDeleteRequest = class _SpecimenDeleteRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string specimen_id = 2;
     */
    this.specimenId = "";
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenDeleteRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenDeleteRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenDeleteRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenDeleteRequest, a, b);
  }
};
_SpecimenDeleteRequest.runtime = import_protobuf4.proto3;
_SpecimenDeleteRequest.typeName = "ccbio.schema.v0.SpecimenDeleteRequest";
_SpecimenDeleteRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "specimen_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SpecimenDeleteRequest = _SpecimenDeleteRequest;
var _SpecimenDeleteResponse = class _SpecimenDeleteResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenDeleteResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenDeleteResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenDeleteResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenDeleteResponse, a, b);
  }
};
_SpecimenDeleteResponse.runtime = import_protobuf4.proto3;
_SpecimenDeleteResponse.typeName = "ccbio.schema.v0.SpecimenDeleteResponse";
_SpecimenDeleteResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenDeleteResponse = _SpecimenDeleteResponse;
var _SpecimenHideTxRequest = class _SpecimenHideTxRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string specimen_id = 2;
     */
    this.specimenId = "";
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenHideTxRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenHideTxRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenHideTxRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenHideTxRequest, a, b);
  }
};
_SpecimenHideTxRequest.runtime = import_protobuf4.proto3;
_SpecimenHideTxRequest.typeName = "ccbio.schema.v0.SpecimenHideTxRequest";
_SpecimenHideTxRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "specimen_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "tx", kind: "message", T: StateActivity }
]);
var SpecimenHideTxRequest = _SpecimenHideTxRequest;
var _SpecimenHideTxResponse = class _SpecimenHideTxResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenHideTxResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenHideTxResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenHideTxResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenHideTxResponse, a, b);
  }
};
_SpecimenHideTxResponse.runtime = import_protobuf4.proto3;
_SpecimenHideTxResponse.typeName = "ccbio.schema.v0.SpecimenHideTxResponse";
_SpecimenHideTxResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenHideTxResponse = _SpecimenHideTxResponse;
var _SpecimenUnHideTxRequest = class _SpecimenUnHideTxRequest extends import_protobuf4.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string specimen_id = 2;
     */
    this.specimenId = "";
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenUnHideTxRequest().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenUnHideTxRequest().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenUnHideTxRequest().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenUnHideTxRequest, a, b);
  }
};
_SpecimenUnHideTxRequest.runtime = import_protobuf4.proto3;
_SpecimenUnHideTxRequest.typeName = "ccbio.schema.v0.SpecimenUnHideTxRequest";
_SpecimenUnHideTxRequest.fields = import_protobuf4.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "specimen_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "tx", kind: "message", T: StateActivity }
]);
var SpecimenUnHideTxRequest = _SpecimenUnHideTxRequest;
var _SpecimenUnHideTxResponse = class _SpecimenUnHideTxResponse extends import_protobuf4.Message {
  constructor(data) {
    super();
    import_protobuf4.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _SpecimenUnHideTxResponse().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _SpecimenUnHideTxResponse().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _SpecimenUnHideTxResponse().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf4.proto3.util.equals(_SpecimenUnHideTxResponse, a, b);
  }
};
_SpecimenUnHideTxResponse.runtime = import_protobuf4.proto3;
_SpecimenUnHideTxResponse.typeName = "ccbio.schema.v0.SpecimenUnHideTxResponse";
_SpecimenUnHideTxResponse.fields = import_protobuf4.proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenUnHideTxResponse = _SpecimenUnHideTxResponse;

// gen/chaincode/ccbio/schema/v0/service_gateway.ts
var SpecimenServiceClient = class {
  constructor(contract, registry) {
    this.registry = "";
    this.contract = contract;
  }
  call(service, method, data) {
    return __async(this, null, function* () {
      const headers = new Headers([]);
      headers.set("content-type", contentType);
      const response = yield fetch(
        `${this.baseUrl}/${service}/${method}`,
        {
          method: "POST",
          headers,
          body: data.toJsonString()
        }
      );
      if (response.status === 200) {
        if (contentType === "application/json") {
          return yield response.json();
        }
        return new Uint8Array(yield response.arrayBuffer());
      }
      throw Error(`HTTP ${response.status} ${response.statusText}`);
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGet
   */
  specimenGet(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenGet",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenGetResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGetList
   */
  specimenGetList(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenGetList",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenGetListResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGetByCollection
   */
  specimenGetByCollection(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenGetByCollection",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenGetByCollectionResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGetHistory
   */
  specimenGetHistory(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenGetHistory",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenGetHistoryResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenCreate
   */
  specimenCreate(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenCreate",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenCreateResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenUpdate
   */
  specimenUpdate(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenUpdate",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenUpdateResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenDelete
   */
  specimenDelete(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenDelete",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenDeleteResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenHideTx
   */
  specimenHideTx(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenHideTx",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenHideTxResponse.fromJson(data);
        })
      );
    });
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenUnHideTx
   */
  specimenUnHideTx(request) {
    return __async(this, null, function* () {
      const promise = this.request(
        "ccbio.schema.v0.SpecimenService",
        "SpecimenUnHideTx",
        request
      );
      return promise.then(
        (data) => __async(this, null, function* () {
          return SpecimenUnHideTxResponse.fromJson(data);
        })
      );
    });
  }
};

// gen/chaincode/ccbio/schema/v0/service_key.ts
var service_key_exports = {};

// gen/chaincode/ccbio/schema/v0/service_reg.ts
var service_reg_exports = {};
__export(service_reg_exports, {
  allTypes: () => allTypes5
});
var allTypes5 = [
  SpecimenGetRequest,
  SpecimenGetResponse,
  SpecimenGetListRequest,
  SpecimenGetListResponse,
  SpecimenGetByCollectionRequest,
  SpecimenGetByCollectionResponse,
  SpecimenGetHistoryRequest,
  SpecimenGetHistoryResponse,
  SpecimenCreateRequest,
  SpecimenCreateResponse,
  SpecimenUpdateRequest,
  SpecimenUpdateResponse,
  SpecimenDeleteRequest,
  SpecimenDeleteResponse,
  SpecimenHideTxRequest,
  SpecimenHideTxResponse,
  SpecimenUnHideTxRequest,
  SpecimenUnHideTxResponse
];

// gen/chaincode/ccbio/schema/v0/state_key.ts
var state_key_exports = {};
__export(state_key_exports, {
  SpecimenKey: () => SpecimenKey
});
function SpecimenKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.specimenId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.specimenId);
  return attr;
}

// gen/chaincode/ccbio/schema/v0/state_reg.ts
var state_reg_exports = {};
__export(state_reg_exports, {
  allTypes: () => allTypes6
});
var allTypes6 = [
  Specimen,
  Specimen_Primary,
  Specimen_Secondary,
  Specimen_Taxon,
  Specimen_Georeference,
  Specimen_Image,
  Specimen_Loan,
  Specimen_Grant
];

// gen/chaincode/sample/v0/items_key.ts
var items_key_exports = {};
__export(items_key_exports, {
  AuthorKey: () => AuthorKey,
  AwardsKey: () => AwardsKey,
  PersonKey: () => PersonKey
});
function AwardsKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.awardName)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.awardName);
  return attr;
}
function AuthorKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.authorId)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.authorId);
  return attr;
}
function PersonKey(item) {
  attr = [];
  if (!(item == null ? void 0 : item.name)) {
    return attr;
  }
  attr.push(item == null ? void 0 : item.name);
  return attr;
}

// gen/chaincode/sample/v0/items_pb.ts
var items_pb_exports = {};
__export(items_pb_exports, {
  Author: () => Author,
  Awards: () => Awards,
  Book: () => Book,
  Degree: () => Degree,
  Group: () => Group,
  Item: () => Item2,
  Person: () => Person
});
var import_protobuf5 = require("@bufbuild/protobuf");
var _Item2 = class _Item2 extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string id = 1;
     */
    this.id = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: int32 quantity = 3;
     */
    this.quantity = 0;
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Item2().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Item2().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Item2().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Item2, a, b);
  }
};
_Item2.runtime = import_protobuf5.proto3;
_Item2.typeName = "sample.Item";
_Item2.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "quantity",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);
var Item2 = _Item2;
var _Group = class _Group extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string name = 1;
     */
    this.name = "";
    /**
     * @generated from field: repeated sample.Item items = 2;
     */
    this.items = [];
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Group().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Group().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Group().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Group, a, b);
  }
};
_Group.runtime = import_protobuf5.proto3;
_Group.typeName = "sample.Group";
_Group.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "items", kind: "message", T: Item2, repeated: true }
]);
var Group = _Group;
var _Book = class _Book extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string isbn = 1;
     */
    this.isbn = "";
    /**
     * @generated from field: string book_title = 2;
     */
    this.bookTitle = "";
    /**
     * @generated from field: string author = 3;
     */
    this.author = "";
    /**
     * @generated from field: int32 year = 4;
     */
    this.year = 0;
    /**
     * @generated from field: string publisher = 5;
     */
    this.publisher = "";
    /**
     * @generated from field: string language = 6;
     */
    this.language = "";
    /**
     * @generated from field: string description = 7;
     */
    this.description = "";
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Book().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Book().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Book().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Book, a, b);
  }
};
_Book.runtime = import_protobuf5.proto3;
_Book.typeName = "sample.Book";
_Book.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "isbn",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "book_title",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "author",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 4,
    name: "year",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  {
    no: 5,
    name: "publisher",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 6,
    name: "language",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 7,
    name: "description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var Book = _Book;
var _Degree = class _Degree extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string degree_name = 1;
     */
    this.degreeName = "";
    /**
     * @generated from field: string institute = 2;
     */
    this.institute = "";
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Degree().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Degree().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Degree().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Degree, a, b);
  }
};
_Degree.runtime = import_protobuf5.proto3;
_Degree.typeName = "sample.Degree";
_Degree.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "degree_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "institute",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "degree_date", kind: "message", T: import_protobuf5.Timestamp }
]);
var Degree = _Degree;
var _Awards = class _Awards extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string award_name = 2;
     */
    this.awardName = "";
    /**
     * @generated from field: string award_description = 4;
     */
    this.awardDescription = "";
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Awards().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Awards().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Awards().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Awards, a, b);
  }
};
_Awards.runtime = import_protobuf5.proto3;
_Awards.typeName = "sample.Awards";
_Awards.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "award_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 3, name: "award_date", kind: "message", T: import_protobuf5.Timestamp },
  {
    no: 4,
    name: "award_description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var Awards = _Awards;
var _Author = class _Author extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string author_id = 2;
     */
    this.authorId = "";
    /**
     * @generated from field: string author_name = 3;
     */
    this.authorName = "";
    /**
     * @generated from field: repeated sample.Book books = 4;
     */
    this.books = [];
    /**
     * Key: degree_name
     *
     * @generated from field: map<string, sample.Degree> degrees = 5;
     */
    this.degrees = {};
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Author().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Author().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Author().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Author, a, b);
  }
};
_Author.runtime = import_protobuf5.proto3;
_Author.typeName = "sample.Author";
_Author.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "author_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "author_name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 4, name: "books", kind: "message", T: Book, repeated: true },
  { no: 5, name: "degrees", kind: "map", K: 9, V: { kind: "message", T: Degree } }
]);
var Author = _Author;
var _Person = class _Person extends import_protobuf5.Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    /**
     * @generated from field: string name = 2;
     */
    this.name = "";
    /**
     * @generated from field: int32 age = 3;
     */
    this.age = 0;
    /**
     * @generated from field: repeated string friends = 4;
     */
    this.friends = [];
    /**
     * @generated from field: repeated sample.Group groups = 5;
     */
    this.groups = [];
    import_protobuf5.proto3.util.initPartial(data, this);
  }
  static fromBinary(bytes, options) {
    return new _Person().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Person().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Person().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return import_protobuf5.proto3.util.equals(_Person, a, b);
  }
};
_Person.runtime = import_protobuf5.proto3;
_Person.typeName = "sample.Person";
_Person.fields = import_protobuf5.proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "name",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 3,
    name: "age",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  },
  { no: 4, name: "friends", kind: "scalar", T: 9, repeated: true },
  { no: 5, name: "groups", kind: "message", T: Group, repeated: true }
]);
var Person = _Person;

// gen/chaincode/sample/v0/items_reg.ts
var items_reg_exports = {};
__export(items_reg_exports, {
  allTypes: () => allTypes7
});
var allTypes7 = [
  Item2,
  Group,
  Book,
  Degree,
  Awards,
  Author,
  Person
];

// operations/well_known.ts
var well_known_exports = {};
var wellKnown = {
  "Create Collection": new Operation({
    action: 11 /* CREATE */,
    itemType: Collection.typeName,
    collectionId: "<>"
  })
};

// utils/factory.ts
var factory_exports = {};
__export(factory_exports, {
  createAuthor: () => createAuthor,
  createAuthorItem: () => createAuthorItem,
  unpackItem: () => unpackItem
});
var import_protobuf7 = require("@bufbuild/protobuf");

// utils/registry.ts
var registry_exports = {};
__export(registry_exports, {
  Registry: () => Registry
});
var import_protobuf6 = require("@bufbuild/protobuf");
var Registry = (0, import_protobuf6.createRegistry)(
  Author,
  Awards,
  Degree,
  Item2,
  Group,
  Person,
  Polices,
  Attribute,
  Collection,
  HiddenTx,
  HiddenTxList,
  History,
  Item,
  ItemKey,
  KeySchema,
  UserMembership,
  Operation,
  PathPolicy,
  Reference,
  Role,
  StateActivity,
  Suggestion
);

// utils/factory.ts
function createAuthor() {
  const author = new Author();
  author.authorName = "John Doe";
  return author;
}
function createAuthorItem() {
  const obj = new Item();
  const author = createAuthor();
  console.log("author", author);
  obj.value = import_protobuf7.Any.pack(author);
  console.log("author", author);
  return obj;
}
function unpackItem(item) {
  var _a;
  const author = new Author();
  return (_a = item.value) == null ? void 0 : _a.unpack(Registry);
}

// utils/fakers.ts
var fakers_exports = {};

// index.ts
var gen = {
  auth: {
    v1: {
      auth_key: auth_key_exports,
      auth_pb: auth_pb_exports,
      auth_reg: auth_reg_exports
    }
  },
  chaincode: {
    auth: {
      common: {
        generic_gateway: generic_gateway_exports,
        generic_key: generic_key_exports,
        generic_pb: generic_pb_exports,
        generic_reg: generic_reg_exports,
        helper_reg: helper_reg_exports
      },
      rbac: {
        schema: {
          v1: {
            rbac_reg: rbac_reg_exports
          }
        }
      }
    },
    ccbio: {
      schema: {
        v0: {
          service_gateway: service_gateway_exports,
          service_key: service_key_exports,
          service_pb: service_pb_exports,
          service_reg: service_reg_exports,
          state_key: state_key_exports,
          state_pb: state_pb_exports,
          state_reg: state_reg_exports
        }
      }
    },
    sample: {
      v0: {
        items_key: items_key_exports,
        items_pb: items_pb_exports,
        items_reg: items_reg_exports
      }
    }
  }
};
var operations = {
  well_known: well_known_exports
};
var utils = {
  factory: factory_exports,
  fakers: fakers_exports,
  registry: registry_exports
};
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  gen,
  operations,
  utils
});
