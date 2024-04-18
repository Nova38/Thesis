import { defineWorkspace } from 'vitest/config'

export default defineWorkspace([
  "./libs/saacs-protos-es/vitest.config.ts",
  "./libs/client/vitest.config.ts",
  // "./apps/biochain-web/vitest.config.ts",
])
