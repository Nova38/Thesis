<script lang="ts" setup></script>

<template>
  <div>
    <PFieldset
      legend="Georeference"
      :toggleable="true"
    >
      <FormKit
        id="georeference"
        type="group"
        name="georeference"
      >
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            id="continent"
            type="text"
            name="continent"
            label="Continent"
            outer-class="min-w-20"
          />
          <FormKit
            id="country"
            type="text"
            name="country"
            label="Country"
            outer-class="min-w-20"
          />
          <FormKit
            id="stateProvince"
            type="text"
            name="stateProvince"
            label="State Province"
            outer-class="min-w-20"
          />
          <FormKit
            id="county"
            type="text"
            name="county"
            label="County"
            outer-class="min-w-20"
          />
          <FormKit
            id="habitat"
            type="text"
            name="habitat"
            label="Habitat"
            outer-class="min-w-20"
          />
          <FormKit
            id="locality"
            type="text"
            name="locality"
            label="Locality"
            outer-class="min-w-20"
          />
        </div>

        <UDivider />
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="number"
            step="any"
            number
            name="longitude"
            validate="number"
            label="longitude"
          />
          <FormKit
            type="number"
            step="any"
            name="latitude"
            validate="number"
            number
            label="latitude"
            outer-class="min-w-20"
          />

          <FormKit
            id="coordinateUncertaintyInMeters"
            type="number"
            name="coordinateUncertaintyInMeters"
            label="Coordinate Uncertainty In Meters"
            outer-class="min-w-20"
          />
        </div>
        <UDivider />

        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            id="georeferenceBy"
            type="text"
            name="georeferenceBy"
            label="Georeference By"
            outer-class="min-w-20"
          />

          <FormKit
            id="geodeticDatum"
            type="text"
            name="geodeticDatum"
            label="Geodetic Datum"
            outer-class="min-w-20"
          />
          <FormKit
            id="georeferenceProtocol"
            type="text"
            name="georeferenceProtocol"
            label="Georeference Protocol"
            outer-class="min-w-20"
          />
        </div>

        <SpecimenFormDate name="georeferenceDate" />

        <UDivider class="py-2" />

        <div class="inline-flex w-full flex-wrap gap-2">
          <FormKit
            type="textarea"
            name="georeferenceRemarks"
            label="Georeference Remarks"
          />

          <FormKit
            type="textarea"
            name="footprintWkt"
            label="Footprint Well Known Type"
          />
          <FormKit
            type="textarea"
            name="notes"
            label="Notes"
          />
        </div>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
