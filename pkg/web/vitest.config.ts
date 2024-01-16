// vitest.config.ts
// import { fileURLToPath } from "node:url";
import { defineVitestConfig } from "@nuxt/test-utils/config";

export default defineVitestConfig({
  test: {
    includeSource: ["src/**/*.{js,ts}"],

    // environment: "nuxt",
    // you can optionally set Nuxt-specific environment options
    // environmentOptions: {
    //   nuxt: {
    //     rootDir: fileURLToPath(new URL('./playground', import.meta.url)),
    //     overrides: {
    //       // other Nuxt config you want to pass
    //     }
    //   }
    // }
  },
});
