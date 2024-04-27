<script lang="ts" setup>
import { toPlainMessage } from '@bufbuild/protobuf'
import type { FormKitNode } from '@formkit/core'
import { reset } from '@formkit/core'
import { useQueryClient, useQuery, useMutation } from '@tanstack/vue-query'

import { createZodPlugin } from '@formkit/zod'
// const zodSchema = z.object({
const toast = useToast()

const mode = ref<FormMode>('view')
const headerColor = computed(() => toModeColor(mode.value ?? 'view'))
const FormDisabled = computed(() => mode.value === 'view')

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

const { _status, data, _error } = useQuery({
  queryKey: [props.collectionId],
  queryFn: async () => {
    const raw = await $fetch('/api/cc/specimens/get', {
      query: {
        collectionId: props.collectionId,
        specimenId: props.specimenId,
      },
    })
    reset(FormId.value, { specimen: raw })
    return raw
  },
})

const mutation = useMutation({
  mutationFn: async (specimen: PlainSpecimen) => {
    const x = new pb.Specimen(specimen)
    console.log(x.toJsonString({ emitDefaultValues: true }))
    return specimen
  },
  onMutate: async (formData: PlainSpecimen) => {
    const x = new pb.Specimen(formData)
    console.log(x.toJsonString({ emitDefaultValues: true }))
    return formData
  },
  onSuccess: async (formData) => {
    toast.add({
      title: 'Success',
      description: JSON.stringify(formData),
    })
    console.log(formData)
  },
})

onMounted(() => {
  reset(FormId.value, data)
})

const [zodPlugin, submitHandler] = createZodPlugin(
  ZSpecimen,
  async (formData) => {
    mutation.mutate(formData)
  },
)
const specimen = ref<PlainSpecimen>(MakeEmptySpecimen())
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
          <div>
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
          </div>
          <UDivider />

          <div class="flex flex-row p-4">
            {{ specimen?.taxon?.genus }}
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
        <SpecimenForm />
      </FormKit>

      <template #footer>
        <FormKit
          type="submit"
          :outer-class="{
            'max-w-[22em]': false,
            'w-full': true,
          }"
          :disabled="FormDisabled"
          input-class="w-full justify-center items-center flex"
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
