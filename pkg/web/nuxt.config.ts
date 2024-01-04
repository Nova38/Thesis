// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    // "@nuxt/ui",
    "@formkit/nuxt",
    "nuxt-quasar-ui",
    "@vueuse/nuxt",
    "@nuxtjs/tailwindcss",
    "@nuxtjs/eslint-module",
    "nuxt-icon",
    "@pinia/nuxt",
  ],
  css: ["~/assets/css/main.css"],
  devServer: {
    port: 8000,
  },

  formkit: {
    // Experimental support for auto loading (see note):
    // autoImport: true,
  },

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
  // ui: {
  //     icons: ["heroicons", "material-symbols", "simple-icons"],
  // },

  quasar: {
    components: {
      defaults: {
        QInput: {},
      },
    },
    config: {},

    extras: {
      fontIcons: [
        "themify",
        // 'line-awesome',
        // 'roboto-font-latin-ext', // this or either 'roboto-font', NEVER both!

        "material-icons", // optional, you are not bound to it
      ],
      font: "roboto-font",
    },
    plugins: ["LoadingBar"],
  },

  runtimeConfig: {
    auth: {
      password: "password",
    },
    fabric: {
      chaincode: {
        channel: "mychannel",
        chaincode: "roles",
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
