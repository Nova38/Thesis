// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: ["~/assets/css/main.css"],
  devServer: {
    port: 8000,
  },
  devtools: {
    enabled: true,
    timeline: {
      enabled: true,
    },
  },
  formkit: {},
  imports: {
    dirs: ["composables/cc/**"],
  },
  modules: [
    "@formkit/nuxt",
    "nuxt-quasar-ui",
    "@vueuse/nuxt",
    "@nuxtjs/eslint-module",
    "nuxt-icon",
    "@pinia/nuxt",
    "@nuxtjs/tailwindcss",
    "nuxt-radash",
    "@nuxt/test-utils/module",
  ],
  nitro: {
    storage: {
      ".data:auth": {
        base: "./.data/auth",
        driver: "fs",
      },
    },
  },
  quasar: {
    components: {
      defaults: {
        QInput: {
          outlined: true,
          dense: true,
          stackLabel: true,
        },
      },
      deepDefaults: true,
    },
    config: {
      loadingBar: {
        color: "secondary",
        position: "bottom",
        size: "4px",
      },
    },
    extras: {
      font: "roboto-font",
      fontIcons: ["themify", "material-icons"],
    },
    plugins: ["LoadingBar", "Notify"],
  },
  appConfig: {
    apiEndpoint: process.env.NUXT_API_URL,
  },
  runtimeConfig: {
    auth: {
      password: process.env.NUXT_AUTH_PASSWORD || "",
    },
    fabric: {
      chaincode: {
        chaincode: process.env.NUXT_CHAINCODE_CHAINCODE || "",
        channel: process.env.NUXT_CHAINCODE_CHANNEL || "",
      },
      peer: {
        grpcOptions: {
          "ssl-target-name-override":
            process.env.NUXT_FABRIC_PEER_GRPC_OPTIONS || "",
        },
        tlsCACerts: {
          pem: process.env.NUXT_FABRIC_PEER_TLS_CA_CERTS_PEM || "",
        },
        url: process.env.NUXT_FABRIC_PEER_URL || "",
      },
      public: {
        credentials: process.env.NUXT_FABRIC_PUBLIC_CREDENTIALS || "",
        key: process.env.NUXT_FABRIC_PUBLIC_KEY || "",
        mspId: process.env.NUXT_FABRIC_PUBLIC_MSPID || "",
      },
    },
    public: {
      api: {
        url: process.env.NUXT_API_URL,
      },
    },
  },
  ssr: false,
});
