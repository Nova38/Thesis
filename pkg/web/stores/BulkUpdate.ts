import { defineStore } from 'pinia'
import Papa, { type ParseResult } from 'papaparse'
import { spec } from 'node:test/reporters'
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

  /**
   *  RawRows is the raw data from the spreadsheet
   *
   */
  const RawRows = shallowRef<UpdateRow[]>()
  const UniqueRawHeaders = ref<string[]>()

  const ProcessingCSV = ref(false)

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
        ProcessingCSV.value = true
        console.table(results.data)

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

    UniqueRawHeaders,
    SpecimenIdHeader,

    ProcessingCSV,

    CollectionId,
    RawRowMap,
    SpecimenMapping,
    RowsSelected,

    LoadFromFile,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useBulkUpdate, import.meta.hot))
