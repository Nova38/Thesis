// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve } from 'pathe'

export default defineNuxtConfig({
  appConfig: {
    apiEndpoint: process.env.NUXT_API_URL,
  },
  colorMode: {
    preference: 'light',
  },

  css: [resolve(__dirname, './assets/css/main.css')],

  // primevue: {
  //   components: {
  //     prefix: 'P',
  //   },
  //   composables: {
  //     exclude: ['useToast'],
  //   },
  //   importPT: { from: resolve(__dirname, 'lib/primevue/presets/wind/') }, // import and apply preset
  //   options: {
  //     unstyled: true,
  //   },
  // },
  // formkit: {
  //   configFile: resolve(__dirname, './formkit.config.ts'),
  // },

  ssr: false,

  debug: true,

  devServer: {
    // https: true,
    port: 8080,
  },
  devtools: {
    disableAuthorization: true,

    enabled: true,
    timeline: {
      enabled: true,
    },
    vscode: {
      enabled: true,
    },
  },

  experimental: {
    // typedPages: false,
    asyncContext: true,
  },

  extends: [['@saacs/ui', { install: true }]],

  imports: {
    dirs: ['composables/cc/**', 'utils/**'],
    presets: [
      {
        from: 'protobuf-es',
        imports: ['PlainMessage'],
      },
      {
        from: '~/lib/pb',
        imports: ['auth', 'ccbio', 'common', 'GlobalRegistry'],
      },
      {
        from: 'defu',
        imports: ['defu'],
      },
    ],
  },
  formkit: {
    // configFile: 'node_modules/@saacs/ui/formkit.config.ts',
  },

  modules: [
    // '@formkit/nuxt',
    // '@nuxt/ui',
    '@vueuse/nuxt',
    '@nuxt/test-utils/module',
    '@pinia/nuxt',
    'nuxt-quasar-ui',
    '@formkit/nuxt',
    '@nuxt/eslint',
    '@hebilicious/vue-query-nuxt',
  ],

  nitro: {
    storage: {
      '.data:auth': {
        base: './.data/auth',
        driver: 'fs',
      },
    },
    esbuild: {
      options: {
        target: 'esnext',
      },
    },
    experimental: {
      tasks: true,
    },
    imports: {
      dirs: ['lib/pb/*'],
      mergeExisting: true,
      // presets: [
      //   {
      //     from: './lib/pb/index.ts',
      //     imports: ['auth', 'ccbio', 'common', 'GlobalRegistry'],
      //   },
      // ],
    },
  },

  quasar: {
    components: {
      deepDefaults: true,
      defaults: {
        QInput: {
          dense: true,
          outlined: true,
          stackLabel: true,
        },
      },
    },
    config: {
      loadingBar: {
        color: 'secondary',
        position: 'bottom',
        size: '4px',
      },
    },
    extras: {
      font: 'roboto-font',
      fontIcons: ['themify', 'material-icons'],
    },
    plugins: ['LoadingBar', 'Notify'],
  },

  runtimeConfig: {
    auth: {
      password: process.env.NUXT_AUTH_PASSWORD || '',
    },
    fabric: {
      chaincode: {
        chaincode: process.env.NUXT_CHAINCODE_CHAINCODE || '',
        channel: process.env.NUXT_CHAINCODE_CHANNEL || '',
      },
      peer: {
        grpcOptions: {
          'ssl-target-name-override':
            process.env.NUXT_FABRIC_PEER_GRPC_OPTIONS || '',
        },
        tlsCACerts: {
          pem: process.env.NUXT_FABRIC_PEER_TLS_CA_CERTS_PEM || '',
        },
        url: process.env.NUXT_FABRIC_PEER_URL || '',
      },
      public: {
        credentials: process.env.NUXT_FABRIC_PUBLIC_CREDENTIALS || '',
        key: process.env.NUXT_FABRIC_PUBLIC_KEY || '',
        mspId: process.env.NUXT_FABRIC_PUBLIC_MSP || '',
      },
    },
    public: {
      api: {
        url: process.env.NUXT_API_URL || '',
      },
    },
  },
  sourcemap: true,

  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: true,
    },
  },
})
