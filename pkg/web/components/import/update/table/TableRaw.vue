<script lang="ts" setup>
const bulk = useBulkUpdate()
</script>

<template>
  <div>
    <UCard />
    <QCardSection>
      <QTable
        v-model:selected="bulk.RowsSelected"
        :columns="bulk.ImportColumns"
        :rows="bulk.RawRows"
        dense
        row-key="id"
        selection="multiple"
        class="max-h-60"
      >
        <template #body-cell-exist="props">
          <QTd :props="props">
            <UBadge :label="props.row.meta.exist" />
          </QTd>
        </template>
        <template #body-cell-status="props">
          <q-td :props="props">
            <UPopover
              :popper="{ adaptive: true }"
              mode="hover"
            >
              <UBadge :label="props.row.status" />
              <q-circular-progress
                v-if="props.row.meta.status === 'loading'"
                color="warn"
                indeterminate
                rounded
                size="15px"
              />
              <template
                v-if="props.row.meta.statusMessage !== ''"
                #panel
              >
                <div class="p-4">
                  <pre wrap>
                  {{ props.row.meta.statusMessage }}</pre>
                </div>
              </template>
            </UPopover>
          </q-td>
        </template>
      </QTable>
    </QCardSection>
  </div>
</template>

<style scoped>
thead tr th {
  position: sticky;
}
</style>
