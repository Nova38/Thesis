'use strict'
exports.__esModule = true
exports.Type = exports.MultiOutput = exports.Name = exports.Label = exports.Include = exports.Metric = exports.Module = void 0
let Module;
(function (Module) {
  Module.Docker = 'docker'
  Module.Process = 'process'
  Module.Prometheus = 'prometheus'
  Module.PrometheusPush = 'prometheus-push'
})(Module = exports.Module || (exports.Module = {}))
let Metric;
(function (Metric) {
  Metric.All = 'all'
  Metric.MaxMemoryMB = 'Max Memory (MB)'
})(Metric = exports.Metric || (exports.Metric = {}))
let Include;
(function (Include) {
  Include.Couch = 'couch'
  Include.Dev = 'dev-.*'
  Include.IncludeDev = 'dev*'
  Include.IncludePeer = 'peer*'
  Include.Orderer = 'orderer'
  Include.Peer = 'peer'
})(Include = exports.Include || (exports.Include = {}))
let Label;
(function (Label) {
  Label.Instance = 'instance'
  Label.Name = 'name'
})(Label = exports.Label || (exports.Label = {}))
let Name;
(function (Name) {
  Name.AvgMemoryMB = 'Avg Memory (MB)'
  Name.CPU = 'CPU (%)'
  Name.DiscReadMB = 'Disc Read (MB)'
  Name.DiscWriteMB = 'Disc Write (MB)'
  Name.EndorseTimeS = 'Endorse Time (s)'
  Name.MaxMemoryMB = 'Max Memory (MB)'
  Name.NetworkInMB = 'Network In (MB)'
  Name.NetworkOutMB = 'Network Out (MB)'
})(Name = exports.Name || (exports.Name = {}))
let MultiOutput;
(function (MultiOutput) {
  MultiOutput.Avg = 'avg'
  MultiOutput.Max = 'max'
  MultiOutput.Sum = 'sum'
})(MultiOutput = exports.MultiOutput || (exports.MultiOutput = {}))
let Type;
(function (Type) {
  Type.CompositeRate = 'composite-rate'
  Type.FixedLoad = 'fixed-load'
  Type.FixedRate = 'fixed-rate'
  Type.LinearRate = 'linear-rate'
})(Type = exports.Type || (exports.Type = {}))
