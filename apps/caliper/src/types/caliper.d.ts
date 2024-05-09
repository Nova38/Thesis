/// <reference types="node" />

declare module '@hyperledger/caliper-core' {
  export class ConnectorInterface extends EventEmitter {
    /**
     * Retrieve the target SUT type.
     * @return {string} The target SUT type.
     */
    getType(): string
    /**
     * Retrieve the worker index.
     * @return {number} The zero-based worker index.
     */
    getWorkerIndex(): number
    /**
     * Initialize the connector and potentially the SUT.
     * @param {boolean} workerInit Indicates whether the initialization happens in the worker process.
     * @async
     */
    init(workerInit: boolean): Promise<void>
    /**
     * Deploy the configured smart contracts to the SUT if applicable.
     * @async
     */
    installSmartContract(): Promise<void>
    /**
     * Return required arguments for reach workers process, e.g., return information generated during an admin phase, such as contract installation.
     * Information returned here is passed to the workers through the messaging protocol on test.
     * @param {number} number The total number of worker processes.
     * @return {Promise<object[]>} Array of data objects, one for each worker process.
     * @async
     */
    prepareWorkerArguments(number: number): Promise<object[]>
    /**
     * Get a connector-specific context object for the workload module of the given round.
     * @param {number} roundIndex The zero-based round index of the test.
     * @param {object} args Arguments prepared by the connector's {@link prepareWorkerArguments} method in the manager process.
     * @return {Promise<object>} The prepared context object for the workload module.
     * @async
     */
    getContext(roundIndex: number, args: object): Promise<object>
    /**
     * Release the current context and related resources previously prepared by the {@link getContext} method.
     * @async
     */
    releaseContext(): Promise<void>
    /**
     * Send one or more requests to the backing SUT.
     * @param {object|object[]} requests The object (or array of objects) containing the options of the request(s).
     * @return {Promise<TxStatus|TxStatus[]>} The result (or an array of them) of the executed request(s).
     * @async
     */
    sendRequests(requests: object | object[]): Promise<TxStatus | TxStatus[]>
    /**
     * Throws and error to signal a not implemented method.
     * @param {string} method The name of the method.
     * @private
     */
    private _throwNotImplemented
  }
  export class ConnectorBase extends ConnectorInterface {
    /**
     * Constructor
     * @param {number} workerIndex The zero-based worker index.
     * @param {string} bcType The target SUT type.
     */
    constructor(workerIndex: number, bcType: string)
    workerIndex: number
    bcType: string
    /**
     * Send one or more requests to the backing SUT.
     * The default implementation handles batching and TX events,
     * and delegates to the {@link _sendSingleRequest} method.
     * @param {object|object[]} requests The object (or array of objects) containing the options of the request(s).
     * @return {Promise<TxStatus|TxStatus[]>} The result (or an array of them) of the executed request(s).
     * @async
     */
    sendRequests(requests: object | object[]): Promise<TxStatus | TxStatus[]>
    /**
     * Raises the "txsSubmitted" event.
     * @param {number} count The number of TXs submitted. Passed to the raised event.
     * @private
     */
    private _onTxsSubmitted
    /**
     * Raises the "txsFinished" event.
     * @param {TxStatus|TxStatus[]} txResults The (array of) TX result(s). Passed to the raised event.
     * @private
     */
    private _onTxsFinished
    /**
     * Send a request to the backing SUT. Must be overridden by derived classes.
     * @param {object} request The object containing the options of the request.
     * @return {Promise<TxStatus>} The array of data about the executed requests.
     * @protected
     * @async
     */
    protected _sendSingleRequest(request: object): Promise<TxStatus>
  }
  export const CaliperWorker: typeof import('./lib/worker/caliper-worker')
  export const Constants: {
    Messages: {
      Types: {
        Register: string
        Connected: string
        AssignId: string
        Assigned: string
        Initialize: string
        Ready: string
        Prepare: string
        Prepared: string
        Test: string
        TxReset: string
        TxUpdate: string
        TestResult: string
        Exit: string
      }
      Targets: {
        All: string
      }
    }
    Factories: {
      Connector: string
      WorkerMessenger: string
      ManagerMessenger: string
    }
    Events: {
      Connector: {
        TxsSubmitted: string
        TxsFinished: string
      }
    }
    AuthComponents: {
      PushGateway: string
      Prometheus: string
    }
  }
  export class TxStatus {
    /**
     * Constructor
     * @param {string} id, transaction id
     */
    constructor(id: string)
    status: {
      id: string
      status: string
      time_create: number
      time_final: number
      result: any
      verified: boolean
      flags: number
      error_messages: any[]
      custom_data: Map<any, any>
    }
    /**
     * Getter of the tx id
     * @return {string}, id
     */
    GetID(): string
    /**
     * Setter of the tx id
     * @param {string} id, id
     */
    SetID(id: string): void
    /**
     * Getter of the tx status
     * @return {string}, status
     */
    GetStatus(): string
    /**
     * Check if the tx has been committed successfully
     * @return {boolean} committed or not
     */
    IsCommitted(): boolean
    /**
     * Set the tx status to 'success'
     * The 'time_final' will also be recorded
     * @param {number} time The epoch time to record for the status change.
     */
    SetStatusSuccess(time: number): void
    /**
     * Getter of the tx creating time
     * @return {int} create time in ms
     */
    GetTimeCreate(): int
    /**
     * Set a new creation time
     * @param {int} newCreationTime new creation time in ms
     */
    SetTimeCreate(newCreationTime: int): void
    /**
     * Getter of the tx final time
     * @return {int} final time in ms
     */
    GetTimeFinal(): int
    /**
     * Set the tx status to 'failed'
     * The 'time_final' will also be recorded
     */
    SetStatusFail(): void
    /**
     * Check if the tx status is verified
     * @return {boolean}, verified or not
     */
    IsVerified(): boolean
    /**
     * Setter of the tx verification state
     * @param {*} isVerified, verified or not
     */
    SetVerification(isVerified: any): void
    /**
     * Getter of the blockchain specified flag
     * @return {any}, flag
     */
    GetFlag(): any
    /**
     * Setter of the blockchain specified flag
     * @param {any} flag, flag to be set
     */
    SetFlag(flag: any): void
    /**
     * Setter of the error message
     * @param {int} idx, index of the error message
     * @param {any} msg, message to be stored
     */
    SetErrMsg(idx: int, msg: any): void
    /**
     * Getter of the error messages
     * @return {Array}, stored messages
     */
    GetErrMsg(): any[]
    /**
     * Setter of blockchain specified submitting result
     * @param {any} result result
     */
    SetResult(result: any): void
    /**
     * Getter of stored submitting result
     * @return {any} result
     */
    GetResult(): any
    /**
     * Getter of the status object
     * @return {object} status object
     */
    Marshal(): object
    /**
     * Set a custom key/value. This sets a key with the custom subsection of this
     * status instance and cannot be used to modify internal keys such as `time_create` or `time_final` for example.
     * @param {string} key key
     * @param {any} value value
     */
    Set(key: string, value: any): void
    /**
     * Get a custom value from the key
     * @param {string} key key
     * @return {any} value
     */
    Get(key: string): any
    /**
     * Gets the Map of custom date recorded for the transaction.
     * @return {Map<string, object>} The map of custom data.
     */
    GetCustomData(): Map<string, object>
  }

  export class CaliperUtils {
    /**
     * Indicates whether the process is a forked/child process, i.e., it has a parent process.
     * @return {boolean} True, if the process has a parent process. Otherwise, false.
     */
    static isForkedProcess(): boolean
    /**
     * Assert that the core configuration file paths are set and exist.
     */
    static assertConfigurationFilePaths(): void
    /**
     * Get the mapping of simple builtin connector names to fully qualified package names.
     * @return {Map<string, string>} The mapping from simple names to package names.
     */
    static getBuiltinConnectorPackageNames(): Map<string, string>
    /**
     * Get the mapping of simple builtin messenger names to fully qualified factory paths.
     * @return {Map<string, string>} The mapping from simple names to factory paths.
     */
    static getBuiltinMessengers(): Map<string, string>
    /**
     * Loads the module at the given path.
     * @param {string} modulePath The path to the module or its name.
     * @param {Function} requireFunction The "require" function (with appropriate scoping) to use to load the module.
     * @return {object} The loaded module.
     */
    static loadModule(modulePath: string, requireFunction?: Function): object
    /**
     * Loads the given function from the given module.
     * @param {object} module The module exporting the function.
     * @param {string} functionName The name of the function.
     * @param {string} moduleName The name of the module.
     * @return {Function} The loaded function.
     */
    static loadFunction(
      module: object,
      functionName: string,
      moduleName: string,
    ): Function
    /**
     * Loads the given function from the module at the given path.
     * @param {Map<string, string>} builtInModules The mapping of built-in module names to their path.
     * @param {string} moduleName The name of the module.
     * @param {string} functionName The name of the function.
     * @param {Function} requireFunction The "require" function (with appropriate scoping) to use to load the module.
     * @return {Function} The loaded function.
     */
    static loadModuleFunction(
      builtInModules: Map<string, string>,
      moduleName: string,
      functionName: string,
      requireFunction?: Function,
    ): Function
    /**
     * Check if a named module is installed and accessible by caliper
     * @param {string} moduleName the name of the module to check
     * @param {Function} requireFunction The "require" function (with appropriate scoping) to use to load the module.
     * @returns {boolean} boolean value for existence of accessible package
     */
    static moduleIsInstalled(
      moduleName: string,
      requireFunction?: Function,
    ): boolean
    /**
     * Utility function to check for singleton values
     * @param {string[]} passedArgs arguments passed by user
     * @param {string[]} uniqueArgs arguments that must be unique
     * @returns {boolean} boolean true if passes check
     */
    static checkSingleton(passedArgs: string[], uniqueArgs: string[]): boolean
    /**
     * Perform a sleep
     * @param {*} ms the time to sleep, in ms
     * @returns {Promise} a completed promise
     */
    static sleep(ms: any): Promise<any>
    /**
     * Simple log method to output to the console
     * @param {any} msg messages to log
     */
    static log(...msg: any): void
    /**
     * Returns a logger configured with the given module name.
     * @param {string} name The name of module who will use the logger.
     * @returns {Logger} The configured logger instance.
     */
    static getLogger(name: string): Logger
    /**
     * Creates an absolute path from the provided relative path if necessary.
     * @param {string} relOrAbsPath The relative or absolute path to convert to an absolute path.
     *                              Relative paths are considered relative to the Caliper root folder.
     * @param {string} root_path root path to use
     * @return {string} The resolved absolute path.
     */
    static resolvePath(relOrAbsPath: string, root_path?: string): string
    /**
     * parse a yaml file.
     * @param {string} filenameOrFilePath the yaml file path
     * @return {object} the parsed data.
     */
    static parseYaml(filenameOrFilePath: string): object
    /**
     * Convert an object to YAML string.
     * @param {object} obj The object to stringify.
     * @return {string} The string YAML content.
     */
    static stringifyYaml(obj: object): string
    /**
     * Parse a YAML conform string into an object.
     * @param {string} stringContent The YAML content.
     * @return {object} The parsed object.
     */
    static parseYamlString(stringContent: string): object
    /**
     * Checks whether the given object is defined and not null.
     * @param {object} object The object to check.
     * @return {boolean} True, if the object is defined and not null. Otherwise false.
     */
    static checkDefined(object: object): boolean
    /**
     * Throws an error if the object is undefined or null.
     * @param {object} object The object to check.
     * @param {string} msg Optional error message to throw in case of unsuccessful check.
     */
    static assertDefined(object: object, msg: string): void
    /**
     * Checks whether the property exists on the object and it isn't undefined or null.
     * @param {object} object The object to check for the property.
     * @param {string} propertyName The name of the property to check.
     * @return {boolean} True, if the property exists and it's defined and not null. Otherwise false.
     */
    static checkProperty(object: object, propertyName: string): boolean
    /**
     * Throws an error if the property doesn't exists on the object or it's undefined or null.
     * @param {object} object The object to check for the property.
     * @param {string} objectName Optional error message to throw in case of an unsuccessful check.
     * @param {string} propertyName The name of the property to check.
     */
    static assertProperty(
      object: object,
      objectName: string,
      propertyName: string,
    ): void
    /**
     * Checks whether any of the given properties exists and is defined and is not null on the given object.
     * @param {object} object The object to check for the properties.
     * @param {string[]} propertyNames The list of property names to check.
     * @return {boolean} True if any of the given properties exists on the object and is defined and is not null. Otherwise false.
     */
    static checkAnyProperty(object: object, ...propertyNames: string[]): boolean
    /**
     * Throws an error if none of the given properties exists and is defined and is not null on the given object.
     * @param {object} object The object to check for the properties.
     * @param {string} objectName The name of the object
     * @param {string[]} propertyNames The list of property names to check.
     */
    static assertAnyProperty(
      object: object,
      objectName: string,
      ...propertyNames: string[]
    ): void
    /**
     * Checks whether all of the given properties exist and are defined and are not null on the given object.
     * @param {object} object The object to check for the properties.
     * @param {string[]} propertyNames The list of property names to check.
     * @return {boolean} True if all of the given properties exist on the object and are defined and are not null. Otherwise false.
     */
    static checkAllProperties(
      object: object,
      ...propertyNames: string[]
    ): boolean
    /**
     * Throws an error if  any of the given properties exists and is defined and is not null on the given object.
     * @param {object} object The object to check for the properties.
     * @param {string} objectName The name of the object for the error message.
     * @param {string[]} propertyNames The list of property names to check.
     */
    static assertAllProperties(
      object: object,
      objectName: string,
      ...propertyNames: string[]
    ): void
    /**
     * Executes the given command asynchronously.
     * @param {string} command The command to execute through a newly spawn shell.
     * @param {boolean} logAction Boolean flag to inform the command being run, default true
     * @return {Promise} The return promise is resolved upon the successful execution of the command, or rejected with an Error instance.
     * @async
     */
    static execAsync(command: string, logAction?: boolean): Promise<any>
    /**
     * Invokes a given command in a spawned child process and attaches all standard IO.
     * @param {string} cmd The command to be run.
     * @param {string[]} args The array of arguments to pass to the command.
     * @param {Map<string, string>} env The key-value pairs of environment variables to set.
     * @param {string} cwd The current working directory to set.
     * @returns {Promise} A Promise that is resolved or rejected.
     */
    static invokeCommand(
      cmd: string,
      args: string[],
      env: Map<string, string>,
      cwd: string,
    ): Promise<any>
    /**
     * Invokes a given command in a spawned child process and returns its output.
     * @param {string} cmd The command to be run.
     * @param {string[]} args The array of arguments to pass to the command.
     * @param {object} env The key-value pairs of environment variables to set.
     * @param {string} cwd The current working directory to set.
     * @returns {Promise} A Promise that is resolved with the command output or rejected with an Error.
     */
    static getCommandOutput(
      cmd: string,
      args: string[],
      env?: object,
      cwd?: string,
    ): Promise<any>
    /**
     * Retrieve user specified flow flags
     * @returns {JSON} a JSON object containing conditioned flow options
     */
    static getFlowOptions(): JSON
    /**
     * Convert milliseconds to seconds
     * @param {number} value to convert
     * @returns {number} the converted value
     */
    static millisToSeconds(value: number): number
    /**
     * Augment the passed URL with basic auth if the settings are present
     * @param {string} urlPath the URL to augment
     * @param {string} component the component being augmented
     * @returns {string} the URL to be used, which may have been augmented with basic auth
     */
    static augmentUrlWithBasicAuth(urlPath: string, component: string): string
  }

  export class Version {
    /**
     * Creates and initializes a new instance of the Version class.
     * @param {string} versionString The version string to encapsulate.
     */
    constructor(versionString: string)
    versionString: string
    /**
     * Checks whether this version equals to the given version.
     * @param {string} version The version string to compare against.
     * @return {boolean} True, if this version equals to the given version. Otherwise false.
     */
    equalsTo(version: string): boolean
    /**
     * Checks whether this version is strictly greater than the given version.
     * @param {string} version The version string to compare against.
     * @return {boolean} True, if this version is strictly greater than the given version. Otherwise false.
     */
    greaterThan(version: string): boolean
    /**
     * Checks whether this version is greater than or equals to the given version.
     * @param {string} version The version string to compare against.
     * @return {boolean} True, if this version is greater than or equals to the given version. Otherwise false.
     */
    greaterThanOrEqualsTo(version: string): boolean
    /**
     * Checks whether this version is strictly less than the given version.
     * @param {string} version The version string to compare against.
     * @return {boolean} True, if this version is strictly less than the given version. Otherwise false.
     */
    lessThan(version: string): boolean
    /**
     * Checks whether this version is less than or equals to the given version.
     * @param {string} version The version string to compare against.
     * @return {boolean} True, if this version is less than or equals to the given version. Otherwise false.
     */
    lessThanOrEqualsTo(version: string): boolean
    /**
     * Returns the string representation of the version.
     * @return {string} The string representation of the version.
     */
    toString(): string
  }

  export declare class WorkloadModuleInterface {
    /**
     * Initialize the workload module with the given parameters.
     * @param {number} workerIndex The 0-based index of the worker instantiating the workload module.
     * @param {number} totalWorkers The total number of workers participating in the round.
     * @param {number} roundIndex The 0-based index of the currently executing round.
     * @param {object} roundArguments The user-provided arguments for the round from the benchmark configuration file.
     * @param {ConnectorBase} sutAdapter The adapter of the underlying SUT.
     * @param {object} sutContext The custom context object provided by the SUT adapter.
     * @async
     */
    initializeWorkloadModule(
      workerIndex: number,
      totalWorkers: number,
      roundIndex: number,
      roundArguments: object,
      sutAdapter: ConnectorBase,
      sutContext: object,
    ): Promise<void>
    /**
     * Assemble the next TX content(s) and submit it to the SUT adapter.
     * @async
     */
    submitTransaction(): Promise<void>
    /**
     * Clean up the workload module at the end of the round.
     * @async
     */
    cleanupWorkloadModule(): Promise<void>
  }

  export declare class WorkloadModuleBase extends WorkloadModuleInterface {
    /**
     * The 0-based index of the worker instantiating the workload module.
     * @type {number}
     */
    workerIndex: number
    /**
     * The total number of workers participating in the round.
     * @type {number}
     */
    totalWorkers: number
    /**
     * The 0-based index of the currently executing round.
     * @type {number}
     */
    roundIndex: number
    /**
     * The user-provided arguments for the round from the benchmark configuration file.
     * @type {object}
     */
    roundArguments: any
    /**
     * The adapter of the underlying SUT.
     * @type {ConnectorBase}
     */
    sutAdapter: ConnectorBase
    /**
     * The custom context object provided by the SUT adapter.
     * @type {object}
     */
    sutContext: any
  }

  export declare class CaliperEngine {
    /**
     * Initializes the CaliperEngine instance.
     * @param {object} benchmarkConfig The benchmark configuration object.
     * @param {object} networkConfig The network configuration object.
     * @param {Function} adapterFactory The factory function for creating an adapter instance.
     */
    constructor(
      benchmarkConfig: object,
      networkConfig: object,
      adapterFactory: Function,
    )
    benchmarkConfig: any
    networkConfig: any
    workspace: any
    returnCode: number
    adapterFactory: Function
    /**
     * Executes the given start/end command with proper checking and error handling.
     * @param {string} commandName Either "start" or "end". Used in the error messages.
     * @param {number} errorStatusStart The last taken error status code. Execution errors will use the next 3 status code.
     * @private
     */
    private _executeCommand
    /**
     * Run the benchmark based on passed arguments
     * @returns {number} the error status of the run
     */
    run(): number
    /**
     * Stops the benchmark run
     */
    stop(): Promise<void>
  }

  export const ConfigUtil: unknown
  export const WorkerMessageHandler: typeof import('./lib/worker/worker-message-handler')
  export const MessengerInterface: typeof import('./lib/common/messengers/messenger-interface')
  export const MonitorOrchestrator: typeof import('./lib/manager/orchestrators/monitor-orchestrator')
  export const RoundOrchestrator: typeof import('./lib/manager/orchestrators/round-orchestrator')
  export const WorkerOrchestrator: typeof import('./lib/manager/orchestrators/worker-orchestrator')
}

declare module '@hyperledger/caliper-fabric' {
  export declare const ConnectorFactory: (
    workerIndex: number,
  ) => Promise<BlockchainConnector>
}
