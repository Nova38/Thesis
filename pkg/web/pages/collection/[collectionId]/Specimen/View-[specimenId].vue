<script lang="ts" setup>
import { Timestamp, createRegistry } from '@bufbuild/protobuf'
import { keys } from 'radash'
import { ccbio } from '~/lib'

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
  }
  catch (error) {
    mode.value = oldMode
    console.error(error)
  }
}
</script>

<template>
  <div class="flex flex-row gap-4 p-4">
    <!-- <QBtn label="Log Diff" @click="logDiff" /> -->
    <div class="basis-size-3/4 min-w-lg">
      <div v-if="spec.data">
        <SpecimenForm
          :enable-edit="mode != 'view'"
          :header-color="toModeColor(mode)"
          :specimen="dirty"
        >
          <template #Header>
            <div>
              <QBar
                :class="modeColor"
                class="p-2"
              >
                <div class="font-bold">
                  Current Mode: {{ modeCapitalized }}
                </div>

                <q-space />
                <q-btn
                  :class="toModeColor('view')"
                  @click="() => (mode = 'view')"
                >
                  View
                  <q-tooltip>Set the current mode to View</q-tooltip>
                </q-btn>
                <q-btn
                  :class="toModeColor('update')"
                  name="Update"
                  @click="() => (mode = 'update')"
                >
                  Update
                  <q-tooltip>Set the current mode to Update</q-tooltip>
                </q-btn>
                <q-btn
                  :class="toModeColor('suggest')"
                  @click="() => (mode = 'suggest')"
                >
                  Suggest Update
                  <q-tooltip>Set the current mode to Suggest Update</q-tooltip>
                </q-btn>
              </QBar>
              <QSeparator class="bg-black" />
            </div>
          </template>
          <template #Footer>
            <div class="flex flex-col">
              <QBtn
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
      <SpecimenTimeline
        :can-hide="$auth.loggedIn || false"
        :history="history"
        class="basis-size-1/4"
      />
      <QCard v-if="false">
        <QBar class="flex flex-row text-lg items-center justify-center">
          Suggestions
        </QBar>
      </QCard>
    </div>
    <QCard> {</QCard>
  </div>
</template>

<style></style>
