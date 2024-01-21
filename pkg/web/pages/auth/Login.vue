<script lang="ts" setup>
const $q = useQuasar()
const form = ref({
  password: '',
  username: '',
})

const router = useRouter()

const submitting = ref(false)

const isPwd = ref(true)

// const that = this;

async function submit(username: string, password: string) {
  submitting.value = true
  const res = await authLogin(username, password)
    .then(() => {
      submitting.value = false
      $q.notify({
        message: 'Login successful',
        type: 'positive',
      })

      router.push('/').catch((err) => {
        console.log(err)
      })
    })
    .catch((err) => {
      submitting.value = false
      console.log(err)
      $q.notify({
        message: 'Login failed',
        type: 'negative',
      })
    })
  console.log(res)
}
</script>

<template>
  <q-form @submit="submit(form.username, form.password)">
    <q-card class="flex flex-row gap-2 p-4 items-center">
      <q-input
        v-model="form.username"
        label="Username"
      />
      <!-- <q-input type="password" v-model="form.password" label="Password" /> -->
      <q-input
        v-model="form.password"
        :type="isPwd ? 'password' : 'text'"
        label="Password"
      >
        <template #append>
          <q-icon
            :name="isPwd ? 'visibility_off' : 'visibility'"
            class="cursor-pointer"
            @click="isPwd = !isPwd"
          />
        </template>
      </q-input>
      <q-card-actions align="right">
        <q-btn
          :loading="submitting"
          color="primary"
          label="Login"
          type="submit"
        />
      </q-card-actions>
    </q-card>
  </q-form>
</template>
