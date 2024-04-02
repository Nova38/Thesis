import {
    Schema,
    findCustomMessageOption,
} from "@bufbuild/protoplugin/ecmascript";
import { getAllMessages, getAllTypes } from "../utils";

import { KeySchema, key_schema } from "../gen/auth/v1/auth_pb";
import { getExtension, hasExtension } from "@bufbuild/protobuf";

export function generateKeySchema(schema: Schema) {
    for (const file of schema.files) {
        const f = schema.generateFile(file.name + "_pb_key.ts");
        // const createRegistry = f.import("createRegistry", "@bufbuild/protobuf");
        const { Message, MessageType, JsonValue } = schema.runtime;
        const MessageAsType = Message.toTypeOnly();

        f.print`export const MessageKeySchema = {`;
        for (const descType of getAllMessages(file)) {
            const options = descType.proto.options;

            if (options && hasExtension(options, key_schema)) {
                const key = getExtension(options, key_schema);
                key.itemType = descType.typeName;
                f.print`  "${descType.typeName}" : ${key.toJsonString({
                    enumAsInteger: true,
                    prettySpaces: 4,
                    emitDefaultValues: true,
                })},`;
            }
        }
        f.print`};`;
    }
}
