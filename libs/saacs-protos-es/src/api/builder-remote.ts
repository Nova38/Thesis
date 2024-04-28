import * as grpc from '@grpc/grpc-js'

import remote from './remote.json' assert { type: 'json' }
import { ofetch } from 'ofetch'
import { createRegistry, IMessageTypeRegistry } from '@bufbuild/protobuf'
import { GenericServiceClient } from '../gen/chaincode/common/generic_pb_gateway.js'
// import { auth, ccbio, common, sample } from "../index.js";
import {
  connect,
  ConnectOptions,
  Identity,
  signers,
} from '@hyperledger/fabric-gateway'
import { auth, objects } from '../gen/auth/v1/index.js'
import { generic, reference } from '../gen/chaincode/common/index.js'
import { ccbio } from '../gen/index.js'
import { sample } from '../gen/index.js'
import * as crypto from 'crypto'

import { promises as fs } from 'fs'

//

const baseurl = 'https://api-biochain.ittc.ku.edu/api'

export const GlobalRegistry: IMessageTypeRegistry = createRegistry(
  ...auth.allMessages,
  ...objects.allMessages,
  ...generic.allMessages,
  ...reference.allMessages,
  ...ccbio.allMessages,
  ...sample.allMessages,
)
export const CLIENT = await newGRPCClient()

export async function newGRPCClient() {
  const peer = remote.peers['org1-peer0.default']
  const tlsCredentials = grpc.credentials.createSsl(
    Buffer.from(peer.tlsCACerts.pem, 'utf-8'),
  )

  const client = await new grpc.Client(
    'peer0-org1.129.237.123.11.nip.io:443',
    tlsCredentials,
  )

  return client
}

export function BuildContract(contract: any) {
  const service = new GenericServiceClient(contract, createRegistry())
  return service
}

const GetUserIdentity = (username: string) => {
  // const certPath = (path: string) => `../../${path}`;

  // const credentials = await fs.readFile(
  //     certPath(remote.organizations.Org1MSP.users[username]),
  // );

  const user = GetUser(username)
  // console.log(user);
  if (!user) {
    throw new Error('User not found')
  }
  // console.log({ username, id: user.credentials, key: user.key });

  const identity: Identity = {
    mspId: 'Org1MSP',
    credentials: Buffer.from(user.credentials),
  }

  const signer = signers.newPrivateKeySigner(
    crypto.createPrivateKey(Buffer.from(user.key)),
  )

  return { identity, signer, private_key: user.key }
}

const getMe = () => {
  const identity = {
    mspId: 'Org1MSP',
    credentials: Buffer.from(
      '\n-----BEGIN CERTIFICATE-----\nMIICVDCCAfugAwIBAgIUAwUUwM/0id5g1gb24QB8d4xN3S8wCgYIKoZIzj0EAwIw\najELMAkGA1UEBhMCRVMxETAPBgNVBAcTCEFsaWNhbnRlMREwDwYDVQQJEwhBbGlj\nYW50ZTEZMBcGA1UEChMQS3VuZyBGdSBTb2Z0d2FyZTENMAsGA1UECxMEVGVjaDEL\nMAkGA1UEAxMCY2EwHhcNMjMwNzA1MTY1MzA0WhcNMjQwNzA2MDcwMzAwWjAhMQ4w\nDAYDVQQLEwVhZG1pbjEPMA0GA1UEAxMGYXRraW5zMFkwEwYHKoZIzj0CAQYIKoZI\nzj0DAQcDQgAESI8A9uNjXcHERmOEyJgCfud2+1717dVi648TIgMBX85gj8N6JQLK\nB2sdAW780gSRkgODmJPszPxs5hDQbrItPaOBxzCBxDAOBgNVHQ8BAf8EBAMCB4Aw\nDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUyJTXiGBlS1+ALdxUw4v1wuWcCtowKwYD\nVR0jBCQwIoAgQr2UdBKNxJyOczPmEBuo/6giqFA13kNru5Oc1fZFoU4wWAYIKgME\nBQYHCAEETHsiYXR0cnMiOnsiaGYuQWZmaWxpYXRpb24iOiIiLCJoZi5FbnJvbGxt\nZW50SUQiOiJhdGtpbnMiLCJoZi5UeXBlIjoiYWRtaW4ifX0wCgYIKoZIzj0EAwID\nRwAwRAIgI1bLXJLS/lLjgBPDRGoW4N5n4olX1QfvHFHaYYsKA1QCIGr1Sd3eW1d1\nVZ/0dwO63nteX2F/JRlzuJtZw4dpo4C9\n-----END CERTIFICATE-----\n',
    ),
  }

  const signer = signers.newPrivateKeySigner(
    crypto.createPrivateKey(
      '\n-----BEGIN PRIVATE KEY-----\nMIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgG/ZAo5b69qT6Y6rM\n/K1p5Z/6ygpNdijIfMbnlr4oAUShRANCAARIjwD242NdwcRGY4TImAJ+53b7XvXt\n1WLrjxMiAwFfzmCPw3olAsoHax0BbvzSBJGSA4OYk+zM/GzmENBusi09\n-----END PRIVATE KEY-----\n',
    ),
  )

  return { identity, signer }
}

export const GetGateway = async (username: string) => {
  // const { identity, signer } = await GetIdentity({ userIdex });

  // let { identity, signer } = getMe();
  let { identity, signer, private_key } = GetUserIdentity(username)
  const client = await newGRPCClient()
  const gateway = connect({
    client,
    identity,
    signer,
  })

  return {
    gateway,
    client,
    identity,
    signer,
    private_key,
    [Symbol.asyncDispose]: async () => {
      console.warn('closing gateway')
      await gateway.close()
      await client.close()
    },
  }
}

function GetUser(username: string) {
  const users = remote.organizations.Org1MSP.users
  for (const [k, value] of Object.entries(users)) {
    // console.log(`${key}: ${value}`);
    if (username == k) {
      const username = k
      const name = value.name
      const password = value.password
      const credentials = value.cert.pem
      const key = value.key.pem
      const mspId = 'Org1MSP'

      return {
        username,
        name,
        password,
        credentials,
        key,
        mspId,
      }
    }
  }
}
export const GetService = async ({
  username,
  channel,
  contractName,
}: {
  username: string
  channel: string
  contractName: string
}) => {
  const connection = await GetGateway(username)

  const network = await connection.gateway.getNetwork(channel)
  const contract = await network.getContract(contractName)
  const service = new GenericServiceClient(
    contract,
    createRegistry(
      ...auth.allMessages,
      ...objects.allMessages,
      ...generic.allMessages,
      ...reference.allMessages,
      ...ccbio.allMessages,
      ...sample.allMessages,
    ),
  )
  return {
    connection,
    contract,
    service,
  }
}
