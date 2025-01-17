import * as grpc from '@grpc/grpc-js'
import users from './users.json' assert { type: 'json' }

// type assert for json import
import {
  connect,
  ConnectOptions,
  Identity,
  signers,
} from '@hyperledger/fabric-gateway'
import connectionProfile from '../../../../network/organizations/peerOrganizations/org1.example.com/connection-org1.json' assert { type: 'json' }

import * as crypto from 'crypto'

import { promises as fs } from 'fs'

import { createRegistry, IMessageTypeRegistry } from '@bufbuild/protobuf'
import { GenericServiceClient } from '../gen/chaincode/common/generic_pb_gateway.js'
// import { auth, ccbio, common, sample } from "../index.js";

import { auth, objects } from '../gen/auth/v1/index.js'
import { generic, reference } from '../gen/chaincode/common/index.js'
import { ccbio } from '../gen/index.js'
import { sample } from '../gen/index.js'
import { GlobalRegistry } from '../gen/global_reg.js'

export const gr: IMessageTypeRegistry = createRegistry()
export const CLIENT = await newGRPCClient()

export async function newGRPCClient() {
  const peer = connectionProfile.peers['peer0.org1.example.com']
  const tlsCredentials = grpc.credentials.createSsl(
    Buffer.from(peer.tlsCACerts.pem),
  )

  const client = await new grpc.Client(
    peer.url.split('//')[1],
    tlsCredentials,
    {
      'grpc.ssl_target_name_override':
        peer.grpcOptions['ssl-target-name-override'],
    },
  )

  return client
}
export async function BuildGateway({ userIdex }: { userIdex: number }) {}
export function BuildContract(contract: any) {
  const service = new GenericServiceClient(contract, createRegistry())
  return service
}

const GetIdentity = async ({ userIdex }: { userIdex: number }) => {
  const certPath = (path: string) => `../../${path}`

  const user = users.identities[userIdex]

  const credentials = await fs.readFile(certPath(user.clientSignedCert))
  const identity: Identity = { mspId: 'Org1MSP', credentials }

  const privateKeyPem = await fs.readFile(certPath(user.clientPrivateKey))

  const privateKey = crypto.createPrivateKey(privateKeyPem)
  console.log('Identity:', {
    user,
    creds: credentials.toString(),
    // identity,
    private: privateKeyPem.toString(),
  })
  const signer = signers.newPrivateKeySigner(privateKey)

  return { identity, signer }
}

export const GetGateway = async ({ userIdex }: { userIdex: number }) => {
  const { identity, signer } = await GetIdentity({ userIdex })

  const client = await newGRPCClient()
  const gateway = connect({ client, identity, signer })

  return {
    gateway,
    client,
    identity,
    signer,
    [Symbol.asyncDispose]: async () => {
      console.log('closing gateway')
      await gateway.close()
      await client.close()
    },
  }
}

export const GetService = async ({
  userIdex,
  channel,
  contractName,
}: {
  userIdex: number
  channel: string
  contractName: string
}) => {
  const connection = await GetGateway({ userIdex })

  const network = await connection.gateway.getNetwork(channel)
  const contract = await network.getContract(contractName)
  const service = new GenericServiceClient(contract, GlobalRegistry)
  return {
    connection,
    contract,
    service,
  }
}
