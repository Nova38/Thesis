<script lang="ts" setup>
import { useFormKitNodeById } from '@formkit/vue'

const mode = defineModel<FormMode>('mode', {
  default: 'view' as FormMode,
})

const headerColor = computed(() => toModeColor(mode.value))

const props = withDefaults(
  defineProps<{
    specimen: PlainSpecimen
  }>(),
  {
    specimen: () => MakeEmptySpecimen(),
  },
)

const specimen = ref<PlainSpecimen>(props.specimen)

const nodeRef = useFormKitNodeById('Main-form', (node) => {
  // perform an effect when the node is available
  node.on('settled.deep', (v) => {
    console.log('settled.deep', v)
    specimen.value = v.payload
  })
})
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
      <SpecimenForm
        :specimen="props.specimen"
        form-prefix="Main"
      />
      <template #footer>
        <FormKit
          type="submit"
          class="w-full"
        >
          Submit
        </FormKit>
      </template>
    </UCard>
  </div>
</template>

<style scoped></style>
