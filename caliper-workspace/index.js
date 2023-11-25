'use strict'
const l = require("es")    

const x = l.utils.factory.createAuthorItem()


console.log(x.toJson({typeRegistry: l.utils.reg.Registry}))
