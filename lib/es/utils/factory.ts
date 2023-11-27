import { Item } from "../gen/auth/v1/auth_pb";
import { Author } from "../gen/chaincode/sample/v0/items_pb";

import { Any, createRegistry, createRegistryFromDescriptors } from "@bufbuild/protobuf";

import { Registry } from "./registry";




export function createAuthor(): Author {
    const author = new Author();
    author.authorName = "John Doe";
    // author.setName("John Doe");
    // author.setAge(42);
    return author;
}

export function createAuthorItem(): Item {
    const obj = new  Item();
    const author = createAuthor();
    console.log("author", author);
    obj.value = Any.pack(author);

    console.log("author", author);

    return obj ;
}



export function unpackItem(item: Item)  {
    const author = new Author();
    
    return item.value?.unpack(Registry) 

}
