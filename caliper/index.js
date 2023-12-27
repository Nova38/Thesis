"use strict";
const hlf = require("saacs-es");
const logger = require("@hyperledger/caliper-core").CaliperUtils.getLogger(
    "dos"
);
const { Any, createRegistry, FieldMask } = require("@bufbuild/protobuf");

const { WorkloadModuleBase } = require("@hyperledger/caliper-core");
const { randomInt } = require("crypto");
const users = ["User1", "User2", "User3", "User4", "User5", "Admin"];

function randomUpdateArgs() {
    let args = new hlf.pb.common.generic.UpdateRequest({
        updateMask: new FieldMask({
            paths: ["name"],
        }),
        item: new hlf.pb.auth.objects.Item({
            key: new hlf.pb.auth.objects.ItemKey({
                collectionId: "collection0",
                itemType: "auth.Collection",
                itemKeyParts: ["collection0"],
            }),
            value: Any.pack(
                new hlf.pb.auth.objects.Collection({
                    collectionId: "collection0",
                    name: `write-${randomInt(0, 100)}`,
                })
            ),
        }),
    });

    console.log(args);

    return args.toJsonString({
        typeRegistry: createRegistry(
            ...hlf.pb.sample.allMessages,
            ...hlf.pb.auth.objects.allMessages
        ),
    });
}

// console.log(randomUpdateArgs());

function randomNewBookArgs() {
    let args = new hlf.pb.common.generic.CreateRequest({
        item: new hlf.pb.auth.objects.Item({
            value: Any.pack(
                new hlf.pb.sample.Book({
                    collectionId: "collection0",
                })
            ),
        }),
    });
    return args.toJsonString({
        typeRegistry: createRegistry(
            ...hlf.pb.sample.allMessages,
            ...hlf.pb.auth.objects.allMessages
        ),
    });
}

// console.log(randomNewBookArgs());


function makeUpdates(){
    let items = hlf.utils.factory.BuildWorkloadSuggestionsOne({
        collectionId: "collection0",
        numItems: 100,
        numSuggestions: 100,
        typeName: "auth.Collection",
        workerIndex: 2,
    });

    return items
}

console.log(makeUpdates());
