<script lang="ts" setup>
// import { ccbio } from 'saacs'
// const { data, pending, error, refresh } = await useGetSpecimen();

// const entry = defineModel<ccbio.SpecimenHistoryEntry>('entry', {
//   default: () => new ccbio.SpecimenHistoryEntry(),
// })

// const isLast = defineModel('isLast', {
//   default: () => false,
//   type: Boolean,
// })
// const hidden = defineModel('hidden', {
//   default: () => false,
//   type: Boolean,
// })

// async function sendHide() {
//   console.log('sendHide', entry.value)
// }

// async function sendUnHide() {
//   console.log('sendUnHide', entry.value)
// }

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

const _emit = defineEmits<{
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
          },
        }"
      >
        <template #header> </template>
        <SpecimenForm
          :specimen="specimen"
          form-prefix="Main"
        />

        <template #footer>
          <FormKit
            type="submit"
            :outer-class="{
              'max-w-[22em]': false,
              'w-full': true,
            }"
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

<style scoped></style>
