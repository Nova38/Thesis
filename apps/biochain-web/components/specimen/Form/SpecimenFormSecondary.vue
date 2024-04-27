<script lang="ts" setup>
// import { useFormKitContext } from '#imports'
import { ccbio } from '#imports'
import { toPlainMessage } from '@bufbuild/protobuf'
import { useFormKitContext } from '@formkit/vue'
const sNode = ref(null)
const collapsed = defineModel<boolean>('collapsed', {
  default: true,
})

const defaultSecondary = toPlainMessage(new ccbio.Specimen_Secondary({}))

const secondary = useFormKitContext('specimen.secondary', (node) => {
  node.node.on('settled.deep', () => {})
})

function addPreparation() {
  if (!secondary.value) return
  const cur = secondary.value.node.value as ccbio.Specimen_Secondary

  const preparations = secondary.value.node.at('$self.preparations')

  console.log('addPreparation', { cur, secondary, preparations })
  if (!preparations) {
    secondary.value.node.input({
      ...cur,
      preparations: {
        '1': new ccbio.Specimen_Secondary_Preparation({}).toJson(),
      },
    })
  } else {
    preparations.input({
      ...cur.preparations,
      [Object.keys(cur.preparations).length + 1]:
        new ccbio.Specimen_Secondary_Preparation({}),
    })
  }
}
</script>

<template>
  <div>
    <PFieldset
      legend="Secondary"
      toggleable
      :collapsed
    >
      <FormKit
        id="secondary"
        ref="sNode"
        v-slot="{ value: v, state: { disabled } }"
        type="group"
        name="secondary"
        :value="defaultSecondary"
      >
        <div class="subGroup">
          <FormKit
            type="number"
            number
            step="any"
            name="weight"
            label="weight"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="weightUnits"
            label="Weight Units"
            outer-class="min-w-20"
          />
          <FormKit
            type="text"
            name="molt"
            label="Molt"
            outer-class="min-w-20"
          />
        </div>
        <div class="subGroup">
          <FormKit
            id="sex"
            type="select"
            name="sex"
            label="Sex"
            outer-class="min-w-20"
            :options="[
              {
                label: 'Not Specified',
                value: 0,
              },
              {
                label: 'Unknown',
                value: 1,
              },
              {
                label: 'Atypical',
                value: 2,
              },
              {
                label: 'Female',
                value: 3,
              },
              {
                label: 'Male',
                value: 4,
              },
            ]"
          />
          <FormKit
            id="age"
            type="select"
            name="age"
            label="Age"
            outer-class="min-w-20"
            :options="[
              {
                label: 'Not Specified',
                value: ccbio.Specimen_Secondary_AGE.AGE_UNDEFINED,
              },
              {
                label: 'Unknown',
                value: ccbio.Specimen_Secondary_AGE.AGE_UNKNOWN,
              },
              {
                label: 'Nest',
                value: ccbio.Specimen_Secondary_AGE.AGE_NEST,
              },
              {
                label: 'Embryo/Egg',
                value: ccbio.Specimen_Secondary_AGE.AGE_EMBRYO_EGG,
              },
              {
                label: 'SubAdult',
                value: ccbio.Specimen_Secondary_AGE.AGE_CHICK_SUBADULT,
              },
              {
                label: 'Adult',
                value: ccbio.Specimen_Secondary_AGE.AGE_ADULT,
              },
              {
                label: 'Contingent',
                value: ccbio.Specimen_Secondary_AGE.AGE_CONTINGENT,
              },
            ]"
          />
        </div>

        <div class="subGroup">
          <FormKit
            type="textarea"
            name="notes"
            label="Notes"
            outer-class="min-w-20"
          />
          <FormKit
            type="textarea"
            name="condition"
            label="Condition"
            outer-class="min-w-20"
          />
        </div>
        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            v-if="v?.preparations"
            id="preparations"
            v-slot="{ value }"
            type="group"
            name="preparations"
          >
            <template
              v-for="(_v, name) in value"
              :key="name"
            >
              <template v-if="typeof name === 'string'">
                <FormKit
                  :id="name"
                  type="group"
                  :name="name"
                >
                  <FormKit
                    id="verbatim"
                    type="text"
                    :label="`Preparation #${name}`"
                    name="verbatim"
                    outer-class="min-w-20"
                  />
                </FormKit>
              </template>
            </template>
          </FormKit>
        </div>

        <div class="flex flex-wrap gap-2">
          <FormKit
            outer-class="w-full flex-grow"
            type="button"
            label="Add Preparation"
            :disabled="disabled"
            @click="addPreparation()"
          />
        </div>
      </FormKit>
      <!-- <DevOnly>
        <Shiki
          v-if="secondary"
          lang="json"
          :code="curValue"
        />
      </DevOnly> -->
    </PFieldset>
  </div>
</template>

<style scoped>
.subGroup {
  @apply flex grow flex-wrap gap-2 py-2;
}
</style>
