import { beforeEach, describe, it } from 'vitest'
import { createPinia, setActivePinia } from 'pinia'

describe('pinia bulk update', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
  })

  console.log('test')
  it('should update all todos', () => {

  })
})
