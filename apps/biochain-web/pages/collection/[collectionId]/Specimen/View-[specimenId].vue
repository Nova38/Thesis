<script lang="ts" setup>
import { ccbio, pb } from '#imports'

import type { PlainMessage } from '@bufbuild/protobuf'

const route = useRoute()

const history = ref<PlainMessage<pb.SpecimenHistory>>(
  new ccbio.SpecimenHistory(),
)

const specimenId = computed(() => route.params?.specimenId.toString() ?? '')
const collectionId = computed(() => route.params?.collectionId.toString() ?? '')

// const specimen =  await $fetch(`/api/cc/specimens/get`)

const getHistory = await useCustomFetch<ccbio.Specimen>(
  `/api/cc/specimens/history`,
  {
    key: makeSpecimenHistoryKey(
      route.params?.collectionId.toString(),
      route.params?.specimenId.toString(),
    ),

    onResponse: async ({ response }) => {
      console.log('history', response._data)

      history.value = new pb.SpecimenHistory(response._data)
    },
    query: {
      collectionId: toValue(route.params?.collectionId.toString()),
      specimenId: toValue(route.params?.specimenId.toString()),
    },
  },
)

// const specimen = ref<PlainSpecimen>(dirty.value)

// const nodeRef = useFormKitNodeById('Main-form', (node) => {
//   // perform an effect when the node is available
//   node.on('settled.deep', (v) => {
//     console.log('settled.deep', v)
//     specimen.value = v.payload
//   })
// })

// const history = await useGetSpecimenHistory();

// async function submitHandler(_value: {
//   specimen: PlainSpecimen
//   mode: FormMode
// }) {
//   const oldMode = mode.value
//   try {
//     mode.value = 'view'
//     // const specimen = new ccbio.Specimen(getCurrent.data.value ?? {});

//     const {differences, mask} = diffCrush(toValue(cur), toValue(dirty), [
//       'lastModified',
//     ])
//     console.log({differences})

//     const req = new ccbio.SpecimenUpdate({
//       mask,
//       specimen: dirty.value,
//     })

//     const response = await useCustomFetch(`/api/cc/specimens/update`, {
//       body: req.toJsonString({
//         typeRegistry: createRegistry(ccbio.Specimen, Timestamp),
//       }),
//       method: 'POST',
//     })
//     console.log('response', toValue(response.data))
//     // dirty.value.fromJson(response.data);
//     getHistory.refresh()
//     getCurrent.refresh()
//     console.log('current', toValue(getCurrent.data))
//   } catch (error) {
//     mode.value = oldMode
//     console.error(error)
//   }
// }
</script>

<template>
  <div class="flex flex-row gap-4 p-4">
    <div class="basis-size-3/4 min-w-lg">
      <SpecimenCardEdit
        :specimen-id
        :collection-id
      />
    </div>

    <div class="flex flex-col gap-4">
      <SpecimenCardTimeline
        :collection-id
        :specimenId
      />
      <!-- <UCard>
        <template #header>
          <div class="flex flex-row items-center justify-center text-lg">
            Suggestions
          </div>
        </template>
      </UCard> -->
    </div>
  </div>
</template>

<style></style>
