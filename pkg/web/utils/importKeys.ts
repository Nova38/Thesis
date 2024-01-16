import { crush, get, set, keys } from "radash";
import { ccbio } from "saacs-es";

const keysForImport = FlattenObject(MakeEmptySpecimen());
console.log("keysForImport", keysForImport);
//   ^?

const notAllowedKeys = ["specimenId", "last_modified_by", "collection_id"];
const SpecimenKeys: Array<string> = keys(keysForImport);

const FilteredSpecimenKeys = SpecimenKeys.filter(
  (k) => !notAllowedKeys.includes(k),
);

export default () => {
  console.log("keysForImport", keysForImport);
  return "Hello Util";
};
