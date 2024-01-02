import { randomUUID } from "uncrypto";
import * as crypto from "crypto";

export interface User {
    id: string;
    createdAt: string;
    username: string;

    password: string;

    mspId: string;
    credentials: string;
    key: string;
}

export function userToIdentity(user: User) {
    if (!user.mspId || !user.credentials) {
        throw createError({
            message: "Users mspId or signCert not found!",
            statusCode: 404,
        });
    }

    return {
        mspId: user.mspId,
        credentials: Buffer.from(user.credentials),
    };
}

export function userToPrivateKey(user: User) {
    if (!user.key) {
        throw createError({
            message: "Users private key not found!",
            statusCode: 404,
        });
    }

    const privateKey = crypto.createPrivateKey(user.key);

    return {
        privateKey,
    };
}

export async function findUserByUsername(username: string): Promise<User> {
    const storage = useStorage(".data:auth");
    const key = getUserKey(username!);

    const user = await storage.getItem(key);
    if (!user) {
        throw createError({ message: "User not found!", statusCode: 404 });
    }
    return user as User;
}

export async function createUser(user: Partial<User>) {
    const key = getUserKey(user.username!);
    if (await useStorage(".data:auth").hasItem(key)) {
        throw createError({
            message: "Username already exists!",
            statusCode: 409,
        });
    }
    console.log("createUser", user);
    return await useStorage(".data:auth").setItem(key, user);
}

export async function updateUserByUsername(
    username: string,
    updates: Partial<User>,
) {
    const storage = useStorage(".data:auth");
    const user = await findUserByUsername(username);
    const key = getUserKey(user.username!);
    return await storage.setItem(key, {
        ...user,
        ...updates,
    });
}

function getUserKey(username: string) {
    return `db:usersDB:${encodeURIComponent(username)}`;
}
