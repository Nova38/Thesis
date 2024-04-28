// import type { AuthSession } from '~/server/utils/session'

// derived from https://github.com/nuxt/examples/blob/main/examples/auth/local/auth/plugins/0.auth.ts

export default defineNuxtPlugin({
  name: 'auth',

  async setup() {
    //   const { data: session, refresh: refreshSession } =
    //     await useCustomFetch<AuthSession>('/api/auth/session')

    const { data: session, refresh } = await useFetch('/api/auth/session')

    const username = computed(() => session.value?.username)
    const redirectTo = useState('authRedirect', () => '/')

    const loggedIn = computed(() => !!username.value)

    const isAdmin = computed(() => username.value === 'thomas')

    addRouteMiddleware(
      'auth',
      (to) => {
        if (to.meta.auth && !loggedIn.value) {
          redirectTo.value = to.path
          return '/auth/login'
        }
      },
      { global: true },
    )

    return {
      provide: {
        auth: {
          loggedIn,
          redirectTo,
          refresh,
          isAdmin,
          username,
        },
      },
    }
  },
})

// // export default defineNuxtPlugin(async (nuxtApp) => {
//   // Skip plugin when rendering error page
//   if (nuxtApp.payload.error) return {}

//   const { data: session, refresh: refreshSession } =
//     await useCustomFetch<AuthSession>('/api/auth/session')

//   const username = useState('username', () => {
//     return session.value?.username
//   })
//   const loggedIn = useState('loggedIn')
//   loggedIn.value = !!username.value

//   // Create a ref to know where to redirect the user when logged in
//   const redirectTo = useState('authRedirect')

//   const updateSession = async () => {
//     const tmp = await refreshSession()
//     console.log(tmp)
//     username.value = session.value?.username ?? ''
//     loggedIn.value = !!username.value
//   }
//   const clearSession = () => {
//     username.value = ''
//     loggedIn.value = ''
//   }

//   /**
//    * Add global route middleware to protect pages using:
//    *
//    * definePageMeta({
//    *  auth: true
//    * })
//    */

//   // addRouteMiddleware(
//   //   'auth',
//   //   (to) => {
//   //     if (to.meta.auth && !loggedIn.value) {
//   //       redirectTo.value = to.path
//   //       return '/login'
//   //     }
//   //   },
//   //   { global: true },
//   // )

//   // if (import.meta.client) {
//   //   watch(loggedIn, async (loggedIn) => {
//   //     if (!loggedIn && currentRoute.meta.auth) {
//   //       redirectTo.value = currentRoute.path
//   //       await navigateTo('/login')
//   //     }
//   //   })
//   // }

//   // if (loggedIn.value && currentRoute.path === '/login')
//   //   await navigateTo(redirectTo.value || '/')

//   return {
//     provide: {},
//   }
// })
