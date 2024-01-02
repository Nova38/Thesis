<script setup lang="ts">
import { ICellRendererParams } from '@ag-grid-community/core';
import { PlainMessage } from '@bufbuild/protobuf';
import { schema } from 'src/lib/ccbio';
import { Specimen_Id } from 'src/lib/ccbio/index.schema';
import { useRoute, useRouter } from 'vue-router';

// import { defineProps, defineEmits } from 'vue';
// /**
//  * z
//  */
const props = defineProps(['params']);

// const props = defineProps<{
//   params: {
//     value: any>;
//     clicked: (value: string) => void;
//   };
// }>();

const router = useRouter();
// console.log(props);

function btnClickedHandler() {
  // alert(`open modal:${props.params.value}}`);
  const currentSpecimen: PlainMessage<schema.Specimen> = props.params?.data;

  console.log(currentSpecimen);
  console.log(router.currentRoute.value.params?.collectionId);

  // const collectionId = currentSpecimen.collection_id;
  // const specimenId = props.params.value;
  //
  router.push(
    `/collection/${router.currentRoute.value.params?.collectionId}/specimen/${currentSpecimen.id?.id}`
  );
}
console.log();
</script>

<template>
  <!-- <q-btn @click="btnClickedHandler(param)">
    <q-icon name="edit" />
  </q-btn> -->
  <q-btn @click="btnClickedHandler">
    <q-icon name="edit" />
  </q-btn>
</template>
