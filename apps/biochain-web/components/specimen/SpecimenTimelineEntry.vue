<script lang="ts" setup>
// import { ccbio } from 'saacs'
const props = defineProps({
  canHide: {
    default: false,
    type: Boolean,
  },
})

// const { data, pending, error, refresh } = await useGetSpecimen();

const entry = defineModel<ccbio.SpecimenHistoryEntry>('entry', {
  default: () => new ccbio.SpecimenHistoryEntry(),
})

const isLast = defineModel('isLast', {
  default: () => false,
  type: Boolean,
})
const hidden = defineModel('hidden', {
  default: () => false,
  type: Boolean,
})

async function sendHide() {
  console.log('sendHide', entry.value)
}

async function sendUnHide() {
  console.log('sendUnHide', entry.value)
}
</script>

<template>
  <QTimelineEntry
    :title="`Transaction ID: ${entry.txId}`"
    icon="ti-pin"
  >
    <!-- <template #title> {{ entry.txId || "Initial State" }} </template> -->
    <template #subtitle>
      {{ entry.timestamp?.toDate() }}
    </template>

    <q-expansion-item
      class="max-w-2xl"
      expand-separator
      icon="ti-bookmark"
      label="Transaction Value"
    >
      <div v-if="!entry.isHidden">
        <SpecimenForm
          :enable-edit="false"
          :specimen="entry?.value"
          :start-open="false"
        />
      </div>
    </q-expansion-item>
    <q-card-actions
      v-if="!isLast && props.canHide"
      class="col justify-center"
    >
      <q-btn
        v-if="hidden"
        color="positive"
        @click="sendUnHide"
      >
        Unhide Transaction
      </q-btn>
      <q-btn
        v-if="!hidden"
        color="negative"
        @click="sendHide"
      >
        Hide Transaction
      </q-btn>
    </q-card-actions>
  </QTimelineEntry>
</template>

<style></style>
