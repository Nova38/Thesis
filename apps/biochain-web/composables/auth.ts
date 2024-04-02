export function useAuth() {
  const nuxtApp = useNuxtApp()
  return nuxtApp.$auth
}
const useLoggedIn = () => useState('loggedIn')

// const { $auth } = useNuxtApp();

export async function authLogin(username: string, password: string) {
  try {
    // const useLoggedIn = () => useState('loggedIn')

    const result = await useCustomFetch('/api/auth/login', {
      body: {
        password,
        username,
      },
      method: 'POST',
    })
    console.log(result)
    if (result.error.value) {
      // useError;
      throw result.error.value
    }
    useLoggedIn().value = true
    useAuth().redirectTo.value = null
    await useAuth().updateSession()
    await navigateTo(useAuth().redirectTo.value || '/')

    return result
  } catch (e) {
    console.log(e)
    createError(e ?? {})
    throw e
  }
}

export async function authRegister(username: string, password: string) {
  await $fetch('/api/auth/register', {
    body: {
      password,
      username,
    },
    method: 'POST',
  })
  return await authLogin(username, password)
}

export async function authLogout() {
  const useLoggedIn = useState('loggedIn')

  await $fetch('/api/auth/logout', {
    method: 'POST',
  })
  useLoggedIn.value = false
  await useAuth().updateSession()
}
