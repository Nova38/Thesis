<script lang="ts" setup>
import { titleCase } from 'scule'
import { type FormKitNode } from '@formkit/core'
import { useFormKitContext } from '@formkit/vue'

const specimen = defineModel<PlainSpecimen>('specimen', {
  default: () => MakeEmptySpecimen(),
})

const props = withDefaults(
  defineProps<{
    specimen: PlainSpecimen
    formPrefix: string
  }>(),
  {
    specimen: () => MakeEmptySpecimen(),
    formPrefix: 'specimen',
  },
)

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

// const form = useFormKitContext((node) => {
//   node.node.on('settled.deep', (v) => {
//     console.log('settled.deep', v)
//     // readValue.value = v
//   })
// })

const FormId = computed(() => `${props.formPrefix}-form`)
</script>

<template>
  <FormKit
    type="form"
    :id="FormId"
    v-model="specimen"
    #default="{ value }"
    dirty-behavior="compare"
    :actions="false"
    :plugins="[autoId, dirtyLabel]"
    validation-visibility="live"
  >
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
  </FormKit>
</template>

<style scoped></style>
