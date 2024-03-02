import { defineStore } from 'pinia'
import { flatten } from 'flat'
import { EmptySpecimenMapping } from '~/utils/objects/Mapping'
import type {
  CSVImportMetadata,
  IdMappedRow,
  PlainSpecimen,
  UpdateRawRow,
} from '~/utils/types'

import { ccbio } from '~/lib'

// CatalogNumber is used to calculate the specimenId uuid to make sure it is unique
// SpecimenId is used to identify the specimen in the database

// Updating the Existing Specimen

export const useBulkUpdate = defineStore('BulkUpdate', () => {
  const CollectionId = ref<string>('')

  const SpecimenIdHeader = ref('primary.specimenId')

  const SpecimenIds = ref<string[]>([])

  // Imported Data
  // ------------------------------

  const ImportedHeaders = ref<string[]>([])

  const ImportColumns = computed(() => {
    // Check if the RawHeaders is empty or if specimenIdHeader is in the RawHeaders and return the columns

    if (
      !ImportedHeaders.value ||
      !ImportedHeaders.value.includes(SpecimenIdHeader.value)
    ) {
      return {
        id: {},
      } as ImportColumns
    }

    const r: ImportCol[] =
      ImportedHeaders.value?.map((header) => {
        return {
          name: header,
          label: header,

          field: (row: UpdateRawRow) => row.raw[header],
          // field: `raw.${header}`,
          colType: header === SpecimenIdHeader.value ? 'id' : 'raw',
        }
      }) || []

    const id = r.find((col: ImportCol) => col.colType === 'id')
    if (!id) throw new Error('SpecimenIdHeader not found in RawHeaders')

    return {
      id,
      raw: r.filter((col: ImportCol) => col.colType === 'raw'),
      meta: r.filter((col: ImportCol) => col.colType === 'meta'),
    }
  })

  const RawRows = shallowRef<UpdateRawRow[]>()

  // Current Data
  // ------------------------------
  const CurrentSpecimen = ref<Map<string, PlainSpecimen>>(new Map())

  //
  // Processing Data
  // ------------------------------

  const SpecimenMapping = ref(EmptySpecimenMapping())

  const MappedRows = computed(() => {
    return (
      RawRows.value?.map((row: UpdateRawRow) => {
        // console.log({ row, SpecimenMapping: SpecimenMapping.value })
        const val = TransformRecordToFlatSpecimen(
          row.raw,
          SpecimenMapping.value,
        )

        console.log({ val })

        return {
          id: row.id,
          val,
        }
      }) || []
    )
  })

  const RowMetadata = ref<Record<string, UpdateRowMeta>>({})

  // ------------------------------
  // Methods
  // ------------------------------

  const fetchSpecimens = async () => {
    console.log('fetching specimens')
    const x = await $fetch('/api/cc/specimens/fullList?collectionId=k')
    console.log(x)

    return x
  }

  const SetMapping = (selected: string | undefined, col: { field: string }) => {
    console.log({ selected, col })

    if (!selected) return

    SpecimenMapping.value = SpecimenMapping.value.map((m) => {
      if (m.newKey === selected) m.oldKey = col.field

      return m
    })
  }

  /**
   * @param csv The rows
   */
  const LoadUpdates = async (csv: CSVImportMetadata) => {
    // Reset The Row variables
    SpecimenIds.value = []
    CurrentSpecimen.value = new Map()

    ImportedHeaders.value = csv.headers
    SpecimenIdHeader.value = csv.specimenIdHeader

    RawRows.value = csv.rows.map((row) => {
      const id = row[SpecimenIdHeader.value]
      if (!id) {
        console.warn(row)
        throw new Error('SpecimenIdHeader not found in RawHeaders')
      }

      const catNum = CatNumToUUID(id)
      SpecimenIds.value.push(catNum)

      return {
        id: catNum,
        raw: row,
      }
    })

    const fullList = await $fetch('/api/cc/specimens/bulk/partialList', {
      method: 'post',
      query: {
        collectionId: CollectionId.value,
      },
      body: {
        specimenIds: SpecimenIds.value,
      },
    })

    console.log(fullList)

    Object.entries(fullList.filteredList).forEach((value) => {
      console.log(value)
      const [id, v] = value[1]

      CurrentSpecimen.value.set(id, new ccbio.Specimen(v))
    })
    console.log(CurrentSpecimen.value)

    // const _selectiveList = await $fetch('/api/cc/specimens/selectiveList', {
    //   body: {
    //     collectionId: CollectionId.value,
    //     specimenIds: SpecimenIds,
    //   },
    // })

    // const list = await $fetch('/api/cc/specimens/fullList', {
    //   body: {
    //     collectionId: CollectionId.value,
    //   },
    // })

    // if (!list) throw new Error('Failed to fetch full list')
  }

  function ClearData() {
    SpecimenIds.value = []
    ImportedHeaders.value = []
    RawRows.value = []
    CurrentSpecimen.value = new Map()
    SpecimenMapping.value = EmptySpecimenMapping()
    RowMetadata.value = {}
  }

  return {
    /**
     * @description CollectionId is the collection that the specimens are being updated
     */
    CollectionId,
    /**
     * @description The header that contains the specimenId for mapping the raw data to the existing specimen
     */
    SpecimenIds,

    // Imported
    // ------------------------------
    /**
     * @description headers of the csv file
     */
    ImportedHeaders,

    /**
     * @description The display columns for the import
     */
    ImportColumns,

    /**
     *  @description  rows from the spreadsheet that are mapped to the specimenId
     */
    RawRows,

    /**
     * @description The header that contains the specimenId for mapping the raw data to the existing specimen
     */
    SpecimenIdHeader,

    // Processing Data
    // ------------------------------
    /**
     * @description The mapping of the imported csvs to specimens
     */
    SpecimenMapping,

    /**
     * @description The mapped values from the raw rows to them in specimen form
     */
    MappedRows,

    /**
     * @description The current specimen that is being updated
     */
    CurrentSpecimen,

    /**
     * @description The metadata for the rows
     */
    RowMetadata,

    LoadUpdates,

    ClearData,
    fetchSpecimens,
    SetMapping,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useBulkUpdate, import.meta.hot))
