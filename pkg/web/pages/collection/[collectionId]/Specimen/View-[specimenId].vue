<template>
  <div class="flex flex-row gap-4 p-4">
    <div class="basis-size-3/4 min-w-lg">
      <FormKit
        id="SpecimenForm"
        v-slot="{ value }"
        type="form"
        :value="data"
        :actions="false"
        :disabled="mode === 'view'"
        @submit="submitHandler"
      >
        <SpecimenForm :specimen="value">
          <template #Header>
            <div>
              <QBar :class="modeColor">
                <div>
                  Current Mode: <QChip>{{ mode }}</QChip>
                </div>

                <q-space />
                <q-btn @click="() => (mode = 'view')">
                  View
                  <q-tooltip>Set the current mode to View</q-tooltip>
                </q-btn>
                <q-btn name="Update" @click="() => (mode = 'update')">
                  Update
                  <q-tooltip>Set the current mode to Update</q-tooltip>
                </q-btn>
                <q-btn @click="() => (mode = 'suggest')">
                  Suggest Update
                  <q-tooltip>Set the current mode to Suggest Update</q-tooltip>
                </q-btn>
              </QBar>
            </div>
          </template>
          <template #footer>
            <div class="">
              <FormKit type="submit" />
            </div>
          </template>
        </SpecimenForm>

        <!-- <pre wrap>{{ new ccbio.Specimen(value) }}</pre> -->
        <pre wrap>{{ diffToFieldMaskPath(value, data) }}</pre>
      </FormKit>
    </div>

    <div>
      <SpecimenTimeline :history="history.data.value" class="basis-size-1/4" />
    </div>
    <div></div>
  </div>
</template>

<script lang="ts" setup>
import { auth, ccbio } from "saacs-es";
import { diff } from "ohash";

// const specimen =  useGetSpecimen();
const route = useRoute();

const { data } = await useCustomFetch<ccbio.Specimen>(`/api/cc/specimens/get`, {
  key: specimenIdKey(),

  query: {
    collectionId: route.params.collectionId,
    specimenId: route.params.specimenId,
  },
});

const mode: Ref<FormMode> = ref<FormMode>("view");
const modeColor = computed(() => toModeColor(mode.value));

const history = await useGetSpecimenHistory();

const submitHandler = async (value: ccbio.Specimen) => {
  console.log("submitHandler", value);

  const response = await useCustomFetch(`/api/cc/specimens/create`, {
    method: "POST",
    body: JSON.stringify(value),
  });
  console.log("response", response);
};
</script>

<style>
.formkit-outer {
  margin: var(0);
}
</style>
