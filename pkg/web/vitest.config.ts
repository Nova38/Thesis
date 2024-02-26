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

import { defineVitestConfig } from '@nuxt/test-utils/config'

// export default defineConfig({
//   test: {
//     // ...
//   },
// })

export default defineVitestConfig({
  test: {
    // coverage: {
    //   reportsDirectory: 'coverage',
    // },
    // includeSource: ['src/**/*.{js,ts}'],
    // environmentOptions: {},
    // api: 15555,
    // // setupFiles: './tests/setup/mocks.ts',
    globals: true,
  },
})
