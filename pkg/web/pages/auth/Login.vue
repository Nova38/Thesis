<script lang="ts" setup>
const $q = useQuasar()
const form = ref({
  username: '',
  password: '',
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
        type: 'positive',
        message: 'Login successful',
      })

      router.push('/').catch((err) => {
        console.log(err)
      })
    })
    .catch((err) => {
      submitting.value = false
      console.log(err)
      $q.notify({
        type: 'negative',
        message: 'Login failed',
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
          type="submit"
          :loading="submitting"
          color="primary"
          label="Login"
        />
      </q-card-actions>
    </q-card>
  </q-form>
</template>
