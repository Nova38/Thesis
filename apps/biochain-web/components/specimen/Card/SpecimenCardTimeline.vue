<script lang="ts" setup>
import { toPlainMessage } from '@bufbuild/protobuf'

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
const FormId = computed(
  () => `${props.formPrefix}-${props.collectionId}-${props.specimenId}-form`,
)

const { isLoading, isSuccess, data } = useQuery({
  queryKey: ['specimen', 'history', props.collectionId, props.specimenId],
  queryFn: async () =>
    await $fetch('/api/cc/specimens/history', {
      query: {
        collectionId: props.collectionId,
        specimenId: props.specimenId,
      },
    }),
})

const hideTransaction = (tx: string) => {
  console.log('hide', tx)
}

const items = computed(() => {
  if (!isSuccess || !data.value) {
    return []
  }
  return data.value?.entries.map((entry) => {
    if (!entry.timestamp) {
      return {
        label: 'Unknown',
        icon: 'i-heroicons-calendar',
        tx: entry.txId,
        specimen: entry.value,
      }
    }
    return {
      label: entry.timestamp ?? 'Unknown',
      icon: 'i-heroicons-calendar',
      tx: entry.txId,
      specimen: entry.value,
    }
  })
})
</script>

<template>
  <div>
    <UCard>
      <template #header>
        <div class="flex flex-row items-center justify-center text-lg">
          History
        </div>
      </template>
      <div v-if="isSuccess && data">
        <UAccordion
          :items="items"
          :ui="{ wrapper: 'flex flex-col w-full' }"
          multiple
        >
          <template #default="{ item, index, open }">
            <UButton
              color="gray"
              variant="ghost"
              class="border-b border-gray-200 dark:border-gray-700"
              :ui="{ rounded: 'rounded-none', padding: { sm: 'p-3' } }"
            >
              <template #leading>
                <div
                  class="bg-primary-500 dark:bg-primary-400 -my-1 flex h-6 w-6 items-center justify-center rounded-full"
                >
                  <UIcon
                    :name="item.icon"
                    class="h-4 w-4 text-white dark:text-gray-900"
                  />
                </div>
              </template>

              <span class="truncate">{{ index + 1 }}. {{ item.label }}</span>

              <template #trailing>
                <UIcon
                  name="i-heroicons-chevron-right-20-solid"
                  class="ms-auto h-5 w-5 transform transition-transform duration-200"
                  :class="[open && 'rotate-90']"
                />
              </template>
            </UButton>
          </template>
          <template #item="{ item }">
            <UCard
              class="min-w-2xl mt-4 max-w-3xl"
              :ui="{ root: 'bg-surface-100 dark:bg-surface-800' }"
            >
              <FormKit
                disabled
                type="form"
                :actions="false"
              >
                <SpecimenForm
                  :specimen="item.specimen"
                  :form-prefix="`specimen-${item.tx}`"
                />
              </FormKit>
              <template
                v-if="$auth.loggedIn.value"
                #footer
              >
                <UButton
                  color="red"
                  variant="ghost"
                  class="w-full"
                  @click="hideTransaction(item.tx)"
                >
                  Hide Transaction
                </UButton>
              </template>
            </UCard>
          </template>
        </UAccordion>
      </div>
    </UCard>
  </div>
</template>

<style scoped></style>
