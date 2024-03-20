import fs from "fs";
import { fileURLToPath } from "url";
import { ofetch } from "ofetch";

import YAML from "yaml";
import { defu } from "defu";
import path from "path";
import { exit } from "process";

export const BASEURL = "https://api-biochain.ittc.ku.edu";

export function ReadFile(name: string): Object {
    // [ true, false, 'maybe', null ]
    const __filename = fileURLToPath(import.meta.url);
    // const currentDirectory = __dirname;
    // get the current directory
    const currentDirectory = path.dirname(__filename);
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
export async function Login(
    username: string,
    password: string
): Promise<string> {
    const res = await ofetch(`${BASEURL + "/auth/login"}`, {
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

    const res = await ofetch(`${BASEURL + "/auth/fullWho"}`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });

    return res;
}

export const USERS = Object.entries(
    ReadFile("merged.yml").organizations.Org1MSP.users
);

async function SetIdentity(
    username: string,
    password: string,
    signKey: string,
    key: string,
    mspId: string,
    name?: string
): Promise<Object> {
    const token = await Login(username, password);

    const res = await ofetch(`${BASEURL + "/auth/setIdentity"}`, {
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

export async function InvokeChaincode(
    username: string,
    password: string,

    chaincode: string,
    method: string,
    args: string
): Promise<Object> {
    const token = await Login(username, password);

    const res = await ofetch(
        `${BASEURL}/contract/demo/${chaincode}/invoke/${method}`,
        {
            method: "POST",
            headers: {
                Authorization: `Bearer ${token}`,
            },
            body: args,
        }
    );

    return res;
}

export async function QueryChaincode(
    username: string,
    password: string,

    chaincode: string,
    method: string,
    args: string
): Promise<Object> {
    const token = await Login(username, password);

    const res = await ofetch(`${BASEURL}/biochain/query/${method}`, {
        method: "post",
        headers: {
            Authorization: `Bearer ${token}`,
        },
        body: args,
    });

    return res;
}

export async function InvokeAs(
    user: number,
    chaincode: string,
    method: string,
    args: string
): Promise<Object> {
    const username = USERS[user][0];
    const password = USERS[user][1].password;

    return await InvokeChaincode(username, password, chaincode, method, args);
}

export async function QueryAs(
    user: number,
    chaincode: string,
    method: string,
    args: string
): Promise<Object> {
    const username = USERS[user][0];
    const password = USERS[user][1].password;

    return await QueryChaincode(username, password, chaincode, method, args);
}
