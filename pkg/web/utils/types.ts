import type { PlainMessage } from '@bufbuild/protobuf'
import type { EmptyOrFlatSpecimenKeys } from './objects/flatten'
import type { ccbio } from '~/lib'

export type BulkMode = 'import' | 'update' | 'hybrid'

export interface BulkImportHeader {
  name: string
  field: string | ((row: UpdateRawRow) => string)
  mapped: string
  isID: boolean
  updated: boolean
}

export interface ImportRowMeta {
  exist: ExistStatus
  status: ProcessingStatus
  statusMessage?: string
}

export interface UpdateRowMeta {
  id: number
  exist: ExistStatus
  status: ProcessingStatus
  catalogNumber?: string
  // currentSpecimen?: ccbio.Specimen
  differences?: Record<string, any>
  statusMessage?: string
}

/**
 * UpdateRow is the row that will be used to update the existing specimen
 * in the database, the first column is the specimenId and the rest of the
 * columns are the raw data from the spreadsheet
 */
export type IdMappedRow = Record<string, Record<string, string>>

export interface UpdateRawRow {
  id: string
  raw: Record<string, string>
}

export interface CSVImportMetadata {
  headers: string[]
  specimenIdHeader: string
  rows: Record<string, string>[]
}

export type ImportColType = 'id' | 'meta' | 'raw'
export interface ImportColumns {
  id: ImportCol
  meta?: ImportCol[]
  raw?: ImportCol[]
}
export interface ImportCol {
  name: string
  label: string
  field_path?: string
  field: ((row: UpdateRawRow) => string) | string
  colType: ImportColType
}

export type ProcessingStatus =
  | 'error'
  | 'loading'
  | 'new'
  | 'pre-existing'
  | 'success'
export type ExistStatus = 'new' | 'pre-existing' | 'unknown'
export type PlainSpecimen = PlainMessage<ccbio.Specimen>
