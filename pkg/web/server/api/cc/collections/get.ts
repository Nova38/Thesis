// import { useChaincode } from "~/server/utils/useChaincode";
import { common, auth } from "saacs-es";
import { z } from "zod";

const bodySchema = z.object({
  collectionId: z.string(),
});
export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event);

  const query = await getValidatedQuery(event, (body) =>
    bodySchema.safeParse(body),
  );
  console.log(query);
  if (!query.success) throw query.error.issues;

  const result = await cc.service.get(
    new common.generic.GetRequest({
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemType: "auth.Collection",
        itemKeyParts: ["collection_id"],
      }),
    }),
  );

  const c = new auth.objects.Collection();
  result.item?.value?.unpackTo(c);
  return c;
});
