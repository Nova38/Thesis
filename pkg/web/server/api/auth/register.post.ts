import { z } from 'zod'

const userSchema = z.object({
  credentials: z.string().optional(),
  key: z.string().optional(),
  mspId: z.string().optional(),
  password: z.string(),
  username: z.string(),
})

export default eventHandler(async (event) => {
  const body = await readValidatedBody(event, body =>
    userSchema.safeParse(body))

  if (!body.success)
    throw body.error.issues
  const { credentials, key, mspId, password, username } = body.data

  const result = await createUser({
    credentials,
    key,
    mspId,
    password: await hash(password),
    username,
  })

  return {
    message: 'Successfully registered!',
    value: result,
  }
})
