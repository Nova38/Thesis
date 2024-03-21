<script lang="ts" setup>
import { ccbio } from 'saacs'

const props = defineProps({
  enableEdit: {
    default: true,
    type: Boolean,
  },
  headerColor: {
    default: toModeColor('view'),
    type: String,
  },
  showEditOptions: {
    default: true,
    type: Boolean,
  },
  showEmpty: {
    default: true,
    type: Boolean,
  },
  showHeader: {
    default: false,
    type: Boolean,
  },
  showUpdater: {
    default: true,
    type: Boolean,
  },
  startOpen: {
    default: true,
    type: Boolean,
  },
})

const sectionHeaderClass = 'text-center section-header bg-gray-200'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

function addPreparation(name: string) {
  if (!specimen?.value) return
  if (!specimen?.value?.secondary)
    specimen.value.secondary = new ccbio.Specimen_Secondary()

  if (!specimen?.value?.secondary?.preparations)
    specimen.value.secondary.preparations = {}

  specimen.value.secondary.preparations[name] =
    new ccbio.Specimen_Secondary_Preparation()
}

function addLoan(name: string) {
  if (!specimen?.value) return
  if (!specimen?.value?.loans) specimen.value.loans = {}

  specimen.value.loans[name] = new ccbio.Specimen_Loan({
    id: name,
  })
}

function addGrant(name: string) {
  if (!specimen?.value) return
  if (!specimen?.value?.grants) specimen.value.grants = {}

  specimen.value.grants[name] = new ccbio.Specimen_Grant({})
}
const newPreparationName = ref('')
const newLoanName = ref('')
const newGrantName = ref('')
// ^?

const secondarySexOptions = [
  {
    label: '',
    value: ccbio.Specimen_Secondary_SEX.SEX_UNDEFINED,
  },
  {
    label: 'Unknown',
    value: ccbio.Specimen_Secondary_SEX.SEX_UNKNOWN,
  },
  {
    label: 'Atypical',
    value: ccbio.Specimen_Secondary_SEX.SEX_ATYPICAL,
  },
  {
    label: 'Male',
    value: ccbio.Specimen_Secondary_SEX.SEX_MALE,
  },
  {
    label: 'Female',
    value: ccbio.Specimen_Secondary_SEX.SEX_FEMALE,
  },
]

const secondaryAgeOptions = [
  {
    label: '',
    value: ccbio.Specimen_Secondary_AGE.AGE_UNDEFINED,
  },
  {
    label: 'Unknown',
    value: ccbio.Specimen_Secondary_AGE.AGE_UNKNOWN,
  },
  {
    label: 'Nest',
    value: ccbio.Specimen_Secondary_AGE.AGE_NEST,
  },
  {
    label: 'Embryo Egg',
    value: ccbio.Specimen_Secondary_AGE.AGE_EMBRYO_EGG,
  },
  {
    label: 'Chick SubAdult',
    value: ccbio.Specimen_Secondary_AGE.AGE_CHICK_SUBADULT,
  },
  {
    label: 'Adult',
    value: ccbio.Specimen_Secondary_AGE.AGE_ADULT,
  },
  {
    label: 'Age Contingent',
    value: ccbio.Specimen_Secondary_AGE.AGE_CONTINGENT,
  },
]

const sections = [
  {
    label: 'Taxon',
    slot: 'taxon',
  },
]
</script>

<template>
  <QCard class="max-w-4xl grow">
    <slot name="Header" />

    <QCardSection :class="headerColor">
      <div class="row">
        <div v-if="specimen?.taxon" class="text-2xl">
          <span v-if="specimen?.taxon?.genus">
            {{ specimen.taxon.genus }}
          </span>
          {{}}
          <span v-if="specimen?.taxon?.species">{{
            specimen?.taxon?.species
          }}</span>
        </div>
        <div class="flex-col ml-auto">
          <div
            v-if="specimen?.primary"
            class="text-subtitle2"
            style="margin-left: auto"
          >
            <q-chip
              v-if="specimen?.primary.catalogNumber"
              :label="`Catalog Number: ${specimen?.primary?.catalogNumber}`"
              class="text-white"
              color="accent"
              square
            />
            <q-chip
              v-if="specimen.primary.tissueNumber"
              :label="`Tissue Number: ${specimen?.primary.tissueNumber}`"
              class="text-white"
              color="accent"
              square
            />
          </div>
        </div>
      </div>
    </QCardSection>

    <div :class="headerColor" class="flex flex-row">
      <q-chip
        v-if="specimen.collectionId"
        :label="`Collection ID: ${specimen.collectionId}`"
        class="text-white flex-grow"
        color="primary"
        square
      />
      <q-chip
        v-if="specimen.specimenId"
        :label="`Specimen ID: ${specimen.specimenId}`"
        class="text-white flex-grows"
        color="primary"
        square
      />
    </div>

    <q-form>
      <UAccordion
        :multiple="true"
        ui=""
        :start-open="startOpen"
        :show-header="showHeader"
        :show-empty="showEmpty"
        :show-edit-options="showEditOptions"
        :items="sections"
      >
        <template #taxon>
          <QCardSection class="flex flex-row gap-1 justify-evenly mb-1">
            <template v-for="(_, key) in specimen.taxon" :key="key">
              <QInput
                v-if="specimen.taxon && typeof specimen.taxon[key] === 'string'"
                :model-value="specimen.taxon[key] as string"
                :disable="!props.enableEdit"
                :label="key"
                class="flex-grow my-1"
              />
            </template>
          </QCardSection>
          <div class="m-4">
            <q-item-label overline>
              <b>Last Updated At</b>
            </q-item-label>
            <q-item-label class="text-bold">
              {{ specimen?.primary?.lastModified?.timestamp }}
            </q-item-label>
          </div>
        </template>
      </UAccordion>
      <QExpansionItem
        :header-class="sectionHeaderClass"
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        default-opened
        expand-icon-toggle
        label="Taxon"
      >
        <QCard>
          <QCardSection class="flex flex-row gap-1 justify-evenly mb-1">
            <template v-for="(_, key) in specimen.taxon" :key="key">
              <QInput
                v-if="specimen.taxon && typeof specimen.taxon[key] === 'string'"
                v-model="specimen.taxon[key]"
                :disable="!props.enableEdit"
                :label="key"
                class="flex-grow my-1"
              />
            </template>
          </QCardSection>
          <div class="m-4">
            <q-item-label overline>
              <b>Last Updated At</b>
            </q-item-label>
            <q-item-label class="text-bold">
              {{ specimen?.primary?.lastModified?.timestamp }}
            </q-item-label>
          </div>
        </QCard>
      </QExpansionItem>

      <q-separator inset />

      <q-expansion-item
        :header-class="sectionHeaderClass"
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        default-opened
        expand-icon-toggle
        label="Primary"
      >
        <div v-if="specimen?.primary">
          <q-card-section class="my-1">
            <!-- <div class="text-h6">Overline</div> -->
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.tissueNumber"
                :disable="!props.enableEdit"
                class="m-1 flex-grow"
                label="Tissue Number"
              />
              <q-input
                v-model="specimen.primary.accessionNumber"
                :disable="!props.enableEdit"
                class="m-1 flex-grow"
                label="Accession Number"
              />
            </div>
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.cataloger"
                :disable="!props.enableEdit"
                class="flex-grow m-1"
                label="Cataloger"
              />
              <q-input
                v-model="specimen.primary.catalogNumber"
                :disable="!props.enableEdit"
                class="flex-grow m-1"
                label="Catalog Number"
              />
            </div>
            <QExpansionItem class="DateExpansion" label="Catalog Date">
              <QCardSection
                v-if="specimen?.primary?.catalogDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.catalogDate.verbatim"
                  :disable="!props.enableEdit"
                  class="flex-grow"
                  label="Catalog Date: Verbatim"
                  type="text"
                />
                <!-- <QInput
                    v-model="specimen.georeference.georeferenceDate.timestamp"
                  :disable="!props.enableEdit"
                  type="date"
                  label="Timestamp"
                    /> -->
                <QInput
                  v-model.number="specimen.primary.catalogDate.year"
                  :disable="!props.enableEdit"
                  label="Catalog Date: Year"
                  number
                  type="number"
                />
                <QInput
                  v-model="specimen.primary.catalogDate.month"
                  :disable="!props.enableEdit"
                  label="Catalog Date: Month"
                  type="text"
                />
                <QInput
                  v-model.number="specimen.primary.catalogDate.day"
                  :disable="!props.enableEdit"
                  label="Catalog Date: Day"
                  number
                  type="number"
                />
              </QCardSection>
            </QExpansionItem>
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.determiner"
                :disable="!props.enableEdit"
                class="flex-grow m-1"
                label="Determiner"
              />
              <q-input
                v-model="specimen.primary.determinedReason"
                :disable="!props.enableEdit"
                class="flex-grow m-1"
                label="Determined Reason"
              />
            </div>
            <QExpansionItem class="DateExpansion mx-4" label="Determined Date">
              <QCardSection
                v-if="specimen?.primary?.determinedDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.determinedDate.verbatim"
                  :disable="!props.enableEdit"
                  class="flex-grow"
                  label="Determined Date: Verbatim"
                  type="text"
                />
                <!-- <QInput
                    v-model="specimen.georeference.georeferenceDate.timestamp"
                  :disable="!props.enableEdit"
                  type="date"
                  label="Timestamp"
                /> -->
                <QInput
                  v-model.number="specimen.primary.determinedDate.year"
                  :disable="!props.enableEdit"
                  label="Determined Date: Year"
                  number
                  type="number"
                />
                <QInput
                  v-model="specimen.primary.determinedDate.month"
                  :disable="!props.enableEdit"
                  label="Determined Date: Month"
                  type="text"
                />
                <QInput
                  v-model.number="specimen.primary.determinedDate.day"
                  :disable="!props.enableEdit"
                  label="Determined Date: Day"
                  number
                  type="number"
                />
              </QCardSection>
            </QExpansionItem>
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.collector"
                :disable="!props.enableEdit"
                class="flex-grow m-1"
                label="Collector"
              />
              <q-input
                v-model="specimen.primary.fieldNumber"
                :disable="!props.enableEdit"
                class="flex-grow m-1"
                label="Field Number"
              />
            </div>

            <QExpansionItem class="DateExpansion" label="Field Date">
              <QCardSection
                v-if="specimen?.primary?.fieldDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.fieldDate.verbatim"
                  :disable="!props.enableEdit"
                  class="flex-grow"
                  label="Field Date: Verbatim"
                  type="text"
                />
                <!-- <QInput
                    v-model="specimen.georeference.georeferenceDate.timestamp"
                  :disable="!props.enableEdit"
                  type="date"
                  label="Timestamp"
                /> -->
                <QInput
                  v-model.number="specimen.primary.fieldDate.year"
                  :disable="!props.enableEdit"
                  label="Field Date: Year"
                  number
                  type="number"
                />
                <QInput
                  v-model="specimen.primary.fieldDate.month"
                  :disable="!props.enableEdit"
                  label="Field Date: Month"
                  type="text"
                />
                <QInput
                  v-model.number="specimen.primary.fieldDate.day"
                  :disable="!props.enableEdit"
                  label="Field Date: Day"
                  number
                  type="number"
                />
              </QCardSection>
            </QExpansionItem>
            <QExpansionItem
              v-if="specimen?.primary?.originalDate"
              class="DateExpansion"
              label="Original Date"
            >
              <QCardSection
                v-if="specimen?.georeference?.georeferenceDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.originalDate.verbatim"
                  :disable="!props.enableEdit"
                  class="flex-grow"
                  label="Original Date: Verbatim"
                  type="text"
                />
                <!-- <QInput
                    v-model="specimen.georeference.georeferenceDate.timestamp"
                  :disable="!props.enableEdit"
                  type="date"
                  label="Timestamp"
                /> -->
                <QInput
                  v-model.number="specimen.primary.originalDate.year"
                  :disable="!props.enableEdit"
                  label="Original Date: Year"
                  number
                  type="number"
                />
                <QInput
                  v-model="specimen.primary.originalDate.month"
                  :disable="!props.enableEdit"
                  label="Original Date: Month"
                  type="text"
                />
                <QInput
                  v-model.number="specimen.primary.originalDate.day"
                  :disable="!props.enableEdit"
                  label="Original Date: Day"
                  number
                  type="number"
                />
              </QCardSection>
            </QExpansionItem>
          </q-card-section>
        </div>
        <!-- <q-card-section>
            <q-item-label overline><b>Primary Last Updater</b></q-item-label>
            <q-item-label class="text-bold">
              {{ specimen?.primary?.lastModified?.timestamp }}
            </q-item-label>
          </q-card-section> -->
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        :header-class="sectionHeaderClass"
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        default-opened
        expand-icon-toggle
        label="Georeference"
      >
        <q-card>
          <q-card-section v-if="specimen?.georeference" class="my-1">
            <div class="row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.continent"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Continent"
              />
              <q-input
                v-model="specimen.georeference.country"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Country"
              />
              <q-input
                v-model="specimen.georeference.stateProvince"
                :disable="!props.enableEdit"
                class="col m-1"
                label="State/Province"
              />
              <q-input
                v-model="specimen.georeference.county"
                :disable="!props.enableEdit"
                class="col m-1"
                label="County"
              />
            </div>

            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.habitat"
                :disable="!props.enableEdit"
                autogrow
                class="col m-1"
                label="Habitat"
                type="textarea"
              />
              <q-input
                v-model="specimen.georeference.locality"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Locality"
              />
            </div>
            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model.number="specimen.georeference.longitude"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Longitude"
                type="number"
              />

              <q-input
                v-model.number="specimen.georeference.latitude"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Latitude"
                type="number"
              />
              <q-input
                v-model.number="
                  specimen.georeference.coordinateUncertaintyInMeters
                "
                :disable="!props.enableEdit"
                class="col m-1"
                label="Coordinate Uncertainty In Meters"
                number
              />
            </div>
            <div class="row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.georeferenceBy"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Georeference By"
              />
              <q-input
                v-model="specimen.georeference.georeferenceProtocol"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Georeference Protocol"
              />
              <q-input
                v-model="specimen.georeference.geodeticDatum"
                :disable="!props.enableEdit"
                class="col m-1"
                label="Geodetic Datum"
              />
            </div>

            <QExpansionItem class="DateExpansion" label="Georeference Date">
              <QCardSection
                v-if="specimen?.georeference?.georeferenceDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.georeference.georeferenceDate.verbatim"
                  :disable="!props.enableEdit"
                  label="Verbatim"
                  type="text"
                />
                <!-- <QInput
                    v-model="specimen.georeference.georeferenceDate.timestamp"
                  :disable="!props.enableEdit"
                  type="date"
                  label="Timestamp"
                /> -->
                <QInput
                  v-model.number="specimen.georeference.georeferenceDate.year"
                  :disable="!props.enableEdit"
                  label="Year"
                  number
                  type="number"
                />
                <QInput
                  v-model="specimen.georeference.georeferenceDate.month"
                  :disable="!props.enableEdit"
                  label="month"
                  type="text"
                />
                <QInput
                  v-model.number="specimen.georeference.georeferenceDate.day"
                  :disable="!props.enableEdit"
                  label="day"
                  number
                  type="number"
                />
              </QCardSection>
            </QExpansionItem>
            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.locationRemarks"
                :disable="!props.enableEdit"
                autogrow
                class="col m-1"
                label="Location Remarks"
                type="textarea"
              />
            </div>
            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.footprintWkt"
                :disable="!props.enableEdit"
                autogrow
                class="col m-1"
                label="Footprint Well Known Type"
                type="textarea"
              />
              <q-input
                v-model="specimen.georeference.notes"
                :disable="!props.enableEdit"
                autogrow
                class="col m-1"
                label="Notes"
                type="textarea"
              />
            </div>
            <!-- <div>
              <q-item-label overline><b>Last Updated At </b></q-item-label>
              <q-item-label class="text-bold">
                {{ specimen?.secondary?.lastModified?.timestamp }}
              </q-item-label>
            </div> -->
          </q-card-section>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        :header-class="sectionHeaderClass"
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        default-opened
        expand-icon-toggle
        label="Secondary"
      >
        <q-card>
          <q-card-section v-if="specimen?.secondary" class="my-1">
            <div class="flex flex-row justify-evenly mb-1 gap-1">
              <q-input
                v-model="specimen.secondary.molt"
                :disable="!props.enableEdit"
                autogrow
                class="my-1 flex-grow"
                label="Molt"
                type="text"
              />
              <QInput
                v-model.number="specimen.secondary.weight"
                :disable="!props.enableEdit"
                autogrow
                class="flex-grow my-1"
                label="Weight"
                type="number"
              />
              <QInput
                v-model="specimen.secondary.weightUnits"
                :disable="!props.enableEdit"
                autogrow
                class="flex-grow my-1"
                label="Weight Units"
                type="text"
              />
            </div>

            <div class="flex flex-row justify-evenly mb-1 gap-1">
              <q-input
                v-model="specimen.secondary.condition"
                :disable="!props.enableEdit"
                autogrow
                class="my-1 flex-grow"
                label="Condition"
                type="textarea"
              />
              <q-input
                v-model="specimen.secondary.notes"
                :disable="!props.enableEdit"
                autogrow
                class="flex-grow my-1"
                label="Notes"
                type="textarea"
              />
            </div>
            <div class="flex flex-row justify-evenly mb-1 gap-1">
              <QSelect
                v-model="specimen.secondary.sex"
                :disable="!props.enableEdit"
                :options="secondarySexOptions"
                autogrow
                class="flex-grow my-1"
                label="Sex"
              />
              <QSelect
                v-model="specimen.secondary.age"
                :disable="!props.enableEdit"
                :options="secondaryAgeOptions"
                autogrow
                class="flex-grow my-1"
                label="Age"
              />
            </div>
          </q-card-section>
          <QCard>
            <QCardSection class="ml-4" label="Preparation">
              <div class="flex flex-row gap-2">
                <div class="font-bold">Preparations:</div>

                <div v-if="props.enableEdit" class="ml-auto">
                  <q-btn color="primary" label="New Preparation" push>
                    <q-popup-proxy class="">
                      <div class="flex flex-row gap-2 items-center p-2">
                        <QInput
                          v-model="newPreparationName"
                          :disable="!props.enableEdit"
                          class="flex-grow my-1"
                          label="New Preparation"
                          type="text"
                        />
                        <QBtn
                          class="h-2"
                          label="Add"
                          @click="addPreparation(newPreparationName)"
                        />
                      </div>
                    </q-popup-proxy>
                  </q-btn>
                </div>
              </div>

              <QCard
                v-for="(_, key) in specimen.secondary?.preparations"
                :key="key"
              >
                <QInput
                  v-model="specimen.secondary.preparations[key].verbatim"
                  :disable="!props.enableEdit"
                  :label="`Preparation: ${key}`"
                  class="flex-grow my-1"
                  type="text"
                >
                  <template v-if="props.enableEdit" #append>
                    <q-icon
                      class="cursor-pointer"
                      name="cancel"
                      @click="delete specimen?.secondary?.preparations[key]"
                    />
                  </template>
                </QInput>
              </QCard>
            </QCardSection>
          </QCard>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        :header-class="sectionHeaderClass"
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        default-opened
        expand-icon-toggle
        label="Loans"
      >
        <q-card>
          <QCardSection class="ml-4 flex flex-col gap-4">
            <q-btn
              v-if="props.enableEdit"
              class="bg-green-300 flex-grow"
              label="New Loans"
              push
            >
              <q-popup-proxy>
                <!-- <q-banner> -->
                <div class="flex flex-row gap-2 p-2 items-center">
                  <QInput
                    v-model="newLoanName"
                    :disable="!props.enableEdit"
                    class="flex-grow my-1"
                    label="New Loan"
                    type="text"
                  />
                  <QBtn class="h-2" label="Add" @click="addLoan(newLoanName)" />
                </div>
                <!-- </q-banner> -->
              </q-popup-proxy>
            </q-btn>
            <QCard
              v-for="(_, key) in specimen.loans"
              :key="key"
              class="flex flex-col gap-2 p-2"
            >
              <template v-if="specimen.loans[key]">
                <QChip :label="`Loan ID: ${key}`" class="flex-grow" />

                <div class="flex flex-row gap-4">
                  <QInput
                    v-model="specimen.loans[key].loanedTo"
                    :disable="!props.enableEdit"
                    class="flex-grow my-1"
                    label="Loaned To"
                    type="text"
                  />
                  <QInput
                    v-model="specimen.loans[key].loanedBy"
                    :disable="!props.enableEdit"
                    class="flex-grow my-1"
                    label="Loaned by"
                    type="text"
                  />
                </div>
                <QInput
                  v-model="specimen.loans[key].description"
                  :disable="!props.enableEdit"
                  class="flex-grow my-1"
                  label="Description"
                  type="textarea"
                />

                <QBtn
                  v-if="props.enableEdit"
                  class="max-h-4 bg-red-300"
                  label="Delete Loan"
                  @click="delete specimen.loans[key]"
                />
              </template>
            </QCard>
          </QCardSection>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        :header-class="sectionHeaderClass"
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        default-opened
        expand-icon-toggle
        label="Grants"
      >
        <QCard>
          <QCardSection class="ml-4 flex flex-col gap-4">
            <q-btn
              v-if="props.enableEdit"
              class="bg-green-300 flex-grow"
              label="New Grants"
              push
            >
              <q-popup-proxy>
                <!-- <q-banner> -->
                <div class="flex flex-row gap-2 p-2 items-center">
                  <QInput
                    v-model="newGrantName"
                    :disable="!props.enableEdit"
                    class="flex-grow my-1"
                    label="New Grant"
                    type="text"
                  />
                  <QBtn
                    class="h-2"
                    label="Add"
                    @click="addGrant(newGrantName)"
                  />
                </div>
                <!-- </q-banner> -->
              </q-popup-proxy>
            </q-btn>
            <QCard
              v-for="(_, key) in specimen.grants"
              :key="key"
              class="flex flex-col gap-2 p-2"
            >
              <template v-if="specimen.grants[key]">
                <QChip :label="`Grant ID: ${key}`" />
                <div class="flex flex-row gap-4">
                  <QInput
                    v-model="specimen.grants[key].grantedTo"
                    :disable="!props.enableEdit"
                    class="flex-grow my-1"
                    label="Grant To"
                    type="text"
                  />
                  <QInput
                    v-model="specimen.grants[key].grantedBy"
                    :disable="!props.enableEdit"
                    class="flex-grow my-1"
                    label="Granted by"
                    type="text"
                  />
                </div>
                <QInput
                  v-model="specimen.grants[key].description"
                  :disable="!props.enableEdit"
                  class="flex-grow my-1"
                  label="Description"
                  type="textarea"
                />

                <QBtn
                  v-if="props.enableEdit"
                  class="max-h-4 bg-red-300"
                  label="Delete Loan"
                  @click="delete specimen.grants[key]"
                />
              </template>
            </QCard>
          </QCardSection>
        </QCard>
      </q-expansion-item>
      <q-separator inset />

      <q-card-section v-if="showUpdater">
        <q-item-label overline>
          <b>Last Updater</b>
        </q-item-label>
        <q-item-label class="text-bold">
          {{ specimen.lastModified?.timestamp }}
        </q-item-label>
      </q-card-section>
      <!-- <q-separator inset /> -->

      <QCardSection>
        <slot name="Footer" />
      </QCardSection>
    </q-form>
  </QCard>
  <!-- <q-btn @click="toggleEdit"> hi </q-btn> -->
</template>

<style scoped>
.DateExpansion {
  @apply text-center  bg-gray-100 mx-4 my-2;
}

.DateSection {
  @apply flex flex-row gap-4 justify-evenly;
}
</style>
