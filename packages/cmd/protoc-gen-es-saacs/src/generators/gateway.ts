import { createEcmaScriptPlugin, runNodeJs } from "@bufbuild/protoplugin";
import {
    GeneratedFile,
    ImportSymbol,
    Schema,
    findCustomEnumOption,
    findCustomMessageOption,
    localName,
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
    Empty,
    MethodKind,
    createDescriptorSet,
    createRegistryFromDescriptors,
} from "@bufbuild/protobuf";

import { pb } from "saacs-es";

const TransactionType = pb.auth.auth.TransactionType;

import * as fs from "fs";

// Modified from the original protoc-gen-ts plugin
export function generateGateway(schema: Schema) {
    for (const file of schema.files) {
        generateFile(schema, file);
    }
}

function generateFile(schema: Schema, file: DescFile) {
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
        f.print(f.jsDoc(service));
        f.print`export class ${localServiceName}Client {`;
        f.print`    private contract: ${contract};`;
        f.print`    private jsonWriteOptions:Partial<${jsonWriteOptions}> = {};`;
        f.print`    private registry: ${registry};`;
        f.print();
        f.print`    constructor(contract: ${contract}, registry: ${registry}) {`;
        f.print`        this.contract = contract;`;
        f.print`        this.registry = registry;`;
        f.print`    }`;
        f.print();
        for (const method of service.methods) {
            if (method.methodKind === MethodKind.Unary) {
                f.print();
                f.print(f.jsDoc(method, "    "));
                const tt: TransactionType | undefined = findCustomEnumOption(
                    method,
                    50556
                );
                f.print(`// ${tt}`);

                const callMethod =
                    tt == TransactionType.INVOKE
                        ? "submitTransaction"
                        : "evaluateTransaction";

                if (method.input.typeName == "google.protobuf.Empty") {
                    f.print`    async ${localName(method)}(): Promise< ${
                        method.output
                    }> {`;
                    f.print`        const results = utf8Decoder.decode(`;
                    f.print`                await this.contract.${callMethod}(`;
                    f.print`                ${f.string(method.name)}`;
                    f.print`            )`;
                    f.print`        )`;
                    f.print`            return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`;
                    f.print`    }`;
                } else {
                    f.print`    async ${localName(method)}(request: ${
                        method.input
                    }, evaluate: boolean ): Promise< ${method.output}> {`;
                    f.print`            const results = utf8Decoder.decode(`;
                    f.print`                await this.contract.${callMethod}(`;
                    f.print`                ${f.string(method.name)},`;
                    f.print`                request.toJsonString(this.jsonWriteOptions)`;
                    f.print`            )`;
                    f.print`            )`;
                    f.print`            return ${method.output}.fromJsonString(results, {typeRegistry: this.registry});`;

                    f.print`    }`;
                }
            }
        }
        f.print`}`;
    }
}
