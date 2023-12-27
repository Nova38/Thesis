import fs from "fs";
import { fileURLToPath } from "url";
import { ofetch } from "ofetch";

import YAML from "yaml";
import { defu } from "defu";
import path from "path";
import { exit } from "process";

const baseurl = "https://api-biochain.ittc.ku.edu";

function readFile(name: string): Object {
    // [ true, false, 'maybe', null ]
    const __filename = fileURLToPath(import.meta.url);
    // const currentDirectory = __dirname;
    const filePath = path.join(__filename, name);
    const file = fs.readFileSync(filePath, "utf8");
    return YAML.parse(file);
}

async function Login(username: string, password: string): Promise<string> {
    const res = await ofetch(`${baseurl + "/auth/login"}`, {
        method: "POST",
        body: JSON.stringify({
            username: username,
            password: password,
        }),
    });

    return res.access_token;
}

async function GetFullUser(
    username: string,
    password: string
): Promise<Object> {
    const token = await Login(username, password);

    const res = await ofetch(`${baseurl + "/auth/fullWho"}`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });

    return res;
}

async function Signup(
    name: string,
    username: string,
    password: string
): Promise<string> {
    const res = await ofetch(`${baseurl + "/auth/signup"}`, {
        method: "POST",
        body: JSON.stringify({
            username: username,
            password: password,
            name: name,
        }),
    });

    return res.access_token;
}

async function SetIdentity(
    username: string,
    password: string,
    signKey: string,
    key: string,
    mspId: string,
    name?: string
): Promise<Object> {
    const token = await Login(username, password);

    const res = await ofetch(`${baseurl + "/auth/setIdentity"}`, {
        method: "POST",
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
            signKey: signKey,
            key: key,
            mspId: mspId,
            name: name,
        }),
    });

    return res;
}

async function Register(
    name: string,
    username: string,
    password: string,
    signKey: string,
    key: string,
    mspId: string
) {
    try {
        await Signup(name, username, password);
    } finally {
        return await SetIdentity(username, password, signKey, key, mspId, name);
    }
}
// for (let user of users) {
//     console.log(user);

//     const data: any = user[1];

//     // const res = await Register(
//     //     data.name,
//     //     user[0],
//     //     data.password,
//     //     data.key.pem,
//     //     data.cert.pem,
//     //     "Org1MSP"
//     // );

//     const res = await SetIdentity(
//         user[0],
//         data.password,
//         data.key.pem,
//         data.cert.pem,
//         "Org1MSP",
//         data.name
//     );

//     console.log(res);
// }
// await Login("admin_1", "HBN8dfz3tvm_nmx");
let users = Object.entries(readFile("merged.yml").organizations.Org1MSP.users);

await GetFullUser(users[0][0], users[0][1].password);
