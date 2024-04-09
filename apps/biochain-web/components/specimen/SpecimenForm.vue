<script lang="ts" setup>
import { titleCase } from 'scule'
import { type FormKitNode } from '@formkit/core'
import { useFormKitContext } from '@formkit/vue'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

const mode = defineModel<FormMode>('mode', {
  default: 'view' as FormMode,
})

const headerColor = computed(() => toModeColor(mode.value))

const submitHandler = async () => {
  try {
    console.log('submitHandler', specimen.value)
  } catch (error) {
    console.error('submitHandler', error)
  }
}

const autoId = (node: FormKitNode) => {
  if (!node.props.id) return
  console.log('AutoPropsFromIdPlugin', node.props.id)
  node.name ??= node.props.id // auto set name to id
  node.props.label ??= titleCase(node.props.id) // auto set label
  if (!['button', 'submit'].includes(node.props.type)) {
    // automatically set help text, but exclude buttons
  }
}

const dirtyLabel = (node: FormKitNode) => {
  if (!node.context) return
  console.log(node.context)
  node.context.classes.label = node.context?.state.dirty
    ? 'text-red-500'
    : 'text-surface-700 dark:text-surface-0/80'
  node.on('commit', (v) => {
    if (node.context?.state.dirty) {
      node.context.help = '* Modified'
    }
    console.log('commit', v)
  })
}

const form = useFormKitContext((node) => {
  node.node.on('settled.deep', (v) => {
    console.log('settled.deep', v)
    // readValue.value = v
  })
})
</script>

<template>
  <NuxtErrorBoundary>
    <template #error="{ error }">
      <div>
        <p>Oops, it looks like something broke</p>
        <pre> {{ error }}</pre>
        <p>{{ error.message }}</p>
        <p><button @click="clearError(error)">clear error</button></p>
      </div>
    </template>

    <FormKit
      type="form"
      v-model="specimen"
      #default="{ value }"
      @submit="submitHandler"
      dirty-behavior="compare"
      :actions="false"
      :plugins="[autoId, dirtyLabel]"
      validation-visibility="live"
    >
      <UCard
        class="min-w-2xl max-w-3xl"
        :ui="{
          header: {
            background: headerColor,
          },
        }"
      >
        <template #header>
          <!-- <PTag :value="specimen.collectionId" /> -->
          <div>
            <div class="flex flex-row">
              {{ specimen?.taxon?.genus }} {{ specimen?.taxon?.species }}
            </div>
            <div class="flex flex-grow flex-row">
              <UBadge
                class="flex-grow"
                color="purple"
                variant="solid"
                size="md"
                :label="`Collection: ${specimen.collectionId}`"
              />
              <UBadge
                class="flex-grow"
                color="red"
                variant="solid"
                size="md"
                v-if="specimen.primary?.catalogNumber"
                :label="`Catalog Number: ${specimen.primary?.catalogNumber}`"
              />
            </div>
          </div>
        </template>
        <template #title> </template>
        <template #subtitle> </template>
        <FormKit
          name="specimenId"
          type="hidden"
          value=""
        />

        <SpecimenFormTaxon />
        <SpecimenFormPrimary />
        <SpecimenFormGeoreference />
        <SpecimenFormSecondary />
        <SpecimenFormLoans />
        <DevOnly>
          <PFieldset
            legend="Value"
            :toggleable="true"
          >
            <NuxtErrorBoundary>
              <Shiki
                lang="json"
                class="max-w-2xl overflow-x-scroll"
                :code="JSON.stringify(value, null, 2)"
              />
            </NuxtErrorBoundary>
          </PFieldset>
        </DevOnly>

        <template #footer>
          <FormKit
            type="submit"
            class="w-full"
          >
            Submit
          </FormKit>
        </template>
      </UCard>
    </FormKit>
  </NuxtErrorBoundary>
</template>

<style scoped></style>
