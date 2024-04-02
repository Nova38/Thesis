// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve } from 'pathe'

export default defineNuxtConfig({
  appConfig: {},
  colorMode: {
    preference: 'light',
  },

  css: [resolve(__dirname, './assets/css/main.css')],

  primevue: {
    components: {
      prefix: 'P',
    },
    composables: {
      exclude: ['useToast'],
    },
    importPT: { from: resolve(__dirname, 'lib/primevue/presets/wind/') }, // import and apply preset
    options: {
      unstyled: true,
    },
  },
  formkit: {
    configFile: resolve(__dirname, './formkit.config.ts'),
  },

  ssr: false,

  debug: false,

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

  modules: [
    '@formkit/nuxt',
    '@nuxt/ui',
    '@vueuse/nuxt',
    '@nuxt/test-utils/module',
    '@pinia/nuxt',
    'nuxt-quasar-ui',
    '@formkit/nuxt',
    'nuxt-primevue',
    '@nuxt/eslint',

    // 'nuxt-radash',
    // 'nuxt-security',
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

  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: true,
    },
  },
})
