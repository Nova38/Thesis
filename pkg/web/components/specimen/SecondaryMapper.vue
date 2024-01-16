<template>
  <UContainer class="flex flex-row items-center gap-2">
    <UCard>
      <template #header>
        <p>Secondary Sex Mappings</p>
      </template>
      <div class="flex flex-row items-center gap-2">
        <div
          v-for="item in SexOptions"
          :key="item.label"
          class="flex flex-row items-center gap-1"
        >
          <div>{{ item.label }} :</div>
          <div>
            <USelectMenu
              v-model="sexList[item.value]"
              searchable
              searchable-placeholder="Select strings to map"
              class="w-full lg:w-30"
              :multiple="true"
              placeholder="Select strings to map"
              :options="$props.sexStrings"
            >
              <!-- <template #label>
                  <div class="">
                    <span v-if="sexStrings[item.value]" class="truncate">{{
                      sexList[item.value].join(", ")
                    }}</span>
                    <span v-else>Select Sex Strings</span>
                  </div>
                </template> -->
            </USelectMenu>
          </div>
        </div>
      </div>
    </UCard>
    <UCard>
      <template #header>
        <p>Secondary Age Mappings</p>
      </template>
      <div class="flex flex-row items-center gap-2">
        <div
          v-for="item in AgeOptions"
          :key="item.label"
          class="flex flex-row items-center gap-1"
        >
          <div>{{ item.label }} :</div>
          <div>
            <USelectMenu
              v-model="ageList[item.value]"
              searchable
              searchable-placeholder="Select strings to map"
              class="w-full lg:w-30"
              :multiple="true"
              placeholder="Select strings to map"
              :options="$props.sexStrings"
            >
              <template #label>
                <div class="">
                  <span v-if="ageStrings[item.value]" class="truncate">{{
                    ageList[item.value].join(", ")
                  }}</span>
                  <span v-else>Select Age Strings</span>
                </div>
              </template>
            </USelectMenu>
          </div>
        </div>
      </div>
    </UCard>
    <div class="flex flex-row items-center gap-2"></div>
  </UContainer>
</template>

<script lang="ts" setup>
console.log(
  Object.keys(ccbio.Specimen_Secondary_SEX).map((item) => {
    return {
      label: item.toString(),
      value: item,
    };
  }),
);

import { ccbio } from "saacs-es";

const sexList = ref([[], [], [], [], []]);
const ageList = ref([[], [], [], [], [], [], []]);

interface Props {
  sexStrings: string[];
  ageStrings: string[];
}
const props = defineProps<Props>();

const SexOptions = [
  { label: "SEX_UNKNOWN", value: 1 },
  { label: "SEX_ATYPICAL", value: 2 },
  { label: "SEX_MALE", value: 3 },
  { label: "SEX_FEMALE", value: 4 },
];

const AgeOptions = [
  { label: "AGE_UNKNOWN", value: 1 },
  { label: "AGE_NEST", value: 2 },
  { label: "AGE_EMBRYO_EGG", value: 3 },
  { label: "AGE_CHICK_SUBADULT", value: 4 },
  { label: "AGE_ADULT", value: 5 },
  { label: "AGE_CONTINGENT", value: 6 },
];

// const SexOptions = Object.keys(ccbio.Specimen_Secondary_SEX).map((item) => {
//   return {
//     label: item.toString(),
//     value: item,
//   };
// });
// console.log(SexOptions);

const sexMapping = defineModel("sexList", {
  default: ref([[], [], [], [], []]),
});
watch(sexList, (newVal) => {
  console.log(newVal);
  // sexMapping.value = newVal;
});

const ageMapping = defineModel("ageList", {
  default: [[], [], [], [], [], [], []],
});

//
</script>

<style></style>
