export default defineNitroPlugin((nitro) => {
  // const startedAt = Date.now()
  // const count = 0
  // nitro.hooks.hook('error', async (error, { event }) => {
  //   console.error(`${event?.path} Application error:`, error)
  // })
  // nitro.hooks.hook('request', (event) => {
  //   // useEvent().context.count = count
  //   // console.log(`count: ${count} - on request: ${event.path}`)
  //   count = count + 1
  // })
  // nitro.hooks.hook('beforeResponse', (event, { body }) => {
  //   console.log('on response', event.path, { body })
  // })
  // // nitro.hooks.hook('afterResponse', (event, { body }) => {
  // //     console.log('on after response', event.path, { body })
  // // })
  // nitro.hooks.hookOnce('close', async () => {
  //   // Will run when nitro is closed
  //   console.log('Closing nitro server...')
  //   await new Promise(resolve => setTimeout(resolve, 500))
  //   console.log('Task is done!')
  // })
  // console.log(`Nitro plug loaded at ${startedAt}`)
})
