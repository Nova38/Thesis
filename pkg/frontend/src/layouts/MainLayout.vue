<template>
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
        <!-- <q-separator vertical inset color="white" /> -->
        <p class="m-1"></p>
        <q-breadcrumbs
          active-color="white"
          style="font-size: 18px"
          class="q-pl-md"
        >
          <q-breadcrumbs-el label="Biochain" to="/" />
          <q-breadcrumbs-el
            v-if="props.collectionId"
            :label="props.collectionId"
            :to="'/collection/' + props.collectionId"
          />
          <q-breadcrumbs-el
            v-if="props.collectionId"
            :label="'Specimen ID: ' + props.specimenId"
            :to="'/specimen/' + props.specimenId"
          />
        </q-breadcrumbs>
        <!-- <q-toolbar-title> Biochain: {{ props.collectionId }} </q-toolbar-title> -->

        <q-space />

        <div v-if="!authStore.isLoggedIn" id="NotLoggedInHeader">
          <q-btn color="teal" label="Login" :to="'/login'" />
          <span style="padding: 0.25rem" />
          <q-btn color="deep-orange" label="Register" :to="'/Register'" />
        </div>

        <div v-if="authStore.isLoggedIn" id="LoggedInHeader">
          <q-chip>
            <q-avatar>
              <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
            </q-avatar>
            {{ authStore.username }}
          </q-chip>

          <q-btn color="red" label="Logout" @click="authStore.logout" />
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      elevated
      class="column"
    >
      <q-toolbar>
        <q-toolbar-title> Collections: </q-toolbar-title>
        <q-space />
      </q-toolbar>

      <q-list>
        <CollectionNavBarItem
          v-for="item in metaStore.collectionNames"
          :key="item"
          :model-value="item"
        />
      </q-list>
      <div v-if="useAuthStore().isLoggedIn" style="margin-top: auto">
        <q-btn
          color="secondary"
          class="full-width"
          label="Register Collection"
          to="/new_collection"
        >
          <q-tooltip class="bg-accent"
            >Create a new Collection in the blockchain</q-tooltip
          >
        </q-btn>
      </div>
    </q-drawer>

    <q-page-container>
      <div>
        <router-view />
      </div>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { PlainMessage } from '@bufbuild/protobuf';
import CollectionNavBarItem from 'src/components/CollectionNavBarItem.vue';
import { schema } from 'src/lib/ccbio';
import { useAuthStore } from 'stores/auth';
import { useMetaStore } from 'stores/meta';

const authStore = useAuthStore();
const metaStore = useMetaStore();
const props = defineProps<{
  collectionId?: string;
  specimenId?: string;
}>();

// watch(
//   () => props.collectionId,
//   () => {
//     if (!props.collectionId) return;

//     metaStore.setCurrentCollection(props.collectionId);
//   }
// );

if (metaStore.collections.items?.length === 0) {
  metaStore.refresh();
}

const leftDrawerOpen = ref(false);

function toggleLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
