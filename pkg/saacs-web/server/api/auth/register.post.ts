export default eventHandler(async (event) => {
    const { username, password } = await readBody(event);

    if (!username || !password) {
        throw createError({
            message: "username or password is empty!",
            statusCode: 401,
        });
    }

    const result = await createUser({
        username: username,
        password: await hash(password),
        mspId: "Org1MSP",
    });

    return {
        message: "Successfully registered!",
    };
});
