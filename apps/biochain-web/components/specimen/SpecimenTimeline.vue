<script setup lang="ts">
import type { pb } from '#imports'
import { ccbio } from '#imports'
import { type PlainMessage } from '@bufbuild/protobuf'

const props = defineProps<{
  history: PlainMessage<pb.SpecimenHistory>
}>()
</script>

<template></template>
