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

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/assert.js
function assert(condition, msg) {
  if (!condition) {
    throw new Error(msg);
  }
}
var FLOAT32_MAX = 34028234663852886e22;
var FLOAT32_MIN = -34028234663852886e22;
var UINT32_MAX = 4294967295;
var INT32_MAX = 2147483647;
var INT32_MIN = -2147483648;
function assertInt32(arg) {
  if (typeof arg !== "number")
    throw new Error("invalid int 32: " + typeof arg);
  if (!Number.isInteger(arg) || arg > INT32_MAX || arg < INT32_MIN)
    throw new Error("invalid int 32: " + arg);
}
function assertUInt32(arg) {
  if (typeof arg !== "number")
    throw new Error("invalid uint 32: " + typeof arg);
  if (!Number.isInteger(arg) || arg > UINT32_MAX || arg < 0)
    throw new Error("invalid uint 32: " + arg);
}
function assertFloat32(arg) {
  if (typeof arg !== "number")
    throw new Error("invalid float 32: " + typeof arg);
  if (!Number.isFinite(arg))
    return;
  if (arg > FLOAT32_MAX || arg < FLOAT32_MIN)
    throw new Error("invalid float 32: " + arg);
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/enum.js
var enumTypeSymbol = Symbol("@bufbuild/protobuf/enum-type");
function getEnumType(enumObject) {
  const t = enumObject[enumTypeSymbol];
  assert(t, "missing enum type on enum object");
  return t;
}
function setEnumType(enumObject, typeName, values, opt) {
  enumObject[enumTypeSymbol] = makeEnumType(typeName, values.map((v) => ({
    no: v.no,
    name: v.name,
    localName: enumObject[v.no]
  })), opt);
}
function makeEnumType(typeName, values, _opt) {
  const names = /* @__PURE__ */ Object.create(null);
  const numbers = /* @__PURE__ */ Object.create(null);
  const normalValues = [];
  for (const value of values) {
    const n = normalizeEnumValue(value);
    normalValues.push(n);
    names[value.name] = n;
    numbers[value.no] = n;
  }
  return {
    typeName,
    values: normalValues,
    // We do not surface options at this time
    // options: opt?.options ?? Object.create(null),
    findName(name) {
      return names[name];
    },
    findNumber(no) {
      return numbers[no];
    }
  };
}
function makeEnum(typeName, values, opt) {
  const enumObject = {};
  for (const value of values) {
    const n = normalizeEnumValue(value);
    enumObject[n.localName] = n.no;
    enumObject[n.no] = n.localName;
  }
  setEnumType(enumObject, typeName, values, opt);
  return enumObject;
}
function normalizeEnumValue(value) {
  if ("localName" in value) {
    return value;
  }
  return Object.assign(Object.assign({}, value), { localName: value.name });
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/message.js
var Message = class {
  /**
   * Compare with a message of the same type.
   */
  equals(other) {
    return this.getType().runtime.util.equals(this.getType(), this, other);
  }
  /**
   * Create a deep copy.
   */
  clone() {
    return this.getType().runtime.util.clone(this);
  }
  /**
   * Parse from binary data, merging fields.
   *
   * Repeated fields are appended. Map entries are added, overwriting
   * existing keys.
   *
   * If a message field is already present, it will be merged with the
   * new data.
   */
  fromBinary(bytes, options) {
    const type = this.getType(), format = type.runtime.bin, opt = format.makeReadOptions(options);
    format.readMessage(this, opt.readerFactory(bytes), bytes.byteLength, opt);
    return this;
  }
  /**
   * Parse a message from a JSON value.
   */
  fromJson(jsonValue, options) {
    const type = this.getType(), format = type.runtime.json, opt = format.makeReadOptions(options);
    format.readMessage(type, jsonValue, opt, this);
    return this;
  }
  /**
   * Parse a message from a JSON string.
   */
  fromJsonString(jsonString, options) {
    let json;
    try {
      json = JSON.parse(jsonString);
    } catch (e) {
      throw new Error(`cannot decode ${this.getType().typeName} from JSON: ${e instanceof Error ? e.message : String(e)}`);
    }
    return this.fromJson(json, options);
  }
  /**
   * Serialize the message to binary data.
   */
  toBinary(options) {
    const type = this.getType(), bin = type.runtime.bin, opt = bin.makeWriteOptions(options), writer = opt.writerFactory();
    bin.writeMessage(this, writer, opt);
    return writer.finish();
  }
  /**
   * Serialize the message to a JSON value, a JavaScript value that can be
   * passed to JSON.stringify().
   */
  toJson(options) {
    const type = this.getType(), json = type.runtime.json, opt = json.makeWriteOptions(options);
    return json.writeMessage(this, opt);
  }
  /**
   * Serialize the message to a JSON string.
   */
  toJsonString(options) {
    var _a;
    const value = this.toJson(options);
    return JSON.stringify(value, null, (_a = options === null || options === void 0 ? void 0 : options.prettySpaces) !== null && _a !== void 0 ? _a : 0);
  }
  /**
   * Override for serialization behavior. This will be invoked when calling
   * JSON.stringify on this message (i.e. JSON.stringify(msg)).
   *
   * Note that this will not serialize google.protobuf.Any with a packed
   * message because the protobuf JSON format specifies that it needs to be
   * unpacked, and this is only possible with a type registry to look up the
   * message type.  As a result, attempting to serialize a message with this
   * type will throw an Error.
   *
   * This method is protected because you should not need to invoke it
   * directly -- instead use JSON.stringify or toJsonString for
   * stringified JSON.  Alternatively, if actual JSON is desired, you should
   * use toJson.
   */
  toJSON() {
    return this.toJson({
      emitDefaultValues: true
    });
  }
  /**
   * Retrieve the MessageType of this message - a singleton that represents
   * the protobuf message declaration and provides metadata for reflection-
   * based operations.
   */
  getType() {
    return Object.getPrototypeOf(this).constructor;
  }
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/message-type.js
function makeMessageType(runtime, typeName, fields, opt) {
  var _a;
  const localName = (_a = opt === null || opt === void 0 ? void 0 : opt.localName) !== null && _a !== void 0 ? _a : typeName.substring(typeName.lastIndexOf(".") + 1);
  const type = {
    [localName]: function(data) {
      runtime.util.initFields(this);
      runtime.util.initPartial(data, this);
    }
  }[localName];
  Object.setPrototypeOf(type.prototype, new Message());
  Object.assign(type, {
    runtime,
    typeName,
    fields: runtime.util.newFieldList(fields),
    fromBinary(bytes, options) {
      return new type().fromBinary(bytes, options);
    },
    fromJson(jsonValue, options) {
      return new type().fromJson(jsonValue, options);
    },
    fromJsonString(jsonString, options) {
      return new type().fromJsonString(jsonString, options);
    },
    equals(a, b) {
      return runtime.util.equals(type, a, b);
    }
  });
  return type;
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/proto-runtime.js
function makeProtoRuntime(syntax, json, bin, util) {
  return {
    syntax,
    json,
    bin,
    util,
    makeMessageType(typeName, fields, opt) {
      return makeMessageType(this, typeName, fields, opt);
    },
    makeEnum,
    makeEnumType,
    getEnumType
  };
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/field.js
var ScalarType;
(function(ScalarType2) {
  ScalarType2[ScalarType2["DOUBLE"] = 1] = "DOUBLE";
  ScalarType2[ScalarType2["FLOAT"] = 2] = "FLOAT";
  ScalarType2[ScalarType2["INT64"] = 3] = "INT64";
  ScalarType2[ScalarType2["UINT64"] = 4] = "UINT64";
  ScalarType2[ScalarType2["INT32"] = 5] = "INT32";
  ScalarType2[ScalarType2["FIXED64"] = 6] = "FIXED64";
  ScalarType2[ScalarType2["FIXED32"] = 7] = "FIXED32";
  ScalarType2[ScalarType2["BOOL"] = 8] = "BOOL";
  ScalarType2[ScalarType2["STRING"] = 9] = "STRING";
  ScalarType2[ScalarType2["BYTES"] = 12] = "BYTES";
  ScalarType2[ScalarType2["UINT32"] = 13] = "UINT32";
  ScalarType2[ScalarType2["SFIXED32"] = 15] = "SFIXED32";
  ScalarType2[ScalarType2["SFIXED64"] = 16] = "SFIXED64";
  ScalarType2[ScalarType2["SINT32"] = 17] = "SINT32";
  ScalarType2[ScalarType2["SINT64"] = 18] = "SINT64";
})(ScalarType || (ScalarType = {}));
var LongType;
(function(LongType2) {
  LongType2[LongType2["BIGINT"] = 0] = "BIGINT";
  LongType2[LongType2["STRING"] = 1] = "STRING";
})(LongType || (LongType = {}));

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/google/varint.js
function varint64read() {
  let lowBits = 0;
  let highBits = 0;
  for (let shift = 0; shift < 28; shift += 7) {
    let b = this.buf[this.pos++];
    lowBits |= (b & 127) << shift;
    if ((b & 128) == 0) {
      this.assertBounds();
      return [lowBits, highBits];
    }
  }
  let middleByte = this.buf[this.pos++];
  lowBits |= (middleByte & 15) << 28;
  highBits = (middleByte & 112) >> 4;
  if ((middleByte & 128) == 0) {
    this.assertBounds();
    return [lowBits, highBits];
  }
  for (let shift = 3; shift <= 31; shift += 7) {
    let b = this.buf[this.pos++];
    highBits |= (b & 127) << shift;
    if ((b & 128) == 0) {
      this.assertBounds();
      return [lowBits, highBits];
    }
  }
  throw new Error("invalid varint");
}
function varint64write(lo, hi, bytes) {
  for (let i = 0; i < 28; i = i + 7) {
    const shift = lo >>> i;
    const hasNext = !(shift >>> 7 == 0 && hi == 0);
    const byte = (hasNext ? shift | 128 : shift) & 255;
    bytes.push(byte);
    if (!hasNext) {
      return;
    }
  }
  const splitBits = lo >>> 28 & 15 | (hi & 7) << 4;
  const hasMoreBits = !(hi >> 3 == 0);
  bytes.push((hasMoreBits ? splitBits | 128 : splitBits) & 255);
  if (!hasMoreBits) {
    return;
  }
  for (let i = 3; i < 31; i = i + 7) {
    const shift = hi >>> i;
    const hasNext = !(shift >>> 7 == 0);
    const byte = (hasNext ? shift | 128 : shift) & 255;
    bytes.push(byte);
    if (!hasNext) {
      return;
    }
  }
  bytes.push(hi >>> 31 & 1);
}
var TWO_PWR_32_DBL = 4294967296;
function int64FromString(dec) {
  const minus = dec[0] === "-";
  if (minus) {
    dec = dec.slice(1);
  }
  const base = 1e6;
  let lowBits = 0;
  let highBits = 0;
  function add1e6digit(begin, end) {
    const digit1e6 = Number(dec.slice(begin, end));
    highBits *= base;
    lowBits = lowBits * base + digit1e6;
    if (lowBits >= TWO_PWR_32_DBL) {
      highBits = highBits + (lowBits / TWO_PWR_32_DBL | 0);
      lowBits = lowBits % TWO_PWR_32_DBL;
    }
  }
  add1e6digit(-24, -18);
  add1e6digit(-18, -12);
  add1e6digit(-12, -6);
  add1e6digit(-6);
  return minus ? negate(lowBits, highBits) : newBits(lowBits, highBits);
}
function int64ToString(lo, hi) {
  let bits = newBits(lo, hi);
  const negative = bits.hi & 2147483648;
  if (negative) {
    bits = negate(bits.lo, bits.hi);
  }
  const result = uInt64ToString(bits.lo, bits.hi);
  return negative ? "-" + result : result;
}
function uInt64ToString(lo, hi) {
  ({ lo, hi } = toUnsigned(lo, hi));
  if (hi <= 2097151) {
    return String(TWO_PWR_32_DBL * hi + lo);
  }
  const low = lo & 16777215;
  const mid = (lo >>> 24 | hi << 8) & 16777215;
  const high = hi >> 16 & 65535;
  let digitA = low + mid * 6777216 + high * 6710656;
  let digitB = mid + high * 8147497;
  let digitC = high * 2;
  const base = 1e7;
  if (digitA >= base) {
    digitB += Math.floor(digitA / base);
    digitA %= base;
  }
  if (digitB >= base) {
    digitC += Math.floor(digitB / base);
    digitB %= base;
  }
  return digitC.toString() + decimalFrom1e7WithLeadingZeros(digitB) + decimalFrom1e7WithLeadingZeros(digitA);
}
function toUnsigned(lo, hi) {
  return { lo: lo >>> 0, hi: hi >>> 0 };
}
function newBits(lo, hi) {
  return { lo: lo | 0, hi: hi | 0 };
}
function negate(lowBits, highBits) {
  highBits = ~highBits;
  if (lowBits) {
    lowBits = ~lowBits + 1;
  } else {
    highBits += 1;
  }
  return newBits(lowBits, highBits);
}
var decimalFrom1e7WithLeadingZeros = (digit1e7) => {
  const partial = String(digit1e7);
  return "0000000".slice(partial.length) + partial;
};
function varint32write(value, bytes) {
  if (value >= 0) {
    while (value > 127) {
      bytes.push(value & 127 | 128);
      value = value >>> 7;
    }
    bytes.push(value);
  } else {
    for (let i = 0; i < 9; i++) {
      bytes.push(value & 127 | 128);
      value = value >> 7;
    }
    bytes.push(1);
  }
}
function varint32read() {
  let b = this.buf[this.pos++];
  let result = b & 127;
  if ((b & 128) == 0) {
    this.assertBounds();
    return result;
  }
  b = this.buf[this.pos++];
  result |= (b & 127) << 7;
  if ((b & 128) == 0) {
    this.assertBounds();
    return result;
  }
  b = this.buf[this.pos++];
  result |= (b & 127) << 14;
  if ((b & 128) == 0) {
    this.assertBounds();
    return result;
  }
  b = this.buf[this.pos++];
  result |= (b & 127) << 21;
  if ((b & 128) == 0) {
    this.assertBounds();
    return result;
  }
  b = this.buf[this.pos++];
  result |= (b & 15) << 28;
  for (let readBytes = 5; (b & 128) !== 0 && readBytes < 10; readBytes++)
    b = this.buf[this.pos++];
  if ((b & 128) != 0)
    throw new Error("invalid varint");
  this.assertBounds();
  return result >>> 0;
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/proto-int64.js
function makeInt64Support() {
  const dv = new DataView(new ArrayBuffer(8));
  const ok = typeof BigInt === "function" && typeof dv.getBigInt64 === "function" && typeof dv.getBigUint64 === "function" && typeof dv.setBigInt64 === "function" && typeof dv.setBigUint64 === "function" && (typeof process != "object" || typeof process.env != "object" || process.env.BUF_BIGINT_DISABLE !== "1");
  if (ok) {
    const MIN = BigInt("-9223372036854775808"), MAX = BigInt("9223372036854775807"), UMIN = BigInt("0"), UMAX = BigInt("18446744073709551615");
    return {
      zero: BigInt(0),
      supported: true,
      parse(value) {
        const bi = typeof value == "bigint" ? value : BigInt(value);
        if (bi > MAX || bi < MIN) {
          throw new Error(`int64 invalid: ${value}`);
        }
        return bi;
      },
      uParse(value) {
        const bi = typeof value == "bigint" ? value : BigInt(value);
        if (bi > UMAX || bi < UMIN) {
          throw new Error(`uint64 invalid: ${value}`);
        }
        return bi;
      },
      enc(value) {
        dv.setBigInt64(0, this.parse(value), true);
        return {
          lo: dv.getInt32(0, true),
          hi: dv.getInt32(4, true)
        };
      },
      uEnc(value) {
        dv.setBigInt64(0, this.uParse(value), true);
        return {
          lo: dv.getInt32(0, true),
          hi: dv.getInt32(4, true)
        };
      },
      dec(lo, hi) {
        dv.setInt32(0, lo, true);
        dv.setInt32(4, hi, true);
        return dv.getBigInt64(0, true);
      },
      uDec(lo, hi) {
        dv.setInt32(0, lo, true);
        dv.setInt32(4, hi, true);
        return dv.getBigUint64(0, true);
      }
    };
  }
  const assertInt64String = (value) => assert(/^-?[0-9]+$/.test(value), `int64 invalid: ${value}`);
  const assertUInt64String = (value) => assert(/^[0-9]+$/.test(value), `uint64 invalid: ${value}`);
  return {
    zero: "0",
    supported: false,
    parse(value) {
      if (typeof value != "string") {
        value = value.toString();
      }
      assertInt64String(value);
      return value;
    },
    uParse(value) {
      if (typeof value != "string") {
        value = value.toString();
      }
      assertUInt64String(value);
      return value;
    },
    enc(value) {
      if (typeof value != "string") {
        value = value.toString();
      }
      assertInt64String(value);
      return int64FromString(value);
    },
    uEnc(value) {
      if (typeof value != "string") {
        value = value.toString();
      }
      assertUInt64String(value);
      return int64FromString(value);
    },
    dec(lo, hi) {
      return int64ToString(lo, hi);
    },
    uDec(lo, hi) {
      return uInt64ToString(lo, hi);
    }
  };
}
var protoInt64 = makeInt64Support();

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/binary-encoding.js
var WireType;
(function(WireType2) {
  WireType2[WireType2["Varint"] = 0] = "Varint";
  WireType2[WireType2["Bit64"] = 1] = "Bit64";
  WireType2[WireType2["LengthDelimited"] = 2] = "LengthDelimited";
  WireType2[WireType2["StartGroup"] = 3] = "StartGroup";
  WireType2[WireType2["EndGroup"] = 4] = "EndGroup";
  WireType2[WireType2["Bit32"] = 5] = "Bit32";
})(WireType || (WireType = {}));
var BinaryWriter = class {
  constructor(textEncoder) {
    this.stack = [];
    this.textEncoder = textEncoder !== null && textEncoder !== void 0 ? textEncoder : new TextEncoder();
    this.chunks = [];
    this.buf = [];
  }
  /**
   * Return all bytes written and reset this writer.
   */
  finish() {
    this.chunks.push(new Uint8Array(this.buf));
    let len = 0;
    for (let i = 0; i < this.chunks.length; i++)
      len += this.chunks[i].length;
    let bytes = new Uint8Array(len);
    let offset = 0;
    for (let i = 0; i < this.chunks.length; i++) {
      bytes.set(this.chunks[i], offset);
      offset += this.chunks[i].length;
    }
    this.chunks = [];
    return bytes;
  }
  /**
   * Start a new fork for length-delimited data like a message
   * or a packed repeated field.
   *
   * Must be joined later with `join()`.
   */
  fork() {
    this.stack.push({ chunks: this.chunks, buf: this.buf });
    this.chunks = [];
    this.buf = [];
    return this;
  }
  /**
   * Join the last fork. Write its length and bytes, then
   * return to the previous state.
   */
  join() {
    let chunk = this.finish();
    let prev = this.stack.pop();
    if (!prev)
      throw new Error("invalid state, fork stack empty");
    this.chunks = prev.chunks;
    this.buf = prev.buf;
    this.uint32(chunk.byteLength);
    return this.raw(chunk);
  }
  /**
   * Writes a tag (field number and wire type).
   *
   * Equivalent to `uint32( (fieldNo << 3 | type) >>> 0 )`.
   *
   * Generated code should compute the tag ahead of time and call `uint32()`.
   */
  tag(fieldNo, type) {
    return this.uint32((fieldNo << 3 | type) >>> 0);
  }
  /**
   * Write a chunk of raw bytes.
   */
  raw(chunk) {
    if (this.buf.length) {
      this.chunks.push(new Uint8Array(this.buf));
      this.buf = [];
    }
    this.chunks.push(chunk);
    return this;
  }
  /**
   * Write a `uint32` value, an unsigned 32 bit varint.
   */
  uint32(value) {
    assertUInt32(value);
    while (value > 127) {
      this.buf.push(value & 127 | 128);
      value = value >>> 7;
    }
    this.buf.push(value);
    return this;
  }
  /**
   * Write a `int32` value, a signed 32 bit varint.
   */
  int32(value) {
    assertInt32(value);
    varint32write(value, this.buf);
    return this;
  }
  /**
   * Write a `bool` value, a variant.
   */
  bool(value) {
    this.buf.push(value ? 1 : 0);
    return this;
  }
  /**
   * Write a `bytes` value, length-delimited arbitrary data.
   */
  bytes(value) {
    this.uint32(value.byteLength);
    return this.raw(value);
  }
  /**
   * Write a `string` value, length-delimited data converted to UTF-8 text.
   */
  string(value) {
    let chunk = this.textEncoder.encode(value);
    this.uint32(chunk.byteLength);
    return this.raw(chunk);
  }
  /**
   * Write a `float` value, 32-bit floating point number.
   */
  float(value) {
    assertFloat32(value);
    let chunk = new Uint8Array(4);
    new DataView(chunk.buffer).setFloat32(0, value, true);
    return this.raw(chunk);
  }
  /**
   * Write a `double` value, a 64-bit floating point number.
   */
  double(value) {
    let chunk = new Uint8Array(8);
    new DataView(chunk.buffer).setFloat64(0, value, true);
    return this.raw(chunk);
  }
  /**
   * Write a `fixed32` value, an unsigned, fixed-length 32-bit integer.
   */
  fixed32(value) {
    assertUInt32(value);
    let chunk = new Uint8Array(4);
    new DataView(chunk.buffer).setUint32(0, value, true);
    return this.raw(chunk);
  }
  /**
   * Write a `sfixed32` value, a signed, fixed-length 32-bit integer.
   */
  sfixed32(value) {
    assertInt32(value);
    let chunk = new Uint8Array(4);
    new DataView(chunk.buffer).setInt32(0, value, true);
    return this.raw(chunk);
  }
  /**
   * Write a `sint32` value, a signed, zigzag-encoded 32-bit varint.
   */
  sint32(value) {
    assertInt32(value);
    value = (value << 1 ^ value >> 31) >>> 0;
    varint32write(value, this.buf);
    return this;
  }
  /**
   * Write a `fixed64` value, a signed, fixed-length 64-bit integer.
   */
  sfixed64(value) {
    let chunk = new Uint8Array(8), view = new DataView(chunk.buffer), tc = protoInt64.enc(value);
    view.setInt32(0, tc.lo, true);
    view.setInt32(4, tc.hi, true);
    return this.raw(chunk);
  }
  /**
   * Write a `fixed64` value, an unsigned, fixed-length 64 bit integer.
   */
  fixed64(value) {
    let chunk = new Uint8Array(8), view = new DataView(chunk.buffer), tc = protoInt64.uEnc(value);
    view.setInt32(0, tc.lo, true);
    view.setInt32(4, tc.hi, true);
    return this.raw(chunk);
  }
  /**
   * Write a `int64` value, a signed 64-bit varint.
   */
  int64(value) {
    let tc = protoInt64.enc(value);
    varint64write(tc.lo, tc.hi, this.buf);
    return this;
  }
  /**
   * Write a `sint64` value, a signed, zig-zag-encoded 64-bit varint.
   */
  sint64(value) {
    let tc = protoInt64.enc(value), sign = tc.hi >> 31, lo = tc.lo << 1 ^ sign, hi = (tc.hi << 1 | tc.lo >>> 31) ^ sign;
    varint64write(lo, hi, this.buf);
    return this;
  }
  /**
   * Write a `uint64` value, an unsigned 64-bit varint.
   */
  uint64(value) {
    let tc = protoInt64.uEnc(value);
    varint64write(tc.lo, tc.hi, this.buf);
    return this;
  }
};
var BinaryReader = class {
  constructor(buf, textDecoder) {
    this.varint64 = varint64read;
    this.uint32 = varint32read;
    this.buf = buf;
    this.len = buf.length;
    this.pos = 0;
    this.view = new DataView(buf.buffer, buf.byteOffset, buf.byteLength);
    this.textDecoder = textDecoder !== null && textDecoder !== void 0 ? textDecoder : new TextDecoder();
  }
  /**
   * Reads a tag - field number and wire type.
   */
  tag() {
    let tag = this.uint32(), fieldNo = tag >>> 3, wireType = tag & 7;
    if (fieldNo <= 0 || wireType < 0 || wireType > 5)
      throw new Error("illegal tag: field no " + fieldNo + " wire type " + wireType);
    return [fieldNo, wireType];
  }
  /**
   * Skip one element on the wire and return the skipped data.
   * Supports WireType.StartGroup since v2.0.0-alpha.23.
   */
  skip(wireType) {
    let start = this.pos;
    switch (wireType) {
      case WireType.Varint:
        while (this.buf[this.pos++] & 128) {
        }
        break;
      case WireType.Bit64:
        this.pos += 4;
      case WireType.Bit32:
        this.pos += 4;
        break;
      case WireType.LengthDelimited:
        let len = this.uint32();
        this.pos += len;
        break;
      case WireType.StartGroup:
        let t;
        while ((t = this.tag()[1]) !== WireType.EndGroup) {
          this.skip(t);
        }
        break;
      default:
        throw new Error("cant skip wire type " + wireType);
    }
    this.assertBounds();
    return this.buf.subarray(start, this.pos);
  }
  /**
   * Throws error if position in byte array is out of range.
   */
  assertBounds() {
    if (this.pos > this.len)
      throw new RangeError("premature EOF");
  }
  /**
   * Read a `int32` field, a signed 32 bit varint.
   */
  int32() {
    return this.uint32() | 0;
  }
  /**
   * Read a `sint32` field, a signed, zigzag-encoded 32-bit varint.
   */
  sint32() {
    let zze = this.uint32();
    return zze >>> 1 ^ -(zze & 1);
  }
  /**
   * Read a `int64` field, a signed 64-bit varint.
   */
  int64() {
    return protoInt64.dec(...this.varint64());
  }
  /**
   * Read a `uint64` field, an unsigned 64-bit varint.
   */
  uint64() {
    return protoInt64.uDec(...this.varint64());
  }
  /**
   * Read a `sint64` field, a signed, zig-zag-encoded 64-bit varint.
   */
  sint64() {
    let [lo, hi] = this.varint64();
    let s = -(lo & 1);
    lo = (lo >>> 1 | (hi & 1) << 31) ^ s;
    hi = hi >>> 1 ^ s;
    return protoInt64.dec(lo, hi);
  }
  /**
   * Read a `bool` field, a variant.
   */
  bool() {
    let [lo, hi] = this.varint64();
    return lo !== 0 || hi !== 0;
  }
  /**
   * Read a `fixed32` field, an unsigned, fixed-length 32-bit integer.
   */
  fixed32() {
    return this.view.getUint32((this.pos += 4) - 4, true);
  }
  /**
   * Read a `sfixed32` field, a signed, fixed-length 32-bit integer.
   */
  sfixed32() {
    return this.view.getInt32((this.pos += 4) - 4, true);
  }
  /**
   * Read a `fixed64` field, an unsigned, fixed-length 64 bit integer.
   */
  fixed64() {
    return protoInt64.uDec(this.sfixed32(), this.sfixed32());
  }
  /**
   * Read a `fixed64` field, a signed, fixed-length 64-bit integer.
   */
  sfixed64() {
    return protoInt64.dec(this.sfixed32(), this.sfixed32());
  }
  /**
   * Read a `float` field, 32-bit floating point number.
   */
  float() {
    return this.view.getFloat32((this.pos += 4) - 4, true);
  }
  /**
   * Read a `double` field, a 64-bit floating point number.
   */
  double() {
    return this.view.getFloat64((this.pos += 8) - 8, true);
  }
  /**
   * Read a `bytes` field, length-delimited arbitrary data.
   */
  bytes() {
    let len = this.uint32(), start = this.pos;
    this.pos += len;
    this.assertBounds();
    return this.buf.subarray(start, start + len);
  }
  /**
   * Read a `string` field, length-delimited data converted to UTF-8 text.
   */
  string() {
    return this.textDecoder.decode(this.bytes());
  }
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/field-wrapper.js
function wrapField(type, value) {
  if (value instanceof Message || !type.fieldWrapper) {
    return value;
  }
  return type.fieldWrapper.wrapField(value);
}
var wktWrapperToScalarType = {
  "google.protobuf.DoubleValue": ScalarType.DOUBLE,
  "google.protobuf.FloatValue": ScalarType.FLOAT,
  "google.protobuf.Int64Value": ScalarType.INT64,
  "google.protobuf.UInt64Value": ScalarType.UINT64,
  "google.protobuf.Int32Value": ScalarType.INT32,
  "google.protobuf.UInt32Value": ScalarType.UINT32,
  "google.protobuf.BoolValue": ScalarType.BOOL,
  "google.protobuf.StringValue": ScalarType.STRING,
  "google.protobuf.BytesValue": ScalarType.BYTES
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/scalars.js
function scalarEquals(type, a, b) {
  if (a === b) {
    return true;
  }
  if (type == ScalarType.BYTES) {
    if (!(a instanceof Uint8Array) || !(b instanceof Uint8Array)) {
      return false;
    }
    if (a.length !== b.length) {
      return false;
    }
    for (let i = 0; i < a.length; i++) {
      if (a[i] !== b[i]) {
        return false;
      }
    }
    return true;
  }
  switch (type) {
    case ScalarType.UINT64:
    case ScalarType.FIXED64:
    case ScalarType.INT64:
    case ScalarType.SFIXED64:
    case ScalarType.SINT64:
      return a == b;
  }
  return false;
}
function scalarDefaultValue(type, longType) {
  switch (type) {
    case ScalarType.BOOL:
      return false;
    case ScalarType.UINT64:
    case ScalarType.FIXED64:
    case ScalarType.INT64:
    case ScalarType.SFIXED64:
    case ScalarType.SINT64:
      return longType == 0 ? protoInt64.zero : "0";
    case ScalarType.DOUBLE:
    case ScalarType.FLOAT:
      return 0;
    case ScalarType.BYTES:
      return new Uint8Array(0);
    case ScalarType.STRING:
      return "";
    default:
      return 0;
  }
}
function scalarTypeInfo(type, value) {
  const isUndefined = value === void 0;
  let wireType = WireType.Varint;
  let isIntrinsicDefault = value === 0;
  switch (type) {
    case ScalarType.STRING:
      isIntrinsicDefault = isUndefined || !value.length;
      wireType = WireType.LengthDelimited;
      break;
    case ScalarType.BOOL:
      isIntrinsicDefault = value === false;
      break;
    case ScalarType.DOUBLE:
      wireType = WireType.Bit64;
      break;
    case ScalarType.FLOAT:
      wireType = WireType.Bit32;
      break;
    case ScalarType.INT64:
      isIntrinsicDefault = isUndefined || value == 0;
      break;
    case ScalarType.UINT64:
      isIntrinsicDefault = isUndefined || value == 0;
      break;
    case ScalarType.FIXED64:
      isIntrinsicDefault = isUndefined || value == 0;
      wireType = WireType.Bit64;
      break;
    case ScalarType.BYTES:
      isIntrinsicDefault = isUndefined || !value.byteLength;
      wireType = WireType.LengthDelimited;
      break;
    case ScalarType.FIXED32:
      wireType = WireType.Bit32;
      break;
    case ScalarType.SFIXED32:
      wireType = WireType.Bit32;
      break;
    case ScalarType.SFIXED64:
      isIntrinsicDefault = isUndefined || value == 0;
      wireType = WireType.Bit64;
      break;
    case ScalarType.SINT64:
      isIntrinsicDefault = isUndefined || value == 0;
      break;
  }
  const method = ScalarType[type].toLowerCase();
  return [wireType, method, isUndefined || isIntrinsicDefault];
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/binary-format-common.js
var unknownFieldsSymbol = Symbol("@bufbuild/protobuf/unknown-fields");
var readDefaults = {
  readUnknownFields: true,
  readerFactory: (bytes) => new BinaryReader(bytes)
};
var writeDefaults = {
  writeUnknownFields: true,
  writerFactory: () => new BinaryWriter()
};
function makeReadOptions(options) {
  return options ? Object.assign(Object.assign({}, readDefaults), options) : readDefaults;
}
function makeWriteOptions(options) {
  return options ? Object.assign(Object.assign({}, writeDefaults), options) : writeDefaults;
}
function makeBinaryFormatCommon() {
  return {
    makeReadOptions,
    makeWriteOptions,
    listUnknownFields(message) {
      var _a;
      return (_a = message[unknownFieldsSymbol]) !== null && _a !== void 0 ? _a : [];
    },
    discardUnknownFields(message) {
      delete message[unknownFieldsSymbol];
    },
    writeUnknownFields(message, writer) {
      const m = message;
      const c = m[unknownFieldsSymbol];
      if (c) {
        for (const f of c) {
          writer.tag(f.no, f.wireType).raw(f.data);
        }
      }
    },
    onUnknownField(message, no, wireType, data) {
      const m = message;
      if (!Array.isArray(m[unknownFieldsSymbol])) {
        m[unknownFieldsSymbol] = [];
      }
      m[unknownFieldsSymbol].push({ no, wireType, data });
    },
    readMessage(message, reader, length, options) {
      const type = message.getType();
      const end = length === void 0 ? reader.len : reader.pos + length;
      while (reader.pos < end) {
        const [fieldNo, wireType] = reader.tag(), field = type.fields.find(fieldNo);
        if (!field) {
          const data = reader.skip(wireType);
          if (options.readUnknownFields) {
            this.onUnknownField(message, fieldNo, wireType, data);
          }
          continue;
        }
        let target = message, repeated = field.repeated, localName = field.localName;
        if (field.oneof) {
          target = target[field.oneof.localName];
          if (target.case != localName) {
            delete target.value;
          }
          target.case = localName;
          localName = "value";
        }
        switch (field.kind) {
          case "scalar":
          case "enum":
            const scalarType = field.kind == "enum" ? ScalarType.INT32 : field.T;
            let read = readScalar;
            if (field.kind == "scalar" && field.L > 0) {
              read = readScalarLTString;
            }
            if (repeated) {
              let arr = target[localName];
              if (wireType == WireType.LengthDelimited && scalarType != ScalarType.STRING && scalarType != ScalarType.BYTES) {
                let e = reader.uint32() + reader.pos;
                while (reader.pos < e) {
                  arr.push(read(reader, scalarType));
                }
              } else {
                arr.push(read(reader, scalarType));
              }
            } else {
              target[localName] = read(reader, scalarType);
            }
            break;
          case "message":
            const messageType = field.T;
            if (repeated) {
              target[localName].push(readMessageField(reader, new messageType(), options));
            } else {
              if (target[localName] instanceof Message) {
                readMessageField(reader, target[localName], options);
              } else {
                target[localName] = readMessageField(reader, new messageType(), options);
                if (messageType.fieldWrapper && !field.oneof && !field.repeated) {
                  target[localName] = messageType.fieldWrapper.unwrapField(target[localName]);
                }
              }
            }
            break;
          case "map":
            let [mapKey, mapVal] = readMapEntry(field, reader, options);
            target[localName][mapKey] = mapVal;
            break;
        }
      }
    }
  };
}
function readMessageField(reader, message, options) {
  const format = message.getType().runtime.bin;
  format.readMessage(message, reader, reader.uint32(), options);
  return message;
}
function readMapEntry(field, reader, options) {
  const length = reader.uint32(), end = reader.pos + length;
  let key, val;
  while (reader.pos < end) {
    let [fieldNo] = reader.tag();
    switch (fieldNo) {
      case 1:
        key = readScalar(reader, field.K);
        break;
      case 2:
        switch (field.V.kind) {
          case "scalar":
            val = readScalar(reader, field.V.T);
            break;
          case "enum":
            val = reader.int32();
            break;
          case "message":
            val = readMessageField(reader, new field.V.T(), options);
            break;
        }
        break;
    }
  }
  if (key === void 0) {
    let keyRaw = scalarDefaultValue(field.K, LongType.BIGINT);
    key = field.K == ScalarType.BOOL ? keyRaw.toString() : keyRaw;
  }
  if (typeof key != "string" && typeof key != "number") {
    key = key.toString();
  }
  if (val === void 0) {
    switch (field.V.kind) {
      case "scalar":
        val = scalarDefaultValue(field.V.T, LongType.BIGINT);
        break;
      case "enum":
        val = 0;
        break;
      case "message":
        val = new field.V.T();
        break;
    }
  }
  return [key, val];
}
function readScalarLTString(reader, type) {
  const v = readScalar(reader, type);
  return typeof v == "bigint" ? v.toString() : v;
}
function readScalar(reader, type) {
  switch (type) {
    case ScalarType.STRING:
      return reader.string();
    case ScalarType.BOOL:
      return reader.bool();
    case ScalarType.DOUBLE:
      return reader.double();
    case ScalarType.FLOAT:
      return reader.float();
    case ScalarType.INT32:
      return reader.int32();
    case ScalarType.INT64:
      return reader.int64();
    case ScalarType.UINT64:
      return reader.uint64();
    case ScalarType.FIXED64:
      return reader.fixed64();
    case ScalarType.BYTES:
      return reader.bytes();
    case ScalarType.FIXED32:
      return reader.fixed32();
    case ScalarType.SFIXED32:
      return reader.sfixed32();
    case ScalarType.SFIXED64:
      return reader.sfixed64();
    case ScalarType.SINT64:
      return reader.sint64();
    case ScalarType.UINT32:
      return reader.uint32();
    case ScalarType.SINT32:
      return reader.sint32();
  }
}
function writeMapEntry(writer, options, field, key, value) {
  writer.tag(field.no, WireType.LengthDelimited);
  writer.fork();
  let keyValue = key;
  switch (field.K) {
    case ScalarType.INT32:
    case ScalarType.FIXED32:
    case ScalarType.UINT32:
    case ScalarType.SFIXED32:
    case ScalarType.SINT32:
      keyValue = Number.parseInt(key);
      break;
    case ScalarType.BOOL:
      assert(key == "true" || key == "false");
      keyValue = key == "true";
      break;
  }
  writeScalar(writer, field.K, 1, keyValue, true);
  switch (field.V.kind) {
    case "scalar":
      writeScalar(writer, field.V.T, 2, value, true);
      break;
    case "enum":
      writeScalar(writer, ScalarType.INT32, 2, value, true);
      break;
    case "message":
      writeMessageField(writer, options, field.V.T, 2, value);
      break;
  }
  writer.join();
}
function writeMessageField(writer, options, type, fieldNo, value) {
  if (value !== void 0) {
    const message = wrapField(type, value);
    writer.tag(fieldNo, WireType.LengthDelimited).bytes(message.toBinary(options));
  }
}
function writeScalar(writer, type, fieldNo, value, emitIntrinsicDefault) {
  let [wireType, method, isIntrinsicDefault] = scalarTypeInfo(type, value);
  if (!isIntrinsicDefault || emitIntrinsicDefault) {
    writer.tag(fieldNo, wireType)[method](value);
  }
}
function writePacked(writer, type, fieldNo, value) {
  if (!value.length) {
    return;
  }
  writer.tag(fieldNo, WireType.LengthDelimited).fork();
  let [, method] = scalarTypeInfo(type);
  for (let i = 0; i < value.length; i++) {
    writer[method](value[i]);
  }
  writer.join();
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/binary-format-proto3.js
function makeBinaryFormatProto3() {
  return Object.assign(Object.assign({}, makeBinaryFormatCommon()), { writeMessage(message, writer, options) {
    const type = message.getType();
    for (const field of type.fields.byNumber()) {
      let value, repeated = field.repeated, localName = field.localName;
      if (field.oneof) {
        const oneof = message[field.oneof.localName];
        if (oneof.case !== localName) {
          continue;
        }
        value = oneof.value;
      } else {
        value = message[localName];
      }
      switch (field.kind) {
        case "scalar":
        case "enum":
          let scalarType = field.kind == "enum" ? ScalarType.INT32 : field.T;
          if (repeated) {
            if (field.packed) {
              writePacked(writer, scalarType, field.no, value);
            } else {
              for (const item of value) {
                writeScalar(writer, scalarType, field.no, item, true);
              }
            }
          } else {
            if (value !== void 0) {
              writeScalar(writer, scalarType, field.no, value, !!field.oneof || field.opt);
            }
          }
          break;
        case "message":
          if (repeated) {
            for (const item of value) {
              writeMessageField(writer, options, field.T, field.no, item);
            }
          } else {
            writeMessageField(writer, options, field.T, field.no, value);
          }
          break;
        case "map":
          for (const [key, val] of Object.entries(value)) {
            writeMapEntry(writer, options, field, key, val);
          }
          break;
      }
    }
    if (options.writeUnknownFields) {
      this.writeUnknownFields(message, writer);
    }
    return writer;
  } });
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/proto-base64.js
var encTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/".split("");
var decTable = [];
for (let i = 0; i < encTable.length; i++)
  decTable[encTable[i].charCodeAt(0)] = i;
decTable["-".charCodeAt(0)] = encTable.indexOf("+");
decTable["_".charCodeAt(0)] = encTable.indexOf("/");
var protoBase64 = {
  /**
   * Decodes a base64 string to a byte array.
   *
   * - ignores white-space, including line breaks and tabs
   * - allows inner padding (can decode concatenated base64 strings)
   * - does not require padding
   * - understands base64url encoding:
   *   "-" instead of "+",
   *   "_" instead of "/",
   *   no padding
   */
  dec(base64Str) {
    let es = base64Str.length * 3 / 4;
    if (base64Str[base64Str.length - 2] == "=")
      es -= 2;
    else if (base64Str[base64Str.length - 1] == "=")
      es -= 1;
    let bytes = new Uint8Array(es), bytePos = 0, groupPos = 0, b, p = 0;
    for (let i = 0; i < base64Str.length; i++) {
      b = decTable[base64Str.charCodeAt(i)];
      if (b === void 0) {
        switch (base64Str[i]) {
          case "=":
            groupPos = 0;
          case "\n":
          case "\r":
          case "	":
          case " ":
            continue;
          default:
            throw Error("invalid base64 string.");
        }
      }
      switch (groupPos) {
        case 0:
          p = b;
          groupPos = 1;
          break;
        case 1:
          bytes[bytePos++] = p << 2 | (b & 48) >> 4;
          p = b;
          groupPos = 2;
          break;
        case 2:
          bytes[bytePos++] = (p & 15) << 4 | (b & 60) >> 2;
          p = b;
          groupPos = 3;
          break;
        case 3:
          bytes[bytePos++] = (p & 3) << 6 | b;
          groupPos = 0;
          break;
      }
    }
    if (groupPos == 1)
      throw Error("invalid base64 string.");
    return bytes.subarray(0, bytePos);
  },
  /**
   * Encode a byte array to a base64 string.
   */
  enc(bytes) {
    let base64 = "", groupPos = 0, b, p = 0;
    for (let i = 0; i < bytes.length; i++) {
      b = bytes[i];
      switch (groupPos) {
        case 0:
          base64 += encTable[b >> 2];
          p = (b & 3) << 4;
          groupPos = 1;
          break;
        case 1:
          base64 += encTable[p | b >> 4];
          p = (b & 15) << 2;
          groupPos = 2;
          break;
        case 2:
          base64 += encTable[p | b >> 6];
          base64 += encTable[b & 63];
          groupPos = 0;
          break;
      }
    }
    if (groupPos) {
      base64 += encTable[p];
      base64 += "=";
      if (groupPos == 1)
        base64 += "=";
    }
    return base64;
  }
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/json-format-common.js
var jsonReadDefaults = {
  ignoreUnknownFields: false
};
var jsonWriteDefaults = {
  emitDefaultValues: false,
  enumAsInteger: false,
  useProtoFieldName: false,
  prettySpaces: 0
};
function makeReadOptions2(options) {
  return options ? Object.assign(Object.assign({}, jsonReadDefaults), options) : jsonReadDefaults;
}
function makeWriteOptions2(options) {
  return options ? Object.assign(Object.assign({}, jsonWriteDefaults), options) : jsonWriteDefaults;
}
function makeJsonFormatCommon(makeWriteField) {
  const writeField = makeWriteField(writeEnum, writeScalar2);
  return {
    makeReadOptions: makeReadOptions2,
    makeWriteOptions: makeWriteOptions2,
    readMessage(type, json, options, message) {
      if (json == null || Array.isArray(json) || typeof json != "object") {
        throw new Error(`cannot decode message ${type.typeName} from JSON: ${this.debug(json)}`);
      }
      message = message !== null && message !== void 0 ? message : new type();
      const oneofSeen = {};
      for (const [jsonKey, jsonValue] of Object.entries(json)) {
        const field = type.fields.findJsonName(jsonKey);
        if (!field) {
          if (!options.ignoreUnknownFields) {
            throw new Error(`cannot decode message ${type.typeName} from JSON: key "${jsonKey}" is unknown`);
          }
          continue;
        }
        let localName = field.localName;
        let target = message;
        if (field.oneof) {
          if (jsonValue === null && field.kind == "scalar") {
            continue;
          }
          const seen = oneofSeen[field.oneof.localName];
          if (seen) {
            throw new Error(`cannot decode message ${type.typeName} from JSON: multiple keys for oneof "${field.oneof.name}" present: "${seen}", "${jsonKey}"`);
          }
          oneofSeen[field.oneof.localName] = jsonKey;
          target = target[field.oneof.localName] = { case: localName };
          localName = "value";
        }
        if (field.repeated) {
          if (jsonValue === null) {
            continue;
          }
          if (!Array.isArray(jsonValue)) {
            throw new Error(`cannot decode field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonValue)}`);
          }
          const targetArray = target[localName];
          for (const jsonItem of jsonValue) {
            if (jsonItem === null) {
              throw new Error(`cannot decode field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonItem)}`);
            }
            let val;
            switch (field.kind) {
              case "message":
                val = field.T.fromJson(jsonItem, options);
                break;
              case "enum":
                val = readEnum(field.T, jsonItem, options.ignoreUnknownFields);
                if (val === void 0)
                  continue;
                break;
              case "scalar":
                try {
                  val = readScalar2(field.T, jsonItem, field.L);
                } catch (e) {
                  let m = `cannot decode field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonItem)}`;
                  if (e instanceof Error && e.message.length > 0) {
                    m += `: ${e.message}`;
                  }
                  throw new Error(m);
                }
                break;
            }
            targetArray.push(val);
          }
        } else if (field.kind == "map") {
          if (jsonValue === null) {
            continue;
          }
          if (Array.isArray(jsonValue) || typeof jsonValue != "object") {
            throw new Error(`cannot decode field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonValue)}`);
          }
          const targetMap = target[localName];
          for (const [jsonMapKey, jsonMapValue] of Object.entries(jsonValue)) {
            if (jsonMapValue === null) {
              throw new Error(`cannot decode field ${type.typeName}.${field.name} from JSON: map value null`);
            }
            let val;
            switch (field.V.kind) {
              case "message":
                val = field.V.T.fromJson(jsonMapValue, options);
                break;
              case "enum":
                val = readEnum(field.V.T, jsonMapValue, options.ignoreUnknownFields);
                if (val === void 0)
                  continue;
                break;
              case "scalar":
                try {
                  val = readScalar2(field.V.T, jsonMapValue, LongType.BIGINT);
                } catch (e) {
                  let m = `cannot decode map value for field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonValue)}`;
                  if (e instanceof Error && e.message.length > 0) {
                    m += `: ${e.message}`;
                  }
                  throw new Error(m);
                }
                break;
            }
            try {
              targetMap[readScalar2(field.K, field.K == ScalarType.BOOL ? jsonMapKey == "true" ? true : jsonMapKey == "false" ? false : jsonMapKey : jsonMapKey, LongType.BIGINT).toString()] = val;
            } catch (e) {
              let m = `cannot decode map key for field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonValue)}`;
              if (e instanceof Error && e.message.length > 0) {
                m += `: ${e.message}`;
              }
              throw new Error(m);
            }
          }
        } else {
          switch (field.kind) {
            case "message":
              const messageType = field.T;
              if (jsonValue === null && messageType.typeName != "google.protobuf.Value") {
                if (field.oneof) {
                  throw new Error(`cannot decode field ${type.typeName}.${field.name} from JSON: null is invalid for oneof field "${jsonKey}"`);
                }
                continue;
              }
              if (target[localName] instanceof Message) {
                target[localName].fromJson(jsonValue, options);
              } else {
                target[localName] = messageType.fromJson(jsonValue, options);
                if (messageType.fieldWrapper && !field.oneof) {
                  target[localName] = messageType.fieldWrapper.unwrapField(target[localName]);
                }
              }
              break;
            case "enum":
              const enumValue = readEnum(field.T, jsonValue, options.ignoreUnknownFields);
              if (enumValue !== void 0) {
                target[localName] = enumValue;
              }
              break;
            case "scalar":
              try {
                target[localName] = readScalar2(field.T, jsonValue, field.L);
              } catch (e) {
                let m = `cannot decode field ${type.typeName}.${field.name} from JSON: ${this.debug(jsonValue)}`;
                if (e instanceof Error && e.message.length > 0) {
                  m += `: ${e.message}`;
                }
                throw new Error(m);
              }
              break;
          }
        }
      }
      return message;
    },
    writeMessage(message, options) {
      const type = message.getType();
      const json = {};
      let field;
      try {
        for (const member of type.fields.byMember()) {
          let jsonValue;
          if (member.kind == "oneof") {
            const oneof = message[member.localName];
            if (oneof.value === void 0) {
              continue;
            }
            field = member.findField(oneof.case);
            if (!field) {
              throw "oneof case not found: " + oneof.case;
            }
            jsonValue = writeField(field, oneof.value, options);
          } else {
            field = member;
            jsonValue = writeField(field, message[field.localName], options);
          }
          if (jsonValue !== void 0) {
            json[options.useProtoFieldName ? field.name : field.jsonName] = jsonValue;
          }
        }
      } catch (e) {
        const m = field ? `cannot encode field ${type.typeName}.${field.name} to JSON` : `cannot encode message ${type.typeName} to JSON`;
        const r = e instanceof Error ? e.message : String(e);
        throw new Error(m + (r.length > 0 ? `: ${r}` : ""));
      }
      return json;
    },
    readScalar: readScalar2,
    writeScalar: writeScalar2,
    debug: debugJsonValue
  };
}
function debugJsonValue(json) {
  if (json === null) {
    return "null";
  }
  switch (typeof json) {
    case "object":
      return Array.isArray(json) ? "array" : "object";
    case "string":
      return json.length > 100 ? "string" : `"${json.split('"').join('\\"')}"`;
    default:
      return String(json);
  }
}
function readScalar2(type, json, longType) {
  switch (type) {
    case ScalarType.DOUBLE:
    case ScalarType.FLOAT:
      if (json === null)
        return 0;
      if (json === "NaN")
        return Number.NaN;
      if (json === "Infinity")
        return Number.POSITIVE_INFINITY;
      if (json === "-Infinity")
        return Number.NEGATIVE_INFINITY;
      if (json === "") {
        break;
      }
      if (typeof json == "string" && json.trim().length !== json.length) {
        break;
      }
      if (typeof json != "string" && typeof json != "number") {
        break;
      }
      const float = Number(json);
      if (Number.isNaN(float)) {
        break;
      }
      if (!Number.isFinite(float)) {
        break;
      }
      if (type == ScalarType.FLOAT)
        assertFloat32(float);
      return float;
    case ScalarType.INT32:
    case ScalarType.FIXED32:
    case ScalarType.SFIXED32:
    case ScalarType.SINT32:
    case ScalarType.UINT32:
      if (json === null)
        return 0;
      let int32;
      if (typeof json == "number")
        int32 = json;
      else if (typeof json == "string" && json.length > 0) {
        if (json.trim().length === json.length)
          int32 = Number(json);
      }
      if (int32 === void 0)
        break;
      if (type == ScalarType.UINT32)
        assertUInt32(int32);
      else
        assertInt32(int32);
      return int32;
    case ScalarType.INT64:
    case ScalarType.SFIXED64:
    case ScalarType.SINT64:
      if (json === null)
        return protoInt64.zero;
      if (typeof json != "number" && typeof json != "string")
        break;
      const long = protoInt64.parse(json);
      return longType ? long.toString() : long;
    case ScalarType.FIXED64:
    case ScalarType.UINT64:
      if (json === null)
        return protoInt64.zero;
      if (typeof json != "number" && typeof json != "string")
        break;
      const uLong = protoInt64.uParse(json);
      return longType ? uLong.toString() : uLong;
    case ScalarType.BOOL:
      if (json === null)
        return false;
      if (typeof json !== "boolean")
        break;
      return json;
    case ScalarType.STRING:
      if (json === null)
        return "";
      if (typeof json !== "string") {
        break;
      }
      try {
        encodeURIComponent(json);
      } catch (e) {
        throw new Error("invalid UTF8");
      }
      return json;
    case ScalarType.BYTES:
      if (json === null || json === "")
        return new Uint8Array(0);
      if (typeof json !== "string")
        break;
      return protoBase64.dec(json);
  }
  throw new Error();
}
function readEnum(type, json, ignoreUnknownFields) {
  if (json === null) {
    return 0;
  }
  switch (typeof json) {
    case "number":
      if (Number.isInteger(json)) {
        return json;
      }
      break;
    case "string":
      const value = type.findName(json);
      if (value || ignoreUnknownFields) {
        return value === null || value === void 0 ? void 0 : value.no;
      }
      break;
  }
  throw new Error(`cannot decode enum ${type.typeName} from JSON: ${debugJsonValue(json)}`);
}
function writeEnum(type, value, emitIntrinsicDefault, enumAsInteger) {
  var _a;
  if (value === void 0) {
    return value;
  }
  if (value === 0 && !emitIntrinsicDefault) {
    return void 0;
  }
  if (enumAsInteger) {
    return value;
  }
  if (type.typeName == "google.protobuf.NullValue") {
    return null;
  }
  const val = type.findNumber(value);
  return (_a = val === null || val === void 0 ? void 0 : val.name) !== null && _a !== void 0 ? _a : value;
}
function writeScalar2(type, value, emitIntrinsicDefault) {
  if (value === void 0) {
    return void 0;
  }
  switch (type) {
    case ScalarType.INT32:
    case ScalarType.SFIXED32:
    case ScalarType.SINT32:
    case ScalarType.FIXED32:
    case ScalarType.UINT32:
      assert(typeof value == "number");
      return value != 0 || emitIntrinsicDefault ? value : void 0;
    case ScalarType.FLOAT:
    case ScalarType.DOUBLE:
      assert(typeof value == "number");
      if (Number.isNaN(value))
        return "NaN";
      if (value === Number.POSITIVE_INFINITY)
        return "Infinity";
      if (value === Number.NEGATIVE_INFINITY)
        return "-Infinity";
      return value !== 0 || emitIntrinsicDefault ? value : void 0;
    case ScalarType.STRING:
      assert(typeof value == "string");
      return value.length > 0 || emitIntrinsicDefault ? value : void 0;
    case ScalarType.BOOL:
      assert(typeof value == "boolean");
      return value || emitIntrinsicDefault ? value : void 0;
    case ScalarType.UINT64:
    case ScalarType.FIXED64:
    case ScalarType.INT64:
    case ScalarType.SFIXED64:
    case ScalarType.SINT64:
      assert(typeof value == "bigint" || typeof value == "string" || typeof value == "number");
      return emitIntrinsicDefault || value != 0 ? value.toString(10) : void 0;
    case ScalarType.BYTES:
      assert(value instanceof Uint8Array);
      return emitIntrinsicDefault || value.byteLength > 0 ? protoBase64.enc(value) : void 0;
  }
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/json-format-proto3.js
function makeJsonFormatProto3() {
  return makeJsonFormatCommon((writeEnum2, writeScalar3) => {
    return function writeField(field, value, options) {
      if (field.kind == "map") {
        const jsonObj = {};
        switch (field.V.kind) {
          case "scalar":
            for (const [entryKey, entryValue] of Object.entries(value)) {
              const val = writeScalar3(field.V.T, entryValue, true);
              assert(val !== void 0);
              jsonObj[entryKey.toString()] = val;
            }
            break;
          case "message":
            for (const [entryKey, entryValue] of Object.entries(value)) {
              jsonObj[entryKey.toString()] = entryValue.toJson(options);
            }
            break;
          case "enum":
            const enumType = field.V.T;
            for (const [entryKey, entryValue] of Object.entries(value)) {
              assert(entryValue === void 0 || typeof entryValue == "number");
              const val = writeEnum2(enumType, entryValue, true, options.enumAsInteger);
              assert(val !== void 0);
              jsonObj[entryKey.toString()] = val;
            }
            break;
        }
        return options.emitDefaultValues || Object.keys(jsonObj).length > 0 ? jsonObj : void 0;
      } else if (field.repeated) {
        const jsonArr = [];
        switch (field.kind) {
          case "scalar":
            for (let i = 0; i < value.length; i++) {
              jsonArr.push(writeScalar3(field.T, value[i], true));
            }
            break;
          case "enum":
            for (let i = 0; i < value.length; i++) {
              jsonArr.push(writeEnum2(field.T, value[i], true, options.enumAsInteger));
            }
            break;
          case "message":
            for (let i = 0; i < value.length; i++) {
              jsonArr.push(wrapField(field.T, value[i]).toJson(options));
            }
            break;
        }
        return options.emitDefaultValues || jsonArr.length > 0 ? jsonArr : void 0;
      } else {
        switch (field.kind) {
          case "scalar":
            return writeScalar3(field.T, value, !!field.oneof || field.opt || options.emitDefaultValues);
          case "enum":
            return writeEnum2(field.T, value, !!field.oneof || field.opt || options.emitDefaultValues, options.enumAsInteger);
          case "message":
            return value !== void 0 ? wrapField(field.T, value).toJson(options) : void 0;
        }
      }
    };
  });
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/util-common.js
function makeUtilCommon() {
  return {
    setEnumType,
    initPartial(source, target) {
      if (source === void 0) {
        return;
      }
      const type = target.getType();
      for (const member of type.fields.byMember()) {
        const localName = member.localName, t = target, s = source;
        if (s[localName] === void 0) {
          continue;
        }
        switch (member.kind) {
          case "oneof":
            const sk = s[localName].case;
            if (sk === void 0) {
              continue;
            }
            const sourceField = member.findField(sk);
            let val = s[localName].value;
            if (sourceField && sourceField.kind == "message" && !(val instanceof sourceField.T)) {
              val = new sourceField.T(val);
            } else if (sourceField && sourceField.kind === "scalar" && sourceField.T === ScalarType.BYTES) {
              val = toU8Arr(val);
            }
            t[localName] = { case: sk, value: val };
            break;
          case "scalar":
          case "enum":
            let copy = s[localName];
            if (member.T === ScalarType.BYTES) {
              copy = member.repeated ? copy.map(toU8Arr) : toU8Arr(copy);
            }
            t[localName] = copy;
            break;
          case "map":
            switch (member.V.kind) {
              case "scalar":
              case "enum":
                if (member.V.T === ScalarType.BYTES) {
                  for (const [k, v] of Object.entries(s[localName])) {
                    t[localName][k] = toU8Arr(v);
                  }
                } else {
                  Object.assign(t[localName], s[localName]);
                }
                break;
              case "message":
                const messageType = member.V.T;
                for (const k of Object.keys(s[localName])) {
                  let val2 = s[localName][k];
                  if (!messageType.fieldWrapper) {
                    val2 = new messageType(val2);
                  }
                  t[localName][k] = val2;
                }
                break;
            }
            break;
          case "message":
            const mt = member.T;
            if (member.repeated) {
              t[localName] = s[localName].map((val2) => val2 instanceof mt ? val2 : new mt(val2));
            } else if (s[localName] !== void 0) {
              const val2 = s[localName];
              if (mt.fieldWrapper) {
                if (
                  // We can't use BytesValue.typeName as that will create a circular import
                  mt.typeName === "google.protobuf.BytesValue"
                ) {
                  t[localName] = toU8Arr(val2);
                } else {
                  t[localName] = val2;
                }
              } else {
                t[localName] = val2 instanceof mt ? val2 : new mt(val2);
              }
            }
            break;
        }
      }
    },
    equals(type, a, b) {
      if (a === b) {
        return true;
      }
      if (!a || !b) {
        return false;
      }
      return type.fields.byMember().every((m) => {
        const va = a[m.localName];
        const vb = b[m.localName];
        if (m.repeated) {
          if (va.length !== vb.length) {
            return false;
          }
          switch (m.kind) {
            case "message":
              return va.every((a2, i) => m.T.equals(a2, vb[i]));
            case "scalar":
              return va.every((a2, i) => scalarEquals(m.T, a2, vb[i]));
            case "enum":
              return va.every((a2, i) => scalarEquals(ScalarType.INT32, a2, vb[i]));
          }
          throw new Error(`repeated cannot contain ${m.kind}`);
        }
        switch (m.kind) {
          case "message":
            return m.T.equals(va, vb);
          case "enum":
            return scalarEquals(ScalarType.INT32, va, vb);
          case "scalar":
            return scalarEquals(m.T, va, vb);
          case "oneof":
            if (va.case !== vb.case) {
              return false;
            }
            const s = m.findField(va.case);
            if (s === void 0) {
              return true;
            }
            switch (s.kind) {
              case "message":
                return s.T.equals(va.value, vb.value);
              case "enum":
                return scalarEquals(ScalarType.INT32, va.value, vb.value);
              case "scalar":
                return scalarEquals(s.T, va.value, vb.value);
            }
            throw new Error(`oneof cannot contain ${s.kind}`);
          case "map":
            const keys = Object.keys(va).concat(Object.keys(vb));
            switch (m.V.kind) {
              case "message":
                const messageType = m.V.T;
                return keys.every((k) => messageType.equals(va[k], vb[k]));
              case "enum":
                return keys.every((k) => scalarEquals(ScalarType.INT32, va[k], vb[k]));
              case "scalar":
                const scalarType = m.V.T;
                return keys.every((k) => scalarEquals(scalarType, va[k], vb[k]));
            }
            break;
        }
      });
    },
    clone(message) {
      const type = message.getType(), target = new type(), any = target;
      for (const member of type.fields.byMember()) {
        const source = message[member.localName];
        let copy;
        if (member.repeated) {
          copy = source.map(cloneSingularField);
        } else if (member.kind == "map") {
          copy = any[member.localName];
          for (const [key, v] of Object.entries(source)) {
            copy[key] = cloneSingularField(v);
          }
        } else if (member.kind == "oneof") {
          const f = member.findField(source.case);
          copy = f ? { case: source.case, value: cloneSingularField(source.value) } : { case: void 0 };
        } else {
          copy = cloneSingularField(source);
        }
        any[member.localName] = copy;
      }
      return target;
    }
  };
}
function cloneSingularField(value) {
  if (value === void 0) {
    return value;
  }
  if (value instanceof Message) {
    return value.clone();
  }
  if (value instanceof Uint8Array) {
    const c = new Uint8Array(value.byteLength);
    c.set(value);
    return c;
  }
  return value;
}
function toU8Arr(input) {
  return input instanceof Uint8Array ? input : new Uint8Array(input);
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/field-list.js
var InternalFieldList = class {
  constructor(fields, normalizer) {
    this._fields = fields;
    this._normalizer = normalizer;
  }
  findJsonName(jsonName) {
    if (!this.jsonNames) {
      const t = {};
      for (const f of this.list()) {
        t[f.jsonName] = t[f.name] = f;
      }
      this.jsonNames = t;
    }
    return this.jsonNames[jsonName];
  }
  find(fieldNo) {
    if (!this.numbers) {
      const t = {};
      for (const f of this.list()) {
        t[f.no] = f;
      }
      this.numbers = t;
    }
    return this.numbers[fieldNo];
  }
  list() {
    if (!this.all) {
      this.all = this._normalizer(this._fields);
    }
    return this.all;
  }
  byNumber() {
    if (!this.numbersAsc) {
      this.numbersAsc = this.list().concat().sort((a, b) => a.no - b.no);
    }
    return this.numbersAsc;
  }
  byMember() {
    if (!this.members) {
      this.members = [];
      const a = this.members;
      let o;
      for (const f of this.list()) {
        if (f.oneof) {
          if (f.oneof !== o) {
            o = f.oneof;
            a.push(o);
          }
        } else {
          a.push(f);
        }
      }
    }
    return this.members;
  }
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/names.js
function localFieldName(protoName, inOneof) {
  const name = protoCamelCase(protoName);
  if (inOneof) {
    return name;
  }
  return safeObjectProperty(safeMessageProperty(name));
}
function localOneofName(protoName) {
  return localFieldName(protoName, false);
}
var fieldJsonName = protoCamelCase;
function protoCamelCase(snakeCase) {
  let capNext = false;
  const b = [];
  for (let i = 0; i < snakeCase.length; i++) {
    let c = snakeCase.charAt(i);
    switch (c) {
      case "_":
        capNext = true;
        break;
      case "0":
      case "1":
      case "2":
      case "3":
      case "4":
      case "5":
      case "6":
      case "7":
      case "8":
      case "9":
        b.push(c);
        capNext = false;
        break;
      default:
        if (capNext) {
          capNext = false;
          c = c.toUpperCase();
        }
        b.push(c);
        break;
    }
  }
  return b.join("");
}
var reservedObjectProperties = /* @__PURE__ */ new Set([
  // names reserved by JavaScript
  "constructor",
  "toString",
  "toJSON",
  "valueOf"
]);
var reservedMessageProperties = /* @__PURE__ */ new Set([
  // names reserved by the runtime
  "getType",
  "clone",
  "equals",
  "fromBinary",
  "fromJson",
  "fromJsonString",
  "toBinary",
  "toJson",
  "toJsonString",
  // names reserved by the runtime for the future
  "toObject"
]);
var fallback = (name) => `${name}$`;
var safeMessageProperty = (name) => {
  if (reservedMessageProperties.has(name)) {
    return fallback(name);
  }
  return name;
};
var safeObjectProperty = (name) => {
  if (reservedObjectProperties.has(name)) {
    return fallback(name);
  }
  return name;
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/private/field.js
var InternalOneofInfo = class {
  constructor(name) {
    this.kind = "oneof";
    this.repeated = false;
    this.packed = false;
    this.opt = false;
    this.default = void 0;
    this.fields = [];
    this.name = name;
    this.localName = localOneofName(name);
  }
  addField(field) {
    assert(field.oneof === this, `field ${field.name} not one of ${this.name}`);
    this.fields.push(field);
  }
  findField(localName) {
    if (!this._lookup) {
      this._lookup = /* @__PURE__ */ Object.create(null);
      for (let i = 0; i < this.fields.length; i++) {
        this._lookup[this.fields[i].localName] = this.fields[i];
      }
    }
    return this._lookup[localName];
  }
};

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/proto3.js
var proto3 = makeProtoRuntime("proto3", makeJsonFormatProto3(), makeBinaryFormatProto3(), Object.assign(Object.assign({}, makeUtilCommon()), {
  newFieldList(fields) {
    return new InternalFieldList(fields, normalizeFieldInfosProto3);
  },
  initFields(target) {
    for (const member of target.getType().fields.byMember()) {
      if (member.opt) {
        continue;
      }
      const name = member.localName, t = target;
      if (member.repeated) {
        t[name] = [];
        continue;
      }
      switch (member.kind) {
        case "oneof":
          t[name] = { case: void 0 };
          break;
        case "enum":
          t[name] = 0;
          break;
        case "map":
          t[name] = {};
          break;
        case "scalar":
          t[name] = scalarDefaultValue(member.T, member.L);
          break;
        case "message":
          break;
      }
    }
  }
}));
function normalizeFieldInfosProto3(fieldInfos) {
  var _a, _b, _c, _d;
  const r = [];
  let o;
  for (const field of typeof fieldInfos == "function" ? fieldInfos() : fieldInfos) {
    const f = field;
    f.localName = localFieldName(field.name, field.oneof !== void 0);
    f.jsonName = (_a = field.jsonName) !== null && _a !== void 0 ? _a : fieldJsonName(field.name);
    f.repeated = (_b = field.repeated) !== null && _b !== void 0 ? _b : false;
    if (field.kind == "scalar") {
      f.L = (_c = field.L) !== null && _c !== void 0 ? _c : LongType.BIGINT;
    }
    f.packed = (_d = field.packed) !== null && _d !== void 0 ? _d : field.kind == "enum" || field.kind == "scalar" && field.T != ScalarType.BYTES && field.T != ScalarType.STRING;
    if (field.oneof !== void 0) {
      const ooname = typeof field.oneof == "string" ? field.oneof : field.oneof.name;
      if (!o || o.name != ooname) {
        o = new InternalOneofInfo(ooname);
      }
      f.oneof = o;
      o.addField(f);
    }
    r.push(f);
  }
  return r;
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/create-registry.js
function createRegistry(...types) {
  const messages = {};
  const enums = {};
  const services = {};
  const registry = {
    /**
     * Add a type to the registry. For messages, the types used in message
     * fields are added recursively. For services, the message types used
     * for requests and responses are added recursively.
     */
    add(type) {
      if ("fields" in type) {
        if (!this.findMessage(type.typeName)) {
          messages[type.typeName] = type;
          for (const field of type.fields.list()) {
            if (field.kind == "message") {
              this.add(field.T);
            } else if (field.kind == "map" && field.V.kind == "message") {
              this.add(field.V.T);
            } else if (field.kind == "enum") {
              this.add(field.T);
            }
          }
        }
      } else if ("methods" in type) {
        if (!this.findService(type.typeName)) {
          services[type.typeName] = type;
          for (const method of Object.values(type.methods)) {
            this.add(method.I);
            this.add(method.O);
          }
        }
      } else {
        enums[type.typeName] = type;
      }
    },
    findMessage(typeName) {
      return messages[typeName];
    },
    findEnum(typeName) {
      return enums[typeName];
    },
    findService(typeName) {
      return services[typeName];
    }
  };
  for (const type of types) {
    registry.add(type);
  }
  return registry;
}

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/google/protobuf/timestamp_pb.js
var Timestamp = class _Timestamp extends Message {
  constructor(data) {
    super();
    this.seconds = protoInt64.zero;
    this.nanos = 0;
    proto3.util.initPartial(data, this);
  }
  fromJson(json, options) {
    if (typeof json !== "string") {
      throw new Error(`cannot decode google.protobuf.Timestamp from JSON: ${proto3.json.debug(json)}`);
    }
    const matches = json.match(/^([0-9]{4})-([0-9]{2})-([0-9]{2})T([0-9]{2}):([0-9]{2}):([0-9]{2})(?:Z|\.([0-9]{3,9})Z|([+-][0-9][0-9]:[0-9][0-9]))$/);
    if (!matches) {
      throw new Error(`cannot decode google.protobuf.Timestamp from JSON: invalid RFC 3339 string`);
    }
    const ms = Date.parse(matches[1] + "-" + matches[2] + "-" + matches[3] + "T" + matches[4] + ":" + matches[5] + ":" + matches[6] + (matches[8] ? matches[8] : "Z"));
    if (Number.isNaN(ms)) {
      throw new Error(`cannot decode google.protobuf.Timestamp from JSON: invalid RFC 3339 string`);
    }
    if (ms < Date.parse("0001-01-01T00:00:00Z") || ms > Date.parse("9999-12-31T23:59:59Z")) {
      throw new Error(`cannot decode message google.protobuf.Timestamp from JSON: must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive`);
    }
    this.seconds = protoInt64.parse(ms / 1e3);
    this.nanos = 0;
    if (matches[7]) {
      this.nanos = parseInt("1" + matches[7] + "0".repeat(9 - matches[7].length)) - 1e9;
    }
    return this;
  }
  toJson(options) {
    const ms = Number(this.seconds) * 1e3;
    if (ms < Date.parse("0001-01-01T00:00:00Z") || ms > Date.parse("9999-12-31T23:59:59Z")) {
      throw new Error(`cannot encode google.protobuf.Timestamp to JSON: must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive`);
    }
    if (this.nanos < 0) {
      throw new Error(`cannot encode google.protobuf.Timestamp to JSON: nanos must not be negative`);
    }
    let z = "Z";
    if (this.nanos > 0) {
      const nanosStr = (this.nanos + 1e9).toString().substring(1);
      if (nanosStr.substring(3) === "000000") {
        z = "." + nanosStr.substring(0, 3) + "Z";
      } else if (nanosStr.substring(6) === "000") {
        z = "." + nanosStr.substring(0, 6) + "Z";
      } else {
        z = "." + nanosStr + "Z";
      }
    }
    return new Date(ms).toISOString().replace(".000Z", z);
  }
  toDate() {
    return new Date(Number(this.seconds) * 1e3 + Math.ceil(this.nanos / 1e6));
  }
  static now() {
    return _Timestamp.fromDate(/* @__PURE__ */ new Date());
  }
  static fromDate(date) {
    const ms = date.getTime();
    return new _Timestamp({
      seconds: protoInt64.parse(Math.floor(ms / 1e3)),
      nanos: ms % 1e3 * 1e6
    });
  }
  static fromBinary(bytes, options) {
    return new _Timestamp().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Timestamp().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Timestamp().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_Timestamp, a, b);
  }
};
Timestamp.runtime = proto3;
Timestamp.typeName = "google.protobuf.Timestamp";
Timestamp.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "seconds",
    kind: "scalar",
    T: 3
    /* ScalarType.INT64 */
  },
  {
    no: 2,
    name: "nanos",
    kind: "scalar",
    T: 5
    /* ScalarType.INT32 */
  }
]);

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/google/protobuf/any_pb.js
var Any = class _Any extends Message {
  constructor(data) {
    super();
    this.typeUrl = "";
    this.value = new Uint8Array(0);
    proto3.util.initPartial(data, this);
  }
  toJson(options) {
    var _a;
    if (this.typeUrl === "") {
      return {};
    }
    const typeName = this.typeUrlToName(this.typeUrl);
    const messageType = (_a = options === null || options === void 0 ? void 0 : options.typeRegistry) === null || _a === void 0 ? void 0 : _a.findMessage(typeName);
    if (!messageType) {
      throw new Error(`cannot encode message google.protobuf.Any to JSON: "${this.typeUrl}" is not in the type registry`);
    }
    const message = messageType.fromBinary(this.value);
    let json = message.toJson(options);
    if (typeName.startsWith("google.protobuf.") || (json === null || Array.isArray(json) || typeof json !== "object")) {
      json = { value: json };
    }
    json["@type"] = this.typeUrl;
    return json;
  }
  fromJson(json, options) {
    var _a;
    if (json === null || Array.isArray(json) || typeof json != "object") {
      throw new Error(`cannot decode message google.protobuf.Any from JSON: expected object but got ${json === null ? "null" : Array.isArray(json) ? "array" : typeof json}`);
    }
    if (Object.keys(json).length == 0) {
      return this;
    }
    const typeUrl = json["@type"];
    if (typeof typeUrl != "string" || typeUrl == "") {
      throw new Error(`cannot decode message google.protobuf.Any from JSON: "@type" is empty`);
    }
    const typeName = this.typeUrlToName(typeUrl), messageType = (_a = options === null || options === void 0 ? void 0 : options.typeRegistry) === null || _a === void 0 ? void 0 : _a.findMessage(typeName);
    if (!messageType) {
      throw new Error(`cannot decode message google.protobuf.Any from JSON: ${typeUrl} is not in the type registry`);
    }
    let message;
    if (typeName.startsWith("google.protobuf.") && Object.prototype.hasOwnProperty.call(json, "value")) {
      message = messageType.fromJson(json["value"], options);
    } else {
      const copy = Object.assign({}, json);
      delete copy["@type"];
      message = messageType.fromJson(copy, options);
    }
    this.packFrom(message);
    return this;
  }
  packFrom(message) {
    this.value = message.toBinary();
    this.typeUrl = this.typeNameToUrl(message.getType().typeName);
  }
  unpackTo(target) {
    if (!this.is(target.getType())) {
      return false;
    }
    target.fromBinary(this.value);
    return true;
  }
  unpack(registry) {
    if (this.typeUrl === "") {
      return void 0;
    }
    const messageType = registry.findMessage(this.typeUrlToName(this.typeUrl));
    if (!messageType) {
      return void 0;
    }
    return messageType.fromBinary(this.value);
  }
  is(type) {
    if (this.typeUrl === "") {
      return false;
    }
    const name = this.typeUrlToName(this.typeUrl);
    let typeName = "";
    if (typeof type === "string") {
      typeName = type;
    } else {
      typeName = type.typeName;
    }
    return name === typeName;
  }
  typeNameToUrl(name) {
    return `type.googleapis.com/${name}`;
  }
  typeUrlToName(url) {
    if (!url.length) {
      throw new Error(`invalid type url: ${url}`);
    }
    const slash = url.lastIndexOf("/");
    const name = slash > 0 ? url.substring(slash + 1) : url;
    if (!name.length) {
      throw new Error(`invalid type url: ${url}`);
    }
    return name;
  }
  static pack(message) {
    const any = new _Any();
    any.packFrom(message);
    return any;
  }
  static fromBinary(bytes, options) {
    return new _Any().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _Any().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _Any().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_Any, a, b);
  }
};
Any.runtime = proto3;
Any.typeName = "google.protobuf.Any";
Any.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "type_url",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  {
    no: 2,
    name: "value",
    kind: "scalar",
    T: 12
    /* ScalarType.BYTES */
  }
]);

// node_modules/.pnpm/@bufbuild+protobuf@1.4.2/node_modules/@bufbuild/protobuf/dist/esm/google/protobuf/field_mask_pb.js
var FieldMask = class _FieldMask extends Message {
  constructor(data) {
    super();
    this.paths = [];
    proto3.util.initPartial(data, this);
  }
  toJson(options) {
    function protoCamelCase2(snakeCase) {
      let capNext = false;
      const b = [];
      for (let i = 0; i < snakeCase.length; i++) {
        let c = snakeCase.charAt(i);
        switch (c) {
          case "_":
            capNext = true;
            break;
          case "0":
          case "1":
          case "2":
          case "3":
          case "4":
          case "5":
          case "6":
          case "7":
          case "8":
          case "9":
            b.push(c);
            capNext = false;
            break;
          default:
            if (capNext) {
              capNext = false;
              c = c.toUpperCase();
            }
            b.push(c);
            break;
        }
      }
      return b.join("");
    }
    return this.paths.map((p) => {
      if (p.match(/_[0-9]?_/g) || p.match(/[A-Z]/g)) {
        throw new Error('cannot encode google.protobuf.FieldMask to JSON: lowerCamelCase of path name "' + p + '" is irreversible');
      }
      return protoCamelCase2(p);
    }).join(",");
  }
  fromJson(json, options) {
    if (typeof json !== "string") {
      throw new Error("cannot decode google.protobuf.FieldMask from JSON: " + proto3.json.debug(json));
    }
    if (json === "") {
      return this;
    }
    function camelToSnake(str) {
      if (str.includes("_")) {
        throw new Error("cannot decode google.protobuf.FieldMask from JSON: path names must be lowerCamelCase");
      }
      const sc = str.replace(/[A-Z]/g, (letter) => "_" + letter.toLowerCase());
      return sc[0] === "_" ? sc.substring(1) : sc;
    }
    this.paths = json.split(",").map(camelToSnake);
    return this;
  }
  static fromBinary(bytes, options) {
    return new _FieldMask().fromBinary(bytes, options);
  }
  static fromJson(jsonValue, options) {
    return new _FieldMask().fromJson(jsonValue, options);
  }
  static fromJsonString(jsonString, options) {
    return new _FieldMask().fromJsonString(jsonString, options);
  }
  static equals(a, b) {
    return proto3.util.equals(_FieldMask, a, b);
  }
};
FieldMask.runtime = proto3;
FieldMask.typeName = "google.protobuf.FieldMask";
FieldMask.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "paths", kind: "scalar", T: 9, repeated: true }
]);

// gen/auth/v1/auth_pb.ts
var TransactionType = /* @__PURE__ */ ((TransactionType2) => {
  TransactionType2[TransactionType2["UNSPECIFIED"] = 0] = "UNSPECIFIED";
  TransactionType2[TransactionType2["INVOKE"] = 1] = "INVOKE";
  TransactionType2[TransactionType2["QUERY"] = 2] = "QUERY";
  return TransactionType2;
})(TransactionType || {});
proto3.util.setEnumType(TransactionType, "auth.TransactionType", [
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
proto3.util.setEnumType(AuthType, "auth.AuthType", [
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
proto3.util.setEnumType(ItemKind, "auth.ItemKind", [
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
proto3.util.setEnumType(Action, "auth.Action", [
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
proto3.util.setEnumType(TxError, "auth.TxError", [
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
var _KeySchema = class _KeySchema extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_KeySchema, a, b);
  }
};
_KeySchema.runtime = proto3;
_KeySchema.typeName = "auth.KeySchema";
_KeySchema.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "item_type",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 2, name: "item_kind", kind: "enum", T: proto3.getEnumType(ItemKind) },
  { no: 3, name: "keys", kind: "message", T: FieldMask }
]);
var KeySchema = _KeySchema;
var _StateActivity = class _StateActivity extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_StateActivity, a, b);
  }
};
_StateActivity.runtime = proto3;
_StateActivity.typeName = "auth.StateActivity";
_StateActivity.fields = proto3.util.newFieldList(() => [
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
  { no: 4, name: "timestamp", kind: "message", T: Timestamp },
  {
    no: 5,
    name: "note",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var StateActivity = _StateActivity;
var _Operation = class _Operation extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Operation, a, b);
  }
};
_Operation.runtime = proto3;
_Operation.typeName = "auth.Operation";
_Operation.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "action", kind: "enum", T: proto3.getEnumType(Action) },
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
  { no: 5, name: "paths", kind: "message", T: FieldMask }
]);
var Operation = _Operation;
var _PathPolicy = class _PathPolicy extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_PathPolicy, a, b);
  }
};
_PathPolicy.runtime = proto3;
_PathPolicy.typeName = "auth.PathPolicy";
_PathPolicy.fields = proto3.util.newFieldList(() => [
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
  { no: 5, name: "actions", kind: "enum", T: proto3.getEnumType(Action), repeated: true }
]);
var PathPolicy = _PathPolicy;
var _Polices = class _Polices extends Message {
  constructor(data) {
    super();
    /**
     * key is the item type
     *
     * @generated from field: map<string, auth.PathPolicy> item_policies = 1;
     */
    this.itemPolicies = {};
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Polices, a, b);
  }
};
_Polices.runtime = proto3;
_Polices.typeName = "auth.Polices";
_Polices.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item_policies", kind: "map", K: 9, V: { kind: "message", T: PathPolicy } }
]);
var Polices = _Polices;
var _Item = class _Item extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Item, a, b);
  }
};
_Item.runtime = proto3;
_Item.typeName = "auth.Item";
_Item.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "key", kind: "message", T: ItemKey },
  { no: 2, name: "value", kind: "message", T: Any }
]);
var Item = _Item;
var _FullItem = class _FullItem extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_FullItem, a, b);
  }
};
_FullItem.runtime = proto3;
_FullItem.typeName = "auth.FullItem";
_FullItem.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "history", kind: "message", T: History },
  { no: 3, name: "suggestions", kind: "message", T: Suggestion, repeated: true },
  { no: 4, name: "references", kind: "message", T: Reference, repeated: true }
]);
var FullItem = _FullItem;
var _HistoryEntry = class _HistoryEntry extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HistoryEntry, a, b);
  }
};
_HistoryEntry.runtime = proto3;
_HistoryEntry.typeName = "auth.HistoryEntry";
_HistoryEntry.fields = proto3.util.newFieldList(() => [
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
  { no: 4, name: "timestamp", kind: "message", T: Timestamp },
  {
    no: 5,
    name: "note",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 6, name: "value", kind: "message", T: Any }
]);
var HistoryEntry = _HistoryEntry;
var _History = class _History extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated auth.HistoryEntry entries = 1;
     */
    this.entries = [];
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_History, a, b);
  }
};
_History.runtime = proto3;
_History.typeName = "auth.History";
_History.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "entries", kind: "message", T: HistoryEntry, repeated: true },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTxList }
]);
var History = _History;
var _ItemKey = class _ItemKey extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ItemKey, a, b);
  }
};
_ItemKey.runtime = proto3;
_ItemKey.typeName = "auth.ItemKey";
_ItemKey.fields = proto3.util.newFieldList(() => [
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
var _ReferenceKey = class _ReferenceKey extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceKey, a, b);
  }
};
_ReferenceKey.runtime = proto3;
_ReferenceKey.typeName = "auth.ReferenceKey";
_ReferenceKey.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "key1", kind: "message", T: ItemKey },
  { no: 2, name: "key2", kind: "message", T: ItemKey }
]);
var ReferenceKey = _ReferenceKey;
var _Reference = class _Reference extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Reference, a, b);
  }
};
_Reference.runtime = proto3;
_Reference.typeName = "auth.Reference";
_Reference.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "reference", kind: "message", T: ReferenceKey },
  { no: 2, name: "item1", kind: "message", T: Item },
  { no: 3, name: "item2", kind: "message", T: Item }
]);
var Reference = _Reference;
var _Collection = class _Collection extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Collection, a, b);
  }
};
_Collection.runtime = proto3;
_Collection.typeName = "auth.Collection";
_Collection.fields = proto3.util.newFieldList(() => [
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
  { no: 4, name: "auth_type", kind: "enum", T: proto3.getEnumType(AuthType) },
  { no: 5, name: "item_types", kind: "scalar", T: 9, repeated: true },
  { no: 6, name: "reference_types", kind: "scalar", T: 9, repeated: true },
  { no: 7, name: "default", kind: "message", T: Polices }
]);
var Collection = _Collection;
var _User = class _User extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_User, a, b);
  }
};
_User.runtime = proto3;
_User.typeName = "auth.User";
_User.fields = proto3.util.newFieldList(() => [
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
var _Suggestion = class _Suggestion extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string suggestion_id = 2;
     */
    this.suggestionId = "";
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Suggestion, a, b);
  }
};
_Suggestion.runtime = proto3;
_Suggestion.typeName = "auth.Suggestion";
_Suggestion.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "primary_key", kind: "message", T: ItemKey },
  {
    no: 2,
    name: "suggestion_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  },
  { no: 5, name: "paths", kind: "message", T: FieldMask },
  { no: 6, name: "value", kind: "message", T: Any }
]);
var Suggestion = _Suggestion;
var _HiddenTx = class _HiddenTx extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HiddenTx, a, b);
  }
};
_HiddenTx.runtime = proto3;
_HiddenTx.typeName = "auth.HiddenTx";
_HiddenTx.fields = proto3.util.newFieldList(() => [
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
  { no: 4, name: "timestamp", kind: "message", T: Timestamp },
  {
    no: 5,
    name: "note",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var HiddenTx = _HiddenTx;
var _HiddenTxList = class _HiddenTxList extends Message {
  constructor(data) {
    super();
    /**
     * The list of hidden txs by tx_id
     *
     * @generated from field: repeated auth.HiddenTx txs = 4;
     */
    this.txs = [];
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HiddenTxList, a, b);
  }
};
_HiddenTxList.runtime = proto3;
_HiddenTxList.typeName = "auth.HiddenTxList";
_HiddenTxList.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "primary_key", kind: "message", T: ItemKey },
  { no: 4, name: "txs", kind: "message", T: HiddenTx, repeated: true }
]);
var HiddenTxList = _HiddenTxList;
var _Role = class _Role extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Role, a, b);
  }
};
_Role.runtime = proto3;
_Role.typeName = "auth.Role";
_Role.fields = proto3.util.newFieldList(() => [
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
var _Attribute = class _Attribute extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Attribute, a, b);
  }
};
_Attribute.runtime = proto3;
_Attribute.typeName = "auth.Attribute";
_Attribute.fields = proto3.util.newFieldList(() => [
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
var _UserMembership = class _UserMembership extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_UserMembership, a, b);
  }
};
_UserMembership.runtime = proto3;
_UserMembership.typeName = "auth.UserMembership";
_UserMembership.fields = proto3.util.newFieldList(() => [
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
var _UserCollectionRoles = class _UserCollectionRoles extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_UserCollectionRoles, a, b);
  }
};
_UserCollectionRoles.runtime = proto3;
_UserCollectionRoles.typeName = "auth.UserCollectionRoles";
_UserCollectionRoles.fields = proto3.util.newFieldList(() => [
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
var _GetCurrentUserResponse = class _GetCurrentUserResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool registerd = 2;
     */
    this.registerd = false;
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_GetCurrentUserResponse, a, b);
  }
};
_GetCurrentUserResponse.runtime = proto3;
_GetCurrentUserResponse.typeName = "auth.common.GetCurrentUserResponse";
_GetCurrentUserResponse.fields = proto3.util.newFieldList(() => [
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
var _GetCurrentFullUserResponse = class _GetCurrentFullUserResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_GetCurrentFullUserResponse, a, b);
  }
};
_GetCurrentFullUserResponse.runtime = proto3;
_GetCurrentFullUserResponse.typeName = "auth.common.GetCurrentFullUserResponse";
_GetCurrentFullUserResponse.fields = proto3.util.newFieldList(() => [
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
var _AuthorizeOperationRequest = class _AuthorizeOperationRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_AuthorizeOperationRequest, a, b);
  }
};
_AuthorizeOperationRequest.runtime = proto3;
_AuthorizeOperationRequest.typeName = "auth.common.AuthorizeOperationRequest";
_AuthorizeOperationRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "operation", kind: "message", T: Operation }
]);
var AuthorizeOperationRequest = _AuthorizeOperationRequest;
var _AuthorizeOperationResponse = class _AuthorizeOperationResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool authorized = 1;
     */
    this.authorized = false;
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_AuthorizeOperationResponse, a, b);
  }
};
_AuthorizeOperationResponse.runtime = proto3;
_AuthorizeOperationResponse.typeName = "auth.common.AuthorizeOperationResponse";
_AuthorizeOperationResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "authorized",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var AuthorizeOperationResponse = _AuthorizeOperationResponse;
var _BootstrapRequest = class _BootstrapRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_BootstrapRequest, a, b);
  }
};
_BootstrapRequest.runtime = proto3;
_BootstrapRequest.typeName = "auth.common.BootstrapRequest";
_BootstrapRequest.fields = proto3.util.newFieldList(() => [
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
var _BootstrapResponse = class _BootstrapResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool success = 1;
     */
    this.success = false;
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_BootstrapResponse, a, b);
  }
};
_BootstrapResponse.runtime = proto3;
_BootstrapResponse.typeName = "auth.common.BootstrapResponse";
_BootstrapResponse.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "success",
    kind: "scalar",
    T: 8
    /* ScalarType.BOOL */
  }
]);
var BootstrapResponse = _BootstrapResponse;
var _CreateCollectionRequest = class _CreateCollectionRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_CreateCollectionRequest, a, b);
  }
};
_CreateCollectionRequest.runtime = proto3;
_CreateCollectionRequest.typeName = "auth.common.CreateCollectionRequest";
_CreateCollectionRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "collection", kind: "message", T: Collection }
]);
var CreateCollectionRequest = _CreateCollectionRequest;
var _CreateCollectionResponse = class _CreateCollectionResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_CreateCollectionResponse, a, b);
  }
};
_CreateCollectionResponse.runtime = proto3;
_CreateCollectionResponse.typeName = "auth.common.CreateCollectionResponse";
_CreateCollectionResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "collection", kind: "message", T: Collection }
]);
var CreateCollectionResponse = _CreateCollectionResponse;
var _CreateUserResponse = class _CreateUserResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_CreateUserResponse, a, b);
  }
};
_CreateUserResponse.runtime = proto3;
_CreateUserResponse.typeName = "auth.common.CreateUserResponse";
_CreateUserResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "user", kind: "message", T: User }
]);
var CreateUserResponse = _CreateUserResponse;
var _GetRequest = class _GetRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_GetRequest, a, b);
  }
};
_GetRequest.runtime = proto3;
_GetRequest.typeName = "auth.common.GetRequest";
_GetRequest.fields = proto3.util.newFieldList(() => [
  { no: 3, name: "item", kind: "message", T: Item }
]);
var GetRequest = _GetRequest;
var _GetResponse = class _GetResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_GetResponse, a, b);
  }
};
_GetResponse.runtime = proto3;
_GetResponse.typeName = "auth.common.GetResponse";
_GetResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item }
]);
var GetResponse = _GetResponse;
var _ListRequest = class _ListRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ListRequest, a, b);
  }
};
_ListRequest.runtime = proto3;
_ListRequest.typeName = "auth.common.ListRequest";
_ListRequest.fields = proto3.util.newFieldList(() => [
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
var _ListResponse = class _ListResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ListResponse, a, b);
  }
};
_ListResponse.runtime = proto3;
_ListResponse.typeName = "auth.common.ListResponse";
_ListResponse.fields = proto3.util.newFieldList(() => [
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
var _ListByCollectionRequest = class _ListByCollectionRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ListByCollectionRequest, a, b);
  }
};
_ListByCollectionRequest.runtime = proto3;
_ListByCollectionRequest.typeName = "auth.common.ListByCollectionRequest";
_ListByCollectionRequest.fields = proto3.util.newFieldList(() => [
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
var _ListByCollectionResponse = class _ListByCollectionResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ListByCollectionResponse, a, b);
  }
};
_ListByCollectionResponse.runtime = proto3;
_ListByCollectionResponse.typeName = "auth.common.ListByCollectionResponse";
_ListByCollectionResponse.fields = proto3.util.newFieldList(() => [
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
var _ListByAttrsRequest = class _ListByAttrsRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ListByAttrsRequest, a, b);
  }
};
_ListByAttrsRequest.runtime = proto3;
_ListByAttrsRequest.typeName = "auth.common.ListByAttrsRequest";
_ListByAttrsRequest.fields = proto3.util.newFieldList(() => [
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
var _ListByAttrsResponse = class _ListByAttrsResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ListByAttrsResponse, a, b);
  }
};
_ListByAttrsResponse.runtime = proto3;
_ListByAttrsResponse.typeName = "auth.common.ListByAttrsResponse";
_ListByAttrsResponse.fields = proto3.util.newFieldList(() => [
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
var _CreateRequest = class _CreateRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_CreateRequest, a, b);
  }
};
_CreateRequest.runtime = proto3;
_CreateRequest.typeName = "auth.common.CreateRequest";
_CreateRequest.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var CreateRequest = _CreateRequest;
var _CreateResponse = class _CreateResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_CreateResponse, a, b);
  }
};
_CreateResponse.runtime = proto3;
_CreateResponse.typeName = "auth.common.CreateResponse";
_CreateResponse.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var CreateResponse = _CreateResponse;
var _UpdateRequest = class _UpdateRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_UpdateRequest, a, b);
  }
};
_UpdateRequest.runtime = proto3;
_UpdateRequest.typeName = "auth.common.UpdateRequest";
_UpdateRequest.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item },
  { no: 3, name: "update_mask", kind: "message", T: FieldMask }
]);
var UpdateRequest = _UpdateRequest;
var _UpdateResponse = class _UpdateResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_UpdateResponse, a, b);
  }
};
_UpdateResponse.runtime = proto3;
_UpdateResponse.typeName = "auth.common.UpdateResponse";
_UpdateResponse.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var UpdateResponse = _UpdateResponse;
var _DeleteRequest = class _DeleteRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string reason = 4;
     */
    this.reason = "";
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_DeleteRequest, a, b);
  }
};
_DeleteRequest.runtime = proto3;
_DeleteRequest.typeName = "auth.common.DeleteRequest";
_DeleteRequest.fields = proto3.util.newFieldList(() => [
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
var _DeleteResponse = class _DeleteResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_DeleteResponse, a, b);
  }
};
_DeleteResponse.runtime = proto3;
_DeleteResponse.typeName = "auth.common.DeleteResponse";
_DeleteResponse.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "item", kind: "message", T: Item }
]);
var DeleteResponse = _DeleteResponse;
var _HistoryRequest = class _HistoryRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HistoryRequest, a, b);
  }
};
_HistoryRequest.runtime = proto3;
_HistoryRequest.typeName = "auth.common.HistoryRequest";
_HistoryRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item }
]);
var HistoryRequest = _HistoryRequest;
var _HistoryResponse = class _HistoryResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HistoryResponse, a, b);
  }
};
_HistoryResponse.runtime = proto3;
_HistoryResponse.typeName = "auth.common.HistoryResponse";
_HistoryResponse.fields = proto3.util.newFieldList(() => [
  { no: 2, name: "history", kind: "message", T: History }
]);
var HistoryResponse = _HistoryResponse;
var _HiddenTxRequest = class _HiddenTxRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HiddenTxRequest, a, b);
  }
};
_HiddenTxRequest.runtime = proto3;
_HiddenTxRequest.typeName = "auth.common.HiddenTxRequest";
_HiddenTxRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item }
]);
var HiddenTxRequest = _HiddenTxRequest;
var _HiddenTxResponse = class _HiddenTxResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HiddenTxResponse, a, b);
  }
};
_HiddenTxResponse.runtime = proto3;
_HiddenTxResponse.typeName = "auth.common.HiddenTxResponse";
_HiddenTxResponse.fields = proto3.util.newFieldList(() => [
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
var _HideTxRequest = class _HideTxRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HideTxRequest, a, b);
  }
};
_HideTxRequest.runtime = proto3;
_HideTxRequest.typeName = "auth.common.HideTxRequest";
_HideTxRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "hidden_tx", kind: "message", T: HiddenTx }
]);
var HideTxRequest = _HideTxRequest;
var _HideTxResponse = class _HideTxResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_HideTxResponse, a, b);
  }
};
_HideTxResponse.runtime = proto3;
_HideTxResponse.typeName = "auth.common.HideTxResponse";
_HideTxResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTxList }
]);
var HideTxResponse = _HideTxResponse;
var _UnHideTxRequest = class _UnHideTxRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string tx_id = 2;
     */
    this.txId = "";
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_UnHideTxRequest, a, b);
  }
};
_UnHideTxRequest.runtime = proto3;
_UnHideTxRequest.typeName = "auth.common.UnHideTxRequest";
_UnHideTxRequest.fields = proto3.util.newFieldList(() => [
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
var _UnHideTxResponse = class _UnHideTxResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_UnHideTxResponse, a, b);
  }
};
_UnHideTxResponse.runtime = proto3;
_UnHideTxResponse.typeName = "auth.common.UnHideTxResponse";
_UnHideTxResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item", kind: "message", T: Item },
  { no: 2, name: "hidden_txs", kind: "message", T: HiddenTxList }
]);
var UnHideTxResponse = _UnHideTxResponse;
var _ReferenceRequest = class _ReferenceRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceRequest, a, b);
  }
};
_ReferenceRequest.runtime = proto3;
_ReferenceRequest.typeName = "auth.common.ReferenceRequest";
_ReferenceRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "reference", kind: "message", T: ReferenceKey }
]);
var ReferenceRequest = _ReferenceRequest;
var _ReferenceResponse = class _ReferenceResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: bool exists = 1;
     */
    this.exists = false;
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceResponse, a, b);
  }
};
_ReferenceResponse.runtime = proto3;
_ReferenceResponse.typeName = "auth.common.ReferenceResponse";
_ReferenceResponse.fields = proto3.util.newFieldList(() => [
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
var _ReferenceByCollectionRequest = class _ReferenceByCollectionRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceByCollectionRequest, a, b);
  }
};
_ReferenceByCollectionRequest.runtime = proto3;
_ReferenceByCollectionRequest.typeName = "auth.common.ReferenceByCollectionRequest";
_ReferenceByCollectionRequest.fields = proto3.util.newFieldList(() => [
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
var _ReferenceByCollectionResponse = class _ReferenceByCollectionResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceByCollectionResponse, a, b);
  }
};
_ReferenceByCollectionResponse.runtime = proto3;
_ReferenceByCollectionResponse.typeName = "auth.common.ReferenceByCollectionResponse";
_ReferenceByCollectionResponse.fields = proto3.util.newFieldList(() => [
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
var _ReferenceByPartialKeyRequest = class _ReferenceByPartialKeyRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceByPartialKeyRequest, a, b);
  }
};
_ReferenceByPartialKeyRequest.runtime = proto3;
_ReferenceByPartialKeyRequest.typeName = "auth.common.ReferenceByPartialKeyRequest";
_ReferenceByPartialKeyRequest.fields = proto3.util.newFieldList(() => [
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
var _ReferenceByPartialKeyResponse = class _ReferenceByPartialKeyResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceByPartialKeyResponse, a, b);
  }
};
_ReferenceByPartialKeyResponse.runtime = proto3;
_ReferenceByPartialKeyResponse.typeName = "auth.common.ReferenceByPartialKeyResponse";
_ReferenceByPartialKeyResponse.fields = proto3.util.newFieldList(() => [
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
var _ReferenceByItemRequest = class _ReferenceByItemRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceByItemRequest, a, b);
  }
};
_ReferenceByItemRequest.runtime = proto3;
_ReferenceByItemRequest.typeName = "auth.common.ReferenceByItemRequest";
_ReferenceByItemRequest.fields = proto3.util.newFieldList(() => [
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
var _ReferenceByItemResponse = class _ReferenceByItemResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceByItemResponse, a, b);
  }
};
_ReferenceByItemResponse.runtime = proto3;
_ReferenceByItemResponse.typeName = "auth.common.ReferenceByItemResponse";
_ReferenceByItemResponse.fields = proto3.util.newFieldList(() => [
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
var _ReferenceCreateRequest = class _ReferenceCreateRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceCreateRequest, a, b);
  }
};
_ReferenceCreateRequest.runtime = proto3;
_ReferenceCreateRequest.typeName = "auth.common.ReferenceCreateRequest";
_ReferenceCreateRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceCreateRequest = _ReferenceCreateRequest;
var _ReferenceCreateResponse = class _ReferenceCreateResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceCreateResponse, a, b);
  }
};
_ReferenceCreateResponse.runtime = proto3;
_ReferenceCreateResponse.typeName = "auth.common.ReferenceCreateResponse";
_ReferenceCreateResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceCreateResponse = _ReferenceCreateResponse;
var _ReferenceDeleteRequest = class _ReferenceDeleteRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceDeleteRequest, a, b);
  }
};
_ReferenceDeleteRequest.runtime = proto3;
_ReferenceDeleteRequest.typeName = "auth.common.ReferenceDeleteRequest";
_ReferenceDeleteRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceDeleteRequest = _ReferenceDeleteRequest;
var _ReferenceDeleteResponse = class _ReferenceDeleteResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_ReferenceDeleteResponse, a, b);
  }
};
_ReferenceDeleteResponse.runtime = proto3;
_ReferenceDeleteResponse.typeName = "auth.common.ReferenceDeleteResponse";
_ReferenceDeleteResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "ref_key", kind: "message", T: ReferenceKey }
]);
var ReferenceDeleteResponse = _ReferenceDeleteResponse;
var _SuggestionRequest = class _SuggestionRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string suggestion_id = 2;
     */
    this.suggestionId = "";
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionRequest, a, b);
  }
};
_SuggestionRequest.runtime = proto3;
_SuggestionRequest.typeName = "auth.common.SuggestionRequest";
_SuggestionRequest.fields = proto3.util.newFieldList(() => [
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
var _SuggestionResponse = class _SuggestionResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionResponse, a, b);
  }
};
_SuggestionResponse.runtime = proto3;
_SuggestionResponse.typeName = "auth.common.SuggestionResponse";
_SuggestionResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionResponse = _SuggestionResponse;
var _SuggestionListRequest = class _SuggestionListRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionListRequest, a, b);
  }
};
_SuggestionListRequest.runtime = proto3;
_SuggestionListRequest.typeName = "auth.common.SuggestionListRequest";
_SuggestionListRequest.fields = proto3.util.newFieldList(() => [
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
var _SuggestionListResponse = class _SuggestionListResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionListResponse, a, b);
  }
};
_SuggestionListResponse.runtime = proto3;
_SuggestionListResponse.typeName = "auth.common.SuggestionListResponse";
_SuggestionListResponse.fields = proto3.util.newFieldList(() => [
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
var _SuggestionListByCollectionRequest = class _SuggestionListByCollectionRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionListByCollectionRequest, a, b);
  }
};
_SuggestionListByCollectionRequest.runtime = proto3;
_SuggestionListByCollectionRequest.typeName = "auth.common.SuggestionListByCollectionRequest";
_SuggestionListByCollectionRequest.fields = proto3.util.newFieldList(() => [
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
var _SuggestionListByCollectionResponse = class _SuggestionListByCollectionResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionListByCollectionResponse, a, b);
  }
};
_SuggestionListByCollectionResponse.runtime = proto3;
_SuggestionListByCollectionResponse.typeName = "auth.common.SuggestionListByCollectionResponse";
_SuggestionListByCollectionResponse.fields = proto3.util.newFieldList(() => [
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
var _SuggestionListByItemRequest = class _SuggestionListByItemRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionListByItemRequest, a, b);
  }
};
_SuggestionListByItemRequest.runtime = proto3;
_SuggestionListByItemRequest.typeName = "auth.common.SuggestionListByItemRequest";
_SuggestionListByItemRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "item_key", kind: "message", T: ItemKey }
]);
var SuggestionListByItemRequest = _SuggestionListByItemRequest;
var _SuggestionListByItemResponse = class _SuggestionListByItemResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated auth.Suggestion suggestions = 1;
     */
    this.suggestions = [];
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionListByItemResponse, a, b);
  }
};
_SuggestionListByItemResponse.runtime = proto3;
_SuggestionListByItemResponse.typeName = "auth.common.SuggestionListByItemResponse";
_SuggestionListByItemResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "suggestions", kind: "message", T: Suggestion, repeated: true }
]);
var SuggestionListByItemResponse = _SuggestionListByItemResponse;
var _SuggestionByPartialKeyRequest = class _SuggestionByPartialKeyRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionByPartialKeyRequest, a, b);
  }
};
_SuggestionByPartialKeyRequest.runtime = proto3;
_SuggestionByPartialKeyRequest.typeName = "auth.common.SuggestionByPartialKeyRequest";
_SuggestionByPartialKeyRequest.fields = proto3.util.newFieldList(() => [
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
var _SuggestionByPartialKeyResponse = class _SuggestionByPartialKeyResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionByPartialKeyResponse, a, b);
  }
};
_SuggestionByPartialKeyResponse.runtime = proto3;
_SuggestionByPartialKeyResponse.typeName = "auth.common.SuggestionByPartialKeyResponse";
_SuggestionByPartialKeyResponse.fields = proto3.util.newFieldList(() => [
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
var _SuggestionCreateRequest = class _SuggestionCreateRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionCreateRequest, a, b);
  }
};
_SuggestionCreateRequest.runtime = proto3;
_SuggestionCreateRequest.typeName = "auth.common.SuggestionCreateRequest";
_SuggestionCreateRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionCreateRequest = _SuggestionCreateRequest;
var _SuggestionCreateResponse = class _SuggestionCreateResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionCreateResponse, a, b);
  }
};
_SuggestionCreateResponse.runtime = proto3;
_SuggestionCreateResponse.typeName = "auth.common.SuggestionCreateResponse";
_SuggestionCreateResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionCreateResponse = _SuggestionCreateResponse;
var _SuggestionDeleteRequest = class _SuggestionDeleteRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionDeleteRequest, a, b);
  }
};
_SuggestionDeleteRequest.runtime = proto3;
_SuggestionDeleteRequest.typeName = "auth.common.SuggestionDeleteRequest";
_SuggestionDeleteRequest.fields = proto3.util.newFieldList(() => [
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
var _SuggestionDeleteResponse = class _SuggestionDeleteResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionDeleteResponse, a, b);
  }
};
_SuggestionDeleteResponse.runtime = proto3;
_SuggestionDeleteResponse.typeName = "auth.common.SuggestionDeleteResponse";
_SuggestionDeleteResponse.fields = proto3.util.newFieldList(() => [
  { no: 4, name: "suggestion", kind: "message", T: Suggestion }
]);
var SuggestionDeleteResponse = _SuggestionDeleteResponse;
var _SuggestionApproveRequest = class _SuggestionApproveRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionApproveRequest, a, b);
  }
};
_SuggestionApproveRequest.runtime = proto3;
_SuggestionApproveRequest.typeName = "auth.common.SuggestionApproveRequest";
_SuggestionApproveRequest.fields = proto3.util.newFieldList(() => [
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
var _SuggestionApproveResponse = class _SuggestionApproveResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SuggestionApproveResponse, a, b);
  }
};
_SuggestionApproveResponse.runtime = proto3;
_SuggestionApproveResponse.typeName = "auth.common.SuggestionApproveResponse";
_SuggestionApproveResponse.fields = proto3.util.newFieldList(() => [
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
  async call(service, method, data) {
    const headers = new Headers([]);
    headers.set("content-type", contentType);
    const response = await fetch(
      `${this.baseUrl}/${service}/${method}`,
      {
        method: "POST",
        headers,
        body: data.toJsonString()
      }
    );
    if (response.status === 200) {
      if (contentType === "application/json") {
        return await response.json();
      }
      return new Uint8Array(await response.arrayBuffer());
    }
    throw Error(`HTTP ${response.status} ${response.statusText}`);
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
  async getCurrentUser(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "GetCurrentUser",
      request
    );
    return promise.then(
      async (data) => GetCurrentUserResponse.fromJson(data)
    );
  }
  /**
   * ──────────────────────────────── Invoke ───────────────────────────────────────
   *
   * @generated from rpc auth.common.GenericService.Bootstrap
   */
  async bootstrap(request) {
    this.contract;
    const promise = this.request(
      "auth.common.GenericService",
      "Bootstrap",
      request.toJsonString({ typeRegistry: this.registry })
    );
    return promise.then(
      async (data) => BootstrapResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.AuthorizeOperation
   */
  async authorizeOperation(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "AuthorizeOperation",
      request
    );
    return promise.then(
      async (data) => AuthorizeOperationResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.CreateUser
   */
  async createUser(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "CreateUser",
      request
    );
    return promise.then(
      async (data) => CreateUserResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.CreateCollection
   */
  async createCollection(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "CreateCollection",
      request
    );
    return promise.then(
      async (data) => CreateCollectionResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.Get
   */
  async get(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "Get",
      request
    );
    return promise.then(
      async (data) => GetResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.List
   */
  async list(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "List",
      request
    );
    return promise.then(
      async (data) => ListResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.ListByCollection
   */
  async listByCollection(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "ListByCollection",
      request
    );
    return promise.then(
      async (data) => ListByCollectionResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.ListByAttrs
   */
  async listByAttrs(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "ListByAttrs",
      request
    );
    return promise.then(
      async (data) => ListByAttrsResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.Create
   */
  async create(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "Create",
      request
    );
    return promise.then(
      async (data) => CreateResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.Update
   */
  async update(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "Update",
      request
    );
    return promise.then(
      async (data) => UpdateResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.Delete
   */
  async delete(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "Delete",
      request
    );
    return promise.then(
      async (data) => DeleteResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.History
   */
  async history(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "History",
      request
    );
    return promise.then(
      async (data) => HistoryResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.HiddenTx
   */
  async hiddenTx(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "HiddenTx",
      request
    );
    return promise.then(
      async (data) => HiddenTxResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.HideTx
   */
  async hideTx(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "HideTx",
      request
    );
    return promise.then(
      async (data) => HideTxResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.UnHideTx
   */
  async unHideTx(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "UnHideTx",
      request
    );
    return promise.then(
      async (data) => UnHideTxResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.Reference
   */
  async reference(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "Reference",
      request
    );
    return promise.then(
      async (data) => ReferenceResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceByItem
   */
  async referenceByItem(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "ReferenceByItem",
      request
    );
    return promise.then(
      async (data) => ReferenceByItemResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceByPartialKey
   */
  async referenceByPartialKey(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "ReferenceByPartialKey",
      request
    );
    return promise.then(
      async (data) => ReferenceByPartialKeyResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceCreate
   */
  async referenceCreate(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "ReferenceCreate",
      request
    );
    return promise.then(
      async (data) => ReferenceCreateResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.ReferenceDelete
   */
  async referenceDelete(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "ReferenceDelete",
      request
    );
    return promise.then(
      async (data) => ReferenceDeleteResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.Suggestion
   */
  async suggestion(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "Suggestion",
      request
    );
    return promise.then(
      async (data) => SuggestionResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionListByCollection
   */
  async suggestionListByCollection(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "SuggestionListByCollection",
      request
    );
    return promise.then(
      async (data) => SuggestionListByCollectionResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionByPartialKey
   */
  async suggestionByPartialKey(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "SuggestionByPartialKey",
      request
    );
    return promise.then(
      async (data) => SuggestionByPartialKeyResponse.fromJson(data)
    );
  }
  /**
   * ──────────────────────────────── Invoke ───────────────────────────────────────
   *
   * @generated from rpc auth.common.GenericService.SuggestionCreate
   */
  async suggestionCreate(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "SuggestionCreate",
      request
    );
    return promise.then(
      async (data) => SuggestionCreateResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionDelete
   */
  async suggestionDelete(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "SuggestionDelete",
      request
    );
    return promise.then(
      async (data) => SuggestionDeleteResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc auth.common.GenericService.SuggestionApprove
   */
  async suggestionApprove(request) {
    const promise = this.request(
      "auth.common.GenericService",
      "SuggestionApprove",
      request
    );
    return promise.then(
      async (data) => SuggestionApproveResponse.fromJson(data)
    );
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
var _Specimen = class _Specimen extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen, a, b);
  }
};
_Specimen.runtime = proto3;
_Specimen.typeName = "ccbio.schema.v0.Specimen";
_Specimen.fields = proto3.util.newFieldList(() => [
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
var _Specimen_Primary = class _Specimen_Primary extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Primary, a, b);
  }
};
_Specimen_Primary.runtime = proto3;
_Specimen_Primary.typeName = "ccbio.schema.v0.Specimen.Primary";
_Specimen_Primary.fields = proto3.util.newFieldList(() => [
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
  { no: 8, name: "field_date", kind: "message", T: Timestamp },
  { no: 9, name: "catalog_date", kind: "message", T: Timestamp },
  { no: 10, name: "determined_date", kind: "message", T: Timestamp },
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
var _Specimen_Secondary = class _Specimen_Secondary extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Secondary, a, b);
  }
};
_Specimen_Secondary.runtime = proto3;
_Specimen_Secondary.typeName = "ccbio.schema.v0.Specimen.Secondary";
_Specimen_Secondary.fields = proto3.util.newFieldList(() => [
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
var _Specimen_Taxon = class _Specimen_Taxon extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Taxon, a, b);
  }
};
_Specimen_Taxon.runtime = proto3;
_Specimen_Taxon.typeName = "ccbio.schema.v0.Specimen.Taxon";
_Specimen_Taxon.fields = proto3.util.newFieldList(() => [
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
var _Specimen_Georeference = class _Specimen_Georeference extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Georeference, a, b);
  }
};
_Specimen_Georeference.runtime = proto3;
_Specimen_Georeference.typeName = "ccbio.schema.v0.Specimen.Georeference";
_Specimen_Georeference.fields = proto3.util.newFieldList(() => [
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
var _Specimen_Image = class _Specimen_Image extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Image, a, b);
  }
};
_Specimen_Image.runtime = proto3;
_Specimen_Image.typeName = "ccbio.schema.v0.Specimen.Image";
_Specimen_Image.fields = proto3.util.newFieldList(() => [
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
var _Specimen_Loan = class _Specimen_Loan extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Loan, a, b);
  }
};
_Specimen_Loan.runtime = proto3;
_Specimen_Loan.typeName = "ccbio.schema.v0.Specimen.Loan";
_Specimen_Loan.fields = proto3.util.newFieldList(() => [
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
  { no: 5, name: "loaned_date", kind: "message", T: Timestamp },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Loan = _Specimen_Loan;
var _Specimen_Grant = class _Specimen_Grant extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Specimen_Grant, a, b);
  }
};
_Specimen_Grant.runtime = proto3;
_Specimen_Grant.typeName = "ccbio.schema.v0.Specimen.Grant";
_Specimen_Grant.fields = proto3.util.newFieldList(() => [
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
  { no: 5, name: "granted_date", kind: "message", T: Timestamp },
  { no: 20, name: "last_modified", kind: "message", T: StateActivity }
]);
var Specimen_Grant = _Specimen_Grant;

// gen/chaincode/ccbio/schema/v0/service_pb.ts
var _SpecimenGetRequest = class _SpecimenGetRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetRequest, a, b);
  }
};
_SpecimenGetRequest.runtime = proto3;
_SpecimenGetRequest.typeName = "ccbio.schema.v0.SpecimenGetRequest";
_SpecimenGetRequest.fields = proto3.util.newFieldList(() => [
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
var _SpecimenGetResponse = class _SpecimenGetResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetResponse, a, b);
  }
};
_SpecimenGetResponse.runtime = proto3;
_SpecimenGetResponse.typeName = "ccbio.schema.v0.SpecimenGetResponse";
_SpecimenGetResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenGetResponse = _SpecimenGetResponse;
var _SpecimenGetListRequest = class _SpecimenGetListRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetListRequest, a, b);
  }
};
_SpecimenGetListRequest.runtime = proto3;
_SpecimenGetListRequest.typeName = "ccbio.schema.v0.SpecimenGetListRequest";
_SpecimenGetListRequest.fields = proto3.util.newFieldList(() => [
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
var _SpecimenGetListResponse = class _SpecimenGetListResponse extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetListResponse, a, b);
  }
};
_SpecimenGetListResponse.runtime = proto3;
_SpecimenGetListResponse.typeName = "ccbio.schema.v0.SpecimenGetListResponse";
_SpecimenGetListResponse.fields = proto3.util.newFieldList(() => [
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
var _SpecimenGetByCollectionRequest = class _SpecimenGetByCollectionRequest extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: string collection_id = 1;
     */
    this.collectionId = "";
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetByCollectionRequest, a, b);
  }
};
_SpecimenGetByCollectionRequest.runtime = proto3;
_SpecimenGetByCollectionRequest.typeName = "ccbio.schema.v0.SpecimenGetByCollectionRequest";
_SpecimenGetByCollectionRequest.fields = proto3.util.newFieldList(() => [
  {
    no: 1,
    name: "collection_id",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var SpecimenGetByCollectionRequest = _SpecimenGetByCollectionRequest;
var _SpecimenGetByCollectionResponse = class _SpecimenGetByCollectionResponse extends Message {
  constructor(data) {
    super();
    /**
     * @generated from field: repeated ccbio.schema.v0.Specimen specimens = 1;
     */
    this.specimens = [];
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetByCollectionResponse, a, b);
  }
};
_SpecimenGetByCollectionResponse.runtime = proto3;
_SpecimenGetByCollectionResponse.typeName = "ccbio.schema.v0.SpecimenGetByCollectionResponse";
_SpecimenGetByCollectionResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimens", kind: "message", T: Specimen, repeated: true }
]);
var SpecimenGetByCollectionResponse = _SpecimenGetByCollectionResponse;
var _SpecimenGetHistoryRequest = class _SpecimenGetHistoryRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetHistoryRequest, a, b);
  }
};
_SpecimenGetHistoryRequest.runtime = proto3;
_SpecimenGetHistoryRequest.typeName = "ccbio.schema.v0.SpecimenGetHistoryRequest";
_SpecimenGetHistoryRequest.fields = proto3.util.newFieldList(() => [
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
var _SpecimenGetHistoryResponse = class _SpecimenGetHistoryResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenGetHistoryResponse, a, b);
  }
};
_SpecimenGetHistoryResponse.runtime = proto3;
_SpecimenGetHistoryResponse.typeName = "ccbio.schema.v0.SpecimenGetHistoryResponse";
_SpecimenGetHistoryResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "history", kind: "message", T: StateActivity }
]);
var SpecimenGetHistoryResponse = _SpecimenGetHistoryResponse;
var _SpecimenCreateRequest = class _SpecimenCreateRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenCreateRequest, a, b);
  }
};
_SpecimenCreateRequest.runtime = proto3;
_SpecimenCreateRequest.typeName = "ccbio.schema.v0.SpecimenCreateRequest";
_SpecimenCreateRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenCreateRequest = _SpecimenCreateRequest;
var _SpecimenCreateResponse = class _SpecimenCreateResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenCreateResponse, a, b);
  }
};
_SpecimenCreateResponse.runtime = proto3;
_SpecimenCreateResponse.typeName = "ccbio.schema.v0.SpecimenCreateResponse";
_SpecimenCreateResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenCreateResponse = _SpecimenCreateResponse;
var _SpecimenUpdateRequest = class _SpecimenUpdateRequest extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenUpdateRequest, a, b);
  }
};
_SpecimenUpdateRequest.runtime = proto3;
_SpecimenUpdateRequest.typeName = "ccbio.schema.v0.SpecimenUpdateRequest";
_SpecimenUpdateRequest.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen },
  { no: 2, name: "update_mask", kind: "message", T: FieldMask }
]);
var SpecimenUpdateRequest = _SpecimenUpdateRequest;
var _SpecimenUpdateResponse = class _SpecimenUpdateResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenUpdateResponse, a, b);
  }
};
_SpecimenUpdateResponse.runtime = proto3;
_SpecimenUpdateResponse.typeName = "ccbio.schema.v0.SpecimenUpdateResponse";
_SpecimenUpdateResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen },
  { no: 2, name: "update_mask", kind: "message", T: FieldMask }
]);
var SpecimenUpdateResponse = _SpecimenUpdateResponse;
var _SpecimenDeleteRequest = class _SpecimenDeleteRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenDeleteRequest, a, b);
  }
};
_SpecimenDeleteRequest.runtime = proto3;
_SpecimenDeleteRequest.typeName = "ccbio.schema.v0.SpecimenDeleteRequest";
_SpecimenDeleteRequest.fields = proto3.util.newFieldList(() => [
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
var _SpecimenDeleteResponse = class _SpecimenDeleteResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenDeleteResponse, a, b);
  }
};
_SpecimenDeleteResponse.runtime = proto3;
_SpecimenDeleteResponse.typeName = "ccbio.schema.v0.SpecimenDeleteResponse";
_SpecimenDeleteResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenDeleteResponse = _SpecimenDeleteResponse;
var _SpecimenHideTxRequest = class _SpecimenHideTxRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenHideTxRequest, a, b);
  }
};
_SpecimenHideTxRequest.runtime = proto3;
_SpecimenHideTxRequest.typeName = "ccbio.schema.v0.SpecimenHideTxRequest";
_SpecimenHideTxRequest.fields = proto3.util.newFieldList(() => [
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
var _SpecimenHideTxResponse = class _SpecimenHideTxResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenHideTxResponse, a, b);
  }
};
_SpecimenHideTxResponse.runtime = proto3;
_SpecimenHideTxResponse.typeName = "ccbio.schema.v0.SpecimenHideTxResponse";
_SpecimenHideTxResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenHideTxResponse = _SpecimenHideTxResponse;
var _SpecimenUnHideTxRequest = class _SpecimenUnHideTxRequest extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenUnHideTxRequest, a, b);
  }
};
_SpecimenUnHideTxRequest.runtime = proto3;
_SpecimenUnHideTxRequest.typeName = "ccbio.schema.v0.SpecimenUnHideTxRequest";
_SpecimenUnHideTxRequest.fields = proto3.util.newFieldList(() => [
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
var _SpecimenUnHideTxResponse = class _SpecimenUnHideTxResponse extends Message {
  constructor(data) {
    super();
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_SpecimenUnHideTxResponse, a, b);
  }
};
_SpecimenUnHideTxResponse.runtime = proto3;
_SpecimenUnHideTxResponse.typeName = "ccbio.schema.v0.SpecimenUnHideTxResponse";
_SpecimenUnHideTxResponse.fields = proto3.util.newFieldList(() => [
  { no: 1, name: "specimen", kind: "message", T: Specimen }
]);
var SpecimenUnHideTxResponse = _SpecimenUnHideTxResponse;

// gen/chaincode/ccbio/schema/v0/service_gateway.ts
var SpecimenServiceClient = class {
  constructor(contract, registry) {
    this.registry = "";
    this.contract = contract;
  }
  async call(service, method, data) {
    const headers = new Headers([]);
    headers.set("content-type", contentType);
    const response = await fetch(
      `${this.baseUrl}/${service}/${method}`,
      {
        method: "POST",
        headers,
        body: data.toJsonString()
      }
    );
    if (response.status === 200) {
      if (contentType === "application/json") {
        return await response.json();
      }
      return new Uint8Array(await response.arrayBuffer());
    }
    throw Error(`HTTP ${response.status} ${response.statusText}`);
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGet
   */
  async specimenGet(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenGet",
      request
    );
    return promise.then(
      async (data) => SpecimenGetResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGetList
   */
  async specimenGetList(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenGetList",
      request
    );
    return promise.then(
      async (data) => SpecimenGetListResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGetByCollection
   */
  async specimenGetByCollection(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenGetByCollection",
      request
    );
    return promise.then(
      async (data) => SpecimenGetByCollectionResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenGetHistory
   */
  async specimenGetHistory(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenGetHistory",
      request
    );
    return promise.then(
      async (data) => SpecimenGetHistoryResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenCreate
   */
  async specimenCreate(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenCreate",
      request
    );
    return promise.then(
      async (data) => SpecimenCreateResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenUpdate
   */
  async specimenUpdate(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenUpdate",
      request
    );
    return promise.then(
      async (data) => SpecimenUpdateResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenDelete
   */
  async specimenDelete(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenDelete",
      request
    );
    return promise.then(
      async (data) => SpecimenDeleteResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenHideTx
   */
  async specimenHideTx(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenHideTx",
      request
    );
    return promise.then(
      async (data) => SpecimenHideTxResponse.fromJson(data)
    );
  }
  /**
   * @generated from rpc ccbio.schema.v0.SpecimenService.SpecimenUnHideTx
   */
  async specimenUnHideTx(request) {
    const promise = this.request(
      "ccbio.schema.v0.SpecimenService",
      "SpecimenUnHideTx",
      request
    );
    return promise.then(
      async (data) => SpecimenUnHideTxResponse.fromJson(data)
    );
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
var _Item2 = class _Item2 extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Item2, a, b);
  }
};
_Item2.runtime = proto3;
_Item2.typeName = "sample.Item";
_Item2.fields = proto3.util.newFieldList(() => [
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
var _Group = class _Group extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Group, a, b);
  }
};
_Group.runtime = proto3;
_Group.typeName = "sample.Group";
_Group.fields = proto3.util.newFieldList(() => [
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
var _Book = class _Book extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Book, a, b);
  }
};
_Book.runtime = proto3;
_Book.typeName = "sample.Book";
_Book.fields = proto3.util.newFieldList(() => [
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
var _Degree = class _Degree extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Degree, a, b);
  }
};
_Degree.runtime = proto3;
_Degree.typeName = "sample.Degree";
_Degree.fields = proto3.util.newFieldList(() => [
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
  { no: 3, name: "degree_date", kind: "message", T: Timestamp }
]);
var Degree = _Degree;
var _Awards = class _Awards extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Awards, a, b);
  }
};
_Awards.runtime = proto3;
_Awards.typeName = "sample.Awards";
_Awards.fields = proto3.util.newFieldList(() => [
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
  { no: 3, name: "award_date", kind: "message", T: Timestamp },
  {
    no: 4,
    name: "award_description",
    kind: "scalar",
    T: 9
    /* ScalarType.STRING */
  }
]);
var Awards = _Awards;
var _Author = class _Author extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Author, a, b);
  }
};
_Author.runtime = proto3;
_Author.typeName = "sample.Author";
_Author.fields = proto3.util.newFieldList(() => [
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
var _Person = class _Person extends Message {
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
    proto3.util.initPartial(data, this);
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
    return proto3.util.equals(_Person, a, b);
  }
};
_Person.runtime = proto3;
_Person.typeName = "sample.Person";
_Person.fields = proto3.util.newFieldList(() => [
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

// utils/registry.ts
var registry_exports = {};
__export(registry_exports, {
  Registry: () => Registry
});
var Registry = createRegistry(
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
  obj.value = Any.pack(author);
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
