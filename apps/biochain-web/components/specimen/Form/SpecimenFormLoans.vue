<script lang="ts" setup>
import { useFormKitContext } from '@formkit/vue'

const loans = useFormKitContext('loans')
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
      />
      <FormKit
        type="group"
        id="loans"
        name="loans"
        #default="{ value }"
      >
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
