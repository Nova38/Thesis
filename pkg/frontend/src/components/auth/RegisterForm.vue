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
const isPwd = ref(true);

const props = defineProps({
  modelValue: {
    type: Object,
    required: true,
  },
});
const submitting = ref(false);

const emit = defineEmits(['update:model-value', 'submit']);

const form = useVModel(props, 'modelValue', emit);

// const that = this;

const submit = async (username: string, email: string, password: string) => {
  submitting.value = true;
  const userNameInUse = await authStore
    .checkUserName(username)
    .then((res) => {
      if (res.data.usernameExists) {
        $q.notify({
          type: 'negative',
          message: 'Username already in use',
        });
        return;
      }
    })
    .catch((err) => {
      console.log(err);
    });

  const res = await authStore
    .register(username, email, password)
    .then(() => {
      submitting.value = false;
      $q.notify({
        type: 'positive',
        message: 'Registration successful',
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
        message: 'Registration failed',
      });
    });
  console.log(res);
};
</script>

<template>
  <q-form @submit="submit(form.username, form.email, form.password)">
    <q-card class="q-pa-md">
      <q-input v-model="form.username" label="Username" />
      <q-input v-model="form.email" label="Email" />
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
          label="Register"
        />
      </q-card-actions>
    </q-card>
  </q-form>
</template>
