import { useChaincode } from "~/server/utils/useChaincode";
import { common, auth } from "saacs-es";

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event);
  const result = await cc.service.get(
    new common.generic.GetRequest({
      key: new auth.objects.ItemKey({
        collectionId: "TestCollection",
        itemType: "auth.Collection",
        itemKeyParts: ["collection_id"],
      }),
    }),
  );

  const c = new auth.objects.Collection();
  result.item?.value?.unpackTo(c);
  return c;
});
