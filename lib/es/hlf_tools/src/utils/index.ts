import * as factory from './factory';
import * as fakers from './fakers';
import * as registry from './registry';


const GlobalRegistry = registry.GlobalRegistry;
export const utils = {
    factory,
    fakers,
    registry,
    GlobalRegistry
};
