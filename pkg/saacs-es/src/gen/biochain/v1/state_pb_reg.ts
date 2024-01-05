import { Date, Researcher, Specimen, Specimen_Georeference, Specimen_Grant, Specimen_Image, Specimen_Loan, Specimen_Primary, Specimen_Secondary, Specimen_Secondary_Preparation, Specimen_Taxon, SpecimenHistory, SpecimenHistoryEntry } from "./state_pb.js";

export const allMessages = [
  SpecimenHistory, 
  SpecimenHistoryEntry, 
  Date, 
  Researcher, 
  Specimen, 
  Specimen_Primary, 
  Specimen_Secondary, 
  Specimen_Secondary_Preparation, 
  Specimen_Taxon, 
  Specimen_Georeference, 
  Specimen_Image, 
  Specimen_Loan, 
  Specimen_Grant, 
];
