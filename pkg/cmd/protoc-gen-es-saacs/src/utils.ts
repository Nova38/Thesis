import {
    DescEnum,
    DescExtension,
    DescFile,
    DescMessage,
    DescService,
    createRegistryFromDescriptors,
} from "@bufbuild/protobuf";
import * as fs from "fs";

export const registry = createRegistryFromDescriptors(
    fs.readFileSync("image.bin")
);
export function* getAllTypes(
    desc: DescFile | DescMessage
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
function* getAllMessages(
    desc: DescFile | DescMessage
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
