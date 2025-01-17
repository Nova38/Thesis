import { defineBuildConfig } from 'unbuild'

export default defineBuildConfig({
  // If entries is not provided, will be automatically inferred from package.json

  clean: true,
  sourcemap: true,

  // Change outDir, default is 'dist'

  // Generates .d.ts declaration file
  declaration: true,

  rollup: {
    // emitCJS: true,
  },
})
