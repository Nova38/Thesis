<script setup lang="ts">
/* __placeholder__ */
import type { ICellRendererParams } from '@ag-grid-community/core'

// import { defineProps, defineEmits } from 'vue';
// /**
//  * z
//  */

const props = defineProps<{
  params: ICellRendererParams
}>()

// const props = defineProps<{
//   params: {
//     value: any>;
//     clicked: (value: string) => void;
//   };
// }>();

// console.log(props);

function btnClickedHandler() {
  // alert(`open modal:${props.params.value}}`);

  console.log('hi')
  console.log(props.params.data)

  useRouter().push({
    name: 'collection-specimen',
    params: {
      collectionId: props.params.data.collection_id,
      specimenId: props.params.data.id,
    },
  })

  // const collectionId = currentSpecimen.collection_id;
  // const specimenId = props.params.value;
  //
  // router.push(
  //   `/collection/${router.currentRoute.value.params?.collectionId}/specimen/${currentSpecimen.id?.id}`,
  // );
}
</script>

<template>
  <div class="">
    <NuxtLink
      :to="`/collection/${useRoute().params?.collectionId.toString()}/Specimen/View-${
        props.params.value.data.id
      }`"
      :no-prefetch="true"
    />
  </div>
  <!-- <q-btn @click="btnClickedHandler(param)">
    <q-icon name="edit" />
  </q-btn> -->
  <!-- <q-btn @click="btnClickedHandler">
    <q-icon name="edit" />
  </q-btn> -->
</template>
