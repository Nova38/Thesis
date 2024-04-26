<script lang="ts" setup>

const collapsed = defineModel<boolean>('collapsed', {
  default: false,
})

</script>

<template>
  <div>
    <PFieldset
      legend="Taxon"
      toggleable
      :collapsed
    >
      <FormKit
        id="Taxon"
        type="group"
        name="taxon"
      >
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            id="kingdom"
            type="text"
            label="Kingdom"
            name="kingdom"
          />
          <FormKit
            id="phylum"
            type="text"
            name="phylum"
            label="Phylum"
          />
          <FormKit
            id="class"
            type="text"
            name="class"
            label="Class"
          />
          <FormKit
            id="order"
            type="text"
            name="order"
            label="Order"
          />
          <FormKit
            id="family"
            type="text"
            name="family"
            label="Family"
          />
          <FormKit
            id="genus"
            type="text"
            name="genus"
            label="Genus"
          />
          <FormKit
            id="species"
            type="text"
            name="species"
            label="Species"
          />
          <FormKit
            id="subspecies"
            type="text"
            name="subspecies"
            label="SubSpecies"
          />
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
