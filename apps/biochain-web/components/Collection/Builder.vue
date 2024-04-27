<script lang="ts" setup></script>

<template>
  <div>
    <UCard>
      <FormKit
        type="form"
        class="w-full"
        :actions="false"
      >
        <FormKit
          type="multi-step"
          class="w-full"
          tab-style="progress"
        >
          <FormKit
            type="step"
            name="Collection Information"
          >
            <!-- component for example brevity. -->
            <StepOne />
          </FormKit>

          <FormKit
            type="step"
            name="Roles"
          >
            <!-- component for example brevity. -->
            <StepTwo />
          </FormKit>

          <FormKit
            type="step"
            name="Users"
          >
            <!-- component for example brevity. -->
            <StepThree />

            <!-- using step slot for submit button-->
            <template #stepNext>
              <FormKit type="submit" />
            </template>
          </FormKit>
        </FormKit>
      </FormKit>
    </UCard>
  </div>
</template>

<style scoped></style>
