<script lang="ts" setup></script>

<template>
  <div>
    <PFieldset
      legend="Primary"
      :toggleable="true"
    >
      <FormKit
        type="group"
        id="primary"
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
          />
          <FormKit
            type="text"
            name="cataloger"
            label="Cataloger"
          />

          <SpecimenFormDate name="catalogDate" />
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

          <SpecimenFormDate name="fieldDate" />
        </div>
        <UDivider />

        <div class="subGroup">
          <FormKit
            type="text"
            name="determiner"
            id="determiner"
            label="Determiner"
          />
          <FormKit
            type="text"
            name="cataloger"
            label="Cataloger"
          />

          <SpecimenFormDate name="determinedDate" />
          <FormKit
            type="textarea"
            name="determinedReason"
            label="Determined Reason"
          />
        </div>
        <SpecimenFormDate name="originalDate" />

        <div class="inline-flex flex-wrap gap-2">
          <div class="inline-flex flex-wrap gap-2"></div>
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped>
.subGroup {
  @apply flex flex-wrap gap-2 py-2;
}
</style>
