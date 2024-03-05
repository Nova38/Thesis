<script lang="ts" setup></script>

<template>
  <div>
    <q-layout view="Hhh Lpr lFf">
      <q-header elevated>
        <q-toolbar>
          <q-btn
            aria-label="Menu"
            dense
            flat
            icon="menu"
            round

            @click="toggleLeftDrawer"
          />
          <span class="w-4" />
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
          <!-- <q-space /> -->
          <AuthUserMenu class="ml-auto" />
        </q-toolbar>
      </q-header>
      <q-drawer
        v-model="leftDrawerOpen"
        bordered
        class="flex-"
        elevated
        show-if-above
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
          <slot name="content" />
          <!-- <NuxtErrorBoundary> -->

          <!-- ... -->
          <!-- <template #error="{ error }"> -->
          <!-- <p>An error occurred: {{ error }}</p> -->
          <!-- </template> -->
          <!-- </NuxtErrorBoundary> -->
        </div>
      </q-page-container>
    </q-layout>
  </div>
</template>

<style scoped></style>
