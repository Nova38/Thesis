<template>
  <div>
    <div>
      <FormKit
        id="SpecimenForm"
        v-slot="{ value }"
        v-model="specimen"
        type="form"
        submit-label="Login"
        @submit="submitHandler"
      >
        <SpecimenForm :specimen="value" />
        <pre wrap>{{ value }}</pre>
        <!-- <pre wrap>{{ new ccbio.Specimen(value) }}</pre> -->
      </FormKit>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ccbio } from "saacs-es";
import { randomUUID } from "uncrypto";

const route = useRoute();
//    ^?

const specimen = ref<ccbio.Specimen>(
  new ccbio.Specimen({
    collectionId: route.params.collectionId as string,
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
    secondary: {},
    specimenId: randomUUID(),
    taxon: {},
  }),
);

const submitHandler = async (value: ccbio.Specimen) => {
  console.log("submitHandler", value);

  const response = await useCustomFetch(`/api/cc/specimens/create`, {
    method: "POST",
    body: JSON.stringify(value),
  });
  console.log("response", response);
};

console.log("specimen", specimen);
</script>

<style></style>
