// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve } from 'pathe'

export default defineNuxtConfig({
  appConfig: {
    apiEndpoint: process.env.NUXT_API_URL,
  },
  colorMode: {
    preference: 'light',
  },

  tailwindcss: {
    exposeConfig: true,
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

  extends: [['@saacs/ui', { install: true }]],

  imports: {
    dirs: ['composables/cc/**', 'utils/**', 'utils/formkit/*'],
    presets: [
      {
        from: 'protobuf-es',
        imports: ['PlainMessage'],
      },
      {
        from: '@saacs/saacs-pb',
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
    options: {},
  },

  modules: [
    '@vueuse/nuxt',
    '@nuxt/test-utils/module',
    '@pinia/nuxt',
    // '@hebilicious/vue-query-nuxt',
    '@formkit/nuxt',
    '@nuxt/eslint',
    'nuxt-build-cache',
    'nuxt-jsoneditor',
    '@nuxt/ui',
    '@formkit/auto-animate',
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

  // quasar: {
  //   components: {
  //     deepDefaults: true,
  //     defaults: {
  //       QInput: {
  //         dense: true,
  //         outlined: true,
  //         stackLabel: true,
  //       },
  //     },
  //   },
  //   config: {
  //     loadingBar: {
  //       color: 'secondary',
  //       position: 'bottom',
  //       size: '4px',
  //     },
  //   },
  //   extras: {
  //     font: 'roboto-font',
  //     fontIcons: ['themify', 'material-icons'],
  //   },
  //   plugins: ['LoadingBar', 'Notify'],
  // },

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

  testUtils: {
    logToConsole: true,
  },

  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: true,
    },
  },
})
