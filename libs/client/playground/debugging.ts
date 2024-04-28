/* eslint-disable unused-imports/no-unused-vars */
// Path to crypto materials: On Windows machine - %USERPROFILE%.vscode\extensions\spydra.hyperledger-fabric-debugger-{version}\fabric\local\organizations\peerOrganizations\org1.debugger.com. On Linux and Mac, it will be - ~/.vscode/extensions/spydra.hyperledger-fabric-debugger-{version}/fabric/local/organizations/peerOrganizations/org1.debugger.com.
// User key directory: users/Org1Admin/msp/keystore
// User certificate path: users/Org1Admin/msp/signcerts/cert.pem
// Gateway Peer endpoint: localhost:5051
// Since the Fabric debugger network does not use TLS, Peer TLS certificate path and Peer SSL Hostname override are not required.
// In addition, ensure that you create the GRPC connection without TLS. In node.js, something like the below will work:

import { promises as fs } from 'node:fs'
import * as crypto from 'node:crypto'
import { join, resolve } from 'pathe'
import * as grpc from '@grpc/grpc-js'

import type {
  Identity,
  Signer,
} from '@hyperledger/fabric-gateway'
import {
  connect,
  signers,
} from '@hyperledger/fabric-gateway'
import type { pb } from '@saacs/saacs-pb'
import type { PlainMessage } from '@bufbuild/protobuf'
import { createBiochainGateway, createUtilGateway } from '../src/fabric/client'

interface UserCrypto { signer: Signer, identity: Identity }

export async function BuildFromBaseDir(path: string) {
  const mspId = 'Org1MSP'
  const peerEndpoint = 'localhost:5051'
  const OrgUsers = await fs.readdir(path)

  async function newGRPCClient() {
    return new grpc.Client(peerEndpoint, grpc.credentials.createInsecure())
  }

  async function getUserCryptoFiles(user: string) {
    const userDir = join(path, user, 'msp')
    const keyDir = join(userDir, 'keystore')
    const keyName = (await fs.readdir(keyDir)).pop() || ''

    const privateKey = await fs.readFile(join(userDir, 'keystore', keyName))
    const credentials = await fs.readFile(join(userDir, 'signcerts', 'cert.pem'))

    return {
      key: privateKey,
      credentials,
    }
  }

  async function getUserCrypto(user: string): Promise<UserCrypto> {
    const { key, credentials } = await getUserCryptoFiles(user)
    const privateKey = crypto.createPrivateKey(key)
    const signer = signers.newPrivateKeySigner(privateKey)

    return {
      signer,
      identity: { mspId, credentials },
    }
  }

  const userCrypto = await Promise.all(
    OrgUsers.map(async (user) => {
      return { username: user, crypto: await getUserCrypto(user) }
    }),
  )

  const Users = userCrypto.reduce(
    (acc: Record<string, UserCrypto>, { username, crypto }) => {
      acc[username] = crypto
      return acc
    },
    {},
  )
  const client = await newGRPCClient()
  console.log('Users:', Users)

  return { Users, client }
}

const _debug_baseDir = join(
  process.env.USERPROFILE || '~',
  '.vscode',
  'extensions',
  'spydra.hyperledger-fabric-debugger-1.0.0',
  'fabric',
  'local',
)
const OrgUsersDir = join(
  _debug_baseDir,
  'organizations',
  'peerOrganizations',
  'org1.debugger.com',
  'users',
)

const userDir = resolve('.', 'infra', 'network', 'organizations', 'peerOrganizations', 'org1.example.com', 'users')

const { Users, client } = await BuildFromBaseDir(userDir)

const gateway = connect({
  client,
  identity: Users['Admin@org1.example.com'].identity,
  signer: Users['Admin@org1.example.com'].signer,
})

const network = await gateway.getNetwork('default')
const contract = network.getContract('saacs-caas')

const utils = createUtilGateway(contract)
// const BiochainClient = createBiochainGateway(contract)

utils.getCurrentUser({})

try {
  const u = await utils.getCurrentUser({})
  console.log(u)

  // const r = await utils.bootstrap({ collection: { collectionId: 'Testing', name: 'Testing', authType: pb.AuthType.ROLE, itemTypes: [
  //   'saacs.biochain.v0.Specimen',
  // ] } })

  // console.log(r)

  const r2 = await utils.getCollectionsList({})
  console.log(r2)
}

catch (error) {
  console.error(error)
}
