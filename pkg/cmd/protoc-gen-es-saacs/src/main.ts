#!/usr/bin/env -S npx tsx

// See here how to run this plugin: https://github.com/bufbuild/protobuf-es/tree/main/pkg/protoplugin-example

import { createEcmaScriptPlugin, runNodeJs } from "@bufbuild/protoplugin";
import {
    ImportSymbol,
    Schema,
    findCustomMessageOption,
} from "@bufbuild/protoplugin/ecmascript";
import type {
    AnyMessage,
    DescMessage,
    FileDescriptorSet,
} from "@bufbuild/protobuf";
import { localName } from "@bufbuild/protoplugin/ecmascript";

import { Empty, MethodKind, createDescriptorSet } from "@bufbuild/protobuf";

import { pb } from "saacs-es";

import { registry } from "./utils";
import { generateRegistry } from "./generators/registry";
import { generateIndex } from "./generators/indexs";
import { generateGateway } from "./generators/gateway";
import { generateKeySchema } from "./generators/keys";

const protocGenReg = createEcmaScriptPlugin({
    name: "protoc-gen-reg",
    version: `v1`,
    generateTs,
    //   parseOptions:
});

runNodeJs(protocGenReg);

function generateTs(schema: Schema) {
    generateGateway(schema);
    generateFabricTest(schema);

    generateRegistry(schema);
    generateIndex(schema);
    generateKeySchema(schema);
}

function populateFields(message: AnyMessage): AnyMessage {
    for (const fieldInfo of message.getType().fields.byNumber()) {
        if (fieldInfo.kind == "message") {
            const subType = registry.findMessage(fieldInfo.T.typeName);

            if (!subType) {
                return new Empty();
            }
            const subMesg = new subType();
            message[fieldInfo.localName] = populateFields(subMesg);
        }
        //   const value = message[fieldInfo.localName];
        //   console.log(`field ${fieldInfo.localName}: ${value}`);
    }

    return message;
}

// for (const file of schema.files) {
//     const f = schema.generateFile(file.name + "_key.ts");
//     f.preamble(file);
//     for (const descType of getAllTypes(file)) {
//         if (descType.kind == "message") {
//             f.print`// ${descType.name} \n`

//             const keySchema = getKeySchema(schema)
//             if (!keySchema) {
//                 return
//             }

//             const options = findCustomMessageOption(descType, 54599, authpb.KeySchema )

//             const {
//                 Message,
//                 JsonValue
//             } = schema.runtime;
//             f.print`// ${Message}`;

//             if (options?.itemKind == authpb.ItemKind.PRIMARY_ITEM) {
//                 f.print`// Primary Item:  ${descType.name}\n`
//                 // f.print`// }`

//                 // ItemKey
//                 // KeyAttributes
//                 // SetKeyAttr
//                 descType.fields.forEach((field) => {
//                     f.print`// name${field.name} ${field.kind} }`

//                 })
//                         options?.keys?.paths?.forEach((path) => {
//                     f.print`    // ${path}   ,`
//                 })

//                 let keyPaths = descType.fields.filter((field) => {
//                     if (options?.keys?.paths?.includes(field.name)) {
//                         return field.name
//                     }
//                 })

//                 f.print`export function ${descType}Key(item : ${descType}): string[] {`
//                 f.print`    attr=[]`

//                 keyPaths.forEach((path) => {
//                     const name = path.proto.jsonName || ""
//                     f.print` if (!item?.${name}) {`
//                     f.print`    return attr`
//                     f.print` }`
//                     f.print`    attr.push(item?.${name})`
//                 })
//                 f.print` return attr`
//                 f.print`}`

// options?.keys?.paths?.forEach((path) => {
//     f.print`// Path: ${path}\n`
// })

// }

//             }
//         }

//     }
// }

function GetKeySchema(message: DescMessage) {
    const options = findCustomMessageOption(
        message,
        54599,
        pb.auth.auth.KeySchema
    );

    if (options) {
        return options;
    }

    return undefined;
}

// function getTransactionType(message: DescMessage) {
//     const options = findCustomMessageOption(
//         message,
//         50556,
//         pb.auth.auth.TransactionType
//     );

//     if (options) {
//         return options;
//     }

//     return undefined;
// }

function generateFabricTest(schema: Schema) {
    for (const file of schema.files) {
        const f = schema.generateFile(file.name + ".fabric");
        // const registry = createRegistryFromDescriptors(createDescriptorSet(schema.proto.sourceFileDescriptors) )

        f.preamble(file);
        const { Message, JsonValue } = schema.runtime;
        // Convert the Message ImportSymbol to a type-only ImportSymbol
        const MessageAsType = Message.toTypeOnly();

        f.print`[`;

        for (const service of file.services) {
            const localServiceName = localName(service);

            for (const method of service.methods) {
                let x = "";
                const inputType = registry.findMessage(method.input.typeName);

                if (inputType) {
                    f.print`// ${method.name} ${inputType.name}`;
                    const input = new inputType();
                    const populated = populateFields(input);
                    x = populated.toJsonString({ emitDefaultValues: false });
                } else {
                    x = "{}";
                }

                if (method.methodKind === MethodKind.Unary) {
                    f.print`    {`;
                    f.print`        "invoke": "${localName(method)}",`;
                    f.print`        "args": [`;
                    f.print`            ${x}`;
                    // (request: ${method.input}, evaluate: bool ): Promise< ${method.output}> {`;
                    f.print`        ]`;
                    f.print`    },`;
                }
            }
        }
        f.print`]`;
    }
}
