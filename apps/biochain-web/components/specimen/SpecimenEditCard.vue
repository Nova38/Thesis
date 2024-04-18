<script lang="ts" setup>
import type { FieldMask } from '#imports'
import { useFormKitNodeById } from '@formkit/vue'

const mode = defineModel<FormMode>('mode', {
  default: 'view' as FormMode,
})
const specimen = defineModel<PlainSpecimen>('specimen', {
  default: MakeEmptySpecimen(),
})
const props = withDefaults(
  defineProps<{
    collectionId: string
    specimenId: string
    formPrefix?: string
    disabled?: boolean
  }>(),
  {
    formPrefix: 'specimen',
  },
)

const emit = defineEmits<{
  (e: 'submit', value: { specimen: PlainSpecimen; mode: FormMode }): void
}>()

const FormId = computed(
  () => `${props.formPrefix}-${props.collectionId}-${props.specimenId}-form`,
)

const headerColor = computed(() => toModeColor(mode.value))

const modeOptions = ref([
  {
    label: 'View',
    value: 'view' as FormMode,
    attrs: {
      mode: 'view',
    },
  },
  {
    label: 'Update',
    value: 'update' as FormMode,
    attrs: {
      mode: 'update',
    },
  },
  {
    label: 'Suggest',
    value: 'suggest' as FormMode,
    attrs: {
      mode: 'suggest',
    },
  },
])

const nodeRef = useFormKitNodeById(FormId.value, (node) => {
  // perform an effect when the node is available
  node.on('settled.deep', (v) => {})

  node.at('mode')?.on('settled', (v) => {
    console.log('mode settled', v)
  })
})

async function submit(FormData) {
  console.log('submitHandler', FormData)
  const x = new Promise((r) => setTimeout(r, 2000))
}
</script>

<template>
  <div>
    <FormKit
      :id="FormId"
      v-slot="{ value }"
      type="form"
      :actions="false"
      :plugins="[AutoPropsFromIdPlugin, DirtyLabelPlugin]"
      validation-visibility="live"
      @submit="submit"
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
          <div>
            <div>
              <FormKit
                v-model="mode"
                type="radio"
                label="Mailing List"
                :options="modeOptions"
                :classes="{
                  fieldset: 'border-none ',
                  wrapper: 'group/wrapper',
                  option:
                    'relative border p-6 rounded-md formkit-checked:border-none',
                  decorator:
                    '$reset absolute top-0 left-0 right-0 bottom-0 rounded-md group-data-[checked=true]/wrapper:ring',
                  decoratorIcon: '$reset hidden',
                  label: '$reset text-lg text-bold mt-0',
                  optionHelp: '$reset text-sm',
                  options: 'grid grid-cols-3 gap-4 items-strech',
                  outer:
                    '$remove:max-w-[20em]bg-gray-100 dark:bg-gray-900 p-3 max-w-100 grow',
                }"
                help="Choose your preferred list"
              />
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
            @click="submit"
            :outer-class="{
              'max-w-[22em]': false,
              'w-full': true,
            }"
            input-class="w-full justify-center items-center flex"
          >
            Send
          </FormKit>
        </template>
      </UCard>
    </FormKit>
  </div>
</template>

<style scoped></style>
