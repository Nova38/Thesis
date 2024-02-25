import antfu from '@antfu/eslint-config'

// import perfectionistNatural from 'eslint-plugin-perfectionist/configs/recommended-natural'
// import eslintConfigPrettier from 'eslint-config-prettier'

import NuxtEslintConfig from './.nuxt/eslint.config.mjs'

export default antfu(
  {
    formatters: {
      css: true,
      html: true,
      markdown: true,
    },
    // ...@antfu/eslint-config options,
    stylistic: true,
    typescript: {
      overrides: [],
      tsconfigPath: 'tsconfig.json',

    },
    vue: true,
    ignores: ['formkit.theme.ts'],
    // vue: {
    //   sfcBlocks: {
    //     blocks: {
    //       script: true,
    //       template: true,
    //       style: true,
    //     },
    //   },
    // },
  },

  // perfectionistNatural,
  // Add the Nuxt rules
  NuxtEslintConfig,
  // eslintConfigPrettier,
  // eslintPluginPrettierRecommended,

  // ...your other rules
  {
    rules: {
      'no-console': 'off',
      'perfectionist/sort-vue-attributes': 'off',
      'perfectionist/sort-imports': 'off',
      'ts/no-redeclare': 'off',
      'ts/no-unsafe-argument': 'off',
      'ts/no-unsafe-assignment': 'off',
      'ts/no-throw-literal': 'off',
      'node/prefer-global/process': 'off',
      '@typescript-eslint/no-explicit-any': 'off',
      'eslint-comments/no-unlimited-disable': 'off',
    },
  },
)
