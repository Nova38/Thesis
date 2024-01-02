export default eventHandler(async (event) => {
    const { username, chainkey } = await readBody(event);

    // eslint-disable-next-line no-console
    console.log(username, chainkey);

    await updateUserByUsername(username, { chainKey: chainkey });
    return {
        message: "Successfully registered!",
    };
});
