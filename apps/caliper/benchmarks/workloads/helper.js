'use strict'
const { ConfigUtil } = require('@hyperledger/caliper-core')

const configLogger =
  require('@hyperledger/caliper-core').CaliperUtils.getLogger('saacs-config')
const runtimeLogger =
  require('@hyperledger/caliper-core').CaliperUtils.getLogger('saacs-runtime')

module.exports.saacsConfig = ConfigUtil.get('saacs')

function GetCollectionId() {
  let collectionId = ConfigUtil.get('saacs-collectionId')
  if (!collectionId) {
    const timestamp = Date.now()

    collectionId = `saacs-collection-${timestamp}`

    ConfigUtil.set('saacs-collectionId', collectionId)

    configLogger.info(
      'Collection ID not found in the configuration, setting new collection ID value',
      'collectionId:',
      collectionId,
    )
  } else {
    configLogger.info(
      'Collection ID found in the configuration, using collection ID value',
      'collectionId:',
      collectionId,
    )
  }

  return collectionId
}

function GetConfig() {
  return ConfigUtil.get('saacs')
}


function SetUserIds(users = {}){

  configLogger.info("SetUserIds", users)

  return ConfigUtil.set('saacs-users', users)
}
function GetUserIds(){
  const users = ConfigUtil.get('saacs-users')

  configLogger.info("GetUserIds", users)

  return users
}


module.exports = {
  logger: runtimeLogger,
  GetCollectionId,
  GetConfig,
  ConfigUtil,
  SetUserIds,
  GetUserIds,
}
// const keys = ConfigUtil.keys
// // logger.info('this', this)
// logger.info("mymodule", ConfigUtil.get('mymodule'));
// logger.info("mymodules", ConfigUtil.get('mymodules'));
// logger.info("mymodules.other", );
// logger.info(JSON.stringify(keys))
// const s = ConfigUtil.get('mymodule-performance-shoudbefast')
