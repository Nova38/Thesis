// module.exports = {
//     root: true,

//     env: {
//         "browser": true,
//         "es2021": true,
//         "node": true
//     },
//     "extends": [
//         "standard-with-typescript",
//         "plugin:vue/vue3-essential",
//         "@nuxt/eslint-config"
//     ],
//     "overrides": [
//         {
//             "env": {
//                 "node": true
//             },
//             "files": [
//                 ".eslintrc.{js,cjs}"
//             ],
//             "parserOptions": {
//                 "sourceType": "script"
//             }
//         }
//     ],
//     "parserOptions": {
//         "ecmaVersion": "latest",
//         "sourceType": "module",
//         parser: "@typescript-eslint/parser",
//     },
//     "plugins": [
//         "vue"
//     ],
//     "rules": {
//     }
// }
module.exports = {
  env: {
    browser: true,
    es2021: true,
    node: true,
  },
  extends: ['@nuxt/eslint-config', 'plugin:prettier/recommended'],
  // ignorePatterns: [],
  parser: 'vue-eslint-parser',

  parserOptions: {
    parser: '@typescript-eslint/parser',
  },
  root: true,
}
