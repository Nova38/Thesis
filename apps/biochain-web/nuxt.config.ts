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

  // formkit: {
  //   configFile: resolve(__dirname, './formkit.config.ts'),
  // },

  ssr: false,

  // debug: true,

  devServer: {
    // https: true,
    port: 8000,
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
    watcher: 'parcel',
  },

  imports: {
    dirs: ['composables/cc/**', 'utils/**', 'utils/formkit/*'],
    presets: [
      {
        from: 'protobuf-es',
        imports: ['PlainMessage'],
      },
      {
        from: '@saacs/saacs-pb',
        imports: ['auth', 'ccbio', 'common', 'GlobalRegistry', 'pb'],
      },
      {
        from: 'defu',
        imports: ['defu'],
      },
    ],
  },

  modules: [
    '@vueuse/nuxt',
    '@nuxt/test-utils/module',
    '@pinia/nuxt',
    '@hebilicious/vue-query-nuxt',
    '@formkit/nuxt',
    '@nuxt/eslint',
    'nuxt-build-cache',
    // 'nuxt-jsoneditor',
    '@nuxt/ui',
    '@formkit/auto-animate',
    'nuxt-primevue',
    'nuxt-shiki',
  ],

  nitro: {
    storage: {
      '.data:auth': {
        base: './.data/db/usersDB',
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
      typescriptBundlerResolution: true,
      openAPI: true,
      database: true,
    },
    imports: {
      mergeExisting: true,
      presets: [
        {
          from: '@saacs/saacs-pb',
          imports: ['auth', 'ccbio', 'common', 'GlobalRegistry', 'pb'],
        },
      ],
    },
  },

  primevue: {
    components: { prefix: 'P' },
    composables: { exclude: ['useToast'] },
    importPT: { from: resolve(__dirname, './primevue/presets/wind/') }, // import and apply preset
    options: { unstyled: true },
  },

  runtimeConfig: {
    auth: {
      password: '',
    },
    fabric: {
      chaincode: {
        chaincode: '',
        channel: '',
      },
      peer: {
        grpcOptions: {
          'ssl-target-name-override': '',
        },
        tlsCACerts: {
          pem: '',
        },
        url: '',
      },
      public: {
        credentials: '',
        key: '',
        msp: '',
      },
    },
  },

  typescript: {
    tsConfig: {
      exclude: ['eslint.config.mjs'],
    },
  },

  sourcemap: true,

  testUtils: {
    logToConsole: true,
  },

  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: true,
    },
  },
})
