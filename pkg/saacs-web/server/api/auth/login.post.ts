import { z } from "zod";

const userSchema = z.object({
    username: z.string(),
    password: z.string(),
});

export default defineEventHandler(async (event) => {
    const session = await useAuthSession(event);

    if (session.data.username) {
        throw createError({
            message: "Already logged in!",
            statusCode: 400,
        });
    }

    const result = await readValidatedBody(event, (body) =>
        userSchema.safeParse(body),
    );

    if (!result.success) throw result.error.issues;

    const { username, password } = result.data;

    const user = await findUserByUsername(username);

    if (!user) {
        throw createError({
            message: "username not found! Please register.",
            statusCode: 401,
        });
    }
    if (!user.password || user.password !== (await hash(password))) {
        throw createError({
            message: "Incorrect password!",
            statusCode: 401,
        });
    }
    await session.update({
        username: user.username,
    });
    return session;
});
