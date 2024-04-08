<script lang="ts" setup>
import { titleCase } from 'scule'
import { type FormKitNode } from '@formkit/core'
import { useFormKitContext } from '@formkit/vue'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

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
}

const curValue = ref('{}')

const form = useFormKitContext((node) => {
  curValue.value = JSON.stringify(node.value, null, 2)
  node.node.on('settled.deep', (v) => {
    curValue.value = JSON.stringify(v.origin.value, null, 2)
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
      :plugins="[autoId, dirtyLabel]"
    >
      <PCard>
        <template #content>
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
              <Shiki
                lang="json"
                :code="JSON.stringify(value, null, 2)"
              />
            </PFieldset>
          </DevOnly>
        </template>
      </PCard>
    </FormKit>
  </NuxtErrorBoundary>
</template>

<style scoped></style>
