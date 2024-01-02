export default eventHandler(async (event) => {
    const session = await useAuthSession(event);
    const { username, password } = await readBody(event);
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
        id: user.id,
        username: user.username,
    });
    return session;
});
