import { BuildGateway, GetGateway, GetService } from "./builder.js";
import { Any, Message, createRegistry } from "@bufbuild/protobuf";
import { GenericServiceClient } from "../gen/chaincode/common/generic_pb_gateway.js";
import {
    BootstrapRequest,
    GetRequest,
    ListByCollectionRequest,
} from "../gen/chaincode/common/generic_pb.js";
import { Collection, ItemKey, Role } from "../gen/auth/v1/objects_pb.js";
import { Action, AuthType } from "../gen/auth/v1/auth_pb.js";
import { GlobalRegistry } from "../utils/registry.js";

async function main() {
    // const network = await gw.gateway.getNetwork("mychannel");
    // const contract = await network.getContract("roles");
    // const service = new GenericServiceClient(contract, createRegistry());

    const { service } = await GetService({
        userIdex: 0,
        channel: "mychannel",
        contractName: "roles",
    });

    const r1 = await service.getCurrentUser();
    console.log(r1);
    // let bootstrapRequest = new BootstrapRequest({
    //     collections: [
    //         new Collection({
    //             name: "TestCollection",
    //             collectionId: "TestCollection",
    //             itemTypes: ["sample"],
    //             default: {},
    //             authType: AuthType.ROLE,
    //         }),
    //     ],
    // });

    // const bootstrap = await service.bootstrap(bootstrapRequest);
    // console.log(bootstrap);

    const response = await service.get(
        new GetRequest({
            key: new ItemKey({
                collectionId: "TestCollection",
                itemType: "auth.Collection",
                itemKeyParts: ["collection_id"],
            }),
        }),
    );
    console.log(response);
    let c = new Collection();
    response.item?.value?.unpackTo(c);
    console.log(c);

    const response2 = await service.get(
        new GetRequest({
            key: new ItemKey({
                collectionId: "TestCollection",
                itemType: "auth.Role",
                itemKeyParts: ["manager"],
            }),
        }),
    );
    console.log(response2);
    let r = new Role();
    response2.item?.value?.unpackTo(r);
    console.log(r);

    const role = new Role({
        collectionId: "TestCollection",
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

    const r3 = await service.create({});

    console.log("hello world");
}

main();
