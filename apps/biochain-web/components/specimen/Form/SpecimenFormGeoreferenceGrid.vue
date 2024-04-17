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
        <div
          class="grid grid-cols-1 content-center gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-3"
        >
          <FormKit
            type="text"
            name="continent"
            id="continent"
            label="Continent"
            outer-class="grid-col-span-3 min-w-20"
            :wrapper-class="{
              non: true,
            }"
          />

          <FormKit
            type="text"
            name="country"
            id="country"
            label="Country"
            class="grid-col-span-3"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="stateProvince"
            id="stateProvince"
            label="State Province"
            outer-class="grid-col-span-3 min-w-20"
          />
          <FormKit
            type="text"
            name="county"
            id="county"
            label="County"
            outer-class="grid-col-span-3 min-w-20"
          />
          <FormKit
            type="text"
            name="habitat"
            id="habitat"
            label="Habitat"
            outer-class="grid-col-span-3 min-w-20"
          />
          <FormKit
            type="text"
            name="locality"
            id="locality"
            label="Locality"
            outer-class="grid-col-span-3 min-w-20"
          />

          <UDivider class="col-span-full col-start-1" />

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
            type="number"
            name="coordinateUncertaintyInMeters"
            id="coordinateUncertaintyInMeters"
            label="Coordinate Uncertainty In Meters"
            outer-class="min-w-20"
          />

          <UDivider class="col-span-full col-start-1" />

          <FormKit
            type="text"
            name="georeferenceBy"
            id="georeferenceBy"
            label="Georeference By"
            outer-class="min-w-20"
          />

          <FormKit
            type="textarea"
            name="geodeticDatum"
            id="geodeticDatum"
            label="Geodetic Datum"
            auto-height
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="georeferenceProtocol"
            id="georeferenceProtocol"
            label="Georeference Protocol"
            outer-class="min-w-20"
          />

          <SpecimenFormDate name="georeferenceDate" />

          <UDivider class="col-span-full col-start-1" />

          <FormKit
            type="textarea"
            auto-height
            name="georeferenceRemarks"
            label="Georeference Remarks"
          />

          <FormKit
            type="textarea"
            auto-height
            name="footprintWkt"
            label="Footprint Well Known Type"
          />
          <FormKit
            auto-height
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
