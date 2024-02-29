// @vitest-environment nuxt
import { beforeEach, describe, expect, it } from 'vitest'
import { createPinia, setActivePinia } from 'pinia'

describe('pinia bulk update', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  console.log('test')
  it('should update all todos', () => {
    const bulk = useBulkUpdate()

    bulk.CollectionId = '123'

    expect(bulk.CollectionId).toBe('123')
  })
})
