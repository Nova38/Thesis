<template>
  <div>
    <!-- <Testtable /> -->
    <!-- <SpecimenTable /> -->
    <div>
      <QBtn @click="() => refresh"></QBtn>
    </div>
    <QTable> </QTable>
    <pre wrap>{{ data }}</pre>
  </div>
</template>

<script lang="ts" setup>
// const links = useBreadcrumbLinks();
const route = useRoute();
// const { data, pending, error, refresh } = useFetch(
//   `/api/cc/specimens/list?collectionId=${route.params.collectionId}`,
// );

const { data, pending, error, refresh } = useCustomFetch(
  `/api/cc/specimens/list?collectionId=${route.params.collectionId}`,
  {
    query: {
      collectionId: route.params.collectionId,
    },
  },
);
</script>

<style></style>
