export default defineTask({
  meta: {
    name: 'cc:',
    description: 'Run database migrations',
  },
  run({ payload, context }) {
    console.log('Running DB migration task...')
    return { result: 'Success' }
  },
})
