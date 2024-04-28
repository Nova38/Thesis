import { expect, it } from 'vitest'
import certToID from './certToID'

it('certTo', () => {
  expect(certToID('017e3825-1f1e-5503-93d3-a0e385683b15')).toEqual(
    'orn:017e3825-1f1e-5503-93d3-a0e385683b15',
  )
})
