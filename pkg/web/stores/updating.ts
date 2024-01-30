import { defineStore } from 'pinia'
import { EmptySpecimenMapping } from '~/utils/objects/Mapping'

// CatalogNumber is used to calculate the specimenId uuid to make sure it is unique
// SpecimenId is used to identify the specimen in the database

// Updating the Existing Specimen

export const useUpdatingStore = defineStore('Updating', () => {
  //
  const CollectionId = ref()

  // RawRows is the raw data from the spreadsheet
  const RawRows = ref([])

  // Mapped by catalogNumber
  const CurrentSpecimenMap = ref(new Map<string, PlainSpecimen>())

  // Mapped by catalogNumber
  const RawRowMap = ref(new Map<string, PlainSpecimen>())

  const SpecimenMapping = ref(EmptySpecimenMapping())

  return {
    RawRows,
    CollectionId,
    CurrentSpecimenMap,
    RawRowMap,
    SpecimenMapping,
  }
})
