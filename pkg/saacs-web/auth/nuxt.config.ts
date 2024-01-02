export default defineNuxtConfig({
    runtimeConfig: {
        auth: {
            name: "nuxt-session",
            password: process.env.NUXT_AUTH_PASSWORD || "",
        },
    },
});
