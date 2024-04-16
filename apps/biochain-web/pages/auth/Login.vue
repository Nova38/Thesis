<script lang="ts" setup>
import type { FormKitNode } from '@formkit/core'

const router = useRouter()

const submitting = ref(false)

const isPwd = ref(true)

// const that = this;
type Form = {
  username: string
  password: string
}

async function submit(data: Form, node: FormKitNode) {
  console.log('submitting', data)
  submitting.value = true
  const res = await authLogin(data.username, data.password)
    .then(() => {
      submitting.value = false

      router.push('/').catch((err) => {
        console.log(err)
      })
    })
    .catch((err) => {
      submitting.value = false
      console.log(err)
    })
  console.log(res)
}
</script>

<template>
  <UCard class="m-4">
    <FormKit
      type="form"
      @submit="submit"
    >
      <FormKit
        type="text"
        name="username"
        label="Username"
        validation="required"
      />
      <FormKit
        type="password"
        name="password"
        label="Password"
        validation="required"
      />
    </FormKit>
  </UCard>
</template>
