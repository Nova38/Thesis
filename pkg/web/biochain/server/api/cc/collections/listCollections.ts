import { GatewayError } from '@hyperledger/fabric-gateway'
import { common } from 'saacs'

// export default defineEventHandler
export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event)

  try {
    return await cc.service.getCollectionsList()
  }
  catch (error) {
    // console.error('Error in listCollections', error)

    if (error instanceof GatewayError) {
      console.log(error.cause.details)

      // createError({
      //   cause: error,
      // })

      return new common.generic.GetCollectionsListResponse({
        collections: [
          {
            collectionId: 'TestingID',
            name: 'TestingName',
          },
        ],
      })
    }
  }
})
