import type { UseFetchOptions } from "#app";
import { defu } from "defu";

export function useCustomFetch<T>(
  url: string | (() => string),
  options: UseFetchOptions<T> = {},
) {
  const config = useRuntimeConfig();

  const defaults: UseFetchOptions<T> = {
    baseURL: config.baseUrl ?? "/api",
    // this overrides the default key generation, which includes a hash of
    // url, method, headers, etc. - this should be used with care as the key
    // is how Nuxt decides how responses should be deduplicated between
    // client and server
    // key: url,

    // set user token if connected

    onResponse() {
      // _ctx.response._data = new myBusinessResponse(_ctx.response._data)
      LoadingBar.stop();
    },

    onRequest() {
      LoadingBar.start();
    },
  };

  // for nice deep defaults, please use unjs/defu
  const params = defu(options, defaults);

  return useFetch(url, params);
}
