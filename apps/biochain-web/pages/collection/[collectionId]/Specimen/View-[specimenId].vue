<script lang="ts" setup>
import { ccbio } from '#imports'
import { Timestamp, createRegistry } from '@bufbuild/protobuf'
import { keys } from 'radash'
// import { ccbio } from 'saacs'
const route = useRoute()

const dirty = ref(MakeEmptySpecimen())

const cur = ref(MakeEmptySpecimen())
const history = ref(new ccbio.SpecimenHistory())

const spec = useNuxtData(
  makeSpecimenKey(
    route.params?.collectionId.toString(),
    route.params?.specimenId.toString(),
  ),
)

const modeCapitalized = computed(
  () => mode.value[0].toUpperCase() + mode.value.slice(1),
)

// await callOnce(async () => {
//   console.log("This will only be logged once");
//   console.log({
//     c: route.params?.collectionId.toString(),
//     s: route.params?.specimenId.toString(),
//   });
//   websiteConfig.value = await $fetch("/api/cc/specimens/get", {
//     query: {
//       collectionId: toValue(route.params?.collectionId.toString()),
//       specimenId: toValue(route.params?.specimenId.toString()),
//     },
//   });
// });

const getCurrent = useCustomFetch<ccbio.Specimen>(`/api/cc/specimens/get`, {
  key: makeSpecimenKey(
    route.params?.collectionId.toString(),
    route.params?.specimenId.toString(),
  ),

  onResponse: async ({ response }) => {
    console.log('current', response._data)
    dirty.value.fromJson(response._data)
    cur.value.fromJson(response._data)
    console.log(keys(response._data))
    console.log(dirty.value)
  },
  query: {
    collectionId: toValue(route.params?.collectionId.toString()),
    specimenId: toValue(route.params?.specimenId.toString()),
  },
})

const getHistory = useCustomFetch<ccbio.Specimen>(`/api/cc/specimens/history`, {
  key: makeSpecimenHistoryKey(
    route.params?.collectionId.toString(),
    route.params?.specimenId.toString(),
  ),

  onResponse: async ({ response }) => {
    console.log('history', response._data)
    history.value.fromJson(response._data)
    console.log(history.value)
  },
  query: {
    collectionId: toValue(route.params?.collectionId.toString()),
    specimenId: toValue(route.params?.specimenId.toString()),
  },
})

// const logDiff = () => {
//   console.log({
//     diff: diff(crush(dirty.value), crush(cur.value)),
//   });
// };

const mode: Ref<FormMode> = ref('view')
const modeColor = computed(() => toModeColor(mode.value))

// const history = await useGetSpecimenHistory();

async function submitHandler() {
  const oldMode = mode.value
  try {
    mode.value = 'view'
    // const specimen = new ccbio.Specimen(getCurrent.data.value ?? {});

    const { differences, mask } = diffCrush(toValue(cur), toValue(dirty), [
      'lastModified',
    ])
    console.log({ differences })

    const req = new ccbio.SpecimenUpdate({
      mask,
      specimen: dirty.value,
    })

    const response = await useCustomFetch(`/api/cc/specimens/update`, {
      body: req.toJsonString({
        typeRegistry: createRegistry(ccbio.Specimen, Timestamp),
      }),
      method: 'POST',
    })
    console.log('response', toValue(response.data))
    // dirty.value.fromJson(response.data);
    getHistory.refresh()
    getCurrent.refresh()
    console.log('current', toValue(getCurrent.data))
  } catch (error) {
    mode.value = oldMode
    console.error(error)
  }
}
</script>

<template>
  <div class="flex flex-row gap-4 p-4">
    <div class="basis-size-3/4 min-w-lg">
      <div v-if="spec.data">
        <SpecimenForm
          :enable-edit="mode !== 'view'"
          :header-color="toModeColor(mode)"
          :specimen="dirty"
        >
          <template #Header>
            <div>
              <div
                class=""
                :class="`${modeColor} flex flex-row items-center gap-2 px-2`"
              >
                <div class="font-bold">Current Mode: {{ modeCapitalized }}</div>

                <div class="flex-grow" />
                <!-- <space /> -->
                <UTooltip text="Set the current mode to View">
                  <UButton
                    :class="`${toModeColor('view')}text-black`"
                    @click="() => (mode = 'view')"
                  >
                    View
                  </UButton>
                </UTooltip>

                <UTooltip text="Set the current mode to Update">
                  <UButton
                    :class="toModeColor('update')"
                    name="Update"
                    @click="() => (mode = 'update')"
                  >
                    Update
                  </UButton>
                </UTooltip>

                <UTooltip text="Set the current mode to Suggest Update">
                  <UButton
                    variant="solid"
                    :class="toModeColor('suggest')"
                    :color="toModeColor('suggest')"
                    @click="() => (mode = 'suggest')"
                  >
                    Suggest Update
                  </UButton>
                </UTooltip>
              </div>
              <UDivider />
            </div>
          </template>
          <template #Footer>
            <div class="flex flex-col">
              <UButton
                v-if="mode !== 'view'"
                :class="modeColor"
                :label="mode === 'update' ? 'Update' : 'Suggest Update'"
                class="flex-grow"
                @click="submitHandler"
              />
            </div>
          </template>
        </SpecimenForm>
      </div>
    </div>
    <div class="flex flex-col gap-4">
      <SpecimenFormKit v-model="dirty" />
    </div>
    <div class="flex flex-col gap-4">
      <SpecimenTimeline
        :can-hide="$auth.loggedIn || false"
        :history="history"
        class="basis-size-1/4"
      />
      <UCard v-if="false">
        <div class="flex flex-row text-lg items-center justify-center">
          Suggestions
        </div>
      </UCard>
    </div>
  </div>
</template>

<style></style>
