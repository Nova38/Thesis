export default defineAppConfig({
  apiEnd: process.env.API_URL ?? '/api',
  theme: {
    primaryColor: 'red',
  },
  ui: {
    primary: 'cyan',
    icons: {
      dynamic: true,
    },

    checkbox: {
      // ring: "!ring-0 focus:!ring-0 focus:!ring-offset-0",
      background: 'bg-gray-400 dark:bg-slate-50 checked:bg-emerald-500',
    },
    card: {
      header: {
        padding: 'px-4 py-2 sm:px-6',
      },
    },
  },
})
