<template>
  <div class="max-w-4xl">
    <QCard>
      <QCardSection :v-if="props.showHeader" class="bg-blue-400">
        <slot name="HeaderBar" />
      </QCardSection>
      <QCardSection class="bg-blue-400">
        <div class="row">
          <div v-if="specimen?.taxon" class="text-2xl text-white">
            <span v-if="specimen?.taxon?.genus">{{
              specimen?.taxon?.genus
            }}</span>
            {{}}
            <span v-if="specimen?.taxon?.species">{{
              specimen?.taxon?.species
            }}</span>
          </div>
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
      </QCardSection>

      <QExpansionItem
        label="Taxon"
        :icon="'ti-pencil-alt'"
        expand-separator
        header-class="text-center text-xl"
        :default-opened="true"
      >
        <QCardSection class="flex flex-row gap-4">
          <FormKit type="group" name="taxon">
            <FormKit type="text" name="kingdom" label="Kingdom" />
            <FormKit type="text" name="phylum" label="Phylum" />
            <FormKit type="text" name="class" label="Class" />
            <FormKit type="text" name="order" label="Order" />
            <FormKit type="text" name="family" label="Family" />
            <FormKit type="text" name="genus" label="Genus" />
            <FormKit type="text" name="species" label="Species" />
            <FormKit type="text" name="subspecies" label="Subspecies" />
          </FormKit>
        </QCardSection>
      </QExpansionItem>

      <QSeparator />

      <QExpansionItem :default-opened="true" label="Primary" expand-separator>
        <QCardSection class="flex flex-col gap-4">
          <FormKit type="group" name="primary">
            <div class="flex flex-row gap-4">
              <FormKit
                type="text"
                name="accessionNumber"
                label="Accession Number"
                placeholder=""
                validation=""
              />
              <FormKit
                type="text"
                name="tissueNumber"
                label="Tissue Number"
                placeholder=""
                validation=""
              />
            </div>

            <div class="flex flex-row gap-4 items-center">
              <FormKit
                type="text"
                name="catalogNumber"
                label="Catalog Number"
                placeholder=""
                validation=""
              />
              <FormKit
                type="text"
                name="cataloger"
                label="Cataloger"
                placeholder=""
                validation=""
              />
              <QExpansionItem label="Catalog Date" class="flex-grow-1">
                <QCardSection class="flex flex-row gap-4">
                  <FormKit type="group" name="catalogDate" label="catalogDate">
                    <FormKit type="text" name="verbatim" label="Verbatim" />
                    <FormKit type="date" name="timestamp" label="Timestamp" />
                    <FormKit type="number" name="year" label="Year" />
                    <FormKit type="text" name="month" label="Month" />
                    <FormKit type="number" name="day" label="Day" />
                  </FormKit>
                </QCardSection>
              </QExpansionItem>
            </div>
            <div class="flex flex-row gap-4 items-center">
              <FormKit
                type="text"
                name="collector"
                label="Collector"
                placeholder=""
                validation=""
              />
              <FormKit
                type="text"
                name="fieldNumber"
                label="Field Number"
                placeholder=""
                validation=""
              />
              <QExpansionItem label="Field Date" class="outline-1">
                <QCardSection class="flex flex-row gap-4">
                  <FormKit type="group" name="fieldDate" label="Field Date">
                    <FormKit
                      type="text"
                      name="verbatim"
                      label="Field Date: Verbatim"
                    />
                    <FormKit
                      type="date"
                      name="timestamp"
                      label="Field Date: Timestamp"
                    />
                    <FormKit type="number" name="year" label="Year" />
                    <FormKit type="text" name="month" label="Month" />
                    <FormKit type="number" name="day" label="Day" />
                  </FormKit>
                </QCardSection>
              </QExpansionItem>
            </div>
            <!-- <span> -->

            <!-- </span> -->
            <div class="flex flex-row gap-4 items-center">
              <FormKit
                type="text"
                name="determiner"
                label="Determiner"
                placeholder=""
                validation=""
              />
              <FormKit
                type="text"
                name="determinedReason"
                label="Determined Reason"
                placeholder=""
                validation=""
              />

              <QExpansionItem label="Determined Date" class="flex-grow-1">
                <QCardSection class="flex flex-row gap-4">
                  <FormKit
                    type="group"
                    name="determinedDate"
                    label="determinedDate"
                  >
                    <FormKit type="text" name="verbatim" label="Verbatim" />
                    <FormKit type="date" name="timestamp" label="Timestamp" />
                    <FormKit type="number" name="year" label="Year" />
                    <FormKit type="text" name="month" label="Month" />
                    <FormKit type="number" name="day" label="Day" />
                  </FormKit>
                </QCardSection>
              </QExpansionItem>
            </div>

            <QExpansionItem label="Original Date" class="flex-grow-1">
              <QCardSection class="flex flex-row gap-4">
                <FormKit type="group" name="originalDate" label="originalDate">
                  <FormKit type="text" name="verbatim" label="Verbatim" />
                  <FormKit type="date" name="timestamp" label="Timestamp" />
                  <FormKit type="number" name="year" label="Year" />
                  <FormKit type="text" name="month" label="month" />
                  <FormKit type="number" name="day" label="day" />
                </FormKit>
              </QCardSection>
            </QExpansionItem>
          </FormKit>
          <QSeparator />
        </QCardSection>
      </QExpansionItem>
      <QExpansionItem
        label="Secondary"
        :expand-icon="props.showEditOptions ? 'edit' : 'arrow_drop_down'"
        :default-opened="true"
        expand-separator
      >
        <QCardSection class="flex flex-row gap-4">
          <FormKit type="group" name="secondary">
            <FormKit
              type="select"
              name="sex"
              label="sex"
              :options="[
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
              ]"
            />
            <FormKit
              type="select"
              name="age"
              label="Age"
              :options="[
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
              ]"
            />

            <FormKit type="number" name="weight" label="Weight" />
            <FormKit type="text" name="weightUnits" label="Weight Units" />
            <FormKit type="text" name="preparation" label="Preparation" />
            <FormKit type="text" name="condition" label="Condition" />
            <FormKit type="text" name="molt" label="Molt" />
            <FormKit
              type="textarea"
              name="notes"
              label="Notes"
              class="flex-grow-1"
            />
          </FormKit>
          <QSeparator />
        </QCardSection>
      </QExpansionItem>

      <QExpansionItem
        label="Georeference"
        :expand-icon="props.showEditOptions ? 'edit' : 'arrow_drop_down'"
        expand-separator
        :default-opened="true"
      >
        <QCardSection class="flex flex-row gap-4">
          <FormKit type="group" name="georeference">
            <FormKit type="text" name="continent" label="Continent" />
            <FormKit type="text" name="country" label="Country" />
            <FormKit type="text" name="stateProvince" label="State/Province" />
            <FormKit type="text" name="county" label="County" />
            <FormKit type="text" name="locality" label="Locality" />
            <FormKit type="text" name="longitude" label="Longitude" />
            <FormKit type="text" name="latitude" label="Latitude" />
            <FormKit type="text" name="habitat" label="Habitat" />
            <FormKit
              type="text"
              name="locationRemarks"
              label="Location Remarks"
            />
            <FormKit
              type="number"
              name="coordinateUncertaintyInMeters"
              label="Coordinate Uncertainty In Meters"
            />
            <FormKit
              type="text"
              name="georeferenceBy"
              label="Georeference By"
            />
            <FormKit
              type="text"
              name="georeferenceProtocol"
              label="georeferenceProtocol"
            />
            <FormKit type="text" name="geodeticDatum" label="Geodetic Datum" />

            <FormKit
              type="textarea"
              name="footprintWkt"
              label="Footprint Wkt"
            />
            <FormKit type="textarea" name="notes" label="notes" />

            <QExpansionItem label="Georeference Date" class="flex-grow-1">
              <QCardSection class="flex flex-row gap-4">
                <FormKit
                  type="group"
                  name="georeferenceDate"
                  label="Field Date"
                >
                  <FormKit type="text" name="verbatim" label="Verbatim" />
                  <FormKit type="date" name="timestamp" label="Timestamp" />
                  <FormKit type="number" name="year" label="Year" />
                  <FormKit type="text" name="month" label="month" />
                  <FormKit type="number" name="day" label="day" />
                </FormKit>
              </QCardSection>
            </QExpansionItem>
          </FormKit>
          <QSeparator />
        </QCardSection>
      </QExpansionItem>

      <QExpansionItem
        label="Loans"
        :expand-icon="props.showEditOptions ? 'edit' : 'arrow_drop_down'"
        expand-separator
        :default-opened="true"
      >
        <QCardSection class="flex flex-row gap-4">
          <FormKit type="group" name="loans"> </FormKit>
          <QSeparator />
        </QCardSection>
      </QExpansionItem>

      <QExpansionItem
        label="Grants"
        :expand-icon="props.showEditOptions ? 'edit' : 'arrow_drop_down'"
        expand-separator
        :default-opened="true"
      >
        <QCardSection class="flex flex-row gap-4">
          <FormKit type="group" name="grants"> </FormKit>
          <QSeparator />
        </QCardSection>
      </QExpansionItem>
      <FormKit type="submit" label="Register" />
    </QCard>

    asdf
  </div>
</template>

<script lang="ts" setup>
import type { PlainMessage } from "@bufbuild/protobuf";
import { ccbio } from "saacs-es";

type rawSpecimen = PlainMessage<ccbio.Specimen>;

// const sp = reactive(new ccbio.Specimen());

// const Specimen = pb.ccbio.Specimen;
// type Specimen = typeof pb.ccbio.Specimen;
export interface Props {
  // specimen: ccbio.Specimen;
  enableEdit?: boolean;
  showUpdater?: boolean;
  showHeader?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  editable: false,
  enableEdit: true,
  showUpdater: true,
  showHeader: true,
  showEditOptions: true,
});

function submitHandler(value: ccbio.Specimen) {
  console.log(value);
  console.log(new ccbio.Specimen(value));
}

const specimen = defineModel<rawSpecimen>("specimen", {
  default: () => new ccbio.Specimen(),
});

// const emit = defineEmits(["update:specimen"]);

// const specimen = useVModel(props, "specimen", emit);
//    ^?

// const dateMask = "M/D/YYYY";

// const getJson = (obj: schema.Specimen) => {
//   return schema.Specimen.toJSON(obj);
// };
</script>

<style></style>
