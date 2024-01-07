<template>
  <div>
    <div>
      <SpecimenForm :specimen="specimen">
        <template #Footer>
          <div>
            <QBtn label="Submit" @click="submitHandler" />
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

<script lang="ts" setup>
import { ccbio } from "saacs-es";
import { randomUUID } from "uncrypto";

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
);

const submitHandler = async () => {
  // console.log("submitHandler", value);

  const response = await useCreateSpecimen(specimen.value);
  console.log("response", response);

  useRouter().push(
    `/collection/${useRouteCollectionId}/specimen/View-${specimen.value.specimenId}`,
  );
};
</script>

<style></style>
