import * as factory from './factory.js';
import * as fakers from './fakers.js';
import * as registry from './registry.js';


const GlobalRegistry = registry.GlobalRegistry;
export const utils = {
    factory,
    fakers,
    registry,
    GlobalRegistry
};
