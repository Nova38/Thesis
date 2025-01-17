// import { fileURLToPath } from 'node:url'
// import antfu from '@antfu/eslint-config'

// import perfectionistNatural from 'eslint-plugin-perfectionist/configs/recommended-natural'
// import eslintConfigPrettier from 'eslint-config-prettier'

// import { dirname, resolve } from 'pathe'
// import NuxtEslintConfig from './.nuxt/eslint.config.mjs'

// const __filename = fileURLToPath(import.meta.url) // get the resolved path to the file
// const __dirname = dirname(__filename) // get the name of the directory

// Get cu
import withNuxt from './.nuxt/eslint.config.mjs'
import eslintPluginPrettierRecommended from 'eslint-plugin-prettier/recommended'

export default withNuxt(eslintPluginPrettierRecommended, {
  files: ['presets/wind/**/*.js'],
  rules: {
    eqeqeq: 'off',
    'no-dupe-keys': 'off',
    'no-unused-vars': 'off',
  },
})

// antfu(
//   {
//     formatters: {
//       css: true,
//       html: true,
//       markdown: true,
//     },
//     // ...@antfu/eslint-config options,
//     stylistic: true,

//     vue: {
//       overrides: {
//         // 'vue/block-order': ['error', '{"order": ["script", "template", "style"]}'],
//       },
//     },
//     ignores: ['formkit.theme.ts'],
//     // vue: {
//     //   sfcBlocks: {
//     //     blocks: {
//     //       script: true,
//     //       template: true,
//     //       style: true,
//     //     },
//     //   },P
//     // },
//   },

//   // perfectionistNatural,
//   // Add the Nuxt rules
//   NuxtEslintConfig,
//   // eslintConfigPrettier,
//   // eslintPluginPrettierRecommended,

//   // ...your other rules
//   {
//     rules: {
//       'no-console': 'off',
//       'perfectionist/sort-vue-attributes': 'off',
//       'perfectionist/sort-imports': 'off',
//       'ts/no-redeclare': 'off',
//       'ts/no-unsafe-argument': 'off',
//       'ts/no-unsafe-assignment': 'off',
//       'ts/no-throw-literal': 'off',
//       'node/prefer-global/process': 'off',
//       '@typescript-eslint/no-explicit-any': 'off',
//       'eslint-comments/no-unlimited-disable': 'off',
//     },
//   },
//   {
//     files: ['presets/wind/**/*.js'],
//     rules: {
//       'eqeqeq': 'off',
//       'no-dupe-keys': 'off',
//       'no-unused-vars': 'off',
//     },
//   },
// )
