<template>
  <q-page class="full-height">
    <q-card class="">
      <q-card-section>
        <h2>Select CSV file to import from</h2>
        <q-file outlined v-model="store.file" accept="csv">
          <template v-slot:prepend>
            <q-icon name="attach_file" />
          </template>
        </q-file>
      </q-card-section>

      <q-card-section v-if="store.headers && store.visibleColumns">
        <div class="q-pa-md">
          <q-table
            v-if="store.data"
            :rows="store.data"
            row-key="index"
            selection="multiple"
            virtual-scroll
            v-model:selected="store.selected"
            :visible-columns="store.visibleColumns"
          >
            <template v-slot:body-cell-status="props">
              <q-td :props="props">
                <div>
                  <!-- <q-badge color="green" :label="props.value"></q-badge> -->
                  <q-badge
                    color="purple"
                    outline
                    v-if="props.value == 'new'"
                    :label="props.value"
                  />
                  <q-badge
                    color="teal"
                    outline
                    v-if="props.value == 'uploading'"
                    :label="props.value"
                  />
                  <q-badge
                    color="red"
                    outline
                    v-if="props.value == 'error'"
                    :label="props.value"
                  >
                    <q-tooltip>
                      {{ props.row.statusMessage }}
                    </q-tooltip>
                  </q-badge>
                  <q-badge
                    color="green"
                    outline
                    v-if="props.value == 'success'"
                    :label="props.value"
                  />
                </div>
                <div class="my-table-details">
                  {{ props.row.details }}
                </div>
              </q-td>
            </template>
          </q-table>
        </div>
      </q-card-section>

      <q-card-section>
        <div class="q-pa-md" v-if="store.mappedArray">
          <q-table :rows="store.mappedArray" :columns="store.mappedColl">
            <template v-slot:header-cell="props">
              <q-th :props="props">
                <div class="">
                  {{ props.col.label }}
                </div>
                <q-select
                  v-model="store.specimenMapping[props.col.name]"
                  label-color="teal-10"
                  dense
                  :options="store.headers"
                >
                  <template
                    v-if="store.specimenMapping[props.col.name]"
                    v-slot:append
                  >
                    <q-icon
                      name="cancel"
                      dense
                      size=".75em"
                      color="red"
                      @click.stop.prevent="store.clearKey(props.col.name)"
                      class="cursor-pointer"
                    />
                  </template>
                </q-select>
              </q-th>
            </template>
          </q-table>
          <q-btn
            color="white"
            text-color="black"
            label="Upload Selected"
            @click="store.upload()"
          />
        </div>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
// types, interfaces, and classes
// import { SpecimenKeys } from 'src/lib/specimen';

import { useImportStore } from 'stores/import';

const store = useImportStore();

// on mount set the file to the file in the store

// Components
// import SpecimenForm from 'src/components/specimen/SpecimenForm.vue';

// vue

// papaparse

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │ Load the file on change and parse it with papaparse                     │
  └─────────────────────────────────────────────────────────────────────────┘
 */
// const file = ref<File | null>(null);
// const headers = ref<string[]>([]);

// watch(file);

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │    Map the headers to the correct specimen fields                       │
  └─────────────────────────────────────────────────────────────────────────┘
 */

// make a list of all the keys in the Specimen class
</script>
