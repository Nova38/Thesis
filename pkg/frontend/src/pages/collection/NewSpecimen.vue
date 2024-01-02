<template>
  <q-page padding>
    <div class="card">
      <specimen-form :specimen="NewSpecimen" :enable-edit="true">
        <template #actions>
          <q-btn
            label="Submit New Specimen"
            class="full-width"
            @click="submitNewSpecimen"
          />
        </template>
      </specimen-form>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import SpecimenForm from 'src/components/specimen/SpecimenForm.vue';
import { schema } from 'src/lib/ccbio';
import { ref } from 'vue';
import { randomUUID } from 'uncrypto';

import { ccApi } from 'src/boot/axios';
import {
  PopulateNestedRequired,
  SchemaToFormSpecimen,
} from 'src/lib/ccbio/utils/specimen';

const props = defineProps<{
  collectionId: string;
}>();

function MakeNewSpecimen() {
  const s = SchemaToFormSpecimen(
    PopulateNestedRequired(
      new schema.Specimen({
        id: new schema.Specimen_Id({
          collectionId: props.collectionId,
          id: randomUUID(),
        }),
      })
    )
  );

  console.log(s);
  return s;
}

const NewSpecimen = ref(MakeNewSpecimen());

watch(
  () => props.collectionId,
  () => {
    NewSpecimen.value = MakeNewSpecimen();
  }
);

function submitNewSpecimen() {
  console.log(NewSpecimen.value);

  // if (cloned.value?.catalog_date) {
  //   cloned.value.catalog_date = unFormatDate(
  //     cloned.value.catalog_date
  //   ).toISOString();
  // }

  // if (cloned.value?.field_date) {
  //   cloned.value.field_date = unFormatDate(
  //     cloned.value.field_date
  //   ).toISOString();
  // }

  // if (cloned.value?.determined_date) {
  //   cloned.value.determined_date = unFormatDate(
  //     cloned.value.determined_date
  //   ).toISOString();
  // }

  const params = new schema.SpecimenCreateRequest({
    id: NewSpecimen.value.id,
    primary: NewSpecimen.value.primary,
    secondary: NewSpecimen.value.secondary,
    taxon: NewSpecimen.value.taxon,
    georeference: NewSpecimen.value.georeference,
    grants: NewSpecimen.value.grants,
    loans: NewSpecimen.value.loans,
    images: NewSpecimen.value.images,
  });

  console.log(params);

  const res = ccApi.specimen.SpecimenCreate(params);

  console.log(res);
}

console.log(NewSpecimen.value);
</script>

<style scoped lang="css">
.card {
  max-width: 850px;
}
</style>
