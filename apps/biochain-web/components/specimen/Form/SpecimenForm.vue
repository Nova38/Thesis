<script lang="ts" setup>
import { titleCase } from 'scule'
import { type FormKitNode } from '@formkit/core'
import { useFormKitContext } from '@formkit/vue'
import { DirtyLabelPlugin } from '#imports'

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
    :plugins="[autoId, DirtyLabelPlugin]"
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
    <SpecimenFormGeoreferenceGrid />
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
