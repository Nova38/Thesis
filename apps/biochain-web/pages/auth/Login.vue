<script lang="ts" setup>
import type { FormKitNode } from '@formkit/core'

// const that = this;
type Form = {
  username: string
  password: string
}
</script>

<template>
  <UCard class="m-4">
    <FormKit
      type="form"
      @submit="(data) => authLogin(data.username, data.password)"
    >
      <FormKit
        type="text"
        name="username"
        label="Username"
        validation="required"
        help="Please enter your username (without t- prefix)"
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
