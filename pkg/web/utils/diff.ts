import { diff } from "ohash";
import { FieldMask } from "@bufbuild/protobuf";

export const diffToFieldMaskPath = (base: any, updated: any) => {
  const paths = diff(base, updated).map((d) => d.key);

  const mask = new FieldMask();
  mask.paths = paths;

  return { paths };
};
