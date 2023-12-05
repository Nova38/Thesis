// import { Author } from "../gen/chaincode/sample/v0/items_pb";

import { Any, createRegistry, createRegistryFromDescriptors } from "@bufbuild/protobuf";


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
    return Math.random() * max;
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
