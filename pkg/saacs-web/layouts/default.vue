<template>
    <div>
        <UCard
            class="h-screen v-full"
            :ui="{
                header: {
                    padding: 'px-2 py-2 sm:px-2',
                },
            }"
        >
            <template #header>
                <div class="flex flex-row items-center">
                    <UBreadcrumb :links="crumb" class="" />
                    <AuthUserMenu class="ml-auto" />
                </div>
            </template>
            <div class="flex flex-row w-full flex-nowrap">
                <div class="p-2 mr-4 flex-none grow-0">
                    <NavSidebar />
                </div>
                <div class="grow h-full">
                    <UCard>
                        <template #header>
                            <div></div>
                        </template>
                        <slot name="content"> </slot>
                    </UCard>
                </div>
            </div>
        </UCard>
    </div>
</template>

<script lang="ts" setup>
// const crumb = [
//     {
//         label: "Home",
//         icon: "i-heroicons-home",
//         to: "/",
//     },
//     {
//         label: "Navigation",
//         icon: "i-heroicons-square-3-stack-3d",
//     },
//     {
//         label: "Breadcrumb",
//         icon: "i-heroicons-link",
//     },
// ];
const crumb = useBreadcrumbLinks();
</script>

<style></style>
