// https://nuxt.com/docs/api/configuration/nuxt-config
import { resolve, join } from 'pathe'

export default defineNuxtConfig({
  devtools: { enabled: true },
  css: [resolve(__dirname, './assets/css/main.css')],

  modules: [
    'nuxt-primevue',
    '@nuxt/ui',
    '@vueuse/nuxt',
    'nuxt-shiki',
    'nuxt-build-cache',
  ],

  tailwindcss: {
    config: {
      content: [
        resolve(__dirname, './primevue/presets/wind/'),
        resolve(__dirname, './components/**/*.{js,vue,ts}'),
        resolve(__dirname, './layouts/**/*.vue'),
        resolve(__dirname, './**/*.vue'),
        resolve(__dirname, './plugins/**/*.{js,ts}'),
        resolve(__dirname, './formkit.theme.ts'),
      ],
    },
    editorSupport: true,
  },

  // formkit: {
  //   configFile: resolve(__dirname, "./formkit.config.ts"),
  // },

  primevue: {
    components: {
      prefix: 'P',
    },
    composables: {
      exclude: ['useToast'],
    },
    importPT: { from: resolve(__dirname, './primevue/presets/wind/') }, // import and apply preset
    options: {
      unstyled: true,
    },
  },
  ui: {
    icons: {},
  },
})
