// import type { AuthSession } from '~/server/utils/session'

export default defineNuxtPlugin({
  name: 'auth',
  async setup() {
    //   const { data: session, refresh: refreshSession } =
    //     await useCustomFetch<AuthSession>('/api/auth/session')

    const session = await $fetch('/api/auth/session')

    const username = useState('username', () => {
      return session?.username
    })
    const loggedIn = useState('loggedIn')
    loggedIn.value = !!username.value

    const updateSession = async () => {
      const tmp = await $fetch('/api/auth/session')
      console.log(tmp)
      username.value = session?.username ?? ''
      loggedIn.value = !!username.value
    }
    const clearSession = () => {
      username.value = ''
      loggedIn.value = ''
    }

    return {
      provide: {
        auth: {
          loggedIn,

          updateSession,
          username,
          clearSession,
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
