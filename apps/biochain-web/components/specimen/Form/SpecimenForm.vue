<script lang="ts" setup>
// import { titleCase } from 'scule'
// import { type FormKitNode } from '@formkit/core'
// import { useFormKitContext } from '@formkit/vue'
import { DirtyLabelPlugin } from '#imports'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

const props = withDefaults(
  defineProps<{
    formPrefix: string
  }>(),
  {
    formPrefix: 'specimen',
    startCollapsed: false,
  },
)

const FormId = computed(() => `${props.formPrefix}-form`)
</script>

<template>
  <FormKit
    :id="FormId"
    v-slot="{ value }"
    v-model="specimen"
    type="group"
    dirty-behavior="compare"
    :actions="false"
    :plugins="[DirtyLabelPlugin]"
    validation-visibility="live"
    class="space-y-4"
  >
    <div class="flex flex-col gap-2">
      <FormKit
        name="specimenId"
        type="hidden"
        value=""
      />

      <SpecimenFormTaxon class="py-1" />
      <SpecimenFormPrimary />
      <SpecimenFormGeoreference />

      <SpecimenFormSecondary />
      <SpecimenFormLoans />

      <DevOnly>
        <PFieldset
          legend="Value"
          :toggleable="true"
          :collapsed="true"
        >
          <NuxtErrorBoundary>
            <Shiki
              lang="json"
              class="max-w-2xl overflow-x-scroll"
              :code="JSON.stringify(value, null, 2)"
            />
          </NuxtErrorBoundary>
        </PFieldset>
      </DevOnly>
    </div>
  </FormKit>
</template>

<style scoped></style>
