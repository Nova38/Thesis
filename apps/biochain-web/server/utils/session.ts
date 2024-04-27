import type { H3Event, SessionConfig } from 'h3'

import crypto from 'uncrypto'

export const sessionConfig: SessionConfig = useRuntimeConfig().auth || {}

export interface AuthSession {
  id: string
  username: string
}

export async function useAuthSession(event: H3Event) {
  const session = await useSession<AuthSession>(event, sessionConfig)
  return session
}

export async function requireAuthSession(event: H3Event) {
  const session = await useAuthSession(event)
  if (!session.data.username) {
    throw createError({
      message: 'Not Authorized',
      statusCode: 401,
    })
  }
  return session
}

export async function hash(str: string) {
  const msgUint8 = new TextEncoder().encode(str)
  const hashBuffer = await crypto.subtle.digest('SHA-512', msgUint8)
  const hashArray = Array.from(new Uint8Array(hashBuffer))
  const hashHex = hashArray.map((b) => b.toString(16).padStart(2, '0')).join('')
  return hashHex
}

export async function requireAdminSession(event: H3Event) {
  const session = await useAuthSession(event)
  if (!session.data.username) {
    throw createError({
      message: 'Not Authorized',
      statusCode: 401,
    })
  }
  return session
}
