<script lang="ts" setup>
import { ccbio } from '#imports'
import { useFormKitContext } from '@formkit/vue'

const loans = useFormKitContext('loans')

function addLoan() {
  console.log('addLoan', loans.value)
  if (!loans.value) return
  const cur = loans.value.node.value as ccbio.Specimen_Loan

  if (!cur) {
    loans.value.node.input({
      '1': new ccbio.Specimen_Loan({
        loanedDate: new ccbio.Date({}),
      }).toJson(),
    })
  } else {
    loans.value.node.input({
      ...cur,
      [Object.keys(cur).length + 1]: new ccbio.Specimen_Loan({
        loanedDate: new ccbio.Date({}),
      }),
    })
  }

  // const preparations = loans.value.node.at('$self.preparations')
}
</script>

<template>
  <div>
    <PFieldset
      legend="Loans"
      :toggleable="true"
    >
      <FormKit
        type="button"
        label="Add Loan"
        @click="addLoan"
      />
      <FormKit
        v-slot="{ value }"
        type="group"
        name="loans"
      >
        <template v-for="(v, name) in value">
          <template v-if="typeof name === 'string'">
            <FormKit
              type="group"
              :name="name"
            >
              <div class="inline-flex flex-wrap gap-2">
                <FormKit
                  type="text"
                  :label="`Loan #${name}`"
                  name="id"
                  outer-class="min-w-20"
                />
                <FormKit
                  type="text"
                  name="loanedBy"
                  label="Loaned By"
                />
                <FormKit
                  type="text"
                  name="loanedTo"
                  label="Loaned By"
                />

                <FormKit
                  type="textarea"
                  name="description"
                  label="Description"
                />
                <SpecimenFormDate name="loanedDate" />
              </div>
            </FormKit>
          </template>
        </template>
        <!-- <template v-for="(v, name) in value">
          <template v-if="typeof name === 'string'">
            <FormKit
              type="group"
              :id="name"
              :name="name"
            >
              <div class="inline-flex gap-4">
                <div
                  class="inline-flex flex-row flex-wrap gap-1 justify-evenly"
                >
                  <FormKit
                    type="text"
                    name="id"
                    id="id"
                    label="id"
                    outer-class="min-w-20"
                  />
                </div>

                <div class="inline-flex flex-wrap gap-2"></div>
              </div>
            </FormKit>
          </template>
        </template> -->
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
