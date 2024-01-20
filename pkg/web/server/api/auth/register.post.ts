import { z } from 'zod'

const userSchema = z.object({
  username: z.string(),
  password: z.string(),
  credentials: z.string().optional(),
  key: z.string().optional(),
  mspId: z.string().optional(),
})

export default eventHandler(async (event) => {
  const body = await readValidatedBody(event, body =>
    userSchema.safeParse(body))

  if (!body.success)
    throw body.error.issues
  const { username, password, credentials, key, mspId } = body.data

  const result = await createUser({
    username,
    password: await hash(password),
    credentials,
    key,
    mspId,
  })

  return {
    value: result,
    message: 'Successfully registered!',
  }
})
