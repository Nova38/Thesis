<script lang="ts" setup></script>

<template>
  <div>
    <PFieldset
      legend="Secondary"
      :toggleable="true"
    >
      <FormKit
        type="group"
        id="secondary"
        name="secondary"
      >
        <div class="inline-flex gap-4">
          <div class="inline-flex flex-row flex-wrap gap-1 justify-evenly">
            <FormKit
              type="text"
              name="tissueNumber"
              id="tissueNumber"
              label="Tissue Number"
              outer-class="min-w-20"
            />
            <FormKit
              type="text"
              name="accessionNumber"
              id="accessionNumber"
              label="Accession Number"
              outer-class="min-w-20"
            />
          </div>
          <div class="inline-flex flex-wrap gap-2">
            <FormKit
              type="text"
              name="cataloger"
              id="cataloger"
              label="Cataloger"
              outer-class="min-w-20"
            />
            <FormKit
              type="text"
              name="catalogNumber"
              id="catalogNumber"
              label="Catalog Number"
              outer-class="min-w-20"
            />
          </div>

          <div class="inline-flex flex-wrap gap-2"></div>
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
