import { unflatten } from 'flat'
import { diff, objectHash } from 'ohash'
import { assign, construct } from 'radash'
import { ccbio } from '@/lib'

export const useBulkStore = defineStore('Bulk', () => {
  const CollectionId = ref<string>('')
  const Mode = ref<BulkMode>('import')

  const SpecimenIdHeader = ref('primary.specimenId')
  const SpecimenIds = ref<string[]>([])

  // CSV Data
  // ------------------------------
  const RawRows = shallowRef<UpdateRawRow[]>([])
  const RawRowsMeta = ref(new Map<string, ImportRowMeta>())
  const RawHeaders = ref<string[]>([])

  const LoadCsv = (csv: CSVImportMetadata) => {
    SpecimenIds.value = []
    RawRowsMeta.value = new Map<string, ImportRowMeta>()

    SpecimenIdHeader.value = csv.specimenIdHeader

    RawHeaders.value = csv.headers

    RawRows.value = csv.rows.map((row) => {
      const id = row[SpecimenIdHeader.value]
      if (!id) {
        console.warn(row)
        throw new Error('SpecimenIdHeader not found in RawHeaders')
      }

      const catNum = CatNumToUUID(id)
      SpecimenIds.value.push(catNum)

      RawRowsMeta.value.set(catNum, {
        exist: 'unknown',
        status: 'loading',
        statusMessage: 'loading from csv',
      })
      console.log({
        id: catNum,
        raw: row,
      })
      return {
        id: catNum,
        raw: row,
      }
    })
  }

  const RawColDefs = ref<BulkImportHeader[]>([])
  // Build the imported Columns
  watch(RawHeaders, () => {
    RawColDefs.value = RawHeaders.value.map((raw) => {
      return {
        name: `${raw}`,
        field: (row: UpdateRawRow) => row?.raw?.[raw] ?? '',

        mapped: '',
        isID: raw === SpecimenIdHeader.value,
        updated: false,
      }
    })
  })

  // Current Data
  // ------------------------------
  // const CurrentSpecimen = ref<Map<string, PlainSpecimen>>(new Map())
  // watch(SpecimenIds, ())
  const CurrentSpecimensEvaluating = ref(false)
  const CurrentSpecimens = computedAsync(
    async () => {
      const specimens = new Map<string, PlainSpecimen>()
      if (SpecimenIds.value.length === 0) return specimens

      console.log(SpecimenIds)

      try {
        const res = await $fetch('/api/cc/specimens/bulk/partialList', {
          method: 'post',
          query: {
            collectionId: CollectionId.value,
          },
          body: {
            specimenIds: Array.from(SpecimenIds.value),
          },
        })

        res.filteredList.forEach(([key, value]) => {
          specimens.set(key, new ccbio.Specimen(value))
        })

        return Object.freeze(specimens)
      } catch (e) {
        console.error(e)
        throw new Error('Failed To fetch specimens from specimenIds', {
          cause: e,
        })
      }
    },
    null,
    { lazy: true, evaluating: CurrentSpecimensEvaluating },
  )

  // Set the row to existing if the Specimen exists, after it is done resolving
  watch(CurrentSpecimens, (resolved) => {
    if (!resolved) return

    RawRowsMeta.value.forEach((meta, id) => {
      if (CurrentSpecimens.value?.has(id)) {
        meta.status = 'pre-existing'
        meta.exist = 'pre-existing'
      } else {
        meta.exist = 'new'
      }
    })

    CurrentSpecimens.value?.forEach((cur) => {
      const meta = RawRowsMeta.value.get(cur.specimenId)
      if (!meta) return
      meta.exist = 'pre-existing'
    })
  })

  const SetMapping = (
    selected: string | undefined,
    col: { name: string; mapped: string },
  ) => {
    if (!selected) return

    const c = RawColDefs.value.find((c) => c.name === col.name)

    if (c === undefined) return
    c.mapped = selected
  }

  // Processed Rows

  const MappedSpecimen = ref<PlainSpecimen[]>([])

  watchEffect(() => {
    MappedSpecimen.value = RawRows.value.map((row) => {
      const mapped: Record<string, string> = {}

      RawColDefs.value.forEach((col) => {
        if (typeof col.field === 'function') mapped[col.mapped] = col.field(row)
        else mapped[col.mapped] = row.raw?.[col.field]
      })

      const current = CurrentSpecimens.value?.get(row.id)

      const update = assign(
        {
          georeference: {
            georeferenceDate: {},
          },
          grants: {},
          loans: {},
          primary: {
            catalogDate: {},
            determinedDate: {},
            fieldDate: {},
            originalDate: {},
          },
          secondary: {
            preparations: {},
          },
          taxon: {},
        },
        construct(mapped),
      )

      const unflat = assign(current ?? {}, update) as PlainSpecimen

      unflat.specimenId = row.id
      unflat.collectionId = CollectionId.value

      console.log(unflat)

      const converted = ZSpecimen.safeParse(unflat)
      const meta = RawRowsMeta.value.get(row.id)
      if (!meta) throw new Error('meta missing')

      if (converted.success) {
        if (meta.status === 'parsing-error') {
          meta.status = 'loading'
          meta.statusMessage = ''
        }

        return new ccbio.Specimen(converted.data)
      } else {
        meta.error = converted.error
        meta.status = 'parsing-error'
        meta.statusMessage = converted.error.toString()
      }
      return new ccbio.Specimen(unflat)
    })
  })

  const differences = computed(() => {
    return MappedSpecimen.value.map((specimen) => {
      const base = CurrentSpecimens.value?.get(specimen.specimenId) ?? {}

      return diff(base, specimen)
      return diffCrush(base, specimen, [])
    })
  })

  const Upload = async () => {
    console.log(differences.value)
  }

  const $reset = () => {
    RawRows.value = []
  }

  return {
    CollectionId,
    Mode,

    RawRows,
    RawRowsMeta,
    RawHeaders,
    SpecimenIds,

    RawColDefs,
    SpecimenIdHeader,

    CurrentSpecimens,
    CurrentSpecimensEvaluating,

    MappedSpecimen,
    differences,

    LoadCsv,
    SetMapping,
    Upload,

    $reset,
  }
})
