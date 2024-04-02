<script lang="ts" setup>
import { titleCase } from 'scule'
import vueJsonPretty from '../../../../libs/nuxt/ui/plugins/vue-json-pretty'
import { type FormKitNode } from '@formkit/core'

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
  node.name = node.props.id // auto set name to id
  node.props.label = titleCase(node.props.id) // auto set label
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
</script>

<template>
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
</template>

<style scoped></style>
