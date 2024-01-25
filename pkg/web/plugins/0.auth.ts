import type { AuthSession } from '~/server/utils/session'

export default defineNuxtPlugin(async (nuxtApp) => {
  // Skip plugin when rendering error page
  if (nuxtApp.payload.error)
    return {}

  const { data: session, refresh: refreshSession }
    = await useCustomFetch<AuthSession>('/api/auth/session')

  const username = useState('username', () => {
    return session.value?.username
  })
  const loggedIn = useState('loggedIn')
  loggedIn.value = !!username.value

  // Create a ref to know where to redirect the user when logged in
  const redirectTo = useState('authRedirect')

  const updateSession = async () => {
    const tmp = await refreshSession()
    console.log(tmp)
    username.value = session.value?.username ?? ''
    loggedIn.value = !!username.value
  }

  /**
   * Add global route middleware to protect pages using:
   *
   * definePageMeta({
   *  auth: true
   * })
   */

  addRouteMiddleware(
    'auth',
    (to) => {
      if (to.meta.auth && !loggedIn.value) {
        redirectTo.value = to.path
        return '/login'
      }
    },
    { global: true },
  )

  const currentRoute = useRoute()

  if (import.meta.client) {
    watch(loggedIn, async (loggedIn) => {
      if (!loggedIn && currentRoute.meta.auth) {
        redirectTo.value = currentRoute.path
        await navigateTo('/login')
      }
    })
  }

  if (loggedIn.value && currentRoute.path === '/login')
    await navigateTo(redirectTo.value || '/')

  return {
    provide: {
      auth: {
        loggedIn,
        redirectTo,
        refreshSession,
        session,
        updateSession,
        username,
      },
    },
  }
})
