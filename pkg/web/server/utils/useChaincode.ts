import type { H3Event } from "h3";
import { sessionConfig } from "./session";
import { User } from "./db";
import { connect, signers } from "@hyperledger/fabric-gateway";
import * as grpc from "@grpc/grpc-js";
import * as crypto from "crypto";

import { auth, ccbio, sample, common } from "saacs-es";
import { IMessageTypeRegistry, createRegistry } from "@bufbuild/protobuf";

export const GlobalRegistry: IMessageTypeRegistry = createRegistry(
  ...auth.auth.allMessages,
  ...auth.objects.allMessages,
  ...common.generic.allMessages,
  ...common.reference.allMessages,
  ...ccbio.allMessages,
  ...sample.allMessages,
);

export interface FabricConfig {
  chaincode: {
    channel: string;
    chaincode: string;
  };
  peer: {
    url: string;
    tlsCACerts: {
      pem: string;
    };
    grpcOptions: {
      "ssl-target-name-override": string;
    };
  };
  public: {
    mspId: string;
    credentials: string;
    key: string;
  };
}

export const fabricConfig: FabricConfig = useRuntimeConfig().fabric || {};
export function userToIdentity(user: User) {
  if (!user.mspId || !user.credentials) {
    throw createError({
      message: "Users mspId or signCert not found!",
      statusCode: 404,
    });
  }

  return {
    mspId: user.mspId,
    credentials: Buffer.from(user.credentials),
  };
}

export function userToPrivateKey(user: User) {
  if (!user.key) {
    throw createError({
      message: "Users private key not found!",
      statusCode: 404,
    });
  }

  const privateKey = crypto.createPrivateKey(user.key);

  return {
    privateKey,
  };
}

export async function newGRPCClient() {
  const tlsCredentials = grpc.credentials.createSsl(
    Buffer.from(fabricConfig.peer.tlsCACerts.pem),
  );

  const client = await new grpc.Client(
    fabricConfig.peer.url.split("//")[1],
    tlsCredentials,
    {
      "grpc.ssl_target_name_override":
        fabricConfig.peer.grpcOptions["ssl-target-name-override"],
    },
  );

  return client;
}

const BuildIdentity = async (event: H3Event) => {
  try {
    const session = await useSession<AuthSession>(event, sessionConfig);

    if (!session.data.username) {
      throw createError({
        message: "Not Authorized",
        statusCode: 401,
      });
    }

    const user = await findUserByUsername(session.data.username);
    const identity = userToIdentity(user);
    const { privateKey } = userToPrivateKey(user);
    const signer = signers.newPrivateKeySigner(privateKey);

    return {
      identity,
      signer,
    };
  } catch (error) {
    const publicUser: User = {
      id: "",
      userId: "",
      createdAt: "",
      username: "",
      password: "",
      mspId: fabricConfig.public.mspId,
      credentials: fabricConfig.public.credentials,
      key: fabricConfig.public.key,
    };
    const identity = userToIdentity(publicUser);
    const { privateKey } = userToPrivateKey(publicUser);
    const signer = signers.newPrivateKeySigner(privateKey);

    return {
      identity,
      signer,
    };
  }
};

const BuildGateway = async (event: H3Event) => {
  const { identity, signer } = await BuildIdentity(event);

  const client = await newGRPCClient();
  const gateway = connect({ client, identity, signer });

  return {
    gateway,
    client,
    [Symbol.asyncDispose]: async () => {
      console.log("closing gateway");
      await gateway.close();
      await client.close();
    },
  };
};

export const useChaincode = async <T extends H3Event>(event: T) => {
  const connection = await BuildGateway(event);

  const network = await connection.gateway.getNetwork(
    fabricConfig.chaincode.channel,
  );
  const contract = await network.getContract(fabricConfig.chaincode.chaincode);

  const service = new common.generic.GenericServiceClient(
    contract,
    GlobalRegistry,
  );

  return {
    connection,
    contract,
    service,
  };
};
