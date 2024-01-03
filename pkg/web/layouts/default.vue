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

          <template v-for="item in crumb" :key="item.label">
            <span class="px-2">
              <NuxtLink :to="item.to"> {{ item.label }} &nbsp; </NuxtLink>
              <Icon name="uil:angle-right" size="2em" />
            </span>
            <!-- <q-separator vertical inset color="white" /> -->
          </template>

          <q-space />
          <AuthUserMenu class="ml-auto" />

          <!-- <div v-if="!authStore.isLoggedIn" id="NotLoggedInHeader">
          <q-btn color="teal" label="Login" :to="'/login'" />
          <span style="padding: 0.25rem" />
          <q-btn color="deep-orange" label="Register" :to="'/Register'" />
        </div> -->

          <!-- <div v-if="authStore.isLoggedIn" id="LoggedInHeader">
          <q-chip>
            <q-avatar>
              <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
            </q-avatar>
            {{ authStore.username }}
          </q-chip>

          <q-btn color="red" label="Logout" @click="authStore.logout" />
        </div> -->
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
          <slot name="content"> </slot>
        </div>
      </q-page-container>
    </q-layout>

    <!-- <UCard
            class="v-full"
            :ui="{
                header: {
                    padding: 'px-2 py-2 sm:px-2',
                },
            }"
        > -->
    <!-- <template #header>
                <div class="flex flex-row items-center"></div>
            </template>
            <div class="flex flex-row w-full flex-nowrap">
                <div class="p-2 mr-4 flex-none grow-0">
                    <NavSidebar />
                </div>
                <div class="grow">
                    <UCard>
                        <template #header>
                            <div></div>
                        </template>
                        <slot name="content"> </slot>
                    </UCard>
                </div>
            </div>
        </UCard> -->
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
