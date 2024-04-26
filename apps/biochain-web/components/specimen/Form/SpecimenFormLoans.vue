<script lang="ts" setup>
import { ccbio } from '#imports'
import { useFormKitContext } from '@formkit/vue'

const loans = useFormKitContext('specimen.loans')

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
        :outer-class="{
            'max-w-[22em]': false,
            'w-full': true,
          }"
        input-class="w-full justify-center items-center flex"
      />
      <FormKit
        v-slot="{ value }"
        type="group"
        name="loans"
      >
        <template v-for="(v, name) in value">
          <template v-if="typeof name === 'string'">
            <FormKit
              :key="name"
              type="group"
              #default="{value}"
              :name="name"
            >
              <div class="inline-flex flex-wrap gap-2">
                <FormKit
                  type="text"
                  :label="`Loan: ${name}`"
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
                <SpecimenFormDate v-if="value?.loanedDate" name="loanedDate" />
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
