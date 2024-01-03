export const useAuth = () => {
  const nuxtApp = useNuxtApp();
  return nuxtApp.$auth;
};
const useLoggedIn = useState("loggedIn");

// const { $auth } = useNuxtApp();

export const authLogin = async (username: string, password: string) => {
  try {
    useCustomFetch("/api/auth/login", {
      method: "POST",
      body: {
        username,
        password,
      },
    });
    useAuth().redirectTo.value = null;
    await useAuth().updateSession();
    await navigateTo(useAuth().redirectTo.value || "/");
    useLoggedIn.value = true;
  } catch (e) {
    console.log(e);
    return e;
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
  await useAuth().updateSession();
  useLoggedIn.value = false;
};
