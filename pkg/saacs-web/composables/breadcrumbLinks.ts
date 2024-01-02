import type { BreadcrumbLink } from "@nuxt/ui/dist/runtime/types";

export const useBreadcrumbLinks = () => {
    const route = useRoute();
    const links = computed(() => {
        const paths = route.path.split("/").filter((p) => p);

        const items: BreadcrumbLink[] = [
            {
                label: "Home",
                to: "/",
                icon: "i-heroicons-home",
            },
        ];

        // Check for the collection route
        if (route.params?.collectionId) {
            items.push({
                label: `Collection: ${route.params.collectionId}`,
                to: `/collection-${route.params.collectionId}`,
                // material-symbols:collections-bookmark-outline-rounded
                icon: "i-material-symbols-collections-bookmark-outline-rounded",
            });
        }

        // Handle the Specimen route
        if (paths.find((p) => p === "Specimen")) {
            // items.push({
            //     label: `Specimen`,
            //     to: `/specimen`,
            //     icon: "i-heroicons-moon",
            // });

            if (route.params?.specimenId) {
                // See if we have a speciemen id
                items.push({
                    label: `Specimen: ${route.params.specimenId}`,
                    to: `/specimen/view-${route.params.specimenId}`,
                    icon: "i-heroicons-moon",
                });
            }
        }

        console.log("items", items);
        console.log("paths", paths);

        return items;
        //
        // const links = paths.map((p, i) => {
        //     const link = {
        //         title: p,
        //         to: "/" + paths.slice(0, i + 1).join("/"),
        //     };
        //     return link;
        // });
        // return links;
    });

    return links;
};
