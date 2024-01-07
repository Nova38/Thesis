import { ccbio } from "saacs-es";

export const MakeEmptySpecimen = () => {
  return new ccbio.Specimen({
    collectionId: "",
    primary: {
      catalogDate: {},
      determinedDate: {},
      fieldDate: {},
      originalDate: {},
    },
    secondary: {
      preparations: {},
    },
    georeference: {
      georeferenceDate: {},
    },
    grants: {},
    taxon: {},
    loans: {},
  });
};
