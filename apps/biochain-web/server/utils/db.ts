export interface User {
  createdAt: string
  credentials: string

  id: string

  certSubject: string

  key: string

  mspId: string
  password: string

  userId: string
  username: string
}

export interface UserInfo {
  username: string
  password: string
}

export interface UserChaincodeIdentity {
  credentials: string
  key: string
  mspId: string
  userId: string
}

export async function getKeys() {
  return await useStorage('.data:auth').getKeys()
}

export async function findUserByUsername(username: string): Promise<User> {
  const storage = useStorage('.data:auth')
  const allKeys = await storage.getKeys()
  const user = await storage.getItem(username)
  if (!user) throw createError({ message: 'User not found!', statusCode: 403 })

  return user as User
}

export async function findUserById(id: string): Promise<User> {
  const storage = useStorage('.data:auth')
  // const key = getUserKey(id)

  const user = await storage.getItem(id)
  if (!user) throw createError({ message: 'User not found!', statusCode: 404 })

  return user as User
}

export async function createUser(user: Partial<User>) {
  if (!user.username) {
    throw createError({
      message: 'Username is required!',
      statusCode: 400,
    })
  }
  if (await useStorage('.data:auth').hasItem(user.username)) {
    throw createError({
      message: 'Username already exists!',
      statusCode: 409,
    })
  }
  console.log('createUser', user)
  return await useStorage('.data:auth').setItem(user.username, user)
}

export async function updateUserByUsername(
  username: string,
  updates: Partial<User>,
) {
  const storage = useStorage('.data:auth')
  const user = await findUserByUsername(username)
  // const key = getUserKey(user.username)

  return await storage.setItem(user.username, {
    ...user,
    ...updates,
  })
}

export async function GetAllUsers() {
  const storage = useStorage('.data:auth')

  return await Promise.all(
    (await storage.getKeys()).map(async (key) => {
      return await useStorage('.data:auth').getItem<User>(key)
    }),
  )
}
