import {
    DescEnum,
    DescExtension,
    DescFile,
    DescMessage,
    DescService,
    createRegistryFromDescriptors,
} from "@bufbuild/protobuf";
import { Schema } from "@bufbuild/protoplugin/ecmascript";
import * as fs from "fs";
import { fileURLToPath } from "url";

import { resolve, dirname } from "pathe";

const __dirname = dirname(fileURLToPath(import.meta.url));
const registryBin = resolve(__dirname, "image.bin");

export const registry = createRegistryFromDescriptors(
    fs.readFileSync(registryBin),
);

export function* GetAllTypesFromSchema(schema: Schema) {
    for (const file of schema.files) {
        for (const descType of getAllTypes(file)) {
            yield descType;
        }
    }
}

export function* getAllTypes(
    desc: DescFile | DescMessage,
): Iterable<DescMessage | DescEnum | DescExtension | DescService> {
    switch (desc.kind) {
        case "file":
            for (const message of desc.messages) {
                yield message;
                yield* getAllTypes(message);
            }
            yield* desc.enums;
            yield* desc.services;
            yield* desc.extensions;
            break;
        case "message":
            for (const message of desc.nestedMessages) {
                yield message;
                yield* getAllTypes(message);
            }
            yield* desc.nestedEnums;
            yield* desc.nestedExtensions;
            break;
    }
}
export function* getAllMessages(
    desc: DescFile | DescMessage,
): Iterable<DescMessage> {
    switch (desc.kind) {
        case "file":
            for (const message of desc.messages) {
                yield message;
                yield* getAllMessages(message);
            }
            break;
        case "message":
            for (const message of desc.nestedMessages) {
                yield message;
                yield* getAllMessages(message);
            }
            break;
    }
}
