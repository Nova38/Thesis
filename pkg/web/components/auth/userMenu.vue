<script lang="ts" setup>
// import type { UButton } from "#build/components";
const loggedIn = useState('loggedIn')

const auth = useAuth()

async function authLogout() {
  await $fetch('/api/auth/logout', {
    method: 'POST',
  })
  loggedIn.value = false

  auth.clearSession()
}
</script>

<template>
  <div>
    <div v-if="!loggedIn">
      <NuxtLink to="/auth/login">
        Login
      </NuxtLink>
    </div>
    <div v-else>
      <UChip>{{ $auth.username.value }} </UChip>
      <UButton @click="authLogout">
        Logout
      </UButton>
    </div>
  </div>
</template>

<style></style>
