<script lang="ts" setup>
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
  }>(),
  {
    formPrefix: 'specimen',
  },
)

const _emit = defineEmits<{
  (e: 'submit', value: { specimen: PlainSpecimen; mode: FormMode }): void
}>()

const FormId = computed(
  () => `${props.formPrefix}-${props.collectionId}-${props.specimenId}-form`,
)

const headerColor = computed(() => toModeColor(mode.value))

const FormDisabled = computed(() => mode.value === 'view')

const modeOptions = ref([
  {
    label: 'View',
    value: 'view' as FormMode,
    attrs: {
      'data-mode': 'view',
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

const _nodeRef = useFormKitNodeById(FormId.value, (node) => {
  // perform an effect when the node is available
  node.on('settled.deep', (v) => {})

  node.at('mode')?.on('settled', (v) => {
    console.log('mode settled', v)
  })
})

async function submit(FormData: unknown) {
  console.log('submitHandler', FormData)
  new Promise((r) => setTimeout(r, 2000))
}
</script>

<template>
  <div>
    <FormKit
      :id="FormId"
      v-slot="{}"
      type="form"
      :actions="false"
      :plugins="[DirtyLabelPlugin]"
      validation-visibility="live"
      @submit="submit"
    >
      <UCard
        class="min-w-2xl max-w-3xl"
        :ui="{
          header: {
            background: headerColor,
            padding: '',
          },
        }"
      >
        <template #header>
          <div>
            <div>
              <FormKit
                v-model="mode"
                type="radio"
                :options="modeOptions"
                :classes="{
                  wrapper: 'group/wrapper',
                  options: 'grid grid-cols-3 items-stretch ',
                  option:
                    '$reset group/option  relative border formkit-checked:border-none    border-none  text-center dark:bg-gray-900 ',
                  outer: '$reset group dark:bg-gray-900 w-full grow px-0 py-0',
                  decorator:
                    '$reset absolute top-0 left-0 right-0 bottom-0  group-data-[checked=true]/wrapper:bg-gray-700/25 dark:group-data-[checked=true]/wrapper:bg-gray-100/50',
                  decoratorIcon: '$reset hidden',
                  label: '!text-md',
                  help: '$reset hidden',
                }"
              />
            </div>
            <UDivider />

            <div class="flex flex-row p-4">
              {{ specimen?.taxon?.genus }} {{ specimen?.taxon?.species }}
            </div>
            <div class="flex flex-grow flex-row">
              <UBadge
                class="flex-grow"
                color="purple"
                variant="solid"
                size="md"
                :ui="{
                  rounded: '',
                }"
                :label="`Collection: ${specimen?.collectionId}`"
              />
              <UBadge
                v-if="specimen.primary?.catalogNumber"
                class="flex-grow"
                color="red"
                variant="solid"
                :ui="{
                  rounded: '',
                }"
                size="md"
                :label="`Catalog Number: ${specimen.primary?.catalogNumber}`"
              />
            </div>
          </div>
        </template>
        <SpecimenForm
          :specimen="specimen"
          :disabled="FormDisabled"
          form-prefix="Main"
        />

        <template #footer>
          <FormKit
            type="submit"
            :outer-class="{
              'max-w-[22em]': false,
              'w-full': true,
            }"
            :disabled="FormDisabled"
            input-class="w-full justify-center items-center flex"
            @click="submit"
          >
            Send
          </FormKit>
        </template>
      </UCard>
    </FormKit>
  </div>
</template>

<style scoped>
input[data-mode='view'] {
  background: #000;
}
</style>
