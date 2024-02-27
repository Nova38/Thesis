<script lang="ts" setup>
const crumb = useBreadcrumbLinks()
</script>

<template>
  <div class="min-w-15">
    <div>
      <PToolbar class="bg-blue-600 grid-cols-3">
        <template #start>
          <!-- <UButton label="Click Me" /> -->
          <PButton class="mr-2">
            <Icon
              name="carbon:menu"
              size="1em"
            />
          </PButton>
          <div class="flex flex-row items-center">
            <template
              v-for="item in crumb"
              :key="item.label"
            >
              <span class="px-2">
                <Icon
                  name="uil:angle-right"
                  size="2em"
                />
              </span>
              <NuxtLink
                :to="item.to"
                active-class="font-bold"
                exact-active-class="font-bold"
              >
                {{ item.label }} &nbsp;
              </NuxtLink>

              <!-- <q-separator vertical inset color="white" /> -->
            </template>
          </div>
        </template>

        <template #center />

        <template #end>
          <ColorModeSelector />
          <AuthUserMenu class="ml-auto" />
        </template>
      </PToolbar>
    </div>
  </div>
</template>

<style></style>
