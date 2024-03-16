import { z } from 'zod'

const userSchema = z.object({
  password: z.string(),
  username: z.string(),
})

export default defineEventHandler(async (event) => {
  const session = await useAuthSession(event)

  if (session.data.username) {
    throw createError({
      message: 'Already logged in!',
      statusCode: 418,
    })
  }

  const result = await readValidatedBody(event, body =>
    userSchema.safeParse(body))

  if (!result.success)
    throw result.error.issues

  const { password, username } = result.data

  const user = await findUserByUsername(username)

  if (!user) {
    throw createError({
      message: 'username not found! Please register.',
      statusCode: 401,
    })
  }
  if (!user.password || user.password !== (await hash(password))) {
    throw createError({
      message: 'Incorrect password!',
      statusCode: 401,
    })
  }
  await session.update({
    username: user.username,
  })
  return session
})
