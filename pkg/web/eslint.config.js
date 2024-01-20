import antfu from "@antfu/eslint-config";
import NuxtEslintConfig from "./.nuxt/eslint.config.mjs";

export default antfu(
  {
    // ...@antfu/eslint-config options,
    stylistic: true,
    typescript: {
      tsconfigPath: "tsconfig.json",
    },
    vue: true,
    formatters: {
      markdown: true,
      css: true,
      html: true,
    },
  },
  {
    rules: {
      "no-console": "warn",
    },
  },

  // Add the Nuxt rules
  NuxtEslintConfig,
  // ...your other rules
);
