import { diff } from 'ohash'
import { cluster, get, set } from 'radash'
// import { ccbio } from 'saacs'
import { ccbio, type PlainSpecimen } from '#imports'
import type { FieldMask } from '@bufbuild/protobuf'
import type { MeterItem } from 'primevue/metergroup'
export const useBulkStore = defineStore('Bulk', () => {
  const toast = useToast()

  const Loading = ref(false)

  const CollectionId = ref<string>('')
  const Mode = ref<BulkMode>('hybrid')

  const SpecimenIdHeader = ref('primary.specimenId')
  const SpecimenIds = ref<string[]>([])

  //
  const MappedSpecimen = ref<PlainSpecimen[]>([])

  // CSV Data
  // ------------------------------
  const RawRows = shallowRef<UpdateRawRow[]>([])
  const RawRowsMeta = ref(new Map<string, ImportRowMeta>())
  const RawHeaders = ref<string[]>([])

  const differences = ref(new Map<string, FieldMask>())

  const processing = ref<{ success: number; fail: number; total: number }>({
    success: 0,
    fail: 0,
    total: 0,
  })

  const ImportStatus = ref<MeterItem[]>([
    { label: 'PreExisting', value: 0, color: 'var(--v-primary)', icon: '' },
    { label: 'New', value: 0, color: 'var(--v-primary)', icon: '' },
    { label: 'Error', value: 0, color: 'var(--v-primary)', icon: '' },
  ])

  const UploadStatus = ref<MeterItem[]>([
    { label: 'Pending', value: 0, color: 'var(--v-primary)', icon: '' },
    { label: 'Uploaded', value: 0, color: 'var(--v-primary)', icon: '' },
    { label: 'Error', value: 0, color: 'var(--v-primary)', icon: '' },
  ])

  interface localInterface {
    collectionId: string
    specimenIdHeader: string
    headers: string[]
    rows: Record<string, string>[]
    rawRowsMeta: Map<string, ImportRowMeta>
    mappedSpecimen: PlainSpecimen[]
    specimenIds: string[]
    rawRows: UpdateRawRow[]
  }
  const LoadCsv = async (csv: CSVImportMetadata) => {
    Loading.value = true

    const local: localInterface = {
      collectionId: CollectionId.value,
      specimenIdHeader: SpecimenIdHeader.value,
      headers: csv.headers,
      rows: csv.rows,
      rawRowsMeta: new Map<string, ImportRowMeta>(),
      mappedSpecimen: [],
      specimenIds: [],
      rawRows: [],
    }

    // SpecimenIds.value = []
    // MappedSpecimen.value = []

    // RawRowsMeta.value = new Map<string, ImportRowMeta>()

    // SpecimenIdHeader.value = csv.specimenIdHeader
    // RawHeaders.value = csv.headers

    local.rawRows = csv.rows.map((row) => {
      const catNum = row[csv.specimenIdHeader]
      if (!catNum) {
        console.warn(row)
        throw new Error('SpecimenIdHeader not found in RawHeaders')
      }

      const id = CatNumToUUID(catNum)
      local.specimenIds.push(id)

      local.rawRowsMeta.set(id, {
        exist: 'unknown',
        status: 'loading',
        statusMessage: 'loading from csv',
      })

      local.mappedSpecimen.push(
        new ccbio.Specimen({
          collectionId: CollectionId.value,
          specimenId: id,
          primary: {
            catalogNumber: catNum,
          },
          secondary: {
            preparations: {
              '1': {},
            },
          },
          georeference: {},
          grants: {
            1: {},
          },
          images: {},
          loans: {
            '1': {},
          },
          taxon: {},
          lastModified: {},
        }),
      )

      return {
        id,
        raw: row,
      }
    })
    Loading.value = false

    SpecimenIds.value = local.specimenIds
    MappedSpecimen.value = local.mappedSpecimen

    RawRowsMeta.value = local.rawRowsMeta

    SpecimenIdHeader.value = csv.specimenIdHeader
    RawHeaders.value = csv.headers

    RawRows.value = local.rawRows
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
      Loading.value = true
      const specimens = new Map<string, PlainSpecimen>()
      if (SpecimenIds.value.length === 0) return specimens

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
          // const c = ccbio.Specimen.fromJsonString(JSON.stringify(value))
          const c = value as PlainSpecimen
          if (c?.primary?.lastModified) {
            c.primary.lastModified = undefined
          }
          if (c?.secondary?.lastModified) {
            c.secondary.lastModified = undefined
          }
          if (c?.georeference?.lastModified) {
            c.georeference.lastModified = undefined
          }
          if (c?.taxon?.lastModified) {
            c.taxon.lastModified = undefined
          }
          if (c?.lastModified) {
            c.lastModified = undefined
          }

          specimens.set(key, Object.freeze(c))
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

      // Set the value of the role to be the preexisting value
      const index = MappedSpecimen.value.findIndex(
        (o) => o.specimenId === cur.specimenId,
      )

      if (index === -1)
        throw new Error(
          `The specimen with id = ${cur.specimenId} was not found in the current rows`,
        )

      const c = new ccbio.Specimen(cur)

      MappedSpecimen.value[index] = c.clone()
    })
  })

  const SetMapping = (
    newMapping: string,
    col: { name: string; mapped: string },
  ) => {
    const def = RawColDefs.value.find((c) => c.name === col.name)
    if (def === undefined)
      throw new Error('Attempted to set non-existent row for mapping')

    // Update the mapping definitions

    // Update the Mapped Specimen
    MappedSpecimen.value = MappedSpecimen.value.map((mapped) => {
      const cur = CurrentSpecimens.value?.get(mapped.specimenId)
      const rawV = RawRows.value.find((x) => x.id === mapped.specimenId)
      if (!rawV) throw new Error('RawRow not found')
      console.group(`SetMapping: ${mapped.specimenId}`)

      const meta = RawRowsMeta.value.get(mapped.specimenId)
      if (!meta) throw new Error('meta missing')
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
        } else {
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
        console.log({
          action: 'Setting Value',
          newMapping,
          raw: rawV.raw,
          rawV: rawV.raw[col.name],
        })
        set(mapped, newMapping, rawV.raw[col.name])
      }

      const parsed = ZSpecimen.safeParse(mapped)

      if (parsed.success === false) {
        meta.status = 'parsing-error'
        meta.statusMessage = parsed.error.toString()
        meta.error = parsed.error
      } else if (meta.status === 'parsing-error') {
        meta.status = 'loading'
        meta.statusMessage = ''
      }
      console.log(mapped)
      console.groupEnd()
      return mapped
    })
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

  watch(MappedSpecimen, () => {
    MappedSpecimen.value.forEach((specimen) => {
      const base = CurrentSpecimens.value?.get(specimen.specimenId) ?? {}
      const meta = RawRowsMeta.value.get(specimen.specimenId)
      if (!meta) throw new Error('meta missing')

      if (meta.status === 'pre-existing') {
        const mask = diffToFieldMaskPath(base, specimen).mask
        console.log('diff', mask)

        differences.value.set(specimen.specimenId, mask)
        meta.statusMessage = mask?.paths.join(', ')
      }
    })
  })

  const importSpecimens = async () => {
    console.log('importing specimens', Mode.value)
    processing.value.total = MappedSpecimen.value.length

    // for (const group of groups) {
    //   const uploads = group.map(async (specimen) => {
    for (const specimen of MappedSpecimen.value) {
      const meta = RawRowsMeta.value.get(specimen.specimenId)
      if (!meta) throw new Error('meta missing')
      if (meta.status === 'parsing-error') {
        console.error('Skipping Specimen due to parsing error', specimen)
        continue
      }

      switch (meta.exist) {
        case 'pre-existing': {
          if (Mode.value === 'import') {
            console.warn('Invalid Mode: can not import in import mode')
            continue
          }

          meta.status = 'submitting'
          const body = {
            mask: differences.value.get(specimen.specimenId)?.paths ?? {},
            specimen: specimen,
          }
          // meta.statusMessage = 'updating specimen'
          try {
            const r = await $fetch('/api/cc/specimens/update', {
              method: 'post',
              body,
            })
            meta.status = 'success'
            meta.statusMessage = 'updated successfully'
            processing.value.success++
            console.log('Imported Specimen', r)

            toast.add({
              id: 'imported',
              title: `Imported Specimen ${specimen.primary?.catalogNumber}`,
              timeout: 5000,
            })
          } catch (error) {
            console.error('Error updating specimen', error)
            meta.status = 'error'
            processing.value.fail++
            if (error instanceof Error) {
              toast.add({
                id: 'failedImported',
                title: `Failed to Imported Specimen ${specimen.primary?.catalogNumber}`,
                timeout: 5000,
                description: error.toString(),
              })
              meta.statusMessage = error.toString()
            }
          }

          break
        }
        case 'new': {
          if (Mode.value === 'update') {
            console.warn('Invalid Mode: can not update in import mode')
            continue
          }

          meta.status = 'submitting'
          meta.statusMessage = 'importing specimen'
          try {
            const r = await $fetch('/api/cc/specimens/create', {
              method: 'post',
              body: new ccbio.Specimen(specimen).toJsonString({
                emitDefaultValues: true,
                enumAsInteger: true,
              }),
            })
            meta.status = 'success'
            meta.statusMessage = 'imported successfully'
            console.log('Imported Specimen', r)
            processing.value.success++
          } catch (error) {
            console.error('Error importing specimen', error)
            meta.status = 'error'
            processing.value.fail++
            if (error instanceof Error) meta.statusMessage = error.toString()
          }

          break
        }
        case 'unknown':
          break

        default:
          break
      }
    }

    // await Promise.all(uploads)
  }

  const UpdateSpecimens = async () => {}

  const Upload = async () => {
    await importSpecimens()
  }

  const $reset = () => {
    RawRows.value = []
    RawRowsMeta.value = new Map()
    RawHeaders.value = []
    SpecimenIds.value = []
    MappedSpecimen.value = []
    differences.value = new Map()
  }

  return {
    CollectionId,
    Mode,

    ImportStatus,
    UploadStatus,

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

    importSpecimens,
    UpdateSpecimens,

    $reset,
  }
})

if (import.meta.hot)
  import.meta.hot.accept(acceptHMRUpdate(useBulkStore, import.meta.hot))
