// https://decipher.dev/30-seconds-of-typescript/docs/flattenObject/
export const FlattenObject = (obj, prefix = "") =>
  Object.keys(obj).reduce((acc, k) => {
    const pre = prefix.length ? prefix + "." : "";
    if (typeof obj[k] === "object")
      Object.assign(acc, FlattenObject(obj[k], pre + k));
    else acc[pre + k] = obj[k];
    return acc;
  }, {});

export const FlattenEmptySpecimen = () => {
  return FlattenObject(
    JSON.parse(MakeEmptySpecimen().toJsonString({ emitDefaultValues: true })),
  );
};

export const FlattedSpecimenKeys = () => {
  return Object.keys(FlattenEmptySpecimen());
};

const structure = {
  primary: {
    catalogNumber: "",
    accessionNumber: "",
    fieldNumber: "",
    tissueNumber: "",
    cataloger: "",
    collector: "",
    determiner: "",
    fieldDate: {
      verbatim: "",
      year: 0,
      month: "",
      day: 0,
    },
    catalogDate: {
      verbatim: "",
      year: 0,
      month: "",
      day: 0,
    },
    determinedDate: {
      verbatim: "",
      year: 0,
      month: "",
      day: 0,
    },
    determinedReason: "",
    originalDate: {
      verbatim: "",
      year: 0,
      month: "",
      day: 0,
    },
  },
  secondary: {
    sex: "SEX_UNDEFINED",
    age: "AGE_UNDEFINED",
    weight: 0,
    weightUnits: "",
    preparations: {},
    condition: "",
    molt: "",
    notes: "",
  },
  taxon: {
    kingdom: "",
    phylum: "",
    class: "",
    order: "",
    family: "",
    genus: "",
    species: "",
    subspecies: "",
  },
  georeference: {
    country: "",
    stateProvince: "",
    county: "",
    locality: "",
    latitude: 0,
    longitude: 0,
    habitat: "",
    continent: "",
    locationRemarks: "",
    coordinateUncertaintyInMeters: 0,
    georeferenceBy: "",
    georeferenceDate: {
      verbatim: "",
      year: 0,
      month: "",
      day: 0,
    },
    georeferenceProtocol: "",
    geodeticDatum: "",
    footprintWkt: "",
    notes: "",
  },
  images: {},
  loans: {},
  grants: {},
};
