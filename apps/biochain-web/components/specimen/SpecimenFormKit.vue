<script lang="ts" setup>
import vueJsonPretty from '../../../../libs/nuxt/ui/plugins/vue-json-pretty'

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
</script>

<template>
  <FormKit
    type="form"
    v-model="specimen"
    #default="{ value }"
    @submit="submitHandler"
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
