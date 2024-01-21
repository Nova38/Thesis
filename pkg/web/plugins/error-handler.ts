export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.config.errorHandler = (error: any, instance, info) => {
    // handle error, e.g. report to a service
    // useQuasar();
    console.log('Error Handler', { error, info, instance })
  }

  // Also possible
  nuxtApp.hook('vue:error', () => {
    // error, instance, info
    // handle error, e.g. report to a service
  })
})
