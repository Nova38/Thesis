import { describe, expect, it } from 'vitest'
import type { FieldMapping } from './Mapping'

describe('fieldMapping', () => {
  it('should have newKey property of type keyof N', () => {
    const mapping: FieldMapping<{ foo: string }, { bar: number }> = {
      newKey: 'foo',
      oldKey: 'bar',
    }
    expect(mapping.newKey).toBe('foo')
  })

  it('should have oldKey property of type keyof O', () => {
    const mapping: FieldMapping<{ foo: string }, { bar: number }> = {
      newKey: 'foo',
      oldKey: 'bar',
    }
    expect(mapping.oldKey).toBe('bar')
  })

  it('should have defaultValue property of type (() => N[keyof N]) | N[keyof N]', () => {
    const mapping1: FieldMapping<{ foo: string }, { bar: number }> = {
      newKey: 'foo',
      oldKey: 'bar',
      defaultValue: 'default',
    }
    expect(mapping1.defaultValue).toBe('default')

    const mapping2: FieldMapping<{ foo: string }, { bar: number }> = {
      newKey: 'foo',
      oldKey: 'bar',
      defaultValue: () => 'default',
    }
    expect(mapping2.defaultValue).toBeInstanceOf(Function)
    expect(mapping2.defaultValue()).toBe('default')
  })

  it('should have transform property of type (value: O[keyof O]) => N[keyof N]', () => {
    const mapping: FieldMapping<{ foo: string }, { bar: number }> = {
      newKey: 'foo',
      oldKey: 'bar',
      transform: value => String(value),
    }
    expect(mapping.transform).toBeInstanceOf(Function)
    expect(mapping.transform(123)).toBe('123')
  })
})
