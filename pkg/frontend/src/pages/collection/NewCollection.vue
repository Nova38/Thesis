<template>
  <q-page padding>
    <q-card class="card">
      <q-card-section id="card-header">
        <h4>Register New Collection:</h4>
      </q-card-section>

      <q-card-section id="card-content">
        <q-input
          outlined
          v-if="collection.id"
          v-model="collection.id.collectionId"
          label="Collection Name"
          hint="This is the name of the collection"
        >
        </q-input>

        <access-control-table
          v-if="collection.accessControl"
          :access-control="collection.accessControl"
          :is-editable="true"
        />
      </q-card-section>

      <q-card-actions>
        <q-btn
          class="full-width"
          color="primary"
          label="Create"
          @click="createCollection"
          :disable="!collection.id"
        />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import { ccApi } from 'src/boot/axios';
import AccessControlTable from 'src/components/collection/AccessControlTable.vue';
import { schema, utils } from 'src/lib/ccbio';

// import AccessControlTable from 'src/components/collection/AccessControlTable.vue';
import { useMetaStore } from 'src/stores/meta';
// import { CC } from 'src/api';

const collection = ref<schema.Collection>(
  new schema.Collection({
    id: { collectionId: 'test' },
    accessControl: utils.MakeBaseAccessControls(),
  })
);

const meta = useMetaStore();
// const CollectionNames = ;

function createCollection() {
  console.log('createCollection');
  console.log(collection.value);

  const payload = new schema.CollectionCreateRequest({
    collection: {
      id: collection.value.id,
      accessControl: { ...collection.value.accessControl },
    },
  });
  console.log(payload);

  // CC.collectionRegister(payload);
  // TODO: create collection

  ccApi.auth.CollectionCreate(payload).then((res) => {
    console.log(res);
    meta.collections.items.push(res);
  });
}
</script>

<style scoped lang="css">
.card {
  max-width: 1000px;
  margin: 20px auto;
}

#card-header {
  text-align: center;
}

h4 {
  margin: 0;
  border-bottom: 2px solid #ddd;
}
</style>
