import { ChaincodeAsUser } from '~/server/utils/useChaincode'
import { GetAllUsers, type User } from '../../utils/db'

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
          console.log('User not found:', key)
          throw new Error('User not found')
        }
        // await using cc = await ChaincodeAsUser(user.username)
        const cc = await ChaincodeAsUser(user.username)
        const { user: ccUser } = await cc.utilService.getCurrentUser({})
        const updatedUser = { ...user, ...ccUser }
        console.log('Updated User:', updatedUser)
        cc.client.close()
        cc.gateway.close()

        return await updateUserByUsername(user.username, updatedUser)
      }),
    )

    return { result: results }
  },
})
