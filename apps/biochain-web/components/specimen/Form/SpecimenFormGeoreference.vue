<script lang="ts" setup></script>

<template>
  <div>
    <PFieldset
      legend="Georeference"
      :toggleable="true"
    >
      <FormKit
        type="group"
        id="georeference"
        name="georeference"
      >
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            name="continent"
            id="continent"
            label="Continent"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="country"
            id="country"
            label="Country"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="stateProvince"
            id="stateProvince"
            label="State Province"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="county"
            id="county"
            label="County"
            outer-class="min-w-20"
          />
        </div>
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            name="habitat"
            id="habitat"
            label="Habitat"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="locality"
            id="locality"
            label="Locality"
            outer-class="min-w-20"
          />
        </div>

        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            number
            name="longitude"
            validate="number"
            label="longitude"
          />
          <FormKit
            type="text"
            name="latitude"
            validate="number"
            number
            label="latitude"
            outer-class="min-w-20"
          />

          <FormKit
            type="number"
            name="coordinateUncertaintyInMeters"
            id="coordinateUncertaintyInMeters"
            label="Coordinate Uncertainty In Meters"
            outer-class="min-w-20"
          />
        </div>

        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            name="georeferenceBy"
            id="georeferenceBy"
            label="Georeference By"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="georeferenceProtocol"
            id="georeferenceProtocol"
            label="Georeference Protocol"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="geodeticDatum"
            id="geodeticDatum"
            label="Geodetic Datum"
            outer-class="min-w-20"
          />
        </div>

        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            type="text"
            name="georeferenceBy"
            id="georeferenceBy"
            label="Georeference By"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="georeferenceProtocol"
            id="georeferenceProtocol"
            label="Georeference Protocol"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="geodeticDatum"
            id="geodeticDatum"
            label="Geodetic Datum"
            outer-class="min-w-20"
          />
        </div>

        <SpecimenFormDate name="georeferenceDate" />
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
