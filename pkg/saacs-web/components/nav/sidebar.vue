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
                        ]"
                        :ui="{
                            active: 'text-primary-500 dark:text-primary-400 border-current font-semibold',
                            inactive:
                                'border-transparent hover:border-gray-400 dark:hover:border-gray-500 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300',
                        }"
                    />
                </div>
            </template>
        </UAccordion>
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

    return data.value.map((collection) => {
        return {
            label: collection.name,
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

console.log(data);
</script>

<style></style>
