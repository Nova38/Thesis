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
    // autoImport: true,
  },

  jsoneditor: {
    componentName: 'JsonEditor',
    options: {
      /**
       *
       * SET GLOBAL OPTIONS
       *
       * */
    },
  },

  modules: [
    // '@formkit/nuxt',
    // '@nuxt/ui',
    '@vueuse/nuxt',
    '@nuxt/test-utils/module',
    '@pinia/nuxt',
    'nuxt-quasar-ui',
    '@formkit/nuxt', // '@hebilicious/vue-query-nuxt',
    '@nuxt/eslint',
    'nuxt-build-cache',
    'nuxt-jsoneditor',
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
      // asyncContext: true,
      database: true,
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
  sourcemap: true,

  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: true,
    },
  },
})
