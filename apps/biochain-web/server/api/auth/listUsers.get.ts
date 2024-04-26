import type { User } from '~/server/utils/db'

export default defineEventHandler(async () => {
  const users = await Promise.all(
    (await useStorage('.data:auth').getKeys()).map((key) =>
      useStorage('.data:auth').getItem(key),
    ),
  )

  return users.map((user) => {
    if (!user) return {}
    const u = user as User
    return {
      key: u?.username,
      username: u?.username,
      id: u?.id,
      mspId: u?.mspId,
    }
  })
})
