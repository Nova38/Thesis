<script lang="ts" setup>
// import { titleCase } from 'scule'
// import { type FormKitNode } from '@formkit/core'
// import { useFormKitContext } from '@formkit/vue'
import { DirtyLabelPlugin } from '#imports'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

// type Sections = "taxon" | "primary" | "secondary" | "georeference" | "loans" | "grants"
</script>

<template>
  <FormKit
    v-model="specimen"
    name="specimen"
    type="group"
    dirty-behavior="compare"
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

      <!-- <DevOnly>
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
      </DevOnly> -->
    </div>
  </FormKit>
</template>

<style scoped></style>
