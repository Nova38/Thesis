import { BuildGateway, GetGateway, GetService, GlobalRegistry } from "./builder.js";
import { Any, Message, createRegistry } from "@bufbuild/protobuf";
import { GenericServiceClient } from "../gen/chaincode/common/generic_pb_gateway.js";
import {
    BootstrapRequest,
    CreateCollectionRequest,
    CreateRequest,
    GetRequest,
    ListByCollectionRequest,
    ListRequest,
} from "../gen/chaincode/common/generic_pb.js";
import { Collection, ItemKey, Role } from "../gen/auth/v1/objects_pb.js";
import { Action, AuthType } from "../gen/auth/v1/auth_pb.js";
import { Specimen } from "../gen/biochain/v1/index.js";
import { construct , omit } from 'radash'
// read object from file at ../sample/biochain/orn.json



// import orn from "../sample/biochain/orn.json" assert { type: "json" };

// import { GlobalRegistry } from "";

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
            collectionId: "KU-1",
            // note: "This is a test collection",
            authType: AuthType.ROLE,
            itemTypes: ["auth.Collection", "auth.Role", "ccbio.schema.v0.Specimen"],
            default: {
                defaultPolicy: {
                    actions: [Action.VIEW],
                    allowSubPaths: false,
                    path: "",
                },
                itemPolicies: {
                    "auth.Collection": {
                        actions: [Action.VIEW],
                        allowSubPaths: false,
                        path: "",
                    },
                    "ccbio.Specimen": {
                        actions: [
                            Action.VIEW,
                            Action.CREATE,
                            Action.VIEW_HISTORY,
                        ],
                        allowSubPaths: false,
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
                collectionId: "KU",
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
        collectionId: "KU-1",
        specimenId: "TestSpecimen",
        })
    const v = Any.pack(specimen);
    const req = new CreateRequest({
        item: {
            value: v
        }
    });
    console.log(req.toJsonString({typeRegistry: service.registry}));

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
        collectionId: "KU",
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
            value: v
        }
    });
    console.log(req.toJsonString({typeRegistry: service.registry}));



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

    const response = await service.listByCollection(new ListByCollectionRequest({
        key: new ItemKey({
            collectionId: "KU",
            itemType: "auth.Role",
            itemKeyParts: ["roleId"],
        }),
    }));

    for (const item of response.items) {
        let role = new Role();
        item.value?.unpackTo(role);
        console.log(role);
    }

    console.log(response);


}

function convertSpecimen(o: any) {
    o["secondary.weight.units"] = o["secondary.weightUnits"];

    const i = omit(o, ["index", "secondary.weight.units"]);
    console.log(i);
    const c =construct(i)
    console.log(c);

    return c;
}

async function importSpecimen() {

    // get first 5 specimens from orn
const specimens = orn.slice(0, 5).map(convertSpecimen);

    console.log(specimens);

    // // create request
    // const req = new CreateRequest({
    //     items: anySpecimens.map((a) => {
    //         return { value: a };
    //     }),
    // });

    // // add to collection
    // const { service, connection, contract } = await GetService({
    //     userIdex: 0,
    //     channel: "mychannel",
    //     contractName: "roles",
    // });

    // const response = await service.create(req);
    // console.log(response);
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

async function main() {
    // await MakeCollection();
    // await ListCollections();
    // await AddRoles();
    // await AddSpecimen();
    // await ListRoles();

    await importSpecimen();
}

main();
