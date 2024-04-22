import antfu from '@antfu/eslint-config'

export default antfu({
  formatters: true,
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
})
