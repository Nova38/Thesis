// @generated by protoc-gen-reg v1 with parameter "target=ts"
// @generated from file chaincode/ccbio/schema/v0/state.proto (package ccbio.schema.v0, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Message } from "@bufbuild/protobuf";
import { Specimen } from "./state_pb.js";

// Specimen 

// Message
// Primary Item:  Specimen

// namecollection_id field }
// namespecimen_id field }
// nameprimary field }
// namesecondary field }
// nametaxon field }
// namegeoreference field }
// nameimages field }
// nameloans field }
// namegrants field }
    // specimen_id   ,
export function SpecimenKey(item : Specimen): string[] {
    attr=[]
 if (!item?.specimenId) {
    return attr
 }
    attr.push(item?.specimenId)
 return attr
}
// Path: specimen_id

// Primary 

// Message
// Secondary 

// Message
// Taxon 

// Message
// Georeference 

// Message
// Image 

// Message
// Loan 

// Message
// Grant 

// Message