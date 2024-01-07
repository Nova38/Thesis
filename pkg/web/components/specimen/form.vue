<script lang="ts" setup>
import { ccbio } from "saacs-es";

const sectionHeaderClass = "text-center section-header bg-gray-200";

const props = defineProps({
  showEditOptions: {
    type: Boolean,
    default: true,
  },
  showHeader: {
    type: Boolean,
    default: false,
  },
  showUpdater: {
    type: Boolean,
    default: true,
  },
  showEmpty: {
    type: Boolean,
    default: true,
  },
  startOpen: {
    type: Boolean,
    default: true,
  },
  enableEdit: {
    type: Boolean,
    default: true,
  },
  headerColor: {
    type: String,
    default: toModeColor("view"),
  },
});

const specimen = defineModel("specimen", {
  default: () => new ccbio.Specimen(),
});

function addPreparation(name: string) {
  if (!specimen?.value) return;
  if (!specimen?.value?.secondary) {
    specimen.value.secondary = new ccbio.Specimen_Secondary();
  }
  if (!specimen?.value?.secondary?.preparations) {
    specimen.value.secondary.preparations = {};
  }
  specimen.value.secondary.preparations[name] =
    new ccbio.Specimen_Secondary_Preparation();
}

function addLoan(name: string) {
  if (!specimen?.value) return;
  if (!specimen?.value?.loans) {
    specimen.value.loans = {};
  }
  specimen.value.loans[name] = new ccbio.Specimen_Loan({
    id: name,
  });
}

function addGrant(name: string) {
  if (!specimen?.value) return;
  if (!specimen?.value?.grants) {
    specimen.value.grants = {};
  }
  specimen.value.grants[name] = new ccbio.Specimen_Grant({});
}
const newPreparationName = ref("");
const newLoanName = ref("");
const newGrantName = ref("");
// ^?

const secondarySexOptions = [
  {
    label: "",
    value: ccbio.Specimen_Secondary_SEX.SEX_UNDEFINED,
  },
  {
    label: "Unknown",
    value: ccbio.Specimen_Secondary_SEX.SEX_UNKNOWN,
  },
  {
    label: "Atypical",
    value: ccbio.Specimen_Secondary_SEX.SEX_ATYPICAL,
  },
  {
    label: "Male",
    value: ccbio.Specimen_Secondary_SEX.SEX_MALE,
  },
  {
    label: "Female",
    value: ccbio.Specimen_Secondary_SEX.SEX_FEMALE,
  },
];

const secondaryAgeOptions = [
  {
    label: "",
    value: ccbio.Specimen_Secondary_AGE.AGE_UNDEFINED,
  },
  {
    label: "Unknown",
    value: ccbio.Specimen_Secondary_AGE.AGE_UNKNOWN,
  },
  {
    label: "Nest",
    value: ccbio.Specimen_Secondary_AGE.AGE_NEST,
  },
  {
    label: "Embryo Egg",
    value: ccbio.Specimen_Secondary_AGE.AGE_EMBRYO_EGG,
  },
  {
    label: "Chick SubAdult",
    value: ccbio.Specimen_Secondary_AGE.AGE_CHICK_SUBADULT,
  },
  {
    label: "Adult",
    value: ccbio.Specimen_Secondary_AGE.AGE_ADULT,
  },
  {
    label: "Age Contingent",
    value: ccbio.Specimen_Secondary_AGE.AGE_CONTINGENT,
  },
];
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
              color="accent"
              square
              class="text-white"
              :label="'Catalog Number: ' + specimen?.primary?.catalogNumber"
            />
            <q-chip
              v-if="specimen.primary.tissueNumber"
              color="accent"
              class="text-white"
              square
              :label="'Tissue Number: ' + specimen?.primary.tissueNumber"
            />
          </div>
        </div>
      </div>
    </QCardSection>
    <div class="flex flex-row" :class="headerColor">
      <q-chip
        v-if="specimen.collectionId"
        color="primary"
        class="text-white flex-grow"
        square
        :label="'Collection ID: ' + specimen.collectionId"
      />
      <q-chip
        v-if="specimen.specimenId"
        color="primary"
        class="text-white flex-grows"
        square
        :label="'Specimen ID: ' + specimen.specimenId"
      />
    </div>

    <q-form>
      <QExpansionItem
        expand-icon-toggle
        default-opened
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        label="Taxon"
        :header-class="sectionHeaderClass"
      >
        <QCard>
          <QCardSection class="flex flex-row gap-1 justify-evenly mb-1">
            <template v-for="(_, key) in specimen.taxon" :key="key">
              <QInput
                v-if="specimen.taxon && typeof specimen.taxon[key] === 'string'"
                v-model="specimen.taxon[key]"
                class="flex-grow my-1"
                :label="key"
                :disable="!props.enableEdit"
              />
            </template>
          </QCardSection>
          <div class="m-4">
            <q-item-label overline><b>Last Updated At</b></q-item-label>
            <q-item-label class="text-bold">
              {{ specimen?.primary?.lastModified?.timestamp }}
            </q-item-label>
          </div>
        </QCard>
      </QExpansionItem>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        label="Primary"
        :header-class="sectionHeaderClass"
      >
        <div v-if="specimen?.primary">
          <q-card-section class="my-1">
            <!-- <div class="text-h6">Overline</div> -->
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.tissueNumber"
                label="Tissue Number"
                class="m-1 flex-grow"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.primary.accessionNumber"
                class="m-1 flex-grow"
                label="Accession Number"
                :disable="!props.enableEdit"
              />
            </div>
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.cataloger"
                class="flex-grow m-1"
                label="Cataloger"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.primary.catalogNumber"
                class="flex-grow m-1"
                label="Catalog Number"
                :disable="!props.enableEdit"
              />
            </div>
            <QExpansionItem label="Catalog Date" class="DateExpansion">
              <QCardSection
                v-if="specimen?.primary?.catalogDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.catalogDate.verbatim"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Catalog Date: Verbatim"
                  class="flex-grow"
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
                  type="number"
                  label="Catalog Date: Year"
                  number
                />
                <QInput
                  v-model="specimen.primary.catalogDate.month"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Catalog Date: Month"
                />
                <QInput
                  v-model.number="specimen.primary.catalogDate.day"
                  :disable="!props.enableEdit"
                  type="number"
                  label="Catalog Date: Day"
                  number
                />
              </QCardSection>
            </QExpansionItem>
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.determiner"
                class="flex-grow m-1"
                label="Determiner"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.primary.determinedReason"
                label="Determined Reason"
                class="flex-grow m-1"
                :disable="!props.enableEdit"
              />
            </div>
            <QExpansionItem label="Determined Date" class="DateExpansion mx-4">
              <QCardSection
                v-if="specimen?.primary?.determinedDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.determinedDate.verbatim"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Determined Date: Verbatim"
                  class="flex-grow"
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
                  type="number"
                  label="Determined Date: Year"
                  number
                />
                <QInput
                  v-model="specimen.primary.determinedDate.month"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Determined Date: Month"
                />
                <QInput
                  v-model.number="specimen.primary.determinedDate.day"
                  :disable="!props.enableEdit"
                  type="number"
                  label="Determined Date: Day"
                  number
                />
              </QCardSection>
            </QExpansionItem>
            <div class="flex flex-row gap-1 justify-evenly mb-1">
              <q-input
                v-model="specimen.primary.collector"
                label="Collector"
                class="flex-grow m-1"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.primary.fieldNumber"
                class="flex-grow m-1"
                label="Field Number"
                :disable="!props.enableEdit"
              />
            </div>

            <QExpansionItem label="Field Date" class="DateExpansion">
              <QCardSection
                v-if="specimen?.primary?.fieldDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.fieldDate.verbatim"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Field Date: Verbatim"
                  class="flex-grow"
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
                  type="number"
                  label="Field Date: Year"
                  number
                />
                <QInput
                  v-model="specimen.primary.fieldDate.month"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Field Date: Month"
                />
                <QInput
                  v-model.number="specimen.primary.fieldDate.day"
                  :disable="!props.enableEdit"
                  type="number"
                  label="Field Date: Day"
                  number
                />
              </QCardSection>
            </QExpansionItem>
            <QExpansionItem
              v-if="specimen?.primary?.originalDate"
              label="Original Date"
              class="DateExpansion"
            >
              <QCardSection
                v-if="specimen?.georeference?.georeferenceDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.primary.originalDate.verbatim"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Original Date: Verbatim"
                  class="flex-grow"
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
                  type="number"
                  label="Original Date: Year"
                  number
                />
                <QInput
                  v-model="specimen.primary.originalDate.month"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Original Date: Month"
                />
                <QInput
                  v-model.number="specimen.primary.originalDate.day"
                  :disable="!props.enableEdit"
                  type="number"
                  label="Original Date: Day"
                  number
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
        expand-icon-toggle
        default-opened
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        label="Georeference"
        :header-class="sectionHeaderClass"
      >
        <q-card>
          <q-card-section v-if="specimen?.georeference" class="my-1">
            <div class="row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.continent"
                label="Continent"
                class="col m-1"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.country"
                label="Country"
                class="col m-1"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.stateProvince"
                label="State/Province"
                class="col m-1"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.county"
                label="County"
                class="col m-1"
                :disable="!props.enableEdit"
              />
            </div>

            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.habitat"
                class="col m-1"
                label="Habitat"
                type="textarea"
                autogrow
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.locality"
                class="col m-1"
                label="Locality"
                :disable="!props.enableEdit"
              />
            </div>
            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model.number="specimen.georeference.longitude"
                class="col m-1"
                type="number"
                label="Longitude"
                :disable="!props.enableEdit"
              />

              <q-input
                v-model.number="specimen.georeference.latitude"
                class="col m-1"
                label="Latitude"
                type="number"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.coordinateUncertaintyInMeters"
                class="col m-1"
                label="Coordinate Uncertainty In Meters"
                :disable="!props.enableEdit"
              />
            </div>
            <div class="row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.georeferenceBy"
                class="col m-1"
                label="Georeference By"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.georeferenceProtocol"
                class="col m-1"
                label="Georeference Protocol"
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.geodeticDatum"
                class="col m-1"
                label="Geodetic Datum"
                :disable="!props.enableEdit"
              />
            </div>

            <QExpansionItem label="Georeference Date" class="DateExpansion">
              <QCardSection
                v-if="specimen?.georeference?.georeferenceDate"
                class="DateSection"
              >
                <QInput
                  v-model="specimen.georeference.georeferenceDate.verbatim"
                  :disable="!props.enableEdit"
                  type="text"
                  label="Verbatim"
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
                  type="number"
                  label="Year"
                  number
                />
                <QInput
                  v-model="specimen.georeference.georeferenceDate.month"
                  :disable="!props.enableEdit"
                  type="text"
                  label="month"
                />
                <QInput
                  v-model.number="specimen.georeference.georeferenceDate.day"
                  :disable="!props.enableEdit"
                  type="number"
                  label="day"
                  number
                />
              </QCardSection>
            </QExpansionItem>
            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.locationRemarks"
                class="col m-1"
                label="Location Remarks"
                type="textarea"
                autogrow
                :disable="!props.enableEdit"
              />
            </div>
            <div class="flex flex-row justify-evenly mb-1">
              <q-input
                v-model="specimen.georeference.footprintWkt"
                class="col m-1"
                label="Footprint Well Known Type"
                type="textarea"
                autogrow
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.georeference.notes"
                class="col m-1"
                label="Notes"
                type="textarea"
                autogrow
                :disable="!props.enableEdit"
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
        expand-icon-toggle
        default-opened
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        label="Secondary"
        :header-class="sectionHeaderClass"
      >
        <q-card>
          <q-card-section v-if="specimen?.secondary" class="my-1">
            <div class="flex flex-row justify-evenly mb-1 gap-1">
              <q-input
                v-model="specimen.secondary.molt"
                class="my-1 flex-grow"
                label="Molt"
                type="text"
                autogrow
                :disable="!props.enableEdit"
              />
              <QInput
                v-model.number="specimen.secondary.weight"
                class="flex-grow my-1"
                label="Weight"
                type="number"
                autogrow
                :disable="!props.enableEdit"
              />
              <QInput
                v-model="specimen.secondary.weightUnits"
                label="Weight Units"
                type="text"
                class="flex-grow my-1"
                autogrow
                :disable="!props.enableEdit"
              />
            </div>

            <div class="flex flex-row justify-evenly mb-1 gap-1">
              <q-input
                v-model="specimen.secondary.condition"
                class="my-1 flex-grow"
                label="Condition"
                type="textarea"
                autogrow
                :disable="!props.enableEdit"
              />
              <q-input
                v-model="specimen.secondary.notes"
                class="flex-grow my-1"
                label="Notes"
                type="textarea"
                autogrow
                :disable="!props.enableEdit"
              />
            </div>
            <div class="flex flex-row justify-evenly mb-1 gap-1">
              <QSelect
                v-model="specimen.secondary.sex"
                class="flex-grow my-1"
                label="Sex"
                :options="secondarySexOptions"
                autogrow
                :disable="!props.enableEdit"
              />
              <QSelect
                v-model="specimen.secondary.age"
                class="flex-grow my-1"
                label="Age"
                :options="secondaryAgeOptions"
                autogrow
                :disable="!props.enableEdit"
              />
            </div>
          </q-card-section>
          <QCard>
            <QCardSection label="Preparation" class="ml-4">
              <div class="flex flex-row gap-2">
                <div class="font-bold">Preparations:</div>

                <div v-if="props.enableEdit" class="ml-auto">
                  <q-btn push color="primary" label="New Preparation">
                    <q-popup-proxy class="">
                      <div class="flex flex-row gap-2 items-center p-2">
                        <QInput
                          v-model="newPreparationName"
                          label="New Preparation"
                          type="text"
                          class="flex-grow my-1"
                          :disable="!props.enableEdit"
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
                  type="text"
                  class="flex-grow my-1"
                  :label="`Preparation: ${key}`"
                  :disable="!props.enableEdit"
                >
                  <template v-if="props.enableEdit" #append>
                    <q-icon
                      name="cancel"
                      class="cursor-pointer"
                      @click="delete specimen.secondary.preparations[key]"
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
        expand-icon-toggle
        default-opened
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        label="Loans"
        :header-class="sectionHeaderClass"
      >
        <q-card>
          <QCardSection class="ml-4 flex flex-col gap-4">
            <q-btn
              v-if="props.enableEdit"
              push
              label="New Loans"
              class="bg-green-300 flex-grow"
            >
              <q-popup-proxy>
                <!-- <q-banner> -->
                <div class="flex flex-row gap-2 p-2 items-center">
                  <QInput
                    v-model="newLoanName"
                    label="New Loan"
                    type="text"
                    class="flex-grow my-1"
                    :disable="!props.enableEdit"
                  />
                  <QBtn label="Add" class="h-2" @click="addLoan(newLoanName)" />
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
                    label="Loaned To"
                    type="text"
                    class="flex-grow my-1"
                    :disable="!props.enableEdit"
                  />
                  <QInput
                    v-model="specimen.loans[key].loanedBy"
                    label="Loaned by"
                    type="text"
                    class="flex-grow my-1"
                    :disable="!props.enableEdit"
                  />
                </div>
                <QInput
                  v-model="specimen.loans[key].description"
                  label="Description"
                  type="textarea"
                  class="flex-grow my-1"
                  :disable="!props.enableEdit"
                />

                <QBtn
                  v-if="props.enableEdit"
                  class="max-h-4 bg-red-300"
                  label="Delete Loan"
                  @click="delete specimen.loans[key]"
                >
                </QBtn>
              </template>
            </QCard>
          </QCardSection>
        </q-card>
      </q-expansion-item>

      <q-separator inset />

      <q-expansion-item
        expand-icon-toggle
        default-opened
        :icon="!props.enableEdit ? 'ti-lock' : 'ti-pencil-alt'"
        label="Grants"
        :header-class="sectionHeaderClass"
      >
        <QCard>
          <QCardSection class="ml-4 flex flex-col gap-4">
            <q-btn
              v-if="props.enableEdit"
              push
              class="bg-green-300 flex-grow"
              label="New Grants"
            >
              <q-popup-proxy>
                <!-- <q-banner> -->
                <div class="flex flex-row gap-2 p-2 items-center">
                  <QInput
                    v-model="newGrantName"
                    label="New Grant"
                    type="text"
                    class="flex-grow my-1"
                    :disable="!props.enableEdit"
                  />
                  <QBtn
                    label="Add"
                    class="h-2"
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
                    label="Grant To"
                    type="text"
                    class="flex-grow my-1"
                    :disable="!props.enableEdit"
                  />
                  <QInput
                    v-model="specimen.grants[key].grantedBy"
                    label="Granted by"
                    type="text"
                    class="flex-grow my-1"
                    :disable="!props.enableEdit"
                  />
                </div>
                <QInput
                  v-model="specimen.grants[key].description"
                  label="Description"
                  type="textarea"
                  class="flex-grow my-1"
                  :disable="!props.enableEdit"
                />

                <QBtn
                  v-if="props.enableEdit"
                  class="max-h-4 bg-red-300"
                  label="Delete Loan"
                  @click="delete specimen.grants[key]"
                >
                </QBtn>
              </template>
            </QCard>
          </QCardSection>
        </QCard>
      </q-expansion-item>
      <q-separator inset />

      <q-card-section v-if="showUpdater">
        <q-item-label overline><b>Last Updater</b></q-item-label>
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
