<script lang="ts" setup>
import { titleCase } from 'scule'
import { type FormKitNode } from '@formkit/core'
import { useFormKitContext } from '@formkit/vue'
import { DirtyLabelPlugin } from '#imports'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

const props = withDefaults(
  defineProps<{
    specimen: PlainSpecimen
    formPrefix: string
  }>(),
  {
    specimen: () => MakeEmptySpecimen(),
    formPrefix: 'specimen',
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
  >
    <FormKit
      name="specimenId"
      type="hidden"
      value=""
    />

    <SpecimenFormTaxon />
    <SpecimenFormPrimary />
    <SpecimenFormGeoreference />

    <SpecimenFormSecondary />
    <SpecimenFormLoans />
    <DevOnly>
      <PFieldset
        legend="Value"
        :toggleable="true"
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
  </FormKit>
</template>

<style scoped></style>
