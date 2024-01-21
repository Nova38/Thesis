export default defineAppConfig({
  apiEnd: process.env.API_URL ?? '/api',
  theme: {
    primaryColor: '#ababab',
  },
  ui: {
    checkbox: {
      // ring: "!ring-0 focus:!ring-0 focus:!ring-offset-0",
      background: 'bg-gray-400 dark:bg-slate-50 checked:bg-emerald-500',
    },
  },
})
