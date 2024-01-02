<template>
  <q-page padding>
    <!-- content -->
    <div class="row">
      <!-- Option  toolbar?-->

      <!-- Current State -->

      <q-card class="q-ma-md">
        <!-- <span v-if="status == 'loading'">Loading...</span>
        <span v-else-if="status == 'error'">Error: {{ error }}</span>
        <span v-else> -->
        <!-- <q-toolbar></q-toolbar> -->
        <SpecimenForm
          :specimen="formSpecimen"
          :enable-edit="edit"
          id="MainSpecimenForm"
          :editPermission="editPerms"
        >
          <template #HeaderBar>
            <q-bar>
              <div>Current Mode: {{ formMode }}</div>

              <q-space />
              <q-btn @click="() => (formMode = 'View')">
                View
                <q-tooltip>Set the current mode to View</q-tooltip>
              </q-btn>
              <q-btn name="Update" @click="() => (formMode = 'Update')">
                Update
                <q-tooltip>Set the current mode to Update</q-tooltip>
              </q-btn>
              <q-btn @click="() => (formMode = 'Suggest')">
                Suggest Update
                <q-tooltip>Set the current mode to Suggest Update</q-tooltip>
              </q-btn>
            </q-bar>
          </template>
        </SpecimenForm>
        <q-card-actions v-if="formMode !== 'View'">
          <q-btn class="full-width" color="secondary" @click="handleSubmit">
            {{ formMode }}
          </q-btn>
        </q-card-actions>
        <!-- </span> -->

        <!-- TODO Form actions -->

        <q-inner-loading :showing="loading">
          <q-spinner-gears size="50px" color="primary" />
        </q-inner-loading>
      </q-card>

      <!-- History -->
      <q-card class="q-pa-lg q-ma-md">
        <q-card-section>
          <div class="text-h6">History</div>
        </q-card-section>
        <q-card-section>
          <!-- <span v-if="historyStatus == 'loading'">Loading...</span>
          <span v-else-if="historyStatus == 'error'"
            >Error: {{ historyError }}</span
          >
          <span v-else> -->
          <q-timeline dense color="secondary" v-if="history">
            <template v-for="tx in history.entries" :key="tx.tx_id">
              <TransactionView
                :tx_id="tx.txId"
                :timestamp="tx.timestamp || new Timestamp()"
                :is_deleted="tx.isDeleted || false"
                :specimen="MaybePlainSchemaToForm(tx.state)"
                :hiddenInfo="extractHiddenInfo(tx.txId) || undefined"
                :isLast="history.entries[history.entries.length - 1] == tx"
              />
            </template>
          </q-timeline>
          <!-- </span> -->
        </q-card-section>
      </q-card>

      <!-- Suggested State -->
      <!-- TODO: All of suggested state -->
    </div>
  </q-page>
</template>

<script setup lang="ts">
import SpecimenForm from 'src/components/specimen/SpecimenForm.vue';
import TransactionView from 'src/components/specimen/TransactionView.vue';

// Schema imports
import { schema } from 'src/lib/ccbio';

import {
  FormSpecimenToSchema,
  SchemaToFormSpecimen,
  MaybePlainSchemaToForm,
  PopulateNestedRequired,
} from 'src/lib/ccbio/utils/specimen';

import { ccApi } from 'src/boot/axios';
import { useMetaStore } from 'src/stores/meta';
import { PlainMessage, Timestamp } from '@bufbuild/protobuf';
// import { BiochainCC } from 'src/api/api';

const meta = useMetaStore();

const props = defineProps<{
  specimenId: string;
  collectionId: string;
}>();

// State
type formStates = 'View' | 'Update' | 'Suggest';

const formMode = ref<formStates>('View');
const loading = ref(false);

const editPerms = computed(() => {
  if (formMode.value == 'View') return new schema.Permissions();
  if (formMode.value == 'Update')
    return meta.currentEditPermissions || new schema.Permissions();
  if (formMode.value == 'Suggest')
    return meta.currentSuggestEditPermissions || new schema.Permissions();

  throw new Error('Invalid formMode');
});

const specimen_id: Ref<PlainMessage<schema.Specimen_Id>> = ref(
  new schema.Specimen_Id({
    id: props.specimenId,
    collectionId: props.collectionId,
  })
);
watch(
  () => props,
  () => {
    console.log(props);
  }
);

watch(
  () => props.specimenId,
  () => {
    console.log(`specimenId changed: ${props.specimenId}}`);
    specimen_id.value = new schema.Specimen_Id({
      id: props.specimenId,
      collectionId: props.collectionId,
    });
  }
);

const edit = computed(() => {
  return formMode.value == 'Update' || formMode.value == 'Suggest';
});



// const specimen_id = computed(() => {
//   console.log(props);
//   return new schema.Specimen_Id({
//     id: props.specimenId,
//     collectionId: props.collectionId,
//   });
// });

const specimenData = ref(new schema.Specimen());
const formSpecimen = computed(() => {
  return SchemaToFormSpecimen(new schema.Specimen(specimenData.value));
});

const history = ref<PlainMessage<schema.Specimen_History>>(
  new schema.Specimen_History()
);

// Methods
function handleSubmit() {
  loading.value = true;

  if (formMode.value == 'View') {
    console.warn('Cannot submit in View mode');
    loading.value = false;
  }
  if (formMode.value == 'Update') {
    const specimen = FormSpecimenToSchema(formSpecimen.value);

    const payload = new schema.SpecimenUpdateRequest({
      specimen: {
        id: specimen_id.value,
        primary: specimen.primary,
        secondary: specimen.secondary,
        georeference: specimen.georeference,
        grants: specimen.grants,
        taxon: specimen.taxon,
        images: specimen.images,
        loans: specimen.loans,
      },
    });

    PopulateNestedRequired(payload);
    console.log(payload);

    // Todo: Notify or update the user on the status of the transaction
    ccApi.specimen
      .SpecimenUpdate(payload)
      .then((res) => {
        console.log(res);
        loading.value = false;
        refreshHistory();
      })
      .catch((err) => {
        console.error(err);
        loading.value = false;
      });
  }
  if (formMode.value == 'Suggest') {
  }
}

function extractHiddenInfo(tx_id: string) {
  if (!specimenData.value) return null;
  if (!specimenData.value?.hiddenTxs) return null;

  const tx = specimenData.value.hiddenTxs[tx_id] || null;
  if (!tx) return null;
  return tx;
}

function refreshHistory() {
  // Get the history
  ccApi.specimen
    .GetSpecimenHistory(
      new schema.GetSpecimenHistoryRequest({
        id: specimen_id.value,
      })
    )
    .then((data) => {
      console.log(data);
      history.value = data;
    })
    .catch((error) => {
      console.error(error);
    });
}

// const { status, data, error } = useQuery(
//   ChaincodeQueries.specimen.getById(specimen_id.value)
// );

// const specimenData = computed(() => {
//   if (!data.value) return schema.Specimen.create();
//   return data.value;
// });

// const {
//   status: historyStatus,
//   data: history,
//   error: historyError,
// } = useQuery(ChaincodeQueries.specimen.getById(specimen_id.value)._ctx.history);

// const {
//   status: suggestedStatus,
//   data: suggested,
//   error: suggestedError,
// } = useQuery(
//   ChaincodeQueries.specimen.getById(specimen_id.value)._ctx.suggested
// );
function useTx(
  x: schema.Specimen_History_Entry | Ref<schema.Specimen_History_Entry>
) {
  return toValue(x);
}
watch(
  props,
  async () => {
    try {
      specimen_id.value = new schema.Specimen_Id({
        id: props.specimenId,
        collectionId: props.collectionId,
      });

      // meta.selectedCollection(props.collectionId);

      ccApi.specimen
        .GetSpecimen(new schema.GetSpecimenRequest({ id: specimen_id.value }))
        .then((data) => {
          specimenData.value = data;
          PopulateNestedRequired(specimenData.value);
          refreshHistory();
          // Get the suggested state
        })
        .catch((error) => {
          console.error(error);
        });

      // get the suggested state
    } catch (error) {}
  },
  {
    immediate: true,
  }
);
</script>

<style scoped>
#MainSpecimenForm {
  max-width: 800px;
}
</style>
