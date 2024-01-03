<template>
  <div>
    <QBtn
      label="Click me"
      @click="
        () => {
          const data = useCustomFetch(
            'https://jsonplaceholder.typicode.com/todos/1',
          );

          console.log('clicked');
          console.log(data);
          // $q.loadingBar.start(1);
          // $q.loadingBar.stop();
        }
      "
    />

    <div>
      {{ $auth.username }}
      <pre>{{ $auth.loggedIn }}</pre>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style></style>
