import { describe, expect, it } from 'vitest'
import { UniqueFields } from './uniqueFields'

describe('uniqueFields', () => {
  it('should return a map of unique fields', () => {
    const arr = [
      { id: 1, name: 'John' },
      { id: 2, name: 'Jane' },
      { id: 3, name: 'John' },
      { id: 4, name: 'Jane' },
    ]

    const result = UniqueFields(arr, ['id', 'name'])
    console.log(result)

    expect(result).toMatchObject({ id: [1, 2, 3, 4] })
    expect(result.id).toEqual([1, 2, 3, 4])
  })
})
