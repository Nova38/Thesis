import type { PlainMessage } from '@bufbuild/protobuf'
import type { ccbio } from '~/lib'

export interface UpdateRowMeta {
  id: number
  exist: ExistStatus
  status: ProcessingStatus
  catalogNumber?: string
  currentSpecimen?: ccbio.Specimen
  differences?: Record<string, any>
  statusMessage?: string
}

export interface UpdateRow {
  id: number
  meta?: UpdateRowMeta
  raw: Record<string, string>
}

export type ProcessingStatus =
  | 'error'
  | 'loading'
  | 'new'
  | 'pre-existing'
  | 'success'
export type ExistStatus = 'new' | 'pre-existing' | 'unknown'
export type PlainSpecimen = PlainMessage<ccbio.Specimen>
