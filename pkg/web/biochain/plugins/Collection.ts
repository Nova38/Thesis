export default defineNuxtPlugin({
  name: 'Collection',
  setup() {
    const collectionId = useState('CurrentCollectionId', () => '')

    addRouteMiddleware(
      'collection',
      (to, from) => {
        const id = to.params?.collectionId

        // console.log('Collection middleware', to, from, id)

        if (id && collectionId.value !== id)
          collectionId.value = id.toString()

        // console.log('Collection middleware', collectionId.value)
      },
      {
        global: true,
      },
    )

    return {
      provide: {
        collectionId,
      },
    }
  },
})
