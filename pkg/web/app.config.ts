export default defineAppConfig({
  theme: {
    primaryColor: "#ababab",
  },
  apiEnd: process.env.API_URL ?? "/api",
});
