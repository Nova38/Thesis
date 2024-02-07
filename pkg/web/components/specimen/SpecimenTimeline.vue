<script lang="ts" setup>
import { ccbio } from '~/lib'

const props = defineProps({
  canHide: {
    default: false,
    type: Boolean,
  },
})

const history = defineModel('history', {
  default: () => new ccbio.SpecimenHistory(),
  type: ccbio.SpecimenHistory,
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
            :can-hide="props.canHide"
            :entry="tx"
          />
        </template>
      </QTimeline>
    </QCard>
  </div>
</template>

<style></style>
