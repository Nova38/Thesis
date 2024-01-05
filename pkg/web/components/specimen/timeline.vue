<template>
  <div>
    <QCard class="p-4">
      TimeLine
      <q-timeline v-if="history" dense color="secondary">
        <template v-for="tx in history.entries" :key="tx.txId">
          <SpecimenTimelineEntry :entry="tx" />
        </template>
      </q-timeline>
    </QCard>
  </div>
</template>

<script lang="ts" setup>
import { auth, ccbio } from "saacs-es";

const history = defineModel("history", {
  default: () => new ccbio.SpecimenHistory(),
});
console.log("history", history);
</script>

<style></style>
