import { ChaincodeAsUser } from '~/server/utils/useChaincode'
import { GetAllUsers } from '../../utils/db'
import {
  GatewayError,
  EndorseError,
  CommitError,
  SubmitError,
} from '@hyperledger/fabric-gateway'

export default defineTask({
  meta: {
    name: 'Update User IDs from chaincode',
    description: '',
  },
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  async run({ payload, context }) {
    console.log('Running Updates User ID task...')
    const allUsers = await GetAllUsers()

    const results = await Promise.allSettled(
      allUsers.map(async (user) => {
        if (!user) {
          // console.log('User not found:', )
          throw new Error('User not found')
        }
        try {
          // await using cc = await ChaincodeAsUser(user.username)
          const cc = await ChaincodeAsUser(user.username)
          const { user: ccUser } = await cc.utilService.getCurrentUser({})

          const certSubject = atob(ccUser?.userId ?? '')

          const updatedUser = { ...user, ...ccUser, certSubject }
          cc.client.close()
          cc.gateway.close()

          await updateUserByUsername(user.username, updatedUser)
          return {
            key: user.username,
            subject: certSubject,
            id: ccUser?.userId,
          }
        } catch (error) {
          if (
            error instanceof GatewayError ||
            error instanceof SubmitError ||
            error instanceof EndorseError
          ) {
            throw {
              user: user.username,
              type: error.constructor.name,
              message: error.message,
              code: error.code,
              details: error.details,
              cause: { code: error.cause.code },
            }
          }
          if (error instanceof CommitError) {
            throw {
              user: user.username,
              message: error.message,
              code: error.code,
              txID: error.transactionId,
            }
          }
        }
      }),
    )

    return { result: results }
  },
})
