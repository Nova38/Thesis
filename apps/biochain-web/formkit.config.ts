import { en } from '@formkit/i18n'
import {
  createLocalStoragePlugin,
  createAutoHeightTextareaPlugin,
  createMultiStepPlugin,
  createFloatingLabelsPlugin,
} from '@formkit/addons'
import '@formkit/addons/css/multistep'

// import { defineFormKitConfig } from "@formkit/vue";
import { genesisIcons } from '@formkit/icons'
import { defaultConfig } from '@formkit/vue'
import { generateClasses } from '@formkit/themes'

import { rootClasses } from './formkit.theme'
import '@formkit/addons/css/floatingLabels'

export default defaultConfig(
  // here we can access `useRuntimeConfig` because
  // our function will be called by Nuxt.
  // const config = useRuntimeConfig();

  {
    locales: { en },
    locale: 'en',
    config: {
      rootClasses,
      classes: generateClasses({}),
    },
    plugins: [
      createAutoHeightTextareaPlugin(),
      createMultiStepPlugin(),
      createLocalStoragePlugin({
        // plugin defaults:
        prefix: 'formkit',
        key: undefined,
        control: undefined,
        maxAge: 3600000, // 1 hour
        debounce: 200,
        beforeSave: undefined,
        beforeLoad: undefined,
      }),

      createFloatingLabelsPlugin({
        // useAsDefault: true, // defaults to false
      }),
    ],
    // theme: "genesis", // will load from CDN and inject into document head

    icons: {
      ...genesisIcons,
    },
  },
)
