# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [auth/v1/auth.proto](#auth_v1_auth-proto)
    - [ACEntry](#auth-ACEntry)
    - [ACEntryTree](#auth-ACEntryTree)
    - [ACEntryTree.ChildrenEntry](#auth-ACEntryTree-ChildrenEntry)
    - [Attribute](#auth-Attribute)
    - [Collection](#auth-Collection)
    - [CombinedList](#auth-CombinedList)
    - [FullObject](#auth-FullObject)
    - [HiddenTx](#auth-HiddenTx)
    - [HiddenTxList](#auth-HiddenTxList)
    - [History](#auth-History)
    - [HistoryEntry](#auth-HistoryEntry)
    - [KeySchema](#auth-KeySchema)
    - [Object](#auth-Object)
    - [ObjectKey](#auth-ObjectKey)
    - [ObjectList](#auth-ObjectList)
    - [ObjectPolicy](#auth-ObjectPolicy)
    - [Operation](#auth-Operation)
    - [PathPolicy](#auth-PathPolicy)
    - [Reference](#auth-Reference)
    - [ReferenceList](#auth-ReferenceList)
    - [Role](#auth-Role)
    - [StateActivity](#auth-StateActivity)
    - [SubObjectKey](#auth-SubObjectKey)
    - [Suggestion](#auth-Suggestion)
    - [SuggestionList](#auth-SuggestionList)
    - [User](#auth-User)

    - [Action](#auth-Action)
    - [ObjectKind](#auth-ObjectKind)
    - [TransactionType](#auth-TransactionType)
    - [TxError](#auth-TxError)

    - [File-level Extensions](#auth_v1_auth-proto-extensions)
    - [File-level Extensions](#auth_v1_auth-proto-extensions)
    - [File-level Extensions](#auth_v1_auth-proto-extensions)

- [chaincode/auth/common/collections.proto](#chaincode_auth_common_collections-proto)
    - [CollectionCreateRequest](#auth-common-CollectionCreateRequest)
    - [CollectionCreateResponse](#auth-common-CollectionCreateResponse)
    - [CollectionGetHistoryRequest](#auth-common-CollectionGetHistoryRequest)
    - [CollectionGetHistoryResponse](#auth-common-CollectionGetHistoryResponse)
    - [CollectionGetListRequest](#auth-common-CollectionGetListRequest)
    - [CollectionGetListResponse](#auth-common-CollectionGetListResponse)
    - [CollectionGetRequest](#auth-common-CollectionGetRequest)
    - [CollectionGetResponse](#auth-common-CollectionGetResponse)
    - [CollectionHideTxRequest](#auth-common-CollectionHideTxRequest)
    - [CollectionHideTxResponse](#auth-common-CollectionHideTxResponse)
    - [CollectionUpdateRequest](#auth-common-CollectionUpdateRequest)
    - [CollectionUpdateResponse](#auth-common-CollectionUpdateResponse)

    - [CollectionService](#auth-common-CollectionService)

- [chaincode/auth/common/generic.proto](#chaincode_auth_common_generic-proto)
    - [CreateRequest](#auth-common-CreateRequest)
    - [CreateResponse](#auth-common-CreateResponse)
    - [DeleteRequest](#auth-common-DeleteRequest)
    - [DeleteResponse](#auth-common-DeleteResponse)
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
    - [SuggestionListByObjectRequest](#auth-common-SuggestionListByObjectRequest)
    - [SuggestionListByObjectResponse](#auth-common-SuggestionListByObjectResponse)
    - [SuggestionListRequest](#auth-common-SuggestionListRequest)
    - [SuggestionListResponse](#auth-common-SuggestionListResponse)
    - [SuggestionRequest](#auth-common-SuggestionRequest)
    - [SuggestionResponse](#auth-common-SuggestionResponse)
    - [UnHideTxRequest](#auth-common-UnHideTxRequest)
    - [UnHideTxResponse](#auth-common-UnHideTxResponse)
    - [UpdateRequest](#auth-common-UpdateRequest)
    - [UpdateResponse](#auth-common-UpdateResponse)

    - [GenericService](#auth-common-GenericService)

- [chaincode/auth/common/users.proto](#chaincode_auth_common_users-proto)
    - [UserCreateRequest](#auth-common-UserCreateRequest)
    - [UserCreateResponse](#auth-common-UserCreateResponse)
    - [UserDeleteRequest](#auth-common-UserDeleteRequest)
    - [UserDeleteResponse](#auth-common-UserDeleteResponse)
    - [UserGetCurrentIdResponse](#auth-common-UserGetCurrentIdResponse)
    - [UserGetCurrentResponse](#auth-common-UserGetCurrentResponse)
    - [UserGetHiddenTxRequest](#auth-common-UserGetHiddenTxRequest)
    - [UserGetHiddenTxResponse](#auth-common-UserGetHiddenTxResponse)
    - [UserGetHistoryRequest](#auth-common-UserGetHistoryRequest)
    - [UserGetHistoryResponse](#auth-common-UserGetHistoryResponse)
    - [UserGetListRequest](#auth-common-UserGetListRequest)
    - [UserGetListResponse](#auth-common-UserGetListResponse)
    - [UserGetRequest](#auth-common-UserGetRequest)
    - [UserGetResponse](#auth-common-UserGetResponse)
    - [UserHideTxRequest](#auth-common-UserHideTxRequest)
    - [UserHideTxResponse](#auth-common-UserHideTxResponse)
    - [UserUpdateRequest](#auth-common-UserUpdateRequest)
    - [UserUpdateResponse](#auth-common-UserUpdateResponse)

    - [UserService](#auth-common-UserService)

- [chaincode/auth/rbac/schema/v1/rbac.proto](#chaincode_auth_rbac_schema_v1_rbac-proto)
    - [Membership](#rbac-schema-v1-Membership)
    - [MembershipCreateRequest](#rbac-schema-v1-MembershipCreateRequest)
    - [MembershipCreateResponse](#rbac-schema-v1-MembershipCreateResponse)
    - [MembershipDeleteRequest](#rbac-schema-v1-MembershipDeleteRequest)
    - [MembershipDeleteResponse](#rbac-schema-v1-MembershipDeleteResponse)
    - [MembershipGetByCollectionRequest](#rbac-schema-v1-MembershipGetByCollectionRequest)
    - [MembershipGetByCollectionResponse](#rbac-schema-v1-MembershipGetByCollectionResponse)
    - [MembershipGetByUserRequest](#rbac-schema-v1-MembershipGetByUserRequest)
    - [MembershipGetByUserResponse](#rbac-schema-v1-MembershipGetByUserResponse)
    - [MembershipGetHistoryRequest](#rbac-schema-v1-MembershipGetHistoryRequest)
    - [MembershipGetHistoryResponse](#rbac-schema-v1-MembershipGetHistoryResponse)
    - [MembershipGetListRequest](#rbac-schema-v1-MembershipGetListRequest)
    - [MembershipGetListResponse](#rbac-schema-v1-MembershipGetListResponse)
    - [MembershipGetRequest](#rbac-schema-v1-MembershipGetRequest)
    - [MembershipGetResponse](#rbac-schema-v1-MembershipGetResponse)
    - [RoleCreateRequest](#rbac-schema-v1-RoleCreateRequest)
    - [RoleCreateResponse](#rbac-schema-v1-RoleCreateResponse)
    - [RoleDeleteRequest](#rbac-schema-v1-RoleDeleteRequest)
    - [RoleDeleteResponse](#rbac-schema-v1-RoleDeleteResponse)
    - [RoleGetByCollectionRequest](#rbac-schema-v1-RoleGetByCollectionRequest)
    - [RoleGetByCollectionResponse](#rbac-schema-v1-RoleGetByCollectionResponse)
    - [RoleGetHistoryRequest](#rbac-schema-v1-RoleGetHistoryRequest)
    - [RoleGetHistoryResponse](#rbac-schema-v1-RoleGetHistoryResponse)
    - [RoleGetListRequest](#rbac-schema-v1-RoleGetListRequest)
    - [RoleGetListResponse](#rbac-schema-v1-RoleGetListResponse)
    - [RoleGetRequest](#rbac-schema-v1-RoleGetRequest)
    - [RoleGetResponse](#rbac-schema-v1-RoleGetResponse)
    - [RoleUpdateRequest](#rbac-schema-v1-RoleUpdateRequest)
    - [RoleUpdateResponse](#rbac-schema-v1-RoleUpdateResponse)

    - [RBACService](#rbac-schema-v1-RBACService)

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
- [Scalar Value Types](#scalar-value-types)



<a name="auth_v1_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## auth/v1/auth.proto



<a name="auth-ACEntry"></a>

### ACEntry
Access Control Entry for use in Radix Tree


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [ObjectPolicy](#auth-ObjectPolicy) | repeated |  |
| view_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="auth-ACEntryTree"></a>

### ACEntryTree
Access Control Entry for use in Hash Tree


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| root | [ObjectPolicy](#auth-ObjectPolicy) |  |  |
| is_leaf | [bool](#bool) |  |  |
| children | [ACEntryTree.ChildrenEntry](#auth-ACEntryTree-ChildrenEntry) | repeated |  |






<a name="auth-ACEntryTree-ChildrenEntry"></a>

### ACEntryTree.ChildrenEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ACEntryTree](#auth-ACEntryTree) |  |  |






<a name="auth-Attribute"></a>

### Attribute



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  | The msp of the organization that this attribute applies to |
| oid | [string](#string) |  | The oid of the attribute |
| value | [string](#string) |  | The value of the attribute required to be satisfied by the user to have the role |
| role_id | [string](#string) |  | The role that the user must have to satisfy the attribute |






<a name="auth-Collection"></a>

### Collection
Collection
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
Note that the types of objects are stored in the default ACEntry

key := {COLLECTION}{COLLECTION_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| object_types | [string](#string) | repeated |  |
| default | [ACEntry](#auth-ACEntry) |  |  |






<a name="auth-CombinedList"></a>

### CombinedList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| objects | [ObjectList](#auth-ObjectList) |  |  |
| references | [ReferenceList](#auth-ReferenceList) |  |  |






<a name="auth-FullObject"></a>

### FullObject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [ObjectKey](#auth-ObjectKey) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |
| history | [History](#auth-History) |  |  |
| suggestions | [SuggestionList](#auth-SuggestionList) |  |  |
| references | [ReferenceList](#auth-ReferenceList) |  |  |






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
Key should be {COLLECTION_ID}{auth.HiddenTxList}{OBJECT_TYPE}{...OBJECT_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ObjectKey](#auth-ObjectKey) |  | The key that is used to store the object |
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
| tx_id | [string](#string) |  |  |
| is_delete | [bool](#bool) |  |  |
| is_hidden | [bool](#bool) |  |  |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| note | [string](#string) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="auth-KeySchema"></a>

### KeySchema



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| namespace | [string](#string) |  |  |
| object_kind | [ObjectKind](#auth-ObjectKind) |  |  |
| keys | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |
| primary_key | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |
| secondary_keys | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="auth-Object"></a>

### Object



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [ObjectKey](#auth-ObjectKey) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="auth-ObjectKey"></a>

### ObjectKey
Object Keys
When converted to its string form it will be:
  - Key := {OBJECT_TYPE}{COLLECTION_ID}{...OBJECT_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| object_type | [string](#string) |  |  |
| object_id_parts | [string](#string) | repeated |  |






<a name="auth-ObjectList"></a>

### ObjectList
Lists
─────────────────────────────────────────────────────────────────────────────────


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| objects | [Object](#auth-Object) | repeated |  |






<a name="auth-ObjectPolicy"></a>

### ObjectPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_type | [string](#string) |  |  |
| object_namespace | [string](#string) |  |  |
| policies | [PathPolicy](#auth-PathPolicy) | repeated |  |






<a name="auth-Operation"></a>

### Operation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [Action](#auth-Action) |  |  |
| collection_id | [string](#string) |  |  |
| namespace | [string](#string) |  |  |
| secondary_namespace | [string](#string) |  |  |
| paths | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="auth-PathPolicy"></a>

### PathPolicy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  |  |
| actions | [Action](#auth-Action) | repeated |  |






<a name="auth-Reference"></a>

### Reference
Reference Keys
{REFERENCE_TYPE}{COLLECTION_ID}[{OBJECT1_TYPE}{...OBJECT1_ID}][{OBJECT2_TYPE}{...OBJECT2_ID}]


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| reference_type | [string](#string) |  |  |
| key_1 | [ObjectKey](#auth-ObjectKey) |  |  |
| key_2 | [ObjectKey](#auth-ObjectKey) |  |  |






<a name="auth-ReferenceList"></a>

### ReferenceList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| objects | [Reference](#auth-Reference) | repeated |  |






<a name="auth-Role"></a>

### Role



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |
| ac | [ACEntry](#auth-ACEntry) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| parent_role_ids | [string](#string) | repeated |  |






<a name="auth-StateActivity"></a>

### StateActivity
Can be used as the history entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| note | [string](#string) |  |  |






<a name="auth-SubObjectKey"></a>

### SubObjectKey
SubKeys
// When converted to its string form it will be:
 - {SUB_OBJECT_TYPE}{COLLECTION_ID}{OBJECT_TYPE}{...OBJECT_ID}{SUB_OBJECT_ID}
     Examples
       - Suggestion := {auth.Suggestion}  {COLLECTION_ID}{OBJECT_TYPE}{...OBJECT_ID}{SUGGESTION_ID}
       - HiddenTxList := {auth.HiddenTxList}{COLLECTION_ID}{OBJECT_TYPE} {...OBJECT_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| sub_object_type | [string](#string) |  |  |
| object_type | [string](#string) |  |  |
| object_id_parts | [string](#string) | repeated | The key is the fields of the object that make up the primary key The value is the value of the field |
| sub_object_id | [string](#string) |  |  |






<a name="auth-Suggestion"></a>

### Suggestion
Key should be
{COLLECTION_ID}{auth.Suggestion}{OBJECT_TYPE}{...OBJECT_ID}{SUGGESTION_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ObjectKey](#auth-ObjectKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| paths | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="auth-SuggestionList"></a>

### SuggestionList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ObjectKey](#auth-ObjectKey) |  |  |
| suggestions | [Suggestion](#auth-Suggestion) | repeated |  |






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








<a name="auth-Action"></a>

### Action
Action - The action to be performed during the operation

# Action Groups
  - UNSPECIFIED: null or not set
  - REGISTER:    Used to register a new user or collection (always allowed)
  - COLLECTION:  Used to manage collections permissions
  - MEMBERSHIP:  Used to manage membership of a collection
  - OBJECT:      Used to manage objects in a collection

| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_UNSPECIFIED | 0 |  |
| ACTION_VIEW | 10 | View the contents of an object |
| ACTION_CREATE | 11 | Create a new object - key must not already exist |
| ACTION_UPDATE | 12 | Update an existing object - key must already exist |
| ACTION_DELETE | 13 | Delete an existing object, key must already exist |
| ACTION_SUGGEST_VIEW | 14 | Suggest a change to an object, key must already exist |
| ACTION_SUGGEST_CREATE | 15 | Suggest a change to an object, key must already exist |
| ACTION_SUGGEST_DELETE | 16 | Delete a suggestion, key must already exist |
| ACTION_SUGGEST_APPROVE | 17 | Approve a suggestion and apply it to the object, key must already exist |
| ACTION_VIEW_HISTORY | 18 |  |
| ACTION_VIEW_HIDDEN_TXS | 19 |  |
| ACTION_HIDE_TX | 20 |  |
| ACTION_CREATE_REFERENCE | 21 |  |
| ACTION_DELETE_REFERENCE | 22 |  |
| ACTION_VIEW_REFERENCE | 23 |  |



<a name="auth-ObjectKind"></a>

### ObjectKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| OBJECT_KIND_UNSPECIFIED | 0 |  |
| OBJECT_KIND_GLOBAL_OBJECT | 1 |  |
| OBJECT_KIND_PRIMARY_OBJECT | 2 | Object&#39;s key := {COLLECTION_ID}{TYPE}[...key_paths] |
| OBJECT_KIND_SUB_OBJECT | 3 | Object&#39;s key := {COLLECTION_ID}{TYPE}&lt;PrimaryKey&gt;{...key_paths} |
| OBJECT_KIND_REFERENCE | 4 |  |



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
| KEY_NOT_FOUND | 4 |  |
| KEY_ALREADY_EXISTS | 5 |  |
| COLLECTION_INVALID_ID | 11 | Collection Errors |
| COLLECTION_UNREGISTERED | 12 |  |
| COLLECTION_ALREADY_REGISTERED | 13 |  |
| COLLECTION_INVALID | 14 |  |
| COLLECTION_INVALID_OBJECT_TYPE | 15 |  |
| COLLECTION_INVALID_ROLE_ID | 16 |  |
| USER_INVALID_ID | 20 | The user does not have permission to perform the operation |
| USER_UNREGISTERED | 21 |  |
| USER_ALREADY_REGISTERED | 22 |  |
| USER_INVALID | 23 |  |
| USER_NO_ROLE | 24 |  |
| USER_DELETED_ROLE | 25 |  |
| USER_PERMISSION_DENIED | 26 |  |
| OBJECT_INVALID_ID | 31 |  |
| OBJECT_UNREGISTERED | 32 |  |
| OBJECT_ALREADY_REGISTERED | 33 |  |
| OBJECT_INVALID | 34 |  |
| INVALID_OBJECT_FIELD_PATH | 35 |  |
| INVALID_OBJECT_FIELD_VALUE | 36 |  |





<a name="auth_v1_auth-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| key_schema | KeySchema | .google.protobuf.MessageOptions | 54599 |  |
| operation | Operation | .google.protobuf.MethodOptions | 57775 |  |
| transaction_type | TransactionType | .google.protobuf.MethodOptions | 50556 |  |







<a name="chaincode_auth_common_collections-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/common/collections.proto



<a name="auth-common-CollectionCreateRequest"></a>

### CollectionCreateRequest
──────────────────-- Invoke ────────────────────────
CollectionCreate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |






<a name="auth-common-CollectionCreateResponse"></a>

### CollectionCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |






<a name="auth-common-CollectionGetHistoryRequest"></a>

### CollectionGetHistoryRequest
CollectionGetHistory


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| show_hidden | [bool](#bool) |  |  |






<a name="auth-common-CollectionGetHistoryResponse"></a>

### CollectionGetHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  | repeated auth.Collection collections = 1; |
| history | [auth.History](#auth-History) |  |  |






<a name="auth-common-CollectionGetListRequest"></a>

### CollectionGetListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |






<a name="auth-common-CollectionGetListResponse"></a>

### CollectionGetListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [auth.Collection](#auth-Collection) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="auth-common-CollectionGetRequest"></a>

### CollectionGetRequest
CollectionGet


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |






<a name="auth-common-CollectionGetResponse"></a>

### CollectionGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |






<a name="auth-common-CollectionHideTxRequest"></a>

### CollectionHideTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| hidden_tx | [auth.HiddenTx](#auth-HiddenTx) |  |  |






<a name="auth-common-CollectionHideTxResponse"></a>

### CollectionHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |






<a name="auth-common-CollectionUpdateRequest"></a>

### CollectionUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |






<a name="auth-common-CollectionUpdateResponse"></a>

### CollectionUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |












<a name="auth-common-CollectionService"></a>

### CollectionService
buf:lint:ignore RPC_NO_DELETE

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CollectionGet | [CollectionGetRequest](#auth-common-CollectionGetRequest) | [CollectionGetResponse](#auth-common-CollectionGetResponse) |  |
| CollectionGetList | [.google.protobuf.Empty](#google-protobuf-Empty) | [CollectionGetListResponse](#auth-common-CollectionGetListResponse) |  |
| CollectionGetHistory | [CollectionGetHistoryRequest](#auth-common-CollectionGetHistoryRequest) | [CollectionGetHistoryResponse](#auth-common-CollectionGetHistoryResponse) |  |
| CollectionCreate | [CollectionCreateRequest](#auth-common-CollectionCreateRequest) | [CollectionCreateResponse](#auth-common-CollectionCreateResponse) |  |
| CollectionUpdate | [CollectionUpdateRequest](#auth-common-CollectionUpdateRequest) | [CollectionUpdateResponse](#auth-common-CollectionUpdateResponse) |  |
| CollectionHideTx | [CollectionHideTxRequest](#auth-common-CollectionHideTxRequest) | [CollectionHideTxResponse](#auth-common-CollectionHideTxResponse) |  |





<a name="chaincode_auth_common_generic-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/common/generic.proto



<a name="auth-common-CreateRequest"></a>

### CreateRequest
──────────────────────────────-- Invoke ────────────────────────────────────-
Create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |
| reason | [string](#string) |  |  |






<a name="auth-common-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-HiddenTxRequest"></a>

### HiddenTxRequest
HiddenTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






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
| object | [auth.Object](#auth-Object) |  |  |
| hidden_tx | [auth.HiddenTx](#auth-HiddenTx) |  |  |






<a name="auth-common-HideTxResponse"></a>

### HideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |
| hidden_txs | [auth.HiddenTxList](#auth-HiddenTxList) |  |  |






<a name="auth-common-HistoryRequest"></a>

### HistoryRequest
History


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |






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
| object | [auth.Object](#auth-Object) |  |  |
| num_attrs | [int32](#int32) |  |  |






<a name="auth-common-ListByAttrsResponse"></a>

### ListByAttrsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| objects | [auth.Object](#auth-Object) | repeated |  |






<a name="auth-common-ListByCollectionRequest"></a>

### ListByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-ListByCollectionResponse"></a>

### ListByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| objects | [auth.Object](#auth-Object) | repeated |  |






<a name="auth-common-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| objects | [auth.Object](#auth-Object) | repeated |  |






<a name="auth-common-SuggestionApproveRequest"></a>

### SuggestionApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_key | [auth.ObjectKey](#auth-ObjectKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="auth-common-SuggestionApproveResponse"></a>

### SuggestionApproveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [auth.Suggestion](#auth-Suggestion) |  |  |
| object | [auth.Object](#auth-Object) |  |  |






<a name="auth-common-SuggestionByPartialKeyRequest"></a>

### SuggestionByPartialKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |
| num_attrs | [int32](#int32) |  |  |
| object_key | [auth.ObjectKey](#auth-ObjectKey) |  |  |
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
| object_key | [auth.ObjectKey](#auth-ObjectKey) |  |  |
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






<a name="auth-common-SuggestionListByObjectRequest"></a>

### SuggestionListByObjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object_key | [auth.ObjectKey](#auth-ObjectKey) |  |  |






<a name="auth-common-SuggestionListByObjectResponse"></a>

### SuggestionListByObjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestions | [auth.Suggestion](#auth-Suggestion) | repeated |  |






<a name="auth-common-SuggestionListRequest"></a>

### SuggestionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  | auth.Object object = 3; |






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
| object_key | [auth.ObjectKey](#auth-ObjectKey) |  |  |
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
| object | [auth.Object](#auth-Object) |  |  |
| tx_id | [string](#string) |  |  |






<a name="auth-common-UnHideTxResponse"></a>

### UnHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |
| hidden_txs | [auth.HiddenTxList](#auth-HiddenTxList) |  |  |






<a name="auth-common-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="auth-common-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| object | [auth.Object](#auth-Object) |  |  |












<a name="auth-common-GenericService"></a>

### GenericService
══════════════════════════════════ Object ═════════════════════════════════════

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#auth-common-GetRequest) | [GetResponse](#auth-common-GetResponse) |  |
| List | [ListRequest](#auth-common-ListRequest) | [ListResponse](#auth-common-ListResponse) |  |
| ListByCollection | [ListByCollectionRequest](#auth-common-ListByCollectionRequest) | [ListByCollectionResponse](#auth-common-ListByCollectionResponse) |  |
| ListByAttrs | [ListByAttrsRequest](#auth-common-ListByAttrsRequest) | [ListByAttrsResponse](#auth-common-ListByAttrsResponse) |  |
| History | [HistoryRequest](#auth-common-HistoryRequest) | [HistoryResponse](#auth-common-HistoryResponse) |  |
| HiddenTx | [HiddenTxRequest](#auth-common-HiddenTxRequest) | [HiddenTxResponse](#auth-common-HiddenTxResponse) |  |
| Create | [CreateRequest](#auth-common-CreateRequest) | [CreateResponse](#auth-common-CreateResponse) |  |
| Update | [UpdateRequest](#auth-common-UpdateRequest) | [UpdateResponse](#auth-common-UpdateResponse) |  |
| Delete | [DeleteRequest](#auth-common-DeleteRequest) | [DeleteResponse](#auth-common-DeleteResponse) |  |
| HideTx | [HideTxRequest](#auth-common-HideTxRequest) | [HideTxResponse](#auth-common-HideTxResponse) |  |
| UnHideTx | [UnHideTxRequest](#auth-common-UnHideTxRequest) | [UnHideTxResponse](#auth-common-UnHideTxResponse) |  |
| Suggestion | [SuggestionRequest](#auth-common-SuggestionRequest) | [SuggestionResponse](#auth-common-SuggestionResponse) | ══════════════════════════════== Suggestions ══════════════════════════════ ──────────────────────────────-- Query ──────────────────────────────────── |
| SuggestionList | [SuggestionListRequest](#auth-common-SuggestionListRequest) | [SuggestionListResponse](#auth-common-SuggestionListResponse) |  |
| SuggestionListByCollection | [SuggestionListByCollectionRequest](#auth-common-SuggestionListByCollectionRequest) | [SuggestionListByCollectionResponse](#auth-common-SuggestionListByCollectionResponse) |  |
| SuggestionByPartialKey | [SuggestionByPartialKeyRequest](#auth-common-SuggestionByPartialKeyRequest) | [SuggestionByPartialKeyResponse](#auth-common-SuggestionByPartialKeyResponse) |  |
| SuggestionCreate | [SuggestionCreateRequest](#auth-common-SuggestionCreateRequest) | [SuggestionCreateResponse](#auth-common-SuggestionCreateResponse) | ──────────────────────────────-- Invoke ─────────────────────────────────-- |
| SuggestionDelete | [SuggestionDeleteRequest](#auth-common-SuggestionDeleteRequest) | [SuggestionDeleteResponse](#auth-common-SuggestionDeleteResponse) |  |
| SuggestionApprove | [SuggestionApproveRequest](#auth-common-SuggestionApproveRequest) | [SuggestionApproveResponse](#auth-common-SuggestionApproveResponse) |  |





<a name="chaincode_auth_common_users-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/common/users.proto



<a name="auth-common-UserCreateRequest"></a>

### UserCreateRequest
UserCreate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="auth-common-UserCreateResponse"></a>

### UserCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |






<a name="auth-common-UserDeleteRequest"></a>

### UserDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="auth-common-UserDeleteResponse"></a>

### UserDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |






<a name="auth-common-UserGetCurrentIdResponse"></a>

### UserGetCurrentIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="auth-common-UserGetCurrentResponse"></a>

### UserGetCurrentResponse
──────────────────-- Query ─────────────────────


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |






<a name="auth-common-UserGetHiddenTxRequest"></a>

### UserGetHiddenTxRequest
UserGetHiddenTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="auth-common-UserGetHiddenTxResponse"></a>

### UserGetHiddenTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| hidden_txs | [auth.HiddenTxList](#auth-HiddenTxList) |  |  |






<a name="auth-common-UserGetHistoryRequest"></a>

### UserGetHistoryRequest
UserGetHistory


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| show_hidden | [bool](#bool) |  |  |






<a name="auth-common-UserGetHistoryResponse"></a>

### UserGetHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |
| history | [auth.History](#auth-History) |  |  |






<a name="auth-common-UserGetListRequest"></a>

### UserGetListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| limit | [uint32](#uint32) |  |  |






<a name="auth-common-UserGetListResponse"></a>

### UserGetListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [auth.User](#auth-User) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="auth-common-UserGetRequest"></a>

### UserGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="auth-common-UserGetResponse"></a>

### UserGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |






<a name="auth-common-UserHideTxRequest"></a>

### UserHideTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| hidden_tx | [auth.HiddenTx](#auth-HiddenTx) |  |  |






<a name="auth-common-UserHideTxResponse"></a>

### UserHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |






<a name="auth-common-UserUpdateRequest"></a>

### UserUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="auth-common-UserUpdateResponse"></a>

### UserUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [auth.User](#auth-User) |  |  |












<a name="auth-common-UserService"></a>

### UserService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UserGetCurrentId | [.google.protobuf.Empty](#google-protobuf-Empty) | [UserGetCurrentIdResponse](#auth-common-UserGetCurrentIdResponse) |  |
| UserGetCurrent | [.google.protobuf.Empty](#google-protobuf-Empty) | [UserGetCurrentResponse](#auth-common-UserGetCurrentResponse) |  |
| UserGet | [UserGetRequest](#auth-common-UserGetRequest) | [UserGetResponse](#auth-common-UserGetResponse) |  |
| UserGetList | [.google.protobuf.Empty](#google-protobuf-Empty) | [UserGetListResponse](#auth-common-UserGetListResponse) |  |
| UserGetHistory | [UserGetHistoryRequest](#auth-common-UserGetHistoryRequest) | [UserGetHistoryResponse](#auth-common-UserGetHistoryResponse) |  |
| UserGetHiddenTx | [UserGetHiddenTxRequest](#auth-common-UserGetHiddenTxRequest) | [UserGetHiddenTxResponse](#auth-common-UserGetHiddenTxResponse) |  |
| UserCreate | [UserCreateRequest](#auth-common-UserCreateRequest) | [UserCreateResponse](#auth-common-UserCreateResponse) |  |
| UserDelete | [UserDeleteRequest](#auth-common-UserDeleteRequest) | [UserDeleteResponse](#auth-common-UserDeleteResponse) |  |
| UserUpdate | [UserUpdateRequest](#auth-common-UserUpdateRequest) | [UserUpdateResponse](#auth-common-UserUpdateResponse) |  |
| UserHideTx | [UserHideTxRequest](#auth-common-UserHideTxRequest) | [UserHideTxResponse](#auth-common-UserHideTxResponse) |  |





<a name="chaincode_auth_rbac_schema_v1_rbac-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## chaincode/auth/rbac/schema/v1/rbac.proto



<a name="rbac-schema-v1-Membership"></a>

### Membership
══════════════════== Membership ══════════════════==
Used for Assigning Roles to Users
!! key := {MEMBERSHIP}{COLLECTION_ID}{ROLE_ID.Id}&lt;{USER_ID.msp_id}{USER_ID.id}&gt;
!! SecondaryKey := {MEMBERSHIP}&lt;{USER_ID.msp_id}{USER_ID.id}&gt;{COLLECTION_ID}{ROLE_ID.Id}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipCreateRequest"></a>

### MembershipCreateRequest
──────────────────-- Invoke ────────────────────────


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipCreateResponse"></a>

### MembershipCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| membership | [Membership](#rbac-schema-v1-Membership) |  |  |






<a name="rbac-schema-v1-MembershipDeleteRequest"></a>

### MembershipDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipDeleteResponse"></a>

### MembershipDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| membership | [Membership](#rbac-schema-v1-Membership) |  |  |






<a name="rbac-schema-v1-MembershipGetByCollectionRequest"></a>

### MembershipGetByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| bookmark | [string](#string) |  |  |
| page_size | [uint32](#uint32) |  |  |






<a name="rbac-schema-v1-MembershipGetByCollectionResponse"></a>

### MembershipGetByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |
| memberships | [Membership](#rbac-schema-v1-Membership) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipGetByUserRequest"></a>

### MembershipGetByUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| bookmark | [string](#string) |  |  |
| page_size | [uint32](#uint32) |  |  |






<a name="rbac-schema-v1-MembershipGetByUserResponse"></a>

### MembershipGetByUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| user | [auth.User](#auth-User) |  |  |
| memberships | [Membership](#rbac-schema-v1-Membership) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipGetHistoryRequest"></a>

### MembershipGetHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipGetHistoryResponse"></a>

### MembershipGetHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| history | [auth.History](#auth-History) |  |  |






<a name="rbac-schema-v1-MembershipGetListRequest"></a>

### MembershipGetListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| page_size | [uint32](#uint32) |  |  |






<a name="rbac-schema-v1-MembershipGetListResponse"></a>

### MembershipGetListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| memberships | [Membership](#rbac-schema-v1-Membership) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipGetRequest"></a>

### MembershipGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-MembershipGetResponse"></a>

### MembershipGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| membership | [Membership](#rbac-schema-v1-Membership) |  |  |






<a name="rbac-schema-v1-RoleCreateRequest"></a>

### RoleCreateRequest
──────────────────-- Invoke ────────────────────────


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-RoleCreateResponse"></a>

### RoleCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [auth.Role](#auth-Role) |  |  |






<a name="rbac-schema-v1-RoleDeleteRequest"></a>

### RoleDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-RoleDeleteResponse"></a>

### RoleDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [auth.Role](#auth-Role) |  |  |






<a name="rbac-schema-v1-RoleGetByCollectionRequest"></a>

### RoleGetByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| bookmark | [string](#string) |  |  |
| page_size | [uint32](#uint32) |  |  |






<a name="rbac-schema-v1-RoleGetByCollectionResponse"></a>

### RoleGetByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [auth.Collection](#auth-Collection) |  |  |
| roles | [auth.Role](#auth-Role) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="rbac-schema-v1-RoleGetHistoryRequest"></a>

### RoleGetHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |
| show_hidden | [bool](#bool) |  |  |






<a name="rbac-schema-v1-RoleGetHistoryResponse"></a>

### RoleGetHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| history | [auth.History](#auth-History) |  |  |






<a name="rbac-schema-v1-RoleGetListRequest"></a>

### RoleGetListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bookmark | [string](#string) |  |  |
| page_size | [uint32](#uint32) |  |  |






<a name="rbac-schema-v1-RoleGetListResponse"></a>

### RoleGetListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| roles | [auth.Role](#auth-Role) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="rbac-schema-v1-RoleGetRequest"></a>

### RoleGetRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |






<a name="rbac-schema-v1-RoleGetResponse"></a>

### RoleGetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [auth.Role](#auth-Role) |  |  |






<a name="rbac-schema-v1-RoleUpdateRequest"></a>

### RoleUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |
| update | [auth.Role](#auth-Role) |  |  |
| mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="rbac-schema-v1-RoleUpdateResponse"></a>

### RoleUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [auth.Role](#auth-Role) |  |  |












<a name="rbac-schema-v1-RBACService"></a>

### RBACService
buf:lint:ignore RPC_NO_DELETE

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| MembershipGetList | [MembershipGetListRequest](#rbac-schema-v1-MembershipGetListRequest) | [MembershipGetListResponse](#rbac-schema-v1-MembershipGetListResponse) |  |
| MembershipGet | [MembershipGetRequest](#rbac-schema-v1-MembershipGetRequest) | [MembershipGetResponse](#rbac-schema-v1-MembershipGetResponse) |  |
| MembershipGetByCollection | [MembershipGetByCollectionRequest](#rbac-schema-v1-MembershipGetByCollectionRequest) | [MembershipGetByCollectionResponse](#rbac-schema-v1-MembershipGetByCollectionResponse) |  |
| MembershipGetByUser | [MembershipGetByUserRequest](#rbac-schema-v1-MembershipGetByUserRequest) | [MembershipGetByUserResponse](#rbac-schema-v1-MembershipGetByUserResponse) |  |
| MembershipGetHistory | [MembershipGetHistoryRequest](#rbac-schema-v1-MembershipGetHistoryRequest) | [MembershipGetHistoryResponse](#rbac-schema-v1-MembershipGetHistoryResponse) |  |
| MembershipCreate | [MembershipCreateRequest](#rbac-schema-v1-MembershipCreateRequest) | [MembershipCreateResponse](#rbac-schema-v1-MembershipCreateResponse) |  |
| MembershipDelete | [MembershipDeleteRequest](#rbac-schema-v1-MembershipDeleteRequest) | [MembershipDeleteResponse](#rbac-schema-v1-MembershipDeleteResponse) |  |
| RoleGetList | [RoleGetListRequest](#rbac-schema-v1-RoleGetListRequest) | [RoleGetListResponse](#rbac-schema-v1-RoleGetListResponse) |  |
| RoleGet | [RoleGetRequest](#rbac-schema-v1-RoleGetRequest) | [RoleGetResponse](#rbac-schema-v1-RoleGetResponse) |  |
| RoleGetByCollection | [RoleGetByCollectionRequest](#rbac-schema-v1-RoleGetByCollectionRequest) | [RoleGetByCollectionResponse](#rbac-schema-v1-RoleGetByCollectionResponse) |  |
| RoleGetHistory | [RoleGetHistoryRequest](#rbac-schema-v1-RoleGetHistoryRequest) | [RoleGetHistoryResponse](#rbac-schema-v1-RoleGetHistoryResponse) |  |
| RoleCreate | [RoleCreateRequest](#rbac-schema-v1-RoleCreateRequest) | [RoleCreateResponse](#rbac-schema-v1-RoleCreateResponse) |  |
| RoleDelete | [RoleDeleteRequest](#rbac-schema-v1-RoleDeleteRequest) | [RoleDeleteResponse](#rbac-schema-v1-RoleDeleteResponse) |  |
| RoleUpdate | [RoleUpdateRequest](#rbac-schema-v1-RoleUpdateRequest) | [RoleUpdateResponse](#rbac-schema-v1-RoleUpdateResponse) |  |





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
