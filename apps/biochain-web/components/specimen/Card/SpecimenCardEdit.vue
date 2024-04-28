<script lang="ts" setup>
import { toPlainMessage } from '@bufbuild/protobuf'
import type { FormKitNode } from '@formkit/core'
import { reset, submitForm } from '@formkit/core'
import { useQueryClient, useQuery, useMutation } from '@tanstack/vue-query'

import { createZodPlugin } from '@formkit/zod'
import { z } from 'zod'
const zodSchema = z.object({
  specimen: ZSpecimen,
})

const toast = useToast()

const mode = ref<FormMode>('view')
const headerColor = computed(() => toModeColor(mode.value ?? 'view'))
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
const props = withDefaults(
  defineProps<{
    collectionId: string
    specimenId: string
    formPrefix?: string
  }>(),
  {
    formPrefix: 'specimen',
  },
)
const FormId = computed(
  () => `${props.formPrefix}-${props.collectionId}-${props.specimenId}-form`,
)

const specimen = ref<PlainSpecimen>(MakeEmptySpecimen())
const { isSuccess, data } = useQuery({
  queryKey: ['specimen', props.collectionId, props.specimenId],
  queryFn: async () => {
    const raw = await $fetch<PlainSpecimen>('/api/cc/specimens/get', {
      query: {
        collectionId: props.collectionId,
        specimenId: props.specimenId,
      },
    })

    const parsed = ZSpecimen.safeParse(raw)
    if (!parsed.success) {
      toast.add({
        title: 'Specimen Failed to load',
        id: 'specimen-status',
        timeout: 5000,
        icon: 'line-md:alert',
      })
      return
    } else {
      if (parsed.data?.primary) {
        parsed.data.primary.originalDate ??= new pb.Date({})
      }
      if (parsed.data?.georeference) {
        parsed.data.georeference.georeferenceDate ??= new pb.Date({})
      }
      if (parsed.data?.georeference) {
        parsed.data.georeference.georeferenceDate ??= new pb.Date({})
      }

      if (parsed.data?.primary?.originalDate === undefined) {
      }
    }

    toast.add({
      title: 'Specimen loaded',
      id: 'specimen-status',
      timeout: 5000,
      icon: 'line-md:square-to-confirm-square-transition',
    })
    reset(FormId.value, { specimen: parsed.data })
    return parsed.data
  },
})

// <!-- const { mutate } = useMutation({
//   mutationFn: (specimen: PlainSpecimen) =>
//     $fetch('/api/cc/specimens/update', {
//       method: 'POST',
//       body: new pb.Specimen(specimen),
//     }),

//   onMutate: async (formData: PlainSpecimen) => {
//     return formData
//   },
//   onSuccess: async (formData) => {
//     toast.add({
//       title: 'Success',
//       description: JSON.stringify(formData),
//     })
//     console.log(formData)
//     reset(FormId.value, formData)
//   },
// }) -->

onMounted(() => {
  reset(FormId.value, data)
})

const [zodPlugin, submitHandler] = createZodPlugin(
  zodSchema,
  async (formData) => {
    console.log(formData)

    return await $fetch('/api/cc/specimens/update', {
      method: 'POST',
      body: {
        specimen: new pb.Specimen(formData.specimen),
      },
    })
  },
)
// const { _data, refresh } = await useFetch('/api/cc/specimens/get', {
//   query: {
//     collectionId: props.collectionId,
//     specimenId: props.specimenId,
//   },
//   onResponse({ response }) {
//     console.log(response)
//     if (response.ok) {
//       console.log(response._data)
//       // reset(FormId.value, { specimen: ZSpecimen.parse(response._data) })
//       toast.add({
//         title: 'Specimen loaded',
//         id: 'specimen-status',
//         timeout: 5000,
//         icon: 'line-md:square-to-confirm-square-transition',
//       })
//     }
//   },
//   onResponseError({ request, response, options }) {
//     toast.add({
//       title: 'Specimen Failed to load',
//       id: 'specimen-status',
//       timeout: 5000,
//       icon: 'line-md:alert',
//     })
//   },
// })
</script>

<template>
  <div>
    <UCard
      class="min-w-2xl max-w-3xl"
      :ui="{
        header: {
          background: headerColor,
          padding: '',
        },
      }"
    >
      <template #header>
        <div>
          <FormKit
            v-model="mode"
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
          />

          <!-- <div>
            <UButton
              label="View"
              @click="
                () => {
                  mode = 'view'
                }
              "
            />
            <UButton
              label="Update"
              @click="
                () => {
                  mode = 'update'
                }
              "
            />
            <UButton
              label=" mode"
              @click="
                () => {
                  mode = 'suggest'
                }
              "
            />
          </div> -->
          <UDivider />

          <div class="flex flex-row p-4">
            {{ data?.taxon?.genus }}
            {{ specimen?.taxon?.species }}
          </div>
          <div class="flex flex-grow flex-row">
            <UBadge
              class="flex-grow"
              color="purple"
              variant="solid"
              size="md"
              :ui="{
                rounded: '',
              }"
              :label="`Collection: ${specimen?.collectionId}`"
            />
            <UBadge
              v-if="data?.primary?.catalogNumber"
              class="flex-grow"
              color="red"
              variant="solid"
              :ui="{
                rounded: '',
              }"
              size="md"
              :label="`Catalog Number: ${data?.primary?.catalogNumber}`"
            />
          </div>
        </div>
      </template>

      <FormKit
        :id="FormId"
        v-slot="{}"
        type="form"
        :disabled="FormDisabled"
        :actions="false"
        :plugins="[DirtyLabelPlugin, zodPlugin]"
        validation-visibility="live"
        @submit="submitHandler"
      >
        <SpecimenForm v-if="isSuccess" />
      </FormKit>

      <template #footer>
        <FormKit
          type="button"
          :outer-class="{
            'max-w-[22em]': false,
            'w-full': true,
          }"
          :disabled="FormDisabled"
          input-class="w-full justify-center items-center flex"
          @click="submitForm(FormId)"
        >
          Send
        </FormKit>
      </template>
    </UCard>
  </div>
</template>

<style scoped>
input[data-mode='view'] {
  background: #000;
}
</style>
