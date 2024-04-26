"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ToItem = void 0;
var protobuf_1 = require("@bufbuild/protobuf");
var saacs_pb_1 = require("@saacs/saacs-pb");
function ToItem(m) {
    return new saacs_pb_1.pb.Item({
        key: saacs_pb_1.key.PrimaryToKeySchema(m),
        value: protobuf_1.Any.pack(m),
    });
}
exports.ToItem = ToItem;
