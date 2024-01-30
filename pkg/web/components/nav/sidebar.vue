<script lang="ts" setup>
// const { data, pending, error } = await useFetch(
//   "/api/cc/collections/listCollections",
// );

const { data, error, pending } = await useCustomFetch(
  '/api/cc/collections/listCollections',
)

const colLinks = computed(() => {
  if (pending.value) return []

  if (error.value) return []

  if (!data.value) return []

  console.log(data.value)
  return data.value.collections.map((collection) => {
    return {
      label: collection.collectionId,
      to: `/collection-${collection.collectionId}`,
    }
  })
})
</script>

<template>
  <div>
    <div v-for="col in colLinks" :key="col.label" class="p-2">
      <QExpansionItem
        :label="col.label"
        caption="Collection"
        expand-separator
        icon="ti-agenda"
      >
        <!-- <template #header>
                    <div class="flex flex-row items-center">
                        <q-icon name="ti-agenda" />
                        <span class="ml-2"> {{ col.label }} </span>
                    </div>
                </template> -->
        <q-card>
          <q-card-section>
            <!-- <RouterLink>View</RouterLink> -->
            <!-- TODO: Add Dashboard Page -->
            <q-item
              v-ripple
              :to="`/collection/${col.label}/SpecimenTable/`"
              clickable
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-view-list-alt" />
              </q-item-section>

              <q-item-section>Specimen Table</q-item-section>
            </q-item>
            <q-item
              v-ripple
              :to="`/collection/${col.label}/AccessControl`"
              clickable
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-dashboard" />
              </q-item-section>

              <q-item-section> Access Control </q-item-section>
            </q-item>
            <q-separator />
            <q-item
              v-ripple
              :to="`/collection/${col.label}/Specimen/New`"
              class=""
              clickable
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-plus" />
              </q-item-section>

              <q-item-section>New Specimen</q-item-section>
            </q-item>

            <q-item
              v-ripple
              :to="`/collection/${col.label}/Specimen/import`"
              clickable
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-import" />
              </q-item-section>

              <q-item-section>Bulk Import</q-item-section>
            </q-item>
            <q-item
              v-ripple
              :to="`/collection/${col.label}/Specimen/update`"
              clickable
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-import" />
              </q-item-section>

              <q-item-section>Bulk Update</q-item-section>
            </q-item>
          </q-card-section>
        </q-card>
      </QExpansionItem>
      <QSeparator />
    </div>
    <!-- </QCardSection>
        </QCard> -->
  </div>
</template>

<style></style>
