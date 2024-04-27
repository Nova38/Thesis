<script lang="ts" setup>
import { toPlainMessage } from '@bufbuild/protobuf'
import type { FormKitNode } from '@formkit/core'
import { reset } from '@formkit/core'
import { useQueryClient, useQuery, useMutation } from '@tanstack/vue-query'

const toast = useToast()
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
const mode = ref<FormMode>('view')
const headerColor = computed(() => toModeColor(mode.value ?? 'view'))
const FormDisabled = computed(() => mode.value === 'view')

const { status, data, error } = useQuery({
  queryKey: [props.collectionId],
  queryFn: async () => {
    const raw = await $fetch('/api/cc/specimens/get', {
      query: {
        collectionId: props.collectionId,
        specimenId: props.specimenId,
      },
    })

    const data = new pb.Specimen()

    return raw
  },
})

async function submitHandler(value: PlainSpecimen, node: FormKitNode) {
  const x = new pb.Specimen(specimen.value)
  console.log(x)
  console.log({ value, node })
  const z = JSON.stringify(value)
  const v = new pb.Specimen({ ...value })
  const s = v.toJsonString()
  console.log({ v, s, x, z })
  try {
    toast.add({
      title: 'Success',
      description: JSON.stringify(value),
    })

    node.reset(res)
  } catch (e) {
    console.error(e)

    toast.add({
      title: 'Error',
      description: `Failed to update specimen: ${e}`,
    })
  }
}

const specimen = ref<PlainSpecimen>(MakeEmptySpecimen())
</script>

<template>
  <div>
    <FormKit
      :id="FormId"
      v-slot="{ value, node }"
      v-model="specimen"
      type="form"
      :disabled="FormDisabled"
      :actions="false"
      :value="data"
      :plugins="[DirtyLabelPlugin]"
      validation-visibility="live"
      @submit="submitHandler"
    >
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
                v-if="specimen?.primary?.catalogNumber"
                class="flex-grow"
                color="red"
                variant="solid"
                :ui="{
                  rounded: '',
                }"
                size="md"
                :label="`Catalog Number: ${specimen?.primary?.catalogNumber}`"
              />
            </div>
          </div>
        </template>

        <div class="flex flex-col gap-2">
          <FormKit
            name="specimenId"
            type="hidden"
            value=""
          />
          <FormKit
            name="collectionId"
            type="hidden"
            value=""
          />
          <SpecimenFormTaxon class="py-1" />
          <SpecimenFormPrimary />
          <SpecimenFormGeoreference />

          <SpecimenFormSecondary />
          <SpecimenFormLoans />
          <SpecimenFormGrants />
        </div>

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
    </FormKit>
  </div>
</template>

<style scoped>
input[data-mode='view'] {
  background: #000;
}
</style>
