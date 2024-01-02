<template>
  <div class="q-pa-md q-gutter-sm">
    <q-expansion-item
      expand-separator
      dense
      icon="ti-agenda"
      :label="data"
      caption="Collection"
    >
      <q-card>
        <q-card-section>
          <!-- <RouterLink>View</RouterLink> -->
          <!-- TODO: Add Dashboard Page -->
          <!-- <q-item clickable v-ripple :to="'/collection/' + data">
            <q-item-section avatar>
              <q-icon color="primary" name="ti-dashboard" />
            </q-item-section>

            <q-item-section>Dashboard</q-item-section>
          </q-item> -->

          <q-item clickable v-ripple :to="'/collection/' + data + '/specimens'">
            <q-item-section avatar>
              <q-icon color="primary" name="ti-view-list-alt" />
            </q-item-section>

            <q-item-section>Specimen Table</q-item-section>
          </q-item>
          <q-item clickable v-ripple :to="'/collection/' + data">
            <q-item-section avatar>
              <q-icon color="primary" name="ti-dashboard" />
            </q-item-section>

            <q-item-section>Access Control</q-item-section>
          </q-item>
          <q-separator />
          <q-item
            clickable
            v-ripple
            :to="'/collection/' + data + '/new/specimen/'"
            class=""
          >
            <q-item-section avatar>
              <q-icon color="primary" name="ti-plus" />
            </q-item-section>

            <q-item-section>New Specimen</q-item-section>
          </q-item>

          <q-item
            clickable
            v-ripple
            :to="'/collection/' + data + '/bulk_import'"
          >
            <q-item-section avatar>
              <q-icon color="primary" name="ti-import" />
            </q-item-section>

            <q-item-section>Bulk Import</q-item-section>
          </q-item>
        </q-card-section>
      </q-card>
    </q-expansion-item>
  </div>
</template>

<script setup lang="ts">
import { useVModel } from '@vueuse/core';
import { computed } from 'vue';

const baseURL = computed(() => {
  return `/collection/${data}`;
});
const props = defineProps<{
  modelValue: string;
}>();
const emit = defineEmits(['update:modelValue']);

const data = useVModel(props, 'modelValue', emit);
</script>
