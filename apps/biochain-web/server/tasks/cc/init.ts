export default defineTask({
  meta: {
    name: 'cc:init',
    description: 'Initialize the collection',
  },
  run({ payload, context }) {
    console.log('Running DB migration task...')
    return { result: 'Success' }
  },
})
