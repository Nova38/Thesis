// Path to crypto materials: On Windows machine - %USERPROFILE%.vscode\extensions\spydra.hyperledger-fabric-debugger-{version}\fabric\local\organizations\peerOrganizations\org1.debugger.com. On Linux and Mac, it will be - ~/.vscode/extensions/spydra.hyperledger-fabric-debugger-{version}/fabric/local/organizations/peerOrganizations/org1.debugger.com.
// User key directory: users/Org1Admin/msp/keystore
// User certificate path: users/Org1Admin/msp/signcerts/cert.pem
// Gateway Peer endpoint: localhost:5051
// Since the Fabric debugger network does not use TLS, Peer TLS certificate path and Peer SSL Hostname override are not required.
// In addition, ensure that you create the GRPC connection without TLS. In node.js, something like the below will work:

import { resolve, join } from 'pathe'
import { promises as fs } from 'fs'
import * as crypto from 'crypto'
import {
  connect,
  Contract,
  Identity,
  Signer,
  signers,
} from '@hyperledger/fabric-gateway'

const baseDir = join(
  process.env.USERPROFILE || '~',
  '.vscode',
  'extensions',
  'spydra.hyperledger-fabric-debugger-1.0.0',
  'fabric',
  'local',
)

const OrgUsersDir = join(
  baseDir,
  'organizations',
  'peerOrganizations',
  'org1.debugger.com',
  'users',
)

type UserCrypto = { signer: Signer; identity: Identity }
const mspId = 'Org1MSP'
const OrgUsers = await fs.readdir(OrgUsersDir)

async function getUserCryptoFiles(user: string) {
  const userDir = join(OrgUsersDir, user, 'msp')
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

console.log(Users)
