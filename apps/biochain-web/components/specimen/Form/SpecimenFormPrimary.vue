<script lang="ts" setup>
const collapsed = defineModel<boolean>('collapsed', {
  default: false,
})
</script>

<template>
  <div>
    <PFieldset
      legend="Primary"
      toggleable
      :collapsed
    >
      <FormKit
        id="primary"
        v-slot="{ value }"
        type="group"
        name="primary"
        :config="{
          classes: {
            outer: 'min-w-20',
          },
        }"
      >
        <div class="subGroup">
          <FormKit
            type="text"
            name="tissueNumber"
            label="Tissue Number"
          />
          <FormKit
            type="text"
            name="accessionNumber"
            label="Accession Number"
          />
        </div>

        <div class="subGroup">
          <FormKit
            type="text"
            name="catalogNumber"
            label="Catalog Number"
            validation="required"
          />
          <FormKit
            type="text"
            name="cataloger"
            label="Cataloger"
          />

          <SpecimenFormDate
            v-if="value?.catalogDate"
            name="catalogDate"
          />
        </div>

        <UDivider />

        <div class="subGroup">
          <FormKit
            type="text"
            name="fieldNumber"
            label="Field Number"
          />
          <FormKit
            type="text"
            name="collector"
            label="Collector"
          />

          <SpecimenFormDate
            v-if="value?.fieldDate"
            name="fieldDate"
          />
        </div>
        <UDivider />

        <div class="subGroup">
          <FormKit
            id="determiner"
            type="text"
            name="determiner"
            label="Determiner"
          />
          <FormKit
            type="text"
            name="cataloger"
            label="Cataloger"
          />

          <SpecimenFormDate
            v-if="value?.determinedDate"
            name="determinedDate"
          />
        </div>
        <div class="subGroup">
          <FormKit
            type="textarea"
            auto-height
            name="determinedReason"
            label="Determined Reason"
          />
        </div>
        <SpecimenFormDate
          v-if="value?.originalDate"
          name="originalDate"
        />

        <div class="inline-flex flex-wrap gap-2">
          <div class="inline-flex flex-wrap gap-2"></div>
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped>
.subGroup {
  @apply container flex flex-wrap gap-2 py-2;
}
</style>
