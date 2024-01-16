import fs from "fs";
import { fileURLToPath } from "url";
import { ofetch } from "ofetch";

import YAML from "yaml";
import { defu } from "defu";
import path from "path";
import { exit } from "process";

const baseurl = "https://api-biochain.ittc.ku.edu/api";

function readFile(name: string): Object {
    // [ true, false, 'maybe', null ]
    const __filename = fileURLToPath(import.meta.url);
    // const currentDirectory = __dirname;
    // get the current directory
    const currentDirectory = path.dirname(__filename);
    const filePath = path.join(currentDirectory, name);
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
    username: string,
    password: string,
    credentials: string,
    key: string,
    mspId: string
): Promise<string> {
    const res = await ofetch(`${baseurl + "/auth/register"}`, {
        method: "POST",
        body: JSON.stringify({
            username: username,
            password: password,
            credentials: credentials,
            key: key,
            mspId: mspId,
        }),
    });

    return res;
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

async function main() {
    let users = Object.entries(
        readFile("merged.yml").organizations.Org1MSP.users
    );

    users.forEach((user) => {
        console.log(
            user["username"],
            user["password"],
            user["credentials"],
            user["key"],
            user["mspId"]
        );
        console.log(res);
    });
}

// await GetFullUser(users[0][0], users[0][1].password);
