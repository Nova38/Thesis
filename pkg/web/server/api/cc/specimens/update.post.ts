import { common, auth, ccbio } from "saacs-es";
import { Any, FieldMask } from "@bufbuild/protobuf";

export type bodySchema = { specimen: ccbio.Specimen; mask: FieldMask };

export default defineEventHandler(async (event) => {
  const cc = await useChaincode(event);

  try {
    const body = await readBody(event);
    console.log();
    console.log({ body });

    const specimen = new ccbio.Specimen().fromJson(body.specimen);
    console.log({ specimen });
    const value = Any.pack(specimen);
    console.log({ value });

    const req = new common.generic.UpdateRequest({
      item: {
        key: new auth.objects.ItemKey({
          collectionId: specimen.collectionId,
          itemType: ccbio.Specimen.typeName,
          itemKeyParts: [specimen.specimenId],
        }),
        value: value,
      },
      updateMask: body.mask,
    });
    const result = await cc.service.update(req);

    console.log(result);
    const unpacked = new ccbio.Specimen();
    result.item?.value?.unpackTo(unpacked);

    return unpacked;
    return req.toJson({ typeRegistry: cc.service.registry });
  } catch (error) {
    console.log(error);
    throw error;
  }
});