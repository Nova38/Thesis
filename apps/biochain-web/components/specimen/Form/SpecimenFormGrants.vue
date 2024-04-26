<script lang="ts" setup>
import { ccbio } from '#imports'
import { useFormKitContext } from '@formkit/vue'


const collapsed = defineModel<boolean>('collapsed', {
  default: true,
})

const grants = useFormKitContext('specimen.grants')

function addGrant() {
  console.log('add grants', grants.value)
  if (!grants.value) return
  const cur = grants.value.node.value as ccbio.Specimen_Loan

  if (!cur) {
    grants.value.node.input({
      '1': new ccbio.Specimen_Grant({
        grantedDate: new pb.Date({}),
      }).toJson(),
    })
  } else {
    grants.value.node.input({
      ...cur,
      [Object.keys(cur).length + 1]: new pb.Specimen_Grant({
        grantedDate: new pb.Date({}),
      }),
    })
  }

}
</script>

<template>
  <div>
    <PFieldset
      legend="Grants"
      :toggleable="true"
      :collapsed
    >
      <FormKit
        type="button"
        label="Add Grant"
        :outer-class="{
            'max-w-[22em]': false,
            'w-full': true,
          }"
        input-class="w-full justify-center items-center flex"
        @click="addGrant"
      />
      <FormKit
        v-slot="{ value }"
        type="group"
        name="grants"
      >
        <template v-for="(v, name) in value">
          <template v-if="typeof name === 'string'">
            <FormKit
              :key="name"
              type="group"
              :name="name"
            >
              <div class="inline-flex flex-wrap gap-2">
                <FormKit
                  type="text"
                  :label="`Grant: ${name}`"
                  name="id"
                  outer-class="min-w-20"
                />
                <FormKit
                  type="text"
                  name="granted_by"
                  label="Granted By"
                />
                <FormKit
                  type="text"
                  name="granted_to"
                  label="Granted To"
                />

                <FormKit
                  type="textarea"
                  name="description"
                  label="Description"
                />
                <SpecimenFormDate v-if="value.grantedDate" name="grantedDate" />
              </div>
            </FormKit>
            <UDivider class="my-2"/>
          </template>
        </template>

      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
