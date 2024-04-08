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
        <div class="inline-flex gap-4 flex-wrap">
          <div class="inline-flex flex-row flex-wrap gap-1 justify-evenly">
            <FormKit
              type="text"
              id="sex"
              name="sex"
              label="sex"
              outer-class="min-w-20"
            />
            <FormKit
              type="text"
              id="notes"
              outer-class="min-w-20"
            />
          </div>
          <div class="inline-flex flex-wrap gap-2">
            <FormKit
              type="number"
              id="weight"
              number
              name="weight"
              label="weight"
              outer-class="min-w-20"
            />
            <FormKit
              type="text"
              id="weightUnits"
              name="weightUnits"
              label="Weight Units"
              outer-class="min-w-20"
            />
          </div>
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
        <FormKit
          type="button"
          label="Add Preparation"
          @click="addPreparation()"
        />
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

<style scoped></style>
