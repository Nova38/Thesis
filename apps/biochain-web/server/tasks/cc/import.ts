// export default defineTask({
//   meta: {
//     name: 'cc:',
//     description: 'Run database migrations',
//   },
//   run({ payload, context }) {
//     const specimens = l.entries
//       .map(s => pb.Specimen.fromJsonString(JSON.stringify(s)))
//       .map(s => ToItem(s))
//       .map(i => new pb.CreateRequest({ item: i }))

//     console.log('Running DB migration task...')

//     return { result: 'Success' }
//   },
// })
