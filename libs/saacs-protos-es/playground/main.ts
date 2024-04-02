// import * as hlf from "../dist/index.mjs";

// import orn from "../sample/biochain/orn.json" assert { type: "json" };

// async function main() {
//     for
// }

// main().catch(console.error);

import type {
    MethodInfoBiDiStreaming,
    MethodInfoClientStreaming,
    MethodInfoServerStreaming,
    MethodInfoUnary,
    PlainMessage,
    ServiceType,
    Message,
    PartialMessage,
} from "@bufbuild/protobuf";
import type { Transport, CallOptions } from "@connectrpc/connect";
import { createPromiseClient } from "@connectrpc/connect";

// import { auth, common } from "../dist/index"

// Tweak this to change the request signature like changing from PlainMessage to PartialMessage.
// type Request<R, I extends Message<I>> =
//     | Strict<R, PlainMessage<I>>
//     | StrictMessage<R, I>;

// prettier-ignore
export type GatewayClient<T extends ServiceType> = {
    [P in keyof T["methods"]]:
    T["methods"][P] extends MethodInfoUnary<infer I, infer O> ? (request: PartialMessage<I>, options?: CallOptions) => Promise<O>
    : never;
};


export type BiochainGateway = GatewayClient<typeof common.generic.GenericServiceClient>;
