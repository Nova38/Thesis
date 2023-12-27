import fs from "fs";
import { fileURLToPath } from "url";
import { ofetch } from "ofetch";

import YAML from "yaml";
import { defu } from "defu";
import path from "path";

function readFile(name: string): Object {
    // [ true, false, 'maybe', null ]
    const __filename = fileURLToPath(import.meta.url);
    const __dirname = path.dirname(__filename);
    const currentDirectory = __dirname;
    const filePath = path.join(currentDirectory, name);
    const file = fs.readFileSync(filePath, "utf8");
    return YAML.parse(file);
}

function writeFile(name: string, data: Object): void {
    // [ true, false, 'maybe', null ]
    const __filename = fileURLToPath(import.meta.url);
    const __dirname = path.dirname(__filename);
    const currentDirectory = __dirname;
    const filePath = path.join(currentDirectory, name);

    const strData = YAML.stringify(data);

    fs.writeFileSync(filePath, strData);
}

let certs = readFile("certs.yml");
let passwd = readFile("passwd.yml");
let config = defu(certs, passwd);

writeFile("merged.yml", config);

let users = config.organizations.Org1MSP.users;
typeof users;

for (const [key, value] of Object.entries(users)) {
    console.log(`${key}, ${value.password}`);
}
