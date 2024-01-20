import { defu } from 'defu'
import type { UseFetchOptions } from '#app'

export function useCustomFetch<T>(
  url: string | (() => string),
  options: UseFetchOptions<T> = {},
) {
  const config = useRuntimeConfig()
  const appConfig = useAppConfig()
  console.log({
    appconfig: appConfig.api,
    base: config.baseUrl,
    public: config.public.api.url,
  })
  console.log(config.public.api.url || '/api')
  const defaults: UseFetchOptions<T> = {
    baseURL: config.public.api.url || '/api',
    // config.public.api.url ?? config.app.baseURL,
    // this overrides the default key generation, which includes a hash of
    // url, method, headers, etc. - this should be used with care as the key
    // is how Nuxt decides how responses should be deduplicated between
    // client and server
    // key: url,

    // set user token if connected

    onResponse() {
      // _ctx.response._data = new myBusinessResponse(_ctx.response._data)
    },

    onRequest() {
      LoadingBar.start()
      LoadingBar.stop()
    },
  }

  // for nice deep defaults, please use unjs/defu
  const params = defu(options, defaults)

  return useFetch(url, params)
}
