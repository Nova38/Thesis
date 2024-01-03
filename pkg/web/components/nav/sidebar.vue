<template>
  <div>
    <UVerticalNavigation :links="baseLinks" :ui="{}" />
    <UDivider
      :ui="{
        size: 'lg',
      }"
    />
    <UAccordion :items="colLinks" :default-open="true" multiple>
      <template #item="{ item }">
        <div class="pl-3">
          <UVerticalNavigation
            :links="[
              {
                label: 'Dashboard',
                to: `/collection-${item.label}/`,

                icon: 'i-heroicons-home',
              },
              {
                label: 'Specimen Table',
                to: `/collection-${item.label}/SpecimenTable`,
                icon: 'i-heroicons-square-3-stack-3d',
              },
              {
                // NewSpecimen
                label: 'New Specimen',
                to: `/collection-${item.label}/NewSpecimen`,
                icon: 'i-heroicons-plus-circle',
              },
            ]"
            :ui="{
              active:
                'text-primary-500 dark:text-primary-400 border-current font-semibold',
              inactive:
                'border-transparent hover:border-gray-400 dark:hover:border-gray-500 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300',
            }"
          />
        </div>
      </template>
    </UAccordion>
    <!-- <QCard class="p-> -->
    <!-- <QCardSection> -->
    <div v-for="col in colLinks" :key="col.label" class="p-2">
      <QExpansionItem
        expand-separator
        icon="ti-agenda"
        :label="col.label"
        caption="Collection"
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
              clickable
              :to="'/collection-' + col.label + '/SpecimenTable/'"
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-view-list-alt" />
              </q-item-section>

              <q-item-section>Specimen Table</q-item-section>
            </q-item>
            <q-item v-ripple clickable :to="'/collection-' + col.label">
              <q-item-section avatar>
                <q-icon color="primary" name="ti-dashboard" />
              </q-item-section>

              <q-item-section> Access Control </q-item-section>
            </q-item>
            <q-separator />
            <q-item
              v-ripple
              clickable
              :to="'/collection-' + col.label + '/NewSpecimen/'"
              class=""
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-plus" />
              </q-item-section>

              <q-item-section>New Specimen</q-item-section>
            </q-item>

            <q-item
              v-ripple
              clickable
              :to="'/collection-' + col.label + '/import'"
            >
              <q-item-section avatar>
                <q-icon color="primary" name="ti-import" />
              </q-item-section>

              <q-item-section>Bulk Import</q-item-section>
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

<script lang="ts" setup>
const { data, pending, error } = await useFetch(
  "/api/cc/collections/listCollections",
);

const colLinks = computed(() => {
  if (pending.value) {
    return [];
  }
  if (error.value) {
    return [];
  }
  if (!data.value) {
    return [];
  }
  console.log(data.value);
  return data.value.collections.map((collection) => {
    return {
      label: collection.collectionId,
      to: `/collection-${collection.collectionId}`,
    };
  });
});

const baseLinks = [
  {
    label: "Home",
    icon: "i-heroicons-home",
    to: "/",
  },
  {
    label: "Collections",
    icon: "i-heroicons-collection",
    to: "/Collections",
  },
  {
    label: "New Collection",
    icon: "i-heroicons-square-3-stack-3d",
    to: "/",
  },
];
</script>

<style></style>
