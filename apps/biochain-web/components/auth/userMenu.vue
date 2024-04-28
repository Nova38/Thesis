<script lang="ts" setup>
// import type { UButton } from "#build/components";

const auth = useAuth()

const UserDisplay = computed(() => {
  if (auth.username.value) return auth.username.value || ''
  return ''
})
</script>

<template>
  <div>
    <div v-if="!auth.loggedIn.value">
      <NuxtLink to="/auth/login"> Login </NuxtLink>
    </div>
    <div v-else>
      <UBadge
        variant="soft"
        class="mx-2"
        color="primary"
        :ui="{ rounded: 'rounded-full' }"
        :label="UserDisplay"
      />

      <UButton @click="authLogout"> Logout </UButton>
    </div>
  </div>
</template>

<style></style>
