import type { ccbio } from 'saacs-es'
import type { PlainMessage } from '@bufbuild/protobuf'

export interface UpdateRowMeta {
  catalogNumber?: string
  currentSpecimen?: ccbio.Specimen
  differences?: Record<string, any>
  exist: ExistStatus
  status: ProcessingStatus
  statusMessage?: string
  uuid: string
}

export type ProcessingStatus = 'error' | 'loading' | 'new' | 'pre-existing' | 'success'
export type ExistStatus = 'new' | 'pre-existing' | 'unknown'
export type PlainSpecimen = PlainMessage<ccbio.Specimen>
