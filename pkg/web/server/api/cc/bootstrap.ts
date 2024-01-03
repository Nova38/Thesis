import { common, auth } from "saacs-es";

export default defineEventHandler(async (event) => {
  const bootstrapRequest = new common.generic.BootstrapRequest({
    collections: [
      new auth.objects.Collection({
        name: "TestCollection",
        collectionId: "TestCollection",
        itemTypes: ["sample"],
        default: {},
        authType: auth.auth.AuthType.ROLE,
      }),
    ],
  });

  const cc = await useChaincode(event);
  const bootstrap = await cc.service.bootstrap(bootstrapRequest);
  console.log(bootstrap);
  return bootstrap;
});
