import type { PlainMessage } from '@bufbuild/protobuf'
import { Timestamp } from '@bufbuild/protobuf'
import { pb } from '@saacs/saacs-pb'

export function NormalizeSpecimen(input: PlainMessage<pb.Specimen>): pb.Specimen {
  const response = new pb.Specimen(input)

  // Make sure all sub felids are initialized
  response.primary ??= new pb.Specimen_Primary()
  response.secondary ??= new pb.Specimen_Secondary()
  response.taxon ??= new pb.Specimen_Taxon()
  response.georeference ??= new pb.Specimen_Georeference()
  response.grants ??= {}
  response.loans ??= {}
  response.images ??= {}

  // Make sure all of the protoDate objects are initialized
  response.primary.catalogDate ??= new pb.Date({ timestamp: new Timestamp() })
  response.primary.determinedDate ??= new pb.Date({ timestamp: new Timestamp() })
  response.primary.fieldDate ??= new pb.Date({ timestamp: new Timestamp() })
  response.primary.originalDate ??= new pb.Date({ timestamp: new Timestamp() })
  response.georeference.georeferenceDate ??= new pb.Date({ timestamp: new Timestamp() })
  response.georeference.georeferenceDate ??= new pb.Date({ timestamp: new Timestamp() })

  // Set the last modified to exist
  response.primary.lastModified ??= new pb.StateActivity()
  response.secondary.lastModified ??= new pb.StateActivity()
  response.taxon.lastModified ??= new pb.StateActivity()
  response.georeference.lastModified ??= new pb.StateActivity()

  // Check the iterable elements
  response.secondary.preparations ??= {}
  for (const key in response.secondary.preparations)
    response.secondary.preparations[key].verbatim ??= ''

  for (const key in response.grants) {
    response.grants[key] ??= new pb.Specimen_Grant()
    response.grants[key].grantedDate ??= new pb.Date({ timestamp: new Timestamp() })
    response.grants[key].lastModified ??= new pb.StateActivity({ timestamp: new Timestamp() })
  }

  for (const key in response.loans) {
    response.loans[key] ??= new pb.Specimen_Loan()
    response.loans[key].loanedDate ??= new pb.Date({ timestamp: new Timestamp() })
    response.loans[key].lastModified ??= new pb.StateActivity({ timestamp: new Timestamp() })
  }

  return response
}
