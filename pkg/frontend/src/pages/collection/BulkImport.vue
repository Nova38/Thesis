<template>
  <q-page class="full-height">
    <q-card class="q-pa-md q-ma-md">
      <q-card-section>
        <h2>Select CSV file to import from</h2>
        <q-file outlined v-model="file" accept=".csv">
          <template v-slot:prepend>
            <q-icon name="attach_file" />
          </template>
        </q-file>
      </q-card-section>

      <q-card-section>
        <q-select label="Data Format"></q-select
      ></q-card-section>

      <q-card-section>
        <q-table
          dense
          :rows="rawData"
          row-key="index"
          selection="multiple"
          v-model:selected="RowsSelected"
        >
          <template v-slot:body-cell-status="props">
            <q-td :props="props">
              <q-chip
                outline
                :color="statusToChipColor(props.row[props.col.field])"
              >
                <span class="q-pr-sm">
                  {{ props.row[props.col.field] }}
                </span>
                <q-circular-progress
                  indeterminate
                  v-if="props.row[props.col.field] == 'loading'"
                  rounded
                  size="15px"
                  color="warn"
                />
              </q-chip>
            </q-td>
          </template>
        </q-table>
      </q-card-section>

      <q-card-section>
        <div v-if="possessedData">
          <q-table dense :rows="possessedData" :columns="MappingHeaders">
            <template v-slot:header-cell="props">
              <q-th :props="props">
                <div class="">
                  {{ props.col.label }}
                </div>
                <q-select
                  v-model="specimenMapping[props.col.label]"
                  label-color="teal-10"
                  label="key"
                  stack-label
                  dense
                  :options="sortedImportHeaders"
                >
                  <template
                    v-if="specimenMapping[props.col.label]"
                    v-slot:append
                  >
                    <q-icon
                      name="cancel"
                      dense
                      size=".75em"
                      color="red"
                      class="cursor-pointer"
                      @click.stop.prevent="clearKey(props.col.label)"
                    />
                  </template>
                </q-select>
                <!-- <div v-if="props.col.label">hi</div> -->
              </q-th>
            </template>
          </q-table>
          <!-- <q-btn
      color="white"
      text-color="black"
      label=""
      @click="store.upload()"
    /> -->
        </div>
      </q-card-section>
      <q-card-section>
        <q-btn
          color="secondary"
          label="Upload Selected"
          class="full-width"
          @click="run"
        />
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup lang="ts">
import { schema } from 'src/lib/ccSchema';
import { get, set } from 'lodash';
import Papa, { ParseResult } from 'papaparse';
import { reactifyObject, useSorted } from '@vueuse/core';
import { MakeEmptySpecimen } from 'src/lib/utils/specimen';
import { BiochainCC } from 'src/api';
import { randomUUID } from 'uncrypto';
import { useMetaStore } from 'src/stores/meta';
import { SpecimenList } from 'src/lib/ccSchema/index.schema';
// import { Specimen } from 'src/lib/ccSchema/index.schema';

const props = defineProps<{
  collectionId: string;
}>();

const cc = new BiochainCC();

const existingSpecimens = ref(SpecimenList.create());

const existingCatalogNumber = computed(() => {
  return existingSpecimens.value.items.map((specimen) => {
    return specimen.primary?.catalog_number || '';
  });
});

watch(
  props,
  async () => {
    console.log('ActiveCollectionId', props.collectionId);

    const payload = schema.CollectionId.fromPartial({
      collection_id: props.collectionId,
    });

    console.log(payload);
    existingSpecimens.value = await cc.specimenGetByCollection(payload);
    console.log(existingSpecimens.value);
  },
  { immediate: true }
);

const DateFields = [
  'primary.field_date',
  'primary.catalog_date',
  'primary.determined_date',
];

// Reactive variables
const file = ref<File | null>(null);

// Table 1
const headers = ref<string[]>([]);
const rawData = ref<Record<string, string>[]>();
const RowsSelected = ref<Record<string, string>[]>([]);
const rowToUUID = ref<Map<number, string>>(new Map());

watch(RowsSelected, () => {
  console.log(RowsSelected.value);
});

// This is a mapping between the specimen keys and the column headers
// <column header>: <specimen key>
const specimenMapping = ref<Record<string, string>>({});

watch(specimenMapping.value, () => {
  console.log('specimenMapping');

  const key = specimenMapping.value['primary.catalog_number'];

  if (key === undefined) {
    return;
  }
  rawData.value?.forEach((item) => {
    console.log(item[key]);
    if (existingCatalogNumber.value.includes(item[key])) {
      item.status = 'pre-existing';
    }
  });
  // for (const item in rawData.value) {
  //   console.log(item);
  //   const value = get(item, key);
  //   console.log(item[key]);
  //   console.log(value);

  //   // if (existingCatalogNumber.value.includes(value)) {
  //   //   item = 'pre-existing';
  // }
});

function clearKey(key: string) {
  console.log(key);
  specimenMapping.value[key] = '';
}

// Table 2
const possessedData = computed(() => {
  if (!RowsSelected.value || !specimenMapping) {
    return;
  }
  const possessedData: schema.Specimen[] = [];
  console.log('run');
  console.log(RowsSelected.value);

  RowsSelected.value?.forEach((data) => {
    const importingSpecimen = MakeEmptySpecimen();

    for (const key in specimenMapping.value) {
      set(importingSpecimen, key, get(data, specimenMapping.value[key]));
    }
    possessedData.push(importingSpecimen);
  });

  return possessedData;
});

const sortedImportHeaders = computed(() => {
  if (!headers.value) {
    return;
  }
  return headers.value.sort();
});

// Set up the mapping objects
const keysForImport = schema.Specimen.fromPartial({
  primary: schema.Specimen_Primary.create({}),
  secondary: schema.Specimen_Secondary.create(),
  taxon: schema.Specimen_Taxon.create(),
  georeference: schema.Specimen_Georeference.create(),
});

const SpecimenKeys: Array<string> = [];

Object.keys(keysForImport).forEach((key) => {
  if (typeof get(keysForImport, key) === 'string') {
    SpecimenKeys.push(key);
  } else if (typeof get(keysForImport, key) === 'object') {
    Object.keys(get(keysForImport, key)).forEach((subKey) => {
      const fullKey = `${key}.${subKey}`;
      SpecimenKeys.push(fullKey);
    });
  }
});

const FilteredSpecimenKeys = SpecimenKeys.filter((key) => {
  switch (key) {
    case 'id':
      return false;
    case 'last_modified_by':
      return false;
    case 'collection_id':
      return false;

    default:
      break;
  }

  return true;
});

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │ Load the file on change and parse it with papaparse                     │
  └─────────────────────────────────────────────────────────────────────────┘
 */
watch(file, (file) => {
  if (!file) {
    return;
  }

  Papa.parse(file, {
    // worker: true,
    header: true,
    complete: (results: ParseResult<Record<string, string>>) => {
      console.log(results);

      rawData.value = results.data.map((value, index) => {
        rowToUUID.value.set(index, randomUUID());

        return reactifyObject(
          Object.assign(
            {
              index: `${index}`,
              status: 'new',
            },
            value
          )
        );
      });
      headers.value = results.meta.fields || [];
      headers.value.unshift('status');

      // for (let i = 0; i < rawData.value.length; i++) {
      //   rawData.value[i].index = `${i}`;
      //   rawData.value[i].status = 'new';
      // }

      // If the data's header row that contains the value of a specimen key map it to the specimen key to start
      for (const key of FilteredSpecimenKeys) {
        if (headers.value.includes(key)) {
          specimenMapping.value[key] = key;
        }
      }
    },
  });
});

/*
  ┌─────────────────────────────────────────────────────────────────────────┐
  │    Map the headers to the correct specimen fields                       │
  └─────────────────────────────────────────────────────────────────────────┘
 */
const makeHeaders = () => {
  const flat: any = [];

  for (const key of FilteredSpecimenKeys) {
    flat.push({
      label: key,
      field: (row: any) => get(row, key),
    });
  }
  return flat;
};

const MappingHeaders = makeHeaders();

function run() {
  console.log('run');
  console.log(rawData.value);

  const cc = new BiochainCC();

  const dataMap = new Map<number, schema.Specimen>();

  // const uploadingArray: Promise[] = [];

  RowsSelected.value?.forEach((data) => {
    if (rawData.value === undefined) {
      return;
    }

    RowsSelected;

    const importingSpecimen = MakeEmptySpecimen();

    // specimenMapping.value.forEach((value, key) => {
    //   set(importingSpecimen, key, get(data, value));
    // });

    for (const key in specimenMapping.value) {
      if (Object.prototype.hasOwnProperty.call(specimenMapping.value, key)) {
        const element = specimenMapping.value[key];

        const value = get(data, element);

        // Convert the dates
        if (DateFields.includes(key)) {
          if (value) {
            set(importingSpecimen, key, new Date(value as string));
            console.log(importingSpecimen);
          }
          continue;
        }

        set(importingSpecimen, key, value);
      }
    }
    if (data.index) {
      importingSpecimen.id = rowToUUID.value.get(parseInt(data.index)) || '';

      dataMap.set(parseInt(data.index), importingSpecimen);
    }

    // DateFields.forEach((field) => {
    // DateFields.forEach((field) => {
    //   if (importingSpecimen[field]) {
    //     importingSpecimen[field] = new Date(importingSpecimen[field] as string);
    //   }
    // });

    console.log(importingSpecimen);

    // possessedData.value.set(data.index, importingSpecimen);
  });

  console.log(dataMap);

  // TODO: Make sure they haven't already been uploaded by checking the catalog number

  dataMap.forEach((value, key) => {
    if (!rawData?.value) {
      return;
    }
    const row = rawData.value.at(key);
    if (!row) {
      return;
    }

    if (row.status === 'success') {
      console.debug('Skipping row as it has already been uploaded', row);
      return;
    }

    const payload = schema.CreateSpecimenRequest.fromPartial({
      specimen_id: schema.SpecimenId.fromPartial({
        collection_id: props.collectionId,
        id: value.id,
      }),
      georeference: value.georeference,
      primary: value.primary,
      secondary: value.secondary,
      taxon: value.taxon,
      grants: value.grants,
      loans: value.loans,
      images: value.images,
    });

    row.status = 'importing';

    cc.specimenCreate(payload)
      .then((res) => {
        console.log(res);
        row.status = 'success';
      })
      .catch((err) => {
        console.log(err);
        row.status = 'error';
      });
  });

  // console.log(possessedData.value);
}
console.log(specimenMapping.value);

type status = 'new' | 'importing' | 'success' | 'error' | 'pre-existing';

function statusToChipColor(status: status) {
  switch (status) {
    case 'new':
      return 'blue';
    case 'pre-existing':
      return 'DeepPink';
    case 'importing':
      return 'orange';
    case 'success':
      return 'green';
    case 'error':
      return 'red';
    default:
      return 'blue';
  }
}
</script>
