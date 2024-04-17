import { crush } from 'radash'
import { beforeEach, describe, expect, it } from 'vitest'

describe('flatten', () => {
  it('should flatten object', () => {
    const s = MakeEmptySpecimen()
    s.collectionId = 'collectionId'
    console.log('s', s)
    const SpecimenKeys = crush(s)

    console.log('SpecimenKeys', SpecimenKeys)
  })
})
//
