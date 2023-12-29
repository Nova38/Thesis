import {
    Schema,
    findCustomMessageOption,
} from "@bufbuild/protoplugin/ecmascript";
import { getAllTypes } from "../utils";

import { KeySchema } from "../gen/auth/v1/auth_pb";
//?^

export function generateKeySchema(schema: Schema) {
    for (const file of schema.files) {
        const f = schema.generateFile(file.name + "_pb_key.ts");
        // const createRegistry = f.import("createRegistry", "@bufbuild/protobuf");
        const { Message, MessageType, JsonValue } = schema.runtime;
        const MessageAsType = Message.toTypeOnly();

        f.print`export const MessageKeySchema = {`;
        for (const descType of getAllTypes(file)) {
            if (descType.kind == "message") {
                const key = findCustomMessageOption(descType, 54599, KeySchema);
                if (key) {
                    f.print`  "${descType.typeName}" : ${key.toJsonString({
                        enumAsInteger: true,
                        prettySpaces: 4,
                        emitDefaultValues: true,
                    })},`;
                    // f.print`${descType.typeName}: ${key},`;
                }
            }
        }
        f.print`};`;
    }
}
