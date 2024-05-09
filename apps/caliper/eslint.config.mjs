import antfu from '@antfu/eslint-config'

export default antfu({
  formatters: false,
  rules: {
    'no-console': 'off',
    'perfectionist/sort-vue-attributes': 'off',
    'ts/no-unused-expressions': 'warn',
    'perfectionist/sort-imports': 'off',
    'ts/no-redeclare': 'warn',
    'ts/no-unused-vars': 'warn',
    'ts/no-unsafe-argument': 'off',
    'ts/no-unused-vars': 'warn',
    'unused-imports/no-unused-vars': 'warn',
    'ts/no-unsafe-assignment': 'off',
    'ts/no-throw-literal': 'off',
    'node/prefer-global/process': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
    'eslint-comments/no-unlimited-disable': 'off',
  },
})
