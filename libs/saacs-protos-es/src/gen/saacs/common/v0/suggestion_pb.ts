// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file saacs/common/v0/suggestion.proto (package saacs.common.v0, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
  BinaryReadOptions,
  FieldList,
  JsonReadOptions,
  JsonValue,
  PartialMessage,
  PlainMessage,
} from '@bufbuild/protobuf'
import { Any, FieldMask, Message, proto3 } from '@bufbuild/protobuf'
import { ItemKey } from './item_pb.js'

/**
 * Key should be
 * {auth.Suggestion}{COLLECTION_ID}{ITEM_TYPE}{...ITEM_ID}{SUGGESTION_ID}
 *
 * @generated from message saacs.common.v0.Suggestion
 */
export class Suggestion extends Message<Suggestion> {
  /**
   * @generated from field: saacs.common.v0.ItemKey primary_key = 1;
   */
  primaryKey?: ItemKey

  /**
   * @generated from field: string suggestion_id = 2;
   */
  suggestionId = ''

  /**
   * @generated from field: google.protobuf.FieldMask paths = 5;
   */
  paths?: FieldMask

  /**
   * @generated from field: google.protobuf.Any value = 6;
   */
  value?: Any

  constructor(data?: PartialMessage<Suggestion>) {
    super()
    proto3.util.initPartial(data, this)
  }

  static readonly runtime: typeof proto3 = proto3
  static readonly typeName = 'saacs.common.v0.Suggestion'
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: 'primary_key', kind: 'message', T: ItemKey },
    {
      no: 2,
      name: 'suggestion_id',
      kind: 'scalar',
      T: 9 /* ScalarType.STRING */,
    },
    { no: 5, name: 'paths', kind: 'message', T: FieldMask },
    { no: 6, name: 'value', kind: 'message', T: Any },
  ])

  static fromBinary(
    bytes: Uint8Array,
    options?: Partial<BinaryReadOptions>,
  ): Suggestion {
    return new Suggestion().fromBinary(bytes, options)
  }

  static fromJson(
    jsonValue: JsonValue,
    options?: Partial<JsonReadOptions>,
  ): Suggestion {
    return new Suggestion().fromJson(jsonValue, options)
  }

  static fromJsonString(
    jsonString: string,
    options?: Partial<JsonReadOptions>,
  ): Suggestion {
    return new Suggestion().fromJsonString(jsonString, options)
  }

  static equals(
    a: Suggestion | PlainMessage<Suggestion> | undefined,
    b: Suggestion | PlainMessage<Suggestion> | undefined,
  ): boolean {
    return proto3.util.equals(Suggestion, a, b)
  }
}
