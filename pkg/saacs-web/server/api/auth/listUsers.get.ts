export default defineEventHandler(async (event) => {
    return useStorage(".data:auth").getKeys();
});
