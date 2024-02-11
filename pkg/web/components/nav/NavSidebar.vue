<script lang="ts" setup>
// const { data, pending, error } = await useFetch(
//   "/api/cc/collections/listCollections",
// );

const { data, error, pending } = await useCustomFetch(
  '/api/cc/collections/listCollections',
  {
    transform: (data) => {
      return data.collections.map((col) => {
        return {
          col,
          id: col.collectionId,
          to: `/collection-${col.collectionId}`,
        }
      })
    },
  },
)

const links = computed(() => {
  if (!data.value)
    return []

  return data.value?.map((col) => {
    const base = `/collection/${col.id}/AccessControl/`

    return {
      label: col.id,
      to: base,
    }
  })
})
</script>

<template>
  <div v-if="!pending || !error">
    <!-- <div v-if="links">
      <UVerticalNavigation :links="links" />
    </div> -->

    <div
      v-for="col in data"
      :key="col.id"
      class="p-2"
    >
      <QExpansionItem
        :label="col.id"
        caption="Collection"
        expand-separator
        icon="ti-agenda"
      >
        <q-card>
          <q-card-section>
            <!-- <RouterLink>View</RouterLink> -->
            <!-- TODO: Add Dashboard Page -->
            <q-item
              v-ripple
              :to="`/collection/${col.id}/SpecimenTable/`"
              clickable
            >
              <q-item-section avatar>
                <q-icon
                  color="primary"
                  name="ti-view-list-alt"
                />
              </q-item-section>

              <q-item-section>Specimen Table</q-item-section>
            </q-item>
            <q-item
              v-ripple
              :to="`/collection/${col.id}/AccessControl`"
              clickable
            >
              <q-item-section avatar>
                <q-icon
                  color="primary"
                  name="ti-dashboard"
                />
              </q-item-section>

              <q-item-section> Access Control </q-item-section>
            </q-item>
            <q-separator />
            <q-item
              v-ripple
              :to="`/collection/${col.id}/Specimen/New`"
              class=""
              clickable
            >
              <q-item-section avatar>
                <q-icon
                  color="primary"
                  name="ti-plus"
                />
              </q-item-section>

              <q-item-section>New Specimen</q-item-section>
            </q-item>

            <q-item
              v-ripple
              :to="`/collection/${col.id}/Bulk/import`"
              clickable
            >
              <q-item-section avatar>
                <q-icon
                  color="primary"
                  name="ti-import"
                />
              </q-item-section>

              <q-item-section>Bulk Import</q-item-section>
            </q-item>
            <q-item
              v-ripple
              :to="`/collection/${col.id}/Bulk/Update`"
              clickable
            >
              <q-item-section avatar>
                <q-icon
                  color="primary"
                  name="ti-import"
                />
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
