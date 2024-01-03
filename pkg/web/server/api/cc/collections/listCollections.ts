import { pb } from "saacs-es";
import { common, auth } from "saacs-es";

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event);

  const result = await cc.service.getCollectionsList();

  return result;

  const cols = [
    new pb.auth.objects.Collection({
      collectionId: "Test 1",
      name: "Test Collection 1",
      authType: pb.auth.auth.AuthType.ROLE,
    }),
    new pb.auth.objects.Collection({
      collectionId: "Test 2",
      name: "Test Collection 2",
      authType: pb.auth.auth.AuthType.ROLE,
    }),
  ];

  return cols;
});
