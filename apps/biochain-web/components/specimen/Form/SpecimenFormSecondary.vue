<script lang="ts" setup>
// import { useFormKitContext } from '#imports'

import { ccbio } from '#imports'
import { useFormKitContext } from '@formkit/vue'

const secondary = useFormKitContext('secondary', (node) => {
  curValue.value = JSON.stringify(node.value, null, 2)
  node.node.on('settled.deep', (v) => {
    curValue.value = JSON.stringify(v.origin.value, null, 2)
  })
})

const curValue = ref('{}')

watch(secondary, (v) => {
  console.log('secondary', v)
  if (!v) return
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
      :toggleable="true"
    >
      <FormKit
        type="group"
        id="secondary"
        name="secondary"
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
            type="select"
            id="sex"
            name="sex"
            label="Sex"
            outer-class="min-w-20"
            :options="[
              {
                label: 'Not Specified',
                value: ccbio.Specimen_Secondary_SEX.SEX_UNDEFINED,
              },
              {
                label: 'Unknown',
                value: ccbio.Specimen_Secondary_SEX.SEX_UNKNOWN,
              },
              {
                label: 'Atypical',
                value: ccbio.Specimen_Secondary_SEX.SEX_ATYPICAL,
              },
              {
                label: 'Female',
                value: ccbio.Specimen_Secondary_SEX.SEX_FEMALE,
              },
              {
                label: 'Male',
                value: ccbio.Specimen_Secondary_SEX.SEX_MALE,
              },
            ]"
          />
          <FormKit
            type="select"
            id="age"
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
            type="group"
            id="preparations"
            name="preparations"
            #default="{ value }"
          >
            <template v-for="(v, name) in value">
              <template v-if="typeof name === 'string'">
                <FormKit
                  type="group"
                  :id="name"
                  :name="name"
                >
                  <FormKit
                    type="text"
                    id="verbatim"
                    :label="`Preparation #${name}`"
                    name="verbatim"
                    outer-class="min-w-20"
                  />
                </FormKit>
              </template>
            </template>
          </FormKit>
        </div>

        <div class="inline-flex flex-wrap gap-2">
          <FormKit
            outer-class="w-full flex-grow"
            type="button"
            label="Add Preparation"
            @click="addPreparation()"
          />
        </div>
      </FormKit>
      <DevOnly>
        <Shiki
          v-if="secondary"
          lang="json"
          :code="curValue"
        />
      </DevOnly>
    </PFieldset>
  </div>
</template>

<style scoped>
.subGroup {
  @apply flex grow flex-wrap gap-2 py-2;
}
</style>
