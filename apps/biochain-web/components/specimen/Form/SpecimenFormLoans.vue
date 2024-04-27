<script lang="ts" setup>
import type { pb } from '#imports'
import { ccbio } from '#imports'
import { useFormKitContext } from '@formkit/vue'
import { type PlainMessage } from '@bufbuild/protobuf'
import { toPlainMessage } from '@bufbuild/protobuf'

const loans = useFormKitContext('specimen.loans')

function addLoan() {
  console.log('addLoan', loans.value)
  if (!loans.value) return
  const cur = loans.value.node.value as PlainMessage<pb.Specimen_Loan>

  if (!cur) {
    loans.value.node.input({
      '1': new ccbio.Specimen_Loan({
        loanedDate: new ccbio.Date({}),
      }),
    })
  } else {
    loans.value.node.input({
      ...cur,
      [Object.keys(cur).length + 1]: toPlainMessage(
        new ccbio.Specimen_Loan({
          loanedDate: new ccbio.Date({}),
        }),
      ),
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
        v-slot="{ value, state: { disabled } }"
        type="group"
        name="loans"
      >
        <UButton
          block
          :disabled="disabled"
          label="Add Preparation"
          @click="addLoan"
        />
        <template v-for="(v, name) in value">
          <template v-if="typeof name === 'string'">
            <FormKit
              :key="name"
              v-slot="{ value: loan }"
              type="group"
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
                <SpecimenFormDate
                  v-if="loan?.loanedDate"
                  name="loanedDate"
                />
              </div>
            </FormKit>
          </template>
        </template>
      </FormKit>
    </PFieldset>
  </div>
</template>

<style scoped></style>
