import { defineStore } from 'pinia'
import Papa, { type ParseResult } from 'papaparse'
import { EmptySpecimenMapping } from '~/utils/objects/Mapping'

// CatalogNumber is used to calculate the specimenId uuid to make sure it is unique
// SpecimenId is used to identify the specimen in the database

// Updating the Existing Specimen

export const useBulkUpdate = defineStore('BulkUpdate', () => {
  //
  const CollectionId = ref()

  /**
   * RawHeaders is the raw headers of the csv file
   */
  const RawHeaders = ref<string[]>()

  /**
   *
   */
  const ImportColumns = computed(() => {
    let r
      = RawHeaders.value?.map((header) => {
        return {
          name: header,
          label: header,
          field: (row: UpdateRow) => row.raw[header],
        }
      }) || []

    r = [
      {
        field: (row: UpdateRow) => row.meta?.status || 'unknown',
        name: 'Status',
        label: 'Status',
      },
      ...r,
    ]

    return r
  })

  /**
   *  RawRows is the raw data from the spreadsheet
   *
   */
  const RawRows = ref<UpdateRow[]>()

  /**
   * @description SpecimenMapping is the mapping of the raw data to the Specimen object
   */
  const SpecimenMapping = ref(EmptySpecimenMapping())

  const RowsSelected = ref([])

  // Mapped by catalogNumber
  const RawRowMap = ref(new Map<string, PlainSpecimen>())

  function LoadFromFile(file: File) {
    Papa.parse(file, {
      header: true,
      worker: true,
      complete: (results: ParseResult<Record<string, string>>) => {
        RawHeaders.value = results.meta.fields

        RawRows.value = results.data.map((row, index) => {
          return {
            id: index,
            meta: {
              exist: 'unknown',
              status: 'new',
              statusMessage: '',
              id: index,
            },
            raw: row,
          }
        })
      },
    })
  }

  return {
    /*
     * Imported
     */
    RawRows,
    ImportColumns,
    RawHeaders,

    CollectionId,
    RawRowMap,
    SpecimenMapping,
    RowsSelected,

    LoadFromFile,
    setStatus,
  }
})

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useBulkUpdate, import.meta.hot))
}
