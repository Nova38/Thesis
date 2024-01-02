import { boot } from 'quasar/wrappers';
import axios, { AxiosError, AxiosInstance } from 'axios';

import { api as cc } from '../lib/ccbio';

import { Notify } from 'quasar';

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
    $api: AxiosInstance;
    $ccApi: cc.CCBioApi;
  }
}

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
// const api = axios.create({ baseURL: 'https://api-biochain.ittc.ku.edu:8080/' });
const api = axios.create({ baseURL: 'https://api-biochain.ittc.ku.edu/' });
const ccApi = new cc.CCBioApi(api);

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API

  app.config.globalProperties.$ccApi = ccApi;
});

api.interceptors.response.use(
  function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
  },
  function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    const e = error as AxiosError;

    const ccFunction = e.request.responseURL.split('/').pop();

    const message = error.response;
    console.log(message);

    // or with a config object:
    Notify.create({
      message: `${ccFunction} ${error.response.data.message}`,
      icon: 'warning_amber',
      color: 'negative',
    });

    return Promise.reject(error);
  }
);

export { api, ccApi };
