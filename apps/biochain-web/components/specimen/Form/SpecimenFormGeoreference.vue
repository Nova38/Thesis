<script lang="ts" setup>

const collapsed = defineModel<boolean>('collapsed', {
  default: false,
})

</script>

<template>
  <div>
    <PFieldset
      legend="Georeference"
      toggleable
      :collapsed
    >
      <FormKit
        id="georeference"
        type="group"
        name="georeference"
        #default="{value}"
      >
        <div
          class="grid grid-cols-1 content-center gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-3"
        >
          <FormKit
            id="continent"
            type="text"
            name="continent"
            label="Continent"
            outer-class="grid-col-span-3 min-w-20"
            :wrapper-class="{
              non: true,
            }"
          />

          <FormKit
            id="country"
            type="text"
            name="country"
            label="Country"
            class="grid-col-span-3"
            outer-class="min-w-20"
          />
          <FormKit
            id="stateProvince"
            type="text"
            name="stateProvince"
            label="State Province"
            outer-class="grid-col-span-3 min-w-20"
          />
          <FormKit
            id="county"
            type="text"
            name="county"
            label="County"
            outer-class="grid-col-span-3 min-w-20"
          />
          <FormKit
            id="habitat"
            type="text"
            name="habitat"
            label="Habitat"
            outer-class="grid-col-span-3 min-w-20"
          />
          <FormKit
            id="locality"
            type="text"
            name="locality"
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
            id="coordinateUncertaintyInMeters"
            type="number"
            name="coordinateUncertaintyInMeters"
            label="Coordinate Uncertainty In Meters"
            outer-class="min-w-20"
          />

          <UDivider class="col-span-full col-start-1" />

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
            auto-height
            outer-class="min-w-20"
          />
          <FormKit
            id="georeferenceProtocol"
            type="textarea"
            name="georeferenceProtocol"
            label="Georeference Protocol"
            outer-class="min-w-50"
          />

          <SpecimenFormDate v-if="value?.georeferenceDate" name="georeferenceDate" />

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
