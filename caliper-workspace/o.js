const esbuild = require('esbuild');

esbuild
    .build({
        entryPoints: ['lib/gen/'],
        outdir: 'libs',
        bundle: true,
        sourcemap: true,
        minify: true,
        platform: 'node',
        target: ['node10.4'],
    })
    .catch(() => process.exit(1));
