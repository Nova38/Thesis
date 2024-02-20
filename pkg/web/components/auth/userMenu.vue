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
      <QChip>{{ $auth.username.value }} </QChip>
      <QBtn @click="authLogout">
        Logout
      </QBtn>
    </div>
  </div>
</template>

<style></style>
