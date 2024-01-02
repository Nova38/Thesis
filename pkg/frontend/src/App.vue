<template>
  <router-view />
</template>

<script setup lang="ts">
import '@ag-grid-community/styles/ag-grid.css';
import '@ag-grid-community/styles/ag-theme-alpine.css';

import '@ag-grid-community/styles/ag-theme-material.css';
import { ClientSideRowModelModule } from '@ag-grid-community/client-side-row-model';

import { ModuleRegistry } from '@ag-grid-community/core';

ModuleRegistry.registerModules([ClientSideRowModelModule]);

import 'virtual:uno.css';
</script>
