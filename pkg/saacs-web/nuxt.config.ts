// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    modules: [
        "@vueuse/nuxt",
        "@pinia/nuxt",
        "@nuxt/ui",
        "nuxt-quasar-ui",
        "@nuxtjs/tailwindcss",
        "@nuxt/test-utils/module",
        "@nuxtjs/eslint-module",
    ],
    nitro: {
        storage: {
            ".data:auth": { driver: "fs", base: "./.data/auth" },
        },
    },
    devtools: {
        enabled: true,

        timeline: {
            enabled: true,
        },
    },
    ssr: false,
    ui: {
        icons: ["heroicons", "material-symbols", "simple-icons"],
    },

    quasar: {
        components: {
            defaults: {
                QInput: {},
            },
        },
    },

    runtimeConfig: {
        auth: {
            password: "password",
        },
        fabric: {
            chaincode: {
                channel: "mychannel",
                chaincode: "saacs",
            },
            peer: {
                url: "grpcs://localhost:7051",
                tlsCACerts: {
                    pem: "",
                },
                grpcOptions: {
                    "ssl-target-name-override": "peer0.org1.example.com",
                },
            },
            public: {
                mspId: "Org1MSP",
                credentials: "",
                key: "",
            },
        },
    },
});
