<script lang="ts" setup>
import type { FieldMask } from '#imports'
import { useFormKitNodeById } from '@formkit/vue'

const mode = defineModel<FormMode>('mode', {
  default: 'view' as FormMode,
})
const specimen = defineModel<PlainSpecimen>('specimen', {
  default: MakeEmptySpecimen(),
})

const emit = defineEmits<{
  (e: 'submit', value: { specimen: PlainSpecimen; mode: FormMode }): void
}>()

const headerColor = computed(() => toModeColor(mode.value))

// const props = withDefaults(
//   defineProps<{
//     specimen: PlainSpecimen
//   }>(),
//   {
//     specimen: () => MakeEmptySpecimen(),
//   },
// )

const specimenValue = ref<PlainSpecimen>()

const nodeRef = useFormKitNodeById('Main-form', (node) => {
  // perform an effect when the node is available
  node.on('settled.deep', (v) => {
    // console.log('settled.deep', v)
    // specimenValue.value = v.payload
    // console.log('specimenValue', specimenValue.value)
  })
})

const submitHandler = () => {
  console.log('submitHandler', specimen.value)
  // emit('submit', {
  //   specimen: specimen.value,
  //   mode: mode.value,
  // })
}
</script>

<template>
  <div>
    <UCard
      class="min-w-2xl max-w-3xl"
      :ui="{
        header: {
          background: headerColor,
        },
      }"
    >
      <template #header>
        <div>
          <div>
            <UButton @click="submitHandler">View</UButton>
            <UButton @click="submitHandler">Update</UButton>
            <UButton @click="submitHandler">Suggest</UButton>
          </div>
          <UDivider />
          <div class="flex flex-row">
            {{ specimen?.taxon?.genus }} {{ specimen?.taxon?.species }}
          </div>
          <div class="flex flex-grow flex-row">
            <UBadge
              class="flex-grow"
              color="purple"
              variant="solid"
              size="md"
              :label="`Collection: ${specimen?.collectionId}`"
            />
            <UBadge
              v-if="specimen.primary?.catalogNumber"
              class="flex-grow"
              color="red"
              variant="solid"
              size="md"
              :label="`Catalog Number: ${specimen.primary?.catalogNumber}`"
            />
          </div>
        </div>
      </template>
      <SpecimenForm
        :specimen="specimen"
        form-prefix="Main"
      />
      <template #footer>
        <FormKit
          type="submit"
          class="w-full"
          @click="submitHandler"
        >
          Submit
        </FormKit>
        <UButton
          block
          @click="submitHandler"
          >Button</UButton
        >
      </template>
    </UCard>
  </div>
</template>

<style scoped></style>
