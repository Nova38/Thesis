import { expect, it } from 'vitest'
import { ccbio } from '~/lib'

it('foo', () => {
  console.log('foo')

  const a = new ccbio.Specimen({
    collectionId: 'foo',
    georeference: {},
    grants: {},
    images: {},
    loans: {},
    primary: {
      catalogNumber: 'foo',
    },
    secondary: {
      preparations: {},
    },
    specimenId: 'bar',
    taxon: {},
  })

  const b = a.clone()

  if (!b.taxon)
    return
  if (!b.primary)
    return
  b.taxon.family = 'bar'
  b.primary.catalogNumber = 'bar'

  const fm = diffToFieldMaskPath(a, b)
  console.log(fm.mask)
  expect(1 + 1).toBe(2)
  expect(2 + 4).toBe(6)
})
