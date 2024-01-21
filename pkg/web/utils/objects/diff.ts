import { FieldMask } from '@bufbuild/protobuf'
import { diff } from 'ohash'
import { crush } from 'radash'
import { snakeCase } from 'scule'

export type diffOp = 'added' | 'removed' | 'updated'
export function toMask(p: string) {
  return p
    .split('.')
    .map(i => snakeCase(i))
    .join('.')
}

export function diffCrush(base: any, updated: any, excludePaths: string[]) {
  const b: Record<string, any> = crush(toValue(base) ?? {})
  const u: Record<string, any> = crush(toValue(updated) ?? {})

  const differences: Record<string, any>[] = []
  const paths: string[] = []
  for (const key in b) {
    if (excludePaths.includes(key))
      continue
    if (b[key] !== u[key]) {
      differences.push({
        key,
        type: 'updated',
        values: { base: b[key], updated: u[key] },
      })
      paths.push(toMask(key))
    }
    // if (!(key in u) && b[key] !== undefined) {
    //   differences.push({
    //     key,
    //     values: { base: b[key] },
    //     type: "removed",
    //   });
    //   paths.push(toMask(key));
    // }
  }

  for (const key in u) {
    if (excludePaths.includes(key))
      continue

    if (!(key in b) && u[key] !== undefined) {
      differences.push({
        key,
        type: 'added',
        value: { values: { updated: u[key] } },
      })
      paths.push(toMask(key))
    }
  }

  const mask = new FieldMask({ paths })

  return { differences, mask, paths }
}

export function diffToFieldMaskPath(base: any, updated: any) {
  const paths = diff(base, updated, {
    excludeKeys: (key: string) => {
      if (key.startsWith('_'))
        return true
      if (key === 'dep')
        return true

      return false
    },
  }).map(d =>
    d.key
      .split('.')
      .map(i => snakeCase(i))
      .join('.'),
  )

  // const t = JSON.stringify({
  //   paths,
  // });

  const mask = new FieldMask({ paths })

  // const mask = FieldMask.fromJson();

  return { mask }
}

// export const SpecimenToForm = (specimen: any) => {
//   let data: any = set(
//     specimen,
//     "primary.DeterminedDateTimestamp",
//     specimen?.primary?.determinedDate?.timestamp
//       ?.toDate()
//       .toLocaleDateString("en-US"),
//   );
//   data = set(
//     data,
//     "primary.OriginalDateTimestamp",
//     specimen?.primary?.originalDate?.timestamp
//       ?.toDate()
//       .toLocaleDateString("en-US"),
//   );
//   data = set(
//     data,
//     "primary.FieldDateTimestamp",
//     specimen?.primary?.fieldDate?.timestamp
//       ?.toDate()
//       .toLocaleDateString("en-US"),
//   );
//   data = set(
//     data,
//     "primary.FieldDateTimestamp",
//     specimen?.primary?.fieldDate?.timestamp
//       ?.toDate()
//       .toLocaleDateString("en-US"),
//   );
//   return data;
// };
// export const FormToSpecimen = (specimen: ccbio.Specimen) => {
//   const DeterminedDateTimestamp = get<string>(
//     specimen,
//     "primary.DeterminedDateTimestamp",
//   );
//   if (DeterminedDateTimestamp && specimen.primary) {
//     const date = new Date(DeterminedDateTimestamp);
//     if (!specimen.primary.determinedDate)
//       specimen.primary.determinedDate = new ccbio.Date();
//     specimen.primary.determinedDate.timestamp = Timestamp.fromDate(date);
//   }

//   const OriginalDateTimestamp = get<string>(
//     specimen,
//     "primary.OriginalDateTimestamp",
//   );
//   if (OriginalDateTimestamp && specimen.primary) {
//     const date = new Date(OriginalDateTimestamp);
//     if (!specimen.primary.determinedDate)
//       specimen.primary.determinedDate = new ccbio.Date();
//     specimen.primary.determinedDate.timestamp = Timestamp.fromDate(date);
//   }

//   const FieldDateTimestamp = get<string>(
//     specimen,
//     "primary.FieldDateTimestamp",
//   );
//   if (FieldDateTimestamp && specimen.primary) {
//     const date = new Date(FieldDateTimestamp);
//     if (!specimen.primary.determinedDate)
//       specimen.primary.determinedDate = new ccbio.Date();
//     specimen.primary.determinedDate.timestamp = Timestamp.fromDate(date);
//   }

//   return specimen;
// };
