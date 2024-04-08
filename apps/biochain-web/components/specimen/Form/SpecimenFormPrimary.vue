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
      >
        <div class="inline-flex gap-2 flex-wrap">
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
