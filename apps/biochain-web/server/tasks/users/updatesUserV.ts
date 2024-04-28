import certToID from '~/utils/cc/certToID'
import type { User } from '../../utils/db'
import { X509Certificate } from 'crypto'

export default defineTask({
  meta: {
    name: 'UpdateUserIds:payload',
    description: 'Updates The User IDs calculated from the certificate',
  },
  async run({ payload, context }) {
    console.log('Running Updates User ID task...')
    const storage = useStorage('.data:auth')

    const users = await Promise.all(
      (await storage.getKeys()).map(async (key) => {
        try {
          const user = await storage.getItem<User>(key)
          if (!user) {
            console.log('User not found:', key)
            return { key, subject: 'no-credentials' }
          }

          if (!user.credentials) {
            console.log('User has no credentials:', key)
            return { key, subject: 'no-credentials' }
          }
          const cert = new X509Certificate(user.credentials)
          if (!cert) {
            console.error('User Cert failed to load:', key)
            return { key, subject: 'no-credentials' }
          }

          const certSubject = certToID(cert)
          const id = btoa(certSubject)

          const updatedUser = { ...user, id: id, certSubject }
          await storage.setItem(key, updatedUser)
          return { key, subject: certSubject, id }
        } catch (error) {
          console.error('User Errored:', key, error)
          return { key, subject: 'no-credentials' }
        }
      }),
    )

    return { result: users }
  },
})
