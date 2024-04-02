import { defineBuildConfig } from "unbuild";

export default defineBuildConfig({
    // If entries is not provided, will be automatically inferred from package.json
    entries: [
        // default
        "./src/index",


        // mkdist builder transpiles file-to-file keeping original sources structure
        {
            builder: "mkdist",
            input: "./src/gen/",
            outDir: "./dist/gen",
            format: "cjs",
        },
        {
            builder: "mkdist",
            input: "./src/gen/",
            outDir: "./dist/gen",
            format: "esm",
        },
    ],

    // Change outDir, default is 'dist'

    // Generates .d.ts declaration file
    declaration: true,

    rollup: {
        emitCJS: true
    }


});
