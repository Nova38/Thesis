<template>
  <div>
    <q-layout view="Hhh Lpr lFf">
      <q-header elevated>
        <q-toolbar>
          <q-btn
            flat
            dense
            round
            icon="menu"
            aria-label="Menu"
            @click="toggleLeftDrawer"
          />
          <span class="w-4" />
          <div class="flex flex-row items-center">
            <template v-for="item in crumb" :key="item.label">
              <span class="px-2">
                <Icon name="uil:angle-right" size="2em" />
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

          <!-- <q-space /> -->
          <AuthUserMenu class="ml-auto" />
        </q-toolbar>
      </q-header>
      <q-drawer
        v-model="leftDrawerOpen"
        show-if-above
        bordered
        elevated
        class="flex-"
      >
        <!-- <q-toolbar>
                    <q-toolbar-title> Collections: </q-toolbar-title>
                    <q-space />
                </q-toolbar> -->
        <NavSidebar />
      </q-drawer>
      <q-page-container>
        <div class="p-2">
          <NuxtLoadingIndicator />
          <NuxtErrorBoundary>
            <slot name="content"> </slot>

            <!-- ... -->
            <template #error="{ error }">
              <p>An error occurred: {{ error }}</p>
            </template>
          </NuxtErrorBoundary>
        </div>
      </q-page-container>
    </q-layout>
  </div>
</template>

<script lang="ts" setup>
// const crumb = [
//     {
//         label: "Home",
//         icon: "i-heroicons-home",
//         to: "/",
//     },
//     {
//         label: "Navigation",
//         icon: "i-heroicons-square-3-stack-3d",
//     },
//     {
//         label: "Breadcrumb",
//         icon: "i-heroicons-link",
//     },
// ];
const crumb = useBreadcrumbLinks();
const leftDrawerOpen = ref(true);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
LoadingBar.setDefaults({
  color: "purple",
  size: "15px",
  position: "bottom",
});
</script>

<style></style>
