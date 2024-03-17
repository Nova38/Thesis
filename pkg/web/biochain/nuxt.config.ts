// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve } from 'pathe'

/* eslint perfectionist/sort-objects: "error" */
export default defineNuxtConfig({
  appConfig: {
    apiEndpoint:
      process.env.NUXT_API_URL || 'https://api-biochain.ittc.ku.edu/',
  },
  colorMode: {
    preference: 'light',
  },

  extends: [['../ui', { install: true }]],
  tailwindcss: {
    exposeConfig: true,
    config: {
      content: [
        './app.vue',
        './formkit.theme.ts',
        './components/**/*.{js,vue,ts}',
        './layouts/**/*.vue',
        './**/*.vue',
        './plugins/**/*.{js,ts}',
        './primevue/presets/**/*.{js,vue,ts}',
        './error.vue',
      ],

    }
  },

  formkit: {
    configFile: resolve(__dirname, './formkit.config.ts'),
  },

  ssr: false,

  debug: true,

  devServer: {
    https: true,
    // https: {
    //   cert: resolve(__dirname, '.dev/RootCA.pem'),
    //   key: resolve(__dirname, '.dev/RootCA.key'),
    // }, // enable HTTPS

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

  eslintConfig: {
    setup: false,
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
        from: 'saacs-es',
        imports: ['auth', 'ccbio', 'common'],
      },
      {
        from: 'defu',
        imports: ['defu'],
      },
    ],
  },

  modules: [
    '@nuxt/test-utils/module',
    'nuxt-module-eslint-config',
    '@pinia/nuxt',
    'nuxt-quasar-ui',
    // '@formkit/nuxt',
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
        mspId: process.env.NUXT_FABRIC_PUBLIC_MSPID || '',
      },
    },
    public: {
      api: {
        url: process.env.NUXT_API_URL || '',
      },
    },
  },



  vite: {
    define: {
      __VUE_PROD_DEVTOOLS__: true,
    },
  },
})
/* eslint perfectionist/sort-objects: "off" */
