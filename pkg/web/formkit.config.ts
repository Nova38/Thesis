import { en } from "@formkit/i18n";
// import { defineFormKitConfig } from "@formkit/vue";
import { rootClasses } from "./formkit.theme";
import { genesisIcons } from "@formkit/icons";
import { defaultConfig } from "@formkit/vue";
import { createFloatingLabelsPlugin } from "@formkit/addons";
import "@formkit/addons/css/floatingLabels";

export default defaultConfig(
  // here we can access `useRuntimeConfig` because
  // our function will be called by Nuxt.
  // const config = useRuntimeConfig();

  {
    locales: { en },
    locale: "en",
    config: {
      rootClasses,
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
);
