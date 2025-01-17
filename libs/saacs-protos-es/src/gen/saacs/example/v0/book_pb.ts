// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file saacs/example/v0/book.proto (package saacs.sample.v0, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message saacs.sample.v0.Book
 */
export class Book extends Message<Book> {
  /**
   * @generated from field: string collection_id = 1;
   */
  collectionId = "";

  /**
   * @generated from field: string isbn = 2;
   */
  isbn = "";

  /**
   * @generated from field: string book_title = 3;
   */
  bookTitle = "";

  /**
   * @generated from field: string author = 4;
   */
  author = "";

  /**
   * @generated from field: int32 year = 5;
   */
  year = 0;

  /**
   * @generated from field: string publisher = 6;
   */
  publisher = "";

  /**
   * @generated from field: string language = 7;
   */
  language = "";

  /**
   * @generated from field: string description = 8;
   */
  description = "";

  constructor(data?: PartialMessage<Book>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "saacs.sample.v0.Book";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "collection_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "isbn", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "book_title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "author", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "year", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "publisher", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "language", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Book {
    return new Book().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Book {
    return new Book().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Book {
    return new Book().fromJsonString(jsonString, options);
  }

  static equals(a: Book | PlainMessage<Book> | undefined, b: Book | PlainMessage<Book> | undefined): boolean {
    return proto3.util.equals(Book, a, b);
  }
}

