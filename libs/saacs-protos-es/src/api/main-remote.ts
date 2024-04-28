import { GetGateway, GetService, GlobalRegistry } from './builder-remote.js'
import { Any, Message, createRegistry } from '@bufbuild/protobuf'
import { GenericServiceClient } from '../gen/chaincode/common/generic_pb_gateway.js'
import {
  BootstrapRequest,
  CreateCollectionRequest,
  CreateRequest,
  GetRequest,
  ListByAttrsRequest,
  ListByCollectionRequest,
  ListRequest,
} from '../gen/chaincode/common/generic_pb.js'
import {
  Collection,
  ItemKey,
  Role,
  UserCollectionRoles,
} from '../gen/auth/v1/objects_pb.js'
import { Action, AuthType } from '../gen/auth/v1/auth_pb.js'
import { Specimen } from '../gen/biochain/v1/index.js'
import { construct, omit, parallel, random, tryit } from 'radash'
import { auth, ccbio } from '../index.js'
// read object from file at ../sample/biochain/orn.json
import z from 'zod'
import { Timestamp, FieldMask } from '@bufbuild/protobuf'
import remote from './remote.json' assert { type: 'json' }

const collectionId = 'KU Ornithology Great Plains'
const channel = 'demo'
const contractName = 'a'
// const collectionId = "KU-Zoology";
const utf8Decoder = new TextDecoder()

// import orn from "../../../biochain/import/ku_orn_cov.json" assert { type: "json" };

// import { GlobalRegistry } from "";

async function getCurrentUser(username: string) {
  const { service, connection, contract } = await GetService({
    username: username,
    channel: channel,
    contractName: contractName,
  })

  const response = await service.getCurrentUser()
  // console.log({ username, response, connection });

  return {
    response,
    cred: connection.identity.credentials.toString(),
    key: connection.private_key.toString(),
  }
}
async function MakeCollection() {
  // const network = await gw.gateway.getNetwork("mychannel");
  // const contract = await network.getContract("roles");
  // const service = new GenericServiceClient(contract, createRegistry());

  console.group('MakeCollection')

  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const r1 = await service.getCurrentUser()
  console.log(r1)

  const collectionRequest = new CreateCollectionRequest({
    collection: new Collection({
      collectionId: collectionId,
      // note: "This is a test collection",
      authType: AuthType.ROLE,
      itemTypes: [
        ccbio.Specimen.typeName,
        Collection.typeName,
        UserCollectionRoles.typeName,
        Role.typeName,
      ],
      default: {
        defaultPolicy: {
          actions: [Action.VIEW, Action.SUGGEST_VIEW, Action.VIEW_HISTORY],
          allowSubPaths: false,
          path: '',
        },
        itemPolicies: {
          [ccbio.Specimen.typeName]: {
            actions: [Action.VIEW, Action.SUGGEST_VIEW, Action.VIEW_HISTORY],
            allowSubPaths: false,
            fullPath: '',
            path: '',
          },
          [UserCollectionRoles.typeName]: {
            actions: [Action.VIEW, Action.SUGGEST_VIEW],
            allowSubPaths: false,
            fullPath: '',
            path: '',
          },
          [Role.typeName]: {
            actions: [Action.VIEW, Action.SUGGEST_VIEW],
            allowSubPaths: false,
            fullPath: '',
            path: '',
          },
        },
      },
    }),
  })
  console.log(collectionRequest.toJsonString())

  const response1 = await contract.submitTransaction(
    'CreateCollection',
    collectionRequest.toJsonString(),
  )
  console.log(response1)

  // const bootstrap = await service.bootstrap(bootstrapRequest);
  // console.log(bootstrap);

  // const response = await service.get(
  //     new GetRequest({
  //         key: new ItemKey({
  //             collectionId: collectionId,
  //             itemType: Collection.typeName,
  //             itemKeyParts: [collectionId],
  //         }),
  //     }),
  // );
  // console.log(response);
  // let c = new Collection();
  // response.item?.value?.unpackTo(c);
  // console.log(c);

  console.groupEnd()
}

async function GetUserId(username: string) {
  const { service, connection, contract } = await GetService({
    username: username,
    channel: channel,
    contractName: contractName,
  })

  const r1 = await service.getCurrentUser()
  return r1.user?.userId
}

async function addUserRoles({
  username,
  roleId,
}: {
  username: string
  roleId: string
}) {
  const admin = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const userId = await GetUserId(username)
  // console.log({ username, userId });
  const userRole = new UserCollectionRoles({
    collectionId: collectionId,
    userId: await GetUserId(username),
    mspId: 'Org1MSP',
    roleIds: [roleId],
  })
  const response = await admin.service.create(
    new CreateRequest({
      item: {
        value: Any.pack(userRole),
      },
    }),
  )
  console.log(response.toJson({ typeRegistry: admin.service.registry }))
  return response
}

async function AddRoles() {
  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const suggestRole = new Role({
    collectionId: collectionId,
    roleId: 'Suggester',
    note: 'This is a test role',
    parentRoleIds: [],
    polices: {
      defaultPolicy: {},
      itemPolicies: {
        [ccbio.Specimen.typeName]: {
          actions: [Action.SUGGEST_VIEW, Action.SUGGEST_CREATE, Action.VIEW],
          allowSubPaths: false,
        },
      },
    },
  })

  const MemberRole = new Role({
    collectionId: collectionId,
    roleId: 'Member',
    note: 'This is a test role',
    parentRoleIds: [],
    polices: {
      defaultPolicy: {},
      itemPolicies: {
        [ccbio.Specimen.typeName]: {
          actions: [
            Action.SUGGEST_VIEW,
            Action.SUGGEST_CREATE,
            Action.VIEW,
            Action.CREATE,
            Action.UPDATE,
            Action.SUGGEST_CREATE,
            Action.SUGGEST_APPROVE,
            Action.VIEW_HISTORY,
          ],
          allowSubPaths: false,
        },
      },
    },
  })
  console.group('AddRoles')
  let req = new CreateRequest()
  console.log(req.toJsonString({ typeRegistry: service.registry }))

  let response = await service.create(
    new CreateRequest({
      item: {
        value: Any.pack(MemberRole),
      },
    }),
  )
  let role = new Role()
  response.item?.value?.unpackTo(role)
  console.log(role)

  console.groupEnd()

  req = new CreateRequest()
  console.log(req.toJsonString({ typeRegistry: service.registry }))

  response = await service.create(
    new CreateRequest({
      item: {
        value: Any.pack(suggestRole),
      },
    }),
  )
  role = new Role()
  response.item?.value?.unpackTo(role)
  console.log(role)
}

async function AddSpecimen() {
  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const specimen = new Specimen({
    collectionId: collectionId,
    specimenId: random(0, 100000).toString(),
    georeference: {
      georeferenceDate: {},
    },
    primary: {
      catalogDate: {},
      determinedDate: {},
      fieldDate: {},
      originalDate: {},
    },
    secondary: {
      preparations: {},
    },
    taxon: {},
    grants: {},
    images: {},
    loans: {},
  })
  const v = Any.pack(specimen)
  const req = new CreateRequest({
    item: {
      value: v,
    },
  })
  console.log(req.toJsonString({ typeRegistry: service.registry }))

  const response = await service.create(req)
  console.log(response)
}

async function ListRoles() {
  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const response = await service.listByAttrs(
    new ListByAttrsRequest({
      key: new ItemKey({
        collectionId: collectionId,
        itemType: 'auth.Role',
        itemKeyParts: [collectionId],
      }),
      numAttrs: 0,
    }),
  )

  for (const item of response.items) {
    let role = new Role()
    item.value?.unpackTo(role)
    console.log(role)
  }

  console.log(response)
}

async function ListUserRoles() {
  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const response = await service.listByAttrs(
    new ListByAttrsRequest({
      key: new ItemKey({
        collectionId: collectionId,
        itemType: UserCollectionRoles.typeName,
        itemKeyParts: [collectionId],
      }),
      numAttrs: 0,
    }),
  )

  for (const item of response.items) {
    let role = new UserCollectionRoles()
    item.value?.unpackTo(role)
    console.log(role)
  }

  console.log(response)
}

async function ListSpecimens() {
  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const response = await service.listByAttrs(
    new ListByAttrsRequest({
      key: new ItemKey({
        collectionId: collectionId,
        itemType: Specimen.typeName,
        itemKeyParts: [collectionId],
      }),
      numAttrs: 0,
      limit: 1,
    }),
  )

  for (const item of response.items) {
    let specimen = new Specimen()
    item.value?.unpackTo(specimen)
    console.log(specimen)
  }

  console.log(response)
}

async function GetUserRoles(username: string = 't-tom') {
  const { service, connection, contract } = await GetService({
    username: username,
    channel: channel,
    contractName: contractName,
  })

  const { response: r1 } = await getCurrentUser(username)
  const user = r1.user
  if (!user) {
    throw new Error('no user')
  }

  const response = await service.get(
    new GetRequest({
      key: {
        collectionId: collectionId,
        itemType: 'auth.UserCollectionRoles',
        itemKeyParts: [user.mspId, user.userId],
      },
    }),
  )
  let roles = new UserCollectionRoles()
  response.item?.value?.unpackTo(roles)
  console.log({ username, roles })
}

async function ListCollections() {
  const { service, connection, contract } = await GetService({
    username: 't-tom',
    channel: channel,
    contractName: contractName,
  })

  const response = await service.getCollectionsList()
  console.log(response)
  // let c = new Collection();
  // console.log(c);
}

// async function AddUserRoles() {
//     // await addUserRoles({ userId: 1, roleId: "Suggester" });
//     await addUserRoles({ userId: 2, roleId: "Manager" });
//     await addUserRoles({ userId: 3, roleId: "Suggester" });
//     await addUserRoles({ userId: 4, roleId: "Manager" });
// }

async function listUsers() {
  const users = remote.organizations.Org1MSP.users
  for (const [k, value] of Object.entries(users)) {
    console.log('username:', k)

    try {
      await getCurrentUser(k)
    } catch (e) {
      console.error(e)
    }
    await getCurrentUser(k)

    // await addUserRoles({ username: k, roleId: "Member" });
  }
}

async function Bootstrap() {
  // await MakeCollection();
  // await AddRoles();
  // await ListRoles();
  // await AddUserRoles();

  // await ListCollections();
  await AddSpecimen()
  // await ListSpecimens();
  await GetUserRoles()
}

async function fuckthis() {
  const r1 = await getCurrentUser('t-tom')
  const r2 = await getCurrentUser('t-town')
  const r3 = await getCurrentUser('town')
  const r4 = await getCurrentUser('f0')
  // const r4 = await getCurrentUser("t-narayani.ku");
  const r5 = await getCurrentUser('t-vicky.argudo19')
  const r6 = await getCurrentUser('vicky.argudo19')
  ;``
  const certs = [r1.cred, r2.cred, r3.cred, r4.cred, r5.cred]
  const ids = [
    r1.response.user?.userId,
    r2.response.user?.userId,
    r3.response.user?.userId,
    r4.response.user?.userId,
    r5.response.user?.userId,
    r6.response.user?.userId,
  ]
  const keys = [r1.key, r2.key, r3.key, r4.key, r5.key]

  // console.log({ certs, ids, keys });

  const UIds = [...new Set(ids)]
  // const UCerts = [...new Set(certs)];
  // const UKeys = [...new Set(certs)];
  console.log({ UIds })
}
async function fuckthisfix() {
  const users = remote.organizations.Org1MSP.users
  const ids = []
  const errs = []
  for (const [k, value] of Object.entries(users)) {
    try {
      const r = await getCurrentUser(k)
      ids.push(r.response.user?.userId)
    } catch (error) {
      errs.push({ k, error })
    }
  }
  console.log(errs)
  const UIds = [...new Set(ids)]

  console.log({ ids, UIds, ids_len: ids.length, uids_len: UIds.length })
}
async function setupUsers() {
  const rawUsers = remote.organizations.Org1MSP.users
  const errs = []
  const p = []

  // for (const [k, value] of Object.entries(users)) {
  //     console.log("setting up user:", k);
  //     try {
  //         if (
  //             k === "t-f.machado.stredel" ||
  //             k === "t-town" ||
  //             k == "t-lhdecicco"
  //         ) {
  //             p.push(
  //                 addUserRoles({
  //                     username: k,
  //                     roleId: "Manager",
  //                 }),
  //             );
  //         } else {
  //             p.push(addUserRoles({ username: k, roleId: "Member" }));
  //         }
  //     } catch (error) {
  //         errs.push({ k, error });
  //     }
  // }
  const [err, users] = await tryit(parallel)(
    27,
    Object.keys(rawUsers),
    async (userId) => {
      console.log('setting up user:', userId)
      if (
        userId === 't-f.machado.stredel' ||
        userId === 't-town' ||
        userId === 't-lhdecicco' ||
        userId === 't-tom'
      ) {
        return addUserRoles({ username: userId, roleId: 'Manager' })
      } else {
        return addUserRoles({ username: userId, roleId: 'Member' })
      }
    },
  )

  console.log({ err, users })
}
async function main() {
  // await fuckthisfix();
  // await GetUserRoles();
  await setupUsers()
  // const r = await addUserRoles({
  //     username: "t-town",
  //     roleId: "Manager",
  // });
  // console.log(r);

  // collectionId = "KU-Zoology";
  // await Bootstrap();
  // await MakeCollection();
  // await ListCollections();
  // await AddSpecimen();
  // await ListCollections();
  // await ListSpecimens();
  // await AddRoles();
  // await AddSpecimen();
  // await ListRoles();
  // await listUsers();
  // await setupUsers();
  // await listUsers();
  // await ListUserRoles();
  // await setupUsers();
  // await importSpecimen();
  // await ListSpecimens();
  // await ListUserRoles();
  // await GetUserRolesList("me");
  // await GetUserRoles("t-tom");
  // await GetUserRoles("t-town");
  // await GetUserRoles("t-vicky.argudo19");
  // await addUserRoles({ username: "t-vicky.argudo19", roleId: "Member" });
  // await GetUserRoles("t-ramosdovalj");
}

await main()
