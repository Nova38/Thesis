<template>
    <div>
        <QCard>
            <QForm>
                <QCardSection>
                    <pre>{{ specimen }}</pre>
                </QCardSection>
            </QForm>
        </QCard>
    </div>
</template>

<script lang="ts" setup>
import type { PlainMessage } from "@bufbuild/protobuf";

import * as ccbio from "saacs-es/dist/gen/biochain/v1";

// const Specimen = pb.ccbio.Specimen;
// type Specimen = typeof pb.ccbio.Specimen;
export interface Props {
    specimen: PlainMessage<ccbio.Specimen>;
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
const emit = defineEmits(["update:specimen"]);

const specimen = useVModel(props, "specimen", emit);
//    ^?

const dateMask = "M/D/YYYY";

// const getJson = (obj: schema.Specimen) => {
//   return schema.Specimen.toJSON(obj);
// };

const DefaultFormProps = ref({
    outlined: true,
    dense: true,
    "stack-label": true,
    // standout: true,
    // rounded: true,
    // filled: true,
});
</script>

<style></style>
