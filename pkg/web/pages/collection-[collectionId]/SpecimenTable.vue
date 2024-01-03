<template>
  <div>
    <!-- <Testtable /> -->
    <!-- <SpecimenTable /> -->

    <QTable></QTable>
  </div>
</template>

<script lang="ts" setup>
// const links = useBreadcrumbLinks();

import { z } from "zod";

z.object({
  id: z.string(),
  name: z.string(),
  age: z.number(),
});
</script>

<style></style>
