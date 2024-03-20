import { diff } from 'ohash'
import { get, set } from 'radash'
import { ccbio } from 'saacs'

export const useBulkStore = defineStore('Bulk', () => {
  const CollectionId = ref<string>('')
  const Mode = ref<BulkMode>('import')

  const SpecimenIdHeader = ref('primary.specimenId')
  const SpecimenIds = ref<string[]>([])

  //
  const MappedSpecimen = ref<PlainSpecimen[]>([])

  // CSV Data
  // ------------------------------
  const RawRows = shallowRef<UpdateRawRow[]>([])
  const RawRowsMeta = ref(new Map<string, ImportRowMeta>())
  const RawHeaders = ref<string[]>([])

  const LoadCsv = async (csv: CSVImportMetadata) => {
    SpecimenIds.value = []
    MappedSpecimen.value = []

    RawRowsMeta.value = new Map<string, ImportRowMeta>()

    SpecimenIdHeader.value = csv.specimenIdHeader

    RawHeaders.value = csv.headers

    RawRows.value = csv.rows.map((row) => {
      const catNum = row[SpecimenIdHeader.value]
      if (!catNum) {
        console.warn(row)
        throw new Error('SpecimenIdHeader not found in RawHeaders')
      }

      const id = CatNumToUUID(catNum)
      SpecimenIds.value.push(id)

      RawRowsMeta.value.set(id, {
        exist: 'unknown',
        status: 'loading',
        statusMessage: 'loading from csv',
      })

      MappedSpecimen.value.push(new ccbio.Specimen({
        collectionId: CollectionId.value,
        specimenId: id,
        primary: {
          catalogNumber: catNum,
        },
        secondary: {},
        georeference: {},
        grants: {},
        images: {},
        loans: {},
        taxon: {},
        lastModified: {},
      }))

      return {
        id,
        raw: row,
      }
    })

    // await nextTick()
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
      if (SpecimenIds.value.length === 0)
        return specimens

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
          specimens.set(key, Object.freeze(ccbio.Specimen.fromJsonString(JSON.stringify(value))))
        })

        return Object.freeze(specimens)
      }
      catch (e) {
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
    if (!resolved)
      return

    RawRowsMeta.value.forEach((meta, id) => {
      if (CurrentSpecimens.value?.has(id)) {
        meta.status = 'pre-existing'
        meta.exist = 'pre-existing'
      }
      else {
        meta.exist = 'new'
      }
    })

    CurrentSpecimens.value?.forEach((cur) => {
      const meta = RawRowsMeta.value.get(cur.specimenId)
      if (!meta)
        return
      meta.exist = 'pre-existing'

      // Set the value of the role to be the preexisting value
      const index = MappedSpecimen.value.findIndex(o => o.specimenId === cur.specimenId)

      if (index === -1)
        throw new Error(`The specimen with id = ${cur.specimenId} was not found in the current rows`)

      const c = new ccbio.Specimen(cur)

      MappedSpecimen.value[index] = c.clone()
    })
  })

  const SetMapping = (
    newMapping: string,
    col: { name: string, mapped: string },
  ) => {
    const def = RawColDefs.value.find(c => c.name === col.name)
    if (def === undefined)
      throw new Error('Attempted to set non-existent row for mapping')

    // Update the mapping definitions

    // Update the Mapped Specimen
    MappedSpecimen.value = MappedSpecimen.value.map((mapped) => {
      const cur = CurrentSpecimens.value?.get(mapped.specimenId)
      const rawV = RawRows.value.find(x => x.id === mapped.specimenId)
      if (!rawV)
        throw new Error('RawRow not found')
      console.group(`SetMapping: ${mapped.specimenId}`)
      // console.log(rawV)

      const meta = RawRowsMeta.value.get(mapped.specimenId)
      if (!meta)
        throw new Error('meta missing')
      // console.log(meta)
      // Empty/Unset
      if (newMapping === '' || newMapping === ' ') {
        // If specimen is current
        if (cur) {
          mapped = set(mapped, def.mapped, get(cur, def.mapped))
          console.log({
            id: mapped.specimenId,
            action: 'Clearing Value: Resting to Current',
            newMapping,
            col: col.name,
            mapped,

          })
        }
        else {
          mapped = set(mapped, def.mapped, '')
          console.log({
            id: mapped.specimenId,

            action: 'Clearing Value: Setting to undefined',
            newMapping,
            col: col.name,
            mapped,
          })
        }
      }

      // Date type
      // Set Any
      else {
        console.log(
          {
            action: 'Setting Value',
            newMapping,
            raw: rawV.raw,
            rawV: rawV.raw[col.name],
          },
        )
        set(mapped, newMapping, rawV.raw[col.name])
      }

      const parsed = ZSpecimen.safeParse(mapped)

      if (parsed.success === false) {
        meta.status = 'parsing-error'
        meta.statusMessage = parsed.error.toString()
        meta.error = parsed.error
      }
      else if (meta.status === 'parsing-error') {
        meta.status = 'loading'
        meta.statusMessage = ''
      }
      console.log(mapped)
      console.groupEnd()
      return mapped
    },
    )
    def.mapped = newMapping

    console.log(col, newMapping, def.mapped, RawColDefs.value)
  }

  // const SetMapping = (
  //   selected: string | undefined,
  //   col: { name: string, mapped: string },
  // ) => {
  //   if (!selected)
  //     return

  //   const c = RawColDefs.value.find(c => c.name === col.name)

  //   if (c === undefined)
  //     return
  //   c.mapped = selected
  // }

  // // Processed Rows

  // watchEffect(() => {
  //   MappedSpecimen.value = RawRows.value.map((row) => {
  //     const mapped: Record<string, string> = {}

  //     RawColDefs.value.forEach((col) => {
  //       if (typeof col.field === 'function')
  //         mapped[col.mapped] = col.field(row)
  //       else mapped[col.mapped] = row.raw?.[col.field]
  //     })

  //     const current = CurrentSpecimens.value?.get(row.id)

  //     const update = assign(
  //       {
  //         georeference: {
  //           georeferenceDate: {},
  //         },
  //         grants: {},
  //         loans: {},
  //         primary: {
  //           catalogDate: {},
  //           determinedDate: {},
  //           fieldDate: {},
  //           originalDate: {},
  //         },
  //         secondary: {
  //           preparations: {},
  //         },
  //         taxon: {},
  //       },
  //       construct(mapped),
  //     )

  //     const unflat = assign(current ?? {}, update) as PlainSpecimen

  //     unflat.specimenId = row.id
  //     unflat.collectionId = CollectionId.value

  //     console.log(unflat)

  //     const converted = ZSpecimen.safeParse(unflat)
  //     const meta = RawRowsMeta.value.get(row.id)
  //     if (!meta)
  //       throw new Error('meta missing')

  //     if (converted.success) {
  //       if (meta.status === 'parsing-error') {
  //         meta.status = 'loading'
  //         meta.statusMessage = ''
  //       }

  //       return new ccbio.Specimen(converted.data)
  //     }
  //     else {
  //       meta.error = converted.error
  //       meta.status = 'parsing-error'
  //       meta.statusMessage = converted.error.toString()
  //     }
  //     return new ccbio.Specimen(unflat)
  //   })
  // })

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
    RawRowsMeta.value = new Map()
    RawHeaders.value = []
    SpecimenIds.value = []
    MappedSpecimen.value = []
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

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useBulkStore, import.meta.hot))
