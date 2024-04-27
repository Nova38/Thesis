<script lang="ts" setup>
import { type PlainSpecimen, ccbio, pb } from '#imports'
import { Timestamp, createRegistry, toPlainMessage } from '@bufbuild/protobuf'
import { keys } from 'radash'
import { useFormKitNodeById } from '@formkit/vue'
import SpecimenEditCard from '~/components/specimen/SpecimenEditCard.vue'
import type { AsyncDataRequestStatus } from '#app/composables/asyncData'

import type { PlainMessage } from '@bufbuild/protobuf'

const route = useRoute()

const history = ref<PlainMessage<pb.SpecimenHistory>>(
  new ccbio.SpecimenHistory(),
)

const specimenId = computed(() => route.params?.specimenId.toString() ?? '')
const collectionId = computed(() => route.params?.collectionId.toString() ?? '')

const modeCapitalized = computed(
  () => mode.value[0].toUpperCase() + mode.value.slice(1),
)

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

const mode: Ref<FormMode> = ref('view')

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
const headerColor = computed(() => toModeColor(mode.value))

const FormDisabled = computed(() => mode.value === 'view')
const modeOptions = ref([
  {
    label: 'View',
    value: 'view' as FormMode,
    attrs: {
      'data-mode': 'view',
    },
  },
  {
    label: 'Update',
    value: 'update' as FormMode,
    attrs: {
      mode: 'update',
    },
  },
  {
    label: 'Suggest',
    value: 'suggest' as FormMode,
    attrs: {
      mode: 'suggest',
    },
  },
])
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
      <UCard>
        <template #header>
          <div class="flex flex-row items-center justify-center text-lg">
            History
          </div>
        </template>
        <SpecimenTimeline
          v-if="history"
          :history="history"
          class="basis-size-1/4"
        />
      </UCard>
      <UCard>
        <template #header>
          <div class="flex flex-row items-center justify-center text-lg">
            Suggestions
          </div>
        </template>
      </UCard>
    </div>
  </div>
</template>

<style></style>

<!-- <FormKit
                v-model="
                type="radio"
                :options="modeOptions"
                :classes="{
                  wrapper: 'group/wrapper',
                  options: 'grid grid-cols-3 items-stretch ',
                  option:
                    '$reset group/option  relative border formkit-checked:border-none    border-none  text-center dark:bg-gray-900 ',
                  outer: '$reset group dark:bg-gray-900 w-full grow px-0 py-0',
                  decorator:
                    '$reset absolute top-0 left-0 right-0 bottom-0  group-data-[checked=true]/wrapper:bg-gray-700/25 dark:group-data-[checked=true]/wrapper:bg-gray-100/50',
                  decoratorIcon: '$reset hidden',
                  label: '!text-md',
                  help: '$reset hidden',
                }"
              />-->
