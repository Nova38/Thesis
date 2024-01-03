<template>
  <QCard class="p-4">
    <FormKit
      id="registration-example"
      v-slot="{ value }"
      type="form"
      :form-class="submitted ? 'hide' : 'show'"
      submit-label="Register"
      :actions="false"
      @submit="submitHandler"
    >
      <h1 class="text-2xl font-bold mb-2">Login!</h1>

      <div class="flex flex-row gap-4">
        <FormKit
          type="text"
          name="username"
          label="Username"
          placeholder=""
          validation="required"
        />
        <FormKit
          type="password"
          name="password"
          label="Password"
          validation="required|"
          placeholder=""
        />
      </div>

      <FormKit type="submit" label="Login" class="" />
      <pre wrap>{{ value }}</pre>
    </FormKit>
  </QCard>
</template>

<script setup lang="ts">
const submitHandler = async (data) => {
  // Let's pretend this is an ajax request:
  await new Promise((r) => setTimeout(r, 1000));

  await authLogin(data.username, data.password);
};
</script>
