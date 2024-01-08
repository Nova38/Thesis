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
import { construct, omit, random } from "radash";
import { auth, ccbio } from "../index.js";
// read object from file at ../sample/biochain/orn.json
import z from "zod";
import { Timestamp, FieldMask } from "@bufbuild/protobuf";

const collectionId = "ku_orn";
// const collectionId = "KU-Zoology";
const utf8Decoder = new TextDecoder();

// import orn from "../../../biochain/import/ku_orn_cov.json" assert { type: "json" };

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

    console.group("MakeCollection");

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
            itemTypes: [
                ccbio.Specimen.typeName,
                Collection.typeName,
                UserCollectionRoles.typeName,
                Role.typeName,
            ],
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
                itemPolicies: {
                    [ccbio.Specimen.typeName]: {
                        actions: [
                            Action.VIEW,
                            Action.SUGGEST_VIEW,
                            Action.VIEW_HISTORY,
                        ],
                        allowSubPaths: false,
                        fullPath: "",
                        path: "",
                    },
                    [UserCollectionRoles.typeName]: {
                        actions: [Action.VIEW, Action.SUGGEST_VIEW],
                        allowSubPaths: false,
                        fullPath: "",
                        path: "",
                    },
                    [Role.typeName]: {
                        actions: [Action.VIEW, Action.SUGGEST_VIEW],
                        allowSubPaths: false,
                        fullPath: "",
                        path: "",
                    },
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
                itemType: Collection.typeName,
                itemKeyParts: [collectionId],
            }),
        }),
    );
    console.log(response);
    let c = new Collection();
    response.item?.value?.unpackTo(c);
    console.log(c);

    console.groupEnd();
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
}

async function GetUserId({ userIdex }: { userIdex: number }) {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const r1 = await service.getCurrentUser();
    return r1.user?.userId;
}

async function addUserRoles({
    userId,
    roleId,
}: {
    userId: number;
    roleId: string;
}) {
    const admin = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const userRole = new UserCollectionRoles({
        collectionId: collectionId,
        userId: await GetUserId({ userIdex: userId }),
        mspId: "Org1MSP",
        roleIds: [roleId],
    });
    const response = await admin.service.create(
        new CreateRequest({
            item: {
                value: Any.pack(userRole),
            },
        }),
    );
    console.log(response.toJson({ typeRegistry: admin.service.registry }));
}

async function AddRoles() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const suggestRole = new Role({
        collectionId: collectionId,
        roleId: "Suggester",
        note: "This is a test role",
        parentRoleIds: [],
        polices: {
            defaultPolicy: {},
            itemPolicies: {
                [ccbio.Specimen.typeName]: {
                    actions: [
                        Action.SUGGEST_VIEW,
                        Action.SUGGEST_CREATE,
                        Action.VIEW,
                    ],
                    allowSubPaths: false,
                },
            },
        },
    });

    console.group("AddRoles");

    const send = async (r: Role) => {
        const req = new CreateRequest();
        console.log(req.toJsonString({ typeRegistry: service.registry }));

        const response = await service.create(
            new CreateRequest({
                item: {
                    value: Any.pack(r),
                },
            }),
        );
        const role = new Role();
        response.item?.value?.unpackTo(role);
        console.log(role);

        return role;
    };
    send(suggestRole);
    console.groupEnd();
}

async function AddSpecimen() {
    const { service, connection, contract } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

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
                itemType: UserCollectionRoles.typeName,
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
            limit: 1,
        }),
    );

    for (const item of response.items) {
        let specimen = new Specimen();
        item.value?.unpackTo(specimen);
        console.log(specimen);
    }

    console.log(response);
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
                collectionId: collectionId,
                itemType: "auth.UserCollectionRoles",
                itemKeyParts: [user.mspId, user.userId],
            },
        }),
    );
    let roles = new UserCollectionRoles();
    response.item?.value?.unpackTo(roles);
    console.log(roles);
}

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

async function AddUserRoles() {
    // await addUserRoles({ userId: 1, roleId: "Suggester" });
    await addUserRoles({ userId: 2, roleId: "Manager" });
    await addUserRoles({ userId: 3, roleId: "Suggester" });
    await addUserRoles({ userId: 4, roleId: "Manager" });
}

async function Bootstrap() {
    await MakeCollection();
    // await AddRoles();
    await ListRoles();
    await AddUserRoles();

    // await ListCollections();
    await AddSpecimen();
    // await ListSpecimens();
    await GetUserRoles();
}

async function main() {
    // collectionId = "KU-Zoology";
    // await Bootstrap();
    await AddSpecimen();

    await ListCollections();
    await ListSpecimens();
    // await getCurrentUser();
    // await MakeCollection();
    // await AddRoles();
    // await AddSpecimen();
    // await ListRoles();
    // await importSpecimen();
    // await ListSpecimens();
    // await ListUserRoles();
    // await GetUserRolesList(0);
    // await GetUserRoles();
}

main();
