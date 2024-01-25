<script lang="ts" setup>

</script>

<template>
  <div>
    <q-card-section>
      <q-table
        v-model:selected="RowsSelected"
        :rows="rawData"
        dense
        row-key="index"
        selection="multiple"
      >
        <template #body-cell-exist="props">
          <QTd :props="props">
            <UBadge
              :color="existToChipColor(RowMeta[props.rowIndex].exist)"
              :label="RowMeta[props.rowIndex].exist"
            />
          </QTd>
        </template>
        <template #body-cell-status="props">
          <q-td :props="props">
            <UPopover
              :popper="{ adaptive: true }"
              mode="hover"
            >
              <UBadge
                :color="statusToChipColor(RowMeta[props.rowIndex].status)"
                :label="RowMeta[props.rowIndex].status"
              />
              <q-circular-progress
                v-if="RowMeta[props.rowIndex].status == 'loading'"
                color="warn"
                indeterminate
                rounded
                size="15px"
              />
              <template
                v-if="RowMeta[props.rowIndex].statusMessage != ''"
                #panel
              >
                <div class="p-4">
                  <pre wrap>
                  {{ RowMeta[props.rowIndex].statusMessage }}</pre>
                </div>
              </template>
            </UPopover>
          </q-td>
        </template>
      </q-table>
    </q-card-section>
  </div>
</template>

<style scoped>

</style>
