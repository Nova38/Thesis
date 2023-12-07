// import { Author } from "../gen/chaincode/sample/v0/items_pb";

import { Any, AnyMessage, createRegistry, createRegistryFromDescriptors } from "@bufbuild/protobuf";
import { Item, ItemKey } from "../gen/auth/v1/objects_pb.js";
import { Book, SimpleItem } from "../gen/sample/v0/items_pb.js";
import { CreateRequest } from "../gen/chaincode/common/generic_pb.js";
import { auth, sample } from "../gen/index.js";

export function ShortToFullTypeName(type: any) {
    if (!type) {
        return SimpleItem.typeName;

    }else  if (type == "simple") {
        return SimpleItem.typeName;
    } else if (type == "book") {
        return Book.typeName;
    }

    return SimpleItem.typeName

}




export function BuildWorkloadItemKeys(arg: {numItems: number, numCollections: number, typeName: string, workerIndex: number}) {

    let ColMap: Record<string, Array<ItemKey>> = {}

    for (let i = 0; i < arg.numCollections; i++) {
        let col = `collection${i}`;
        ColMap[i] = [];

        for (let j = 0; j < arg.numItems; j++) {
            let id = `worker:${arg.workerIndex}-item:${j}`;
            let key = new ItemKey({
                itemType: ShortToFullTypeName(arg.typeName),
                itemKeyParts: [id],
                collectionId: col
            })



            ColMap[i].push(key);
        }
    }

    return ColMap;
}

export function ToCreateRequestString(obj: AnyMessage) {
    let v = Any.pack(obj)
    let item = new Item({
        value: v,
    })
    let arg = new CreateRequest({
        item: item,
    })

    return arg.toJsonString({typeRegistry: createRegistry(... sample.allMessages, ...auth.auth.allMessages, ...auth.objects.allMessages)})

}

export function BuildWorkloadItem(arg: {numItems: number, numCollections: number, typeName: string, workerIndex: number}) {

    let ColMap: Record<string, Array<AnyMessage>> = {}

    for (let i = 0; i < arg.numCollections; i++) {
        let col = `collection${i}`;
        ColMap[i] = [];

        for (let j = 0; j < arg.numItems; j++) {
            let id = `worker:${arg.workerIndex}-item:${j}`;
            let key = new ItemKey({
                itemType: ShortToFullTypeName(arg.typeName),
                itemKeyParts: [id],
                collectionId: col
            })
            if (arg.typeName == "book" ) {
                ColMap[i].push( new Book({
                    collectionId: col,
                    isbn: id,
                    author: "author",
                    language: "en",
                    bookTitle: "title",
                    description: "description",
                    year: 2023,
                    publisher: "publisher",
                }))

            }else  {
                ColMap[i].push( new SimpleItem({
                    collectionId: col,
                    id: id,
                    quantity: j,
                    name: `item:${j}`,
                }))
            }

        }
    }
    return ColMap;

}






export function randomUser(){
    const users = ['User1', 'User2', 'User3', 'User4', 'User5', 'Admin']
    return users[Math.floor(Math.random() * users.length)];
}

export function modCollectionId(numCollections: number, mod: number){
    const collections = []
    for (let i = 0; i < numCollections; i++) {
        collections.push(`collection${i}`)
    }
    return collections[mod % numCollections];
}

export function randomCollection(numCollections: number){

    const collections = []
    for (let i = 0; i < numCollections; i++) {
        collections.push(`Collection${i}`)
    }
    return collections[Math.floor(Math.random() * collections.length)];
}


export function BuildCollection(types: any[]) {


}
export function randomInt( max: number): number {
    // faker.seed(seed);
    return Math.floor(Math.random() * max);
}



// export function createAuthor(): Author {
//     const author = new Author();
//     author.authorName = "John Doe";
//     // author.setName("John Doe");
//     // author.setAge(42);
//     return author;
// }

// export function createAuthorItem(): Item {
//     const obj = new  Item();
//     const author = createAuthor();
//     console.log("author", author);
//     obj.value = Any.pack(author);

//     console.log("author", author);

//     return obj ;
// }



// export function unpackItem(item: Item)  {
//     const author = new Author();

//     return item.value?.unpack(Registry)

// }
