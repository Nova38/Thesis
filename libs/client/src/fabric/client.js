"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createPromiseClient = exports.createUtilGateway = exports.createBiochainGateway = void 0;
var protobuf_1 = require("@bufbuild/protobuf");
var saacs_pb_1 = require("@saacs/saacs-pb");
function createBiochainGateway(contract) {
    return createPromiseClient(saacs_pb_1.chaincode.chaincode.ItemService, contract);
}
exports.createBiochainGateway = createBiochainGateway;
function createUtilGateway(contract) {
    return createPromiseClient(saacs_pb_1.chaincode.utils.UtilsService, contract);
}
exports.createUtilGateway = createUtilGateway;
/**
 * Create a PromiseClient for the given service, invoking RPCs through the
 * given transport.
 */
function createPromiseClient(service, contract) {
    return Object.fromEntries(Object.entries(service.methods).map(function (_a) {
        var name = _a[0], method = _a[1];
        return [
            name,
            createContractFn(contract, method),
        ];
    }));
}
exports.createPromiseClient = createPromiseClient;
function createContractFn(contract, method) {
    var utf8Decoder = new TextDecoder();
    return function (input, options) {
        return __awaiter(this, void 0, void 0, function () {
            var params, decode;
            var _a;
            return __generator(this, function (_b) {
                params = ((_a = options === null || options === void 0 ? void 0 : options.serializer) !== null && _a !== void 0 ? _a : 'json') === 'json'
                    ? new method.I(input).toJsonString({ emitDefaultValues: true, typeRegistry: saacs_pb_1.GlobalRegistry })
                    : new method.I(input).toBinary();
                decode = function (reply) {
                    var _a;
                    return ((_a = options === null || options === void 0 ? void 0 : options.serializer) !== null && _a !== void 0 ? _a : 'json') === 'json'
                        ? method.O.fromJsonString(utf8Decoder.decode(reply), { typeRegistry: saacs_pb_1.GlobalRegistry })
                        : method.O.fromBinary(reply);
                };
                console.log('params', params);
                try {
                    if (method.idempotency === protobuf_1.MethodIdempotency.Idempotent
                        || method.idempotency === protobuf_1.MethodIdempotency.NoSideEffects) {
                        return [2 /*return*/, contract.evaluateTransaction(method.name, params).then(function (e) {
                                console.log(e);
                                return decode(e);
                            })];
                    }
                    return [2 /*return*/, contract.submitTransaction(method.name, params).then(decode)];
                }
                catch (error) {
                    console.error('Error:', error);
                    throw error;
                }
                return [2 /*return*/];
            });
        });
    };
}
