export interface Workload {
  test: Test
  monitors?: Monitors
}

export interface Monitors {
  transaction?: Transaction[]
  resource: Resource[]
}
export interface Resource {
  module: Module
  options: ResourceOptions
}

export enum Module {
  Docker = 'docker',
  Process = 'process',
  Prometheus = 'prometheus',
  PrometheusPush = 'prometheus-push',
}

export interface Test {
  name?: string
  description?: null | string
  workers: Workers
  rounds: Round[]
}

export interface Round {
  label: string
  contractId?: string
  txDuration?: number
  rateControl: RateControl
  workload: WorkloadClass
  description?: string
  chaincodeID?: string
  txNumber?: number
}
export interface WorkloadClass {
  module: string
  arguments?: Arguments
}

export interface Arguments {
}

export interface ResourceOptions {
  interval?: number
  containers?: string[]
  url?: string
  charting?: Charting
  metrics?: Metrics
  processes?: Process[]
}

export interface Charting {
  polar: Bar
  bar: Bar
}

export interface Bar {
  metrics: Metric[]
}

export enum Metric {
  All = 'all',
  MaxMemoryMB = 'Max Memory (MB)',
}

export interface Metrics {
  include: Include[]
  queries: Query[]
}

export enum Include {
  Couch = 'couch',
  Dev = 'dev-.*',
  IncludeDev = 'dev*',
  IncludePeer = 'peer*',
  Orderer = 'orderer',
  Peer = 'peer',
}

export interface Query {
  name: Name
  query: string
  step: number
  label: Label
  statistic: MultiOutput
  multiplier?: number
}

export enum Label {
  Instance = 'instance',
  Name = 'name',
}

export enum Name {
  AvgMemoryMB = 'Avg Memory (MB)',
  CPU = 'CPU (%)',
  DiscReadMB = 'Disc Read (MB)',
  DiscWriteMB = 'Disc Write (MB)',
  EndorseTimeS = 'Endorse Time (s)',
  MaxMemoryMB = 'Max Memory (MB)',
  NetworkInMB = 'Network In (MB)',
  NetworkOutMB = 'Network Out (MB)',
}

export enum MultiOutput {
  Avg = 'avg',
  Max = 'max',
  Sum = 'sum',
}

export interface Process {
  command: string
  arguments: string
  multiOutput: MultiOutput
}

export interface Transaction {
  module: string
  options?: TransactionOptions
}

export interface TransactionOptions {
  pushInterval?: number
  pushUrl?: string
  loggerModuleName?: string
  messageLevel?: string
}

export interface RateControl {
  type: Type
  opts: RateControlOpts
}

export interface RateControlOpts {
  tps?: number
  transactionLoad?: number
  startingTps?: number
  finishingTps?: number
  weights?: number[]
  rateControllers?: RateController[]
  startTps?: number
}

export interface RateController {
  type: Type
  opts: RateControllerOpts
}

export interface RateControllerOpts {
  tps: number
}

export enum Type {
  CompositeRate = 'composite-rate',
  FixedLoad = 'fixed-load',
  FixedRate = 'fixed-rate',
  LinearRate = 'linear-rate',
}

export interface Write {
  count: number
  writeMode: string
}

export interface Workers {
  number: number
}
