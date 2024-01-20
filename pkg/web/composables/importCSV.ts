import type { ParseResult } from 'papaparse'
import Papa from 'papaparse'

export function useImportCSV(file: Ref<File | any>) {
  const data = ref<ParseResult<Record<string, string>>>()
  const headers = ref<string[]>([])

  if (!file || !file.value)
    return

  Papa.parse(toValue(file), {
    // worker: true,
    header: true,
    complete: (results: ParseResult<Record<string, string>>) => {
      console.log(results)
    },
  })

  return ref(data)
}
