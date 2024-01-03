<template>
  <div>
    <QCard>
      <QCardSection :v-if="props.showHeader">
        <slot name="HeaderBar" />
      </QCardSection>
      <QCardSection> </QCardSection>
    </QCard>
  </div>
</template>

<script lang="ts" setup>
// import type { PlainMessage } from "@bufbuild/protobuf";

import { ccbio } from "saacs-es";

// const sp = reactive(new ccbio.Specimen());

// const Specimen = pb.ccbio.Specimen;
// type Specimen = typeof pb.ccbio.Specimen;
export interface Props {
  specimen: ccbio.Specimen;
  enableEdit?: boolean;
  showUpdater?: boolean;
  showHeader?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  editable: false,
  enableEdit: true,
  showUpdater: true,
  showHeader: true,
  showEditOptions: true,
});
// const emit = defineEmits(["update:specimen"]);

// const specimen = useVModel(props, "specimen", emit);
//    ^?

// const dateMask = "M/D/YYYY";

// const getJson = (obj: schema.Specimen) => {
//   return schema.Specimen.toJSON(obj);
// };
</script>

<style></style>
