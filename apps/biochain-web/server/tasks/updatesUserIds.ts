import certToID from '~/utils/cc/certToID'
import { User } from '../utils/db'
import { X509Certificate } from 'crypto'

export default defineTask({
  meta: {
    name: 'UpdateUserIds:payload',
    description: 'Updates The User IDs calculated from the certificate',
  },
  async run({ payload, context }) {
    console.log('Running Updates User ID task...')
    const storage = useStorage('.data:auth')

    const keys = await storage.getKeys()
    console.log('Keys:', keys)
    for (const key of keys) {
      const user = await storage.getItem<User>(key)
      if (!user) {
        console.log('User not found:', key)
        continue
      }

      if (!user.credentials) {
        console.log('User has no credentials:', key)
        continue
      }

      try {
        const cert = new X509Certificate(user.credentials)
        if (!cert) {
          console.error('User Cert failed to load:', key)
          continue
        }

        console.log(`${user.username}:`)

        console.log(certToID(cert))
      } catch (error) {
        console.error('User Errored:', key, error)
        continue
      }
    }

    return { result: { keys } }
  },
})
