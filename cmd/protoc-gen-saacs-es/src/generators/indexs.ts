import { Schema } from "@bufbuild/protoplugin/ecmascript";

export function generateIndex(schema: Schema) {
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
            // f.print(`export * from "./${base}_pb_reg.js"`);
        }

        if (file.services.length > 0) {
            f.print(`export * from "./${base}_pb_gateway.js"`);
        }
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
    });

}
