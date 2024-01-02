<template>
  <q-timeline-entry>
    <template v-slot:title> {{ tx_date }} </template>
    <template v-slot:subtitle
      >Transaction ID: {{ props.tx_id || 'Initial State' }}
    </template>

    <q-expansion-item expand-separator icon="perm_identity" label="State Value">
      <SpecimenForm
        :specimen="props.specimen"
        style="max-width: 700px"
        :enable-edit="false"
        v-if="props.specimen && !props.hiddenInfo"
      />
      <div v-else-if="props.hiddenInfo">
        Hidden because: {{ props.hiddenInfo.notes }}
      </div>
    </q-expansion-item>

    <q-card-actions class="col justify-center" v-if="!props.isLast">
      <q-btn v-if="hidden" color="positive" @click="sendUnHide">
        Unhide Transaction
      </q-btn>
      <q-btn v-if="!hidden" color="negative" @click="sendHide">
        Hide Transaction
      </q-btn>
    </q-card-actions>
  </q-timeline-entry>
</template>

<script setup lang="ts">
import { Timestamp } from '@bufbuild/protobuf';
import SpecimenForm from 'src/components/specimen/SpecimenForm.vue';

import { schema } from 'src/lib/ccbio';
import { FormSpecimen } from 'src/lib/ccbio/utils/specimen';

// useDateFormat();

const tx_date = computed(() => {
  const date = props.timestamp.toDate();
  return date.toLocaleString();
});

const emit = defineEmits<{
  (e: 'hideTx', id: string): void;
  (e: 'unHideTx', value: string): void;
}>();

const props = defineProps<{
  tx_id: string;
  timestamp: Timestamp;
  is_deleted: boolean;
  specimen: FormSpecimen | null;
  hiddenInfo?: schema.Specimen_HiddenTx;
  isLast: boolean;
}>();

const hidden = computed(() => {
  return !props.specimen;
});

function sendHide() {
  emit('hideTx', props.tx_id);
}

function sendUnHide() {
  emit('unHideTx', props.tx_id);
}
</script>
