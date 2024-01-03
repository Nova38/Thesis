<template>
  <div>
    <div v-if="!loggedIn">
      <NuxtLink to="/auth/login">Login</NuxtLink>
    </div>
    <div v-else>
      <QChip>{{ $auth.username.value }} </QChip>
      <QBtn color="red" @click="authLogout">Logout</QBtn>
    </div>
  </div>
</template>

<script lang="ts" setup>
// import type { UButton } from "#build/components";
const loggedIn = useState("loggedIn");
</script>

<style></style>
