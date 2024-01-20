// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  css: ["~/assets/css/main.css"],
  devServer: {
    port: 8000,
    https: true,
  },
  devtools: {
    enabled: true,
    timeline: {
      enabled: true,
    },
  },
  imports: {
    dirs: ["composables/cc/**"],
    presets: [
      {
        from: "protobuf-es",
        imports: ["PlainMessage"],
      },
      {
        from: "saacs-es",
        imports: ["auth", "ccbio"],
      },
      {
        from: "defu",
        imports: ["defu"],
      },
    ],
  },
  ui: {
    global: true,
    icons: {},
  },
  modules: [
    "nuxt-quasar-ui",
    "@nuxt/ui",
    "@vueuse/nuxt",
    "nuxt-module-eslint-config",

    // "@nuxtjs/eslint-module",
    "@pinia/nuxt",
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

  sourcemap: true,

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
    apiEndpoint:
      process.env.NUXT_API_URL || "https://api-biochain.ittc.ku.edu/",
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
  routeRules: {
    "/api/**": {
      // enable CORS
      cors: true, // if enabled, also needs cors-preflight-request.ts Nitro middleware to answer CORS preflight requests
      headers: {
        // CORS headers
        "Access-Control-Allow-Origin": "*", // 'http://example:6006', has to be set to the requesting domain that you want to send the credentials back to
        "Access-Control-Allow-Methods": "*", // 'GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS'
        "Access-Control-Allow-Credentials": "true",
        "Access-Control-Allow-Headers": "*", // 'Origin, Content-Type, Accept, Authorization, X-Requested-With'
        "Access-Control-Expose-Headers": "*",
        // 'Access-Control-Max-Age': '7200', // 7200 = caching 2 hours (Chromium default), https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age#directives
      },
    },
  },
});
