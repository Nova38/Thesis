import { defineBuildConfig } from 'unbuild'

import { BundleSchemas } from './src/build/schema'

export default defineBuildConfig({
  // If entries is not provided, will be automatically inferred from package.json
  entries: [
    // default
    './src/index',

    // mkdist builder transpiles file-to-file keeping original sources structure
    {
      builder: 'mkdist',
      input: './src/gen/',
      outDir: './dist/gen',
      format: 'cjs',
    },
    {
      builder: 'mkdist',
      input: './src/gen/',
      outDir: './dist/gen',
      format: 'esm',
    },
    {
      builder: 'mkdist',
      input: './src/schema/',
      outDir: './dist/schema/',
    },
  ],

  clean: true,
  sourcemap: true,

  // Change outDir, default is 'dist'

  // Generates .d.ts declaration file
  declaration: true,
  hooks: {
    'build:before': async () => {
      console.info('Bundling schemas ...')
      BundleSchemas()
      console.info('Bundling schemas done')
    },
  },
  rollup: {
    emitCJS: true,
  },
})
