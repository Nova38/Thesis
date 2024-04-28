import { z } from 'zod'

export default defineEventHandler(async (event) => {
  const session = await useAuthSession(event)

  if (session.data.username) {
    throw createError({
      message: 'Already logged in!',
      statusCode: 418,
    })
  }

  const result = await readValidatedBody(
    event,
    z.object({
      password: z.string(),
      username: z.string(),
    }).parse,
  )

  const { password, username } = result

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
