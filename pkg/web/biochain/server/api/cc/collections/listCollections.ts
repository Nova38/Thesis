import { GatewayError } from '@hyperledger/fabric-gateway'
import { GetCollectionsListResponse } from '~/lib/pb/types_pb'

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

      return new GetCollectionsListResponse({
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
