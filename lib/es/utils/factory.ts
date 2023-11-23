import { Object$ } from "../gen/auth/v1/auth_pb";
import { Author } from "../gen/chaincode/sample/v0/items_pb";

import { Any, createRegistry, createRegistryFromDescriptors } from "@bufbuild/protobuf";


export const registry = createRegistry(
    Author
)

export function createAuthor(): Author {
    const author = new Author();
    author.authorName = "John Doe";
    // author.setName("John Doe");
    // author.setAge(42);
    return author;
}

export function createAuthorObject(): Object$ {
    const obj = new Object$();
    const author = createAuthor();

    obj.value = Any.pack(author);
    // author.setAuthorName("John Doe");
    // author.setName("John Doe");
    // author.setAge(42);
    return obj ;
}
