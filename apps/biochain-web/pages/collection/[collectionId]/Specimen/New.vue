<script lang="ts" setup>
//    ^?
const nuxtApp = useNuxtApp()

const specimen = ref(
  new ccbio.Specimen({
    collectionId: nuxtApp.$collectionId.value,
    specimenId: '',

    georeference: {
      georeferenceDate: {},
    },
    grants: {},
    images: {},
    loans: {},
    primary: {
      catalogNumber: '',
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
    taxon: {},
  }),
)

watchDeep(specimen, (value, oldValue) => {
  if (value?.primary?.catalogNumber) {
    console.log('catalogNumber', value.primary.catalogNumber)

    nextTick(
      () =>
        (value.specimenId = CatNumToUUID(value?.primary?.catalogNumber || '')),
    )
  }

  console.log('specimen', value.specimenId)
})

const api = useCustomFetch(`/api/cc/specimens/create`, {
  method: 'POST',
  immediate: false,

  onRequest: ({ options }) => {
    options.body = new ccbio.Specimen(toValue(specimen)).toJsonString({
      emitDefaultValues: true,
      enumAsInteger: true,
    })
  },
})

async function submitHandler() {
  // console.log("submitHandler", value);

  try {
    console.log(api.data.value)

    const response = await api.execute()
    console.log({ api, response, specimen: specimen.value })
    if (api.status.value === 'success') {
      nuxtApp.$router.push(
        `/collection/${nuxtApp.$collectionId.value}/specimen/View-${specimen.value.specimenId}`,
      )
    }
  } catch (error) {
    console.error('error', error)
  }
}
</script>

<template>
  <div>
    <div>
      <SpecimenForm :specimen="specimen">
        <template #Footer>
          <div>
            <UButton
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
