import { auth, common } from '~/lib'

export default defineEventHandler(async (event) => {
  const bootstrapRequest = new common.generic.BootstrapRequest({
    collections: [
      new auth.objects.Collection({
        authType: auth.auth.AuthType.ROLE,
        collectionId: 'TestCollection',
        default: {},
        itemTypes: ['sample'],
        name: 'TestCollection',
      }),
    ],
  })

  const cc = await useChaincode(event)
  const bootstrap = await cc.service.bootstrap(bootstrapRequest)
  console.log(bootstrap)
  return bootstrap
})
