import { Generator } from '@runtyping/runtypes'

const generator = new Generator({
  targetFile: './src/runtypes.ts',
  // optional: runtypeFormat / typeFormat (see above)
})

generator
  .generate([
    { file: './src/types.ts', type: 'Foo' },
    { file: 'json/my-json-schema.json', type: 'ExampleType' },
  ])
  .then((file) => file.save())
