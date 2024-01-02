<template>
  <q-page padding class="full-height">
    <div class="row">
      <div v-if="collection.value" class="card">
        <span v-if="collection.value.access_control">
          <access-control-table
            :accessControl="collection.value.access_control"
            :isEditable="false"
          />
        </span>
      </div>

      <user-list :collectionId="props.collectionId" />
      <!-- todo:user role assignment -->
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { useMetaStore } from 'src/stores/meta';

import AccessControlTable from 'src/components/collection/AccessControlTable.vue';
import { ccApi } from 'boot/axios';
import { schema } from 'src/lib/ccbio';
import { reactiveComputed } from '@vueuse/core';
import UserList from 'src/components/collection/UserList.vue';

const splitterProportion = ref(50); // start at 50%

const meta = useMetaStore();

const props = defineProps<{
  collectionId: string;
}>();

const collection = reactiveComputed(() => {
  const current = meta.getCollection(props.collectionId);
  if (current) {
    return {
      value: new schema.Collection({ ...current }),
    };
  } else {
    return {
      value: undefined,
    };
  }
});
</script>

<style scoped lang="css">
.card {
  max-width: 850px;
}
</style>
