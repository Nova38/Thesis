// vitest.config.ts
// import { fileURLToPath } from "node:url";

// export default defineVitestConfig({
//   test: {
//     environment: 'nuxt',
//     api: 9999,

//
//     // you can optionally set Nuxt-specific environment options
//     // environmentOptions: {
//     //   nuxt: {
//     //     rootDir: fileURLToPath(new URL('./playground', import.meta.url)),
//     //     overrides: {
//     //       // other Nuxt config you want to pass
//     //     }
//     //   }
//     // }
//   },
// })

import { fileURLToPath } from 'node:url'
import { defineVitestConfig } from '@nuxt/test-utils/config'

export default defineVitestConfig({
  test: {
    coverage: {
      reportsDirectory: 'coverage',
    },
    includeSource: ['src/**/*.{js,ts}'],
    environmentOptions: {
      nuxt: {
        rootDir: fileURLToPath(new URL('./', import.meta.url)),
        domEnvironment:
          (process.env.VITEST_DOM_ENV as 'happy-dom' | 'jsdom') ?? 'happy-dom',
        mock: {
          indexedDb: true,
        },
      },
    },
    // setupFiles: './tests/setup/mocks.ts',
    globals: true,
  },
})
