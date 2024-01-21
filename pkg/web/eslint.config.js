import antfu from '@antfu/eslint-config'
import perfectionistNatural from 'eslint-plugin-perfectionist/configs/recommended-natural'

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
      overrides: [

      ],
      tsconfigPath: 'tsconfig.json',
    },
    vue: true,

  },

  perfectionistNatural,
  // Add the Nuxt rules
  NuxtEslintConfig,
  // ...your other rules
  {
    rules: {
      'no-console': 'off',
      'perfectionist/sort-vue-attributes': 'off',
      'ts/no-redeclare': 'off',
    },
  },
)
