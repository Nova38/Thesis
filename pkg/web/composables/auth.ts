export const useAuth = () => {
  const nuxtApp = useNuxtApp();
  return nuxtApp.$auth;
};
const useLoggedIn = useState("loggedIn");

// const { $auth } = useNuxtApp();

export const authLogin = async (username: string, password: string) => {
  try {
    const result = await useCustomFetch("/api/auth/login", {
      method: "POST",
      body: {
        username,
        password,
      },
    });
    console.log(result);
    if (result.error.value) {
      // useError;
      throw result.error.value;
    }
    useLoggedIn.value = true;
    useAuth().redirectTo.value = null;
    await useAuth().updateSession();
    await navigateTo(useAuth().redirectTo.value || "/");

    return result;
  } catch (e) {
    console.log(e);
    createError(e ?? {});
    throw e;
  }
};

export const authRegister = async (username: string, password: string) => {
  await $fetch("/api/auth/register", {
    method: "POST",
    body: {
      username,
      password,
    },
  });
  return await authLogin(username, password);
};

export const authLogout = async () => {
  await $fetch("/api/auth/logout", {
    method: "POST",
  });
  useLoggedIn.value = false;
  await useAuth().updateSession();
};
