<script lang="ts" setup>
const bulk = useBulkStore()

const { processing } = storeToRefs(bulk)
</script>

<template>
  <div>
    <UCard
      v-if="bulk.RawRows.length"
      class="m-2"
    >
      <template #header>
        <div class="w-full object-center">
          <PSelectButton
            v-model="bulk.Mode"
            :options="['import', 'update', 'hybrid']"
            aria-labelledby="basic"
          />
        </div>
      </template>
      <div>
        <UMeterGroup
          :min="0"
          :max="bulk.RawRows.length"
          size="md"
          indicator
          icon="i-heroicons-minus"
        >
          <UMeter
            :value="processing.success"
            color="green"
            label="Success"
          />
          <UMeter
            :value="processing.fail"
            color="red"
            label="Errors"
          />
        </UMeterGroup>
      </div>
      <template #footer>
        <UButton
          block
          class="mb-4"
          :disabled="!bulk.RawRows.length"
          @click="bulk.Upload"
        >
          Upload
        </UButton>
      </template>
    </UCard>
  </div>
</template>

<style scoped></style>
