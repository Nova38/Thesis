<script lang="ts" setup>
import { ref } from 'vue';
import { useVModel } from '@vueuse/core';

import { useQuasar } from 'quasar';
import { useAuthStore } from 'src/stores/auth';
import { useRouter, useRoute } from 'vue-router';

const authStore = useAuthStore();

const $q = useQuasar();

const router = useRouter();
// const route = useRoute();
const props = defineProps({
  modelValue: {
    type: Object,
    required: true,
  },
});
const submitting = ref(false);

const emit = defineEmits(['update:model-value', 'submit']);

const form = useVModel(props, 'modelValue', emit);
const isPwd = ref(true);

// const that = this;

const submit = async (username: string, password: string) => {
  submitting.value = true;
  const res = await authStore
    .login(username, password)
    .then(() => {
      submitting.value = false;
      $q.notify({
        type: 'positive',
        message: 'Login successful',
      });

      router.push('/').catch((err) => {
        console.log(err);
      });
    })
    .catch((err) => {
      submitting.value = false;
      console.log(err);
      $q.notify({
        type: 'negative',
        message: 'Login failed',
      });
    });
  console.log(res);
};
</script>

<template>
  <q-form @submit="submit(form.username, form.password)">
    <q-card class="q-pa-md">
      <q-input v-model="form.username" label="Username" />
      <!-- <q-input type="password" v-model="form.password" label="Password" /> -->
      <q-input
        v-model="form.password"
        :type="isPwd ? 'password' : 'text'"
        label="Password"
      >
        <template v-slot:append>
          <q-icon
            :name="isPwd ? 'visibility_off' : 'visibility'"
            class="cursor-pointer"
            @click="isPwd = !isPwd"
          />
        </template>
      </q-input>
      <q-card-actions align="right">
        <q-btn
          type="submit"
          :loading="submitting"
          color="primary"
          label="Login"
        />
      </q-card-actions>
    </q-card>
  </q-form>
</template>
