<script lang="ts" setup>
import { titleCase } from 'scule'
import { FilterMatchMode, FilterOperator } from 'primevue/api'
import type {
  DataTableFilterMeta,
  DataTableOperatorFilterMetaData,
} from 'primevue/datatable'
import { first } from 'radash'

const props = defineProps<{
  specimenList: PlainSpecimen[]
}>()

const ColDefs = ref([
  {
    field: 'specimenId',
    filter: true,
    headerName: 'Specimen ID',
    name: 'SpecimenID',
    pin: 'left',
    sortable: true,
  },
  {
    align: 'left',
    field: 'primary.catalogNumber',
    filter: true,
    headerName: 'Catalog Number',
    sortable: true,
  },
  {
    align: 'left',
    field: 'primary.accessionNumber',
    headerName: 'Accession Number',
    name: 'Accession Number',
    sortable: true,
  },
  {
    align: 'left',
    field: 'primary.fieldNumber',
    headerName: 'Field Number',
    name: 'Field Number',
    sortable: true,
  },
  {
    align: 'left',
    field: 'primary.tissueNumber',
    headerName: 'Tissue Number',
    name: 'tissueNumber',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.cataloger',
    headerName: 'Cataloger',
    name: 'cataloger',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.collector',
    headerName: 'Collector',
    name: 'collector',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.determiner',
    headerName: 'Determiner',
    name: 'determiner',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.fieldDate.verbatim',
    headerName: 'Field Date',
    name: 'Field Date',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.determinedDate.verbatim',
    headerName: 'Determined Date',
    name: 'Determined Date',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.determinedReason',
    headerName: 'Determined Reason',
    name: 'determinedReason',
    sortable: true,
  },
  {
    align: 'left',
    columnGroupShow: 'open',
    field: 'primary.originalDate.verbatim',
    headerName: 'Original Date',
    name: 'originalDate',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.kingdom',
    headerName: 'Kingdom',
    name: 'kingdom',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.phylum',
    headerName: 'Phylum',
    name: 'phylum',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.class',
    headerName: 'Class',
    name: 'class',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.order',
    headerName: 'Order',
    name: 'order',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.family',
    headerName: 'Family',
    name: 'family',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.genus',
    headerName: 'Genus',
    name: 'genus',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.species',
    headerName: 'Species',
    name: 'species',
    sortable: true,
  },
  {
    align: 'left',
    field: 'taxon.subspecies',
    headerName: 'Subspecies',
    name: 'subspecies',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.sex',
    headerName: 'Sex',
    name: 'sex',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.age',
    headerName: 'Age',
    name: 'age',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.weight',
    headerName: 'weight',
    name: 'weight',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.weightUnits',
    headerName: 'Weight Units',
    name: 'weightUnits',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.condition',
    headerName: 'condition',
    name: 'condition',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.molt',
    headerName: 'Molt',
    name: 'molt',
    sortable: true,
  },
  {
    align: 'left',
    field: 'secondary.notes',
    headerName: 'Secondary Notes',
    name: 'secondaryNotes',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.country',
    headerName: 'Country',
    name: 'country',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.stateProvince',
    headerName: 'State Province',
    name: 'stateProvince',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.county',
    headerName: 'County',
    name: 'county',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.locality',
    headerName: 'Locality',
    name: 'locality',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.latitude',
    headerName: 'Latitude',
    name: 'latitude',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.longitude',
    headerName: 'Longitude',
    name: 'longitude',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.habitat',
    headerName: 'Habitat',
    name: 'habitat',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.continent',
    headerName: 'Continent',
    name: 'continent',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.locationRemarks',
    headerName: 'Location Remarks',
    name: 'locationRemarks',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.coordinateUncertaintyInMeters',
    headerName: 'Coordinate Uncertainty In Meters',
    name: 'coordinateUncertaintyInMeters',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.georeferenceBy',
    headerName: 'Georeference By',
    name: 'georeferenceBy',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.georeferenceDate.verbatim',
    headerName: 'Georeference Date',
    name: 'GeoreferenceDate',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.georeferenceProtocol',
    headerName: 'Georeference Protocol',
    name: 'georeferenceProtocol',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.geodeticDatum',
    headerName: 'Geodetic Datum',
    name: 'geodeticDatum',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.footprintWkt',
    headerName: 'Footprint Wkt',
    name: 'footprintWkt',
    sortable: true,
  },
  {
    align: 'left',
    field: 'georeference.notes',
    headerName: 'Notes',
    name: 'notes',
    sortable: true,
  },
])

/**
 * The columns that are currently selected
 */
const SelectedColumns = useState('SelectedColumns', () => {
  return ColDefs.value.slice(0, 5)
})
function onToggle(val: unknown[]) {
  SelectedColumns.value = ColDefs.value.filter((col) => val.includes(col))
}

// Reduce the ColDefs to extract the column groups based on the the first part of the
// field name separated by a dot (.), if the field does not contain a dot, it is placed in the meta group
// Also, keep track of the number of columns in each group to calculate the colspan for the header
const ColGroups = computed(() => {
  // ColDefs.value.map()
  const groups = SelectedColumns.value.reduce(
    (acc, col) => {
      let [group, ...children] = col.field.split('.')
      if (children.length === 0) group = 'Meta'

      if (!acc[group]) {
        acc[group] = {
          headerName: titleCase(group),
          children: 0,
        }
      }

      acc[group].children += 1
      return acc
    },
    {} as Record<string, { headerName: string; children: number }>,
  )

  groups.Meta = {
    headerName: 'Meta',
    children: groups.Meta ? groups.Meta.children + 1 : 1,
  }
  return Object.values(groups)
})

const filterFields = computed(() => {
  return SelectedColumns.value.map((col) => col.field)
})

const filters = useState<DataTableFilterMeta>('filters', () => {
  return {
    global: { value: null, matchMode: FilterMatchMode.CONTAINS },

    specimenId: {
      operator: FilterOperator.AND,
      constraints: [{ value: '', matchMode: FilterMatchMode.STARTS_WITH }],
    },
    'primary.catalogNumber': {
      operator: FilterOperator.AND,
      constraints: [{ value: '', matchMode: FilterMatchMode.STARTS_WITH }],
    },
  }
})

function initFilter() {
  filters.value = filterFields.value.reduce(
    (acc: Record<string, DataTableOperatorFilterMetaData>, field) => {
      acc[field] = {
        operator: FilterOperator.AND,
        constraints: [{ value: '', matchMode: FilterMatchMode.CONTAINS }],
      }
      return acc
    },
    {} as Record<string, DataTableOperatorFilterMetaData>,
  )
}

const pageState = useState<{
  first: number
  sortField: string
  sortOrder: string[]
}>('pageState', () => {
  return { first: 0, sortField: 'catalogNumber', sortOrder: [] }
})
onMounted(() => {
  if (!pageState.value.first) {
    pageState.value.first = 0
  }
})

// initFilter()
// const filterss = ref({
//   'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
//   'name': { value: null, matchMode: FilterMatchMode.STARTS_WITH },
//   'country.name': { value: null, matchMode: FilterMatchMode.STARTS_WITH },
//   'representative': { value: null, matchMode: FilterMatchMode.IN },
//   'status': { value: null, matchMode: FilterMatchMode.EQUALS },
//   'verified': { value: null, matchMode: FilterMatchMode.EQUALS },
// })

// onMounted(() => {
//   filters.value = filterFields.value.map((field) => {
//     return {
//       field,
//       value: null,
//       matchMode: FilterMatchMode.CONTAINS,
//     }
//   })
// })
</script>

<template>
  <div>
    <PDataTable
      size="small"
      removable-sort
      sort-mode="multiple"
      column-resize-mode="expand"
      data-key="specimenId"
      :first="pageState?.first ?? 0"
      :value="props.specimenList"
      :paginator="true"
      :rows="10"
      :rows-per-page-options="[5, 10, 25, 50, 100]"
      :global-filter-fields="filterFields"
      filter-display="menu"
      :filters="filters"
      :pt="{}"
      @page="({ first }) => (pageState.first = first)"
    >
      <template #header>
        <div style="text-align: left">
          <PMultiSelect
            :model-value="SelectedColumns"
            :options="ColDefs"
            class="w-full"
            option-label="headerName"
            display="chip"
            placeholder="Select Columns"
            @update:model-value="onToggle"
          />
          <!-- <USelectMenu
            :model-value="SelectedColumns" :options="ColDefs" option-label="header" display="chip"
            placeholder="Select Columns" @update:model-value="onToggle"

          /> -->
        </div>
      </template>

      <PColumnGroup type="header">
        <PRow>
          <PColumn
            v-for="group in ColGroups"
            :key="group.headerName"
            :header="group.headerName"
            :colspan="group.children"
          />
        </PRow>
        <PRow>
          <PColumn header="Actions" />
          <PColumn
            v-for="col in SelectedColumns"
            :key="col.field"
            :field="col.field"
            :header="col.headerName"
            sortable
            header-class="text-nowrap"
          >
            <template #filter="{ filterModel }">
              <PInputText
                v-model="filterModel.value"
                mode="text"
              />
            </template>
            <template #filterclear="data">
              <UButton
                label="Clear"
                @click="
                  async () => {
                    data.filterCallback()

                    await nextTick()
                    filters[data.field] = data.filterModel
                  }
                "
                severity="secondary"
              />
            </template>
            <template #filterapply="data">
              <UButton
                label="Apply"
                @click="
                  async () => {
                    data.filterCallback()

                    await nextTick()
                    filters[data.field] = data.filterModel
                  }
                "
                severity="success"
              />
            </template>
          </PColumn>
        </PRow>
      </PColumnGroup>
      <PColumn header="Actions">
        <template #body="{ data }">
          <NuxtLink
            :to="`/collection/${$route.params.collectionId}/Specimen/View-${data.specimenId}`"
          >
            View
          </NuxtLink>
        </template>
        <!-- <NuxtLink
          :no-prefetch="true"
          :to="`/collection/${$route.params.collectionId}/Specimen/View-${props.row.specimenId}`"
        /> -->
      </PColumn>

      <PColumn
        v-for="col in SelectedColumns"
        :key="col.field"
        :field="col.field"
      />
    </PDataTable>
  </div>
</template>

<style></style>
