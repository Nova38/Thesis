<script lang="ts" setup>
//    ^?
import type { FormKitNode } from '@formkit/core'
import { pb, type PlainSpecimen } from '#imports'

const nuxtApp = useNuxtApp()
const toast = useToast()

const specimen = ref({
  specimen: new ccbio.Specimen({
    collectionId: nuxtApp.$collectionId.value,
    specimenId: '',

    georeference: {
      georeferenceDate: {},
    },
    grants: {},
    images: {},
    loans: {
      initial: {
        loanedDate: {},
        description: '',
        loanedTo: '',
        id: '',
        loanedBy: '',
      },
    },
    primary: {
      catalogNumber: '',
      catalogDate: {},
      determinedDate: {},
      fieldDate: {},
      originalDate: {},
    },
    secondary: {
      preparations: {
        initial: {},
      },
    },
    taxon: {},
  }),
})

// watchDeep(specimen, (value, _oldValue) => {
//   if (value?.specimen?.primary?.catalogNumber) {
//     console.log('catalogNumber', value.specimen?.primary.catalogNumber)
//
//     nextTick(
//       () =>
//         (value.specimen.specimenId = CatNumToUUID(value.specimen?.primary?.catalogNumber || '')),
//     )
//   }
//
//   console.log('specimen', value.specimenId)
// })

//  const api = useCustomFetch(`/api/cc/specimens/create`, {
//    method: 'POST',
//    immediate: false,
//
//    onRequest: ({ options }) => {
//      options.body = new ccbio.Specimen(toValue(specimen)).toJsonString({
//        emitDefaultValues: true,
//      })
//    },
//  })

async function submitHandler(
  specimen: { specimen: PlainSpecimen },
  node: FormKitNode,
) {
  console.log({ specimen, node })

  const s = specimen.specimen
  s.specimenId = CatNumToUUID(s?.primary?.catalogNumber ?? '')

  console.log({ s, node })

  const m = toast.add({
    title: 'Creating Specimen',
    description: 'Please wait...',
    timeout: 2000,
  })

  try {
    const value = await $fetch('/api/cc/specimens/create', {
      method: 'POST',
      body: new ccbio.Specimen(s).toJsonString({ enumAsInteger: true }),
    })

    if (value) {
      nuxtApp.$router.push(
        `/collection/${value.unpacked.collectionId}/specimen/View-${value.unpacked.specimenId}`,
      )
    }
  } catch (error) {
    console.log('error', error)
    node.setErrors(['Something went wrong with the server, please try again'])
    // comment out this line and refresh after submit
    // to see how values would otherwise be lost.
  }
}
</script>

<template>
  <div>
    <FormKit
      id="NewSpecimen"
      v-slot="{ disabled }"
      v-model="specimen"
      type="form"
      :plugins="[]"
      submit-label="Create Character"
      :actions="false"
      @submit="submitHandler"
    >
      <UCard class="min-w-2xl max-w-3xl">
        <template #header>
          <h3 class="text-lg font-semibold">New Specimen</h3>
        </template>

        <SpecimenForm />

        <template #footer>
          <FormKit
            type="submit"
            :outer-class="{
              'max-w-[22em]': false,
              'w-full': true,
            }"
            input-class="w-full justify-center items-center flex"
            :disabled="disabled"
          >
            Send
          </FormKit>
        </template>
      </UCard>
    </FormKit>
  </div>
</template>

<style></style>
