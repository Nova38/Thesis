<script lang="ts" setup></script>

<template>
  <div class="min-w-15">
    <div>
      <PToolbar class="bg-blue-600 grid-cols-3">
        <template #start>
          <UButton label="Click Me" />
          <PButton
            icon="pi pi-print"
            class="mr-2"
            severity="secondary"
          />
          <PButton
            icon="pi pi-upload"
            severity="secondary"
          />
        </template>

        <template #center />

        <template #end>
          <AuthUserMenu class="ml-auto" />
        </template>
      </PToolbar>
    </div>
  </div>
</template>

<style></style>
