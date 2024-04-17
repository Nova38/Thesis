<script lang="ts" setup>
import { crush, keys } from 'radash'

const bulk = useBulkStore()

const MappingOptions = ref<
  {
    group: string
    options: { label: string; value: string }[]
  }[]
>([
  {
    group: 'Empty',
    options: [{ label: ' ', value: ' ' }],
  },
])

const options = FlattedSpecimenKeys.filter(
  (x: string) => !(x === 'collectionId' || x === 'specimenId'),
)
  .map((x: string) => {
    return {
      group: x.split('.')[0],
      options: [{ label: x, value: x }],
    }
  })
  .reduce(
    (
      acc: { group: string; options: { label: string; value: string }[] }[],
      x: { group: string; options: { label: string; value: string }[] },
    ) => {
      const group = acc.find((y) => y.group === x.group)
      if (group) group.options.push(...x.options)
      else acc.push(x)

      return acc
    },
    MappingOptions.value,
  )
  .map((x) => {
    x.options.sort((a, b) => a.label.localeCompare(b.label))
    return x
  })

const rows = ref()

onMounted(() => {
  console.log('bulk.RawColDefs', bulk.RawColDefs)
  console.log('options', options)

  rows.value = bulk.RawRows
})

// MappingOptions.value = [''].concat(FlattedSpecimenKeys)
</script>

<!-- data-key="id"
      paginator
      show-gridlines
      striped-rows
      size="small"
      :rows="10"
      :rows-per-page-options="[5, 10, 20, 50]" -->
<template>
  <div>
    <PDataTable
      :value="rows"
      data-key="id"
      show-gridlines
      striped-rows
      size="small"
      scrollable
      scrollHeight="800px"
      :virtualScrollerOptions="{ itemSize: 46 }"
    >
      <PColumnGroup type="header">
        <PRow>
          <PColumn
            header="Meta"
            :colspan="2"
          />
          <PColumn
            header="Raw Fields"
            :colspan="bulk.RawColDefs.length"
          />
        </PRow>

        <PRow>
          <PColumn
            header="Existence"
            field="id"
          />
          <PColumn
            header="Status"
            field="id"
          />
          <PColumn
            v-for="col of bulk.RawColDefs"
            :key="col.name"
            :field="col.field"
          >
            <template #header="">
              <div class="text-nowrap">
                <FormKit
                  type="select"
                  :label="col.name"
                  :name="`map-${col.name}`"
                  :options="options"
                  label-class="text-nowrap "
                  @input="(val) => bulk.SetMapping(val as string, col)"
                />
                <!-- @input="(val) =>  " -->
                <!--  -->
              </div>
            </template>
          </PColumn>
        </PRow>
      </PColumnGroup>

      <PColumn
        header="Existence"
        field="id"
      >
        <template #body="row">
          <!-- <UBadge class="text-nowrap" :label="bulk.RawRowsMeta.get(row.data?.id)?.exist" /> -->
          <ImportMetaExistence :row="row.data?.id" />
        </template>
      </PColumn>
      <PColumn
        header="Status"
        field="id"
      >
        <template #body="row">
          <ImportMetaStatus :row="row.data?.id" />
        </template>
      </PColumn>
      <PColumn
        v-for="col of bulk.RawColDefs"
        :key="col.name"
        :field="col.field"
      />
    </PDataTable>
  </div>
</template>

<style scoped></style>
