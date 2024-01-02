<template>
  <q-card class="componentWrapper q-ma-lg">
    <!-- Header -->

    <q-card-section>
      <div class="row items-center">
        <div class="col">
          <h4 class="text-h6">Users</h4>
          <!-- {{ metaStore.users }} -->
        </div>
      </div>
    </q-card-section>
    <q-card-section>
      <q-table
        :columns="columnDef"
        :rows="metaStore.users.items"
        row-key="name"
      ></q-table>
    </q-card-section>
    <!-- Role Sections, Users are clickable and open a dialog -->
  </q-card>
</template>

<script script setup lang="ts">
import { QTableProps } from 'quasar';
import { schema } from 'src/lib/ccbio';
import { useMetaStore } from 'src/stores/meta';

const props = defineProps<{
  collectionId: string;
}>();

const metaStore = useMetaStore();

const columnDef: Ref<QTableProps['columns']> = ref([
  {
    name: 'name',
    label: 'name',
    field: 'name',
    sortable: true,
  },
  {
    name: 'msp_id',
    label: 'msp_id',
    field: 'msp_id',
    // cellRenderer: TableButton,
    sortable: true,
  },
  {
    name: 'roles',
    label: 'roles',
    field: (row) => {
      return (
        row.memberships[metaStore.selectedCollectionId] ||
        schema.Role.PUBLIC_UNSPECIFIED
      );
    },
    sortable: true,
  },
]);

// const users = useQuery(
//   ChaincodeQueries.collection.getById(collectionId.value)._ctx.users
// );
</script>

<style scoped lang="css">
#componentWrapper {
  margin: 0;
}
</style>
