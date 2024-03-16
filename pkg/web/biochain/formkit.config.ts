import { en } from '@formkit/i18n'

// import { defineFormKitConfig } from "@formkit/vue";
import { genesisIcons } from '@formkit/icons'
import { defaultConfig } from '@formkit/vue'
import { createFloatingLabelsPlugin } from '@formkit/addons'
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
