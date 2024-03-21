/**
 * Returns a Map containing unique values for each field in the array of objects.
 * If headers are provided, only those fields will be considered for uniqueness.
 * If headers are not provided, the keys from the first object in the array will be used as headers.
 *
 * @param arr - The array of objects to find unique values from.
 * @param headers - Optional. The fields to consider for uniqueness.
 * @returns A Map containing unique values for each field that is unique.
 */
export function UniqueFields<T extends object>(
  arr: T[],
  headers?: (keyof T)[],
) {
  // If headers is not provided, get the keys from the first object in the array
  const keys = headers ?? (Object.keys(arr[0]) as (keyof T)[])
  console.log(keys)
  const uniqueField = new Map<keyof T, Set<string>>()
  keys.forEach((header: keyof T) => {
    uniqueField.set(header, new Set())
  })

  arr.forEach((row) => {
    keys.forEach((header: keyof T) => {
      const key = uniqueField.get(header)
      if (!key) return

      // if the value is not in the set, add it to the set. Else remove the
      // header from the map

      switch (key?.has(row[header] as string)) {
        case true:
          uniqueField.delete(header)
          break
        case false:
          key.add(row[header] as string)
          break
      }
    })
  })

  // Reduce the uniqueField map into a object with only the unique fields and their values
  const fields = Object.fromEntries(
    [...uniqueField.entries()].map(([k, v]) => [k, Array.from(v)]),
  )
  return fields
}
