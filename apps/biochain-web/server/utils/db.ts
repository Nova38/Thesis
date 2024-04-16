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

export interface UserChaincodeIdentity {
  credentials: string
  key: string
  mspId: string
  userId: string
  username: string
}

export async function findUserByUsername(username: string): Promise<User> {
  const storage = useStorage('.data:auth')
  const key = getUserKey(username)

  const user = await storage.getItem(key)
  if (!user) throw createError({ message: 'User not found!', statusCode: 404 })

  return user as User
}

export async function findUserById(id: string): Promise<User> {
  const storage = useStorage('.data:auth')
  const key = getUserKey(id)

  const user = await storage.getItem(key)
  if (!user) throw createError({ message: 'User not found!', statusCode: 404 })

  return user as User
}

export async function createUser(user: Partial<User>) {
  const key = getUserKey(user.username!)
  if (await useStorage('.data:auth').hasItem(key)) {
    throw createError({
      message: 'Username already exists!',
      statusCode: 409,
    })
  }
  console.log('createUser', user)
  return await useStorage('.data:auth').setItem(key, user)
}

export async function updateUserByUsername(
  username: string,
  updates: Partial<User>,
) {
  const storage = useStorage('.data:auth')
  const user = await findUserByUsername(username)
  const key = getUserKey(user.username)

  return await storage.setItem(key, {
    ...user,
    ...updates,
  })
}

function getUserKey(username: string) {
  return `db:usersDB:${encodeURIComponent(username)}`
}
