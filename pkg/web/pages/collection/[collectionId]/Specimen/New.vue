<template>
  <div>
    <div>
      <FormKit
        id="SpecimenForm"
        v-slot="{ value }"
        type="form"
        :value="specimen"
        :actions="false"
        @submit="submitHandler"
      >
        <FormKit type="meta" name="notes" :value="extraData" />
        <SpecimenForm :specimen="value">
          <template #footer>
            <div>
              <FormKit
                type="submit"
                label="Submit New Specimen"
                class="justify-center"
              />
            </div>
          </template>
        </SpecimenForm>

        <!-- <pre wrap>{{ value }}</pre> -->
      </FormKit>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ccbio } from "saacs-es";
import { randomUUID } from "uncrypto";

const route = useRoute();
//    ^?

const specimen = ref(
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

  const router = useRouter();
  router.push(
    `/collection/${route.params.collectionId}/specimen/View-${value.specimenId}`,
  );

  console.log("response", response);
};
const extraData = {
  hair: "gold",
  eyes: "blue",
  weight: "215lb",
  height: "6ft 3in",
  hands: "tiny",
  cool: false,
};
console.log("specimen", specimen);
</script>

<style></style>
