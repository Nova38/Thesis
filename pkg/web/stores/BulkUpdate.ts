import { defineStore } from 'pinia'
import { EmptySpecimenMapping } from '~/utils/objects/Mapping'

// CatalogNumber is used to calculate the specimenId uuid to make sure it is unique
// SpecimenId is used to identify the specimen in the database

// Updating the Existing Specimen

export interface UpdateRowArgs {
  headers: string[]
  specimenIdHeader: string
  rows: Record<string, string>[]
}

export const useBulkUpdate = defineStore('BulkUpdate', () => {
  //
  const CollectionId = ref()

  /**
   * RawHeaders is the raw headers of the csv file
   */
  const RawHeaders = ref<string[]>()

  /**
   *  RawRows is the raw data from the spreadsheet
   *
   */
  const RawRowMap = ref<Record<string, Record<string, string>>>({})

  /**
   *  SpecimenIds is the specimenIds from the raw data
   */
  const SpecimenIds = ref<string[]>([])

  const Rows = ref<UpdateRow[]>()

  /**
   * SpecimenIdHeader is the header that contains the specimenId for mapping
   * the raw data to the existing specimen
   */
  const SpecimenIdHeader = ref('primary.specimenId')

  /**
   *
   */
  const ImportColumns = computed(() => {
    // Check if the RawHeaders is empty or if specimenIdHeader is in the RawHeaders and return the columns

    if (
      !RawHeaders.value ||
      !RawHeaders.value.includes(SpecimenIdHeader.value)
    ) {
      return {
        id: {},
      } as ImportColumns
    }

    let r: ImportCol[] =
      RawHeaders.value?.map((header) => {
        return {
          name: header,
          label: header,
          field: (row: UpdateRow) => row.raw[header],
          colType: header === SpecimenIdHeader.value ? 'id' : 'raw',
        }
      }) || []

    r = [
      {
        field: (row: UpdateRow) => row.meta?.status || 'unknown',
        name: 'Status',
        label: 'Status',
        colType: 'meta',
      },
      ...r,
    ]

    const id = r.find((col: ImportCol) => col.colType === 'id')
    if (!id) throw new Error('SpecimenIdHeader not found in RawHeaders')

    return {
      id,
      raw: r.filter((col: ImportCol) => col.colType === 'raw'),
      meta: r.filter((col: ImportCol) => col.colType === 'meta'),
    }
  })

  // const UniqueRawHeaders = ref<string[]>()

  // const ProcessingCSV = ref(false)

  /**
   * @description SpecimenMapping is the mapping of the raw data to the Specimen object
   */
  const SpecimenMapping = ref(EmptySpecimenMapping())

  const LoadUpdates = async (arg: UpdateRowArgs) => {
    RawHeaders.value = arg.headers
    SpecimenIdHeader.value = arg.specimenIdHeader

    // Reset The Row variables
    RawRowMap.value = {}

    arg.rows.forEach((row) => {
      const id = row[SpecimenIdHeader.value]
      if (!id) {
        console.warn(row)
        throw new Error('SpecimenIdHeader not found in RawHeaders')
      }

      const catNum = CatNumToUUID(id)
      SpecimenIds.value.push(catNum)

      RawRowMap.value[catNum] = row
    })

    const list = await $fetch('/api/cc/specimens/selectiveList', {
      body: {
        collectionId: CollectionId.value,
        specimenIds: SpecimenIds,
      },
    })

    if (!list) throw new Error('Failed to fetch full list')

    // asign the current data to a value
  }
  return {
    /*
     * Imported
     */
    Rows,
    ImportColumns,
    RawHeaders,

    LoadUpdates,

    // UniqueRawHeaders,
    SpecimenIdHeader,

    // ProcessingCSV,

    CollectionId,
    RawRowMap,
    SpecimenMapping,
    // RowsSelected,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useBulkUpdate, import.meta.hot))
