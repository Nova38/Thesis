export default defineEventHandler(async () => {
  return useStorage(".data:auth").getKeys();
});
