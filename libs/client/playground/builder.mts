/* eslint-disable node/prefer-global/buffer */
// type assert for json import
import * as crypto from 'node:crypto'
import { promises as fs } from 'node:fs'
import type {
  Identity,
} from '@hyperledger/fabric-gateway'
import {
  ConnectOptions,
  connect,
  signers,
} from '@hyperledger/fabric-gateway'

import * as grpc from '@grpc/grpc-js'

import type { IMessageTypeRegistry } from '@bufbuild/protobuf'
import { createRegistry } from '@bufbuild/protobuf'
import { GlobalRegistry, chaincode } from '@saacs/saacs-pb'
import connectionProfile from '../../../infra/network/organizations/peerOrganizations/org1.example.com/connection-org1.json' assert { type: 'json' }
// import { auth, ccbio, common, sample } from "../index.js";

import users from './users.json' assert { type: 'json' }

export const gr: IMessageTypeRegistry = createRegistry(

)
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

export function BuildContract(contract: any) {
  const service = new chaincode.chaincode.ItemServiceClient(contract, createRegistry())
  return service
}

async function GetIdentity({ userIdex }: { userIdex: number }) {
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

export async function GetGateway({ userIdex }: { userIdex: number }) {
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

export async function GetService({
  userIdex,
  channel,
  contractName,
}: {
  userIdex: number
  channel: string
  contractName: string
}) {
  const connection = await GetGateway({ userIdex })

  const network = await connection.gateway.getNetwork(channel)
  const contract = await network.getContract(contractName)

  const service = new chaincode.chaincode.ItemServiceClient(
    contract,
    GlobalRegistry,
  )

  const utilService = new chaincode.utils.UtilsServiceClient(
    contract,
    GlobalRegistry,
  )
  return {
    connection,
    contract,
    service,
    utilService,
  }
}

const s = await GetService({ userIdex: 0, channel: 'default', contractName: 'saacs-caas' })
const r = await s.utilService.getCurrentUser({})
console.log(r)
