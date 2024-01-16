import type { ParseResult } from "papaparse";
import Papa from "papaparse";
import { crush, get, set, keys } from "radash";
import { ccbio } from "saacs-es";
import { randomUUID } from "uncrypto";

export const useImportCSV = (file: Ref<File | any>) => {
  const data = ref<ParseResult<Record<string, string>>>();
  const headers = ref<string[]>([]);

  if (!file || !file.value) {
    return;
  }

  Papa.parse(toValue(file), {
    // worker: true,
    header: true,
    complete: (results: ParseResult<Record<string, string>>) => {
      console.log(results);
    },
  });

  return ref(data);
};
