<template>
  <div>
    <QTimelineEntry>
      <template #title> {{ entry.timestamp }} </template>
      <template #subtitle
        >Transaction ID: {{ entry.txId || "Initial State" }}
      </template>

      <q-expansion-item
        expand-separator
        icon="perm_identity"
        label="State Value"
      >
        <div v-if="!entry.isHidden">
          <FormKit
            id="SpecimenForm"
            v-slot="{ value }"
            :value="entry.value"
            :actions="false"
            type="form"
            disabled
          >
            <SpecimenForm
              :specimen="value"
              :start-open="false"
              class="max-w-3xl"
            />
          </FormKit>
        </div>
      </q-expansion-item>

      <q-card-actions v-if="!isLast" class="col justify-center">
        <q-btn v-if="hidden" color="positive" @click="sendUnHide">
          Unhide Transaction
        </q-btn>
        <q-btn v-if="!hidden" color="negative" @click="sendHide">
          Hide Transaction
        </q-btn>
      </q-card-actions>
    </QTimelineEntry>
  </div>
</template>

<script lang="ts" setup>
import { auth, ccbio } from "saacs-es";
const { data, pending, error, refresh } = await useGetSpecimen();

const entry = defineModel<ccbio.SpecimenHistoryEntry>("entry", {
  default: () => new ccbio.SpecimenHistoryEntry(),
});

const isLast = defineModel("isLast", {
  default: () => false,
});
const hidden = defineModel("hidden", {
  default: () => false,
});

async function sendHide() {
  console.log("sendHide", entry.value);
}

async function sendUnHide() {
  console.log("sendUnHide", entry.value);
}
</script>

<style></style>
