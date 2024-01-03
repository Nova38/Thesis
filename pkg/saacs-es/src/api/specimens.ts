// import orn from "../sample/biochain/orn.json" assert { type: "json" };
import { parse } from "csv";
import * as fs from "fs";
import path from "path";

import { fileURLToPath } from "url";
const __filename = fileURLToPath(import.meta.url); // get the resolved path to the file
const __dirname = path.dirname(__filename); // get the name of the directory
import { construct, omit } from "radash";

import * as orn from "./orn.json" assert { type: "json" };
import { Specimen, Specimen_Primary, Date } from "../gen/biochain/v1/index.js";

async function main() {
    const specimens = orn.default.map((element) => {
        // const omitFields = ;

        let i = construct(omit(element, ["index"]));
        // console.log(i);

        const s = new Specimen(i);

        if (!s.primary) s.primary = new Specimen_Primary();
        if (!s.primary.originalDate) s.primary.originalDate = new Date();

        s.primary.originalDate.verbatim = element["primary.originalDate"];

        console.log(s);
        return s;
    });
    console.log(specimens);
}

main().catch(console.error);
