import {
    BuildGateway,
    GetGateway,
    GetService,
    GlobalRegistry,
} from "./builder.js";
import { Any, Message, createRegistry } from "@bufbuild/protobuf";
import { GenericServiceClient } from "../gen/chaincode/common/generic_pb_gateway.js";
import {
    BootstrapRequest,
    CreateCollectionRequest,
    CreateRequest,
    GetRequest,
    ListByAttrsRequest,
    ListByCollectionRequest,
    ListRequest,
} from "../gen/chaincode/common/generic_pb.js";
import {
    Collection,
    ItemKey,
    Role,
    UserCollectionRoles,
} from "../gen/auth/v1/objects_pb.js";
import { Action, AuthType } from "../gen/auth/v1/auth_pb.js";
import { Specimen } from "../gen/biochain/v1/index.js";
import { construct, omit } from "radash";
// read object from file at ../sample/biochain/orn.json

const collectionId = "KU-Testing";

// import orn from "../sample/biochain/orn.json" assert { type: "json" };

// import { GlobalRegistry } from "";

async function getCurrentUser() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const r1 = await service.getCurrentUser();
    console.log(r1);

    return r1;
}

async function MakeCollection() {
    // const network = await gw.gateway.getNetwork("mychannel");
    // const contract = await network.getContract("roles");
    // const service = new GenericServiceClient(contract, createRegistry());

    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const r1 = await service.getCurrentUser();
    console.log(r1);

    const collectionRequest = new CreateCollectionRequest({
        collection: new Collection({
            collectionId: collectionId,
            // note: "This is a test collection",
            authType: AuthType.ROLE,
            itemTypes: ["auth.Collection", "auth.Role", "biochain.v1.Specimen"],
            default: {
                defaultPolicy: {
                    actions: [
                        Action.VIEW,
                        Action.SUGGEST_VIEW,
                        Action.VIEW_HISTORY,
                    ],
                    allowSubPaths: false,
                    path: "",
                },
            },
        }),
    });

    const response1 = await contract.submitTransaction(
        "CreateCollection",
        collectionRequest.toJsonString(),
    );
    console.log(response1);

    // const bootstrap = await service.bootstrap(bootstrapRequest);
    // console.log(bootstrap);

    const response = await service.get(
        new GetRequest({
            key: new ItemKey({
                collectionId: collectionId,
                itemType: "auth.Collection",
                itemKeyParts: ["collection_id"],
            }),
        }),
    );
    console.log(response);
    let c = new Collection();
    response.item?.value?.unpackTo(c);
    console.log(c);

    // const response2 = await service.get(
    //     new GetRequest({
    //         key: new ItemKey({
    //             collectionId: "TestCollection",
    //             itemType: "auth.Role",
    //             itemKeyParts: ["manager"],
    //         }),
    //     }),
    // );
    // console.log(response2);
    // let r = new Role();
    // response2.item?.value?.unpackTo(r);
    // console.log(r);

    // const role = new Role({
    //     collectionId: "TestCollection",
    //     roleId: "TestRole",
    //     note: "This is a test role",
    //     parentRoleIds: [],
    //     polices: {
    //         defaultPolicy: {
    //             actions: [Action.VIEW],
    //             allowSubPaths: false,
    //             path: "",
    //         },
    //     },
    // });

    // const r3 = await service.create({});

    console.log("hello world");
}

async function AddSpecimen() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const specimen = new Specimen({
        collectionId: collectionId,
        specimenId: "310d45f4-98ae-407a-869d-83200dc5376e",
        georeference: {
            georeferenceDate: {},
        },
        primary: {
            catalogDate: {},
            determinedDate: {},
            fieldDate: {},
            originalDate: {},
        },
        secondary: {},
        taxon: {},
        grants: {},
        images: {},
        loans: {},
    });
    const v = Any.pack(specimen);
    const req = new CreateRequest({
        item: {
            value: v,
        },
    });
    console.log(req.toJsonString({ typeRegistry: service.registry }));

    const response = await service.create(req);
    console.log(response);
}

async function AddRoles() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const role = new Role({
        collectionId: collectionId,
        roleId: "TestRole",
        note: "This is a test role",
        parentRoleIds: [],
        polices: {
            defaultPolicy: {
                actions: [Action.VIEW],
                allowSubPaths: false,
                path: "",
            },
        },
    });

    const v = Any.pack(role);

    const req = new CreateRequest({
        item: {
            value: v,
        },
    });
    console.log(req.toJsonString({ typeRegistry: service.registry }));

    // console.log(req.toJsonString(service.jsonWriteOptions));

    const response = await service.create(req);

    console.log(response);
}

async function ListRoles() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const response = await service.listByCollection(
        new ListByCollectionRequest({
            key: new ItemKey({
                collectionId: collectionId,
                itemType: "auth.Role",
                itemKeyParts: [collectionId],
            }),
        }),
    );

    for (const item of response.items) {
        let role = new Role();
        item.value?.unpackTo(role);
        console.log(role);
    }

    console.log(response);
}

async function ListUserRoles() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const response = await service.listByCollection(
        new ListByCollectionRequest({
            key: new ItemKey({
                collectionId: collectionId,
                itemType: "auth.UserCollectionRoles",
                itemKeyParts: [""],
            }),
        }),
    );

    for (const item of response.items) {
        let role = new UserCollectionRoles();
        item.value?.unpackTo(role);
        console.log(role);
    }

    console.log(response);
}

async function ListSpecimens() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const response = await service.listByAttrs(
        new ListByAttrsRequest({
            key: new ItemKey({
                collectionId: collectionId,
                itemType: Specimen.typeName,
                itemKeyParts: [collectionId],
            }),
            numAttrs: 0,
        }),
    );

    for (const item of response.items) {
        let specimen = new Specimen();
        item.value?.unpackTo(specimen);
        console.log(specimen);
    }

    console.log(response);
}

async function GetUserRolesList(numAttrs: number) {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const result = await service.listByAttrs(
        new ListByAttrsRequest({
            key: new ItemKey({
                collectionId: collectionId,
                itemType: UserCollectionRoles.typeName,
                itemKeyParts: [
                    "Org1MSP",
                    "eDUwOTo6Q049b3JnMWFkbWluLE9VPWFkbWluLE89SHlwZXJsZWRnZXIsU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUzo6Q049Y2Eub3JnMS5leGFtcGxlLmNvbSxPPW9yZzEuZXhhbXBsZS5jb20sTD1EdXJoYW0sU1Q9Tm9ydGggQ2Fyb2xpbmEsQz1VUw==",
                ],
            }),
            numAttrs: numAttrs,
        }),
    );
    const roles: any[] = [];
    for (const item of result.items) {
        const role = new UserCollectionRoles();
        item.value?.unpackTo(role);
        roles.push(role);
    }
    // response.item?.value?.unpackTo(roles);
    console.log(roles);
}
async function GetUserRoles() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const { user } = await getCurrentUser();
    if (!user) {
        throw new Error("no user");
    }

    const response = await service.get(
        new GetRequest({
            key: {
                collectionId: "33KU-1",
                itemType: "auth.UserCollectionRoles",
                itemKeyParts: [user.mspId, user.userId],
            },
        }),
    );
    let roles = new UserCollectionRoles();
    response.item?.value?.unpackTo(roles);
    console.log(roles);
}

// function convertSpecimen(o: any) {
//     o["secondary.weight.units"] = o["secondary.weightUnits"];

//     const i = omit(o, ["index", "secondary.weight.units"]);
//     console.log(i);
//     const c =construct(i)
//     console.log(c);

//     return c;
// }

// async function importSpecimen() {

//     // get first 5 specimens from orn
// const specimens = orn.slice(0, 5).map(convertSpecimen);

//     console.log(specimens);

//     // // create request
//     // const req = new CreateRequest({
//     //     items: anySpecimens.map((a) => {
//     //         return { value: a };
//     //     }),
//     // });

//     // // add to collection
//     // const { service, connection, contract } = await GetService({
//     //     userIdex: 0,
//     //     channel: "mychannel",
//     //     contractName: "roles",
//     // });

//     // const response = await service.create(req);
//     // console.log(response);
// }

async function ListCollections() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const response = await service.getCollectionsList();
    console.log(response);
    // let c = new Collection();
    // console.log(c);
}

async function main() {
    await getCurrentUser();
    await MakeCollection();
    // await ListCollections();
    await AddRoles();
    await AddSpecimen();
    // await ListRoles();
    // await importSpecimen();
    // await ListSpecimens();
    // await ListUserRoles();
    // await GetUserRolesList(0);
    // await GetUserRoles();
}

main();
