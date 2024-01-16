import { common, auth, ccbio } from "saacs-es";
import { z } from "zod";

const querySchema = z.object({
  collectionId: z.string(),
  limit: z.number().optional(),
});

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event);

  const query = await getValidatedQuery(event, (body) =>
    querySchema.safeParse(body),
  );
  console.log(query);
  if (!query.success) throw query.error.issues;

  let bookmark = "";
  let lastBookmark = "-";
  const specimenMap: Record<string, any> = {};

  while (bookmark !== lastBookmark) {
    const result = await cc.service.listByAttrs(
      new common.generic.ListByAttrsRequest({
        key: new auth.objects.ItemKey({
          collectionId: query.data.collectionId,
          itemType: ccbio.Specimen.typeName,
          itemKeyParts: [query.data.collectionId],
        }),
        numAttrs: 0,
        limit: query.data.limit ?? 1000,
        bookmark: bookmark ?? "",
      }),
    );
    lastBookmark = bookmark;
    bookmark = result.bookmark;

    result.items.forEach((i) => {
      const s = new ccbio.Specimen();
      i.value?.unpackTo(s);
      specimenMap[s.specimenId] = s.toJson({ emitDefaultValues: true });
    });

    if (lastBookmark === bookmark) {
      console.log("no change", bookmark, lastBookmark);
      break;
    }
  }
  // console.log("1");

  // console.log("2");

  // console.log(result);

  return { specimenMap, bookmark: bookmark };
});
