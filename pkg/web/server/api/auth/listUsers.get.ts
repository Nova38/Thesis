export default defineEventHandler(async () => {
  const keys = await useStorage('.data:auth').getKeys()
  const users = []
  for (const key of keys) {
    const user = await useStorage('.data:auth').getItem(key)
    console.log({ user })
    users.push(user)
  }
  return users
})
