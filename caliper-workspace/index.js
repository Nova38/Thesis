'use strict'
const l = require("es")    

const x = new l.utils.createAuthorObject()



console.log(x.toJson({typeRegistry: l.utils.registry}))
