/**
 *  This is a scratch pad for building the bulk import functionality.
 *
 *  The bulk import goal is to take a array of flat objects and convert them into
 *    a collection of Specimen objects.
 *
 *  Key points:
 *  - The input data is an array of flat objects with unknown keys at compile time.
 *  - The input will also include a object that maps the flat object keys to the Specimen object keys.
 *  - The output data is a collection of Specimen objects.
 *
 *  - The Specimen object is a complex object with nested objects and arrays.
 *  - None of the Specimen properties should be undefined.
 *
 */
// import { schema } from '../src/lib/ccSchema/index';

// type rawRow = Record<string, string>;

// This is a type mapping the nested Specimen object keys to the flat object keys.
