import { Schema } from "@bufbuild/protoplugin/ecmascript";
import { getAllTypes } from "../utils";

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
export function generateRegistry(schema: Schema) {
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
