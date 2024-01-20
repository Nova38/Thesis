<script lang="ts" setup>
import { ccbio } from 'saacs-es'

const props = defineProps({
  canHide: {
    type: Boolean,
    default: false,
  },
})

const history = defineModel('history', {
  default: () => new ccbio.SpecimenHistory(),
})
console.log('history', history)
</script>

<template>
  <div>
    <QBar class="flex flex-row text-lg items-center justify-center">
      Specimen History
    </QBar>
    <QCard class="p-4 flex flex-col justify-center">
      <QTimeline
        v-if="history"
        dense
      >
        <template
          v-for="tx in history.entries"
          :key="tx.txId"
        >
          <SpecimenTimelineEntry
            :entry="tx"
            :can-hide="props.canHide"
          />
        </template>
      </QTimeline>
    </QCard>
  </div>
</template>

<style></style>
