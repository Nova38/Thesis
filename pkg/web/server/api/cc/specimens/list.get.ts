import { common, auth, ccbio } from "saacs-es";
import { z } from "zod";

const querySchema = z.object({
  collectionId: z.string(),
});

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event);

  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  );
  console.log(query);
  if (!query.success) throw query.error.issues;

  console.log("1");

  const result = await cc.service.listByAttrs(
    new common.generic.ListByAttrsRequest({
      key: new auth.objects.ItemKey({
        collectionId: query.data.collectionId,
        itemType: ccbio.Specimen.typeName,
        itemKeyParts: [query.data.collectionId],
      }),
      numAttrs: 0,
    }),
  );
  console.log("2");

  // console.log(result);
  const specimen = result.items.map((i) => {
    const s = new ccbio.Specimen();
    i.value?.unpackTo(s);
    return s.toJson({ emitDefaultValues: true });
  });

  // console.log(specimen);

  return specimen;
});
