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
        QInput: {},
      },
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
  runtimeConfig: {
    auth: {
      password: "password",
    },
    fabric: {
      chaincode: {
        chaincode: "roles",
        channel: "mychannel",
      },
      peer: {
        grpcOptions: {
          "ssl-target-name-override": "peer0.org1.example.com",
        },
        tlsCACerts: {
          pem: "",
        },
        url: "grpcs://localhost:7051",
      },
      public: {
        credentials: "",
        key: "",
        mspId: "Org1MSP",
      },
    },
  },
  ssr: false,
});
