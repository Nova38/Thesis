# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [saacs/common/v0/enums.proto](#saacs_common_v0_enums-proto)
    - [Action](#saacs-common-v0-Action)
    - [ItemKind](#saacs-common-v0-ItemKind)
    - [TransactionType](#saacs-common-v0-TransactionType)

- [saacs/auth/v0/policy.proto](#saacs_auth_v0_policy-proto)
    - [PathPolicy](#saacs-auth-v0-PathPolicy)
    - [PathPolicy.SubPathsEntry](#saacs-auth-v0-PathPolicy-SubPathsEntry)
    - [Polices](#saacs-auth-v0-Polices)
    - [Polices.ItemPoliciesEntry](#saacs-auth-v0-Polices-ItemPoliciesEntry)

- [saacs/auth/v0/type.proto](#saacs_auth_v0_type-proto)
    - [AuthType](#saacs-auth-v0-AuthType)

    - [File-level Extensions](#saacs_auth_v0_type-proto-extensions)

- [saacs/common/v0/item.proto](#saacs_common_v0_item-proto)
    - [Item](#saacs-common-v0-Item)
    - [ItemKey](#saacs-common-v0-ItemKey)
    - [KeySchema](#saacs-common-v0-KeySchema)
    - [ReadWriteSet](#saacs-common-v0-ReadWriteSet)

- [saacs/common/v0/operation.proto](#saacs_common_v0_operation-proto)
    - [Operation](#saacs-common-v0-Operation)

- [saacs/common/v0/options.proto](#saacs_common_v0_options-proto)
    - [File-level Extensions](#saacs_common_v0_options-proto-extensions)
    - [File-level Extensions](#saacs_common_v0_options-proto-extensions)
    - [File-level Extensions](#saacs_common_v0_options-proto-extensions)

- [saacs/auth/v0/attribute.proto](#saacs_auth_v0_attribute-proto)
    - [Attribute](#saacs-auth-v0-Attribute)

- [saacs/auth/v0/collection.proto](#saacs_auth_v0_collection-proto)
    - [Collection](#saacs-auth-v0-Collection)

- [saacs/auth/v0/identity.proto](#saacs_auth_v0_identity-proto)
    - [UserDirectMembership](#saacs-auth-v0-UserDirectMembership)

- [saacs/auth/v0/roles.proto](#saacs_auth_v0_roles-proto)
    - [Role](#saacs-auth-v0-Role)
    - [RoleIDList](#saacs-auth-v0-RoleIDList)
    - [UserCollectionRoles](#saacs-auth-v0-UserCollectionRoles)
    - [UserGlobalRoles](#saacs-auth-v0-UserGlobalRoles)
    - [UserGlobalRoles.RolesEntry](#saacs-auth-v0-UserGlobalRoles-RolesEntry)

- [saacs/auth/v0/models.proto](#saacs_auth_v0_models-proto)
    - [AuthModel](#saacs-auth-v0-AuthModel)
    - [Model](#saacs-auth-v0-Model)
    - [Model.Attribute](#saacs-auth-v0-Model-Attribute)
    - [Model.GlobalRoles](#saacs-auth-v0-Model-GlobalRoles)
    - [Model.Identity](#saacs-auth-v0-Model-Identity)
    - [Model.Roles](#saacs-auth-v0-Model-Roles)

- [saacs/common/v0/activity.proto](#saacs_common_v0_activity-proto)
    - [StateActivity](#saacs-common-v0-StateActivity)

- [saacs/common/v0/history.proto](#saacs_common_v0_history-proto)
    - [HiddenOptions](#saacs-common-v0-HiddenOptions)
    - [HiddenTx](#saacs-common-v0-HiddenTx)
    - [HiddenTxList](#saacs-common-v0-HiddenTxList)
    - [History](#saacs-common-v0-History)
    - [History.HiddenTxsByMspIdEntry](#saacs-common-v0-History-HiddenTxsByMspIdEntry)
    - [HistoryEntry](#saacs-common-v0-HistoryEntry)
    - [HistoryOptions](#saacs-common-v0-HistoryOptions)

- [saacs/biochain/v0/state.proto](#saacs_biochain_v0_state-proto)
    - [Date](#saacs-biochain-v0-Date)
    - [Researcher](#saacs-biochain-v0-Researcher)
    - [Specimen](#saacs-biochain-v0-Specimen)
    - [Specimen.Georeference](#saacs-biochain-v0-Specimen-Georeference)
    - [Specimen.Grant](#saacs-biochain-v0-Specimen-Grant)
    - [Specimen.GrantsEntry](#saacs-biochain-v0-Specimen-GrantsEntry)
    - [Specimen.Image](#saacs-biochain-v0-Specimen-Image)
    - [Specimen.ImagesEntry](#saacs-biochain-v0-Specimen-ImagesEntry)
    - [Specimen.Loan](#saacs-biochain-v0-Specimen-Loan)
    - [Specimen.LoansEntry](#saacs-biochain-v0-Specimen-LoansEntry)
    - [Specimen.Primary](#saacs-biochain-v0-Specimen-Primary)
    - [Specimen.Secondary](#saacs-biochain-v0-Specimen-Secondary)
    - [Specimen.Secondary.Preparation](#saacs-biochain-v0-Specimen-Secondary-Preparation)
    - [Specimen.Secondary.PreparationsEntry](#saacs-biochain-v0-Specimen-Secondary-PreparationsEntry)
    - [Specimen.Taxon](#saacs-biochain-v0-Specimen-Taxon)
    - [SpecimenHistory](#saacs-biochain-v0-SpecimenHistory)
    - [SpecimenHistoryEntry](#saacs-biochain-v0-SpecimenHistoryEntry)
    - [SpecimenList](#saacs-biochain-v0-SpecimenList)
    - [SpecimenMap](#saacs-biochain-v0-SpecimenMap)
    - [SpecimenMap.SpecimensEntry](#saacs-biochain-v0-SpecimenMap-SpecimensEntry)
    - [SpecimenUpdate](#saacs-biochain-v0-SpecimenUpdate)

    - [Specimen.Secondary.AGE](#saacs-biochain-v0-Specimen-Secondary-AGE)
    - [Specimen.Secondary.SEX](#saacs-biochain-v0-Specimen-Secondary-SEX)

- [saacs/common/v0/suggestion.proto](#saacs_common_v0_suggestion-proto)
    - [Suggestion](#saacs-common-v0-Suggestion)

- [saacs/common/v0/packing.proto](#saacs_common_v0_packing-proto)
    - [FullItem](#saacs-common-v0-FullItem)
    - [Pagination](#saacs-common-v0-Pagination)

- [saacs/chaincode/v0/chaincode.proto](#saacs_chaincode_v0_chaincode-proto)
    - [BatchCreateRequest](#saacs-chaincode-v0-BatchCreateRequest)
    - [BatchCreateResponse](#saacs-chaincode-v0-BatchCreateResponse)
    - [CreateRequest](#saacs-chaincode-v0-CreateRequest)
    - [CreateResponse](#saacs-chaincode-v0-CreateResponse)
    - [DeleteRequest](#saacs-chaincode-v0-DeleteRequest)
    - [DeleteResponse](#saacs-chaincode-v0-DeleteResponse)
    - [GetFullRequest](#saacs-chaincode-v0-GetFullRequest)
    - [GetFullResponse](#saacs-chaincode-v0-GetFullResponse)
    - [GetHiddenTxRequest](#saacs-chaincode-v0-GetHiddenTxRequest)
    - [GetHiddenTxResponse](#saacs-chaincode-v0-GetHiddenTxResponse)
    - [GetHistoryRequest](#saacs-chaincode-v0-GetHistoryRequest)
    - [GetHistoryResponse](#saacs-chaincode-v0-GetHistoryResponse)
    - [GetRequest](#saacs-chaincode-v0-GetRequest)
    - [GetResponse](#saacs-chaincode-v0-GetResponse)
    - [GetSuggestionRequest](#saacs-chaincode-v0-GetSuggestionRequest)
    - [GetSuggestionResponse](#saacs-chaincode-v0-GetSuggestionResponse)
    - [HideTxRequest](#saacs-chaincode-v0-HideTxRequest)
    - [HideTxResponse](#saacs-chaincode-v0-HideTxResponse)
    - [ListByAttrsRequest](#saacs-chaincode-v0-ListByAttrsRequest)
    - [ListByAttrsResponse](#saacs-chaincode-v0-ListByAttrsResponse)
    - [ListRequest](#saacs-chaincode-v0-ListRequest)
    - [ListResponse](#saacs-chaincode-v0-ListResponse)
    - [SuggestionApproveRequest](#saacs-chaincode-v0-SuggestionApproveRequest)
    - [SuggestionApproveResponse](#saacs-chaincode-v0-SuggestionApproveResponse)
    - [SuggestionByPartialKeyRequest](#saacs-chaincode-v0-SuggestionByPartialKeyRequest)
    - [SuggestionByPartialKeyResponse](#saacs-chaincode-v0-SuggestionByPartialKeyResponse)
    - [SuggestionCreateRequest](#saacs-chaincode-v0-SuggestionCreateRequest)
    - [SuggestionCreateResponse](#saacs-chaincode-v0-SuggestionCreateResponse)
    - [SuggestionDeleteRequest](#saacs-chaincode-v0-SuggestionDeleteRequest)
    - [SuggestionDeleteResponse](#saacs-chaincode-v0-SuggestionDeleteResponse)
    - [SuggestionListByCollectionRequest](#saacs-chaincode-v0-SuggestionListByCollectionRequest)
    - [SuggestionListByCollectionResponse](#saacs-chaincode-v0-SuggestionListByCollectionResponse)
    - [SuggestionListByItemRequest](#saacs-chaincode-v0-SuggestionListByItemRequest)
    - [SuggestionListByItemResponse](#saacs-chaincode-v0-SuggestionListByItemResponse)
    - [SuggestionListRequest](#saacs-chaincode-v0-SuggestionListRequest)
    - [SuggestionListResponse](#saacs-chaincode-v0-SuggestionListResponse)
    - [UnHideTxRequest](#saacs-chaincode-v0-UnHideTxRequest)
    - [UnHideTxResponse](#saacs-chaincode-v0-UnHideTxResponse)
    - [UpdateRequest](#saacs-chaincode-v0-UpdateRequest)
    - [UpdateResponse](#saacs-chaincode-v0-UpdateResponse)

    - [ItemService](#saacs-chaincode-v0-ItemService)

- [saacs/common/v0/user.proto](#saacs_common_v0_user-proto)
    - [User](#saacs-common-v0-User)

- [saacs/chaincode/v0/events.proto](#saacs_chaincode_v0_events-proto)
    - [OperationsPerformed](#saacs-chaincode-v0-OperationsPerformed)

- [saacs/chaincode/v0/utils.proto](#saacs_chaincode_v0_utils-proto)
    - [AuthorizeOperationRequest](#saacs-chaincode-v0-AuthorizeOperationRequest)
    - [AuthorizeOperationResponse](#saacs-chaincode-v0-AuthorizeOperationResponse)
    - [BootstrapRequest](#saacs-chaincode-v0-BootstrapRequest)
    - [BootstrapResponse](#saacs-chaincode-v0-BootstrapResponse)
    - [GetCollectionAuthModelRequest](#saacs-chaincode-v0-GetCollectionAuthModelRequest)
    - [GetCollectionsListRequest](#saacs-chaincode-v0-GetCollectionsListRequest)
    - [GetCollectionsListResponse](#saacs-chaincode-v0-GetCollectionsListResponse)
    - [GetCurrentFullUserResponse](#saacs-chaincode-v0-GetCurrentFullUserResponse)
    - [GetCurrentUserRequest](#saacs-chaincode-v0-GetCurrentUserRequest)
    - [GetCurrentUserResponse](#saacs-chaincode-v0-GetCurrentUserResponse)

    - [UtilsService](#saacs-chaincode-v0-UtilsService)

- [saacs/common/v0/errors.proto](#saacs_common_v0_errors-proto)
    - [ErrorWrapper](#saacs-common-v0-ErrorWrapper)

    - [TxError](#saacs-common-v0-TxError)

- [saacs/example/v0/book.proto](#saacs_example_v0_book-proto)
    - [Book](#saacs-sample-v0-Book)

- [saacs/example/v0/items.proto](#saacs_example_v0_items-proto)
    - [Group](#saacs-sample-v0-Group)
    - [SimpleItem](#saacs-sample-v0-SimpleItem)

- [saacs/example/v0/nested.proto](#saacs_example_v0_nested-proto)
    - [ItemWithNestedKey](#saacs-sample-v0-ItemWithNestedKey)
    - [Nested](#saacs-sample-v0-Nested)

- [Scalar Value Types](#scalar-value-types)



<a name="saacs_common_v0_enums-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/enums.proto





<a name="saacs-common-v0-Action"></a>

### Action


| Name | Number | Description |
| ---- | ------ | ----------- |
| ACTION_UNSPECIFIED | 0 | Should throw an error if used |
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
| ACTION_UNHIDE_TX | 21 |  |
| ACTION_HIDE_MSP_TX | 22 |  |
| ACTION_UNHIDE_MSP_TX | 23 |  |
| ACTION_VIEW_MSP_HIDDEN_TX | 24 |  |



<a name="saacs-common-v0-ItemKind"></a>

### ItemKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| ITEM_KIND_UNSPECIFIED | 0 |  |
| ITEM_KIND_PRIMARY_ITEM | 2 | ITEM_KIND_GLOBAL_ITEM = 1; Item&#39;s key := {COLLECTION_ID}{TYPE}[...key_paths] |
| ITEM_KIND_SUB_ITEM | 3 | Item&#39;s key := {COLLECTION_ID}{TYPE}&lt;PrimaryKey&gt;{...key_paths} |
| ITEM_KIND_REFERENCE | 4 |  |



<a name="saacs-common-v0-TransactionType"></a>

### TransactionType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSACTION_TYPE_UNSPECIFIED | 0 |  |
| TRANSACTION_TYPE_INVOKE | 1 |  |
| TRANSACTION_TYPE_QUERY | 2 |  |










<a name="saacs_auth_v0_policy-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/policy.proto



<a name="saacs-auth-v0-PathPolicy"></a>

### PathPolicy
This message is the tree node for operations on the state item


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | The path is a sub path of a field mask |
| full_path | [string](#string) |  |  |
| allow_sub_paths | [bool](#bool) |  |  |
| sub_paths | [PathPolicy.SubPathsEntry](#saacs-auth-v0-PathPolicy-SubPathsEntry) | repeated | The key is a valid sub path in the type of state item |
| actions | [saacs.common.v0.Action](#saacs-common-v0-Action) | repeated | If the policy is not set than use a parent policy unless nested policy is set |






<a name="saacs-auth-v0-PathPolicy-SubPathsEntry"></a>

### PathPolicy.SubPathsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathPolicy](#saacs-auth-v0-PathPolicy) |  |  |






<a name="saacs-auth-v0-Polices"></a>

### Polices



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_policies | [Polices.ItemPoliciesEntry](#saacs-auth-v0-Polices-ItemPoliciesEntry) | repeated | key is the item type |
| default_policy | [PathPolicy](#saacs-auth-v0-PathPolicy) |  | Default policy for all items |
| default_excluded_types | [string](#string) | repeated | The types that are excluded from the default policy |






<a name="saacs-auth-v0-Polices-ItemPoliciesEntry"></a>

### Polices.ItemPoliciesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [PathPolicy](#saacs-auth-v0-PathPolicy) |  |  |















<a name="saacs_auth_v0_type-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/type.proto





<a name="saacs-auth-v0-AuthType"></a>

### AuthType


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTH_TYPE_UNSPECIFIED | 0 | Invalid authentication type/Non Specified |
| AUTH_TYPE_NONE | 1 | No authentication, all users have access to all collections actions |
| AUTH_TYPE_IDENTITY | 2 | Identity-based authentication, Identities are stored per user per collection |
| AUTH_TYPE_ROLE | 3 | Role-based authentication, Roles are stored per user per collection |
| AUTH_TYPE_GLOBAL_ROLE | 4 | Global role-based authentication, Roles are stored per user |
| AUTH_TYPE_ATTRIBUTE | 5 | Attribute-based authentication, Attributes are stored per msp per collection |





<a name="saacs_auth_v0_type-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| auth_type | AuthType | .google.protobuf.MessageOptions | 55888 |  |







<a name="saacs_common_v0_item-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/item.proto



<a name="saacs-common-v0-Item"></a>

### Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [ItemKey](#saacs-common-v0-ItemKey) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |






<a name="saacs-common-v0-ItemKey"></a>

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
| item_kind | [ItemKind](#saacs-common-v0-ItemKind) |  |  |
| item_key_parts | [string](#string) | repeated |  |






<a name="saacs-common-v0-KeySchema"></a>

### KeySchema



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_type | [string](#string) |  | The item type of the key |
| item_kind | [ItemKind](#saacs-common-v0-ItemKind) |  | The kind of item that the key is for |
| properties | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  | The paths that make up the key |






<a name="saacs-common-v0-ReadWriteSet"></a>

### ReadWriteSet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| read_set | [Item](#saacs-common-v0-Item) | repeated |  |
| write_set | [Item](#saacs-common-v0-Item) | repeated |  |















<a name="saacs_common_v0_operation-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/operation.proto



<a name="saacs-common-v0-Operation"></a>

### Operation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action | [Action](#saacs-common-v0-Action) |  |  |
| collection_id | [string](#string) |  |  |
| item_type | [string](#string) |  |  |
| paths | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |















<a name="saacs_common_v0_options-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/options.proto







<a name="saacs_common_v0_options-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| key_schema | KeySchema | .google.protobuf.MessageOptions | 54599 |  |
| operation | Operation | .google.protobuf.MethodOptions | 57775 |  |
| transaction_type | TransactionType | .google.protobuf.MethodOptions | 50556 |  |







<a name="saacs_auth_v0_attribute-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/attribute.proto



<a name="saacs-auth-v0-Attribute"></a>

### Attribute
─────────────────────────────────────────────────────────────────────────────────
An attribute is used to define permissions via the value of the attribute in
the users certificate for a given msp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  | The msp of the organization that this attribute applies to |
| oid | [string](#string) |  | The oid of the attribute |
| value | [string](#string) |  | The value of the attribute required to be satisfied by the user to have the role |
| polices | [Polices](#saacs-auth-v0-Polices) |  | The Permission that the user will have if they have the attribute |















<a name="saacs_auth_v0_collection-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/collection.proto



<a name="saacs-auth-v0-Collection"></a>

### Collection
Collection
┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈
Note that the types of items are stored in the default ACLEntry

key := {COLLECTION}{COLLECTION_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | The key for the ledger |
| name | [string](#string) |  |  |
| auth_type | [AuthType](#saacs-auth-v0-AuthType) |  |  |
| item_types | [string](#string) | repeated |  |
| default | [Polices](#saacs-auth-v0-Polices) |  |  |
| use_auth_parents | [bool](#bool) |  |  |















<a name="saacs_auth_v0_identity-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/identity.proto



<a name="saacs-auth-v0-UserDirectMembership"></a>

### UserDirectMembership
Identity Auth Object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | The collection that the user is a member of |
| msp_id | [string](#string) |  | The msp of the organization that the user&#39;s certificate is from |
| user_id | [string](#string) |  | The id of the user from the certificate |
| polices | [Polices](#saacs-auth-v0-Polices) |  | The Permissions that the user will have |
| note | [string](#string) |  |  |















<a name="saacs_auth_v0_roles-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/roles.proto



<a name="saacs-auth-v0-Role"></a>

### Role
Shared Auth Object for Role Based Authentication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| role_id | [string](#string) |  |  |
| polices | [Polices](#saacs-auth-v0-Polices) |  |  |
| note | [string](#string) |  |  |
| parent_role_ids | [string](#string) | repeated |  |






<a name="saacs-auth-v0-RoleIDList"></a>

### RoleIDList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role_id | [string](#string) | repeated |  |






<a name="saacs-auth-v0-UserCollectionRoles"></a>

### UserCollectionRoles
Auth Object For RBAC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  | The collection that the user is a member of |
| msp_id | [string](#string) |  | The msp of the organization that the user&#39;s certificate is from |
| user_id | [string](#string) |  | The id of the user from the certificate |
| role_ids | [string](#string) | repeated | The roles that the user has in the collection |
| note | [string](#string) |  |  |






<a name="saacs-auth-v0-UserGlobalRoles"></a>

### UserGlobalRoles
Auth Object For Embedded RBAC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| msp_id | [string](#string) |  | The msp of the organization that the user&#39;s certificate is from |
| user_id | [string](#string) |  | The id of the user from the certificate |
| roles | [UserGlobalRoles.RolesEntry](#saacs-auth-v0-UserGlobalRoles-RolesEntry) | repeated | The roles that the user has in the collection key is the collection id value is the list of rolesIds |






<a name="saacs-auth-v0-UserGlobalRoles-RolesEntry"></a>

### UserGlobalRoles.RolesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [RoleIDList](#saacs-auth-v0-RoleIDList) |  |  |















<a name="saacs_auth_v0_models-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/auth/v0/models.proto



<a name="saacs-auth-v0-AuthModel"></a>

### AuthModel



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| identity | [Model.Identity](#saacs-auth-v0-Model-Identity) |  |  |
| roles | [Model.Roles](#saacs-auth-v0-Model-Roles) |  |  |
| global_roles | [Model.GlobalRoles](#saacs-auth-v0-Model-GlobalRoles) |  |  |
| attribute | [Model.Attribute](#saacs-auth-v0-Model-Attribute) |  |  |






<a name="saacs-auth-v0-Model"></a>

### Model







<a name="saacs-auth-v0-Model-Attribute"></a>

### Model.Attribute



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#saacs-auth-v0-Collection) |  |  |
| attribute | [Model.Attribute](#saacs-auth-v0-Model-Attribute) | repeated |  |






<a name="saacs-auth-v0-Model-GlobalRoles"></a>

### Model.GlobalRoles



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#saacs-auth-v0-Collection) |  |  |
| roles | [Role](#saacs-auth-v0-Role) | repeated |  |
| user_collection_roles | [UserGlobalRoles](#saacs-auth-v0-UserGlobalRoles) | repeated |  |






<a name="saacs-auth-v0-Model-Identity"></a>

### Model.Identity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#saacs-auth-v0-Collection) |  |  |
| user_direct_membership | [UserDirectMembership](#saacs-auth-v0-UserDirectMembership) | repeated |  |






<a name="saacs-auth-v0-Model-Roles"></a>

### Model.Roles



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [Collection](#saacs-auth-v0-Collection) |  |  |
| roles | [Role](#saacs-auth-v0-Role) | repeated |  |
| user_collection_roles | [UserCollectionRoles](#saacs-auth-v0-UserCollectionRoles) | repeated |  |















<a name="saacs_common_v0_activity-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/activity.proto



<a name="saacs-common-v0-StateActivity"></a>

### StateActivity
Can be used as the history entry


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  | The transaction id that caused the change |
| msp_id | [string](#string) |  | The msp of the user that caused the change |
| user_id | [string](#string) |  | The id of the user that caused the change |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp of the change |
| note | [string](#string) |  | A note about the change |
| object | [google.protobuf.Any](#google-protobuf-Any) |  | object |
| mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |















<a name="saacs_common_v0_history-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/history.proto



<a name="saacs-common-v0-HiddenOptions"></a>

### HiddenOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| include | [bool](#bool) |  |  |
| msp_ids | [string](#string) | repeated |  |






<a name="saacs-common-v0-HiddenTx"></a>

### HiddenTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  |  |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| note | [string](#string) |  |  |






<a name="saacs-common-v0-HiddenTxList"></a>

### HiddenTxList
Key should be {COLLECTION_ID}{auth.HiddenTxList}{?msp_id}{ITEM_TYPE}{...ITEM_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ItemKey](#saacs-common-v0-ItemKey) |  | The key that is used to store the item |
| msp_id | [string](#string) |  |  |
| txs | [HiddenTx](#saacs-common-v0-HiddenTx) | repeated | The list of hidden txs by tx_id |






<a name="saacs-common-v0-History"></a>

### History



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [HistoryEntry](#saacs-common-v0-HistoryEntry) | repeated |  |
| hidden_txs | [HiddenTxList](#saacs-common-v0-HiddenTxList) |  |  |
| hidden_txs_by_msp_id | [History.HiddenTxsByMspIdEntry](#saacs-common-v0-History-HiddenTxsByMspIdEntry) | repeated | The key is the msp_id of the group that is hiding the tx |






<a name="saacs-common-v0-History-HiddenTxsByMspIdEntry"></a>

### History.HiddenTxsByMspIdEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [HiddenTxList](#saacs-common-v0-HiddenTxList) |  |  |






<a name="saacs-common-v0-HistoryEntry"></a>

### HistoryEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  | The transaction id that caused the change |
| is_delete | [bool](#bool) |  | Whether the item was deleted |
| is_hidden | [bool](#bool) |  | Whether the transaction was hidden |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp of the change |
| note | [string](#string) |  | A note about the change |
| value | [google.protobuf.Any](#google-protobuf-Any) |  | The value of the item |






<a name="saacs-common-v0-HistoryOptions"></a>

### HistoryOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| include | [bool](#bool) |  |  |
| hidden | [HiddenOptions](#saacs-common-v0-HiddenOptions) |  |  |















<a name="saacs_biochain_v0_state-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/biochain/v0/state.proto



<a name="saacs-biochain-v0-Date"></a>

### Date
────────────────────────────────────────────────--
Specimen
────────────────────────────────────────────────--


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| verbatim | [string](#string) |  |  |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| year | [int32](#int32) |  |  |
| month | [string](#string) |  |  |
| day | [int32](#int32) |  |  |






<a name="saacs-biochain-v0-Researcher"></a>

### Researcher



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| first_name | [string](#string) |  |  |
| last_name | [string](#string) |  |  |
| middle_name | [string](#string) |  |  |






<a name="saacs-biochain-v0-Specimen"></a>

### Specimen



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| specimen_id | [string](#string) |  |  |
| primary | [Specimen.Primary](#saacs-biochain-v0-Specimen-Primary) |  |  |
| secondary | [Specimen.Secondary](#saacs-biochain-v0-Specimen-Secondary) |  |  |
| taxon | [Specimen.Taxon](#saacs-biochain-v0-Specimen-Taxon) |  |  |
| georeference | [Specimen.Georeference](#saacs-biochain-v0-Specimen-Georeference) |  |  |
| images | [Specimen.ImagesEntry](#saacs-biochain-v0-Specimen-ImagesEntry) | repeated |  |
| loans | [Specimen.LoansEntry](#saacs-biochain-v0-Specimen-LoansEntry) | repeated |  |
| grants | [Specimen.GrantsEntry](#saacs-biochain-v0-Specimen-GrantsEntry) | repeated |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-Georeference"></a>

### Specimen.Georeference



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| country | [string](#string) |  |  |
| state_province | [string](#string) |  |  |
| county | [string](#string) |  |  |
| locality | [string](#string) |  |  |
| latitude | [double](#double) |  |  |
| longitude | [double](#double) |  |  |
| habitat | [string](#string) |  |  |
| continent | [string](#string) |  | Georeference.continent -string 32 characters |
| location_remarks | [string](#string) |  | Georeference.locationRemarks -string 100-1k characters |
| coordinate_uncertainty_in_meters | [int32](#int32) |  | Georeference.coordinateUncercaintyInMeters -integer 7 digits |
| georeference_by | [string](#string) |  | Georeference.georeferenceBy -string 64 characters |
| georeference_date | [Date](#saacs-biochain-v0-Date) |  | Georeference.georeferenceDate -string MM/DD/YYYY |
| georeference_protocol | [string](#string) |  | Georeference.georeferenceProtocol -string 256 chars (weblink) |
| geodetic_datum | [string](#string) |  | Georeference.geodeticDatum -string 16 characters |
| footprint_wkt | [string](#string) |  | Georeference.footprintWKT -string 10-100k characters |
| notes | [string](#string) |  |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-Grant"></a>

### Specimen.Grant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| granted_by | [string](#string) |  |  |
| granted_to | [string](#string) |  |  |
| granted_date | [Date](#saacs-biochain-v0-Date) |  |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-GrantsEntry"></a>

### Specimen.GrantsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Grant](#saacs-biochain-v0-Specimen-Grant) |  |  |






<a name="saacs-biochain-v0-Specimen-Image"></a>

### Specimen.Image
Mapped Types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| url | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| hash | [string](#string) |  |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-ImagesEntry"></a>

### Specimen.ImagesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Image](#saacs-biochain-v0-Specimen-Image) |  |  |






<a name="saacs-biochain-v0-Specimen-Loan"></a>

### Specimen.Loan



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| description | [string](#string) |  |  |
| loaned_by | [string](#string) |  |  |
| loaned_to | [string](#string) |  |  |
| loaned_date | [Date](#saacs-biochain-v0-Date) |  |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-LoansEntry"></a>

### Specimen.LoansEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Loan](#saacs-biochain-v0-Specimen-Loan) |  |  |






<a name="saacs-biochain-v0-Specimen-Primary"></a>

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
| field_date | [Date](#saacs-biochain-v0-Date) |  |  |
| catalog_date | [Date](#saacs-biochain-v0-Date) |  |  |
| determined_date | [Date](#saacs-biochain-v0-Date) |  |  |
| determined_reason | [string](#string) |  |  |
| original_date | [Date](#saacs-biochain-v0-Date) |  |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-Secondary"></a>

### Specimen.Secondary



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sex | [Specimen.Secondary.SEX](#saacs-biochain-v0-Specimen-Secondary-SEX) |  |  |
| age | [Specimen.Secondary.AGE](#saacs-biochain-v0-Specimen-Secondary-AGE) |  |  |
| weight | [double](#double) |  |  |
| weight_units | [string](#string) |  |  |
| preparations | [Specimen.Secondary.PreparationsEntry](#saacs-biochain-v0-Specimen-Secondary-PreparationsEntry) | repeated |  |
| condition | [string](#string) |  |  |
| molt | [string](#string) |  |  |
| notes | [string](#string) |  |  |
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-Specimen-Secondary-Preparation"></a>

### Specimen.Secondary.Preparation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| verbatim | [string](#string) |  |  |






<a name="saacs-biochain-v0-Specimen-Secondary-PreparationsEntry"></a>

### Specimen.Secondary.PreparationsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen.Secondary.Preparation](#saacs-biochain-v0-Specimen-Secondary-Preparation) |  |  |






<a name="saacs-biochain-v0-Specimen-Taxon"></a>

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
| last_modified | [saacs.common.v0.StateActivity](#saacs-common-v0-StateActivity) |  |  |






<a name="saacs-biochain-v0-SpecimenHistory"></a>

### SpecimenHistory
option go_package =
&#34;github.com/nova38/saacs/gen/lib/biochain/ccbio/schema/v1;schemav1&#34;;


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [SpecimenHistoryEntry](#saacs-biochain-v0-SpecimenHistoryEntry) | repeated |  |
| hidden_txs | [saacs.common.v0.HiddenTxList](#saacs-common-v0-HiddenTxList) |  |  |






<a name="saacs-biochain-v0-SpecimenHistoryEntry"></a>

### SpecimenHistoryEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tx_id | [string](#string) |  | The transaction id that caused the change |
| is_delete | [bool](#bool) |  | Whether the item was deleted |
| is_hidden | [bool](#bool) |  | Whether the transaction was hidden |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The timestamp of the change |
| note | [string](#string) |  | A note about the change |
| value | [Specimen](#saacs-biochain-v0-Specimen) |  | The value of the item |






<a name="saacs-biochain-v0-SpecimenList"></a>

### SpecimenList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimens | [Specimen](#saacs-biochain-v0-Specimen) | repeated |  |






<a name="saacs-biochain-v0-SpecimenMap"></a>

### SpecimenMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimens | [SpecimenMap.SpecimensEntry](#saacs-biochain-v0-SpecimenMap-SpecimensEntry) | repeated |  |
| bookmark | [string](#string) |  |  |






<a name="saacs-biochain-v0-SpecimenMap-SpecimensEntry"></a>

### SpecimenMap.SpecimensEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Specimen](#saacs-biochain-v0-Specimen) |  |  |






<a name="saacs-biochain-v0-SpecimenUpdate"></a>

### SpecimenUpdate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specimen | [Specimen](#saacs-biochain-v0-Specimen) |  |  |
| mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |








<a name="saacs-biochain-v0-Specimen-Secondary-AGE"></a>

### Specimen.Secondary.AGE
Secondary.age -field with limited options: NEST, EMBRYO_EGG,
CHICK_SUBADULT, ADULT, UNKNOWN, CONTINGENT, blank

| Name | Number | Description |
| ---- | ------ | ----------- |
| AGE_UNDEFINED | 0 |  |
| AGE_UNKNOWN | 1 |  |
| AGE_NEST | 2 |  |
| AGE_EMBRYO_EGG | 3 |  |
| AGE_CHICK_SUBADULT | 4 |  |
| AGE_ADULT | 5 |  |
| AGE_CONTINGENT | 6 |  |



<a name="saacs-biochain-v0-Specimen-Secondary-SEX"></a>

### Specimen.Secondary.SEX


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEX_UNDEFINED | 0 |  |
| SEX_UNKNOWN | 1 |  |
| SEX_ATYPICAL | 2 |  |
| SEX_MALE | 3 |  |
| SEX_FEMALE | 4 |  |










<a name="saacs_common_v0_suggestion-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/suggestion.proto



<a name="saacs-common-v0-Suggestion"></a>

### Suggestion
Key should be
{auth.Suggestion}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUGGESTION_ID}


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| primary_key | [ItemKey](#saacs-common-v0-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| paths | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |
| value | [google.protobuf.Any](#google-protobuf-Any) |  |  |















<a name="saacs_common_v0_packing-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/packing.proto



<a name="saacs-common-v0-FullItem"></a>

### FullItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [Item](#saacs-common-v0-Item) |  |  |
| history | [History](#saacs-common-v0-History) |  |  |
| suggestions | [Suggestion](#saacs-common-v0-Suggestion) | repeated |  |






<a name="saacs-common-v0-Pagination"></a>

### Pagination



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_size | [int32](#int32) |  |  |
| bookmark | [string](#string) |  |  |















<a name="saacs_chaincode_v0_chaincode-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/chaincode/v0/chaincode.proto



<a name="saacs-chaincode-v0-BatchCreateRequest"></a>

### BatchCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [saacs.common.v0.Item](#saacs-common-v0-Item) | repeated |  |






<a name="saacs-chaincode-v0-BatchCreateResponse"></a>

### BatchCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [saacs.common.v0.Item](#saacs-common-v0-Item) | repeated |  |






<a name="saacs-chaincode-v0-CreateRequest"></a>

### CreateRequest
──────────────────────────────── Invoke
─────────────────────────────────────── Create


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |






<a name="saacs-chaincode-v0-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |






<a name="saacs-chaincode-v0-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| reason | [string](#string) |  |  |






<a name="saacs-chaincode-v0-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |






<a name="saacs-chaincode-v0-GetFullRequest"></a>

### GetFullRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  | saacs.common.v0.Item item = 3; |
| history_options | [saacs.common.v0.HistoryOptions](#saacs-common-v0-HistoryOptions) |  |  |
| include_suggestions | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-GetFullResponse"></a>

### GetFullResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |
| history | [saacs.common.v0.History](#saacs-common-v0-History) |  |  |
| suggestions | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) | repeated |  |






<a name="saacs-chaincode-v0-GetHiddenTxRequest"></a>

### GetHiddenTxRequest
GetHiddenTx


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |
| msp_ids | [string](#string) | repeated |  |






<a name="saacs-chaincode-v0-GetHiddenTxResponse"></a>

### GetHiddenTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| hidden_txs | [saacs.common.v0.HiddenTx](#saacs-common-v0-HiddenTx) | repeated |  |






<a name="saacs-chaincode-v0-GetHistoryRequest"></a>

### GetHistoryRequest
GetHistory


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| history_options | [saacs.common.v0.HiddenOptions](#saacs-common-v0-HiddenOptions) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-GetHistoryResponse"></a>

### GetHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| history | [saacs.common.v0.History](#saacs-common-v0-History) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-GetRequest"></a>

### GetRequest
Get


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  | saacs.common.v0.Item item = 3; |






<a name="saacs-chaincode-v0-GetResponse"></a>

### GetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |






<a name="saacs-chaincode-v0-GetSuggestionRequest"></a>

### GetSuggestionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |






<a name="saacs-chaincode-v0-GetSuggestionResponse"></a>

### GetSuggestionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) |  |  |






<a name="saacs-chaincode-v0-HideTxRequest"></a>

### HideTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| hidden_tx | [saacs.common.v0.HiddenTx](#saacs-common-v0-HiddenTx) |  |  |
| for_msp | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-HideTxResponse"></a>

### HideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| hidden_txs | [saacs.common.v0.HiddenTxList](#saacs-common-v0-HiddenTxList) |  |  |
| for_msp | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-ListByAttrsRequest"></a>

### ListByAttrsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| num_attrs | [int32](#int32) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-ListByAttrsResponse"></a>

### ListByAttrsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [saacs.common.v0.Item](#saacs-common-v0-Item) | repeated |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-ListRequest"></a>

### ListRequest
List of a type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [saacs.common.v0.Item](#saacs-common-v0-Item) | repeated |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-SuggestionApproveRequest"></a>

### SuggestionApproveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="saacs-chaincode-v0-SuggestionApproveResponse"></a>

### SuggestionApproveResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) |  |  |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |






<a name="saacs-chaincode-v0-SuggestionByPartialKeyRequest"></a>

### SuggestionByPartialKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| num_attrs | [int32](#int32) |  |  |
| suggestion_id | [string](#string) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-SuggestionByPartialKeyResponse"></a>

### SuggestionByPartialKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |
| suggestions | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) | repeated |  |






<a name="saacs-chaincode-v0-SuggestionCreateRequest"></a>

### SuggestionCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) |  |  |






<a name="saacs-chaincode-v0-SuggestionCreateResponse"></a>

### SuggestionCreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) |  |  |






<a name="saacs-chaincode-v0-SuggestionDeleteRequest"></a>

### SuggestionDeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| suggestion_id | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="saacs-chaincode-v0-SuggestionDeleteResponse"></a>

### SuggestionDeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestion | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) |  |  |






<a name="saacs-chaincode-v0-SuggestionListByCollectionRequest"></a>

### SuggestionListByCollectionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-SuggestionListByCollectionResponse"></a>

### SuggestionListByCollectionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestions | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) | repeated |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-SuggestionListByItemRequest"></a>

### SuggestionListByItemRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item_key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-SuggestionListByItemResponse"></a>

### SuggestionListByItemResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |
| suggestions | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) | repeated |  |






<a name="saacs-chaincode-v0-SuggestionListRequest"></a>

### SuggestionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  | saacs.common.v0.Item item = 3; |






<a name="saacs-chaincode-v0-SuggestionListResponse"></a>

### SuggestionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suggestions | [saacs.common.v0.Suggestion](#saacs-common-v0-Suggestion) | repeated |  |
| pagination | [saacs.common.v0.Pagination](#saacs-common-v0-Pagination) |  |  |






<a name="saacs-chaincode-v0-UnHideTxRequest"></a>

### UnHideTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| for_msp | [bool](#bool) |  |  |
| tx_id | [string](#string) |  |  |






<a name="saacs-chaincode-v0-UnHideTxResponse"></a>

### UnHideTxResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [saacs.common.v0.ItemKey](#saacs-common-v0-ItemKey) |  |  |
| hidden_txs | [saacs.common.v0.HiddenTxList](#saacs-common-v0-HiddenTxList) |  |  |
| for_msp | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |
| update_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) |  |  |






<a name="saacs-chaincode-v0-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| item | [saacs.common.v0.Item](#saacs-common-v0-Item) |  |  |












<a name="saacs-chaincode-v0-ItemService"></a>

### ItemService
rpc CreateCollection(CreateCollectionRequest) returns
(CreateCollectionResponse) {
  option (saacs.common.v0.transaction_type) = TRANSACTION_TYPE_INVOKE;
  option (saacs.common.v0.operation) = {action: ACTION_CREATE, item_type:
  &#34;Collection&#34;};
}
══════════════════════════════════ Item
═════════════════════════════════════

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Get | [GetRequest](#saacs-chaincode-v0-GetRequest) | [GetResponse](#saacs-chaincode-v0-GetResponse) |  |
| GetFull | [GetFullRequest](#saacs-chaincode-v0-GetFullRequest) | [GetFullResponse](#saacs-chaincode-v0-GetFullResponse) |  |
| List | [ListRequest](#saacs-chaincode-v0-ListRequest) | [ListResponse](#saacs-chaincode-v0-ListResponse) |  |
| ListByAttrs | [ListByAttrsRequest](#saacs-chaincode-v0-ListByAttrsRequest) | [ListByAttrsResponse](#saacs-chaincode-v0-ListByAttrsResponse) |  |
| Create | [CreateRequest](#saacs-chaincode-v0-CreateRequest) | [CreateResponse](#saacs-chaincode-v0-CreateResponse) |  |
| Update | [UpdateRequest](#saacs-chaincode-v0-UpdateRequest) | [UpdateResponse](#saacs-chaincode-v0-UpdateResponse) |  |
| Delete | [DeleteRequest](#saacs-chaincode-v0-DeleteRequest) | [DeleteResponse](#saacs-chaincode-v0-DeleteResponse) |  |
| GetHistory | [GetHistoryRequest](#saacs-chaincode-v0-GetHistoryRequest) | [GetHistoryResponse](#saacs-chaincode-v0-GetHistoryResponse) |  |
| GetHiddenTx | [GetHiddenTxRequest](#saacs-chaincode-v0-GetHiddenTxRequest) | [GetHiddenTxResponse](#saacs-chaincode-v0-GetHiddenTxResponse) |  |
| HideTx | [HideTxRequest](#saacs-chaincode-v0-HideTxRequest) | [HideTxResponse](#saacs-chaincode-v0-HideTxResponse) |  |
| UnHideTx | [UnHideTxRequest](#saacs-chaincode-v0-UnHideTxRequest) | [UnHideTxResponse](#saacs-chaincode-v0-UnHideTxResponse) |  |
| GetSuggestion | [GetSuggestionRequest](#saacs-chaincode-v0-GetSuggestionRequest) | [GetSuggestionResponse](#saacs-chaincode-v0-GetSuggestionResponse) |  |
| SuggestionListByCollection | [SuggestionListByCollectionRequest](#saacs-chaincode-v0-SuggestionListByCollectionRequest) | [SuggestionListByCollectionResponse](#saacs-chaincode-v0-SuggestionListByCollectionResponse) |  |
| SuggestionListByItem | [SuggestionListByItemRequest](#saacs-chaincode-v0-SuggestionListByItemRequest) | [SuggestionListByItemResponse](#saacs-chaincode-v0-SuggestionListByItemResponse) |  |
| SuggestionByPartialKey | [SuggestionByPartialKeyRequest](#saacs-chaincode-v0-SuggestionByPartialKeyRequest) | [SuggestionByPartialKeyResponse](#saacs-chaincode-v0-SuggestionByPartialKeyResponse) |  |
| SuggestionCreate | [SuggestionCreateRequest](#saacs-chaincode-v0-SuggestionCreateRequest) | [SuggestionCreateResponse](#saacs-chaincode-v0-SuggestionCreateResponse) | ──────────────────────────────── Invoke ─────────────────────────────────────── |
| SuggestionDelete | [SuggestionDeleteRequest](#saacs-chaincode-v0-SuggestionDeleteRequest) | [SuggestionDeleteResponse](#saacs-chaincode-v0-SuggestionDeleteResponse) |  |
| SuggestionApprove | [SuggestionApproveRequest](#saacs-chaincode-v0-SuggestionApproveRequest) | [SuggestionApproveResponse](#saacs-chaincode-v0-SuggestionApproveResponse) |  |





<a name="saacs_common_v0_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/user.proto



<a name="saacs-common-v0-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msp_id | [string](#string) |  |  |
| user_id | [string](#string) |  |  |















<a name="saacs_chaincode_v0_events-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/chaincode/v0/events.proto



<a name="saacs-chaincode-v0-OperationsPerformed"></a>

### OperationsPerformed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operations | [saacs.common.v0.Operation](#saacs-common-v0-Operation) | repeated |  |
| user | [saacs.common.v0.User](#saacs-common-v0-User) |  |  |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |















<a name="saacs_chaincode_v0_utils-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/chaincode/v0/utils.proto



<a name="saacs-chaincode-v0-AuthorizeOperationRequest"></a>

### AuthorizeOperationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation | [saacs.common.v0.Operation](#saacs-common-v0-Operation) |  |  |






<a name="saacs-chaincode-v0-AuthorizeOperationResponse"></a>

### AuthorizeOperationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorized | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-BootstrapRequest"></a>

### BootstrapRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [saacs.auth.v0.Collection](#saacs-auth-v0-Collection) |  |  |






<a name="saacs-chaincode-v0-BootstrapResponse"></a>

### BootstrapResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |
| collection | [saacs.auth.v0.Collection](#saacs-auth-v0-Collection) |  |  |






<a name="saacs-chaincode-v0-GetCollectionAuthModelRequest"></a>

### GetCollectionAuthModelRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |






<a name="saacs-chaincode-v0-GetCollectionsListRequest"></a>

### GetCollectionsListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| empty | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-GetCollectionsListResponse"></a>

### GetCollectionsListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [saacs.auth.v0.Collection](#saacs-auth-v0-Collection) | repeated |  |






<a name="saacs-chaincode-v0-GetCurrentFullUserResponse"></a>

### GetCurrentFullUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [saacs.common.v0.User](#saacs-common-v0-User) |  |  |
| registered | [bool](#bool) |  |  |
| user_collection_roles | [saacs.auth.v0.UserCollectionRoles](#saacs-auth-v0-UserCollectionRoles) | repeated |  |
| user_memberships | [saacs.auth.v0.UserDirectMembership](#saacs-auth-v0-UserDirectMembership) | repeated |  |






<a name="saacs-chaincode-v0-GetCurrentUserRequest"></a>

### GetCurrentUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| empty | [bool](#bool) |  |  |






<a name="saacs-chaincode-v0-GetCurrentUserResponse"></a>

### GetCurrentUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [saacs.common.v0.User](#saacs-common-v0-User) |  |  |
| registered | [bool](#bool) |  |  |












<a name="saacs-chaincode-v0-UtilsService"></a>

### UtilsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCurrentUser | [GetCurrentUserRequest](#saacs-chaincode-v0-GetCurrentUserRequest) | [GetCurrentUserResponse](#saacs-chaincode-v0-GetCurrentUserResponse) |  |
| Bootstrap | [BootstrapRequest](#saacs-chaincode-v0-BootstrapRequest) | [BootstrapResponse](#saacs-chaincode-v0-BootstrapResponse) |  |
| AuthorizeOperation | [AuthorizeOperationRequest](#saacs-chaincode-v0-AuthorizeOperationRequest) | [AuthorizeOperationResponse](#saacs-chaincode-v0-AuthorizeOperationResponse) |  |
| GetCollectionsList | [GetCollectionsListRequest](#saacs-chaincode-v0-GetCollectionsListRequest) | [GetCollectionsListResponse](#saacs-chaincode-v0-GetCollectionsListResponse) |  |





<a name="saacs_common_v0_errors-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/common/v0/errors.proto



<a name="saacs-common-v0-ErrorWrapper"></a>

### ErrorWrapper



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [TxError](#saacs-common-v0-TxError) |  |  |
| error | [google.protobuf.Struct](#google-protobuf-Struct) |  |  |
| message | [string](#string) |  |  |








<a name="saacs-common-v0-TxError"></a>

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










<a name="saacs_example_v0_book-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/example/v0/book.proto



<a name="saacs-sample-v0-Book"></a>

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















<a name="saacs_example_v0_items-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/example/v0/items.proto



<a name="saacs-sample-v0-Group"></a>

### Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| group_id | [string](#string) |  |  |
| item1 | [SimpleItem](#saacs-sample-v0-SimpleItem) |  |  |
| item2 | [SimpleItem](#saacs-sample-v0-SimpleItem) |  |  |






<a name="saacs-sample-v0-SimpleItem"></a>

### SimpleItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| quantity | [int32](#int32) |  |  |















<a name="saacs_example_v0_nested-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## saacs/example/v0/nested.proto



<a name="saacs-sample-v0-ItemWithNestedKey"></a>

### ItemWithNestedKey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection_id | [string](#string) |  |  |
| id | [string](#string) |  |  |
| nested | [Nested](#saacs-sample-v0-Nested) |  |  |
| value | [string](#string) |  |  |






<a name="saacs-sample-v0-Nested"></a>

### Nested



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| part1 | [string](#string) |  |  |
| part2 | [string](#string) |  |  |















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
