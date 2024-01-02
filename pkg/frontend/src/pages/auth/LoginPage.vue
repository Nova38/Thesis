<template>
  <q-page
    padding
    class="window-height window-width column justify-center items-center"
  >
    <login-form v-model="form" @submit="onSubmit" />
  </q-page>
</template>

<script setup lang="ts">
import LoginForm from 'src/components/auth/LoginForm.vue';
import { useAuthStore } from 'src/stores/auth';
import { ref } from 'vue';

const authStore = useAuthStore();
const form = ref({
  username: '',
  password: '',
});

const onSubmit = () => {
  authStore.login(form.value.username, form.value.password);
};
</script>
