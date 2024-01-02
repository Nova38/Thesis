<template>
  <q-page
    padding
    class="window-height window-width row justify-center items-center"
  >
    <register-form v-model="form" @submit="onSubmit" />
    <!-- <pre>{{ form }}</pre> -->
  </q-page>
</template>

<script setup lang="ts">
import RegisterForm from 'src/components/auth/RegisterForm.vue';
import { useAuthStore } from 'src/stores/auth';
import { ref } from 'vue';

const authStore = useAuthStore();
const form = ref({
  username: '',
  email: '',
  password: '',
});

const onSubmit = () => {
  // authStore.login(form.value.username, form.value.password);
};
</script>
