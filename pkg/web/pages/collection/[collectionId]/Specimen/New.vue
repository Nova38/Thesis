<script lang="ts" setup>
import { randomUUID } from 'uncrypto'
import { ccbio } from '~/lib'

//    ^?

const specimen = ref(
  new ccbio.Specimen({
    collectionId: useRouteCollectionId,
    georeference: {
      georeferenceDate: {},
    },
    grants: {},
    images: {},
    loans: {},
    primary: {
      catalogDate: {},
      determinedDate: {},
      fieldDate: {},
      originalDate: {},
    },
    secondary: {
      preparations: {
        t: {},
      },
    },
    specimenId: randomUUID(),
    taxon: {},
  }),
)

async function submitHandler() {
  // console.log("submitHandler", value);

  const response = await useCreateSpecimen(specimen.value)
  console.log('response', response)

  useRouter().push(
    `/collection/${useRouteCollectionId}/specimen/View-${specimen.value.specimenId}`,
  )
}
</script>

<template>
  <div>
    <div>
      <SpecimenForm :specimen="specimen">
        <template #Footer>
          <div>
            <QBtn
              label="Submit"
              @click="submitHandler"
            />
          </div>
        </template>
      </SpecimenForm>

    <!-- <pre wrap>{{ value }}</pre> -->
    </div>
    <div>
      <pre wrap>{{ specimen }}</pre>
    </div>
  </div>
</template>

<style></style>
