import { Buffer } from 'node:buffer'
import * as crypto from 'node:crypto'
import type { IMessageTypeRegistry } from '@bufbuild/protobuf'
import type { H3Event } from 'h3'

import { createRegistry } from '@bufbuild/protobuf'
import { Client, credentials } from '@grpc/grpc-js'
import { connect, signers } from '@hyperledger/fabric-gateway'
import type { User } from './db'

import { sessionConfig } from './session'
import { auth, ccbio, common, sample } from '~/lib'
import { GlobalRegistry } from '~/lib/pb/global_reg'

export interface FabricConfig {
  chaincode: {
    chaincode: string
    channel: string
  }
  peer: {
    grpcOptions: {
      'ssl-target-name-override': string
    }
    tlsCACerts: {
      pem: string
    }
    url: string
  }
  public: {
    credentials: string
    key: string
    mspId: string
  }
}

export const fabricConfig: FabricConfig = useRuntimeConfig().fabric || {}
export function userToIdentity(user: User) {
  if (!user.mspId || !user.credentials) {
    throw createError({
      message: 'Users mspId or signCert not found!',
      statusCode: 404,
    })
  }

  return {
    credentials: Buffer.from(user.credentials),
    mspId: user.mspId,
  }
}

export function userToPrivateKey(user: User) {
  if (!user.key) {
    throw createError({
      message: 'Users private key not found!',
      statusCode: 404,
    })
  }

  const privateKey = crypto.createPrivateKey(user.key)

  return {
    privateKey,
  }
}

export async function newGRPCClient() {
  const tlsCredentials = credentials.createSsl(
    Buffer.from(fabricConfig.peer.tlsCACerts.pem),
  )

  const client = new Client(
    fabricConfig.peer.url.split('//')[1],
    tlsCredentials,
    {
      'grpc.ssl_target_name_override':
        fabricConfig.peer.grpcOptions['ssl-target-name-override'],
    },
  )

  return client
}

async function BuildIdentity(event: H3Event) {
  try {
    const session = await useSession<AuthSession>(event, sessionConfig)

    if (!session.data.username) {
      throw createError({
        message: 'Not Authorized',
        statusCode: 401,
      })
    }

    const user = await findUserByUsername(session.data.username)
    const identity = userToIdentity(user)
    const { privateKey } = userToPrivateKey(user)
    const signer = signers.newPrivateKeySigner(privateKey)

    return {
      identity,
      signer,
    }
  } catch (error) {
    const publicUser: User = {
      createdAt: '',
      credentials: fabricConfig.public.credentials,
      id: '',
      key: fabricConfig.public.key,
      mspId: fabricConfig.public.mspId,
      password: '',
      userId: '',
      username: '',
    }
    const identity = userToIdentity(publicUser)
    const { privateKey } = userToPrivateKey(publicUser)
    const signer = signers.newPrivateKeySigner(privateKey)

    return {
      identity,
      signer,
    }
  }
}

async function BuildGateway(event: H3Event) {
  const { identity, signer } = await BuildIdentity(event)

  const client = await newGRPCClient()
  const gateway = connect({ client, identity, signer })

  return {
    [Symbol.asyncDispose]: async () => {
      console.log('closing gateway')
      gateway.close()
      client.close()
    },
    client,
    gateway,
  }
}

export async function useChaincode<T extends H3Event>(event: T) {
  const connection = await BuildGateway(event)

  const network = connection.gateway.getNetwork(fabricConfig.chaincode.channel)
  const contract = network.getContract(fabricConfig.chaincode.chaincode)

  const service = new common.generic.GenericServiceClient(
    contract,
    GlobalRegistry,
  )

  return {
    connection,
    contract,
    service,
  }
}
