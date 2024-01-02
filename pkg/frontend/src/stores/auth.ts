import { defineStore, acceptHMRUpdate } from 'pinia';
import { api } from 'boot/axios';
import { schema } from 'src/lib/ccbio';
import { ccApi } from 'boot/axios';
import { PlainMessage } from '@bufbuild/protobuf';

export const useAuthStore = defineStore('auth', () => {
  const username = ref<string>('');
  const access_token = ref<string>('');
  const isLoggedIn = computed(() => !!access_token.value);
  const blockchainUser = ref<PlainMessage<schema.User>>(new schema.User());

  // Initially load state from local storage
  const localStorageAuth = {
    username: localStorage.getItem('username') || '',
    access_token: localStorage.getItem('access_token') || '',
  };
  if (localStorageAuth.username && localStorageAuth.access_token) {
    access_token.value = localStorageAuth.access_token;
    username.value = localStorageAuth.username;
  }

  if (localStorage.getItem('who')) {
    blockchainUser.value = JSON.parse(localStorage.getItem('who') || '');
  }

  // getters
  const header = computed(() => {
    return { Authorization: `Bearer ${access_token.value}` };
  });

  // Actions

  async function login(name: string, password: string) {
    const res = await api.post('/auth/login', {
      username: name,
      password: password,
    });

    if (!res.data.access_token) {
      return res;
    }

    username.value = name;
    access_token.value = res.data.access_token;

    localStorage.setItem('access_token', res.data.access_token);
    localStorage.setItem('username', name);

    setHeaders();

    const who = await ccApi.auth.GetCurrentUser();
    blockchainUser.value = who;
    console.log(who);

    if (who) {
      // assuming `who` has a `user` property
      localStorage.setItem('who', JSON.stringify(who));
      // username.value = who.user;
    }
  }

  function setHeaders() {
    if (access_token.value) {
      api.defaults.headers.common[
        'Authorization'
      ] = `Bearer ${access_token.value}`;
      console.log('headers set');
    }
  }

  async function refreshWho() {
    const who = await ccApi.auth.GetCurrentUser();
    return (blockchainUser.value = who);
  }

  async function logout() {
    username.value = '';
    access_token.value = '';
    localStorage.removeItem('auth');
  }
  async function checkUserName(username: string) {
    // const res = await api.post('/auth/checkUsername', {
    //   username,
    // });
    console.log(username);

    // TODO: check username on blockchain

    return false;
  }
  return {
    username,
    access_token,
    isLoggedIn,
    login,
    logout,
    header,
    blockchainUser,
    refreshWho,
    checkUserName,
    setHeaders,
  };
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useAuthStore, import.meta.hot));
}
