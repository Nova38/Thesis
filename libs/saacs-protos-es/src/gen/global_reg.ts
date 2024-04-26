import type { IMessageTypeRegistry } from "@bufbuild/protobuf";
import { PathPolicy, Polices } from "./saacs/auth/v0/policy_pb.js";
import { Item, ItemKey, KeySchema, ReadWriteSet } from "./saacs/common/v0/item_pb.js";
import { Operation } from "./saacs/common/v0/operation_pb.js";
import { Attribute } from "./saacs/auth/v0/attribute_pb.js";
import { Collection } from "./saacs/auth/v0/collection_pb.js";
import { UserDirectMembership } from "./saacs/auth/v0/identity_pb.js";
import { Role, RoleIDList, UserCollectionRoles, UserGlobalRoles } from "./saacs/auth/v0/roles_pb.js";
import { AuthModel, Model, Model_Attribute, Model_GlobalRoles, Model_Identity, Model_Roles } from "./saacs/auth/v0/models_pb.js";
import { StateActivity } from "./saacs/common/v0/activity_pb.js";
import { HiddenOptions, HiddenTx, HiddenTxList, History, HistoryEntry, HistoryOptions } from "./saacs/common/v0/history_pb.js";
import { Date, Researcher, Specimen, Specimen_Georeference, Specimen_Grant, Specimen_Image, Specimen_Loan, Specimen_Primary, Specimen_Secondary, Specimen_Secondary_Preparation, Specimen_Taxon, SpecimenHistory, SpecimenHistoryEntry, SpecimenList, SpecimenMap, SpecimenUpdate } from "./saacs/biochain/v0/state_pb.js";
import { Suggestion } from "./saacs/common/v0/suggestion_pb.js";
import { FullItem, Pagination } from "./saacs/common/v0/packing_pb.js";
import { BatchCreateRequest, BatchCreateResponse, CreateRequest, CreateResponse, DeleteRequest, DeleteResponse, GetFullRequest, GetFullResponse, GetHiddenTxRequest, GetHiddenTxResponse, GetHistoryRequest, GetHistoryResponse, GetRequest, GetResponse, GetSuggestionRequest, GetSuggestionResponse, HideTxRequest, HideTxResponse, ListByAttrsRequest, ListByAttrsResponse, ListRequest, ListResponse, SuggestionApproveRequest, SuggestionApproveResponse, SuggestionByPartialKeyRequest, SuggestionByPartialKeyResponse, SuggestionCreateRequest, SuggestionCreateResponse, SuggestionDeleteRequest, SuggestionDeleteResponse, SuggestionListByCollectionRequest, SuggestionListByCollectionResponse, SuggestionListByItemRequest, SuggestionListByItemResponse, SuggestionListRequest, SuggestionListResponse, UnHideTxRequest, UnHideTxResponse, UpdateRequest, UpdateResponse } from "./saacs/chaincode/v0/chaincode_pb.js";
import { User } from "./saacs/common/v0/user_pb.js";
import { OperationsPerformed } from "./saacs/chaincode/v0/events_pb.js";
import { AuthorizeOperationRequest, AuthorizeOperationResponse, BootstrapRequest, BootstrapResponse, GetCollectionAuthModelRequest, GetCollectionsListRequest, GetCollectionsListResponse, GetCurrentFullUserResponse, GetCurrentUserRequest, GetCurrentUserResponse } from "./saacs/chaincode/v0/utils_pb.js";
import { ErrorWrapper } from "./saacs/common/v0/errors_pb.js";
import { Book } from "./saacs/example/v0/book_pb.js";
import { Group, SimpleItem } from "./saacs/example/v0/items_pb.js";
import { ItemWithNestedKey, Nested } from "./saacs/example/v0/nested_pb.js";

import { createRegistry } from "@bufbuild/protobuf";
export const GlobalRegistry: IMessageTypeRegistry = createRegistry(
  Polices,
  PathPolicy,
  ItemKey,
  Item,
  KeySchema,
  ReadWriteSet,
  Operation,
  Attribute,
  Collection,
  UserDirectMembership,
  Role,
  RoleIDList,
  UserCollectionRoles,
  UserGlobalRoles,
  Model,
  Model_Identity,
  Model_Roles,
  Model_GlobalRoles,
  Model_Attribute,
  AuthModel,
  StateActivity,
  HiddenTx,
  HiddenTxList,
  HistoryEntry,
  History,
  HistoryOptions,
  HiddenOptions,
  SpecimenHistory,
  SpecimenHistoryEntry,
  SpecimenUpdate,
  Date,
  Researcher,
  Specimen,
  Specimen_Primary,
  Specimen_Secondary,
  Specimen_Secondary_Preparation,
  Specimen_Taxon,
  Specimen_Georeference,
  Specimen_Image,
  Specimen_Loan,
  Specimen_Grant,
  SpecimenList,
  SpecimenMap,
  Suggestion,
  FullItem,
  Pagination,
  GetRequest,
  GetResponse,
  GetFullRequest,
  GetFullResponse,
  ListRequest,
  ListResponse,
  ListByAttrsRequest,
  ListByAttrsResponse,
  CreateRequest,
  CreateResponse,
  BatchCreateRequest,
  BatchCreateResponse,
  UpdateRequest,
  UpdateResponse,
  DeleteRequest,
  DeleteResponse,
  GetHistoryRequest,
  GetHistoryResponse,
  GetHiddenTxRequest,
  GetHiddenTxResponse,
  HideTxRequest,
  HideTxResponse,
  UnHideTxRequest,
  UnHideTxResponse,
  GetSuggestionRequest,
  GetSuggestionResponse,
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
  SuggestionApproveResponse,
  User,
  OperationsPerformed,
  GetCurrentUserRequest,
  GetCurrentUserResponse,
  GetCurrentFullUserResponse,
  AuthorizeOperationRequest,
  AuthorizeOperationResponse,
  GetCollectionsListRequest,
  GetCollectionsListResponse,
  BootstrapRequest,
  BootstrapResponse,
  GetCollectionAuthModelRequest,
  ErrorWrapper,
  Book,
  SimpleItem,
  Group,
  ItemWithNestedKey,
  Nested,
);
