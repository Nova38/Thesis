# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth/v1/auth.proto](#auth_v1_auth-proto)
    - [Attribute](#auth-Attribute)
    - [Collection](#auth-Collection)
    - [FullItem](#auth-FullItem)
    - [HiddenTx](#auth-HiddenTx)
    - [HiddenTxList](#auth-HiddenTxList)
    - [History](#auth-History)
    - [HistoryEntry](#auth-HistoryEntry)
    - [Item](#auth-Item)
    - [ItemKey](#auth-ItemKey)
    - [KeySchema](#auth-KeySchema)
    - [Operation](#auth-Operation)
    - [PathPolicy](#auth-PathPolicy)
    - [PathPolicy.SubPathsEntry](#auth-PathPolicy-SubPathsEntry)
    - [Polices](#auth-Polices)
    - [Polices.ItemPoliciesEntry](#auth-Polices-ItemPoliciesEntry)
    - [Reference](#auth-Reference)
    - [ReferenceKey](#auth-ReferenceKey)
    - [Role](#auth-Role)
    - [StateActivity](#auth-StateActivity)
    - [Suggestion](#auth-Suggestion)
    - [User](#auth-User)
    - [UserCollectionRoles](#auth-UserCollectionRoles)
    - [UserMembership](#auth-UserMembership)

    - [Action](#auth-Action)
    - [AuthType](#auth-AuthType)
    - [ItemKind](#auth-ItemKind)
    - [TransactionType](#auth-TransactionType)
    - [TxError](#auth-TxError)

    - [File-level Extensions](#auth_v1_auth-proto-extensions)
    - [File-level Extensions](#auth_v1_auth-proto-extensions)
    - [File-level Extensions](#auth_v1_auth-proto-extensions)

- [chaincode/auth/common/generic.proto](#chaincode_auth_common_generic-proto)
    - [AuthorizeOperationRequest](#auth-common-AuthorizeOperationRequest)
    - [AuthorizeOperationResponse](#auth-common-AuthorizeOperationResponse)
    - [BootstrapRequest](#auth-common-BootstrapRequest)
    - [BootstrapResponse](#auth-common-BootstrapResponse)
    - [CreateCollectionRequest](#auth-common-CreateCollectionRequest)
    - [CreateCollectionResponse](#auth-common-CreateCollectionResponse)
    - [CreateRequest](#auth-common-CreateRequest)
    - [CreateResponse](#auth-common-CreateResponse)
    - [CreateUserResponse](#auth-common-CreateUserResponse)
    - [DeleteRequest](#auth-common-DeleteRequest)
    - [DeleteResponse](#auth-common-DeleteResponse)
    - [GetCurrentFullUserResponse](#auth-common-GetCurrentFullUserResponse)
    - [GetCurrentUserResponse](#auth-common-GetCurrentUserResponse)
    - [GetRequest](#auth-common-GetRequest)
    - [GetResponse](#auth-common-GetResponse)
    - [HiddenTxRequest](#auth-common-HiddenTxRequest)
    - [HiddenTxResponse](#auth-common-HiddenTxResponse)
    - [HideTxRequest](#auth-common-HideTxRequest)
    - [HideTxResponse](#auth-common-HideTxResponse)
    - [HistoryRequest](#auth-common-HistoryRequest)
    - [HistoryResponse](#auth-common-HistoryResponse)
    - [ListByAttrsRequest](#auth-common-ListByAttrsRequest)
    - [ListByAttrsResponse](#auth-common-ListByAttrsResponse)
    - [ListByCollectionRequest](#auth-common-ListByCollectionRequest)
    - [ListByCollectionResponse](#auth-common-ListByCollectionResponse)
    - [ListRequest](#auth-common-ListRequest)
    - [ListResponse](#auth-common-ListResponse)
    - [ReferenceByCollectionRequest](#auth-common-ReferenceByCollectionRequest)
    - [ReferenceByCollectionResponse](#auth-common-ReferenceByCollectionResponse)
    - [ReferenceByItemRequest](#auth-common-ReferenceByItemRequest)
    - [ReferenceByItemResponse](#auth-common-ReferenceByItemResponse)
    - [ReferenceByPartialKeyRequest](#auth-common-ReferenceByPartialKeyRequest)
    - [ReferenceByPartialKeyResponse](#auth-common-ReferenceByPartialKeyResponse)
    - [ReferenceCreateRequest](#auth-common-ReferenceCreateRequest)
    - [ReferenceCreateResponse](#auth-common-ReferenceCreateResponse)
    - [ReferenceDeleteRequest](#auth-common-ReferenceDeleteRequest)
    - [ReferenceDeleteResponse](#auth-common-ReferenceDeleteResponse)
    - [ReferenceRequest](#auth-common-ReferenceRequest)
    - [ReferenceResponse](#auth-common-ReferenceResponse)
    - [SuggestionApproveRequest](#auth-common-SuggestionApproveRequest)
    - [SuggestionApproveResponse](#auth-common-SuggestionApproveResponse)
    - [SuggestionByPartialKeyRequest](#auth-common-SuggestionByPartialKeyRequest)
    - [SuggestionByPartialKeyResponse](#auth-common-SuggestionByPartialKeyResponse)
    - [SuggestionCreateRequest](#auth-common-SuggestionCreateRequest)
    - [SuggestionCreateResponse](#auth-common-SuggestionCreateResponse)
    - [SuggestionDeleteRequest](#auth-common-SuggestionDeleteRequest)
    - [SuggestionDeleteResponse](#auth-common-SuggestionDeleteResponse)
    - [SuggestionListByCollectionRequest](#auth-common-SuggestionListByCollectionRequest)
    - [SuggestionListByCollectionResponse](#auth-common-SuggestionListByCollectionResponse)
    - [SuggestionListByItemRequest](#auth-common-SuggestionListByItemRequest)
    - [SuggestionListByItemResponse](#auth-common-SuggestionListByItemResponse)
    - [SuggestionListRequest](#auth-common-SuggestionListRequest)
    - [SuggestionListResponse](#auth-common-SuggestionListResponse)
    - [SuggestionRequest](#auth-common-SuggestionRequest)
    - [SuggestionResponse](#auth-common-SuggestionResponse)
    - [UnHideTxRequest](#auth-common-UnHideTxRequest)
    - [UnHideTxResponse](#auth-common-UnHideTxResponse)
    - [UpdateRequest](#auth-common-UpdateRequest)
    - [UpdateResponse](#auth-common-UpdateResponse)

    - [GenericService](#auth-common-GenericService)

- [chaincode/auth/common/helper.proto](#chaincode_auth_common_helper-proto)
- [chaincode/auth/rbac/schema/v1/rbac.proto](#chaincode_auth_rbac_schema_v1_rbac-proto)
- [chaincode/ccbio/schema/v0/state.proto](#chaincode_ccbio_schema_v0_state-proto)
    - [Specimen](#ccbio-schema-v0-Specimen)
    - [Specimen.Georeference](#ccbio-schema-v0-Specimen-Georeference)
    - [Specimen.Grant](#ccbio-schema-v0-Specimen-Grant)
    - [Specimen.GrantsEntry](#ccbio-schema-v0-Specimen-GrantsEntry)
    - [Specimen.Image](#ccbio-schema-v0-Specimen-Image)
    - [Specimen.ImagesEntry](#ccbio-schema-v0-Specimen-ImagesEntry)
    - [Specimen.Loan](#ccbio-schema-v0-Specimen-Loan)
    - [Specimen.LoansEntry](#ccbio-schema-v0-Specimen-LoansEntry)
    - [Specimen.Primary](#ccbio-schema-v0-Specimen-Primary)
    - [Specimen.Secondary](#ccbio-schema-v0-Specimen-Secondary)
    - [Specimen.Taxon](#ccbio-schema-v0-Specimen-Taxon)

- [chaincode/ccbio/schema/v0/service.proto](#chaincode_ccbio_schema_v0_service-proto)
    - [SpecimenCreateRequest](#ccbio-schema-v0-SpecimenCreateRequest)
    - [SpecimenCreateResponse](#ccbio-schema-v0-SpecimenCreateResponse)
    - [SpecimenDeleteRequest](#ccbio-schema-v0-SpecimenDeleteRequest)
    - [SpecimenDeleteResponse](#ccbio-schema-v0-SpecimenDeleteResponse)
    - [SpecimenGetByCollectionRequest](#ccbio-schema-v0-SpecimenGetByCollectionRequest)
    - [SpecimenGetByCollectionResponse](#ccbio-schema-v0-SpecimenGetByCollectionResponse)
    - [SpecimenGetHistoryRequest](#ccbio-schema-v0-SpecimenGetHistoryRequest)
    - [SpecimenGetHistoryResponse](#ccbio-schema-v0-SpecimenGetHistoryResponse)
    - [SpecimenGetListRequest](#ccbio-schema-v0-SpecimenGetListRequest)
    - [SpecimenGetListResponse](#ccbio-schema-v0-SpecimenGetListResponse)
    - [SpecimenGetRequest](#ccbio-schema-v0-SpecimenGetRequest)
    - [SpecimenGetResponse](#ccbio-schema-v0-SpecimenGetResponse)
    - [SpecimenHideTxRequest](#ccbio-schema-v0-SpecimenHideTxRequest)
    - [SpecimenHideTxResponse](#ccbio-schema-v0-SpecimenHideTxResponse)
    - [SpecimenUnHideTxRequest](#ccbio-schema-v0-SpecimenUnHideTxRequest)
    - [SpecimenUnHideTxResponse](#ccbio-schema-v0-SpecimenUnHideTxResponse)
    - [SpecimenUpdateRequest](#ccbio-schema-v0-SpecimenUpdateRequest)
    - [SpecimenUpdateResponse](#ccbio-schema-v0-SpecimenUpdateResponse)

    - [SpecimenService](#ccbio-schema-v0-SpecimenService)

- [chaincode/sample/v0/items.proto](#chaincode_sample_v0_items-proto)
    - [Book](#sample-Book)
    - [Group](#sample-Group)
    - [SimpleItem](#sample-SimpleItem)

- [Scalar Value Types](#scalar-value-types)



<a name="auth_v1_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/auth.proto



<a name="auth-Attribute"></a>

### Attribute
An attribute is used to define permissions via the value of the attribute in the
users certificate for a given msp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  | The msp of the organization that this attribute applies to |
| oid | [string](#string) |  | The oid of the attribute |
| value | [string](#string) |  | The value of the attribute required to be satisfied by the user to have the role |
| polices | [Polices](#auth-Polices) |  | The Permission that the user will have if they have the attribute |






<a name="auth-Collection"></a>

### Collection
Collection
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
Note that the types of items are stored in the default ACLEntry

key := {COLLECTION}{COLLECTION_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | The key for the ledger |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| auth_type | [AuthType](#auth-AuthType) |  |  |
| item_types | [string](#string) | repeated |  |
| reference_types | [string](#string) | repeated | [(buf.validate.field).repeated.items = { string: {prefix: &#34;type.googleapis.com/&#34;} }]; |
| admin_key | [string](#string) | repeated |  |
| default | [Polices](#auth-Polices) |  |  |






<a name="auth-FullItem"></a>

### FullItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#auth-Item) |  |  |
| history | [History](#auth-History) |  |  |
| suggestions | [Suggestion](#auth-Suggestion) | repeated |  |
| references | [Reference](#auth-Reference) | repeated |  |






<a name="auth-HiddenTx"></a>

### HiddenTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| note | [string](#string) |  |  |






<a name="auth-HiddenTxList"></a>

### HiddenTxList
Key should be {COLLECTION_ID}{auth.HiddenTxList}{ITEM_TYPE}{...ITEM_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ItemKey](#auth-ItemKey) |  | The key that is used to store the item |
| txs | [HiddenTx](#auth-HiddenTx) | repeated | The list of hidden txs by tx_id |






<a name="auth-History"></a>

### History



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [HistoryEntry](#auth-HistoryEntry) | repeated |  |
| hidden_txs | [HiddenTxList](#auth-HiddenTxList) |  |  |






<a name="auth-HistoryEntry"></a>

### HistoryEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  | The transaction id that caused the change |
| is_delete | [bool](#bool) |  | Whether the item was deleted |
| is_hidden | [bool](#bool) |  | Whether the transaction was hidden |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp of the change |
| note | [string](#string) |  | A note about the change |
| value | [google.protobuf.Any](#google-protobuf-Any) |  | The value of the item |






<a name="auth-Item"></a>

### Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [ItemKey](#auth-ItemKey) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="auth-ItemKey"></a>

### ItemKey
Keys
─────────────────────────────────────────────────────────────────────────────────────
Item Keys
When converted to its string form it will be:
- Key := {ITEM_TYPE}{COLLECTION_ID}{...ITEM_ID}

Reference Keys
Used to store references to items for case like a user having a role
When converted to its string form it will be:
{Ref}{REFERENCE_TYPE}{COLLECTION_ID}[{ITEM1_TYPE}{...ITEM1_ID}][{ITEM2_TYPE}{...ITEM2_ID}]


SubKeys
When converted to its string form it will be:
{SUB_ITEM_TYPE}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUB_ITEM_ID}
Examples
- Suggestion := {auth.Suggestion}  {COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUGGESTION_ID}
- HiddenTxList := {auth.HiddenTxList}{COLLECTION_ID}{ITEM_TYPE} {...ITEM_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| item_type | [string](#string) |  |  |
| item_id_parts | [string](#string) | repeated |  |






<a name="auth-KeySchema"></a>

### KeySchema



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_type | [string](#string) |  | The item type of the key |
| item_kind | [ItemKind](#auth-ItemKind) |  | The kind of item that the key is for |
| keys | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  | The paths that make up the key |






<a name="auth-Operation"></a>

### Operation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [Action](#auth-Action) |  |  |
| collection_id | [string](#string) |  |  |
| item_type | [string](#string) |  |  |
| secondary_item_type | [string](#string) |  |  |
| paths | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="auth-PathPolicy"></a>

### PathPolicy
This message is the tree node for operations on the state item


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | The path is a sub path of a field mask |
| full_path | [string](#string) |  |  |
| allow_sub_paths | [bool](#bool) |  |  |
| sub_paths | [PathPolicy.SubPathsEntry](#auth-PathPolicy-SubPathsEntry) | repeated | The key is a valid sub path in the type of state item |
| actions | [Action](#auth-Action) | repeated | If the policy is not set than use a parent policy unless nested policy is set |






<a name="auth-PathPolicy-SubPathsEntry"></a>

### PathPolicy.SubPathsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathPolicy](#auth-PathPolicy) |  |  |






<a name="auth-Polices"></a>

### Polices



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_policies | [Polices.ItemPoliciesEntry](#auth-Polices-ItemPoliciesEntry) | repeated | key is the item type |






<a name="auth-Polices-ItemPoliciesEntry"></a>

### Polices.ItemPoliciesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathPolicy](#auth-PathPolicy) |  |  |






<a name="auth-Reference"></a>

### Reference
Used to return the values of the items that are referenced


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reference | [ReferenceKey](#auth-ReferenceKey) |  |  |
| item1 | [Item](#auth-Item) |  |  |
| item2 | [Item](#auth-Item) |  |  |






<a name="auth-ReferenceKey"></a>

### ReferenceKey
Reference Keys
{auth.Reference}{REFERENCE_TYPE}{COLLECTION_ID}[{ITEM1_TYPE}{...ITEM1_ID}][{ITEM2_TYPE}{...ITEM2_ID}]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key1 | [ItemKey](#auth-ItemKey) |  | string reference_type = 2; |
| key2 | [ItemKey](#auth-ItemKey) |  |  |






<a name="auth-Role"></a>

### Role



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |
| polices | [Polices](#auth-Polices) |  |  |
| description | [string](#string) |  |  |
| parent_role_ids | [string](#string) | repeated |  |






<a name="auth-StateActivity"></a>

### StateActivity
Can be used as the history entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  | The transaction id that caused the change |
| msp_id | [string](#string) |  | The msp of the user that caused the change |
| user_id | [string](#string) |  | The id of the user that caused the change |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp of the change |
| note | [string](#string) |  | A note about the change |






<a name="auth-Suggestion"></a>

### Suggestion
Key should be
{auth.Suggestion}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUGGESTION_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ItemKey](#auth-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| paths | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="auth-User"></a>

### User
User
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈

key := {USER}{USER_ID.msp_id}{USER_ID.id}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="auth-UserCollectionRoles"></a>

### UserCollectionRoles



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | The collection that the user is a member of |
| msp_id | [string](#string) |  | The msp of the organization that the user&#39;s certificate is from |
| user_id | [string](#string) |  | The id of the user from the certificate |
| role_ids | [string](#string) | repeated | The roles that the user has in the collection |






<a name="auth-UserMembership"></a>

### UserMembership
Membership is used to store permissions for a user in a collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | The collection that the user is a member of |
| msp_id | [string](#string) |  | The msp of the organization that the user&#39;s certificate is from |
| user_id | [string](#string) |  | The id of the user from the certificate |
| polices | [Polices](#auth-Polices) |  | The Permissions that the user will have |








<a name="auth-Action"></a>

### Action
Action - The action to be performed during the operation

# Action Groups
  - UNSPECIFIED: null or not set
  - REGISTER:    Used to register a new user or collection (always allowed)
  - COLLECTION:  Used to manage collections permissions
  - MEMBERSHIP:  Used to manage membership of a collection
  - ITEM:      Used to manage items in a collection

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_UNSPECIFIED | 0 |  |
| ACTION_UTILITY | 1 |  |
| ACTION_VIEW | 10 | View the contents of an item |
| ACTION_CREATE | 11 | Create a new item - key must not already exist |
| ACTION_UPDATE | 12 | Update an existing item - key must already exist - potential has paths |
| ACTION_DELETE | 13 | Delete an existing item, key must already exist |
| ACTION_SUGGEST_VIEW | 14 | Suggest a change to an item, key must already exist |
| ACTION_SUGGEST_CREATE | 15 | Suggest a change to an item, key must already exist |
| ACTION_SUGGEST_DELETE | 16 | Delete a suggestion, key must already exist |
| ACTION_SUGGEST_APPROVE | 17 | Approve a suggestion and apply it to the item, key must already exist |
| ACTION_VIEW_HISTORY | 18 |  |
| ACTION_VIEW_HIDDEN_TXS | 19 |  |
| ACTION_HIDE_TX | 20 |  |
| ACTION_REFERENCE_CREATE | 21 |  |
| ACTION_REFERENCE_DELETE | 22 |  |
| ACTION_REFERENCE_VIEW | 23 |  |



<a name="auth-AuthType"></a>

### AuthType


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTH_TYPE_UNSPECIFIED | 0 |  |
| AUTH_TYPE_NONE | 1 |  |
| AUTH_TYPE_ROLE | 2 |  |
| AUTH_TYPE_IDENTITY | 3 |  |



<a name="auth-ItemKind"></a>

### ItemKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| ITEM_KIND_UNSPECIFIED | 0 |  |
| ITEM_KIND_GLOBAL_ITEM | 1 |  |
| ITEM_KIND_PRIMARY_ITEM | 2 | Item&#39;s key := {COLLECTION_ID}{TYPE}[...key_paths] |
| ITEM_KIND_SUB_ITEM | 3 | Item&#39;s key := {COLLECTION_ID}{TYPE}&lt;PrimaryKey&gt;{...key_paths} |
| ITEM_KIND_REFERENCE | 4 |  |



<a name="auth-TransactionType"></a>

### TransactionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSACTION_TYPE_UNSPECIFIED | 0 |  |
| TRANSACTION_TYPE_INVOKE | 1 |  |
| TRANSACTION_TYPE_QUERY | 2 |  |



<a name="auth-TxError"></a>

### TxError


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 |  |
| REQUEST_INVALID | 1 |  |
| RUNTIME | 2 |  |
| RUNTIME_BAD_OPS | 3 |  |
| KEY_NOT_FOUND | 4 | The provided key is not in the world state |
| KEY_ALREADY_EXISTS | 5 | The provided key is already in the world state |
| COLLECTION_INVALID_ID | 11 | The collection id is invalid |
| COLLECTION_UNREGISTERED | 12 | The collection is not registered and thus cannot be accessed |
| COLLECTION_ALREADY_REGISTERED | 13 | The collection is already registered and thus cannot be registered again |
| COLLECTION_INVALID | 14 | The collection is invalid (e.g. the collection does not have a default ACLEntry) |
| COLLECTION_INVALID_ITEM_TYPE | 15 | The item type in the collection is invalid |
| COLLECTION_INVALID_ROLE_ID | 16 | The role id in the collection is invalid |
| USER_INVALID_ID | 20 | The user does not have permission to perform the operation |
| USER_UNREGISTERED | 21 | The certificate is not registered as a user and thus cannot be used |
| USER_ALREADY_REGISTERED | 22 | The certificate is already registered as a user and thus cannot be registered again |
| USER_INVALID | 23 | The user is invalid |
| USER_NO_ROLE | 24 | The user does not have a role |
| USER_PERMISSION_DENIED | 26 | USER_DELETED_ROLE = 25; The user does not have permission to perform the operation |
| ITEM_INVALID_ID | 31 | The Item&#39;s key is invalid |
| ITEM_UNREGISTERED | 32 | The Item is not registered and thus cannot be accessed |
| ITEM_ALREADY_REGISTERED | 33 | The Item is already registered and thus cannot be registered again |
| ITEM_INVALID | 34 | The Item is invalid |
| INVALID_ITEM_FIELD_PATH | 35 | The item field path is invalid for the item type |
| INVALID_ITEM_FIELD_VALUE | 36 | The value at the item field path is invalid for the item type |





<a name="auth_v1_auth-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| key_schema | KeySchema | .google.protobuf.MessageOptions | 54599 |  |
| operation | Operation | .google.protobuf.MethodOptions | 57775 |  |
| transaction_type | TransactionType | .google.protobuf.MethodOptions | 50556 |  |







<a name="chaincode_auth_common_generic-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/common/generic.proto



<a name="auth-common-AuthorizeOperationRequest"></a>

### AuthorizeOperationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation | [auth.Operation](#auth-Operation) |  |  |






<a name="auth-common-AuthorizeOperationResponse"></a>

### AuthorizeOperationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorized | [bool](#bool) |  |  |






<a name="auth-common-BootstrapRequest"></a>

### BootstrapRequest
──────────────────────────────── Invoke ───────────────────────────────────────


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| default_types | [string](#string) | repeated |  |
| add_default_setup | [bool](#bool) |  |  |






<a name="auth-common-BootstrapResponse"></a>

### BootstrapResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="auth-common-CreateCollectionRequest"></a>

### CreateCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |






<a name="auth-common-CreateCollectionResponse"></a>

### CreateCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |






<a name="auth-common-CreateRequest"></a>

### CreateRequest
──────────────────────────────── Invoke ───────────────────────────────────────
Create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |






<a name="auth-common-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |
| reason | [string](#string) |  |  |






<a name="auth-common-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-GetCurrentFullUserResponse"></a>

### GetCurrentFullUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |
| registered | [bool](#bool) |  |  |
| user_collection_roles | [auth.UserCollectionRoles](#auth-UserCollectionRoles) | repeated |  |
| user_memberships | [auth.UserMembership](#auth-UserMembership) | repeated |  |






<a name="auth-common-GetCurrentUserResponse"></a>

### GetCurrentUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |
| registered | [bool](#bool) |  |  |






<a name="auth-common-GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-HiddenTxRequest"></a>

### HiddenTxRequest
HiddenTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-HiddenTxResponse"></a>

### HiddenTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| hidden_txs | [auth.HiddenTx](#auth-HiddenTx) | repeated |  |






<a name="auth-common-HideTxRequest"></a>

### HideTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |
| hidden_tx | [auth.HiddenTx](#auth-HiddenTx) |  |  |






<a name="auth-common-HideTxResponse"></a>

### HideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |
| hidden_txs | [auth.HiddenTxList](#auth-HiddenTxList) |  |  |






<a name="auth-common-HistoryRequest"></a>

### HistoryRequest
History


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-HistoryResponse"></a>

### HistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| history | [auth.History](#auth-History) |  | repeated auth. s = 1; |






<a name="auth-common-ListByAttrsRequest"></a>

### ListByAttrsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| item | [auth.Item](#auth-Item) |  |  |
| num_attrs | [int32](#int32) |  |  |






<a name="auth-common-ListByAttrsResponse"></a>

### ListByAttrsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| items | [auth.Item](#auth-Item) | repeated |  |






<a name="auth-common-ListByCollectionRequest"></a>

### ListByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-ListByCollectionResponse"></a>

### ListByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| items | [auth.Item](#auth-Item) | repeated |  |






<a name="auth-common-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| items | [auth.Item](#auth-Item) | repeated |  |






<a name="auth-common-ReferenceByCollectionRequest"></a>

### ReferenceByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| collection_id | [string](#string) |  |  |






<a name="auth-common-ReferenceByCollectionResponse"></a>

### ReferenceByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| references | [auth.ReferenceKey](#auth-ReferenceKey) | repeated |  |






<a name="auth-common-ReferenceByItemRequest"></a>

### ReferenceByItemRequest
Get all of the collections


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| collection_id | [string](#string) |  |  |
| item_key | [auth.ItemKey](#auth-ItemKey) |  |  |






<a name="auth-common-ReferenceByItemResponse"></a>

### ReferenceByItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| references | [auth.ReferenceKey](#auth-ReferenceKey) | repeated |  |






<a name="auth-common-ReferenceByPartialKeyRequest"></a>

### ReferenceByPartialKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| reference | [auth.ReferenceKey](#auth-ReferenceKey) |  |  |






<a name="auth-common-ReferenceByPartialKeyResponse"></a>

### ReferenceByPartialKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| references | [auth.ReferenceKey](#auth-ReferenceKey) | repeated |  |






<a name="auth-common-ReferenceCreateRequest"></a>

### ReferenceCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref_key | [auth.ReferenceKey](#auth-ReferenceKey) |  |  |






<a name="auth-common-ReferenceCreateResponse"></a>

### ReferenceCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref_key | [auth.ReferenceKey](#auth-ReferenceKey) |  |  |






<a name="auth-common-ReferenceDeleteRequest"></a>

### ReferenceDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref_key | [auth.ReferenceKey](#auth-ReferenceKey) |  |  |






<a name="auth-common-ReferenceDeleteResponse"></a>

### ReferenceDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ref_key | [auth.ReferenceKey](#auth-ReferenceKey) |  |  |






<a name="auth-common-ReferenceRequest"></a>

### ReferenceRequest
════════════════════════════════ References ═════════════════════════════════════
──────────────────────────────── Query ──────────────────────────────────────────


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reference | [auth.ReferenceKey](#auth-ReferenceKey) |  | buf:lint:ignore FIELD_SAME_TYPE |






<a name="auth-common-ReferenceResponse"></a>

### ReferenceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exists | [bool](#bool) |  |  |
| reference | [auth.Reference](#auth-Reference) |  |  |






<a name="auth-common-SuggestionApproveRequest"></a>

### SuggestionApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [auth.ItemKey](#auth-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="auth-common-SuggestionApproveResponse"></a>

### SuggestionApproveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [auth.Suggestion](#auth-Suggestion) |  |  |
| item | [auth.Item](#auth-Item) |  |  |






<a name="auth-common-SuggestionByPartialKeyRequest"></a>

### SuggestionByPartialKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| num_attrs | [int32](#int32) |  |  |
| item_key | [auth.ItemKey](#auth-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |






<a name="auth-common-SuggestionByPartialKeyResponse"></a>

### SuggestionByPartialKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| suggestions | [auth.Suggestion](#auth-Suggestion) | repeated |  |






<a name="auth-common-SuggestionCreateRequest"></a>

### SuggestionCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [auth.Suggestion](#auth-Suggestion) |  |  |






<a name="auth-common-SuggestionCreateResponse"></a>

### SuggestionCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [auth.Suggestion](#auth-Suggestion) |  |  |






<a name="auth-common-SuggestionDeleteRequest"></a>

### SuggestionDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [auth.ItemKey](#auth-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="auth-common-SuggestionDeleteResponse"></a>

### SuggestionDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [auth.Suggestion](#auth-Suggestion) |  |  |






<a name="auth-common-SuggestionListByCollectionRequest"></a>

### SuggestionListByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| collection_id | [string](#string) |  |  |






<a name="auth-common-SuggestionListByCollectionResponse"></a>

### SuggestionListByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| suggestions | [auth.Suggestion](#auth-Suggestion) | repeated |  |






<a name="auth-common-SuggestionListByItemRequest"></a>

### SuggestionListByItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [auth.ItemKey](#auth-ItemKey) |  |  |






<a name="auth-common-SuggestionListByItemResponse"></a>

### SuggestionListByItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestions | [auth.Suggestion](#auth-Suggestion) | repeated |  |






<a name="auth-common-SuggestionListRequest"></a>

### SuggestionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  | auth.Item item = 3; |






<a name="auth-common-SuggestionListResponse"></a>

### SuggestionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| suggestions | [auth.Suggestion](#auth-Suggestion) | repeated |  |






<a name="auth-common-SuggestionRequest"></a>

### SuggestionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [auth.ItemKey](#auth-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |






<a name="auth-common-SuggestionResponse"></a>

### SuggestionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [auth.Suggestion](#auth-Suggestion) |  |  |






<a name="auth-common-UnHideTxRequest"></a>

### UnHideTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |
| tx_id | [string](#string) |  |  |






<a name="auth-common-UnHideTxResponse"></a>

### UnHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |
| hidden_txs | [auth.HiddenTxList](#auth-HiddenTxList) |  |  |






<a name="auth-common-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="auth-common-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [auth.Item](#auth-Item) |  |  |












<a name="auth-common-GenericService"></a>

### GenericService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCurrentUser | [.google.protobuf.Empty](#google-protobuf-Empty) | [GetCurrentUserResponse](#auth-common-GetCurrentUserResponse) | ══════════════════════════════════ Helper ═════════════════════════════════════ ────────────────────────────────── Query ────────────────────────────────────── rpc GetAllTypes(google.protobuf.Empty) returns (GetAllTypesResponse) { option (auth.transaction_type) = TRANSACTION_TYPE_QUERY; option (auth.operation) = {action: ACTION_UTILITY}; } |
| CreateUser | [.google.protobuf.Empty](#google-protobuf-Empty) | [CreateUserResponse](#auth-common-CreateUserResponse) |  |
| AuthorizeOperation | [AuthorizeOperationRequest](#auth-common-AuthorizeOperationRequest) | [AuthorizeOperationResponse](#auth-common-AuthorizeOperationResponse) |  |
| Get | [GetRequest](#auth-common-GetRequest) | [GetResponse](#auth-common-GetResponse) |  |
| List | [ListRequest](#auth-common-ListRequest) | [ListResponse](#auth-common-ListResponse) |  |
| ListByCollection | [ListByCollectionRequest](#auth-common-ListByCollectionRequest) | [ListByCollectionResponse](#auth-common-ListByCollectionResponse) |  |
| ListByAttrs | [ListByAttrsRequest](#auth-common-ListByAttrsRequest) | [ListByAttrsResponse](#auth-common-ListByAttrsResponse) |  |
| Create | [CreateRequest](#auth-common-CreateRequest) | [CreateResponse](#auth-common-CreateResponse) |  |
| Update | [UpdateRequest](#auth-common-UpdateRequest) | [UpdateResponse](#auth-common-UpdateResponse) |  |
| Delete | [DeleteRequest](#auth-common-DeleteRequest) | [DeleteResponse](#auth-common-DeleteResponse) |  |
| History | [HistoryRequest](#auth-common-HistoryRequest) | [HistoryResponse](#auth-common-HistoryResponse) |  |
| HiddenTx | [HiddenTxRequest](#auth-common-HiddenTxRequest) | [HiddenTxResponse](#auth-common-HiddenTxResponse) |  |
| HideTx | [HideTxRequest](#auth-common-HideTxRequest) | [HideTxResponse](#auth-common-HideTxResponse) |  |
| UnHideTx | [UnHideTxRequest](#auth-common-UnHideTxRequest) | [UnHideTxResponse](#auth-common-UnHideTxResponse) |  |
| Reference | [ReferenceRequest](#auth-common-ReferenceRequest) | [ReferenceResponse](#auth-common-ReferenceResponse) |  |
| ReferenceByItem | [ReferenceByItemRequest](#auth-common-ReferenceByItemRequest) | [ReferenceByItemResponse](#auth-common-ReferenceByItemResponse) |  |
| ReferenceByPartialKey | [ReferenceByPartialKeyRequest](#auth-common-ReferenceByPartialKeyRequest) | [ReferenceByPartialKeyResponse](#auth-common-ReferenceByPartialKeyResponse) |  |
| ReferenceCreate | [ReferenceCreateRequest](#auth-common-ReferenceCreateRequest) | [ReferenceCreateResponse](#auth-common-ReferenceCreateResponse) |  |
| ReferenceDelete | [ReferenceDeleteRequest](#auth-common-ReferenceDeleteRequest) | [ReferenceDeleteResponse](#auth-common-ReferenceDeleteResponse) |  |
| Suggestion | [SuggestionRequest](#auth-common-SuggestionRequest) | [SuggestionResponse](#auth-common-SuggestionResponse) |  |
| SuggestionListByCollection | [SuggestionListByCollectionRequest](#auth-common-SuggestionListByCollectionRequest) | [SuggestionListByCollectionResponse](#auth-common-SuggestionListByCollectionResponse) |  |
| SuggestionByPartialKey | [SuggestionByPartialKeyRequest](#auth-common-SuggestionByPartialKeyRequest) | [SuggestionByPartialKeyResponse](#auth-common-SuggestionByPartialKeyResponse) |  |
| SuggestionCreate | [SuggestionCreateRequest](#auth-common-SuggestionCreateRequest) | [SuggestionCreateResponse](#auth-common-SuggestionCreateResponse) | ──────────────────────────────── Invoke ─────────────────────────────────────── |
| SuggestionDelete | [SuggestionDeleteRequest](#auth-common-SuggestionDeleteRequest) | [SuggestionDeleteResponse](#auth-common-SuggestionDeleteResponse) |  |
| SuggestionApprove | [SuggestionApproveRequest](#auth-common-SuggestionApproveRequest) | [SuggestionApproveResponse](#auth-common-SuggestionApproveResponse) |  |





<a name="chaincode_auth_common_helper-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/common/helper.proto












<a name="chaincode_auth_rbac_schema_v1_rbac-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/rbac/schema/v1/rbac.proto












<a name="chaincode_ccbio_schema_v0_state-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/ccbio/schema/v0/state.proto



<a name="ccbio-schema-v0-Specimen"></a>

### Specimen



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| specimen_id | [string](#string) |  |  |
| primary | [Specimen.Primary](#ccbio-schema-v0-Specimen-Primary) |  |  |
| secondary | [Specimen.Secondary](#ccbio-schema-v0-Specimen-Secondary) |  |  |
| taxon | [Specimen.Taxon](#ccbio-schema-v0-Specimen-Taxon) |  |  |
| georeference | [Specimen.Georeference](#ccbio-schema-v0-Specimen-Georeference) |  |  |
| images | [Specimen.ImagesEntry](#ccbio-schema-v0-Specimen-ImagesEntry) | repeated |  |
| loans | [Specimen.LoansEntry](#ccbio-schema-v0-Specimen-LoansEntry) | repeated |  |
| grants | [Specimen.GrantsEntry](#ccbio-schema-v0-Specimen-GrantsEntry) | repeated |  |






<a name="ccbio-schema-v0-Specimen-Georeference"></a>

### Specimen.Georeference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| country | [string](#string) |  |  |
| state_province | [string](#string) |  |  |
| county | [string](#string) |  |  |
| locality | [string](#string) |  |  |
| latitude | [string](#string) |  |  |
| longitude | [string](#string) |  |  |
| habitat | [string](#string) |  |  |
| notes | [string](#string) | repeated |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-Specimen-Grant"></a>

### Specimen.Grant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| granted_by | [string](#string) |  |  |
| granted_to | [string](#string) |  |  |
| granted_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-Specimen-GrantsEntry"></a>

### Specimen.GrantsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Grant](#ccbio-schema-v0-Specimen-Grant) |  |  |






<a name="ccbio-schema-v0-Specimen-Image"></a>

### Specimen.Image
Mapped Types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| url | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| hash | [string](#string) |  |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-Specimen-ImagesEntry"></a>

### Specimen.ImagesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Image](#ccbio-schema-v0-Specimen-Image) |  |  |






<a name="ccbio-schema-v0-Specimen-Loan"></a>

### Specimen.Loan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| loaned_by | [string](#string) |  |  |
| loaned_to | [string](#string) |  |  |
| loaned_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-Specimen-LoansEntry"></a>

### Specimen.LoansEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Loan](#ccbio-schema-v0-Specimen-Loan) |  |  |






<a name="ccbio-schema-v0-Specimen-Primary"></a>

### Specimen.Primary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| catalog_number | [string](#string) |  |  |
| accession_number | [string](#string) |  |  |
| field_number | [string](#string) |  |  |
| tissue_number | [string](#string) |  |  |
| cataloger | [string](#string) |  |  |
| collector | [string](#string) |  |  |
| determiner | [string](#string) |  |  |
| field_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| catalog_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| determined_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| determined_reason | [string](#string) |  |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-Specimen-Secondary"></a>

### Specimen.Secondary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| preparation | [string](#string) |  |  |
| condition | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-Specimen-Taxon"></a>

### Specimen.Taxon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kingdom | [string](#string) |  |  |
| phylum | [string](#string) |  |  |
| class | [string](#string) |  |  |
| order | [string](#string) |  |  |
| family | [string](#string) |  |  |
| genus | [string](#string) |  |  |
| species | [string](#string) |  |  |
| subspecies | [string](#string) |  |  |
| last_modified | [auth.StateActivity](#auth-StateActivity) |  |  |















<a name="chaincode_ccbio_schema_v0_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/ccbio/schema/v0/service.proto



<a name="ccbio-schema-v0-SpecimenCreateRequest"></a>

### SpecimenCreateRequest
SpecimenCreate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |






<a name="ccbio-schema-v0-SpecimenCreateResponse"></a>

### SpecimenCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |






<a name="ccbio-schema-v0-SpecimenDeleteRequest"></a>

### SpecimenDeleteRequest
SpecimenDelete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| specimen_id | [string](#string) |  |  |






<a name="ccbio-schema-v0-SpecimenDeleteResponse"></a>

### SpecimenDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |






<a name="ccbio-schema-v0-SpecimenGetByCollectionRequest"></a>

### SpecimenGetByCollectionRequest
SpecimenGetByCollection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |






<a name="ccbio-schema-v0-SpecimenGetByCollectionResponse"></a>

### SpecimenGetByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimens | [Specimen](#ccbio-schema-v0-Specimen) | repeated |  |






<a name="ccbio-schema-v0-SpecimenGetHistoryRequest"></a>

### SpecimenGetHistoryRequest
SpecimenGetHistory


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| specimen_id | [string](#string) |  |  |
| include_hidden | [bool](#bool) |  |  |






<a name="ccbio-schema-v0-SpecimenGetHistoryResponse"></a>

### SpecimenGetHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| history | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-SpecimenGetListRequest"></a>

### SpecimenGetListRequest
SpecimenGetList


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| page_size | [int32](#int32) |  |  |






<a name="ccbio-schema-v0-SpecimenGetListResponse"></a>

### SpecimenGetListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| specimens | [Specimen](#ccbio-schema-v0-Specimen) | repeated |  |






<a name="ccbio-schema-v0-SpecimenGetRequest"></a>

### SpecimenGetRequest
Evaluate
SpecimenGet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | Specimen.Id id = 1 [(buf.validate.field).required = true]; |
| specimen_id | [string](#string) |  |  |






<a name="ccbio-schema-v0-SpecimenGetResponse"></a>

### SpecimenGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |






<a name="ccbio-schema-v0-SpecimenHideTxRequest"></a>

### SpecimenHideTxRequest
SpecimenHideTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| specimen_id | [string](#string) |  |  |
| tx | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-SpecimenHideTxResponse"></a>

### SpecimenHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |






<a name="ccbio-schema-v0-SpecimenUnHideTxRequest"></a>

### SpecimenUnHideTxRequest
SpecimenUnHideTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| specimen_id | [string](#string) |  |  |
| tx | [auth.StateActivity](#auth-StateActivity) |  |  |






<a name="ccbio-schema-v0-SpecimenUnHideTxResponse"></a>

### SpecimenUnHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |






<a name="ccbio-schema-v0-SpecimenUpdateRequest"></a>

### SpecimenUpdateRequest
SpecimenUpdate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="ccbio-schema-v0-SpecimenUpdateResponse"></a>

### SpecimenUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#ccbio-schema-v0-Specimen) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |












<a name="ccbio-schema-v0-SpecimenService"></a>

### SpecimenService
Specimen functions

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SpecimenGet | [SpecimenGetRequest](#ccbio-schema-v0-SpecimenGetRequest) | [SpecimenGetResponse](#ccbio-schema-v0-SpecimenGetResponse) |  |
| SpecimenGetList | [SpecimenGetListRequest](#ccbio-schema-v0-SpecimenGetListRequest) | [SpecimenGetListResponse](#ccbio-schema-v0-SpecimenGetListResponse) |  |
| SpecimenGetByCollection | [SpecimenGetByCollectionRequest](#ccbio-schema-v0-SpecimenGetByCollectionRequest) | [SpecimenGetByCollectionResponse](#ccbio-schema-v0-SpecimenGetByCollectionResponse) |  |
| SpecimenGetHistory | [SpecimenGetHistoryRequest](#ccbio-schema-v0-SpecimenGetHistoryRequest) | [SpecimenGetHistoryResponse](#ccbio-schema-v0-SpecimenGetHistoryResponse) |  |
| SpecimenCreate | [SpecimenCreateRequest](#ccbio-schema-v0-SpecimenCreateRequest) | [SpecimenCreateResponse](#ccbio-schema-v0-SpecimenCreateResponse) |  |
| SpecimenUpdate | [SpecimenUpdateRequest](#ccbio-schema-v0-SpecimenUpdateRequest) | [SpecimenUpdateResponse](#ccbio-schema-v0-SpecimenUpdateResponse) |  |
| SpecimenDelete | [SpecimenDeleteRequest](#ccbio-schema-v0-SpecimenDeleteRequest) | [SpecimenDeleteResponse](#ccbio-schema-v0-SpecimenDeleteResponse) |  |
| SpecimenHideTx | [SpecimenHideTxRequest](#ccbio-schema-v0-SpecimenHideTxRequest) | [SpecimenHideTxResponse](#ccbio-schema-v0-SpecimenHideTxResponse) |  |
| SpecimenUnHideTx | [SpecimenUnHideTxRequest](#ccbio-schema-v0-SpecimenUnHideTxRequest) | [SpecimenUnHideTxResponse](#ccbio-schema-v0-SpecimenUnHideTxResponse) |  |





<a name="chaincode_sample_v0_items-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/sample/v0/items.proto



<a name="sample-Book"></a>

### Book



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| isbn | [string](#string) |  |  |
| book_title | [string](#string) |  |  |
| author | [string](#string) |  |  |
| year | [int32](#int32) |  |  |
| publisher | [string](#string) |  |  |
| language | [string](#string) |  |  |
| description | [string](#string) |  |  |






<a name="sample-Group"></a>

### Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| group_id | [string](#string) |  |  |
| item1 | [SimpleItem](#sample-SimpleItem) |  |  |
| item2 | [SimpleItem](#sample-SimpleItem) |  |  |






<a name="sample-SimpleItem"></a>

### SimpleItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| quantity | [int32](#int32) |  |  |















## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

