#!/usr/bin/env -S npx tsx

// See here how to run this plugin: https://github.com/bufbuild/protobuf-es/tree/main/packages/protoplugin-example

import { createEcmaScriptPlugin, runNodeJs } from "@bufbuild/protoplugin";
import {
    ImportSymbol,
    Schema,
    findCustomMessageOption,
} from "@bufbuild/protoplugin/ecmascript";
import type {
    AnyMessage,
    DescEnum,
    DescExtension,
    DescFile,
    DescMessage,
    DescService,
    FileDescriptorSet,
} from "@bufbuild/protobuf";
import {
    literalString,
    makeJsDoc,
    localName,
} from "@bufbuild/protoplugin/ecmascript";

import {
    Empty,
    MethodKind,
    createDescriptorSet,
    createRegistryFromDescriptors,
} from "@bufbuild/protobuf";

import { pb } from "saacs-es";

import * as fs from "fs";

const registry = createRegistryFromDescriptors(fs.readFileSync("image.bin"));

const protocGenReg = createEcmaScriptPlugin({
    name: "protoc-gen-reg",
    version: `v1`,
    generateTs,
    //   parseOptions:
});

runNodeJs(protocGenReg);

function* getAllTypes(
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

function generateTs(schema: Schema) {
    // generateGateway(schema);
    generateFabricTest(schema);

    generateRegistry(schema);
    generateIndex(schema);
}

function generateIndex(schema: Schema) {
    // export * as auth from './auth/v1/index.js'
    // export * as common from './chaincode/common/index.js'
    // export * as ccbio from './biochain/v1/index.js'
    // export * as sample from './sample/v0/index.js'

    // let baseIndexFile = schema.generateFile("index.ts");
    // b.print(`export * from "./index_pb_reg.js"`)
    // baseIndexFile.print(`// ${schema.files.length} files`)

    const folders = new Map<string, string[]>();

    for (const file of schema.files) {
        // const folder = file.name.split("/").slice(0, -1)
        const p = file.name.split("/");
        const base = p.pop() || "";
        const folder = p.join("/") + "/";

        if (folders.has(folder)) {
            folders.get(folder)?.push(base);
        } else {
            folders.set(folder, [base]);
        }

        const f = schema.generateFile(folder + "index_" + base + "_pb.ts");

        // f.preamble(file);
        if (file.messages.length > 0) {
            f.print(`export * from "./${base}_pb.js"`);
            f.print(`export * from "./${base}_pb_reg.js"`);
        }

        // if (file.services.length > 0) {
        //     f.print(`export * from "./${base}_pb_gateway.js"`);
        // }
    }

    folders.forEach((files, folder) => {
        const f = schema.generateFile(folder + "index.ts");

        if (files.length == 1) {
            const file = files[0];
            f.print(`export * from "./index_${file}_pb.js"`);
        } else {
            files.forEach((file) => {
                f.print(`export * as ${file} from "./index_${file}_pb.js"`);
            });
        }

        // folder.split("/").forEach((file) => {
        //     if (file != "" && !file.match(/^v[0-9]/) && !file.match(/^chaincode/)) {
        //         baseIndexFile.print(`export * as ${folder} from "./${folder}index.js"`)
        //     }

        // // baseIndexFile.print(`export * as ${files} from "./${folder}index.js"`)
        // })
    });

    //  // remove the last part of the path
    //  const f = schema.generateFile(folder + "index.ts");
    //  f.preamble(file);

    //  f.print`//${folder}`
}

function populateFeilds(message: AnyMessage): AnyMessage {
    for (const fieldInfo of message.getType().fields.byNumber()) {
        if (fieldInfo.kind == "message") {
            const subType = registry.findMessage(fieldInfo.T.typeName);

            if (!subType) {
                return new Empty();
            }
            const subMesg = new subType();
            message[fieldInfo.localName] = populateFeilds(subMesg);
        }
        //   const value = message[fieldInfo.localName];
        //   console.log(`field ${fieldInfo.localName}: ${value}`);
    }

    return message;
}

// function genTypes(schema: Schema) {
//     let des: DescMessage[] | undefined;

//     for (const file of schema.files) {
//         for (const descType of getAllTypes(file)) {
//             if (descType.kind == "message") {
//                 if (des == undefined) {
//                     des = [descType]
//                 } else {
//                     des.push(descType)
//                 }
//             }
//         }
//     }

//     let f = schema.generateFile("types_pb.ts");

//     if (des == undefined) {
//         return
//     }

//     for (const msg of des) {
//         f.print(`export class ${msg.name} extends Message {`)
//         for (const field of msg.fields) {
//             f.print(`    ${field.name}: ${field.kind};`)
//         }
//         f.print(`}`)
//     }
// }

function generateRegistry(schema: Schema) {
    for (const file of schema.files) {
        const f = schema.generateFile(file.name + "_pb_reg.ts");
        // f.preamble(file);
        const createRegistry = f.import("createRegistry", "@bufbuild/protobuf");
        const { Message, MessageType, JsonValue } = schema.runtime;
        // Convert the Message ImportSymbol to a type-only ImportSymbol
        const MessageAsType = Message.toTypeOnly();

        f.print`export const allMessages = [`;
        for (const descType of getAllTypes(file)) {
            if (descType.kind == "message") {
                f.print`  ${descType}, `;
            }
        }
        f.print`];`;

        // f.print("export const registry = ", createRegistry, "(");
        // for (const descType of getAllTypes(file)) {
        //   if (descType.kind == "message") {
        //     f.print("  ", descType, ", ");
        //   }
        // }
        // f.print(");");
    }
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

// Modified from the original protoc-gen-ts plugin
function generateGateway(schema: Schema) {
    for (const file of schema.files) {
        const f = schema.generateFile(file.name + "_pb_gateway.ts");

        const contract = f.import("Contract", "@hyperledger/fabric-gateway");
        const registry = f.import("IMessageTypeRegistry", "@bufbuild/protobuf");
        const jsonWriteOptions = f.import(
            "JsonWriteStringOptions",
            "@bufbuild/protobuf"
        );
        // const jsonReadOptions = f.import(
        //     "JsonReadOptions",
        //     "@bufbuild/protobuf"
        // );

        f.print("const utf8Decoder = new TextDecoder();");
        // f.import("TextDecoder")
        // f.preamble(file);
        const { Message, JsonValue, IMessageTypeRegistry } = schema.runtime;
        // Convert the Message ImportSymbol to a type-only ImportSymbol
        const MessageAsType = Message.toTypeOnly();
        for (const service of file.services) {
            const localServiceName = localName(service);
            f.print(makeJsDoc(service));
            f.print`export class ${localServiceName}Client {`;
            f.print`    private contract: ${contract};`;
            f.print`    private jsonWriteOptions:Partial<${jsonWriteOptions}> = {};`;
            f.print`    private registry: ${registry};`;
            // f.print`    private registry: ${registry}  = '';`;
            f.print();
            f.print`    constructor(contract: ${contract}, registry: ${registry}) {`;
            f.print`        this.contract = contract;`;
            f.print`        this.registry = registry;`;
            f.print`    }`;
            f.print();
            for (const method of service.methods) {
                if (method.methodKind === MethodKind.Unary) {
                    f.print();
                    f.print(makeJsDoc(method, "    "));

                    if (method.input.typeName == "google.protobuf.Empty") {
                        f.print`    async ${localName(method)}(evaluate: boolean): Promise< ${method.output}> {`;
                        f.print`        if (evaluate) {`;
                        f.print`            const results = utf8Decoder.decode(`;
                        f.print`                await this.contract.evaluateTransaction(`;
                        f.print`                ${literalString(method.name)}`;
                        f.print`            )`;
                        f.print`            )`;
                        f.print`            return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`;
                        f.print`        } else {`;
                        f.print`            const results = utf8Decoder.decode(`;
                        f.print`                    await this.contract.submitTransaction(`;
                        f.print`                ${literalString(method.name)}`;
                        f.print`            )`;
                        f.print`            )`;
                        f.print`            return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`;
                        f.print`        }`;

                        f.print`    }`;
                    } else {
                        f.print`    async ${localName(method)}(request: ${
                            method.input
                        }, evaluate: boolean ): Promise< ${method.output}> {`;
                        f.print`        if (evaluate) {`;
                        f.print`            const results = utf8Decoder.decode(`;
                        f.print`                await this.contract.evaluateTransaction(`;
                        f.print`                ${literalString(method.name)},`;
                        f.print`                request.toJsonString(this.jsonWriteOptions)`;
                        f.print`            )`;
                        f.print`            )`;
                        f.print`            return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`;
                        f.print`        } else {`;
                        f.print`            const results = utf8Decoder.decode(`;
                        f.print`                    await this.contract.submitTransaction(`;
                        f.print`                ${literalString(method.name)},`;
                        f.print`                request.toJsonString(this.jsonWriteOptions)`;
                        f.print`            )`;
                        f.print`            )`;
                        f.print`            return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`;
                        f.print`        }`;

                        f.print`    }`;
                    }


                }
            }
            f.print`}`;
        }
    }
}

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
                    const populated = populateFeilds(input);
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
