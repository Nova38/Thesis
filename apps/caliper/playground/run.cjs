const jiti = require('jiti')(__filename, { esmResolve: true })

const lib = jiti('./index.ts')
lib.test()
const x = jiti('@saacs/saacs-pb').pb.Book.fromJson({})

console.log(x.toJsonString({}))
